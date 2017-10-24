<template>
    <div class="activity-component">
        <div v-if="isLoaded">
            <legend>Kid Activity</legend>
            <router-link :to="`/kid`">< Back</router-link>

            <h3>
                Total Activity: {{ activityCount }}
            </h3>
            <md-table-card>
                <md-table>
                    <md-table-header>
                        <md-table-row>
                            <md-table-head>ID</md-table-head>
                            <md-table-head>Type</md-table-head>
                            <md-table-head>Steps</md-table-head>
                            <md-table-head>Date</md-table-head>
                        </md-table-row>
                    </md-table-header>

                    <md-table-body>
                        <md-table-row v-for="(a, index) in activity" :key="index" :md-item="a">
                            <md-table-cell>{{ a.ID }}</md-table-cell>
                            <md-table-cell>{{ a.Type }}</md-table-cell>
                            <md-table-cell>{{ a.Steps }}</md-table-cell>
                            <md-table-cell>{{ a.ReceivedDate }}</md-table-cell>
                        </md-table-row>
                    </md-table-body>
                </md-table>
                <md-table-pagination
                        :md-size="max"
                        :md-total="activityCount"
                        :md-page="page"
                        md-label="Activities"
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

  export default {
    name: "Activity",
    props: ['macId'],
    data: () => {
      return {
        isLoaded: false,
        page: 1,
        max: 20
      }

    },

    created: function () {
      this.$store.dispatch('getActivity', {
        macId: this.macId,
        page: this.page,
        max: this.max
      }).then(() => {
        this.isLoaded = true;
      })

    },

    computed: {
      ...mapGetters([
        'activityCount',
        'activity'
      ])
    },

    methods: {
      onPagination: function(context) {
        this.$store.dispatch('getActivity', {
          macId: this.macId,
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
    .activity-component {
        position: relative;
        width: 100%;
    }
</style>
