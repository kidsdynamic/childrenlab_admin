<template>
    <div class="kid-component">
        <legend>Kid</legend>
        <h3>
            Total Kid: {{totalKidCount}}
        </h3>
        <md-table-card v-if="kids.length > 0">

            <md-input-container md-clearable class="searchInput">
                <label>Search Email</label>
                <md-input v-model="searchEmail" @keyup.native="search"></md-input>
            </md-input-container>
            <md-table>
                <md-table-header>
                    <md-table-row>
                        <md-table-head>ID</md-table-head>
                        <md-table-head>Detail</md-table-head>
                        <md-table-head>Activity Count</md-table-head>
                        <md-table-head>Name</md-table-head>
                        <md-table-head>Mac ID</md-table-head>
                        <md-table-head>Version</md-table-head>
                        <md-table-head>Parent</md-table-head>
                        <md-table-head>Date Created</md-table-head>
                        <md-table-head>Option</md-table-head>
                    </md-table-row>
                </md-table-header>

                <md-table-body>
                    <md-table-row v-for="(a, index) in kids" :key="index" :md-item="a">
                        <md-table-cell>{{ a.id }}</md-table-cell>
                        <md-table-cell>
                            <router-link :to="`/kid/activity/${a.id}`">Activity</router-link>
                            |
                            <router-link :to="`/kid/rawActivity/${a.macId}`">Raw</router-link>
                            |
                            <router-link :to="`/kid/battery/${a.macId}`">Battery</router-link>
                        </md-table-cell>
                        <md-table-cell>{{ a.activity }}</md-table-cell>
                        <md-table-cell>{{ a.name.length > 8 ? `${a.name.substring(0, 8)}...` : a.name }}</md-table-cell>
                        <md-table-cell>{{ a.macId }}</md-table-cell>
                        <md-table-cell>{{ a.firmwareVersion }}</md-table-cell>
                        <md-table-cell>{{ a.parentEmail }}</md-table-cell>
                        <md-table-cell>{{ a.dateCreated }}</md-table-cell>
                        <md-table-cell>

                            <md-button class="md-icon-button md-accent" @click="deleteKid(a.macId)">
                                <md-icon class="md-warn">delete</md-icon>
                            </md-button>
                        </md-table-cell>
                    </md-table-row>
                </md-table-body>
            </md-table>
            <md-table-pagination
                    :md-size="max"
                    :md-total="totalKidCount"
                    :md-page="kidPage"
                    md-label="Users"
                    md-separator="of"
                    :md-page-options="false"
                    @pagination="onPagination"></md-table-pagination>

        </md-table-card>
        <div class="loader" v-if="loading">
            <md-spinner md-indeterminate></md-spinner>
        </div>
    </div>


</template>

<script>
  import {mapGetters} from 'vuex'

  export default {
    name: "Kid",
    data: () => {
      return {
        max: 20,
        page: 1,
        searchEmail: '',
        loading: true,
      }

    },

    created: function () {
      this.updateKidList();
    },

    computed: {
      ...mapGetters([
        'kids',
        'totalKidCount',
        'kidPage'
      ])
    },

    methods: {
      onPagination(context) {
        this.page = context.page;

        this.updateKidList();
      },

      deleteKid(macId) {
        let retVal = confirm(`Do you want to delete mac ID: ${macId}?`);
        if (retVal === true) {
          this.$store.dispatch('deleteKid', macId);
          return true;
        }
        else {
          return false;
        }

      },

      search() {
        this.updateKidList();
      },

      updateKidList: function () {
        this.loading = true;
        this.$store.dispatch('getKids', {
          page: this.page,
          max: this.max,
          searchEmail: this.searchEmail,
        }).then(() => {
          window.scrollTo(0, 0);
          this.loading = false;
        })
      }
    },
  }

</script>

<style lang="scss" scoped>
    .kid-component {
        .searchInput {
            width: 250px;
            left: 10px;
        }
    }
</style>
