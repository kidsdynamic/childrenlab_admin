import Vue from 'vue';
import Vuex from 'vuex';
import auth from './auth.store';
import dashboardStore from './dashboard.store';
import userStore from './user.store';

Vue.use(Vuex);

const store = new Vuex.Store({
  modules: {
    auth,
    dashboardStore,
    userStore
  }
});

export default store;