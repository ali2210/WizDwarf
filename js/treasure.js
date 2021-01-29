
function Send() {
	// body...

	document.getElementById('in').style.visibility="hidden";
	document.getElementById('send').style.visibility="visible";

}

function Receive(){
	document.getElementById('btn-rece').style.visibility="hidden";
	document.getElementById('rece-txt').style.visibility="visible";
}

function CancelForm(){
	document.getElementById('in').style.visibility="visible";
	document.getElementById('send').style.visibility="hidden";
	document.getElementById('btn-rece').style.visibility="visible";
	document.getElementById('rece-txt').style.visibility="hidden";
}
