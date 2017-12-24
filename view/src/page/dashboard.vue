<template>
    <div class="dashboard-component">
        <legend>Dashboard</legend>

        <div v-if="dashboardActivity.length > 0">
            <h3>
                Daily Signup
            </h3>
            <md-card>
                <daily-signup :sourceData="signup"></daily-signup>
            </md-card>

            <h3>
                Total Activities: {{totalActivityCount}}
            </h3>
            <md-table-card>
                <md-card>
                    <activity-chart :sourceData="dashboardActivity"></activity-chart>
                </md-card>
            </md-table-card>

            <h3>
                Total Events: {{totalEventCount}}
            </h3>
            <md-table-card>
                <md-card>
                    <event-chart :sourceData="dashboardEvent"></event-chart>
                </md-card>
            </md-table-card>
        </div>
        <div v-else>
            <div class="loader">
                <md-spinner md-indeterminate></md-spinner>
            </div>
        </div>
    </div>


</template>

<script>
  import Vue from 'vue';
  import {mapGetters} from 'vuex'
  import dailySignup from '../charts/signupChart'
  import activityChart from '../charts/activityChart'
  import eventChart from '../charts/eventChart'

  export default {
    name: "Dashboard",
    data: () => {
      return {}

    },

    components: {
      dailySignup,
      activityChart,
      eventChart,
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
        'dashboardActivity',
        'dashboardEvent',
        'totalEventCount'
      ])
    },

    methods: {},
  }

</script>

<style lang="scss" scoped>
    h3 {

    }
</style>
