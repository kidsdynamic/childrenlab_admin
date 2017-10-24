import axios from 'axios';

const auth = {
  state: {
    name: window.localStorage.getItem('name') || '',
    token: window.localStorage.getItem('token') || '',
  },
  mutations: {
    setCredential(state, loggedIn) {
      state.name = loggedIn.name;
      state.token = loggedIn.token;

      window.localStorage.setItem('name', state.name);
      window.localStorage.setItem('token', state.token);

    },

    removeCredential(state) {
      state.name = null;
      state.token = null;
      window.localStorage.removeItem('name');
      window.localStorage.removeItem('token');
    }
  },
  actions: {
    login({commit}, credential) {
      return axios.post("/admin/api/login", {
        name: credential.name,
        password: credential.password
      }).then((result) => {
        if (result.status === 200) {
          commit('setCredential', {
            name: result.data.username,
            token: result.data.access_token
          });

        }
      })
    },

    logout({commit}) {
      return commit('removeCredential');
    }
  }
};
export default auth;