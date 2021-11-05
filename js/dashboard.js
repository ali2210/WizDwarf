


const parent_div = document.getElementsByClassName('sign')[0];
const para_div = document.getElementsByClassName('signed_message')[0];
console.log(para_div, parent_div);
$('#signed').on('change', () =>{
  
  console.log("props:" , ($('#signed').prop('checked')));
  if ($('#signed').prop('checked', true)){
    
    // create div, icon and button
    const div = document.createElement('div');
    const verifiedBtn = document.createElement('button');
    const icon = document.createElement('i');
    

    // adjust button props
    verifiedBtn.className = 'btn btn-light verified';
    verifiedBtn.type = 'button';
    verifiedBtn.style.position ='relative';
    verifiedBtn.style.top = '11px';
    verifiedBtn.style.left = '32px';
    verifiedBtn.style.color = 'lawngreen';

    //adjust icon properties
    icon.className = 'fa fa-check-circle';
    icon.ariaHidden = true;

    // icon render
    verifiedBtn.appendChild(icon);

    // user action message
    para_div.style.visibility = 'visible';
    para_div.innerHTML = 'The message have been signed. Please verify your message. 0.25Eth';
    para_div.appendChild(verifiedBtn);
    
    //  newly created div render
    div.appendChild(para_div);

    // message props
    parent_div.style.position = 'relative';
    parent_div.style.left = '-23px';
    parent_div.style.width = '249px';
    parent_div.style.boxShadow = '5px 5px white';
    parent_div.style.zIndex = '2';
    parent_div.style.borderRadius = '35px';
    parent_div.style.top = '-1px';
    parent_div.style.display = 'flex'; 
    
    // message box render
    parent_div.appendChild(div);
  }
  if ($('#signed').prop('checked', false)) {
    parent_div.children[0].style.visibility = 'hidden';
    console,log("checked false");
  }
});

$("#file").on("change", () => {
  if ($("#file")[0].files.length !== 0) {
    $(".selected-file")[0].innerHTML = `Seleected File: ${$("#file")[0].files[0].name}`;
    $(".selected-file").addClass("file-uploaded");
  }
});
