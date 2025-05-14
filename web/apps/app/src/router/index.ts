import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      // component: () => import('../views/HomeView.vue'),
      redirect: '/sign-in',
    },
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/AboutView.vue'),
    },
    {
      path: '/sign-in',
      name: 'signIn',
      component: () => import('@/components/auth/SignIn.vue'),
    },
    {
      path: '/reset-password',
      name: 'resetPassword',
      component: () => import('@/components/auth/ResetPassword.vue'),
    }
  ],
})

export default router
