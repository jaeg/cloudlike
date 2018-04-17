var c=document.getElementById("game");
var ctx=c.getContext("2d");

var sock = null;
var wsuri = "ws://127.0.0.1:1234";

var tileWidth = 16;
var tileHeight = 24;

var tileMultiplier = 2;


sock = new WebSocket(wsuri);

sock.onopen = function() {
    console.log("connected to " + wsuri);
     sock.send(JSON.stringify({type:"viewSize",width:c.width/(tileWidth*tileMultiplier),height:c.height/(tileHeight*tileMultiplier)}))
}

sock.onclose = function(e) {
    console.log("connection closed (" + e.code + ")");
    sock = null
}
drawBuffer = []
sock.onmessage = function(e) {
    console.log(e.data);
    var payload = e.data;
    var index = payload.indexOf(":");
    var type = payload.substr(0,index);
    var json = payload.substr(index+1);

    var data = JSON.parse(json)
    if (type == "view") {
      drawBuffer = data
    }
}

function handleKeyPressed(e) {
  sock.send(JSON.stringify({type:"input",key:String.fromCharCode(e.keyCode)}));
}

document.addEventListener("keydown", handleKeyPressed, false);


var fps = {	startTime : 0,	frameNumber : 0,	getFPS : function(){		this.frameNumber++;		var d = new Date().getTime(),			currentTime = ( d - this.startTime ) / 1000,			result = Math.floor( ( this.frameNumber / currentTime ) );		if( currentTime > 1 ){			this.startTime = new Date().getTime();			this.frameNumber = 0;		}		return result;	}	};
var fpsDiv = document.getElementById("fps");
//Main Loop
var draw = function() {
  fpsDiv.innerHTML = "FPS:" + fps.getFPS();
  ctx.clearRect(0, 0, c.width, c.height);
  ctx.imageSmoothingEnabled = false;
  if (drawBuffer != null) {
    for (var x = 0; x < drawBuffer.length; x++) {
      for (var y = 0; y < drawBuffer[x].length; y++) {
        if (drawBuffer[x][y] == null) {
          ctx.fillStyle = "black"
          ctx.fillRect(x*tileWidth*tileMultiplier,y*tileHeight*tileMultiplier,tileWidth*tileMultiplier,tileHeight*tileMultiplier);
        } else {
          var tileType = drawBuffer[x][y].TileType
          if (tileType === 1)
            ctx.fillStyle = "green";
          if (tileType === 2)
            ctx.fillStyle = "blue";
          ctx.fillRect(x*tileWidth*tileMultiplier,y*tileHeight*tileMultiplier,tileWidth*tileMultiplier,tileHeight*tileMultiplier);
        }
      }
    }
  }
}

//Render stuff
function step() {
  draw();
  window.requestAnimationFrame(step);
}
window.requestAnimationFrame(step);
