import axios from 'axios'
import { urls } from '../constants/urls'

const genKey = () => Math.random().toString(36).substr(2, 5)

// const renameProp = (oldProp, newProp, { [oldProp]: old, ...others }) => ({
//   [newProp]: old,
//   ...others
// })

export const general = {
  namespaced: true,

  state () {
    return {
      type: '2d',
      selectedMfs: {
        '2d': {
          [genKey()]: { mf: null }
        },
        '3d': {
          [genKey()]: { mf: null }
        }
      }
    }
  },

  getters: {
    getSelectedMfsByType (state) {
      return (type = state.type) => {
        return Object.entries(state.selectedMfs[type])
      }
    },

    getSelectedMfByKeyAndType (state) {
      return (k, type = state.type) => {
        return state.selectedMfs[type][k]?.mf
      }
    }
  },

  mutations: {
    setType (state, type) {
      state.type = type
    },

    addMfByType (state, { k }) {
      state.selectedMfs[state.type][k] = { mf: null }
    },

    setMfByType (state, { k, v }) {
      state.selectedMfs[state.type][k].mf = v
    },

    removeMfByTypeAndKey (state, { k }) {
      delete state.selectedMfs[state.type][k]
    }
  },

  actions: {
    addMf ({ commit }) {
      const k = genKey()

      commit('addMfByType', { k })
    },

    removeMf ({ commit }, k) {
      commit('removeMfByTypeAndKey', { k })
    },

    fetchSettings ({ commit }) {
      axios.get(urls.settings)
        .then(({ data }) => {
          commit('setMembershipFuncs', data.data.membership_funcs)
        })
    }
  }
}
