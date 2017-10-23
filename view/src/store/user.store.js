import axios from 'axios';

const userStore = {
  state: {
    user: []
  },
  getters: {
    user: state => state.user,

  },
  mutations: {
    updateUser(state, data) {
      console.log(data);
      state.user = data;
    }
  },
  actions: {
    getUser({commit}) {
      return axios.get("/admin/api/user", {}).then((result) => {
        if (result.status === 200) {
          commit('updateUser', result.data);

        }
      })
    }
  }
};
export default userStore;