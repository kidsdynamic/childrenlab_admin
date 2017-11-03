import axios from 'axios';

const fwStore = {
  state: {
    fwList: [],
  },
  getters: {
    fwList: state => state.fwList,

  },
  mutations: {
    updateFwList(state, result) {
      state.fwList = result;
    }
  },
  actions: {
    getAllVersion({commit}, params) {
      return axios.get(`/admin/api/fw`, {}).then((result) => {
        if (result.status === 200) {
          commit('updateFwList', result.data);

        }
      })
    },

    uploadNewVersion({commit}, params) {
      const formData = new FormData();
      formData.append('fileA', params.fileA);
      formData.append('fileB', params.fileB);
      formData.append('versionName', params.version);

      return axios.post('/v1/admin/fwFile', formData, {
        headers: {
          'enctype': 'multipart/form-data',
        }
      })
    },

    updateActivation({ commit }, params) {
      return axios.put('/admin/api/fw/updateActivation', {
        id: params.id,
        active: params.active
      }).then((result) => {
        commit('updateFwList', result.data);
      })
    }
  }
};
export default fwStore;