<template>
  <div
    v-if="result.length"
    class="w-full text-center"
  >
    <div class="text-3xl mb-3 font-mono">
      Результат
    </div>
    <div class="text-xl">
      {{ `${result.join(' | ')}` }}
    </div>
  </div>
</template>

<script>
export default {
  props: {
    result: {
      type: Array,
      default: () => []
    }
  }
}
</script>
