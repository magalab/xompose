import { createApp } from 'vue'
import './style.css'
import App from './App.vue'

import { router } from '@/router'

import i18n from '@/i18n'

import 'virtual:uno.css'

createApp(App).use(i18n).use(router).mount('#app')
