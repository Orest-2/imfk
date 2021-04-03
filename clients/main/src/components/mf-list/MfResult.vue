<template>
  <div
    v-if="result.length"
    class="w-full text-center"
  >
    <div class="text-3xl mb-3 font-mono">
      Результат
    </div>
    <div class="text-xl">
      {{ `${resf.join(' | ')}` }}
    </div>
  </div>
</template>

<script>
import { computed } from 'vue'
export default {
  props: {
    result: {
      type: Array,
      default: () => []
    }
  },

  setup (props) {
    const resf = computed(() => props.result.map(e => e.toFixed(4)))

    return {
      resf
    }
  }
}
</script>
