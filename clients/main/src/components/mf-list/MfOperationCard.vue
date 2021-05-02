<template>
  <div class="py-1rem">
    <div
      class="-lg:px-1rem flex space-x-2"
      :class="{
        'mb-1rem': operands.length > 1
      }"
    >
      <mf-selector
        :mfid="mfid"
        :operand="0"
      />

      <mf-operation-selector :mfid="mfid" />

      <mf-selector
        :mfid="mfid"
        :operand="1"
      />

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
      v-if="operands.length > 1"
      class="flex flex-wrap lg:flex-row flex-col-reverse"
    >
      <polot :mfid="mfid" />

      <div class="lg:w-1/2 pl-1rem -lg:px-1rem sm:w-full">
        <danger-alert class="mb-2" />

        <PolotParams :mfid="mfid" />

        <mf-params
          v-if="type==='2d'"
          :mfid="mfid"
          :operand="0"
        />

        <mf-params
          v-if="type==='2d'"
          :mfid="mfid"
          :operand="1"
          hide-plot-params
        />

        <!-- <mf-3-d-params
          v-if="type==='3d'"
          :mfid="mfid"
        /> -->
        <!--
        <mf-eval-data
          v-if="type==='2d'"
          :mfid="mfid"
          @eval="evalData"
        />

        <mf-3-d-eval-data
          v-if="type==='3d'"
          :mfid="mfid"
          @eval="evalData"
        /> -->
      </div>

      <mf-result :mfid="mfid" />
    </div>
  </div>
</template>

<script>
import { useStore } from 'vuex'
import { computed, watch } from 'vue'

import { createDebounce } from '../../utils/debounce'

import MfSelector from './MfSelector.vue'
import MfOperationSelector from './MfOperationSelector.vue'
import Polot from '../plot/Polot.vue'
import DangerAlert from '../alerts/DangerAlert.vue'
import MfParams from './MfParams.vue'
// import Mf3DParams from './Mf3DParams.vue'
// import MfEvalData from './MfEvalData.vue'
// import Mf3DEvalData from './Mf3DEvalData.vue'
import MfResult from './MfResult.vue'
import PolotParams from '../plot/PolotParams.vue'

export default {
  components: {
    MfSelector,
    MfOperationSelector,
    Polot,
    DangerAlert,
    MfParams,
    // Mf3DParams,
    // MfEvalData,
    // Mf3DEvalData,
    MfResult,
    PolotParams
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
    const operands = computed(() => selectedMfData.value?.operands || [])
    // const data = computed(() => selectedMfData.value?.evalData)

    const debounce = createDebounce()

    // const evalData = () => {
    //   const transpose = matrix => matrix.reduce(
    //     ($, row) => row.map((_, i) => [...($[i] || []), row[i]]),
    //     []
    //   )

    //   const payload = {
    //     membership_func_id: selectedMf.value.id,
    //     func_params: selectedMfData.value?.funcParams,
    //     in_data: type.value === '3d' ? transpose(data.value) : data.value
    //   }

    //   store.dispatch('general/evalData', { k: props.mfid, payload })
    // }

    watch(
      [
        () => operands.value,
        () => selectedMfData.value?.operation,
        () => selectedMfData.value?.plotParams
      ],
      ([operands, operation, plotParams]) => {
        if (operands.length > 1 && operation && plotParams) {
          const data = operands.map(operand => ({
            membership_func_id: operand.mf.id,
            func_params: operand.funcParams,
            plot_params: plotParams
          }))

          const payload = { data }

          debounce(
            () => store.dispatch('general/operationMakePlot', { k: props.mfid, payload }),
            1000
          )
        }
      },
      { deep: true }
    )

    return {
      type,
      operands
      // evalData
    }
  }
}
</script>
