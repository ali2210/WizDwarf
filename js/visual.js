  if (!! window.WebGLRenderingContext){
    var canvas = document.createElement("canvas"),
    name = ["webgl", "experimental-webgl", "moz-webgl", "webkit-3d"],
    context = false;


    for (var i in name){
      try {
          context = canvas.getContext(name[i]);
          if (context && typeof context.getParameter === "function") {
            console.log(name[i]);
          }
      } catch (e) {

      }
      alert("webgl disable");

    }
  }
  alert("webgl not supported");
