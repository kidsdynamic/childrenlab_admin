import VueRouter from 'vue-router';
import Login from '../page/login.vue';
import store from '../store/store';
import Dashboard from '../page/dashboard.vue';
import User from '../page/user.vue';
import Kid from '../page/kid.vue';
import Event from '../page/event.vue';
import RawActivity from '../page/rawActivity.vue';
import Activity from '../page/activity.vue';
import Battery from '../page/battery.vue';
import FwManagement from '../page/fwManagement.vue';
import ErrorDatas from '../page/ErrorDatas.vue';


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
    },
    {
      path: '/kid',
      component: Kid
    },
    {
      path: '/kid/rawActivity/:macId',
      component: RawActivity,
      props: true
    },
    {
      path: '/kid/activity/:macId',
      component: Activity,
      props: true
    },
    {
      path: '/kid/battery/:macId',
      component: Battery,
      props: true
    },
    {
      path: '/event',
      component: Event,
      props: true
    },
    {
      path: '/fw',
      component: FwManagement,
      props: true
    },
    {
      path: '/errorData',
      component: ErrorDatas,
      props: true
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