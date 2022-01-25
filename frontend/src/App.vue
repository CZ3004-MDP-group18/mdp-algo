<template>
  <div id="app">

    <router-view/>
  </div>
</template>

<script>
//import VueCanvas from "./Stimulation.vue";

export default {
  name: "App",
  components: {
    //VueCanvas
  }
};
</script>

<style>
body,
html {
  margin: 0;
  padding: 0;
}

#app {
  margin-top: 60px;
  color: #2c3e50;
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  text-align: center;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
</style>
