<template>
  <div
    v-if="resf && resf.length"
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
import { computed, watch } from 'vue'
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

    const selectedMfData = computed(() => store.getters['general/getSelectedMfDataByKeyAndType'](props.mfid))
    const resf = computed(() => selectedMfData.value?.evalDataResult.map(e => e.toFixed(4)))

    watch(
      [
        () => selectedMfData.value?.funcParams,
        () => selectedMfData.value?.plotParams,
        () => selectedMfData.value?.mf,
        () => selectedMfData.value?.evalData
      ],
      () => {
        store.commit('general/setMfEvalDataResultByType', { k: props.mfid, v: [] })
      },
      { deep: true }
    )

    return {
      resf
    }
  }
}
</script>
