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

      <b-card>
        <b-button pill variant = "outline-secondary"
                  :pressed.sync = "mapDescriptor"
                  @click = "setInput">Map Descriptor</b-button>
        <b-button pill variant = "outline-secondary"
                  :pressed.sync = "selfInput"
                  @click = "setInput">Self Input</b-button>
      </b-card>


      <b-container v-if="mapDescriptor">
        <textarea v-model="descriptor" placeholder="Enter grid with obstacle direction" style="width:500px; height: 300px"></textarea>
        <b-card>
<!--          <b-button variant = "outline-primary" @click="getObstacle">Plot Obstacles</b-button>-->
          <b-button variant = "info" @click="updateGrid(gridArray)">Plot Obstacles</b-button>
          <b-button variant = "info" @click="getPath">Get Path</b-button>
        </b-card>
      </b-container>

      <b-container v-if="selfInput">

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
        <b-button variant = "info" @click="getPath">Get Path</b-button>
      </b-card>
      </b-container>

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
import axios  from "axios";
export default {
  name: "VueCanvas",
  components: {
    "vue-p5": p5,

  },

  data() {
    return {
      noOfObstacles: 5,
      direction: ["North", "South", "East", "West"],
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
    ["O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"],
  ],
      // Recieve Actual algo from backend
      pathGrid: null,
      count:0,

      // Descriptor Alg0
      mapDescriptor: false,
      descriptor: null,
      descriptorObstacle: [],

      // Self input plot
      selfInput: false,
      obstacle1: {
          row: null,
          column: null,
          direction: null
        },
        obstacle2: {
          row: null,
          column: null,
          direction: null
        },
        obstacle3: {
          row: null,
          column: null,
          direction: null
        },
        obstacle4: {
          row: null,
          column: null,
          direction: null
        },
        obstacle5: {
          row: null,
          column: null,
          direction: null
        },
    }
  },

  methods: {
    sendData: async function(){
      if (this.selfInput) {
        await axios.post('http://localhost:3000/set_arena', {arena: this.gridArray});
      }

      if (this.mapDescriptor){
        await axios.post('http://localhost:3000/set_arena', {arena: JSON.parse(JSON.stringify(this.gridArray))});
      }

      //let data = res.data;
      //console.log(data);
    },

    async getData() {
      try {
        const response = await axios.post(
            "http://localhost:3000/hamilton_path"
        );
        // JSON responses are automatically parsed.
        console.log(response.data);
        var grid =  response.data.path;
        console.log(grid);
        //return grid;
        this.pathGrid = grid;
        JSON.parse(JSON.stringify(this.pathGrid));
        console.log(this.pathGrid);
      } catch (error) {
        console.log(error);
      }
    },

    setup: function (canvas){
      canvas.createCanvas(this.canvasSize, this.canvasSize);
      //canvas.parent("vue-canvas");
      //p5.createGrid();
      //canvas.position(50, 5);
      console.log("hi");
    },

    draw: function(canvas){
      canvas.background('lightblue');
      canvas.frameRate(2.5);
      //Draw Grid
      for (let x = 0; x < this.canvasSize; x += this.canvasSize / this.resolution) {
        for (let y = 0; y < this.canvasSize; y += this.canvasSize / this.resolution) {
          canvas.stroke(0);
          canvas.strokeWeight(0.05);
          canvas.line(x, 0, x, this.canvasSize);
          canvas.line(0, y, this.canvasSize, y);
        }
      }

      //Draw Robot
      //canvas.fill("lightgreen");
      //canvas.rect(0, 510, 90, 90);

      if (this.pathGrid === null){
        canvas.fill("lightgreen");
        canvas.rect(0, 510, 90, 90);
        canvas.fill("darkgreen");
        canvas.ellipse(45, 525, 30, 30)
      } else{
        let n = JSON.parse(JSON.stringify(this.pathGrid)).length;


          if (this.count < n) {

            let xvalue = (JSON.parse(JSON.stringify(this.pathGrid))[this.count].cell.x - 1) * 30;
            let yvalue = (JSON.parse(JSON.stringify(this.pathGrid))[this.count].cell.y - 1) * 30;

            let headX;
            let headY;

            //Account head direction
            let headD = JSON.parse(JSON.stringify(this.pathGrid))[this.count].direction;
            switch (headD) {
              case 0: { //East
                headX = (xvalue) + 75;
                headY = (yvalue) + 45;
                break;
              }
              case 1: {// North
                headX = (xvalue) + 45;
                headY = (yvalue) + 15;
                break;
              }
              case 2: { // West
                headX = (xvalue) + 15;
                headY = (yvalue) + 45;
                break;
              }
              case 3: { // South
                headX = (xvalue) + 45;
                headY = (yvalue) + 75;
                break;
              }
            }

            // draw ellipsis to show head
            canvas.fill("lightgreen");
            canvas.rect(xvalue, yvalue, 90, 90);
            canvas.fill("darkgreen");
            canvas.ellipse(headX, headY, 30, 30);
            this.count++;
          } else {
            this.count = 0;
          }
        }

      //Draw Obstacles
      // Draw Direction
      if (this.selfInput) {

        ////Obstacle 1
        let X1 = (parseInt(this.obstacle1.column) - 1) * 30;
        let Y1 = (parseInt(this.obstacle1.row) - 1) * 30;
        let dir1 = this.obstacle1.direction;

        let headX1;
        let headY1;

        switch (dir1) {
          case 'East': { //East
            headX1 = (X1) + 25;
            headY1 = (Y1) + 15;
            break;
          }
          case 'North': {// North
            headX1 = (X1) + 15;
            headY1 = (Y1) + 5;
            break;
          }
          case'West': { // West
            headX1 = (X1) + 5;
            headY1 = (Y1) + 15;
            break;
          }
          case 'South': { // South
            headX1 = (X1) + 15;
            headY1 = (Y1) + 25;
            break;
          }
        }

        canvas.fill(155);
        canvas.rect(X1, Y1, 30, 30);
        canvas.fill(55);
        canvas.ellipse(headX1, headY1, 10, 10);

        ////Obstacle 2
        let X2 = (parseInt(this.obstacle2.column) - 1) * 30;
        let Y2 = (parseInt(this.obstacle2.row) - 1) * 30;
        let dir2 = this.obstacle2.direction;

        let headX2;
        let headY2;

        switch (dir2) {
          case 'East': { //East
            headX2 = (X2) + 25;
            headY2 = (Y2) + 15;
            break;
          }
          case 'North': {// North
            headX2 = (X2) + 15;
            headY2 = (Y2) + 5;
            break;
          }
          case'West': { // West
            headX2 = (X2) + 5;
            headY2 = (Y2) + 15;
            break;
          }
          case 'South': { // South
            headX2 = (X2) + 15;
            headY2 = (Y2) + 25;
            break;
          }
        }

        canvas.fill(155);
        canvas.rect(X2, Y2, 30, 30);
        canvas.fill(55);
        canvas.ellipse(headX2, headY2, 10, 10);

        ////Obstacle 3
        let X3 = (parseInt(this.obstacle3.column) - 1) * 30;
        let Y3 = (parseInt(this.obstacle3.row) - 1) * 30;
        let dir3 = this.obstacle3.direction;

        let headX3;
        let headY3;

        switch (dir3) {
          case 'East': { //East
            headX3 = (X3) + 25;
            headY3 = (Y3) + 15;
            break;
          }
          case 'North': {// North
            headX3 = (X3) + 15;
            headY3 = (Y3) + 5;
            break;
          }
          case'West': { // West
            headX3 = (X3) + 5;
            headY3 = (Y3) + 15;
            break;
          }
          case 'South': { // South
            headX3 = (X3) + 15;
            headY3= (Y3) + 25;
            break;
          }
        }

        canvas.fill(155);
        canvas.rect(X3, Y3, 30, 30);
        canvas.fill(55);
        canvas.ellipse(headX3, headY3, 10, 10);

        ////Obstacle 4
        let X4 = (parseInt(this.obstacle4.column) - 1) * 30;
        let Y4 = (parseInt(this.obstacle4.row) - 1) * 30;
        let dir4 = this.obstacle4.direction;

        let headX4;
        let headY4;

        switch (dir4) {
          case 'East': { //East
            headX4 = (X4) + 25;
            headY4 = (Y4) + 15;
            break;
          }
          case 'North': {// North
            headX4 = (X4) + 15;
            headY4 = (Y4) + 5;
            break;
          }
          case'West': { // West
            headX4 = (X4) + 5;
            headY4 = (Y4) + 15;
            break;
          }
          case 'South': { // South
            headX4 = (X4) + 15;
            headY4= (Y4) + 25;
            break;
          }
        }

        canvas.fill(155);
        canvas.rect(X4, Y4, 30, 30);
        canvas.fill(55);
        canvas.ellipse(headX4, headY4, 10, 10);

        ////Obstacle 5
        let X5 = (parseInt(this.obstacle5.column) - 1) * 30;
        let Y5 = (parseInt(this.obstacle5.row) - 1) * 30;
        let dir5 = this.obstacle5.direction;

        let headX5;
        let headY5;

        switch (dir5) {
          case 'East': { //East
            headX5 = (X5) + 25;
            headY5 = (Y5) + 15;
            break;
          }
          case 'North': {// North
            headX5 = (X5) + 15;
            headY5 = (Y5) + 5;
            break;
          }
          case 'West': { // West
            headX5 = (X5) + 5;
            headY5 = (Y5) + 15;
            break;
          }
          case 'South': { // South
            headX5 = (X5) + 15;
            headY5= (Y5) + 25;
            break;
          }
        }

        canvas.fill(155);
        canvas.rect(X5, Y5, 30, 30);
        canvas.fill(55);
        canvas.ellipse(headX5, headY5, 10, 10);
      }

      if (this.mapDescriptor){

        for (let i =0; i< JSON.parse(JSON.stringify(this.descriptorObstacle)).length; i++){
          let xvalue = (JSON.parse(JSON.stringify(this.descriptorObstacle))[i][0]) * 30;
          let yvalue = (JSON.parse(JSON.stringify(this.descriptorObstacle))[i][1]) * 30;
          let obstacleD = JSON.parse(JSON.stringify(this.descriptorObstacle))[i][2];

          let headX;
          let headY;

          switch (obstacleD) {
            case 'E': { //East
              headX = (xvalue) + 25;
              headY = (yvalue) + 15;
              break;
            }
            case 'N': {// North
              headX = (xvalue) + 15;
              headY = (yvalue) + 5;
              break;
            }
            case'W': { // West
              headX = (xvalue) + 5;
              headY = (yvalue) + 15;
              break;
            }
            case 'S': { // South
              headX = (xvalue) + 15;
              headY = (yvalue) + 25;
              break;
            }
          }

          canvas.fill(155);
          canvas.rect(xvalue, yvalue, 30, 30);
          canvas.fill(55);
          canvas.ellipse(headX, headY, 10, 10);
        }

      }

    },

    updateGrid: function (gridArray) {

      if (this.selfInput) {


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

      if (this.mapDescriptor){
        const splitString = this.descriptor.replace(/(\r\n|\n|\r| )/gm, "").split("");
        console.log(splitString);

        const size = 20;
        console.assert(splitString.length === size * size, "Map Descriptor Length is not correct");

        const arrayOfArrays = [];
        for (var i=0; i<splitString.length; i+=size) {
          arrayOfArrays.push(splitString.slice(i,i+size));
        }
        console.log(arrayOfArrays);

        this.gridArray = arrayOfArrays;
        console.log(JSON.parse(JSON.stringify(this.gridArray)));

        for (let i = 0; i<JSON.parse(JSON.stringify(this.gridArray)).length; i++){
          for(let j = 0; j<JSON.parse(JSON.stringify(this.gridArray)).length; j++){
            let obstacle = [];
            let D = JSON.parse(JSON.stringify(this.gridArray))[i][j];
            if (D === 'N'|| D==='S' || D === 'E' || D==='W') {
              obstacle.push(j); // X
              obstacle.push(i); // Y
              obstacle.push(D); // Direction
              this.descriptorObstacle.push(obstacle);
            }

          }
        }
        console.log(this.descriptorObstacle);
      }

      this.sendData();
      //this.getData();
      //this.pathGrid = JSON.parse(JSON.stringify(this.pathGrid));

    },

    getPath: function(){
      this.getData();
    },

    setInput: function(){
      if (this.mapDescriptor){
        this.selfInput = false;
      }
      if(this.selfInput){
        this.mapDescriptor = false;
      }
    }

  },

  mounted() {

  }
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
