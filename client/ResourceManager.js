/*The resource manager controls the loading of game resources.  loadResources has optional callback and errorCallback functions to inform the rest of the program when
loading is complete.
Example sources object:
		var sources = {
			tiles: "images/spritesheet_tiles.png",
			characters: "images/spritesheet_characters.png",
			items: "images/spritesheet_items.png",
			particles: "images/spritesheet_particles.png"
		}
*/
var ResourceManager = function() {
    this.totalResources = 0;
    this.failedResources = 0;
    this.successfulResources = 0;
    this.resources = {};
    this.bufferCanvas = document.createElement("canvas")
}

ResourceManager.prototype.loadResources = function(sources, callback, errorCallback) {
    var self = this;
    for (var src in sources) {
        this.totalResources++;
    }

    for (var src in sources) {
        this.resources[src] = new Image();
        this.resources[src].onload = function() {
            self.successfulResources++;
            if (callback) {
                if (self.totalResources == self.successfulResources) {
                    callback();
                }
            }
        };
        this.resources[src].onerror = function() {
            self.failedResources++;
            if (errorCallback) {
                errorCallback(this);
            }
            if (callback) {
              if (self.totalResources == (self.successfulResources + self.failedResources)) {
                  callback();
              }
            }
        }
        this.resources[src].src = sources[src];
        this.resources[src].tints = [];

    }
}

ResourceManager.prototype.getResource = function(resourceName) {
    if (this.resources[resourceName] != undefined) {
        return this.resources[resourceName];
    }

    return null;
}

ResourceManager.prototype.getTintedImage = function(resourceName,r,g,b) {
  var image = this.getResource(resourceName);
  if (image != null) {
    if (image.tints[r+","+g+","+b] == undefined) {
        this.bufferCanvas.width = image.width;
        this.bufferCanvas.height = image.height;
        var bufferCtx = this.bufferCanvas.getContext("2d");
        //bufferCtx.save();
        bufferCtx.clearRect(0, 0, this.bufferCanvas.width, this.bufferCanvas.height);
        bufferCtx.imageSmoothingEnabled = false;
        bufferCtx.drawImage(image, 0, 0);
        bufferCtx.globalCompositeOperation = "source-in";
        bufferCtx.fillStyle = "rgb("+r+","+g+","+b+")";
        bufferCtx.rect(0, 0, this.bufferCanvas.width, this.bufferCanvas.height);
        bufferCtx.fill();
        //bufferCtx.restore();

        var img = new Image();
        img.src = this.bufferCanvas.toDataURL();

        image.tints[r+","+g+","+b] = img;
        return (img);
    }

    return image.tints[r+","+g+","+b]
  }

  return null;
}
