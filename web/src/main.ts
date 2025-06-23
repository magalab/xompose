import { createApp } from 'vue'
import '@/style.css'
import App from '@/App.vue'
import { createPinia } from 'pinia'
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'
import router from '@/router'

const pinia = createPinia()
pinia.use(piniaPluginPersistedstate)

// unocss
import 'virtual:uno.css'


createApp(App).
    use(pinia).
    use(router).
    mount('#app')
