import axios from 'axios';

const dashboardStore = {
  state: {
    totalUserCount: 0,
    signup: [],
    totalActivityCount: 0,
    dashboardActivity: [],
    dashboardActivityByEventDate: [],
  },
  getters: {
    signup: state => state.signup,
    totalUserCount: state => state.totalUserCount,
    totalActivityCount: state => state.totalActivityCount,
    dashboardActivity: state => state.dashboardActivity,
    dashboardActivityByEventDate: state => state.dashboardActivityByEventDate

  },
  mutations: {
    updateDashboard(state, data) {
      state.signup = data.signup;
      state.totalUserCount = data.totalUserCount;
      state.totalActivityCount = data.totalActivityCount;
      state.dashboardActivity = data.activity;
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