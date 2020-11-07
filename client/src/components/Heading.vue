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
          :src="require('../assets/weather.svg')"
          class="my-3"
          contain
          height="200"
        />
      </v-col>
    </v-row>
  </div>
</template>

<script>
export default {};
</script>

<style lang="scss" scoped></style>
