<template>
    <div class="dashboard-component">
        <legend>Dashboard</legend>

        <div class="">
            <h3>
                Daily Signup: {{totalUserCount}}
            </h3>
            <md-table>
                <md-table-header>
                    <md-table-row>
                        <md-table-head>Sign-up Count</md-table-head>
                        <md-table-head>Date</md-table-head>
                    </md-table-row>
                </md-table-header>

                <md-table-body>
                    <md-table-row v-for="(s, index) in signup" :key="index">
                        <md-table-cell>{{ s.SignupCount }}</md-table-cell>
                        <md-table-cell>{{ s.Date }}</md-table-cell>
                    </md-table-row>
                </md-table-body>
            </md-table>

            <h3>
                Total Activities: {{totalActivityCount}}
            </h3>
            <md-table>
                <md-table-header>
                    <md-table-row>
                        <md-table-head>Activity Count</md-table-head>
                        <md-table-head>Indoor Steps</md-table-head>
                        <md-table-head>Outdoor Steps</md-table-head>
                        <md-table-head>User Count</md-table-head>
                        <md-table-head>Date</md-table-head>
                    </md-table-row>
                </md-table-header>

                <md-table-body>
                    <md-table-row v-for="(a, index) in activity" :key="index">
                        <md-table-cell>{{ a.ActivityCount }}</md-table-cell>
                        <md-table-cell>{{ a.IndoorSteps }}</md-table-cell>
                        <md-table-cell>{{ a.OutdoorSteps }}</md-table-cell>
                        <md-table-cell>{{ a.UserCount }}</md-table-cell>
                        <md-table-cell>{{ a.Date }}</md-table-cell>
                    </md-table-row>
                </md-table-body>
            </md-table>
        </div>
    </div>


</template>

<script>
  import Vue from 'vue';
  import {mapGetters} from 'vuex'

  export default {
    name: "Dashboard",
    data: () => {
      return {}

    },

    created: function () {
      this.$store.dispatch('getDashboard').then(() => {
      })
    },

    computed: {
      ...mapGetters([
        'signup',
        'totalUserCount',
        'totalActivityCount',
        'activity'
      ])
    },

    methods: {}
  }

</script>

<style lang="scss" scoped>
    h3 {

    }
</style>
