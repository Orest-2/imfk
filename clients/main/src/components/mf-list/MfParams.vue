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
          v-for="(n, i) in selectedMf.params_count"
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
import { computed } from 'vue'
import { useStore } from 'vuex'

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
    const plotParams = computed(() => selectedMfData.value?.plotParams)
    const params = computed(() => selectedMfData.value?.funcParams)

    return {
      selectedMf,
      plotParams,
      params
    }
  }
}
</script>
