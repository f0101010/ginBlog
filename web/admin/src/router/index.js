import { createRouter, createWebHashHistory } from 'vue-router'
import LoginView from '../views/LoginView.vue'
import AdminView from '../views/AdminView.vue'

//页面路由组件
import HomeIndex from '../components/admin/HomeIndex.vue'
import AddArt from '../components/article/AddArt.vue'
import ArtList from '../components/article/ArtList.vue'
import CateList from '../components/category/CateList.vue'
import UserList from '../components/user/UserList.vue'

const routes = [
  {
    path: '/login',
    name: 'login',
    component: LoginView,
  },
  {
    path: '/',
    name: 'admin',
    component: AdminView,
    children: [
      { path: 'index', component: HomeIndex },
      { path: 'addart', component: AddArt },
      { path: 'addart/:id', component: AddArt, props: true },
      { path: 'artlist', component: ArtList },
      { path: 'catelist', component: CateList },
      { path: 'userlist', component: UserList },
    ]
  },

]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

router.beforeEach((to, from, next) => {
  const token = window.sessionStorage.getItem('token')
  if (to.path === '/login') return next()
  if (!token || to.path === '/admin') {
    next('/login')
  } else {
    next()
  }
})

export default router
