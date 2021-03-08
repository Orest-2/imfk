<template>
  <div class="py-1rem">
    <div
      class="-lg:px-1rem flex space-x-2"
      :class="{
        'mb-1rem': selectedMf
      }"
    >
      <mf-selector
        v-model="selectedMf"
        :type="type"
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
      v-if="selectedMf"
      class="flex lg:flex-row flex-col-reverse"
    >
      <polot
        :traces="plotTraces"
        :type="type"
      />

      <div class="lg:w-1/2 pl-1rem -lg:px-1rem sm:w-full">
        <danger-alert
          ref="alerts"
          class="mb-2"
        />

        <mf-params
          v-if="type==='2d'"
          v-model="params"
          :selected-mf="selectedMf"
        />

        <mf-3-d-params
          v-if="type==='3d'"
          v-model="params"
          v-model:modelPlotParams="plotParams"
          :selected-mf="selectedMf"
        />

        <fieldset class="mb-2">
          <legend class="font-mono mb-1">
            Введіть інформацію, що потрібно формалізувати
          </legend>
          <div class="flex flex-wrap space md:-mx-1">
            <div class="flex items-center w-full md:w-1/8 md:mx-1 md:mb-0 mb-1">
              <input
                type="text"
                class="border-3 border-gray-500 rounded w-full p-1"
              >
            </div>

            <div class="flex items-center w-full md:w-1/8 md:mx-1 md:mb-0 mb-1">
              <input
                type="text"
                class="border-3 border-gray-500 rounded w-full p-1"
              >
            </div>

            <div class="flex items-center w-full md:w-1/8 md:mx-1 md:mb-0 mb-1">
              <input
                type="text"
                class="border-3 border-gray-500 rounded w-full p-1"
              >
            </div>
          </div>
        </fieldset>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, watch } from 'vue'
import axios from 'axios'
import { createDebounce } from '../../utils/debounce'
import Polot from '../plot/Polot.vue'
import MfParams from './MfParams.vue'
import MfSelector from './MfSelector.vue'
import DangerAlert from '../alerts/DangerAlert.vue'
import Mf3DParams from './Mf3DParams.vue'

export default {
  components: { Polot, MfParams, MfSelector, DangerAlert, Mf3DParams },

  props: {
    showDelBtn: Boolean,
    type: {
      type: String,
      default: '2d'
    }
  },

  emits: ['remove'],

  setup () {
    const alerts = ref(null)
    const selectedMf = ref(null)
    const params = ref([])
    const plotParams = ref([])
    const plotTraces = ref([{ x: [], y: [] }])

    const debounce = createDebounce()

    const makePlot = ({ payload }) => {
      axios.post(
        `http://localhost:1447/api/v1/mf/plot/${selectedMf.value.type}`,
        payload
      )
        .then(({ data }) => {
          plotTraces.value = [data.data]

          alerts?.value?.clearErrors()
        })
        .catch(err => {
          plotTraces.value = [{ x: [], y: [] }]
          alerts?.value?.responseErrorHandler(err)
        })
    }

    watch(
      [params, plotParams],
      () => {
        const payload = {
          membership_func_id: selectedMf.value.id,
          func_params: params.value,
          plot_params: plotParams.value
        }
        debounce(() => makePlot({ payload }), 1000)
      },
      { deep: true }
    )

    return {
      selectedMf,
      params,
      plotParams,
      plotTraces,
      alerts
    }
  }
}
</script>
