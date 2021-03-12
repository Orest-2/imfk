<template>
  <div>
    <div class="lg:container mx-auto min-h-screen bg-white">
      <div class="pt-10px">
        <p class="text-3xl text-center font-mono">
          Функції належності
        </p>
      </div>

      <div class="font-mono flex justify-center divide-x-2 divide-black border-b-2 p-15px border-true-gray-500">
        <div
          class="px-15px"
          role="button"
          :class="{ 'text-true-gray-400 hover:text-black': type !== '2d' }"
          @click="type = '2d'"
        >
          Однієї змінної
        </div>

        <div
          role="button"
          class="px-15px"
          :class="{ 'text-true-gray-400 hover:text-black': type !== '3d' }"
          @click="type = '3d'"
        >
          Багатьох змінних
        </div>
      </div>

      <mf-list
        v-if="type==='2d'"
        type="2d"
      />

      <mf-list
        v-if="type==='3d'"
        type="3d"
      />
    </div>
  </div>
</template>

<script>
import { ref } from 'vue'
import { useStore } from 'vuex'

import MfList from './components/mf-list/MfList.vue'

export default {
  components: { MfList },

  setup () {
    const store = useStore()
    const type = ref('2d')

    store.dispatch('settings/fetchSettings')

    return {
      type
    }
  }
}
</script>
