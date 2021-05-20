<template>
  <fieldset class="mb-2">
    <legend class="font-mono mb-1">
      {{ $t('general.function_parameters') }}
    </legend>
    <ul class="list-lowerlatin list-inside">
      <div class="flex flex-wrap space md:-mx-1">
        <div
          v-for="(n, i) in params.length"
          :key="n+i"
          class="flex items-center w-full md:mx-1 mb-1"
        >
          <span class="min-w-4 mr-1">
            {{ paramsLetters[selectedMf.code][i] }}
          </span>
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
import { computed } from 'vue'
import { useStore } from 'vuex'

import { paramsLetters } from '../../constants/default-mf-params'

export default {
  props: {
    mfid: {
      type: String,
      default: ''
    }
  },

  setup (props) {
    const store = useStore()

    const selectedMfData = computed(() => store.getters['general/getSelectedMfDataByKeyAndType'](props.mfid))
    const selectedMf = computed(() => selectedMfData.value?.mf)
    const params = computed(() => selectedMfData.value?.funcParams)

    return {
      selectedMf,
      paramsLetters,
      params
    }
  }
}
</script>
