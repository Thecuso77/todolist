import { createRouter, createWebHistory } from 'vue-router'
import UserDetail from '../components/UserDetail.vue'
import UserCreate from '../components/UserCreate.vue'
import UserLogin from '../components/UserLogin.vue'

const routes = [
  {
    path: '/users/:id',
    name: 'user',
    component: UserDetail
  },
  {
    path: '/users/create',
    name: 'userCreate',
    component: UserCreate
  },
  {
    path: '/users/login',
    name: 'userLogin',
    component: UserLogin
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
  linkActiveClass: "text-blue-500"
})

export default router