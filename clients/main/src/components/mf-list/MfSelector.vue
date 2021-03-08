<template>
  <fieldset class="w-full">
    <legend class="font-mono mb-1">
      Виберіть функцію належності
    </legend>
    <select
      v-model="selectedMf"
      class="border-3 border-gray-500 rounded w-full p-1.2"
    >
      <option :value="null">
        Нічого не вибрано
      </option>
      <option
        v-for="mf in membershipFuncs"
        :key="mf.id"
        :value="mf"
      >
        {{ mf.name }}
      </option>
    </select>
  </fieldset>
</template>

<script>
import { computed } from 'vue'
import { useStore } from 'vuex'
import { useModelWrapper } from '../../utils/modelWrapper'

export default {
  props: {
    modelValue: {
      type: Object,
      default: () => null
    },
    type: {
      type: String,
      default: '2d'
    }
  },

  emits: ['update:modelValue'],

  setup (props, { emit }) {
    const store = useStore()
    const membershipFuncs = computed(() => store.getters['settings/getMembershipFuncsByType'](props.type))

    const selectedMf = useModelWrapper(props, emit)

    return {
      selectedMf,
      membershipFuncs
    }
  }
}
</script>
