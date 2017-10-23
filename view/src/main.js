import Vue from 'vue'
import VueRouter from 'vue-router';
import axios from 'axios';
import App from './App.vue'
import router from './router/router';
import store from './store/store';

let VueMaterial = require('vue-material');
Vue.use(VueMaterial);
Vue.use(VueRouter);

Vue.material.registerTheme({
    default: {
        primary: 'grey',
        accent: 'blue'
    },
});

axios.defaults.headers.common['x-auth-token'] = store.state.auth.token || '';

new Vue({
  el: '#app',
  router,
  store,
  render: h => h(App)
});
