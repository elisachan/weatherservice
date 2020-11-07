<template>
  <v-container>
    <h1>what the weather</h1>
    <v-row class="text-center">
      <v-col cols="12">
        <v-img
          :src="require('../assets/weather.svg')"
          class="my-3"
          contain
          height="200"
        />
      </v-col>

      <v-col class="mb-4">
     
      </v-col>




    </v-row>
  </v-container>
</template>

<script>
  export default {
    name: 'WeatherSearch'
  }
</script>
