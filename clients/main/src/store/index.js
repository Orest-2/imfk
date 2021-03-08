import { createStore } from 'vuex'

import { settings } from './settings'

export const store = createStore({
  modules: {
    settings
  }
})
