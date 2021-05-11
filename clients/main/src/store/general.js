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

const defaultMfOperationObj = {
  operands: [],
  operation: 'intersection',
  plotParams: [1, 0.01],
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
        return state.selectedMfs[type][k]?.mf || null
      }
    },

    getOperandByKeyAndOperandIndxAndType (state) {
      return (k, operandIndx, type = state.type) => {
        return state.selectedMfs[type][k]?.operands[operandIndx] || null
      }
    }
  },

  mutations: {
    setType (state, type) {
      state.type = type
    },

    addMfByType (state, { k, data }) {
      state.selectedMfs[state.type][k] = data
    },

    setMfByType (state, { k, v }) {
      state.selectedMfs[state.type][k].mf = v

      if (defaultParams[v?.code]) {
        state.selectedMfs[state.type][k].funcParams = defaultParams[v.code]
      } else {
        state.selectedMfs[state.type][k].funcParams = Array(v?.params_count).fill(0)
      }
    },

    setOperandByType (state, { k, i, v }) {
      state.selectedMfs[state.type][k].operands[i] = {
        mf: v,
        funcParams: [],
        evalData: []
      }

      if (defaultParams[v?.code]) {
        state.selectedMfs[state.type][k].operands[i].funcParams = defaultParams[v.code]
      } else {
        state.selectedMfs[state.type][k].operands[i].funcParams = Array(v?.params_count).fill(0)
      }
    },

    setMfOpretionByType (state, { k, v }) {
      state.selectedMfs[state.type][k].operation = v
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
      const data = dc(defaultMfObj)

      commit('addMfByType', { k, data })
    },

    addMfOperation ({ commit }) {
      const k = genKey()
      const data = dc(defaultMfOperationObj)

      commit('addMfByType', { k, data })
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

    operationMakePlot ({ state, getters, commit, dispatch }, { k, payload }) {
      const { operation } = getters.getSelectedMfDataByKeyAndType(k)

      axios.post(
        urls.operationPlot.replace(':type', state.type).replace(':operation', operation),
        payload
      )
        .then(({ data }) => {
          dispatch('error/clearErrors', null, { root: true })

          const v = data.data.map((el, i) => ({
            x: el.x,
            y: el.y,
            name: el.membership_func_id ? `функця ${i + 1}` : 'результат',
            color: el.membership_func_id ? 'blue' : 'red',
            dash: el.membership_func_id ? 'dot' : 'solid'
          }))

          commit('setMfPlotTracesByType', { k, v })
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
    },

    operationEvalData ({ state, getters, commit, dispatch }, { k, payload }) {
      const { operation } = getters.getSelectedMfDataByKeyAndType(k)

      axios.post(
        urls.operationEval.replace(':type', state.type).replace(':operation', operation),
        payload
      )
        .then(({ data }) => {
          dispatch('error/clearErrors', null, { root: true })

          commit('setMfEvalDataResultByType', { k, v: [...data.data[2].result] })
        })
        .catch(err => {
          dispatch('error/setErrors', err, { root: true })
        })
    }
  }
}
