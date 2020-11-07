<template>
  <v-container>
   <Heading />
    <v-row class="text-center">
      <v-col cols="12">
        <v-btn @click="logout">Logout</v-btn>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>

import firebase from 'firebase'
import Heading from '../components/Heading.vue'
export default {
  name: "Home",
  components: {
    Heading
  },
  data() {
    return {
      email: "",
    };
  },
  methods: {
    logout() {
      firebase
        .auth()
        .signOut()
        .then(() => {
          this.$router.replace("login");
        });
    },
  },
};
</script>
