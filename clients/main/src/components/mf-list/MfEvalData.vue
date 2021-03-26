<template>
  <fieldset class="mb-2">
    <legend class="font-mono mb-1">
      Введіть інформацію, що потрібно формалізувати
    </legend>
    <div class="flex flex-wrap space md:-mx-1">
      <div
        v-for="(_, i) in data"
        :key="i"
        class="flex items-center w-full md:w-1/8 md:mx-1 mb-1"
      >
        <input
          v-model.number="data[i]"
          type="number"
          class="border-3 border-gray-500 rounded w-full p-1"
        >
      </div>
    </div>

    <div class="flex items-center w-full my-1">
      <button
        class="mx-auto px-15px py-5px border-3 border-gray-500 rounded w-full font-mono hover:bg-true-gray-100 bg-true-gray-200"
        @click="data.pop()"
      >
        -
      </button>
      <div class="flex flex-wrap space mx-1 md:w-1/1 text-center">
        <div
          class="flex items-center w-full"
        >
          <input
            v-model.number="count"
            type="number"
            class="border-3 border-gray-500 rounded w-full p-1"
          >
        </div>
      </div>
      <button
        class="mx-auto px-15px py-5px border-3 border-gray-500 rounded w-full font-mono hover:bg-true-gray-100 bg-true-gray-200"
        @click="data.push(0)"
      >
        +
      </button>
    </div>

    <div class="flex py-2">
      <button
        class="mx-auto px-15px py-5px border-3 border-gray-500 rounded w-full font-mono hover:bg-true-gray-100 bg-true-gray-200"
        @click="$emit('eval')"
      >
        Виконати
      </button>
    </div>
  </fieldset>
</template>

<script>
import { computed, onMounted } from 'vue'
import { useModelWrapper } from '../../utils/modelWrapper'
import { qs } from '../../utils/query'

export default {
  props: {
    modelValue: {
      type: Array,
      default: () => []
    }
  },

  emits: ['update:modelValue', 'eval'],

  setup (props, { emit }) {
    const data = useModelWrapper(props, emit)

    const count = computed({
      get () {
        return data.value.length
      },
      set (cnt) {
        const l = data.value.length

        if (l < cnt) {
          data.value.push(...Array(cnt - l).fill(0))
        }
        if (l > cnt) {
          data.value = data.value.slice(0, cnt)
        }
      }
    })

    onMounted(() => {
      if (qs.params.get('test') === '1') {
        data.value = [0.415, 0.350, 0.613, 0.283, 0.927]
      } else {
        data.value = [0]
      }
    })

    return {
      data,
      count
    }
  }
}
</script>
