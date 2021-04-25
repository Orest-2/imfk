import { createStore } from 'vuex'

import { general } from './general'
import { settings } from './settings'
import { error } from './error'

export const store = createStore({
  modules: {
    general,
    settings,
    error
  }
})
