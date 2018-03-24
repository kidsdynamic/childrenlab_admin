<template>
    <div class="activity-component">
        <div v-if="isLoaded">
            <legend>Error Data</legend>
            <h3>
                Mac ID without Kid (Lose data) - Total: {{ macIdNoKid.length }}
            </h3>
            <md-table-card>
                <md-table>
                    <md-table-header>
                        <md-table-row>
                            <md-table-head>Mac ID</md-table-head>
                            <md-table-head>Activity Count</md-table-head>
                            <md-table-head>Last Uploaded</md-table-head>
                        </md-table-row>
                    </md-table-header>

                    <md-table-body>
                        <md-table-row v-for="(a, index) in macIdNoKid" :key="index" :md-item="a">
                            <md-table-cell>{{ a.mac_id }}</md-table-cell>
                            <md-table-cell>{{ a.activity_count }}</md-table-cell>
                            <md-table-cell>{{ a.last_upload }}</md-table-cell>
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
  import moment from 'moment'

  export default {
    name: "Error Data",
    data: () => {
      return {
        isLoaded: false,
      }
    },

    created: function () {
      this.$store.dispatch('getMacIdNoKid')
        .then(() => {
          this.isLoaded = true;
        })
    },

    computed: {
      ...mapGetters([
        'macIdNoKid',
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
