<template>
  <div>
    <vue-p5
        @setup=setup
        @draw=draw></vue-p5>

    <b-container>{{obstacleCoord}}</b-container>

  </div>
</template>

<script>
import p5 from "vue-p5";
import axios  from "axios";

export default {
  name: "VueCanvas",
  components: {
    "vue-p5": p5
  },

  data(){
    return{
      pathGrid:null,
      canvasSize: 600,
      arenaSize: 200,
      resolution: 20,

    }
  },
  props:{obstacleCoord: Array},

  methods:{
    async getData() {
      try {
        const response = await axios.post(
            "http://localhost:3000/hamilton_path"
        );
        // JSON responses are automatically parsed.
        this.pathGrid = response.data;
        console.log(this.pathGrid);
      } catch (error) {
        console.log(error);
      }
    },

    setup: function (canvas ){
      canvas.createCanvas(this.canvasSize, this.canvasSize);
      //canvas.parent("vue-canvas");
      //p5.createGrid();
      //canvas.position(50, 5);
      console.log("hi");
      console.log(this.obstacleCoord);
    },

    draw: function(canvas) {
      canvas.background('lightblue');
      //Draw Grid
      for (let x = 0; x < this.canvasSize; x += this.canvasSize / this.resolution) {
        for (let y = 0; y < this.canvasSize; y += this.canvasSize / this.resolution) {
          canvas.stroke(0);
          canvas.strokeWeight(0.05);
          canvas.line(x, 0, x, this.canvasSize);
          canvas.line(0, y, this.canvasSize, y);
        }
      }
    }
  },

  mounted() {
    this.getData();
  }

}
</script>

<style scoped>

</style>
