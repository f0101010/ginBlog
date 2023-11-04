import { createRouter, createWebHashHistory } from 'vue-router'

const HomeView = () =>
  import(/* webpackChunkName: "group-index" */ '../views/HomeView.vue')
const ArticleList = () =>
  import(/* webpackChunkName: "group-article" */ '../components/ArticleList.vue')
const ArticleDetails = () =>
  import(/* webpackChunkName: "group-detail" */ '../components/ArticleDetails.vue')
const CategoryList = () =>
  import(/* webpackChunkName: "group-category" */ '../components/CategoryList.vue')
const SearchResult = () =>
  import(/* webpackChunkName: "group-search" */ '../components/SearchResult.vue')


const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView,
    children: [
      {
        path: '/',
        component: ArticleList,
        meta: {
          title: '欢迎来到GInBlog'
        }
      },
      {
        path: '/article/detail/:id',
        component: ArticleDetails,
        meta: {
          title: '文章详情'
        },
        props: true
      },

      {
        path: '/category/:cid',
        component: CategoryList,
        meta: {
          title: '分类信息'
        },
        props: true
      },
      {
        path: '/search/:title',
        component: SearchResult,
        meta: {
          title: '搜索结果'
        },
        props: true
      },
    ]
  },
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

router.beforeEach((to, from, next) => {
  if (to.meta.title) {
    document.title = to.meta.title
  }
  next()
})

export default router
