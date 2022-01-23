<template>
  <div>
    <h2>Stimulated Arena</h2>
    <div>
      <vue-p5
          @setup=setup
          @draw=draw></vue-p5>
    </div>

    <b-container>
      <b-row>
        <b-col>No of Obstacles:</b-col>
        <b-col>
          <b-form-input v-model="noOfObstacles" placeholder="Enter number of obstacles"></b-form-input>
        </b-col>
      </b-row>

      <b-card><h4>Obstacle 1</h4>
        <b-form-input v-model="obstacle1.row" placeholder="Row number"></b-form-input>
        <b-form-input v-model="obstacle1.column" placeholder="Column number"></b-form-input>
        <b-form-input list="my-list-id" v-model = "obstacle1.direction" placeholder="Direction"></b-form-input>
        <b-card-text >Row: {{ obstacle1.row }}, Column: {{obstacle1.column}}, Direction: {{obstacle1.direction}}</b-card-text>
      </b-card>

      <b-card><h4>Obstacle 2</h4>
        <b-form-input v-model="obstacle2.row" placeholder="Row number"></b-form-input>
        <b-form-input v-model="obstacle2.column" placeholder="Column number"></b-form-input>

        <b-form-input list="my-list-id" v-model = "obstacle2.direction" placeholder="Direction"></b-form-input>
        <b-card-text >Row: {{ obstacle2.row }}, Column: {{obstacle2.column}}, Direction: {{obstacle2.direction}}</b-card-text>
      </b-card>

      <b-card><h4>Obstacle 3</h4>
        <b-form-input v-model="obstacle3.row" placeholder="Row number"></b-form-input>
        <b-form-input v-model="obstacle3.column" placeholder="Column number"></b-form-input>

        <b-form-input list="my-list-id" v-model = "obstacle3.direction" placeholder="Direction"></b-form-input>
        <b-card-text >Row: {{ obstacle3.row }}, Column: {{obstacle3.column}}, Direction: {{obstacle3.direction}}</b-card-text>
      </b-card>

      <b-card><h4>Obstacle 4</h4>
        <b-form-input v-model="obstacle4.row" placeholder="Row number"></b-form-input>
        <b-form-input v-model="obstacle4.column" placeholder="Column number"></b-form-input>

        <b-form-input list="my-list-id" v-model = "obstacle4.direction" placeholder="Direction"></b-form-input>
        <b-card-text >Row: {{ obstacle4.row }}, Column: {{obstacle4.column}}, Direction: {{obstacle4.direction}}</b-card-text>
      </b-card>

      <b-card><h4>Obstacle 5</h4>
        <b-form-input v-model="obstacle5.row" placeholder="Row number"></b-form-input>
        <b-form-input v-model="obstacle5.column" placeholder="Column number"></b-form-input>

        <b-form-input list="my-list-id" v-model = "obstacle5.direction" placeholder="Direction"></b-form-input>
        <b-card-text >Row: {{ obstacle5.row }}, Column: {{obstacle5.column}}, Direction: {{obstacle5.direction}}</b-card-text>
      </b-card>

      <b-card>
        <b-button variant = "info" @click="updateGrid(gridArray)">Send Data for Algo</b-button>
      </b-card>

    </b-container>

    <datalist id="my-list-id">
      <option
          v-for="direction in direction" :key="direction">{{ direction }}
      </option>
    </datalist>

    <div id="vue-canvas"></div>
  </div>
</template>

<script>
import p5 from "vue-p5";
export default {
  name: "VueCanvas",
  components: {
    "vue-p5": p5
  },

  data() {
    return {
      noOfObstacles: 5,
      direction: ["North", "South", "East", "West"],
      obstacle1: {
        row:null,
        column:null,
        direction: null
      },
      obstacle2: {
        row:null ,
        column:null,
        direction: null
      },
      obstacle3: {
        row:null ,
        column:null,
        direction: null
      },
      obstacle4: {
        row:null ,
        column:null,
        direction: null
      },
      obstacle5: {
        row:null ,
        column:null,
        direction: null

      },
      canvasSize: 600,
      arenaSize: 200,
      resolution: 20,
      gridArray: [
    ["O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"],
    ["O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"],
    ["O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"],
    ["O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"],
    ["O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"],
    ["O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"],
    ["O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"],
    ["O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"],
    ["O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"],
    ["O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"],
    ["O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"],
    ["O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"],
    ["O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"],
    ["O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"],
    ["O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"],
    ["O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"],
    ["O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"],
    ["O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"],
    ["O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"],
    ["O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"],]
    }
  },

  computed:{
    grid(){
      return this.gridArray;
    }
  },

  methods: {
    sendData: function(){

    },


    setup: function (canvas){
      canvas.createCanvas(this.canvasSize, this.canvasSize);
      //canvas.parent("vue-canvas");
      //p5.createGrid();
      canvas.position(50, 5);
    },

    draw: function(canvas){
      canvas.background('lightblue');
      //Draw Grid
      for (var x = 0; x < this.canvasSize; x += this.canvasSize / this.resolution) {
        for (var y = 0; y < this.canvasSize; y += this.canvasSize / this.resolution) {
          canvas.stroke(0);
          canvas.strokeWeight(0.05);
          canvas.line(x, 0, x, this.canvasSize);
          canvas.line(0, y, this.canvasSize, y);
        }
      }

      //Draw Robot
      canvas.fill("lightgreen");
      canvas.rect(0, 510, 90, 90);

      //Draw Obstacles
      ////Obstacle 1
      let X1 = (parseInt(this.obstacle1.row) -1);
      let Y1 = (parseInt(this.obstacle1.column) -1);

      X1 = X1 * 30;
      Y1 = Y1 * 30;


      canvas.fill(155);
      canvas.rect(X1, Y1, 30, 30);

      ////Obstacle 2
      let X2 = (parseInt(this.obstacle2.row) -1);
      let Y2 = (parseInt(this.obstacle2.column) -1);

      X2 = X2 * 30;
      Y2 = Y2 * 30;


      canvas.fill(155);
      canvas.rect(X2, Y2, 30, 30);

      ////Obstacle 3
      let X3 = (parseInt(this.obstacle3.row) -1);
      let Y3 = (parseInt(this.obstacle3.column) -1);

      X3 = X3 * 30;
      Y3 = Y3 * 30;


      canvas.fill(155);
      canvas.rect(X3, Y3, 30, 30);

      ////Obstacle 4
      let X4 = (parseInt(this.obstacle4.row) -1);
      let Y4 = (parseInt(this.obstacle4.column) -1);

      X4 = X4 * 30;
      Y4 = Y4 * 30;


      canvas.fill(155);
      canvas.rect(X4, Y4, 30, 30);

      ////Obstacle 5
      let X5 = (parseInt(this.obstacle5.row) -1);
      let Y5 = (parseInt(this.obstacle5.column) -1);

      X5 = X5 * 30;
      Y5 = Y5 * 30;


      canvas.fill(155);
      canvas.rect(X5, Y5, 30, 30);

    },

    updateGrid: function (gridArray) {

      // Change Obstacle 1
      let row1 = this.obstacle1.row - 1;
      let col1 = this.obstacle1.column - 1;
      gridArray[row1][col1] = this.obstacle1.direction.charAt(0);

      // Change Obstacle 2
      let row2 = this.obstacle2.row - 1;
      let col2 = this.obstacle2.column - 1;
      gridArray[row2][col2] = this.obstacle2.direction.charAt(0);

      // Change Obstacle 3
      let row3 = this.obstacle3.row - 1;
      let col3 = this.obstacle3.column - 1;
      gridArray[row3][col3] = this.obstacle3.direction.charAt(0);

      // Change Obstacle 4
      let row4 = this.obstacle4.row - 1;
      let col4 = this.obstacle4.column - 1;
      gridArray[row4][col4] = this.obstacle4.direction.charAt(0);

      // Change Obstacle 5
      let row5 = this.obstacle5.row - 1;
      let col5 = this.obstacle5.column - 1;
      gridArray[row5][col5] = this.obstacle5.direction.charAt(0);
    }
  },

  /*mounted() {
    const script = function (p5) {
      //var speed = 2;
      //var posX = 35;
      let canvasSize = 600;
      let arenaSize =200;
      let scaleFactor = canvasSize/arenaSize;
      let resolution = 20;
      let color = "lightblue";
      //let padding = arenaSize/resolution;

      //let grid = [];
      //let rows = 0;
      //let nextRow = 0;
      let canvas;

      // NOTE: Set up is here
      p5.setup = () => {
        canvas = p5.createCanvas(canvasSize, canvasSize);
        canvas.parent("vue-canvas");
        //p5.createGrid();
        canvas.position(50, 100);
      };
      // NOTE: Draw is here
      p5.draw = () => {
        p5.background('lightblue');
        //p5.showGrid();
        for (var x = 0; x < canvasSize; x += canvasSize / resolution) {
          for (var y = 0; y < canvasSize; y += canvasSize / resolution) {
            p5.stroke(0);
            p5.strokeWeight(1);
            p5.line(x, 0, x, canvasSize);
            p5.line(0, y, canvasSize, y);
          }
        }

        //p5.fill(155)
        //p5.rect(120,120,30,30);
        // Plot obstacle
        this.mouseClicked();

        /!*if(p5.mouseIsPressed){
          p5.triggerMouse(p5.mouseX, p5.mouseY);
        }*!/


        //const degree = p5.frameCount * 3;
        //const y = p5.sin(p5.radians(degree)) * 50;

        //p5.push();
        //p5.translate(0, p5.height / 2);
        //p5.fill(66, 184, 131);
        //p5.stroke(53, 73, 94);
        //p5.strokeWeight(5);
        //p5.ellipse(posX, y, 50, 50);
        //p5.pop();
        /!*posX += speed;

        if (posX > p5.width - 35 || posX < 35) {
          speed *= -1;
        }*!/
      };

      //Mouse clicking Event
      p5.mouseClicked = () => {
        //p5.triggerMouse(p5.mouseX, p5.mouseY);
        if (color !== 155 ){
          color = 155;
          //this.noOfObstacles +=1;
        } else{
          color = "lightblue";
          //this.noOfObstacles -=1;
        }
        p5.triggerMouse(p5.mouseX, p5.mouseY);
      }

      // Create obstacle object
      p5.triggerMouse = (xvalue, yvalue) => {
        let x = xvalue/ scaleFactor;
        let y = yvalue/ scaleFactor;

        x = parseInt(x/10, 10) * 10 *scaleFactor
        y = parseInt(y/10, 10) * 10 *scaleFactor

        if(xvalue <= canvas.width && yvalue <= canvas.height)
        {
          p5.fill(color)
          p5.rect(x, y, 30, 30);

          /!*for(var i = 0; i < grid.length; i++)
          {
            if(p5.mouseX >= grid[i].x && p5.mouseX <= grid[i].x+grid[i].size)
            {
              if(p5.mouseY >= grid[i].y && p5.mouseY <= grid[i].y+grid[i].size)
              {
                grid[i].trigger();
              }
            }
          }*!/
        }
      }

      /!*p5.showGrid = () => {
        for(var i = 0; i < grid.length; i++)
        {
          grid[i].show();
        }
      };

      p5.mouseDragged = () => { p5.triggerMouse(); };

      p5.Pixel = () => {
        this.size = canvasSize / resolution;
        this.color = 255;

        this.x = nextRow * this.size;
        this.y = rows * this.size;

        grid.push(this);

        if(grid.length % resolution === 0) rows++;

        nextRow++;
        if(nextRow === resolution) nextRow = 0;

        this.show = function() {
          p5.stroke(0);
          p5.strokeWeight(2);
          p5.fill(this.color);
          p5.square(this.x, this.y, this.size);
        }
      };

      p5.createGrid = () => {
        for (var i = 0; i < resolution * resolution; i++) {
          new p5.Pixel();
        }
      }*!/

    };
    // NOTE: Use p5 as an instance mode
    const P5 = require("p5");
    new P5(script);
  }*/
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
#vue-canvas {
  display: block;
  margin: 0 auto;
  padding: 0;
  width: 500px;
  height: 500px;
  border-radius: 20px;
  overflow: hidden;
}
</style>
