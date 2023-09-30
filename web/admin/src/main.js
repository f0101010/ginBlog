import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import { Button, Form, Input, message, Layout } from 'ant-design-vue'
import axios from 'axios'
import 'ant-design-vue/dist/reset.css'


const app = createApp(App)

message.config({
    top: `100px`,
    duration: 2,
    maxCount: 3,
})

axios.defaults.baseURL = 'http://localhost:3000/api/v1'
app.config.globalProperties.$http = axios
app.config.globalProperties.$message = message



app.use(store)
    .use(Form)
    .use(Button)
    .use(Input)
    .use(router)
    .use(Layout)
    .mount('#app')
