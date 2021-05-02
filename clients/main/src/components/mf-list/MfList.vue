<template>
  <div class="divide-y-2 divide-true-gray-500 mb-2rem">
    <div
      v-for="([k, v]) in mfs"
      :key="k"
    >
      <mf-card
        v-if="!v.operation"
        :mfid="k"
        :show-del-btn="mfs.length > 1"
        @remove="removeMf(k)"
      />
      <mf-operation-card
        v-if="v.operation"
        :mfid="k"
        :show-del-btn="mfs.length > 1"
        @remove="removeMf(k)"
      />
    </div>

    <div class="flex py-15px">
      <button
        class="mx-15px px-15px py-5px border-3 border-gray-500 rounded w-full font-mono hover:bg-true-gray-100 bg-true-gray-200"
        @click="addMf"
      >
        + Добавити функцію належності
      </button>

      <button
        v-if="type === '2d'"
        class="mx-15px px-15px py-5px border-3 border-gray-500 rounded w-full font-mono hover:bg-true-gray-100 bg-true-gray-200"
        @click="addMfOperation"
      >
        + Добавити операцію
      </button>
    </div>
  </div>
</template>

<script>
import { computed } from 'vue'
import { useStore } from 'vuex'

import MfCard from './MfCard.vue'
import MfOperationCard from './MfOperationCard.vue'

export default {
  components: { MfCard, MfOperationCard },

  setup () {
    const store = useStore()

    const type = computed(() => store.state.general.type)
    const mfs = computed(() => store.getters['general/getSelectedMfsByType']())

    const addMf = () => { store.dispatch('general/addMf') }
    const addMfOperation = () => { store.dispatch('general/addMfOperation') }
    const removeMf = (k) => { store.dispatch('general/removeMf', k) }

    return {
      type,
      mfs,
      addMf,
      addMfOperation,
      removeMf
    }
  }
}
</script>
