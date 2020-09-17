
function VerifyCountryStatus() {
    var country = document.getElementById('country');
    var status = document.getElementById('status_country');
    if (country) {
      console.log(country.readOnly);
      consle.log(status.checked);
    }
    alert("Filed empty");
}
