

/**
 * declaration variables
 * userProfileForm specified web form along with avatar option
 */
const userProfileForm = document.getElementsByClassName('profile-form')[0];

/**
 * Editform input fields
 */
const editForm = document.getElementsByClassName('edit')[0];
const editchoice = document.getElementsByClassName('col')[0];

/**
 * Avatar Upload or capture via stream
 */
userProfileForm.children[0].children[0].children[0].children[0].addEventListener("click", event=>{
    if(event.target){
        editForm.style.top = "700px";
     }
    
});


/**
 *  Edit-Option  allow user to add information which is not in database
 *  Edit your Name 
 */ 
editchoice.children[0].children[1].addEventListener("click", event =>{
    console.log(editForm.children[0].firstElementChild.children[0].readOnly);
    editForm.children[0].firstElementChild.children[0].readOnly = false;
    editForm.children[0].firstElementChild.children[0].addEventListener("change", function(){
        console.log(editForm.children[0].firstElementChild.children[0].value.length);
        if(editForm.children[0].firstElementChild.children[0].value.length>0){
            editForm.children[0].firstElementChild.children[0].style.color = "green";
            editForm.children[0].firstElementChild.children[1].style.color = "green";
        }else if(editForm.children[0].firstElementChild.children[0].value.length == 0){
            editForm.children[0].firstElementChild.children[0].style.color = "red";
            editForm.children[0].firstElementChild.children[1].style.color = "red";
        }
    });
});

/**
 * Edit your Lastname
 */

editchoice.children[1].children[1].addEventListener("click", event =>{
    console.log(editForm.children[0].children[1].firstElementChild.readOnly);
    editForm.children[0].children[1].firstElementChild.readOnly = false;
    editForm.children[0].children[1].children[0].addEventListener("change", function(){
        console.log(editForm.children[0].children[1].firstElementChild.value.length);
        if(editForm.children[0].children[1].firstElementChild.value.length>0){
            editForm.children[0].children[1].firstElementChild.style.color = "green";
            editForm.children[0].children[1].lastElementChild.style.color = "green";
        }else if(editForm.children[0].children[1].firstElementChild.value.length == 0){
            editForm.children[0].children[1].firstElementChild.style.color = "red";
            editForm.children[0].children[1].lastElementChild.style.color = "red";
        }
    });
});

/**
 * Edit Your Gender For Males
 */


editchoice.children[2].children[0].children[0].addEventListener("click", event=>{
    console.log(editForm.children[0].children[2].children[0].children[0].children[0].style.border);
    editForm.children[0].children[2].children[0].children[0].style.border = "1px green dotted";
    editForm.children[0].children[2].children[0].children[0].style.color = "green";
    editForm.children[0].children[2].children[0].children[0].style.backgroundColor = "white";
});

/**
 * Edit Your Gender For Female
 */

editchoice.children[2].children[1].children[0].addEventListener("click", event=>{
    console.log("female",editForm.children[0].children[2].children[1].children[0].children[0].style.border);
    editForm.children[0].children[2].children[1].children[0].style.border = "1px green dotted";
    editForm.children[0].children[2].children[1].children[0].style.color = "green";
    editForm.children[0].children[2].children[1].children[0].style.backgroundColor = "white";
});

