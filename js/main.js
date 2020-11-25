function GeolocationPermission(){

    var loc = document.getElementById("checkbox-geo"); 

        navigator.permissions.query({name: 'geolocation'}).then(function(permissionStatus){
        console.log('geolocation permission status :' +  permissionStatus.state);
       
        if (!(permissionStatus.state == "granted")){
                  var perm = Geolocation_Permission_Success(permissionStatus.state);
                  permissionStatus.state = perm;
                  loc.checked = true;

        }else{
            var info = Geolocation_Permission_Success(permissionStatus.state);
            console.log('permission already granted ' + info);
            loc.checked = true;
        }

    })
}

function NotificationPermission(){
    var notify = document.getElementById("checkbox-notification");
    
        if (!('Notification' in window)){
        alert('Notification is not supported by your browser');
        notify.checked=false;
        }else if (!(Notification.permission == "granted")){
            var text , perm = Notification_Success_Permission(Notification.permission);
            Notification.permission = perm;
            console.log('Desktop Notification activate...' +  text);
            notify.checked =true;

        }
}

function AmbientLightPermission(){
    var light = document.getElementById('checkbox-ambient');

        if(!('AmbientLightSensor' in window)){
            alert('Ambient Light Sensor is not supported by your browser');
            light.checked = false;
        }else{
            const ambientSensor = new AmbientLightSensor();
            ambientSensor.onreading = () => {
            console.log('current light level:' +  ambientSensor.illuminance);
            light.checked = true;
            }
            ambientSensor.onerror = (event) => {
            console.log(event.error.name, event.error.message);
            }
            ambientSensor.start();
        }
}    



function Geolocation_Permission_Success(geoFlag){
    if (geoFlag == 'denied'){
        geoFlag = 'granetd';
        return geoFlag;
    }
    return geoFlag;
}

function Notification_Success_Permission(perm){
    if (perm == 'denied'){
        perm.state = 'granted'
        var note = new Notification('Notification activate');
        return note, perm;
    }else{
         var note = new Notification('Good to see you ');
    }
    return note, perm;
}
