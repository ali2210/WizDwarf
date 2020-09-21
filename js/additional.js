function VerifyCountryStatus() {

  // ENTER COUNTRY NAME , if country name is not empty then mark it
    var country = document.getElementById('homeland');
    var status = document.getElementById('status_country');
    var button = document.getElementById('Homeland');

    if ((country.value).length > 0){
        country.readOnly = true;
        status.checked = true;
        button.style.visibility = "hidden";
    }else{
        alert('Enter Country name! ');
    }
}

function VerifyCityStatus() {
  var city = document.getElementById('mytown');
  var status = document.getElementById('status_city');
  var button = document.getElementById('Hometown');

  if ((city.value).length > 0){
      city.readOnly = true;
      status.checked = true;
      button.style.visibility = "hidden";
  }else{
      alert('Enter city name! ');
  }
}

function PostalStatus() {
  var postal = document.getElementById('area');
  var status = document.getElementById('status_postal');
  var button = document.getElementById('buttonTown');

  if ((postal.value).length > 0){
      postal.readOnly = true;
      status.checked = true;
      button.style.visibility = "hidden";
  }else{
      alert('Enter postal name! ');
  }
}
