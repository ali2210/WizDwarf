
	// body...
	const video = document.querySelector('video');
	var promiseOld =  function(constraints, sucessCallback,errorCallback){
		


		var getUserMedia = (navigator.getUserMedia || 
							navigator.webkitGetUserMedia ||
							navigator.mozGetUserMedia);
		if(!getUserMedia){
			return Promise.reject(new Error('getUserMedia is not implemented in this browser'));

		}

		return new Promise(function(sucessCallback,errorCallback){
			getUserMedia.call(navigator,constraints,sucessCallback,errorCallback);
		});

		if (navigator.mediaDevices == undefined) {
			navigator.mediaDevices = {};
		}

		if (navigator.mediaDevices.getUserMedia === undefined){
			navigator.mediaDevices.getUserMedia = promiseOld;
		}

		var constraints = {audio : true , video : {width : 580 ,height : 320}
	    };

		navigator.mediaDevices.getUserMedia(constraints).then(function(stream){
			video.src = window.URL.createObjectURL(stream);
			video.onloadedmetadata = function(e){
				video.play();
			};
		})
			.catch(function(err){
				console.log(err.name +":" + err.message);
			});
	}

	/*navigator.mediaDevices.getUserMedia({video: true , audio: false}
	)
	   .then(function(stream) {
	   	// body...
	   	  video.srcObject = stream;
	   	  video.play();
	   })
	   .catch(function(err) {
	   	// body...
	   	console.log("Error:"+ err);
	   });*/

			
