<template>
  <div>
    <div
      v-for="(err, i) in errorList"
      :key="i"
      class="border-3 p-4 rounded border-red-500 bg-red-200"
    >
      {{ err.detail }}
    </div>
  </div>
</template>

<script>
import { computed } from 'vue'
import { useStore } from 'vuex'

export default {
  setup () {
    const store = useStore()

    const errorList = computed(() => store.state.error.list)

    return {
      store,
      errorList
    }
  }
}
</script>
