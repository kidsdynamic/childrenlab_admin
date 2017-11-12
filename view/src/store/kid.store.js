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
    },
    deleteKid(state, macId) {
      state.totalKidCount -= 1;
      for(let i = 0; i < state.kids.length; i++) {
        if(state.kids[i].MacID === macId) {
          state.kids.splice(i, 1);
        }
      }
    }
  },
  actions: {
    getKids({commit}, params) {
      return axios.get("/admin/api/kid", {
        params: {
          max: params.max,
          page: params.page,
          searchEmail: params.searchEmail,
        }
      }).then((result) => {
        if (result.status === 200) {
          commit('updateKids', result.data);

        }
      })
    },

    deleteKid({ commit }, macId) {
      return axios.delete(`/admin/api/kid/deleteMacID?macId=${macId}`).then((result) => {
        commit('deleteKid', macId);
      });
    }
  }
};
export default kidStore;