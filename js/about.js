(function() {
    // The width and height of the captured photo. We will set the
    // width to the value defined here, but the height will be
    // calculated based on the aspect ratio of the input stream.
  
    var width = 320;    // We will scale the photo width to this
    var height = 0;     // This will be computed based on the input stream
  
    // |streaming| indicates whether or not we're currently streaming
    // video from the camera. Obviously, we start at false.
  
    var streaming = false;
  
    // The various HTML elements we need to configure or control. These
    // will be set by the startup() function.
  
    var video = null;
    var canvas = null;
    var photo = null;
    
    var startbutton = null;
    var savePic = null;
  
    function startup() {
      video = document.getElementById('video');
      canvas = document.getElementById('canvas');
      photo = document.getElementById('photo');
      
      startbutton = document.getElementById('startbutton');
      savePic = document.getElementById('savePic');

      navigator.mediaDevices.getUserMedia({video: true, audio: false})
      .then(function(stream) {
        video.srcObject = stream;
        video.play();
      })
      .catch(function(err) {
        console.log("An error occurred: " + err);
      });
  
      video.addEventListener('canplay', function(ev){
        if (!streaming) {
          height = video.videoHeight / (video.videoWidth/width);
        
          // Firefox currently has a bug where the height can't be read from
          // the video, so we will make assumptions if this happens.
        
          if (isNaN(height)) {
            height = width / (4/3);
          }
        
          video.setAttribute('width', width);
          video.setAttribute('height', height);
          canvas.setAttribute('width', width);
          canvas.setAttribute('height', height);
          streaming = true;
        }
      }, false);
  
      startbutton.addEventListener('click', function(ev){
        takepicture();
        ev.preventDefault();
      }, false);

      // savePic  want to take retro picture
      savePic.addEventListener('click', function(ev){
        RetroStillPicture();
        ev.preventDefault();
      }, false);
  
      
     clearphoto();
    }

    
    // Fill the photo with an indication that none has been
    // captured.
  
    function clearphoto() {
      var context = canvas.getContext('2d');
      context.fillStyle = "#AAA";
      context.fillRect(0, 0, canvas.width, canvas.height);
  
      var data = canvas.toDataURL('image/png');
      photo.setAttribute('src', data);
      stopRoutine();
      
    }

    // video streaming live or not 
    function stopRoutine(){
        setTimeout(() =>{
            canvas.style.visibility = "hidden";
           
           setTimeout(() =>{
              canvas.style.backgroundColor = "aquamarine";
              canvas.style.visibility = "visible";
              // console.log("state1:", canvas.style.visibility);         
           }, 100);
       }, 1000/2);
    }

    // clear all other html componemts after still picture achieve

    function RetroStillPicture(){
        var context = canvas.getContext('2d');
        if (width && height){
            canvas.width = width;
            canvas.height=  height;
            context.drawImage(video, 0, 0, canvas.width, canvas.height);
            setTimeout(() => {
                canvas.style.backgroundColor = "aquamarine";
                photo.style.backgroundColor = "aquamarine";
                photo.setAttribute('src'," ");
                video.setAttribute('src', " ");
                video.style.visibility = "hidden";
                startbutton.style.visibility = "hidden";
                savePic.style.visibility = "hidden";
            }, 100);

        }
    }
    // Capture a photo by fetching the current contents of the video
    // and drawing it into a canvas, then converting that to a PNG
    // format data URL. By drawing it on an offscreen canvas and then
    // drawing that to the screen, we can change its size and/or apply
    // other changes before drawing it.
  
    function takepicture() {
      var context = canvas.getContext('2d');
      if (width && height) {
        canvas.width = width;
        canvas.height = height;
        context.drawImage(video, 0, 0, width, height);
      
        var data = canvas.toDataURL('image/png');
        photo.setAttribute('src', data);
      } else {
        clearphoto();
      }
    }

   
    
  
    // Set up our event listener to run the startup process
    // once loading is complete.
    window.addEventListener('load', startup, false);
  })();

  // variables declaration 


(function(){
    
    let network = null; 
    var metamaskBtn = null;
    var message = null;
    var node = null;
    var textNode = null;
    var ethAccs = null;
 

  async function connectFunc(){
    network = await detectEthereumProvider();    
    message = document.getElementsByClassName('server-message'); 
    metamaskBtn = document.getElementsByClassName('metabits');
   
    if (network === window.ethereum && window.ethereum.isMetaMask){
      metamaskBtn[0].style.borderStyle = "dotted";
      metamaskBtn[0].style.borderColor = "teal";
      metamaskBtn[0].style.backgroundColor = "lightgreen";
      window.ethereum.isConnected();
      const version = await window.ethereum.request({
        method : 'net_version'
      }); 
      message[0].innerHTML = "Greeat !Metamask already installed:\t" + version 
    }else{
      message[0].innerHTML = "Please install metamask-exstension on your browser.";
      node = document.createElement("a");
      node.href = "https://metamask.io/";
      textNode = document.createTextNode("Metamask Download");
      node.appendChild(textNode);
      message[0].appendChild(node);
    }
    
  }
  

  
  window.addEventListener('load', connectFunc, true);
})();  

async function Metamasklogin(){
  message = document.getElementsByClassName('server-message');
  metamaskBtn = document.getElementsByClassName('metabits');
  const chainId = await ethereum.request({ method: 'eth_chainId' });
  
  handleChainChanged(chainId);

  ethereum.on('chainChanged', handleChainChanged);

  function handleChainChanged(_chainId) {
    // We recommend reloading the page, unless you must do otherwise
    window.location.reload();
  }
  
ethereum
  .request({ method: 'eth_accounts' })
  .then(handleAccountsChanged)
  .catch((err) => {
    // Some unexpected error.
    // For backwards compatibility reasons, if no accounts are available,
    // eth_accounts will return an empty array.
    console.error(err);
  });
  ethereum.on('accountsChanged', handleAccountsChanged);
  message[0].innerHTML = "Metamask Connected";
  // window.ethereum.enable();
  const metamaskAppInit = await window.ethereum.request({
    method : 'eth_requestAccounts'
  })
  
}
  
function handleAccountsChanged(accounts){
  
 
  if(accounts.length === 0){
    message[0].innerHTML = "No account connected with metamask: \t" + (!window.ethereum.isConnected());
  }else{
    if(accounts[0] !== ethAccs){
      ethAccs = accounts[0];     
    }
  }
}  