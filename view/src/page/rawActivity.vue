<template>
    <div class="activity-component">
        <div v-if="isLoaded">
            <legend>Kid Activity</legend>
            <router-link :to="`/kid`">< Back</router-link>

            <h3>
                Total Raw Activity: {{ rawActivityCount }}
            </h3>
            <md-table-card>
                <md-table>
                    <md-table-header>
                        <md-table-row>
                            <md-table-head>ID</md-table-head>
                            <md-table-head>Indoor</md-table-head>
                            <md-table-head>Outdoor</md-table-head>
                            <md-table-head>Time</md-table-head>
                            <md-table-head>Timezone Offset</md-table-head>
                            <md-table-head>Date Created</md-table-head>
                        </md-table-row>
                    </md-table-header>

                    <md-table-body>
                        <md-table-row v-for="(a, index) in rawActivity" :key="index" :md-item="a">
                            <md-table-cell>{{ a.ID }}</md-table-cell>
                            <md-table-cell>{{ a.Indoor }}</md-table-cell>
                            <md-table-cell>{{ a.Outdoor }}</md-table-cell>
                            <md-table-cell>{{ a.Time }}</md-table-cell>
                            <md-table-cell>{{ a.TimeZoneOffset }}</md-table-cell>
                            <md-table-cell>{{ a.DateCreated }}</md-table-cell>
                        </md-table-row>
                    </md-table-body>
                </md-table>
            </md-table-card>
        </div>

        <div class="loading" v-if="!isLoaded">
            <md-spinner md-indeterminate></md-spinner>
        </div>

    </div>
</template>

<script>
  import Vue from 'vue';
  import {mapGetters} from 'vuex'

  export default {
    name: "RawActivity",
    props: ['macId'],
    data: () => {
      return {
        isLoaded: false
      }

    },

    created: function () {
      this.$store.dispatch('getRawActivity', this.macId).then(() => {
        this.isLoaded = true;
      })

    },

    computed: {
      ...mapGetters([
        'rawActivityCount',
        'rawActivity'
      ])
    },

    methods: {

    }
  }

</script>

<style lang="scss" scoped>
    .activity-component {
        position: relative;
        width: 100%;
    }
</style>
