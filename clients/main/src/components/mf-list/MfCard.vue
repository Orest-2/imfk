<template>
  <div class="py-1rem">
    <div
      class="-lg:px-1rem flex space-x-2"
      :class="{
        'mb-1rem': selectedMf
      }"
    >
      <mf-selector :mfid="mfid" />

      <div
        v-if="showDelBtn"
        class="md:w-1/10 self-end"
      >
        <button
          class="mx-auto px-15px py-5px border-3 border-red-500 rounded w-full font-mono hover:bg-red-300 bg-red-200"
          @click="$emit('remove')"
        >
          X
        </button>
      </div>
    </div>

    <div
      v-if="selectedMf"
      class="flex flex-wrap lg:flex-row flex-col-reverse"
    >
      <polot :mfid="mfid" />

      <div class="lg:w-1/2 pl-1rem -lg:px-1rem sm:w-full">
        <danger-alert class="mb-2" />

        <mf-params
          v-if="type==='2d'"
          :mfid="mfid"
        />

        <mf-3-d-params
          v-if="type==='3d'"
          :mfid="mfid"
        />

        <mf-eval-data
          v-if="type==='2d'"
          :mfid="mfid"
          @eval="evalData"
        />

        <mf-3-d-eval-data
          v-if="type==='3d'"
          :mfid="mfid"
          @eval="evalData"
        />
      </div>

      <mf-result :mfid="mfid" />
    </div>
  </div>
</template>

<script>
import { useStore } from 'vuex'
import { computed, watch } from 'vue'

import { createDebounce } from '../../utils/debounce'

import Polot from '../plot/Polot.vue'
import MfParams from './MfParams.vue'
import MfSelector from './MfSelector.vue'
import DangerAlert from '../alerts/DangerAlert.vue'
import Mf3DParams from './Mf3DParams.vue'
import MfEvalData from './MfEvalData.vue'
import Mf3DEvalData from './Mf3DEvalData.vue'
import MfResult from './MfResult.vue'

export default {
  components: {
    Polot,
    MfParams,
    MfSelector,
    DangerAlert,
    Mf3DParams,
    MfEvalData,
    Mf3DEvalData,
    MfResult
  },

  props: {
    showDelBtn: Boolean,
    mfid: {
      type: String,
      default: ''
    }
  },

  emits: ['remove'],

  setup (props) {
    const store = useStore()

    const type = computed(() => store.state.general.type)
    const selectedMfData = computed(() => store.getters['general/getSelectedMfDataByKeyAndType'](props.mfid))
    const selectedMf = computed(() => selectedMfData.value?.mf)
    const data = computed(() => selectedMfData.value?.evalData)

    const debounce = createDebounce()

    const evalData = () => {
      const transpose = matrix => matrix.reduce(
        ($, row) => row.map((_, i) => [...($[i] || []), row[i]]),
        []
      )

      const payload = {
        membership_func_id: selectedMf.value.id,
        func_params: selectedMfData.value?.funcParams,
        in_data: type.value === '3d' ? transpose(data.value) : data.value
      }

      store.dispatch('general/evalData', { k: props.mfid, payload })
    }

    watch(
      [
        () => selectedMfData.value?.funcParams,
        () => selectedMfData.value?.plotParams
      ],
      ([funcParams, plotParams]) => {
        if (funcParams && plotParams) {
          const payload = {
            membership_func_id: selectedMf.value.id,
            func_params: funcParams,
            plot_params: plotParams
          }

          debounce(
            () => store.dispatch('general/makePlot', { k: props.mfid, payload }),
            1000
          )
        }
      },
      { deep: true }
    )

    return {
      type,
      selectedMf,
      evalData
    }
  }
}
</script>
