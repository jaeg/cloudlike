var c=document.getElementById("game");
var ctx=c.getContext("2d");

var sock = null;
var wsuri = "ws://127.0.0.1:1234";

var tileWidth = 16;
var tileHeight = 16;

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
flip = false
drawBuffer = []
entities = []
player = {}

sock.onmessage = function(e) {

    var payload = e.data;
    var index = payload.indexOf(":");
    var type = payload.substr(0,index);
    var json = payload.substr(index+1);

    var data = JSON.parse(json)
    if (type == "view") {
      drawBuffer = data
    }

    if (type == "entities") {
      entities = data
    }

    if (type == "player") {
      player = data
      flip = !flip;
    }

    if (type == "commandQueue") {
      document.getElementById("commandQueue").innerHTML = data
    }

    if (type == "messageLog") {
      if (data != null) {
        document.getElementById("messageLog").innerHTML = ""
        //Only show the last 10 messages
        var start = data.length - 10
        if (start < 0) {
          start = 0
        }
        for (var i = start; i < data.length; i++){
          document.getElementById("messageLog").innerHTML += data[i] + "</br>";
        }
      }
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
          var spriteIndex = drawBuffer[x][y].SpriteIndex
          var tile = getTileLocation(spriteIndex, 16)
          drawTile(x*tileWidth*tileMultiplier,y*tileHeight*tileMultiplier,tileWidth,tileHeight,"terrain", tile.X,tile.Y, tileWidth,tileHeight);
        }
      }
    }
  }


  if (entities != null) {
    for (var i = 0; i < entities.length; i++) {
      entity = entities[i]
      coords = toCanvasCoords(entity.X, entity.Y)

      var x = coords.x
      var y = coords.y
      var tile = {}
      if (entity.Direction > -1) {
        tile = getTileLocation(entity.SpriteIndex + entity.Direction, 16, true)
      } else {
        tile = getTileLocation(entity.SpriteIndex, 16, true)
      }
      drawTile(x, y, tileWidth, tileHeight, entity.Resource, tile.X, tile.Y, tileWidth, tileHeight);
    }
  }
}

function drawTile(x,y, width, height, image, subx,suby, subwidth, subheight, r,g,b) {
  img = resourceManager.getResource(image);
  if (img != null)
    ctx.drawImage(img, subx * subwidth, suby * subheight, subwidth, subheight, x, y, width * tileMultiplier, height * tileMultiplier);
}

function toCanvasCoords(x,y) {
  if (player != null) {
    var coords = {x:null, y:null};
    coords.x = c.width/2 + (x - player.X) * tileWidth * tileMultiplier
    coords.y = c.height/2 + (y - player.Y) * tileHeight * tileMultiplier
    return coords;
  }
  return null;
}

function getTileLocation(index, numberAcross, useFlip) {
  var tile = {};

  if (useFlip) {
    if (flip) {
      index += numberAcross
    }
  }
  //figure out tile location
  tile.X = index;
  tile.Y = 0;
  while (tile.X > numberAcross-1) {
    tile.X -= numberAcross;
    tile.Y++;
  }

  return tile;
}


resourceManager = new ResourceManager();
resourceManager.loadResources({terrain:"assets/space/tiny_galaxy_world.png", item:"assets/space/tiny_galaxy_items.png", npc:"assets/space/tiny_galaxy_monsters.png", fx: "assets/space/tiny_galaxy_fx.png"}, step)

//Render stuff
function step() {
  draw();
  window.requestAnimationFrame(step);
}
