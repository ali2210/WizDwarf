

/**
 * declaration variables
 * userProfileForm specified web form along with avatar option
 */
const userProfileForm = document.getElementsByClassName('profile-form')[0];

/**
 * Editform input fields
 */
const editForm = document.getElementsByClassName('edit')[0];
const editchoice = document.getElementsByClassName('col')[0];

/**
 * Avatar Upload or capture via stream
 */
userProfileForm.children[0].children[0].children[0].children[0].addEventListener("click", event=>{
    if(event.target){
        editForm.style.top = "700px";
        avatarSelector.style.left= "452px";
     }
    
});


/**
 *  Edit-Option  allow user to add information which is not in database
 *  Edit your Name 
 */ 
editchoice.children[0].children[1].addEventListener("click", event =>{
    console.log(editForm.children[0].firstElementChild.children[0].readOnly);
    editForm.children[0].firstElementChild.children[0].readOnly = false;
    editForm.children[0].firstElementChild.children[0].addEventListener("change", function(){
        console.log(editForm.children[0].firstElementChild.children[0].value.length);
        if(editForm.children[0].firstElementChild.children[0].value.length>0){
            editForm.children[0].firstElementChild.children[0].style.color = "green";
            editForm.children[0].firstElementChild.children[1].style.color = "green";
            editchoice.children[11].children[0].disabled = false;
            editchoice.children[11].children[0].style.color = "green";
        }else if(editForm.children[0].firstElementChild.children[0].value.length == 0){
            editForm.children[0].firstElementChild.children[0].style.color = "red";
            editForm.children[0].firstElementChild.children[1].style.color = "red";
            editchoice.children[11].children[0].style.color = "red";
            editchoice.children[11].children[0].disabled = true;
        }
    });
});

/**
 * Edit your Lastname
 */

editchoice.children[1].children[1].addEventListener("click", event =>{
    console.log(editForm.children[0].children[1].firstElementChild.readOnly);
    editForm.children[0].children[1].firstElementChild.readOnly = false;
    editForm.children[0].children[1].children[0].addEventListener("change", function(){
        console.log(editForm.children[0].children[1].firstElementChild.value.length);
        if(editForm.children[0].children[1].firstElementChild.value.length>0){
            editForm.children[0].children[1].firstElementChild.style.color = "green";
            editForm.children[0].children[1].lastElementChild.style.color = "green";
            editchoice.children[11].children[0].disabled = false;
            editchoice.children[11].children[0].style.color = "green";
        }else if(editForm.children[0].children[1].firstElementChild.value.length == 0){
            editForm.children[0].children[1].firstElementChild.style.color = "red";
            editForm.children[0].children[1].lastElementChild.style.color = "red";
            editchoice.children[11].children[0].style.color = "red";
            editchoice.children[11].children[0].disabled = true;
        }
    });
});

/**
 * Edit Your Gender For Males
 */


editchoice.children[2].children[0].children[0].addEventListener("click", event=>{
    console.log(editForm.children[0].children[2].children[0].children[0].children[0].style.border);
    editForm.children[0].children[2].children[0].children[0].style.border = "1px green dotted";
    editForm.children[0].children[2].children[0].children[0].style.color = "green";
    editForm.children[0].children[2].children[0].children[0].style.backgroundColor = "white";
});

/**
 * Edit Your Gender For Female
 */

editchoice.children[2].children[1].addEventListener("click", event=>{
    console.log("female",editForm.children[0].children[2].children[1].children[0].children[0].style.border);
    editForm.children[0].children[2].children[1].children[0].style.border = "1px green dotted";
    editForm.children[0].children[2].children[1].children[0].style.color = "green";
    editForm.children[0].children[2].children[1].children[0].style.backgroundColor = "white";
});

/**
 * Edit your Address
 */

editchoice.children[3].children[1].addEventListener("click", event =>{
    console.log(editForm.children[0].children[3].firstElementChild.readOnly);
    editForm.children[0].children[3].firstElementChild.readOnly = false;
    editForm.children[0].children[3].firstElementChild.addEventListener("change", function(){
        console.log("address", editForm.children[0].children[3].firstElementChild.value.length);
        if(editForm.children[0].children[3].firstElementChild.value.length>0){
            editForm.children[0].children[3].firstElementChild.style.color = "green";
            editForm.children[0].children[3].lastElementChild.style.color = "green";
            editchoice.children[11].children[0].disabled = false;
            editchoice.children[11].children[0].style.color = "green";
        }else if(editForm.children[0].children[3].firstElementChild.value.length == 0){
            editForm.children[0].children[3].firstElementChild.style.color = "red";
            editForm.children[0].children[3].lastElementChild.style.color = "red";
            editchoice.children[11].children[0].style.color = "red";
            editchoice.children[11].children[0].disabled = true;
        }
    });
});

/**
 * Edit your Apartment
 */
editchoice.children[4].children[1].addEventListener("click", event =>{
    console.log(editForm.children[0].children[4].firstElementChild.readOnly);
    editForm.children[0].children[4].firstElementChild.readOnly = false;
    editForm.children[0].children[4].firstElementChild.addEventListener("change", function(){
        console.log("address", editForm.children[0].children[4].firstElementChild.value.length);
        if(editForm.children[0].children[4].firstElementChild.value.length>0){
            editForm.children[0].children[4].firstElementChild.style.color = "green";
            editForm.children[0].children[4].lastElementChild.style.color = "green";
            editchoice.children[11].children[0].disabled = false;
            editchoice.children[11].children[0].style.color = "green";
        }else if(editForm.children[0].children[4].firstElementChild.value.length == 0){
            editForm.children[0].children[4].firstElementChild.style.color = "red";
            editForm.children[0].children[4].lastElementChild.style.color = "red";
            editchoice.children[11].children[0].style.color = "red";
            editchoice.children[11].children[0].disabled = true;
        }
    });
});

/**
 * Edit your Reside Country-name
 */

editchoice.children[5].children[1].addEventListener("click", event =>{
    console.log(editForm.children[0].children[5].firstElementChild.readOnly);
    editForm.children[0].children[5].firstElementChild.readOnly = false;
    editForm.children[0].children[5].firstElementChild.addEventListener("change", function(){
        console.log("address", editForm.children[0].children[5].firstElementChild.value.length);
        if(editForm.children[0].children[5].firstElementChild.value.length>0){
            editForm.children[0].children[5].firstElementChild.style.color = "green";
            editForm.children[0].children[5].lastElementChild.style.color = "green";
            editchoice.children[11].children[0].disabled = false;
            editchoice.children[11].children[0].style.color = "green";
        }else if(editForm.children[0].children[5].firstElementChild.value.length == 0){
            editForm.children[0].children[5].firstElementChild.style.color = "red";
            editForm.children[0].children[5].lastElementChild.style.color = "red";
            editchoice.children[11].children[0].style.color = "red";
            editchoice.children[11].children[0].disabled = true;
        }
    });
});


/**
 * Edit Your City-name
 */
editchoice.children[6].children[1].addEventListener("click", event =>{
    console.log(editForm.children[0].children[6].firstElementChild.readOnly);
    editForm.children[0].children[6].firstElementChild.readOnly = false;
    editForm.children[0].children[6].firstElementChild.addEventListener("change", function(){
        console.log("address", editForm.children[0].children[6].firstElementChild.value.length);
        if(editForm.children[0].children[6].firstElementChild.value.length>0){
            editForm.children[0].children[6].firstElementChild.style.color = "green";
            editForm.children[0].children[6].lastElementChild.style.color = "green";
            editchoice.children[11].children[0].disabled = false;
            editchoice.children[11].children[0].style.color = "green";
        }else if(editForm.children[0].children[6].firstElementChild.value.length == 0){
            editForm.children[0].children[6].firstElementChild.style.color = "red";
            editForm.children[0].children[6].lastElementChild.style.color = "red";
            editchoice.children[11].children[0].style.color = "red";
            editchoice.children[11].children[0].disabled = true;
        }
    });
});

/**
 * Edit your zip-code
 */

editchoice.children[9].children[1].addEventListener("click", event =>{
    console.log(editForm.children[0].children[8].firstElementChild.readOnly);
    editForm.children[0].children[9].firstElementChild.readOnly = false;
    editForm.children[0].children[9].firstElementChild.addEventListener("change", function(){
        console.log("address", editForm.children[0].children[9].firstElementChild.value.length);
        if(editForm.children[0].children[9].firstElementChild.value.length>0){
            editForm.children[0].children[9].firstElementChild.style.color = "green";
            editForm.children[0].children[9].lastElementChild.style.color = "green";
            editchoice.children[11].children[0].disabled = false;
            editchoice.children[11].children[0].style.color = "green";
        }else if(editForm.children[0].children[9].firstElementChild.value.length == 0){
            editForm.children[0].children[9].firstElementChild.style.color = "red";
            editForm.children[0].children[9].lastElementChild.style.color = "red";
            editchoice.children[11].children[0].style.color = "red";
            editchoice.children[11].children[0].disabled = true;
        }
    });
});

/**
 * Edit Your Twitter-Account url 
 */
editchoice.children[10].children[1].addEventListener("click", event =>{
    console.log(editForm.children[0].children[10].firstElementChild.readOnly);
    editForm.children[0].children[10].firstElementChild.readOnly = false;
    editForm.children[0].children[10].firstElementChild.addEventListener("change", function(){
        console.log("address", editForm.children[0].children[10].firstElementChild.value.length);
        if(editForm.children[0].children[10].firstElementChild.value.length>0){
            editForm.children[0].children[10].firstElementChild.style.color = "green";
            editForm.children[0].children[10].lastElementChild.style.color = "green";
            editchoice.children[11].children[0].disabled = false;
            editchoice.children[11].children[0].style.color = "green";
        }else if(editForm.children[0].children[10].firstElementChild.value.length == 0){
            editForm.children[0].children[10].firstElementChild.style.color = "red";
            editForm.children[0].children[10].lastElementChild.style.color = "red";
            editchoice.children[11].children[0].style.color = "red";
        }
    });
});


/**
 * Upload your Image 
 */
const uploadImage = document.getElementsByClassName('camera')[0];
const avatarSelector = document.getElementsByClassName('avatar')[0];

let video = null;

uploadImage.children[0].children[1].addEventListener("change", (event) =>{
    console.log(event.target.files[0]);
    const preview = document.createElement('img');
    preview.src = URL.createObjectURL(event.target.files[0]);
    avatarSelector.appendChild(preview);
    avatarSelector.children[1].style.width ="159px";
    avatarSelector.children[1].style.position = "absolute";
    avatarSelector.children[1].style.top = "46px";
    avatarSelector.children[1].style.left = "-14px"; 
    avatarSelector.children[1].style.height = "54px";
});

/**
 * Motion Picture to still picture
 * 
 */
const gallerySelector = document.getElementsByClassName('gallery')[0];
let stop = null;
uploadImage.children[0].children[0].addEventListener("click", (event) =>{
    
    
    let width = 214;
    let height = 160;

    /**
     * Change Video button with play button with additional parameters
     */
    
    uploadImage.children[0].children[0].style.color = "white";
    uploadImage.children[0].children[0].className = "btn btn-success";
    uploadImage.children[0].children[0].style.position = "absolute";
    uploadImage.children[0].children[0].style.top = "30px";
    uploadImage.children[0].children[0].style.left = "-96px";
    uploadImage.children[0].children[0].style.borderRadius ="10px";
    
    /**
     * Create video divider 
     */
    
    video = document.createElement('video');
    video.id = "webcam";
    video.className = "webcam-forward";
    
    /**
     * Adjust video size
     */
    uploadImage.appendChild(video);
    uploadImage.children[0].children[0].disabled = true;
    

    /**
     * Create stop button & adjust button parameters 
     */

    stop = document.createElement('button');
    stop.className = "btn btn-success";
    
    let index = document.createElement('i');
    index.className = "fa fa-stop";
    stop.appendChild(index);
    
    stop.style.position = "absolute";
    stop.style.left = "207px";
    stop.style.top = "155px";
    stop.style.borderRadius = "10px";
    uploadImage.appendChild(stop);

    /**
     * video window size
     */
    video.style.width = width+"px";
    video.style.height = height+"px";
    
    /**
     * additional parameters such as camera-motion toward users or envorinment, audio parameters.
     * In case video-streaming fit window size then video play automatically, otherwise window convey 
     * error message.   
     */

    navigator.mediaDevices.getUserMedia({
        video:  {facingMode: "user"},
        audio: false,
    }).then((stream) =>{
        video.srcObject = stream;
        video.play();
    }).catch((err)=>{
        video.innerHTML = "Video is not supported";
    });

    /**
     * Stop button allow users to stop camera-motion and adjust as portrait.
     */

    stop.addEventListener("click", (stop_video)=>{
        console.log("video:",video); 
        video.paused = true;
        video.pause();
        uploadImage.children[0].children[0].disabled = true;
        uploadImage.children[2].disabled = true;
        gallerySelector.children[4].children[0].style.visibility = "hidden";
        stop_video.preventDefault();
    },true);
    

    /**
     * Camera button allow users to capture picture. 
     * Capture Picture is an snapshot of video frame per second FPS.
     */
    
    gallerySelector.children[4].addEventListener("click", (capture)=>{
        console.log("capture");
        capturingPhoto();
        capture.preventDefault();
    },true);

    function capturingPhoto(){
       
        let width = 214;
       let height = 160;
       
        /**
         * Create still picture frames
         */
       var canvas_Session = new Array(4);
        
        for (let index = 0; index < canvas_Session.length; index++) {
            canvas_Session[index] = document.createElement("canvas");
            canvas_Session[index].id = "canvas"+index;
            canvas_Session[index].style.position = "absolute";
            canvas_Session[index].style.width = "50px";
            canvas_Session[index].style.height = "47px";
            gallerySelector.appendChild(canvas_Session[index]);
        }

        /**
         * Still pictures frame adjust window size
         */
        
        canvas_Session[0].style.left = "46px";
        canvas_Session[0].style.top = "37px";

        canvas_Session[1].style.left = "134px";
        canvas_Session[1].style.top = "37px";
        
        canvas_Session[2].style.left = "46px";
        canvas_Session[2].style.top = "119px";
        
        canvas_Session[3].style.left = "134px";
        canvas_Session[3].style.top = "118px";
        
        var frames_per_sec = new Array(4); 
        
        console.log(canvas_Session,width,height);
        
        if(width && height){

            let context1 = canvas_Session[0].getContext('2d');
            let context2 = canvas_Session[1].getContext('2d');
            let context3 = canvas_Session[2].getContext('2d');
            let context4 = canvas_Session[3].getContext('2d');

            /**
             * Camera motion frames transforme data into portarit or 2-Dimensional image 
             */

            context1.drawImage(video, 0, 0, width, height);
            context2.drawImage(video, 0, 0, width, height);
            context3.drawImage(video, 0, 0, width, height);
            context4.drawImage(video, 0, 0, width, height);

            /**
             * Portrait-image add type of image (.png, .jpeg, .svg)
             */
            var data = canvas_Session[0].toDataURL('image/png');
            

            /**
             * Create image-divider 
             */
            let avatar_build = document.createElement('img');
            avatar_build.id = "avatar";
            avatar_build.src = data;
            avatar_build.ref = "pictures";
            avatar_build.name = "avatars_gen";
            console.log(avatar_build);
            
            /**
             * Portriat-image display 
             */
            avatarSelector.appendChild(avatar_build);
            avatarSelector.style.border = "none";
        }else{
            clearCanvas()
        }
    }

    function clearCanvas(){
        
        let context1 = canvas_Session[0].getContext('2d');
        let context2 = canvas_Session[1].getContext('2d');
        let context3 = canvas_Session[2].getContext('2d');
        let context4 = canvas_Session[3].getContext('2d');

        context1.fillStyle = "#AAA";
        context2.fillStyle = "#AAA";
        context3.fillStyle = "#AAA";
        context4.fillStyle = "#AAA";

        context1.fillRect(0,0,canvas_Session[0].width, canvas_Session[0].height);
        context2.fillRect(0,0,canvas_Session[1].width, canvas_Session[1].height);
        context3.fillRect(0,0,canvas_Session[2].width, canvas_Session[2].height);
        context4.fillRect(0,0,canvas_Session[3].width, canvas_Session[3].height);

        let data = canvas_Session[0].toDataURL('image/png');
        avatarSelector.setAttribute('src', data);

        avatarSelector.children[1].style.width ="153px";
        avatarSelector.children[1].style.position = "absolute";
        avatarSelector.children[1].style.top = "29px";
        avatarSelector.children[1].style.left = "23px"; 
        avatarSelector.children[1].style.height = "80px";
        avatarSelector.style.border = "none";
    }
});

const btn_submit = document.getElementsByClassName("btn-save")[0];
const pusher_channel = {
    appid : "1265511",
    cluster : "mt1",
    key : "65993b3c66b5317411a5",
    secret : "4f8bf3faf121d9c8dadf",
};

btn_submit.addEventListener("click", (event) => {
    if (avatarSelector.children[1].id != ""){
        Pusher.logToConsole = true; //pusher log active 
        let pusher_channel_credentials = new Pusher(pusher_channel.key, 
            {cluster : pusher_channel.cluster, 
              encrypted:true
            });
        console.log("Pusher_channels logs:", pusher_channel_credentials);
        
        let channel_subcribe = pusher_channel_credentials.subscribe('encrypted-photo-stream');
        console.log("subcribe:",channel_subcribe);
        
        channel_subcribe.bind('photos-bytes', (pusher_data) => {
            
            console.log("Data logs:", pusher_data.data.items); 
        });
    }                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              
}) 


