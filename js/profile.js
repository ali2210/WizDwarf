
// This codebase desgin according to mozilla open source license.
// Redistribution , contribution and improve codebase under license
// convensions. @contact Ali Hassan AliMatrixCode@protonmail.com

/**
 * declaration variables
 * userProfileForm specified web form along with avatar option
 */
const userProfileForm = document.getElementsByClassName('profile-form')[0];

/**
 * Editform input fields
 */

// global marcos
const editForm = document.getElementsByClassName('edit')[0];
const editchoice = document.getElementsByClassName('col')[0];
const _ed01 = document.getElementsByClassName('_ed01')[0];
const _b01 = document.getElementsByClassName('btn-mycontainer')[0];

/**
 * Avatar Upload or capture via stream
 */

userProfileForm.children[0].children[0].children[0].children[0].addEventListener("click", event=>{
    
  if(_ed01.style.top === '' && event.target){
      _ed01.style.position = 'relative';
      _ed01.style.top = '758px';
      _ed01.style.left = '217px';
  }else if(_ed01.style.top === '758px'){
        _ed01.style.position = 'relative';
        _ed01.style.top = '167px';
        _ed01.style.left = '217px';
  }
    
});


/**
 *  Edit-Option  allow user to add information which is not in database
 *  Edit your Name 
 */ 
editchoice.children[0].children[1].addEventListener("click", event =>{
    
    editchoice.children[0].children[0].readOnly = false;
    console.log(editchoice.children[0].children[0].readOnly);
    
    editchoice.children[0].children[0].addEventListener("change", function(){

        console.log(editchoice.children[0].children[0].value);
        if(editchoice.children[0].children[0].value !== "{{.Name}}" && editchoice.children[0].children[0].value.length>0){
           
            editchoice.children[0].children[0].style.color = "green";
            editchoice.children[0].children[1].style.color = "green";
            editchoice.children[0].children[0].style.border = "1px transparent";
           
            _b01.children[0].children[0].disabled = false;
            _b01.children[0].children[0].style.color = "green";
        }else if(editchoice.children[0].children[0].value.length === 0 || editchoice.children[0].children[0].value.length < 0){
            
            editchoice.children[0].children[0].style.color = "red";
            editchoice.children[0].children[1].style.color = "red";
            editchoice.children[0].children[0].style.border = "1px transparent";

            _b01.children[0].children[1].style.color = "red";
            _b01.children[0].children[0].disabled = true;
        }
    });
});

/**
 * Edit your Lastname
 */

editchoice.children[1].children[1].addEventListener("click", event =>{
    
    editchoice.children[1].children[0].readOnly = false;
    console.log(editchoice.children[1].children[0].readOnly);
    editchoice.children[1].children[0].addEventListener("change", function(){
        
        if(editchoice.children[1].children[0].value !== "{{.LastName}}" && editchoice.children[1].children[0].value.length>0){
            
            editchoice.children[1].children[1].style.color = "green";
            editchoice.children[1].children[0].style.color = "green";
            editchoice.children[1].children[0].style.border = "1px transparent";
            
            _b01.children[0].children[0].disabled = false;
            _b01.children[0].children[0].style.color = "green";
        }else if(editchoice.children[1].children[0].value.length === 0 || editchoice.children[1].children[0].value.length < 0) {
            
            editchoice.children[1].children[1].style.color = "red";
            editchoice.children[1].children[0].style.color = "red";
            editchoice.children[0].children[0].style.border = "1px transparent";
            
            _b01.children[0].children[1].style.color = "red";
            _b01.children[0].children[0].disabled = true;
        }
    });
});

/**
 * Edit Your Gender For Males
 */


editchoice.children[2].children[0].addEventListener("click", event=>{
    
    editchoice.children[2].children[0].children[0].style.border = "1px green dotted";
    console.log(editchoice.children[2].children[0].children[0].style.border);

    editchoice.children[2].children[0].children[0].style.color = "green";
    editchoice.children[2].children[0].children[0].style.border = "1px transparent";
    editchoice.children[2].children[0].children[0].style.backgroundColor = "white";
});

/**
 * Edit Your Gender For Female
 */

editchoice.children[2].children[1].addEventListener("click", event=>{
    
    editchoice.children[2].children[1].children[0].style.border = "1px green dotted";
    console.log("female",editchoice.children[2].children[1].children[0].style.border);
    
    editchoice.children[2].children[1].children[0].style.color = "red";
    editchoice.children[2].children[1].children[0].style.border = "1px transparent";
    editchoice.children[2].children[1].children[0].style.backgroundColor = "white";
});

/**
 * Edit your Address
 */

editchoice.children[3].children[1].addEventListener("click", event =>{
    
    editchoice.children[3].children[0].readOnly = false;
    console.log(editchoice.children[3].children[0].readOnly);

    editchoice.children[3].children[0].addEventListener("change", function(){
    
        if(editchoice.children[3].children[0].value !== "{{.Address}}" && editchoice.children[3].children[0].value.length>0){

            editchoice.children[3].children[1].style.color = "green";
            editchoice.children[3].children[0].style.color = "green";
            editchoice.children[3].children[0].style.border = "1px transparent";
            
            _b01.children[0].children[0].disabled = false;
            _b01.children[0].children[0].style.color = "green";
        }else if(editchoice.children[3].children[0].value.length == 0 || editchoice.children[3].children[0].value.length < 0) {
            
            editchoice.children[3].children[1].style.color = "red";
            editchoice.children[3].children[0].style.color = "red";
            editchoice.children[3].children[0].style.border = "1px transparent";
            
            _b01.children[0].children[0].disabled = true;
            _b01.children[0].children[1].style.color = "red";
        }
    });
});

/**
 * Edit your Apartment
 */
editchoice.children[4].children[1].addEventListener("click", event =>{

    editchoice.children[4].children[0].readOnly = false;
    console.log(editchoice.children[4].children[0].readOnly);

    editchoice.children[4].children[0].addEventListener("change", function(){
        
        if(editchoice.children[4].children[0].value !== "{{.Appartment}}" && editchoice.children[4].children[0].value.length>0){
            
            editchoice.children[4].children[0].style.color = "green";
            editchoice.children[4].children[1].style.color = "green";
            editchoice.children[4].children[0].style.border = "1px transparent";

            _b01.children[0].children[0].disabled = false;
            _b01.children[0].children[0].style.color = "green";
        }else if(editchoice.children[4].children[0].value.length == 0 || editchoice.children[4].children[0].value.length < 0){
            
            editchoice.children[4].children[0].style.color = "red";
            editchoice.children[4].children[0].style.color = "red";
            editchoice.children[4].children[0].style.border = "1px transparent";
            
            _b01.children[0].children[0].disabled = true;
            _b01.children[0].children[1].style.color = "red";
        }
    });
});

/**
 * Edit your Reside Country-name
 */

editchoice.children[5].children[1].addEventListener("click", event =>{
    
    editchoice.children[5].children[0].readOnly = false;
    console.log(editchoice.children[5].children[0].readOnly);
    
    editchoice.children[5].children[0].addEventListener("change", function(){
        
        if(editchoice.children[5].children[0].value !== "{{.Country}}" && editchoice.children[5].children[0].value.length>0){
            
            editchoice.children[5].children[0].style.color = "green";
            editchoice.children[5].children[1].style.color = "green";
            editchoice.children[5].children[0].style.border = "1px transparent";
            
            _b01.children[0].children[0].disabled = false;
            _b01.children[0].children[0].style.color = "green";
        }else if(editchoice.children[5].children[0].value.length == 0 || editchoice.children[5].children[0].value.length < 0){
            
            editchoice.children[5].children[0].style.color = "red";
            editchoice.children[5].children[1].style.color = "red";
            editchoice.children[5].children[0].style.border = "1px transparent";
            
            _b01.children[0].children[0].disabled = true;
            _b01.children[0].children[1].style.color = "red";
        }
    });
});


/**
 * Edit Your City-name
 */
editchoice.children[6].children[1].addEventListener("click", event =>{
    
    editchoice.children[6].children[0].readOnly = false;
    console.log(editchoice.children[6].children[0].readOnly);

    editchoice.children[6].children[0].addEventListener("change", function(){
    
        
        if(editchoice.children[6].children[0].value !== "{{.City}}" && editchoice.children[6].children[0].value.length>0){
            
            editchoice.children[6].children[0].style.color = "green";
            editchoice.children[6].children[1].style.color = "green";
            editchoice.children[6].children[0].style.border = "1px transparent";
            
            _b01.children[0].children[0].disabled = false;
            _b01.children[0].children[0].style.color = "green";
        }else if(editchoice.children[6].children[0].value.length == 0 || editchoice.children[6].children[0].value.length < 0){
            
            editchoice.children[6].children[0].style.color = "red";
            editchoice.children[6].children[1].style.color = "red";
            editchoice.children[6].children[0].style.border = "1px transparent";

            _b01.children[0].children[0].disabled = true;
            _b01.children[0].children[1].style.color = "red";
        }
    });
});

/**
 * Edit your zip-code
 */

editchoice.children[9].children[1].addEventListener("click", event =>{
    
    editchoice.children[9].children[0].readOnly = false;
    
    console.log(editchoice.children[9].children[0].readOnly);
    
    editchoice.children[9].children[0].addEventListener("change", function(){
        
        if(editchoice.children[9].children[0].value !== "{{.Email}}" && editchoice.children[9].children[0].value.length>0){
            
            editchoice.children[9].children[0].style.color = "green";
            editchoice.children[9].children[1].style.color = "green";
            editchoice.children[9].children[0].style.border = "1px transparent";
            
            _b01.children[0].children[0].disabled = false;
            _b01.children[0].children[0].style.color = "green";
        }else if(editchoice.children[9].children[0].value.length == 0 || editchoice.children[9].children[0].value.length < 0){
            
            editchoice.children[9].children[0].style.color = "red";
            editchoice.children[9].children[1].style.color = "red";
            editchoice.children[9].children[0].style.border = "1px transparent";
            
            _b01.children[0].children[0].disabled = true;
            _b01.children[0].children[1].style.color = "red";
        }
    });
});

/**
 * Edit Your Twitter-Account url 
 */
editchoice.children[10].children[1].addEventListener("click", event =>{
    
    editchoice.children[10].children[0].readOnly = false;
    console.log(editchoice.children[10].children[0].readOnly);

    editchoice.children[10].children[0].addEventListener("change", function(){
        
        if(editchoice.children[10].children[0].value.length>0){
            
            editchoice.children[10].children[0].style.color = "green";
            editchoice.children[10].children[1].style.color = "green";
            editchoice.children[10].children[0].style.border = "1px transparent";
            
            _b01.children[0].children[0].disabled = false;
            _b01.children[0].children[0].style.color = "green";
        }else if(editchoice.children[10].children[0].value.length == 0 || editchoice.children[10].children[0].value.length < 0){
            
            editchoice.children[10].children[0].style.color = "red";
            editchoice.children[10].children[1].style.color = "red";
            editchoice.children[10].children[0].style.border = "1px transparent";
            
            _b01.children[0].children[0].disabled = true;
            _b01.children[0].children[1].style.color = "red";
        }
    });
});


/**
 * Upload your Image 
 */

// global marcos 
const uploadImage = document.getElementsByClassName('float-card')[0];
const avatarSelector = document.getElementsByClassName('avatar')[0];

let video = null;

let preview =  null;
let image_size  = null;
uploadImage.children[0].addEventListener("change", (event) =>{
    
    // preview object created 
    preview = document.createElement('img');

    // preview object linked to image
    preview.src = URL.createObjectURL(event.target.files[0]);
    
    // preview object append with dom-element
    uploadImage.children[2].appendChild(preview);
    
    // preview object params
    preview.style.position ="relative";
    preview.style.left = "106px";
    preview.style.top = "98px";
    
    // preview object size  
    image_size = event.target.files[0].size;

    
});

// let profile_pic = ;
const dte_picker = document.getElementsByClassName("date-picker")[0];
uploadImage.children[4].addEventListener("click", (event) => {
    
    // information must be true before event happen
    if(preview.src != "" && image_size <= 2745){
        
        // image param
         avatarSelector.children[1].className = "profilo-avatar";
         avatarSelector.children[1].id = "profile";
         avatarSelector.children[1].style.position = "relative";
         avatarSelector.children[1].style.left = "106px";
         avatarSelector.children[1].style.top = "98px";
         avatarSelector.children[1].src = preview.src;
         
        //  scale image object
         if (image_size <= 1500){
             avatarSelector.children[1].style.transform = 'scale(3.7)';
         }else if (image_size <= 2500){
             avatarSelector.children[1].style.transform = 'scale(2.7)';
         }else{
             avatarSelector.children[1].style.transform = 'scale(1.7)';
         }

        //  date object
         const date = new Date();
         const [year, month, today] = [date.getFullYear(), date.getMonth(), date.getDate()];
         
         dte_picker.style.visibility = 'visible';
         dte_picker.value = ' '+today+'-'+month+'-'+year+' ';
        
    }else if(image_size > 2745){
        alert("File size must be less than 2745", image_size);
    }
});

// global marcos 
const btn_submit = document.getElementsByClassName("btn-save")[0];
const pusher_channel = {
    appid : "1265511",
    cluster : "mt1",
    key : "65993b3c66b5317411a5",
    secret : "4f8bf3faf121d9c8dadf",
};

// channels event handlers
btn_submit.addEventListener("click", (event) => {
    
    if (avatarSelector.children[1].id != ""){
     
        // console for development
        Pusher.logToConsole = true; //pusher log active 
        
        // pusher object
        let pusher_channel_credentials = new Pusher(pusher_channel.key, 
            {cluster : pusher_channel.cluster, 
              encrypted:true
            });
        console.log("Pusher_channels logs:", pusher_channel_credentials);
        
        // channel objects 
        let channel_subcribe = pusher_channel_credentials.subscribe('encrypted-photo-stream');
        console.log("subcribe:",channel_subcribe);
        
        // channel bind with data 
        channel_subcribe.bind('photos-bytes', (pusher_data) => {
            
            console.log("Data logs:", pusher_data.data.items); 
        });
    }                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              
}) 




// /**
//  * Motion Picture to still picture
//  * 
//  */
// const gallerySelector = document.getElementsByClassName('gallery')[0];
// let stop = null;
// uploadImage.children[0].children[0].addEventListener("click", (event) =>{
    
    
//     let width = 214;
//     let height = 160;

//     /**
//      * Change Video button with play button with additional parameters
//      */
    
//     uploadImage.children[0].children[0].style.color = "white";
//     uploadImage.children[0].children[0].className = "btn btn-success";
//     uploadImage.children[0].children[0].style.position = "absolute";
//     uploadImage.children[0].children[0].style.top = "30px";
//     uploadImage.children[0].children[0].style.left = "-96px";
//     uploadImage.children[0].children[0].style.borderRadius ="10px";
    
//     /**
//      * Create video divider 
//      */
    
//     video = document.createElement('video');
//     video.id = "webcam";
//     video.className = "webcam-forward";
    
//     /**
//      * Adjust video size
//      */
//     uploadImage.appendChild(video);
//     uploadImage.children[0].children[0].disabled = true;
    

//     /**
//      * Create stop button & adjust button parameters 
//      */

//     stop = document.createElement('button');
//     stop.className = "btn btn-success";
    
//     let index = document.createElement('i');
//     index.className = "fa fa-stop";
//     stop.appendChild(index);
    
//     stop.style.position = "absolute";
//     stop.style.left = "207px";
//     stop.style.top = "155px";
//     stop.style.borderRadius = "10px";
//     uploadImage.appendChild(stop);

//     /**
//      * video window size
//      */
//     video.style.width = width+"px";
//     video.style.height = height+"px";
    
//     /**
//      * additional parameters such as camera-motion toward users or envorinment, audio parameters.
//      * In case video-streaming fit window size then video play automatically, otherwise window convey 
//      * error message.   
//      */

//     navigator.mediaDevices.getUserMedia({
//         video:  {facingMode: "user"},
//         audio: false,
//     }).then((stream) =>{
//         video.srcObject = stream;
//         video.play();
//     }).catch((err)=>{
//         video.innerHTML = "Video is not supported";
//     });

//     /**
//      * Stop button allow users to stop camera-motion and adjust as portrait.
//      */

//     stop.addEventListener("click", (stop_video)=>{
//         console.log("video:",video); 
//         video.paused = true;
//         video.pause();
//         uploadImage.children[0].children[0].disabled = true;
//         uploadImage.children[2].disabled = true;
//         gallerySelector.children[4].children[0].style.visibility = "hidden";
//         stop_video.preventDefault();
//     },true);
    

//     /**
//      * Camera button allow users to capture picture. 
//      * Capture Picture is an snapshot of video frame per second FPS.
//      */
    
//     gallerySelector.children[4].addEventListener("click", (capture)=>{
//         console.log("capture");
//         capturingPhoto();
//         capture.preventDefault();
//     },true);

//     function capturingPhoto(){
       
//         let width = 214;
//        let height = 160;
       
//         /**
//          * Create still picture frames
//          */
//        var canvas_Session = new Array(4);
        
//         for (let index = 0; index < canvas_Session.length; index++) {
//             canvas_Session[index] = document.createElement("canvas");
//             canvas_Session[index].id = "canvas"+index;
//             canvas_Session[index].style.position = "absolute";
//             canvas_Session[index].style.width = "50px";
//             canvas_Session[index].style.height = "47px";
//             gallerySelector.appendChild(canvas_Session[index]);
//         }

//         /**
//          * Still pictures frame adjust window size
//          */
        
//         canvas_Session[0].style.left = "46px";
//         canvas_Session[0].style.top = "37px";

//         canvas_Session[1].style.left = "134px";
//         canvas_Session[1].style.top = "37px";
        
//         canvas_Session[2].style.left = "46px";
//         canvas_Session[2].style.top = "119px";
        
//         canvas_Session[3].style.left = "134px";
//         canvas_Session[3].style.top = "118px";
        
//         var frames_per_sec = new Array(4); 
        
//         console.log(canvas_Session,width,height);
        
//         if(width && height){

//             let context1 = canvas_Session[0].getContext('2d');
//             let context2 = canvas_Session[1].getContext('2d');
//             let context3 = canvas_Session[2].getContext('2d');
//             let context4 = canvas_Session[3].getContext('2d');

//             /**
//              * Camera motion frames transforme data into portarit or 2-Dimensional image 
//              */

//             context1.drawImage(video, 0, 0, width, height);
//             context2.drawImage(video, 0, 0, width, height);
//             context3.drawImage(video, 0, 0, width, height);
//             context4.drawImage(video, 0, 0, width, height);

//             /**
//              * Portrait-image add type of image (.png, .jpeg, .svg)
//              */
//             var data = canvas_Session[0].toDataURL('image/png');
            

//             /**
//              * Create image-divider 
//              */
//             let avatar_build = document.createElement('img');
//             avatar_build.id = "avatar";
//             avatar_build.src = data;
//             avatar_build.ref = "pictures";
//             avatar_build.alt = "avatars-image"
//             avatar_build.name = "avatars_gen";
//             console.log(avatar_build);
            
//             /**
//              * Portriat-image display 
//              */
//             avatarSelector.appendChild(avatar_build);
//             avatarSelector.style.border = "none";
//         }else{
//             clearCanvas()
//         }
//     }

//     function clearCanvas(){
        
//         let context1 = canvas_Session[0].getContext('2d');
//         let context2 = canvas_Session[1].getContext('2d');
//         let context3 = canvas_Session[2].getContext('2d');
//         let context4 = canvas_Session[3].getContext('2d');

//         context1.fillStyle = "#AAA";
//         context2.fillStyle = "#AAA";
//         context3.fillStyle = "#AAA";
//         context4.fillStyle = "#AAA";

//         context1.fillRect(0,0,canvas_Session[0].width, canvas_Session[0].height);
//         context2.fillRect(0,0,canvas_Session[1].width, canvas_Session[1].height);
//         context3.fillRect(0,0,canvas_Session[2].width, canvas_Session[2].height);
//         context4.fillRect(0,0,canvas_Session[3].width, canvas_Session[3].height);

//         let data = canvas_Session[0].toDataURL('image/png');
//         avatarSelector.setAttribute('src', data);

//         avatarSelector.children[1].style.width ="153px";
//         avatarSelector.children[1].style.position = "absolute";
//         avatarSelector.children[1].style.top = "29px";
//         avatarSelector.children[1].style.left = "23px"; 
//         avatarSelector.children[1].style.height = "80px";
//         avatarSelector.style.border = "none";
//     }
// });

