<template>
  <div>
    <div id="vue-canvas">
      <h3>{{noOfObstacles}}</h3>
    </div>
    <h3></h3>
  </div>
</template>

<script>
export default {
  name: "VueCanvas",

  data() {
    return {
      noOfObstacles: 0,
      direction: null,
      coordX: 0,
      coordY: 0
    }
  },

  mounted() {
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
        canvas.position(50, 50);
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

        /*if(p5.mouseIsPressed){
          p5.triggerMouse(p5.mouseX, p5.mouseY);
        }*/


        //const degree = p5.frameCount * 3;
        //const y = p5.sin(p5.radians(degree)) * 50;

        //p5.push();
        //p5.translate(0, p5.height / 2);
        //p5.fill(66, 184, 131);
        //p5.stroke(53, 73, 94);
        //p5.strokeWeight(5);
        //p5.ellipse(posX, y, 50, 50);
        //p5.pop();
        /*posX += speed;

        if (posX > p5.width - 35 || posX < 35) {
          speed *= -1;
        }*/
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

          /*for(var i = 0; i < grid.length; i++)
          {
            if(p5.mouseX >= grid[i].x && p5.mouseX <= grid[i].x+grid[i].size)
            {
              if(p5.mouseY >= grid[i].y && p5.mouseY <= grid[i].y+grid[i].size)
              {
                grid[i].trigger();
              }
            }
          }*/
        }
      }

      /*p5.showGrid = () => {
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
      }*/

    };
    // NOTE: Use p5 as an instance mode
    const P5 = require("p5");
    new P5(script);
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
