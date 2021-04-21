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

export default {
  props: {
    mfid: {
      type: String,
      default: ''
    }
  },

  setup (props) {
    const store = useStore()

    const selectedMf = computed({
      get () {
        return store.getters['general/getSelectedMfByKeyAndType'](props.mfid)
      },
      set (val) {
        store.commit('general/setMfByType', { k: props.mfid, v: val })
      }
    })
    const membershipFuncs = computed(() => store.getters['settings/getMembershipFuncsByType'](store.state.general.type))

    return {
      selectedMf,
      membershipFuncs
    }
  }
}
</script>
