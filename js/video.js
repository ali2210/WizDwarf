

// global variables 
var video = null;
// var canvas = null;



		video = document.getElementById('video');
		// canvas = document.getElementById('canvas');

let stream =  await navigator.mediaDevices.getUserMedia({video: true, audio :false});
let rec =  new RecordRTCPromisesHandler(stream, {
	type : video,
});

rec .startRecording();

const sleep = m => new Promise(r => setTimeout(r, m));
await sleep(3000);


await rec.stopRecording();
let blob = await rec.getBlob();
invokeSaveAsDialog(blob); 

		