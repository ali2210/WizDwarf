
$("#file").on("change", () => {
  if ($("#file")[0].files.length !== 0) {
    $(".selected-file")[0].innerHTML = `Seleected File: ${$("#file")[0].files[0].name}`;
    $(".selected-file").addClass("file-uploaded")
  }
})

const bodyAlertSys = document.getElementsByClassName("container-alert")[0];
const childLeft = bodyAlertSys.children[0];
const childRight = bodyAlertSys.children[1];
const closeFailBtn = childLeft.children[2];
const closeSuccessBtn = childRight.children[1];

function onrequestaction() {
  childLeft.style.visibility = "hidden";
}

closeFailBtn.addEventListener('click', onrequestaction, false);

function onrequestsuccess() {
  childRight.style.visibility = "hidden";
}

closeSuccessBtn.addEventListener('click', onrequestsuccess, false);




