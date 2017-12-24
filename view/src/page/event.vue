<template>
    <div class="kid-component">
        <legend>Event</legend>
        <h3>
            Total Event: {{totalEventCount}}
        </h3>
        <md-table-card v-if="events.length > 0">
            <md-table>
                <md-table-header>
                    <md-table-row>
                        <md-table-head>ID</md-table-head>
                        <md-table-head>Name</md-table-head>
                        <md-table-head>Start</md-table-head>
                        <md-table-head>End</md-table-head>
                        <md-table-head>Color</md-table-head>
                        <md-table-head>Email</md-table-head>
                        <md-table-head>Description</md-table-head>
                        <md-table-head>Alert</md-table-head>
                        <md-table-head>Repeat</md-table-head>
                        <md-table-head>Timezone</md-table-head>
                    </md-table-row>
                </md-table-header>

                <md-table-body>
                    <md-table-row v-for="(a, index) in events" :key="index" :md-item="a">
                        <md-table-cell>{{ a.id }}</md-table-cell>
                        <md-table-cell>{{ a.name }}</md-table-cell>
                        <md-table-cell>{{ a.start }}</md-table-cell>
                        <md-table-cell>{{ a.end }}</md-table-cell>
                        <md-table-cell>
                            <div v-if="a.color" class="color" :style="`backgroundColor: ${a.color}`"></div>
                        </md-table-cell>
                        <md-table-cell>{{ a.email }}</md-table-cell>
                        <md-table-cell>{{ a.description }}</md-table-cell>
                        <md-table-cell>{{ a.alert }}</md-table-cell>
                        <md-table-cell>{{ a.repeat }}</md-table-cell>
                        <md-table-cell>{{ a.timezoneOffset }}</md-table-cell>
                    </md-table-row>
                </md-table-body>
            </md-table>
            <md-table-pagination
                    :md-size="max"
                    :md-total="totalEventCount"
                    :md-page="eventPage"
                    md-label="Events"
                    md-separator="of"
                    :md-page-options="false"
                    @pagination="onPagination"></md-table-pagination>
        </md-table-card>

        <div class="emptyState" v-show="!loading && events.length == 0">
            No Data
        </div>


        <div class="loader" v-if="loading">
            <md-spinner md-indeterminate></md-spinner>
        </div>

        <md-snackbar md-position="center" :md-duration="3000" :md-active.sync="showMessage" md-persistent>
            <span>{{ message }}</span>
            <md-button class="md-primary" @click="showSnackbar = false">Retry</md-button>
        </md-snackbar>
    </div>


</template>

<script>
  import {mapGetters} from 'vuex'

  export default {
    name: "Event",
    data: () => {
      return {
        max: 20,
        page: 1,
        searchText: '',
        loading: true,
        showMessage: false,
        message: '',
      }

    },
    created: function () {
      this.updateEventList();
    },

    computed: {
      ...mapGetters([
        'events',
        'totalEventCount',
        'eventPage'
      ])
    },

    methods: {
      onPagination(context) {
        this.page = context.page;

        this.updateEventList();
      },


      updateEventList() {
        this.loading = true;
        this.$store.dispatch('getEvents', {
          page: this.page,
          max: this.max,
          searchField: this.searchField,
          searchText: this.searchText,
        }).then(() => {
          window.scrollTo(0, 0);
          this.loading = false;
        })
      },
    },
  }

</script>

<style lang="scss" scoped>
    .kid-component {
        .filter-container {
            display: flex;
            padding: 15px 0 0 5px;
        }
        .searchInput {
            width: 250px;
            left: 30px;
        }

        .searchField {
            left: 10px;
            width: 120px;
        }
    }
    .emptyState {
        font-size: 2em;
        text-align: center;
        background-color: darkgray;
        color: white;
        padding: 10px;
        border-radius: 5px;
    }
    .color {
        width: 10px;
        height: 10px;
    }
</style>
