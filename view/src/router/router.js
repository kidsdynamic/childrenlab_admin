import VueRouter from 'vue-router';
import Login from '../page/login.vue';
import store from '../store/store';
import Dashboard from '../page/dashboard.vue';
import User from '../page/user.vue';


const router = new VueRouter({
  mode: 'history',
  routes: [
    {
      path: '/',
      redirect: 'dashboard'
    },
    {
      path: '/login',
      component: Login,
    },
    {
      path: '/dashboard',
      component: Dashboard,
    },
    {
      path: '/user',
      component: User,
    }
  ]
});

router.beforeEach((to, from, next) => {
  if(to.path === "/login") {
    if(store.state.auth.token === '') {
      next();
    } else {
      next('/');
    }
    return;
  }

  if(store.state.auth.token === '') {
    next('/login');
  } else {
    next();
  }
});

export default router;