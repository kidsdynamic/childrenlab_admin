<template>
    <div id="app">

        <md-toolbar class="md-title">
            <header>Kid Dyanmic Admin Server</header>
        </md-toolbar>

        <div class="container">
            <div class="menu" v-if="isLoggedIn">
                <router-link to="/dashboard">
                    <md-button class="md-raised md-dense">Dashboard</md-button>
                </router-link>
                <router-link to="/user">
                    <md-button class="md-raised md-dense">User</md-button>
                </router-link>
                <router-link to="/kid">
                    <md-button class="md-raised md-dense">Kid</md-button>
                </router-link>
                <router-link to="/event">
                    <md-button class="md-raised md-dense">Event</md-button>
                </router-link>
                <router-link to="/fw">
                    <md-button class="md-raised md-dense">FW Management</md-button>
                </router-link>
                <router-link to="/errorData">
                    <md-button class="md-raised md-dense">Error Data</md-button>
                </router-link>
                <md-button class="md-raised md-warn md-dense" @click="logout">Logout</md-button>
            </div>
            <router-view></router-view>
        </div>
    </div>
</template>

<script>

  export default {
    components: {},
    computed: {
      isLoggedIn: {
        get() {
          return this.$store.state.auth.token !== '';
        }
      }
    },
    methods: {
      logout: function () {
        this.$store.dispatch('logout').then(() => {
          window.location = '/login';
        });
      }
    }
  }
</script>

<style src="vue-material/dist/vue-material.css"></style>
<style lang="scss">
    html {
        scroll-behavior: smooth;
    }

    body {
        font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
        margin: 0;
        padding: 0;
    }

    #app {
        padding-bottom: 15px;

        legend {
            padding-bottom: 5px;
            margin-bottom: 20px;
            border-bottom: 1px solid #e5e5e5;
            font-size: 1.7em;
            text-align: center;

        }

        .md-title {
            width: 100%;
            background-color: #FD733D;
            position: relative;
            top: 0;
            left: 0;
            font-size: 2em;
            padding: 5px 15px;
            box-shadow: 0 3px 5px -1px rgba(0, 0, 0, .2), 0 6px 10px 0 rgba(0, 0, 0, .14), 0 1px 18px 0 rgba(0, 0, 0, .12);
            color: white;
            margin-bottom: 1em;
        }

        .container {
            padding: 0 35px;
        }

        .menu {
            margin-bottom: 15px;
        }
    }

    .loading {
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        background-color: rgba(255, 255, 255, .5);
        justify-content: center;
        align-items: center;
        text-align: center;
    }

    .loader {
        display: flex;
        justify-content: center;
        padding: 15px;
    }


</style>
