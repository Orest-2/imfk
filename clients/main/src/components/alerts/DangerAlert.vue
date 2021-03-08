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
import { ref } from 'vue'

export default {
  setup () {
    const errorList = ref([])

    const clearErrors = () => {
      errorList.value = []
    }

    const responseErrorHandler = (data) => {
      if (data?.response) {
        const { error } = data?.response?.data || {}

        clearErrors()

        errorList.value.push({ detail: error })
      }
    }

    return {
      errorList,
      responseErrorHandler,
      clearErrors
    }
  }
}
</script>
