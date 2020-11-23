


// button
var submit = document.getElementById('submBtn');

function EditEvent(){
    var trash = document.getElementById('trash');
        trash.style.background = "red";
        submit.style.visibility = "visible";    
}


function EmailEditEvent(){
    var email = document.getElementById('inputEmail4'); 
    email.readonly = false;
    if (email.textContent.length>0){
        EditEvent();
    }

}



function NameEditEvent(){
    var name_in = document.getElementById('name'); 
    name_in.readonly = false;
    if(name_in.textContent.length>0){
        EditEvent();
    }
    
}

function FamilyNameEditEvent(){
    var fname_in = document.getElementById('fname'); 
    fname_in.readonly = false;
    if(fname_in.textContent.length>0){
        EditEvent();
    }
}

function ResidenceEditEvent(){
    var res_in = document.getElementById('inputAddress');
    res_in.readonly = false;
    if(res_in.textContent.length>0){
        EditEvent();
    }
}

function SubResEditEvent(){
    var sres_in = document.getElementById('inputAddress2');
    sres_in.readonly = false;
    if(sres_in.textContent.length>0){
        EditEvent();
    }
}

function CountryEditEvent(){
    var coun_in = document.getElementById('country');
    coun_in.readonly = false;
    if(coun_in.textContent.length>0){
        EditEvent();
    }
}

function TelEditEvent(){
    var tel_in = document.getElementById('phone');
    tel_in.readonly = false;
    if(tel_in.textContent.length>0){
        EditEvent();
    }
}

function ZipEditEvent(){
    var zip_in = document.getElementById('inputZip'); 
    zip_in.readonly = false;
    if(zip_in.textContent.length>0){
        EditEvent();
    }
}