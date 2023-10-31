import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import axios from 'axios'
import vuetify from './plugins/vuetify'



const app = createApp(App)

axios.defaults.baseURL = 'http://localhost:3000/api/v1'

app.config.globalProperties.$http = axios

app.use(store).use(router).use(vuetify).mount('#app')
