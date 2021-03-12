<template>
  <div class="divide-y-2 divide-true-gray-500 mb-2rem">
    <mf-box
      v-for="(n, i) in mfcnt"
      :key="n"
      :show-del-btn="mfcnt.length > 1"
      :type="type"
      @remove="removeMf(i)"
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
import { ref } from 'vue'

import MfBox from './MfCard.vue'

export default {
  components: { MfBox },

  props: {
    type: {
      type: String,
      default: '2d'
    }
  },

  setup () {
    const mfcnt = ref([0])

    const addMf = () => {
      mfcnt.value.push(mfcnt.value[mfcnt.value.length - 1] + 1)
    }

    const removeMf = (i) => {
      mfcnt.value.splice(i, 1)
    }

    return {
      mfcnt,
      addMf,
      removeMf
    }
  }
}
</script>
