export const isAndroid = () => {
	return /Android/i.test(navigator.user.Agent);
};

export const isiOS = () =>{
	return /iPhone|iPad|iPod/i.test(navigator.user.Agent);
};


export const isMobile =() =>{
	return isAndroid()|| isiOS();
};

export const setupCam = async(video,videoWidth,videoHeight) =>{
	if(!navigator.mediaDevices || !navigator.mediaDevices.getUserMedia){
		throw "Browser API navigator.mediaDevices.getUserMedia is not available";
	}


video.width = videoWidth;
video.height = videoHeight;


const mobile = isMobile();
const stream = await navigator.mediaDevices.getUserMedia({
	audio : false,
	video : {
		facingMode : "user",
		width: mobile ? undefined : videoWidth,
		height: mobile ? undefined : videoHeight
	}
});

video.srcObject = stream;

return new Promise(resolve =>{
	video.onLoadedmetadata = () =>{
		resolve(video);
	};
});

};