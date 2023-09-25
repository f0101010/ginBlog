import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import axios from 'axios'
import { Button, Form, Input, Message } from 'ant-design-vue'
import 'ant-design-vue/dist/reset.css'


const app = createApp(App)


Message.config({
    top: `60px`,
    duration: 2,
    maxCount: 3,
})

axios.defaults.baseURL = "http://localhost:3000/api/v1"
app.use(Button)
app.use(Form)
app.use(Input)
app.config.globalProperties.$message = Message
app.config.globalProperties.$http = axios
app.use(store).use(router).mount('#app')
