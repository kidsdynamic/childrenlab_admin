import Vue from 'vue';
import Vuex from 'vuex';
import auth from './auth.store';
import dashboardStore from './dashboard.store';
import userStore from './user.store';
import kidStore from './kid.store';
import activityStore from './activity.store';
import batteryStore from './battery.store';
import fwStore from './fw.store';
import eventStore from "./event.store";

Vue.use(Vuex);

const store = new Vuex.Store({
  modules: {
    auth,
    dashboardStore,
    userStore,
    kidStore,
    activityStore,
    batteryStore,
    fwStore,
    eventStore
  }
});

export default store;