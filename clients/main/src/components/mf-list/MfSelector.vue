<template>
  <fieldset class="w-full">
    <legend class="font-mono mb-1">
      {{ $t('general.select_the_membership_function') }}
    </legend>
    <select
      v-model="selectedMf"
      class="border-3 border-gray-500 rounded w-full p-1.2"
    >
      <option
        :value="null"
        selected
      >
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

    const membershipFuncs = computed(() => store.getters['settings/getMembershipFuncsByType'](store.state.general.type))
    const selectedMf = computed({
      get () {
        if (props.operand >= 0) {
          return store.getters['general/getOperandByKeyAndOperandIndxAndType'](props.mfid, props.operand)?.mf || null
        }
        return store.getters['general/getSelectedMfByKeyAndType'](props.mfid)
      },
      set (val) {
        if (props.operand >= 0) {
          store.commit('general/setOperandByType', { k: props.mfid, i: props.operand, v: val })
          return
        }
        store.commit('general/setMfByType', { k: props.mfid, v: val })
      }
    })

    return {
      selectedMf,
      membershipFuncs
    }
  }
}
</script>
