function GeolocationPermission(){

    var loc = document.getElementById("geo");

    if (loc.checked){ 

        navigator.permissions.query({name: 'geolocation'}).then(function(permissionStatus){
        console.log('geolocation permission status :' +  permissionStatus.state);
       
        if (!(permissionStatus.state == "granted")){
                  var perm = Geolocation_Permission_Success(permissionStatus.state, loc.checked);
                  permissionStatus.state = perm;
        }else{
            var info = Geolocation_Permission_Success(permissionStatus.state, loc.checked);
            console.log('permission already granted ' + info);
        }

        })
    }
}

function NotificationPermission(){
    var notify = document.getElementById("notify");
    if (notify.checked){
    
        if (!('Notification' in window)){
        alert('Notification is not supported by your browser');
        }else if (Notification.permission == "granted"){
            var text , perm = Notification_Success_Permission(notify, Notification.permission);
            Notification.permission = perm;
            console.log('Desktop Notification activate...' +  text);
        }
    }else{
            if(Notification.permission == 'denied'){
                var text , perm =  Notification_Success_Permission(notify, Notification.permission);
                Notification.permission = perm;
                console.log('Desktop Notification activate...' +  text);
            }
        }
    }

function AmbientLightPermission(){
    var light = document.getElementById('light');

    if(light.checked){
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
}



function Geolocation_Permission_Success(gFlag, loc){
    if (loc == true && gFlag == 'denied'){
        gFlag = 'granetd';
        return gFlag;
    }else{
        if (loc == true && gFlag == 'granted') 
            return gFlag;
    }
}

function Notification_Success_Permission(notify, perm){
    if (notify == true && perm == 'granted'){
        var note = new Notification('Good to see you');
        return note, perm;
    }else if (notify == true && perm == 'denied'){
        perm = 'granted';
        var note = new Notification('Good to see you');
        return note, perm;
    }
}
