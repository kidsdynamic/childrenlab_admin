<template>
    <div class="login-component" v-if="!loggedIn">
        <legend>Admin Server</legend>
        <form>
            <!--<label class="error-message label-danger label">Please enter valid credential</label>-->
            <md-input-container>
                <label>Name</label>
                <md-input v-model="name" required></md-input>
            </md-input-container>
            <md-input-container md-has-password>
                <label>Password</label>
                <md-input type="password" v-model="password" @keyup.native="enter" required></md-input>
            </md-input-container>
            <md-button class="md-raised md-warn" @click="clicked">Login</md-button>

            <div class="loading" v-if="loading">
                <md-spinner md-indeterminate></md-spinner>
            </div>
        </form>

    </div>


</template>

<script>
  import Vue from 'vue';

  export default {
    name: "Login",
    data: () => {
      return {
        name: '',
        password: '',
        loading: false
      }
    },
    computed: {
      loggedIn: {
        get() {
          return this.$store.state.auth.token !== '';
        }
      }
    },

    methods: {
      clicked: function () {
        this.loading = true;
        this.$store.dispatch('login', {
          name: this.name,
          password: this.password
        }).then(() => {
          this.loading = false;
          window.location = "/"
        });
      },

      enter: function(key) {
        if(key.code === 'Enter') {
          this.clicked();
        }
      }
    }
  };

</script>

<style lang="scss" scoped>
    .login-component {
        display: flex;
        flex-direction: column;
        align-items: center;

        form {
            width: 250px;
            text-align: center;
            position: relative;
        }
    }

</style>
