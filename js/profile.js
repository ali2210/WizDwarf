


// button
var submit = document.getElementById('submBtn');

function EditEvent(){
    var trash = document.getElementById('trash');
        trash.style.background = "green";
        submit.style.visibility = "visible";    
}


function EmailEditEvent(){
    var email = document.getElementById('inputEmail4');
    if (email.textContent != " "){
        email.textContent = " ";
    } 
    EditEvent();
    email.readonly = false;
}



function NameEditEvent(){
    var name_in = document.getElementById('name');
    if (name_in.textContent != " "){
        name_in.textContent = " ";
    } 
    EditEvent();
    name_in.readonly = false;
}

function FamilyNameEditEvent(){
    var fname_in = document.getElementById('fname');
    if (fname_in.textContent != " "){
        fname_in.textContent = " ";
    } 
    EditEvent();
    fname_in.readonly = false;
}

function ResidenceEditEvent(){
    var res_in = document.getElementById('inputAddress');
    if (res_in.textContent != " "){
        res_in.textContent = " ";
    } 
    EditEvent();
    res_in.readonly = false;
}

function SubResEditEvent(){
    var sres_in = document.getElementById('inputAddress2');
    if (sres_in.textContent != " "){
        sres_in.textContent = " ";
    } 
    EditEvent();
    sres_in.readonly = false;
}

function CountryEditEvent(){
    var coun_in = document.getElementById('country');
    if (coun_in.textContent != " "){
        coun_in.textContent = " ";
    } 
    EditEvent();
    coun_in.readonly = false;
}

function TelEditEvent(){
    var tel_in = document.getElementById('phone');
    if (tel_in.textContent != " "){
        tel_in.textContent = " ";
    } 
    EditEvent();
    tel_in.readonly = false;
}

function ZipEditEvent(){
    var zip_in = document.getElementById('inputZip');
    if (zip_in.textContent != " "){
        zip_in.textContent = " ";
    } 
    EditEvent();
    zip_in.readonly = false;
}