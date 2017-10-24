import axios from 'axios';

const activityStore = {
  state: {
    rawActivityCount: 0,
    rawActivity: [],
    activityCount: 0,
    activity: [],
  },
  getters: {
    rawActivityCount: state => state.rawActivityCount,
    rawActivity: state => state.rawActivity,
    activity: state => state.activity,
    activityCount: state => state.activityCount,

  },
  mutations: {
    updateRawActivity(state, result) {
      state.rawActivity = result.data;
      state.rawActivityCount = result.totalCount;
    },
    updateActivity(state, result) {
      state.activity = result.data;
      state.activityCount = result.totalCount;
    }
  },
  actions: {
    getRawActivity({commit}, macId) {
      return axios.get(`/admin/api/kid/rawActivity/${macId}`, {
      }).then((result) => {
        if (result.status === 200) {
          commit('updateRawActivity', result.data);

        }
      })
    },

    getActivity({commit}, params) {
      return axios.get(`/admin/api/kid/activity/${params.macId}`, {
        params: {
          max: params.max,
          page: params.page
        }
      }).then((result) => {
        if (result.status === 200) {
          commit('updateActivity', result.data);

        }
      })
    }
  }
};
export default activityStore;