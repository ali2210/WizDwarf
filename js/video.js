(function() {
	var width = 320; 
	var height = 0;
    var	streaming = false;
    var video = null;
    var canvas = null;


function startup() {
	// body...
	video = document.getElementById('video');
	canvas = document.getElementById('canvas');


	// navigator 
	navigator.mediaDevices.getUserMedia({video: true, audio :false})
		.then(function(stream){

			// video source intialize...
			video.srcObject = stream;

			// video start play
			video.play();
		})
		.catch(function(err){

			// console output
			console.log("Error:" + err)
		});

		video.addEventListener('canplay', function(ev){
			if (!streaming) {
				height = video.videoHeight / (video.videoWidth/width);

				if(isNaN(height)){
					height = width/(4/3);
				}

				video.setAttribute('width',width);
				video.setAttribute('height',height);
				canvas.setAttribute('width', width);
				canvas.setAttribute('height',height);
				streaming = true;
			}

		}, false);

	}

	window.addEventListener('load',startup,false);
})

