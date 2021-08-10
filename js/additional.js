
const marker_location = document.getElementsByClassName("location-div")[0];

marker_location.children[0].addEventListener("click", function(){
  marker_location.children[1].style.visibility = "visible";
  // marker_location.children[1].innerHTML = "x, y unknown";
  navigator.permissions.query({name:'geolocation'}).then(function(permissionStatus){
    if (!(permissionStatus.state == "granted")) {
        permissionStatus.onchange = function(){
        marker_location.children[0].children[0].style.color="red";
        // marker_location.children[0].children[0].innerHTML = `permission denied`;
      };

      function error(err){
        console.warn(`ERROR(${err.code}): ${err.message}`);
      };

      var option = {
        enableHighAccuracy: true,
      };
      
      navigator.geolocation.getCurrentPosition((position, error, option)=>{
        marker_location.children[0].children[0].style.color="darkgreen";
        marker_location.children[1].style.color= "purple";
        var lati = `${position.coords.latitude}`;
        var logi = `${position.coords.longitude}`;
        marker_location.children[1].value = lati +` , `+ logi; 
      });
    }
  })
});
