<template>
    <div class="activity-component">
        <div v-if="isLoaded">
            <legend>Kid Battery</legend>
            <router-link :to="`/kid`">< Back</router-link>

            <h3>
                Total Battery: {{ totalBatteryCount }}
            </h3>
            <md-table-card>
                <md-table>
                    <md-table-header>
                        <md-table-row>
                            <md-table-head>Mac ID</md-table-head>
                            <md-table-head>Battery Life</md-table-head>
                            <md-table-head>Date</md-table-head>
                        </md-table-row>
                    </md-table-header>

                    <md-table-body>
                        <md-table-row v-for="(a, index) in battery" :key="index" :md-item="a">
                            <md-table-cell>{{ a.MacID }}</md-table-cell>
                            <md-table-cell>{{ a.BatteryLife }}</md-table-cell>
                            <md-table-cell>{{ toDate(a.DateReceived) }}</md-table-cell>
                        </md-table-row>
                    </md-table-body>
                </md-table>
                <md-table-pagination
                        :md-size="max"
                        :md-total="totalBatteryCount"
                        :md-page="page"
                        md-label="Battery"
                        md-separator="of"
                        :md-page-options="false"
                        @pagination="onPagination"></md-table-pagination>
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
  import moment from 'moment'

  export default {
    name: "Battery",
    props: ['macId'],
    data: () => {
      return {
        isLoaded: false,
        page: 1,
        max: 20
      }

    },

    created: function () {
      this.$store.dispatch('getBattery', {
        macId: this.macId,
        page: this.page,
        max: this.max
      }).then(() => {
        this.isLoaded = true;
      })

    },

    computed: {
      ...mapGetters([
        'battery',
        'totalBatteryCount'
      ])
    },

    methods: {
      onPagination: function(context) {
        this.$store.dispatch('getBattery', {
          max: context.size,
          page: context.page
        }).then(() => {
          window.scrollTo(0, 0);
        })
      },
      toDate: function(time) {
        const dateTime = moment(time*1000);
        const formatted = dateTime.format("YYYY-MM-D HH:mm:ss");
        return formatted
      }
    }
  }

</script>

<style lang="scss" scoped>
    .activity-component {
        position: relative;
        width: 100%;
    }
</style>
