<template>
  <div id="app">
    <ejs-diagram id="diagram"  :width='width' :height='height' :nodes='nodes' ></ejs-diagram>
  </div>
</template>

<script>
import Vue from 'vue';
import { DiagramPlugin } from '@syncfusion/ej2-vue-diagrams';

Vue.use(DiagramPlugin);
let nodes = [
  {
    id: "node1",
    width: 50,
    height: 50,
    offsetX: 500,
    offsetY: 800,
    annotations: [
      {
        content: "start"
      }
    ]
  }
]
export default {
  name: 'app',
  data () {
    return {
      width: "1000",
      height: "1000",
      nodes: nodes
    }
  }
}
</script>

<style>
@import "../node_modules/@syncfusion/ej2-vue-diagrams/styles/material.css";
</style>
