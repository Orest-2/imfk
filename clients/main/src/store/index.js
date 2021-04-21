import { createStore } from 'vuex'

import { general } from './general'
import { settings } from './settings'

export const store = createStore({
  modules: {
    general,
    settings
  }
})
