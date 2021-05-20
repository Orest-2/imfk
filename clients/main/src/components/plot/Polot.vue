<template>
  <div
    ref="plotEl"
    class="lg:w-1/2 md:h-150 h-100 w-full"
  />
</template>

<script>
import { computed, onMounted, ref, watch } from 'vue'
import { useStore } from 'vuex'
import { useI18n } from 'vue-i18n'

export default {
  props: {
    mfid: {
      type: String,
      default: ''
    }
  },

  setup (props) {
    const { t: $t } = useI18n({ useScope: 'global' })
    const store = useStore()

    const plotEl = ref(null)

    const type = computed(() => store.state.general.type)
    const selectedMfData = computed(() => store.getters['general/getSelectedMfDataByKeyAndType'](props.mfid))
    const plotTraces = computed(() => selectedMfData.value?.plotTraces || [])

    const data = computed(() => {
      return plotTraces.value.map(t => {
        const trace = {
          x: t.x,
          y: t.y
        }

        if (t.name) {
          trace.name = $t(`general.${t.name}`)
        }

        if (type.value === '3d') {
          trace.z = t.z || []
          trace.type = 'surface'
        } else {
          trace.line = {
            color: t.color || 'red',
            width: 2
          }
          if (t.dash) {
            trace.line.dash = t.dash
          }
        }

        return trace
      })
    })

    const layout = {
      dragmode: type.value === '2d' ? 'pan' : 'turntable'
    }

    const config = {
      scrollZoom: true,
      responsive: true
    }

    onMounted(() => {
      Plotly.newPlot(
        plotEl.value,
        data.value,
        layout,
        config
      )
    })

    watch(
      data,
      () => {
        Plotly.newPlot(
          plotEl.value,
          data.value,
          layout,
          config
        )
      }
    )

    return {
      plotEl
    }
  }
}
</script>
