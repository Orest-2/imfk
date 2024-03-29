<template>
  <fieldset class="mb-2">
    <legend class="font-mono mb-1">
      {{ $t('general.enter_the_information_you_want_to_formalize') }}
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
        {{ $t('general.number_of_data_sets') }}
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
        {{ $t('general.number_of_arguments') }}
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
        {{ $t('general.evaluate') }}
      </button>
    </div>
  </fieldset>
</template>

<script>
import { computed, nextTick, onMounted, watch } from 'vue'
import { useStore } from 'vuex'

import { qs } from '../../utils/query'

export default {
  props: {
    mfid: {
      type: String,
      default: ''
    }
  },

  emits: ['eval'],

  setup (props) {
    const store = useStore()

    const type = computed(() => store.state.general.type)
    const selectedMfData = computed(() => store.getters['general/getSelectedMfDataByKeyAndType'](props.mfid))
    const selectedMf = computed(() => selectedMfData.value?.mf)
    const data = computed(() => selectedMfData.value?.evalData)

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
          store.commit(
            'general/setMfEvalDataByType',
            { k: props.mfid, v: data.value.slice(0, cnt) }
          )
        }
      }
    })

    const upadteParams = () => {
      if (type.value === '3d' && data.value) {
        const params = selectedMfData.value?.funcParams

        nextTick(() => {
          const cnt = data.value.length || 1

          params.forEach((row, i) => {
            const l = row.length || 2

            if (l < cnt) {
              params[i] = [...params[i], ...Array(cnt - l).fill(0)]
            }
            if (l > cnt) {
              params[i] = row.slice(0, cnt)
            }
          })
        })
      }
    }

    watch(
      selectedMf,
      upadteParams,
      { deep: true }
    )

    watch(
      data,
      upadteParams,
      { deep: true }
    )

    onMounted(() => {
      if (qs.params.get('test') === '1') {
        store.commit(
          'general/setMfEvalDataByType',
          {
            k: props.mfid,
            v: [
              [0.87, 0.82, 0.60, 0.77, 0.69],
              [0.66, 0.83, 0.71, 0.98, 0.91],
              [0.78, 0.40, 0.54, 0.85, 0.82]
            ]
          }
        )
      } else {
        store.commit(
          'general/setMfEvalDataByType',
          { k: props.mfid, v: [[0, 0], [0, 0]] }
        )
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
