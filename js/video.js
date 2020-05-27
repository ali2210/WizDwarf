

// global variables 
var video = null;
var canvas = null;



		video = document.getElementById('video');
		canvas = document.getElementById('canvas');


		if (navigator.mediaDevices.getUserMedia){
			navigator.mediaDevices.getUserMedia({video:true, audio:false})
			.then(function (stream) {
				// body...
				video.srcObject = stream;
				video.play();
			})
			.catch(function(err){
				console.log("Error:"+ err)
			})
		}