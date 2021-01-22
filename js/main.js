var option ={
    enableHighAccuracy : true,
    timeout :5000,
};

var options = {
    title : "@notification received from wizdwarfs",
    body : "Good Morning",
    links : "https://peaceful-island-75545.herokuapp.com/", 
    vibrate : [200,200,200],
};

// html components
var loc = document.getElementsByClassName("checkbox-geo"); 
var notify = document.getElementById("wizNotify");
var screenLight = document.getElementById("wizLight");

// functions
function GeolocationPermission(){

    // Geolocation permission
        navigator.permissions.query({name: 'geolocation'}).then(function(permissionStatus){
            console.log(' status :' +  permissionStatus.state);
            loc[0].checked = false;
        if (!(permissionStatus.state == "granted")){
            permissionStatus.onchange = function() {
                console.log('geolocation permission state has changed to ', this.state);
               
              };
               
            navigator.geolocation.getCurrentPosition((position, error,option) =>{
                alert('Current Geolocation:'+ `${position.coords.latitude} ` + `${position.coords.longitude}`);
                loc[0].checked = true;
            });
                
                  
        }

    })
}
function error(err){
    console.warn(`ERROR(${err.code}): ${err.message}`);
}

function NotificationPermission(){
    

    // Notification Permission
        if (!('Notification' in window)){
        alert('Notification is not supported by your browser');
        }else{
             if (!(Notification.permission == "granted")){
            
            var permission = Notification.permission;
            Notification.requestPermission().then(function (permission) {

                // If the user accepts, let's create a notification
                if (Notification.permission === "granted") {
                  var notification = new Notification("Message Receive: ",options);
                  notification.vibrate;
                  
                }else{
                    notify.checked = false;
                    alert("Notification request permission Block");
                    
                }
              });
             }
        }

}    

function AmbientLightPermission(){
    

        if(!('AmbientLightSensor' in window)){
            wizLight.checked = false;
            alert('Ambient Light Sensor is not supported by your browser');
            
        }else{
            const ambientSensor = new AmbientLightSensor();
            ambientSensor.onreading = () => {
                console.log('current light level:' +  ambientSensor.illuminance);
            }
            ambientSensor.onerror = (event) => {
                console.log(event.error.name, event.error.message);
            }
                ambientSensor.start();
        }
} 




