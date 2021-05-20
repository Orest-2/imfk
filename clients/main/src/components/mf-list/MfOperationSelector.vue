<template>
  <fieldset class="w-full">
    <legend class="font-mono mb-1">
      {{ $t('general.select_an_operation') }}
    </legend>
    <select
      v-model="selectedMfOpretion"
      class="border-3 border-gray-500 rounded w-full p-1.2"
    >
      <option
        v-for="o in operations"
        :key="o.value"
        :value="o.value"
      >
        {{ o.text }}
      </option>
    </select>
  </fieldset>
</template>

<script>
import { computed } from 'vue'
import { useStore } from 'vuex'
import { useI18n } from 'vue-i18n'

export default {
  props: {
    mfid: {
      type: String,
      default: ''
    }
  },

  setup (props) {
    const { t } = useI18n({ useScope: 'global' })
    const store = useStore()

    const operations = computed(() => [
      {
        value: 'intersection',
        text: t('operations.intersection')
      },
      {
        value: 'association',
        text: t('operations.association')
      },
      {
        value: 'difference',
        text: t('operations.difference')
      },
      {
        value: 'symmetrical_difference',
        text: t('operations.symmetrical_difference')
      },
      {
        value: 'disjunctive_sum',
        text: t('operations.disjunctive_sum')
      }
    ])

    const selectedMfData = computed(() => store.getters['general/getSelectedMfDataByKeyAndType'](props.mfid))

    const selectedMfOpretion = computed({
      get () {
        return selectedMfData.value?.opretion || 'intersection'
      },
      set (val) {
        store.commit('general/setMfOpretionByType', { k: props.mfid, v: val })
      }
    })

    return {
      selectedMfOpretion,
      operations
    }
  }
}
</script>
