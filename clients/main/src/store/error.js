export const error = {
  namespaced: true,

  state () {
    return {
      list: []
    }
  },

  mutations: {
    setErrorList (state, data) {
      state.list = data
    }
  },

  actions: {
    setErrors ({ commit }, data) {
      if (data?.response) {
        const { error } = data?.response?.data || { error: 'Internal server error' }

        commit('setErrorList', [{ detail: error }])
      } else {
        commit('setErrorList', [{ detail: data }])
      }
    },

    clearErrors ({ commit }) {
      commit('setErrorList', [])
    }
  }
}
