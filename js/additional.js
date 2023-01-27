const marker_location = document.getElementsByClassName('location-div')[0];
var option = {
  enableHighAccuracy: true,
};

let latitude = null;
let longitude = null;

marker_location.children[0].addEventListener("click", function(){
  marker_location.children[1].style.visibility = "visible";
  navigator.permissions.query({name:'geolocation'}).then(function(permissionStatus){
    if (!(permissionStatus.state == "granted")) {
        permissionStatus.onchange = function(){
        marker_location.children[0].children[0].style.color="red";
      }
      
      navigator.geolocation.getCurrentPosition((position, error, option)=>{
        marker_location.children[0].children[0].style.color="darkgreen";
        marker_location.children[1].style.color= "purple";
        latitude = `${position.coords.latitude}`;
        longitude = `${position.coords.longitude}`;
        marker_location.children[1].value = latitude +` , `+ longitude; 
      });
    }else{
      marker_location.children[0].children[0].style.color = 'red';
      marker_location.children[1].style.visibility = 'hidden'; 
      alert('Please reset your location. Your Last Location:'+' '+ latitude+ longitude);
    }
  });
});

function error(err){
  console.warn(`ERROR(${err.code}): ${err.message}`);
};
