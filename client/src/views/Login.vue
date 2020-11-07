<template>
  <v-container>
    <v-row justify="center">
      <v-col>
        <h1>what the weather</h1>
      </v-col>
    </v-row>
    <v-row class="text-center">
      <v-col cols="12">
        <v-img
          :src="require('../assets/weather.svg')"
          class="my-3"
          contain
          height="200"
        />
      </v-col>
    </v-row>
    <v-row>
      <v-col col="4">
        <v-text-field
          solo
          type="text"
          v-model="email"
          placeholder="Enter Email"
        ></v-text-field>

        <v-text-field
          solo
          type="password"
          v-model="password"
          placeholder="Enter Password"
        ></v-text-field>

        <v-btn large @click="login">Sign in</v-btn>
        <v-row>
            <v-col>
                <router-link to="/register">Create an account</router-link>
            </v-col>
        </v-row>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import firebase from "firebase";
export default {
  name: "Login",
  data() {
    return {
      email: "",
      password: "",
    };
  },
  methods: {
    login: function() {
      firebase
        .auth()
        .signInWithEmailAndPassword(this.email, this.password)
        .then(
          (user) => {
            console.log("user", user);
            this.$router.replace("home");
          },
          (err) => {
            alert("Oops. " + err.message);
          }
        );
    },
  },
};
</script>
