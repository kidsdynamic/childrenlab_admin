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
axios.defaults.validateStatus = (status) => {
  return (status >= 200 && status < 300) || status === 403;
}
axios.interceptors.response.use((response) => {
  if(response.status === 403) {
    store.dispatch('logout');
    router.go('/login');
  } else {
    return response;
  }
}, (error) => {
  console.error(error);
});

new Vue({
  el: '#app',
  router,
  store,
  render: h => h(App)
});
