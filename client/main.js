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

sock.onmessage = function(e) {
    console.log(e.data);
}

function handleKeyPressed(e) {
  sock.send(JSON.stringify({type:"input",key:String.fromCharCode(e.keyCode)}));
}

document.addEventListener("keydown", handleKeyPressed, false);
