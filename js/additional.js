function VerifyCountryStatus() {
    // ENTER COUNTRY NAME , if country name is not empty then mark it
    VerifyLocationFields(document.getElementById('homeland'), document.getElementById('Homeland'));
}

function VerifyCityStatus() {
    VerifyLocationFields(document.getElementById('mytown'), document.getElementById('Hometown'));
}

function PostalStatus() {
    VerifyLocationFields(document.getElementById('area'), document.getElementById('buttonTown'));
}

function StreetStatus() {
    VerifyLocationFields(document.getElementById('address'), document.getElementById('streetAddress'));
}

function NumStatus() {
    VerifyLocationFields(document.getElementById('snum'), document.getElementById('snumButton'));
}

function StateStatus(){
    VerifyLocationFields(document.getElementById('state'), document.getElementById('stateButton'));
}

function SaveLocation(){
  var map = document.getElementById('saveButton');
  var geocoder = new  mapkit.Geocoder({
    language : "en-GB",
    getsUserLocation : true
  });

  var street = VerifiedStatusandAppendAddressFormat(document.getElementById('address'), document.getElementById('streetAddress'));
  var num = VerifiedStatusandAppendAddressFormat(document.getElementById('snum'), document.getElementById('snumButton'));
  var city = VerifiedStatusandAppendAddressFormat(document.getElementById('mytown'), document.getElementById('Hometown'));
  var state = VerifiedStatusandAppendAddressFormat(document.getElementById('state'), document.getElementById('stateButton'));
  var postal = VerifiedStatusandAppendAddressFormat(document.getElementById('area'), document.getElementById('buttonTown'));
  var country = VerifiedStatusandAppendAddressFormat(document.getElementById('homeland'), document.getElementById('Homeland'));

  var locapi = {
    Street : street,
    Route : num,
    City: city,
    State : state,
    Postal: postal,
    Country: country
  };

  var address = formattedAddress(locapi);
  geocoder.lookup(address, (data,error) =>{
    if (error) {
      console.error(error);
    }else{
      console.log(data);
      var long = data.results[0].coordinates.longitude;
      var lat = data.results[0].coordinates.latitude;
      console.log("longitude:"+ long);
      console.log("latitude:" + latitude);
    }

  });
}

function VerifyLocationFields(inputField, verifyButton) {
    var checkIcon = verifyButton.parentElement.getElementsByClassName("verified-icon")[0];
    var uncheckedIcon = verifyButton.parentElement.getElementsByClassName("unveriied-icon")[0];
    if ((inputField.value).length > 0){
        inputField.readOnly = true;
        checkIcon.classList.remove('d-none');
        uncheckedIcon.classList.add('d-none');
        verifyButton.innerHTML = 'verified';
        verifyButton.classList.remove('btn-primary', 'btm-success', 'btn-danger');
        verifyButton.classList.add('btn-success');
    } else {
        checkIcon.classList.add('d-none');
        uncheckedIcon.classList.remove('d-none');
        verifyButton.innerHTML = 'not verified';
        verifyButton.classList.remove('btn-primary', 'btm-success', 'btn-danger');
        verifyButton.classList.add('btn-danger');
    }
}

function VerifiedStatusandAppendAddressFormat(inputField, verifyButton){
  if (inputField.value.length > 0 && Status(verifyButton)) {
    return inputField.value
  }
}

function Status(verifyButton){
  if (verifyButton.innerHTML == "verified"){
    return true;
  }
  return false;
}

function formattedAddress(locapi) {
  var x = locapi['Street']+','+locapi['Route']+','+locapi['City']+','+locapi['State']+','+  locapi['Postal']+','+locapi['Country'];
  console.log("string :", x)
  return x
}
