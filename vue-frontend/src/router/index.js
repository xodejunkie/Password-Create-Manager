// src/router/index.js
import { createRouter, createWebHistory } from 'vue-router';
import HomePage from '../components/HomePage.vue';
import CreatePassword from '../components/CreatePassword.vue';
import GeneratePassword from '../components/GeneratePassword.vue';
import DatabaseView from '../components/DatabaseView.vue';

const routes = [
  {
    path: '/',
    name: 'HomePage',
    component: HomePage
  },
  {
    path: '/create-password',
    name: 'CreatePassword',
    component: CreatePassword
  },
  {
    path: '/generate-password',
    name: 'GeneratePassword',
    component: GeneratePassword
  },
  {
    path: '/database',
    name: 'DatabaseView',
    component: DatabaseView
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;
