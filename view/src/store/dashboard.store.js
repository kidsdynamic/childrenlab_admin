import axios from 'axios';

const dashboardStore = {
  state: {
    totalUserCount: 0,
    signup: [],
    totalActivityCount: 0,
    activity: [],
    activityByEventDate: [],
  },
  getters: {
    signup: state => state.signup,
    totalUserCount: state => state.totalUserCount,
    totalActivityCount: state => state.totalActivityCount,
    activity: state => state.activity,
    activityByEventDate: state => state.activityByEventDate

  },
  mutations: {
    updateDashboard(state, data) {
      console.log(data);
      state.signup.push(...data.signup);
      state.totalUserCount = data.totalUserCount;
      state.totalActivityCount = data.totalActivityCount;
      state.activity = data.activity;
      state.activityByEventDate = data.activityByEventDate;
    }
  },
  actions: {
    getDashboard({commit}) {
      return axios.get("/admin/api/dashboard", {}).then((result) => {
        if (result.status === 200) {
          commit('updateDashboard', result.data);

        }
      })
    }
  }
};
export default dashboardStore;