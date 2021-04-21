<template>
  <fieldset class="mb-2">
    <legend class="font-mono mb-1">
      Введіть інформацію, що потрібно формалізувати
    </legend>
    <div class="flex flex-wrap space md:-mx-1">
      <div
        v-for="(_, i) in data"
        :key="i"
        class="flex flex-wrap items-center w-full md:mx-1 mb-1"
      >
        <div
          v-for="(__, j) in data[i]"
          :key="`${i}+${j}`"
          class="flex items-center w-full md:w-1/8 md:mx-1 mb-1"
        >
          <input
            v-model.number="data[i][j]"
            type="number"
            class="border-3 border-gray-500 rounded w-full p-1"
          >
        </div>
      </div>
    </div>

    <fieldset class="mb-1">
      <legend class="font-mono mb-1">
        Кількість наборів даних
      </legend>
      <div class="flex items-center w-full mb-1">
        <button
          class="mx-auto px-15px py-5px border-3 border-gray-500 rounded w-full font-mono hover:bg-true-gray-100 bg-true-gray-200"
          @click="countCol--"
        >
          -
        </button>
        <div class="flex flex-wrap space mx-1 md:w-1/1 text-center">
          <div
            class="flex items-center w-full"
          >
            <input
              v-model.number="countCol"
              type="number"
              class="border-3 border-gray-500 rounded w-full p-1"
            >
          </div>
        </div>
        <button
          class="mx-auto px-15px py-5px border-3 border-gray-500 rounded w-full font-mono hover:bg-true-gray-100 bg-true-gray-200"
          @click="countCol++"
        >
          +
        </button>
      </div>
    </fieldset>

    <fieldset class="mb-1">
      <legend class="font-mono mb-1">
        Кількість аргументів
      </legend>
      <div class="flex items-center w-full mb-1">
        <button
          class="mx-auto px-15px py-5px border-3 border-gray-500 rounded w-full font-mono hover:bg-true-gray-100 bg-true-gray-200"
          @click="countRow--"
        >
          -
        </button>
        <div class="flex flex-wrap space mx-1 md:w-1/1 text-center">
          <div
            class="flex items-center w-full"
          >
            <input
              v-model.number="countRow"
              type="number"
              class="border-3 border-gray-500 rounded w-full p-1"
            >
          </div>
        </div>
        <button
          class="mx-auto px-15px py-5px border-3 border-gray-500 rounded w-full font-mono hover:bg-true-gray-100 bg-true-gray-200"
          @click="countRow++"
        >
          +
        </button>
      </div>
    </fieldset>

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
      default: () => [[]]
    }
  },

  emits: ['update:modelValue', 'eval'],

  setup (props, { emit }) {
    const data = useModelWrapper(props, emit)

    const countCol = computed({
      get () {
        return data.value[0]?.length || 2
      },
      set (cnt) {
        data.value.forEach((row, i) => {
          const l = row.length || 2

          if (l < cnt) {
            row.push(...Array(cnt - l).fill(0))
          }
          if (l > cnt) {
            data.value[i] = row.slice(0, cnt)
          }
        })
      }
    })

    const countRow = computed({
      get () {
        return data.value.length || 1
      },
      set (cnt) {
        const l = data.value.length || 1

        if (l < cnt) {
          data.value.push(Array(countCol.value).fill(0))
        }
        if (l > cnt) {
          data.value = data.value.slice(0, cnt)
        }
      }
    })

    onMounted(() => {
      if (qs.params.get('test') === '1') {
        data.value = [
          [0.87, 0.82, 0.60, 0.77, 0.69],
          [0.66, 0.83, 0.71, 0.98, 0.91],
          [0.78, 0.40, 0.54, 0.85, 0.82]
        ]
      } else {
        data.value = [[0, 0], [0, 0]]
      }
    })

    return {
      data,
      countCol,
      countRow
    }
  }
}
</script>
