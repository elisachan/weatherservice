<template>
  <div>
    <v-row justify="center">
      <v-col>
        <h1>what the weather</h1>
      </v-col>
    </v-row>
    <v-row class="text-center">
      <v-col cols="12">
        <v-img
          v-show="weather == 'everything'"
          :src="require('../assets/weather.svg')"
          class="my-3"
          contain
          height="200"
        />
        <v-img
          v-show="weather == 'Clouds'"
          :src="require('../assets/cloudy.svg')"
          class="my-3"
          contain
          height="200"
        />
        <v-img
          v-show="weather == 'Rain'"
          :src="require('../assets/rainy-7.svg')"
          class="my-3"
          contain
          height="200"
        />
        <v-img
          v-show="weather == 'Clear'"
          :src="require('../assets/day.svg')"
          class="my-3"
          contain
          height="200"
        />
      </v-col>
    </v-row>
  </div>
</template>

<script>
export default {
  props: {
    weather: {
      type: String,
      default: "everything",
    },
  },
};
</script>
