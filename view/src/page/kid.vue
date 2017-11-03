<template>
    <div class="dashboard-component">
        <legend>Kid</legend>
        <h3>
            Total Kid: {{totalKidCount}}
        </h3>
        <md-table-card>

            <md-table>
                <md-table-header>
                    <md-table-row>
                        <md-table-head>ID</md-table-head>
                        <md-table-head>Activity</md-table-head>
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
                        <md-table-cell>{{ a.ID }}</md-table-cell>
                        <md-table-cell>
                            <router-link :to="`/kid/activity/${a.MacID}`">Activity</router-link>|
                            <router-link :to="`/kid/rawActivity/${a.MacID}`">Raw Data</router-link>|
                            <router-link :to="`/kid/battery/${a.MacID}`">Battery</router-link>
                        </md-table-cell>
                        <md-table-cell>{{ a.Name.length > 8 ? `${a.Name.substring(0, 8)}...` : a.Name }}</md-table-cell>
                        <md-table-cell>{{ a.MacID }}</md-table-cell>
                        <md-table-cell>{{ a.FirmwareVersion }}</md-table-cell>
                        <md-table-cell>{{ a.ParentEmail }}</md-table-cell>
                        <md-table-cell>{{ a.DateCreated }}</md-table-cell>
                        <md-table-cell>
                            <md-button class="md-warn md-dense" @click="deleteKid(a.MacID)">Delete</md-button>
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
    </div>


</template>

<script>
  import Vue from 'vue';
  import {mapGetters} from 'vuex'

  export default {
    name: "Kid",
    data: () => {
      return {
        max: 20,
        page: 1,
        confirmationText: "",
      }

    },

    created: function () {
      this.$store.dispatch('getKids', {
        page: this.kidPage,
        max: this.max
      })
    },

    computed: {
      ...mapGetters([
        'kids',
        'totalKidCount',
        'kidPage'
      ])
    },

    methods: {
      onPagination: function(context) {
        this.page = context.page;
        this.max = context.max;

        this.updateKidList();
      },

      deleteKid: function(macId) {
        let retVal = confirm(`Do you want to delete mac ID: ${macId}?`);
        if( retVal === true ){
          this.$store.dispatch('deleteKid', macId);
          return true;
        }
        else{
          return false;
        }

      },

      updateKidList: function() {
        this.$store.dispatch('getKids', {
          page: this.page,
          max: this.max
        }).then(() => {
          window.scrollTo(0, 0);
        })
      }
    }
  }

</script>

<style lang="scss" scoped>
    .md-table .md-table-cell .md-button {
        min-width: 88px;
    }
</style>
