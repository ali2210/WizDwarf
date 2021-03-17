const inputCheckbox = document.getElementsByClassName("span-input-choice")[0];
const spanBox = document.getElementsByClassName("container-choice")[0];
// const form1 = document.getElementsByClassName("form0")[0];

inputCheckbox.addEventListener('change', function() {
    const pay = spanBox.children[0];
    const token = spanBox.children[1];
    // const paypalBtn = form1.children[1];
    // const tokenBtn = form1.children[2];
    if (this.checked) {
        token.style.visibility = "hidden";
        pay.style.visibility = "visible";
        pay.style.color = "green";
        pay.style.marginLeft = "60px";
        pay.style.marginTop = "40px";
        pay.children[1].style.color = "red";
        pay.children[1].style.visibility = "visible";
        token.children[1].style.visibility = "hidden";
        pay.children[0].classList.add(pay.children[1]);
        pay.children[0].classList.remove(pay.children[1]);

    } else {
        token.style.visibility = "visible";
        pay.style.visibility = "hidden";
        token.style.color = "orange";
        token.style.marginLeft = "40px";
        token.style.marginTop = "20px";
        token.children[1].style.color = "red";
        token.children[1].style.visibility = "visible";
        pay.children[1].style.visibility = "hidden";
        token.children[0].classList.add(token.children[1]);
        pay.children[0].classList.remove(pay.children[1]);
    }
});

const inputCheckbox3 = document.getElementsByClassName("vmpay")[0];
const spanBox3 = document.getElementsByClassName("vm-container")[0];
inputCheckbox3.addEventListener('change', function() {

    const pay = spanBox3.children[0];
    const token = spanBox3.children[1];
    if (this.checked) {
        token.style.visibility = "hidden";
        pay.style.visibility = "visible";
        pay.style.color = "green";
        pay.style.marginLeft = "60px";
        pay.style.marginTop = "40px";
        pay.children[1].style.color = "red";
        pay.children[1].style.visibility = "visible";
        token.children[1].style.visibility = "hidden";
        pay.children[0].classList.add(pay.children[1]);
        pay.children[0].classList.remove(pay.children[1]);
    } else {
        token.style.visibility = "visible";
        pay.style.visibility = "hidden";
        token.style.color = "orange";
        token.style.marginLeft = "40px";
        token.style.marginTop = "20px";
        token.children[1].style.color = "red";
        token.children[1].style.visibility = "visible";
        pay.children[1].style.visibility = "hidden";
        token.children[0].classList.add(token.children[1]);
        pay.children[0].classList.remove(pay.children[1]);
    }
});

const inputCheckbox2 = document.getElementsByClassName("paycluster")[0];
const spanBox2 = document.getElementsByClassName("container-cluster")[0];
inputCheckbox2.addEventListener('change', function() {
    const pay = spanBox2.children[0];
    const token = spanBox2.children[1];
    if (this.checked) {
        token.style.visibility = "hidden";
        pay.style.visibility = "visible";
        pay.style.color = "green";
        pay.style.marginLeft = "60px";
        pay.style.marginTop = "40px";
        pay.children[1].style.color = "red";
        pay.children[1].style.visibility = "visible";
        token.children[1].style.visibility = "hidden";
        pay.children[0].classList.add(pay.children[1]);
        pay.children[0].classList.remove(pay.children[1]);
    } else {
        token.style.visibility = "visible";
        pay.style.visibility = "hidden";
        token.style.color = "orange";
        token.style.marginLeft = "40px";
        token.style.marginTop = "20px";
        token.children[1].style.color = "red";
        token.children[1].style.visibility = "visible";
        pay.children[1].style.visibility = "hidden";
        token.children[0].classList.add(token.children[1]);
        pay.children[0].classList.remove(pay.children[1]);
    }
});