function VerifyCountryStatus() {
    var country = document.getElementById('homeland');
    var status = document.getElementById("status_country");

    if ((country.value).length > 0){
        country.readOnly = true;
        status.checked = true;
    }else{
        alert('Enter Country name! ');
    }


}
