import cv from "opencv.js";


console.log("Opencv started...")
let video = document.getElementById("videoInput");
video.height = 640;
video.width = 480;

navigator.mediaDevices.getUserMedia({video:true, audio:file}).then(function(stream){
	video.srcObject = stream;
	video.play();


	let src = new cv.Mat(video.height, video.width, cv.CV_8UC4);
	let out = new cv.Mat(video.height, video.width, cv.CV_8UC1);
	let cap = new cv.VideoCapture(video);


	const FPS = 30;
	function processVideo(){
		try{
			let date = Date.now();
			cap.read(src);
			cv.cvtColor(src, out, cv.COLOR_RGBA2GRAY);
			cv.imshow("videoOutput", out);

			let delay = 1000 / FPS -(Date.now() - date);
			setTimeout(processVideo,delay);
		}catch(err){
			console.error(err);
		}
	}

	setTimeout(processVideo, 0);
})
.catch(function(err){
	console.log("An error occured!"+ err)
});