
let width = 500,
  height = 0,
  streaming = false;
	// body...
	const video = document.getElementById("video");
		// vendrURL = window.URL || window.webkitURL; 

		// navigator.getMedia = navigator.getUserMedia ||
		//                      navigator.WebkitGetUserMedia ||
		//                      navigator.msGetUserMedia;

		navigator.mediaDevices.getUserMedia({video:true,audio:false}
		).then(function(stream){
			video.srcObject = stream;
			video.play();
		}).catch(function(err){
			console.log(err);
		});
			
