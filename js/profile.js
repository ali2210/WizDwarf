


// button
var submit = document.getElementById('submBtn');

function EditEvent(){
    var trash = document.getElementById('trash');
    if(trash.checked == true){
        trash.style.background = "green";
        submit.style.visibility = "visible";
    }    
}


function EmailEditEvent(){
    var email = document.getElementById('inputEmail4');
    EditEvent();
    email.readonly = false;
}

function PasswordEditEvent(){
    var password_in = document.getElementById('inputPassword4');
    EditEvent();
    password_in.readonly = false;
}

function NameEditEvent(){
    var name_in = document.getElementById('name');
    EditEvent();
    name_in.readonly = false;
}

function FamilyNameEditEvent(){
    var fname_in = document.getElementById('fname');
    EditEvent();
    fname_in.readonly = false;
}

function ResidenceEditEvent(){
    var res_in = document.getElementById('inputAddress');
    EditEvent();
    res_in.readonly = false;
}

function SubResEditEvent(){
    var sres_in = document.getElementById('inputAddress2');
    EditEvent();
    sres_in.readonly = false;
}

function CountryEditEvent(){
    var coun_in = document.getElementById('country');
    EditEvent();
    coun_in.readonly = false;
}

function TelEditEvent(){
    var tel_in = document.getElementById('phone');
    EditEvent();
    tel_in.readonly = false;
}

function ZipEditEvent(){
    var zip_in = document.getElementById('inputZip');
    EditEvent();
    zip_in.readonly = false;
}