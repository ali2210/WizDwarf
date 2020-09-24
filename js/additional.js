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
