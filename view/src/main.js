import Vue from 'vue'
import VueRouter from 'vue-router';
import axios from 'axios';
import App from './App.vue'
import router from './router/router';
import store from './store/store';
import favicon from './assets/favicon.ico';

let VueMaterial = require('vue-material');
Vue.use(VueMaterial);
Vue.use(VueRouter);

Vue.material.registerTheme({
});

axios.defaults.headers.common['x-auth-token'] = store.state.auth.token || '';
axios.interceptors.response.use((response) => {
  if(response.status === 403) {
    store.dispatch('logout');
    router.go('/login');
  } else {
    return response;
  }
}, (error) => {
  if(error.indexOf('403') !== -1) {
    store.dispatch('logout');
    router.go('/login');
  }

});

new Vue({
  el: '#app',
  router,
  store,
  render: h => h(App)
});
