import axios from 'axios';

const eventStore = {
  state: {
    totalEventCount: 0,
    events: [],
    eventPage: 1,
  },
  getters: {
    events: state => state.events,
    totalEventCount: state => state.totalEventCount,
    eventPage: state => state.eventPage,

  },
  mutations: {
    updateEvent(state, result) {
      state.events = result.data;
      state.totalEventCount = result.totalCount;
      state.eventPage = result.page;
    },
  },
  actions: {
    getEvents({commit}, params) {
      return axios.get("/admin/api/event", {
        params: {
          max: params.max,
          page: params.page,
          searchField: params.searchField,
          searchText: params.searchText,
        }
      }).then((result) => {
        if (result.status === 200) {
          commit('updateEvent', result.data);

        }
      })
    },
  }
};
export default eventStore;