import './assets/main.css'

import {createApp} from 'vue'
import {createPinia} from 'pinia'

import App from './App.vue'
import router from './router'
import {setupI18n} from "@/i18n";

const app = createApp(App)

app.use(createPinia())
app.use(router)
setupI18n(app)
app.mount('#app')
