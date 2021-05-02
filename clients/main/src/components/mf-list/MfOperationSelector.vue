<template>
  <fieldset class="w-full">
    <legend class="font-mono mb-1">
      Виберіть операцію
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

export default {
  props: {
    mfid: {
      type: String,
      default: ''
    }
  },

  setup (props) {
    const store = useStore()

    const operations = [
      {
        value: 'intersection',
        text: 'Перетин'
      },
      {
        value: 'association',
        text: 'Об\'єднання'
      },
      {
        value: 'difference',
        text: 'Різниця'
      },
      {
        value: 'symmetrical_difference',
        text: 'Симетрична різниця'
      },
      {
        value: 'disjunctive_sum',
        text: 'Диз\'юнктивна сума'
      }
    ]

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
