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
      algoReceived: false,
      pathGrid: null,
      count:0
    }
  },


  computed:{

  },

  methods: {
    sendData: async function(){
      await axios.post('http://localhost:3000/set_arena', {arena: this.gridArray});

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

    setup: function (canvas ){
      canvas.createCanvas(this.canvasSize, this.canvasSize);
      //canvas.parent("vue-canvas");
      //p5.createGrid();
      //canvas.position(50, 5);
      console.log("hi");
    },

    draw: function(canvas){
      canvas.background('lightblue');
      canvas.frameRate(3);
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

      //let i = 0;

      if (this.pathGrid === null){
        canvas.fill("lightgreen");
        canvas.rect(0, 510, 90, 90);
        canvas.fill("darkgreen");
        canvas.ellipse(45, 525, 30, 30)
      } else{
        let n = JSON.parse(JSON.stringify(this.pathGrid)).length;

        if (this.count < n) {

          let xvalue = JSON.parse(JSON.stringify(this.pathGrid))[this.count].cell.x - 1;
          let yvalue = JSON.parse(JSON.stringify(this.pathGrid))[this.count].cell.y - 1;
          let headX
          let headY

          //Account head direction
          let headD = JSON.parse(JSON.stringify(this.pathGrid))[this.count].direction;
          switch (headD){
            case 0: { //East
              headX = (xvalue*30) + 75;
              headY = (yvalue*30) + 45;
              break;
            }
            case 1: {// North
              headX = (xvalue * 30) + 45;
              headY = (yvalue * 30) + 15;
              break;
            }
            case 2: { // West
              headX = (xvalue*30) + 15;
              headY = (yvalue*30) + 45;
              break;
            }
            case 3:{ // South
              headX = (xvalue*30) + 45;
              headY = (yvalue*30) + 75;
              break;
            }
          }

          // draw ellipsis to show head
          canvas.fill("lightgreen");
          canvas.rect(xvalue * 30, yvalue * 30, 90, 90);
          canvas.fill("darkgreen");
          canvas.ellipse(headX, headY, 30, 30);
          this.count++;
        } else{
          this.count =0;
        }
      }

      //Draw Obstacles
      // Draw Direction
      ////Obstacle 1
      let X1 = (parseInt(this.obstacle1.column) -1);
      let Y1 = (parseInt(this.obstacle1.row) -1);

      canvas.fill(155);
      canvas.rect(X1*30, Y1*30, 30, 30);

      ////Obstacle 2
      let X2 = (parseInt(this.obstacle2.column) -1);
      let Y2 = (parseInt(this.obstacle2.row) -1);

      canvas.fill(155);
      canvas.rect(X2*30, Y2*30, 30, 30);

      ////Obstacle 3
      let X3 = (parseInt(this.obstacle3.column) -1);
      let Y3 = (parseInt(this.obstacle3.row) -1);

      canvas.fill(155);
      canvas.rect(X3*30, Y3*30, 30, 30);

      ////Obstacle 4
      let X4 = (parseInt(this.obstacle4.column) -1);
      let Y4 = (parseInt(this.obstacle4.row) -1);

      canvas.fill(155);
      canvas.rect(X4*30, Y4*30, 30, 30);

      ////Obstacle 5
      let X5 = (parseInt(this.obstacle5.column) -1);
      let Y5 = (parseInt(this.obstacle5.row) -1);

      canvas.fill(155);
      canvas.rect(X5*30, Y5*30, 30, 30);

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

      /*let data = [];
      data.push(this.obstacle1);
      data.push(this.obstacle2);
      data.push(this.obstacle3);
      data.push(this.obstacle4);
      data.push(this.obstacle5);
      this.obstacleCoord = data;*/

      this.sendData();
      this.getData();
      this.pathGrid = JSON.parse(JSON.stringify(this.pathGrid));
      //console.log(this.pathGrid);
      //console.log(this.obstacleCoord);

    },
  },

  mounted() {
    //this.getData();
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
