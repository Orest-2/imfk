import axios from 'axios'
import { defaultParams } from '../constants/default-mf-params'
import { urls } from '../constants/urls'

const genKey = () => Math.random().toString(36).substr(2, 5)

const dc = (o) => JSON.parse(JSON.stringify(o))

const defaultMfObj = {
  mf: null,
  plotParams: [1, 0.01],
  funcParams: [],
  plotTraces: [],
  evalData: [],
  evalDataResult: []
}

export const general = {
  namespaced: true,

  state () {
    return {
      type: '2d',
      selectedMfs: {
        '2d': {
          [genKey()]: dc(defaultMfObj)
        },
        '3d': {
          [genKey()]: dc(defaultMfObj)
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

    getSelectedMfDataByKeyAndType (state) {
      return (k, type = state.type) => {
        return state.selectedMfs[type][k]
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
      state.selectedMfs[state.type][k] = dc(defaultMfObj)
    },

    setMfByType (state, { k, v }) {
      state.selectedMfs[state.type][k].mf = v

      if (defaultParams[v?.code]) {
        state.selectedMfs[state.type][k].funcParams = defaultParams[v.code]
      } else {
        Array(v?.params_count).fill(0)
      }
    },

    setMfPlotTracesByType (state, { k, v }) {
      state.selectedMfs[state.type][k].plotTraces = v
    },

    setMfEvalDataByType (state, { k, v }) {
      state.selectedMfs[state.type][k].evalData = v
    },

    setMfEvalDataResultByType (state, { k, v }) {
      if (state.selectedMfs[state.type][k]) {
        state.selectedMfs[state.type][k].evalDataResult = v
      }
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

    makePlot ({ state, commit, dispatch }, { k, payload }) {
      axios.post(
        urls.plot.replace(':type', state.type),
        payload
      )
        .then(({ data }) => {
          dispatch('error/clearErrors', null, { root: true })

          commit('setMfPlotTracesByType', { k, v: [data.data] })
        })
        .catch(err => {
          dispatch('error/setErrors', err, { root: true })
        })
    },

    evalData ({ state, commit, dispatch }, { k, payload }) {
      axios.post(
        urls.eval.replace(':type', state.type),
        payload
      )
        .then(({ data }) => {
          dispatch('error/clearErrors', null, { root: true })

          commit('setMfEvalDataResultByType', { k, v: [...data.data.result] })
        })
        .catch(err => {
          dispatch('error/setErrors', err, { root: true })
        })
    }
  }
}
