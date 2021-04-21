<template>
  <div class="divide-y-2 divide-true-gray-500 mb-2rem">
    <mf-box
      v-for="([k]) in mfs"
      :key="k"
      :mfid="k"
      :show-del-btn="mfs.length > 1"
      @remove="removeMf(k)"
    />

    <div class="flex p-15px">
      <button
        class="mx-auto px-15px py-5px border-3 border-gray-500 rounded w-full font-mono hover:bg-true-gray-100 bg-true-gray-200"
        @click="addMf"
      >
        + Добавити нову функцію належності
      </button>
    </div>
  </div>
</template>

<script>
import { computed } from 'vue'
import { useStore } from 'vuex'

import MfBox from './MfCard.vue'

export default {
  components: { MfBox },

  setup () {
    const store = useStore()

    const mfs = computed(() => store.getters['general/getSelectedMfsByType']())

    const addMf = () => { store.dispatch('general/addMf') }
    const removeMf = (k) => { store.dispatch('general/removeMf', k) }

    return {
      mfs,
      addMf,
      removeMf
    }
  }
}
</script>
