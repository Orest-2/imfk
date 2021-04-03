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
          v-for="(n, i) in params.length"
          :key="n+i"
          class="flex items-center w-full md:mx-1 mb-1"
        >
          <li class="mr-1" />
          <input
            v-for="(__, j) in params[i]"
            :key="`${i}+${j}`"
            v-model.number="params[i][j]"
            type="number"
            step="1"
            class="border-3 border-gray-500 rounded md:w-1/4 p-1 mx-1"
          >
        </div>
      </div>
    </ul>
  </fieldset>
</template>

<script>
import { onMounted, watch } from 'vue'
import { useModelWrapper } from '../../utils/modelWrapper'

const defaultParams = {
  conemf: [[0.5, 0.5], [0.4, 0.4]],
  pyrammf: [[0.5, 0.5], [0.4, 0.4]],
  trappyrammf: [[0.5, 0.5], [0.2, 1], [0.4, 0.4]],
  gsigmf: [[0.5, 0.5], [0, 10]],
  gbell3dmf: [[0.5, 0.5], [0.2, 0.2], [2.5, 2.5]],
  gauss3dmf: [[0.5, 0.5], [0.25, 0.25]],
  hyperbolmf: [[0.5, 0.5], [0.25, 0.25]],
  ellipsmf: [[0.5, 0.5], [0.4, 0.4]]
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
    const params = useModelWrapper(props, emit)
    const plotParams = useModelWrapper(props, emit, 'modelPlotParams')

    const initparams = () => {
      const dp = defaultParams[props.selectedMf.code]

      if (dp) {
        params.value = [...dp]
      } else {
        params.value = Array(props.selectedMf.params_count).fill(0).map(() => [0, 0])
      }

      plotParams.value = [1, 0.01]
    }

    watch(() => props.selectedMf, initparams)

    onMounted(initparams)

    return {
      params,
      plotParams
    }
  }
}
</script>
