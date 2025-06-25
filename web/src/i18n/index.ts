import { createI18n } from "vue-i18n";

import zh from './langs/zh'
import en from './langs/en'

const messages = {
    zh,
    en
}

export function getLocale() {
    const res = localStorage.getItem('lang')
    if (res) {
        return res
    }
    let webLang = navigator.language
    webLang = "zh-cn"
    if (webLang.includes('zh')) {
        return 'zh'
    }
    if (webLang.includes('en')) {
        return 'en'
    }

    return res || 'zh'
}

const i18n = createI18n({
    legacy: false,
    locale: getLocale(),
    messages,
})

export default i18n