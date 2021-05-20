<template>
  <fieldset class="mb-2">
    <legend class="font-mono mb-1">
      {{ $t('general.function_parameters') }} {{ operand >= 0 && operand+1 || '' }}
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
    },
    operand: {
      type: Number,
      default: -1
    }
  },

  setup (props) {
    const store = useStore()

    const selectedMfData = computed(() => store.getters['general/getSelectedMfDataByKeyAndType'](props.mfid))
    const operands = computed(() => selectedMfData.value?.operands || [])
    const selectedMf = computed(() => ((operands.value.length && operands.value[props.operand]) || selectedMfData.value)?.mf)
    const params = computed(() => ((operands.value.length && operands.value[props.operand]) || selectedMfData.value)?.funcParams)

    return {
      selectedMf,
      params
    }
  }
}
</script>
