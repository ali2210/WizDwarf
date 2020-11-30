var option ={
    enableHighAccuracy : true,
    timeout :5000,
};

var options = {
    title : "@notification received from wizdwarfs",
    body : "Good Morning",
    links : "https://peaceful-island-75545.herokuapp.com/home", 
    vibrate : [200,200,200],
};


function GeolocationPermission(){

    // var loc = document.getElementById("checkbox-geo") 

        navigator.permissions.query({name: 'geolocation'}).then(function(permissionStatus){
            console.log(' status :' +  permissionStatus.state);
       
        if (!(permissionStatus.state == "granted")){
            permissionStatus.onchange = function() {
                console.log('geolocation permission state has changed to ', this.state);
              };
               
            navigator.geolocation.getCurrentPosition((position, error,option) =>{
                alert('Current Geolocation:');
                console.log(`${position.coords.latitude}`,`${position.coords.longitude}`);
            });
                
                  
        }else{
            alert("Your browser already set as :", Geolocation_Permission_Success(permissionStatus.state));
            
        }

    })
}
function error(err){
    console.warn(`ERROR(${err.code}): ${err.message}`);
}

function NotificationPermission(){
    var notify = document.getElementById("checkbox-notification");
        if (!('Notification' in window)){
        alert('Notification is not supported by your browser');
        }else if (!(Notification.permission == "granted")){
            console.log(' status :' +  Notification.permission);
            var permission = Notification.permission;
            Notification.requestPermission().then(function (permission) {

                // If the user accepts, let's create a notification
                if (Notification.permission === "granted") {
                  var notification = new Notification("Message Receive: ",options);
                  notification.vibrate;
                }else{
                    alert("Notification request permission Block");
                }
              });
        }else{
            if (Notification.permission === "granted") {
                var notification = new Notification("Message Receive: ",options);
                console.log(notification.vibrate);
              }
        }

}    

function AmbientLightPermission(){
    var light = document.getElementById('checkbox-ambient');

        if(!('AmbientLightSensor' in window)){
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




