import Vue from 'vue'
import VueRouter from 'vue-router'
import firebase from 'firebase';

import Home from '@/views/Home';
import Login from '@/views/Login';
import Register from '@/views/Register';

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home,
    meta: {
      requiresAuth: true
    }
  },
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/register',
    name: 'Register',
    component: Register
  }
]

const router = new VueRouter({
  routes
})

// base router logic to go to login page if not logged 
router.beforeEach((to, from, next) => {
  const requiresAuth = to.matched.some(record => record.meta.requiresAuth);
  if (requiresAuth) {
    const currentUser = firebase.auth().currentUser;
    if (!currentUser){
      next('login')
    }
  } else {
    next()
  }
});

export default router
