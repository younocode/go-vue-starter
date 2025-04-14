import type { App } from 'vue'
import { createI18n } from 'vue-i18n'
import messages from '@intlify/unplugin-vue-i18n/messages'

export type ILocale = 'en' | 'zh-CN'
export const defaultLocale: ILocale = 'zh-CN'

// Import i18n resources
export const i18n = createI18n({
  locale: defaultLocale,
  messages,
})

function setI18nLanguage(locale: ILocale = defaultLocale) {
  i18n.global.locale = locale
}

export function loadLanguage(locale: ILocale = defaultLocale) {
  if (i18n.global.locale === locale) {
    return
  }

  setI18nLanguage(locale)
}

export function setupI18n(app: App<Element>, locale: ILocale = defaultLocale) {
  app.use(i18n)
  loadLanguage(locale)
}