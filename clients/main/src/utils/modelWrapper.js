import { computed } from 'vue'

export const useModelWrapper = (props, emit, name = 'modelValue') => {
  return computed({
    get: () => props[name],
    set: (value) => {
      emit(`update:${name}`, value)
    }
  })
}
