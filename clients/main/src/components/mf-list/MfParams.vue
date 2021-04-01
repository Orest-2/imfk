<template>
  <fieldset class="mb-2">
    <legend class="font-mono mb-1">
      Параметри графіка
    </legend>
    <div class="flex flex-wrap space md:-mx-1">
      <div class="flex items-center w-full md:w-1/8 md:mx-1 md:mb-0 mb-1">
        <input
          v-model.number="plotParams[0]"
          class="border-3 border-gray-500 rounded w-full p-1"
          type="number"
          step="1"
        >
      </div>

      <div class="flex items-center w-full md:w-1/8 md:mx-1 md:mb-0 mb-1">
        <input
          v-model.number="plotParams[1]"
          class="border-3 border-gray-500 rounded w-full p-1"
          type="number"
          step="0.1"
        >
      </div>
    </div>
  </fieldset>

  <fieldset class="mb-2">
    <legend class="font-mono mb-1">
      Параметри функції
    </legend>
    <ul class="list-lowerlatin list-inside">
      <div class="flex flex-wrap space md:-mx-1">
        <div
          v-for="(n, i) in pm"
          :key="n+i"
          class="flex items-center w-full md:w-1/8 md:mx-1 md:mb-0 mb-1"
        >
          <li class="mr-1" />
          <input
            v-model.number="params[i]"
            type="number"
            step="0.2"
            class="border-3 border-gray-500 rounded w-full p-1"
          >
        </div>
      </div>
    </ul>
  </fieldset>
</template>

<script>
import { nextTick, onMounted, ref, watch } from 'vue'
import { useModelWrapper } from '../../utils/modelWrapper'

const defaultParams = {
  trimf: [2 / 10, 5 / 10, 8 / 10],
  trapmf: [2 / 10, 4 / 10, 6 / 10, 7.5 / 10],
  szmf: [2.5 / 10, 7.5 / 10],
  hzmf: [2.5 / 10, 7.5 / 10],
  zsigmf: [-20 * 10, 5 / 10],
  ssigmf: [20, 5 / 10],
  zlinemf: [2.5 / 10, 7.5 / 10],
  slinemf: [2.5 / 10, 7.5 / 10],
  hsmf: [2.5 / 10, 7.5 / 10],
  ssmf: [2.5 / 10, 7.5 / 10],
  gbellmf: [2 / 10, 4, 5 / 10],
  gaussmf: [1 / 10, 5 / 10]
}

export default {
  props: {
    modelValue: {
      type: Array,
      default: () => []
    },
    modelPlotParams: {
      type: Array,
      default: () => []
    },
    selectedMf: {
      type: Object,
      required: true,
      default: () => null
    }
  },

  emits: ['update:modelValue', 'update:modelPlotParams'],

  setup (props, { emit }) {
    const pm = ref(0)
    const params = useModelWrapper(props, emit)
    const plotParams = useModelWrapper(props, emit, 'modelPlotParams')

    const initparams = () => {
      const dp = defaultParams[props.selectedMf.code]

      if (dp) {
        params.value = [...dp]
      } else {
        params.value = Array(props.selectedMf.params_count).fill(0)
      }

      plotParams.value = [10 / 10, 0.1 / 10]

      nextTick(() => {
        pm.value = props.selectedMf.params_count
      })
    }

    watch(() => props.selectedMf, initparams)

    onMounted(initparams)

    return {
      params,
      plotParams,
      pm
    }
  }
}
</script>
