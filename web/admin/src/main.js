import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import { Button, Form, Input, message, Layout, Menu, Card, Table, Row, Col, ConfigProvider, Modal, Switch, Select, Upload } from 'ant-design-vue'
import axios from 'axios'
import 'ant-design-vue/dist/reset.css'
import './assets/css/style.css'


const app = createApp(App)

message.config({
    top: `100px`,
    duration: 2,
    maxCount: 3,
})

axios.defaults.baseURL = 'http://localhost:3000/api/v1/'

axios.interceptors.request.use(config => {
    config.headers.Authorization = `Bearer ${window.sessionStorage.getItem('token')}`
    return config
})
app.config.globalProperties.$http = axios
app.config.globalProperties.$message = message
app.config.globalProperties.$confirm = Modal.confirm




app.use(store)
    .use(Form)
    .use(Button)
    .use(Input)
    .use(router)
    .use(Layout)
    .use(Menu)
    .use(Card)
    .use(Table)
    .use(Row)
    .use(Col)
    .use(ConfigProvider)
    .use(Modal)
    .use(Switch)
    .use(Select)
    .use(Upload)
    .mount('#app')
