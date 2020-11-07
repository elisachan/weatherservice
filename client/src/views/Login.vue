<template>
  <v-container>
    <Heading />
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
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <router-link to="/register">Create an account</router-link>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import firebase from "firebase";
import Heading from "../components/Heading.vue";
export default {
  name: "Login",
  data() {
    return {
      email: "",
      password: "",
    };
  },
  components: {
    Heading,
  },
  methods: {
    login: function() {
      firebase
        .auth()
        .signInWithEmailAndPassword(this.email, this.password)
        .then(
          () => {
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
