<template>
    <div class="activity-component">
        <div v-if="isLoaded">
            <legend>FW Management</legend>

            <div class="headerContainer">
                <h3>
                    Total Version: {{ fwList.length  }}
                </h3>
                <md-input-container>
                    <label for="filter">Filter</label>
                    <md-select name="filter" id="filter" @change="updateList">
                        <md-option v-for="filter in filterList" :key="filter" :value="filter">{{ filter }}</md-option>
                    </md-select>
                </md-input-container>
            </div>

            <div class="uploadSection">
                <md-switch v-model="showUploadSection" class="md-primary">Upload New Firmware</md-switch>

                <transition name="fade" mode="out-in">
                <div key="2" class="uploadFile" v-if="showUploadSection">
                    <md-card>
                        <md-card-content>
                            <md-input-container>
                                <label>Version</label>
                                <md-input v-model="version"></md-input>
                            </md-input-container>
                            <md-input-container>
                                <label>File A</label>
                                <md-file @selected="setFileA"></md-file>
                            </md-input-container>
                            <md-input-container>
                                <label>File B</label>
                                <md-file @selected="setFileB"></md-file>
                            </md-input-container>
                        </md-card-content>
                        <md-card-actions>
                            <md-button class="md-raised md-dense" @click="uploadFiles">Upload</md-button>
                        </md-card-actions>
                    </md-card>
                </div>
                </transition>
            </div>


            <md-table-card>
                <md-table>
                    <md-table-header>
                        <md-table-row>
                            <md-table-head>ID</md-table-head>
                            <md-table-head>Version</md-table-head>
                            <md-table-head>File A</md-table-head>
                            <md-table-head>File B</md-table-head>
                            <md-table-head>Uploaded Date</md-table-head>
                            <md-table-head>Active</md-table-head>
                        </md-table-row>
                    </md-table-header>

                    <md-table-body>
                        <md-table-row v-for="(a, index) in localFwList" :key="index" :md-item="a">
                            <md-table-cell>{{ a.ID }}</md-table-cell>
                            <md-table-cell>{{ a.Version }}</md-table-cell>
                            <md-table-cell>
                                <a :href="a.FileAURL">
                                    {{ a.Version }}_A
                                </a>
                            </md-table-cell>
                            <md-table-cell>
                                <a :href="a.FileBURL">
                                    {{ a.Version }}_B
                                </a>
                            </md-table-cell>
                            <md-table-cell>{{ a.UploadedDate }}</md-table-cell>
                            <md-table-cell>
                                <md-switch v-model="a.Active" @change="updateActivation(a.ID, !a.Active)" class="md-primary"></md-switch>
                            </md-table-cell>
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
    name: "FwManagement",
    data: () => {
      return {
        isLoaded: false,
        filterList: [
          'All',
          'Japanese',
          'English',
          'OLD'
        ],
        fileA: null,
        fileB: null,
        version: null,
        showUploadSection: false
      }

    },

    created: function () {
      this.fetchData();

    },

    computed: {
      ...mapGetters([
        'fwList',
      ]),
      localFwList: function() {
        return this.fwList;
      }
    },

    methods: {
      fetchData() {
        this.$store.dispatch('getAllVersion', {}).then(() => {
          this.showUploadSection = false;
          this.isLoaded = true;
        })
      },
      updateList(selected) {
        if (selected === 'English') {
          this.localFwList = this.fwList.filter(
            fw => fw.Version.indexOf('-E') !== -1
          );
        } else if (selected === 'Japanese') {
          this.localFwList = this.fwList.filter(
            fw => fw.Version.indexOf('-J') !== -1
          );
        } else if (selected === 'OLD') {
          this.localFwList = this.fwList.filter(
            fw => fw.Version.indexOf('-A') !== -1
          );
        } else {
          this.localFwList = this.fwList;
        }
      },

      uploadFiles() {
        this.isLoaded = false;
        this.$store.dispatch('uploadNewVersion', {
          fileA: this.fileA[0],
          fileB: this.fileB[0],
          version: this.version,
        }).then(() => {
          this.isLoaded = true;
          this.fetchData();
        })
      },

      setFileA(fileA) {
        this.fileA = fileA;
      },

      setFileB(fileB) {
        this.fileB = fileB;
      },

      updateActivation(id, active) {
        this.$store.dispatch('updateActivation', {
          id: id,
          active: active,
        })
      }
    }
  }

</script>

<style lang="scss" scoped>

    .md-input-container {
        max-width: 150px;
    }
    .headerContainer {
        display: flex;
        justify-content: space-between;
    }
    .uploadSection {
        display: flex;
        flex-direction: column;
        align-items: center;
    }
    .toggleButton {
        display: inline-block;
        margin-bottom: 15px;
    }
    .uploadFile {
        width: 400px;
        display: inline-block;
        margin-bottom: 25px;

        .md-input-container {
            max-width: inherit;
        }
    }

</style>

<style lang="scss">
    .md-table .md-table-cell .md-table-cell-container {
        padding: 0 32px 0 24px;
    }
</style>
