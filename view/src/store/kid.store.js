import axios from 'axios';

const kidStore = {
  state: {
    totalKidCount: 0,
    kids: [],
    kidPage: 1,
  },
  getters: {
    kids: state => state.kids,
    totalKidCount: state => state.totalKidCount,
    kidPage: state => state.kidPage,

  },
  mutations: {
    updateKids(state, result) {
      state.kids = result.data;
      state.totalKidCount = result.totalKidCount;
      state.kidPage = result.page;
    }
  },
  actions: {
    getKids({commit}, pagination) {
      return axios.get("/admin/api/kid", {
        params: {
          max: pagination.max,
          page: pagination.page
        }
      }).then((result) => {
        if (result.status === 200) {
          commit('updateKids', result.data);

        }
      })
    }
  }
};
export default kidStore;