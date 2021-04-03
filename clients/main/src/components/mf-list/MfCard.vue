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
      class="flex flex-wrap lg:flex-row flex-col-reverse"
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
          v-model:modelPlotParams="plotParams"
          :selected-mf="selectedMf"
        />

        <mf-3-d-params
          v-if="type==='3d'"
          v-model="params"
          v-model:modelPlotParams="plotParams"
          :selected-mf="selectedMf"
        />

        <mf-eval-data
          v-if="type==='2d'"
          v-model="data"
          @eval="evalData"
        />

        <mf-3-d-eval-data
          v-if="type==='3d'"
          v-model="data"
          @eval="evalData"
        />
      </div>

      <mf-result :result="result" />
    </div>
  </div>
</template>

<script>
import { nextTick, ref, watch } from 'vue'
import axios from 'axios'
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
  components: { Polot, MfParams, MfSelector, DangerAlert, Mf3DParams, MfEvalData, Mf3DEvalData, MfResult },

  props: {
    showDelBtn: Boolean,
    type: {
      type: String,
      default: '2d'
    }
  },

  emits: ['remove'],

  setup (props) {
    const alerts = ref(null)
    const selectedMf = ref(null)
    const params = ref([])
    const result = ref([])
    const data = ref([0])
    const plotParams = ref([])
    const plotTraces = ref([{ x: [], y: [] }])

    const debounce = createDebounce()

    const makePlot = ({ payload }) => {
      axios.post(
        `http://localhost:1447/api/v1/mf/${selectedMf.value.type}/plot`,
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

    const evalData = () => {
      const transpose = matrix => matrix.reduce(
        ($, row) => row.map((_, i) => [...($[i] || []), row[i]]),
        []
      )

      const payload = {
        membership_func_id: selectedMf.value.id,
        func_params: params.value,
        in_data: props.type === '3d' ? transpose(data.value) : data.value
      }

      axios.post(
        `http://localhost:1447/api/v1/mf/${selectedMf.value.type}/eval`,
        payload
      )
        .then(({ data }) => {
          result.value = [...data.data.result]

          alerts?.value?.clearErrors()
        })
        .catch(err => {
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

        result.value = []
      },
      { deep: true }
    )

    const upadteParams = () => {
      if (props.type === '3d') {
        nextTick(() => {
          const cnt = data.value.length || 1

          params.value.forEach((row, i) => {
            const l = row.length || 2

            if (l < cnt) {
              params.value[i] = [...params.value[i], ...Array(cnt - l).fill(0)]
              console.log(params.value[i])
            }
            if (l > cnt) {
              params.value[i] = row.slice(0, cnt)
            }
          })
        })
      }
    }

    watch(
      data,
      upadteParams,
      { deep: true }
    )
    watch(
      selectedMf,
      upadteParams,
      { deep: true }
    )

    return {
      selectedMf,
      params,
      plotParams,
      plotTraces,
      data,
      evalData,
      result,
      alerts
    }
  }
}
</script>
