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
        <v-btn large @click="signUp">Create an account</v-btn>
      </v-col>
    </v-row>
     <v-row>
        <v-col>
          <router-link to="/login">Go back to login</router-link>
        </v-col>
      </v-row>
  </v-container>
</template>

<script>
import firebase from "firebase";
import Heading from "../components/Heading.vue";

export default {
  name: "Register",
  components: {
    Heading,
  },
  data() {
    return {
      email: "",
      password: "",
    };
  },
  methods: {
    signUp() {
      firebase
        .auth()
        .createUserWithEmailAndPassword(this.email, this.password)
        .then(
          () => {
            this.$router.replace("home");
          },
          (err) => {
            alert("Oops! " + err.message);
          }
        );
    },
  },
};
</script>

<style scoped></style>
