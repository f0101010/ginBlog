import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import { Button, Form, Input } from 'ant-design-vue'
import axios from 'axios'
import 'ant-design-vue/dist/reset.css'


const app = createApp(App)

axios.defaults.baseURL = 'http://localhost:3000/api/v1'
app.config.globalProperties.$http = axios



app.use(store)
    .use(Form)
    .use(Button)
    .use(Input)
    .use(router)
    .mount('#app')
