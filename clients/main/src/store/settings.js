import axios from 'axios'
import { urls } from '../constants/urls'

export const settings = {
  namespaced: true,

  state () {
    return {
      membership_funcs: null
    }
  },

  getters: {
    getMembershipFuncsByType (state) {
      return (type) => {
        return state.membership_funcs?.filter(el => el.type === type) || []
      }
    }
  },

  mutations: {
    setMembershipFuncs (state, data) {
      state.membership_funcs = data
    }
  },

  actions: {
    fetchSettings ({ commit }) {
      axios.get(urls.settings)
        .then(({ data }) => {
          commit('setMembershipFuncs', data.data.membership_funcs)
        })
    }
  }
}
