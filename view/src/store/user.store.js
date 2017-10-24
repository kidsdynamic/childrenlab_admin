import axios from 'axios';

const userStore = {
  state: {
    totalUser: 0,
    user: [],
    userPage: 1,
  },
  getters: {
    user: state => state.user,
    totalUser: state => state.totalUser,
    userPage: state => state.page,

  },
  mutations: {
    updateUser(state, result) {
      state.user = result.data;
      state.totalUser = result.total;
    }
  },
  actions: {
    getUser({commit}, pagination) {
      return axios.get("/admin/api/user", {
        params: {
          max: pagination.max,
          page: pagination.page
        }
      }).then((result) => {
        if (result.status === 200) {
          commit('updateUser', result.data);

        }
      })
    }
  }
};
export default userStore;