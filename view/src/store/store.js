import Vue from 'vue';
import Vuex from 'vuex';
import auth from './auth.store';
import dashboardStore from './dashboard.store';
import userStore from './user.store';
import kidStore from './kid.store';
import activityStore from './activity.store';
import batteryStore from './battery.store';

Vue.use(Vuex);

const store = new Vuex.Store({
  modules: {
    auth,
    dashboardStore,
    userStore,
    kidStore,
    activityStore,
    batteryStore
  }
});

export default store;