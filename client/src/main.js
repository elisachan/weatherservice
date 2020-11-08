import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import vuetify from "./plugins/vuetify";
import firebase from "firebase";

Vue.config.productionTip = false;

let app = "";
const config = {
  apiKey: "AIzaSyBaxD7KMjRUwvxIKOoRelKRmTkOIKgX1uM",
  authDomain: "thinkific-weather-3bbfb.firebaseapp.com",
  databaseURL: "https://thinkific-weather-3bbfb.firebaseio.com",
  projectId: "thinkific-weather-3bbfb",
  storageBucket: "thinkific-weather-3bbfb.appspot.com",
  messagingSenderId: "572598306782",
  appId: "1:572598306782:web:017bc9ccfeb70c7a4e3f4f",
  measurementId: "G-XHCKRK5RPF",
};

firebase.initializeApp(config);

firebase.auth().onAuthStateChanged(() => {
  if (!app) {
    /* eslint-disable no-new */
    app = new Vue({
      router,
      vuetify,
      render: (h) => h(App),
    }).$mount("#app");
  }
});
