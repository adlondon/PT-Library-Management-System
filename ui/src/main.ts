import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import HomePage from './components/HomePage.vue'
import { createRouter, createWebHistory, type RouteRecordRaw } from 'vue-router'

const routes: RouteRecordRaw[] = [{ path: '/', component: HomePage }]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

createApp(App).use(router).mount('#app')
