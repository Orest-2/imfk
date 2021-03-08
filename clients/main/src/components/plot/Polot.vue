<template>
  <div
    ref="plotEl"
    class="lg:w-1/2 md:h-150 h-100 w-full"
  />
</template>

<script>
import { computed, onMounted, ref, watch } from 'vue'

export default {
  props: {
    traces: {
      type: Array,
      default: () => []
    },
    type: {
      type: String,
      default: '2d'
    }
  },

  setup (props) {
    const plotEl = ref(null)

    const data = computed(() => {
      return props.traces.map(t => {
        const trace = {
          x: t.x,
          y: t.y
        }

        if (props.type === '3d') {
          trace.z = t.z || []
          trace.type = 'surface'
        } else {
          trace.line = {
            color: 'red',
            width: 2
          }
        }

        return trace
      })
    })

    const layout = {
      dragmode: 'pan'
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
