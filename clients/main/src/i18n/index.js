import { createI18n } from 'vue-i18n'

import { messages } from './messages'

export const i18n = createI18n({
  locale: 'en', // set locale
  fallbackLocale: 'en',
  messages
})
