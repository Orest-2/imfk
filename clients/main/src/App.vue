<template>
  <div>
    <div class="lg:container mx-auto min-h-screen bg-white">
      <div class="pt-10px">
        <div class=" text-center">
          <span class="text-3xl font-mono">
            {{ $t('general.membership_functions') }}
          </span>
          <select
            v-model="$i18n.locale"
            class="float-right"
          >
            <option value="en">
              EN
            </option>
            <option value="ua">
              UA
            </option>
            <option value="hu">
              HU
            </option>
          </select>
        </div>
      </div>

      <div class="font-mono flex justify-center divide-x-2 divide-black border-b-2 p-15px border-true-gray-500">
        <div
          class="px-15px"
          role="button"
          :class="{ 'text-true-gray-400 hover:text-black': type !== '2d' }"
          @click="type = '2d'"
        >
          {{ $t('general.one_variable') }}
        </div>

        <div
          role="button"
          class="px-15px"
          :class="{ 'text-true-gray-400 hover:text-black': type !== '3d' }"
          @click="type = '3d'"
        >
          {{ $t('general.many_variables') }}
        </div>
      </div>

      <mf-list />
    </div>
  </div>
</template>

<script>
import { computed } from 'vue'
import { useStore } from 'vuex'

import MfList from './components/mf-list/MfList.vue'

export default {
  components: { MfList },

  setup () {
    const store = useStore()

    const type = computed({
      get () {
        return store.state.general.type
      },
      set (val) {
        store.commit('general/setType', val)
      }
    })

    store.dispatch('settings/fetchSettings')

    return {
      type
    }
  }
}
</script>
