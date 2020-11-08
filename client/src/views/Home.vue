<template>
  <v-container>
    <Heading :weather="weatherHeading"/>
    <v-row>
      <v-text-field
        solo
        type="text"
        v-model="city"
        placeholder="Enter your City"
      ></v-text-field>
    </v-row>
    <v-row>
      <h5>{{weatherInfo}}</h5>
    </v-row>
    <v-btn large @click="submitCity">Submit for weather data</v-btn>
    <v-row class="text-center">
      <v-col cols="12">
        <v-btn @click="logout">Logout</v-btn>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import firebase from "firebase";
import Heading from "../components/Heading.vue";
// import weatherservice from "../services/weather.js"
export default {
  name: "Home",
  components: {
    Heading,
  },
  data() {
    return {
      city: "",
      weatherInfo: "",
      weatherHeading:"everything",
      baseURL: "http://localhost:8082" // used for dev 
    };
  },
  methods: {
    submitCity() {
      fetch(`${this.baseURL}/getweather?city=${this.city}`)
      .then(response => response.json())
      .then(data =>{
        this.weatherInfo = data
        this.weatherHeading = data.weather[0].main
      })
    },

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
