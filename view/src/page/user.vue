<template>
    <div class="dashboard-component">
        <legend>User</legend>

        <h3>
            Total User: {{totalUser}}
        </h3>
        <md-table-card>

            <md-table>
                <md-table-header>
                    <md-table-row>
                        <md-table-head>ID</md-table-head>
                        <md-table-head>Email</md-table-head>
                        <md-table-head>Name(F L)</md-table-head>
                        <md-table-head>Profile</md-table-head>
                        <md-table-head>iOS ID</md-table-head>
                        <md-table-head>Android ID</md-table-head>
                        <md-table-head>Country</md-table-head>
                        <md-table-head>Sign-up Date</md-table-head>
                    </md-table-row>
                </md-table-header>

                <md-table-body>
                    <md-table-row v-for="(a, index) in user" :key="index" :md-item="a">
                        <md-table-cell>{{ a.ID }}</md-table-cell>
                        <md-table-cell>{{ a.Email }}</md-table-cell>
                        <md-table-cell>{{ a.FirstName }} {{ a.LastName }}</md-table-cell>
                        <md-table-cell>{{ a.Profile }}</md-table-cell>
                        <md-table-cell>{{ a.RegistrationID ? a.RegistrationID.substring(0,3) : '' }}</md-table-cell>
                        <md-table-cell>{{ a.AndroidRegistrationToken ? a.AndroidRegistrationToken.substring(0,3) : '' }}</md-table-cell>
                        <md-table-cell>{{ a.SignUpCountryCode }}</md-table-cell>
                        <md-table-cell>{{ a.DateCreated }}</md-table-cell>
                    </md-table-row>
                </md-table-body>
            </md-table>
            <md-table-pagination
                    :md-size="max"
                    :md-total="totalUser"
                    :md-page="userPage"
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
    name: "User",
    data: () => {
      return {
        max: 20,
        page: 1
      }

    },

    created: function () {
      this.$store.dispatch('getUser', {
        max: this.max,
        page: this.page
      }).then(() => {
      })
    },

    computed: {
      ...mapGetters([
        'user',
        'totalUser',
        'userPage',
      ])
    },

    methods: {
      onPagination: function(context) {
        this.$store.dispatch('getUser', {
          max: context.size,
          page: context.page
        }).then(() => {
          window.scrollTo(0, 0);
        })
      }
    }
  }

</script>

<style lang="scss" scoped>
    h3 {

    }
</style>
