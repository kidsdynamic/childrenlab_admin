import axios from 'axios';

const errorStore = {
  state: {
    macIdNoKid: [],
  },
  getters: {
    macIdNoKid: state => state.macIdNoKid,

  },
  mutations: {
    updateMacIdNoKid(state, result) {
      state.macIdNoKid = result;
    },
  },
  actions: {
    getMacIdNoKid({commit}) {
      return axios.get("/admin/api/errorData", {})
        .then((result) => {
          if (result.status === 200) {
            commit('updateMacIdNoKid', result.data);

          }
        })
    },

  }
};
export default errorStore;