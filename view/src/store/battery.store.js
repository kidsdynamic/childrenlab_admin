import axios from 'axios';

const batteryStore = {
  state: {
    totalBatteryCount: 0,
    battery: [],
  },
  getters: {
    battery: state => state.battery,
    totalBatteryCount: state => state.totalBatteryCount,

  },
  mutations: {
    updateBattery(state, result) {
      state.battery = result.data;
      state.totalBatteryCount = result.totalCount;
    }
  },
  actions: {
    getBattery({commit}, params) {
      return axios.get(`/admin/api/kid/battery/${params.macId}`, {
        params: {
          max: params.max,
          page: params.page
        }
      }).then((result) => {
        if (result.status === 200) {
          commit('updateBattery', result.data);

        }
      })
    }
  }
};
export default batteryStore;