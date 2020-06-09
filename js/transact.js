

function ChangePic() {
	// body...

	var img = document.getElementById('change').src = "/images/Female.png";
	console.log(img);
	document.getElementById('female').disabled = false;
	var female = document.getElementById('female').style.visibility="visible";
	var male = document.getElementById('male').style.visibility="hidden";


}

function Reverse(){
	var reset = document.getElementById('change').src = "/images/myAvatar.png";
	console.log(reset);
	var male = document.getElementById('male').style.visibility="visible";
	var female = document.getElementById('female').style.visibility="hidden";
}

function PayByEth(){
	document.getElementById('btn').style.visibility = "hidden";
	document.getElementById('add').style.visibility = "visible";
	document.getElementById('your').style.visibility = "visible";
	document.getElementById('time').style.visibility = "visible";
	document.getElementById('transact').style.visibility = "visible";
	document.getElementById('transact').disabled = false;
	document.getElementById('ok').style.visibility = "visible";
	document.getElementById('ok').disabled = false;
	document.getElementById('cancel').style.visibility = "visible";
	document.getElementById('cancel').disabled = false;
}

function ReceiveByEth() {
	// body...
	document.getElementById('btn-rec-0').style.visibility = "hidden";
	document.getElementById('recepit').style.visibility = "visible";
	document.getElementById('your-add').style.visibility = "visible";
	document.getElementById('gift').style.visibility = "visible";
	document.getElementById('gift').disabled = false;
	document.getElementById('reset').style.visibility = "visible";
	document.getElementById('reset').disabled = false;
}

function Cancel(){
	document.getElementById('btn').style.visibility = "visible";
	document.getElementById('add').style.visibility = "hidden";
	document.getElementById('your').style.visibility = "hidden";
	document.getElementById('time').style.visibility = "hidden";
	document.getElementById('transact').style.visibility = "hidden";
	document.getElementById('transact').disabled = true;
	document.getElementById('ok').style.visibility = "hidden";
	document.getElementById('ok').disabled = true;
	document.getElementById('cancel').style.visibility = "hidden";
	document.getElementById('cancel').disabled = true;	
}

function ClearWindow() {
	// body...
	document.getElementById('btn-rec-0').style.visibility = "visible";
	document.getElementById('recepit').style.visibility = "hidden";
	document.getElementById('your-add').style.visibility = "hidden";
	document.getElementById('gift').style.visibility = "hidden";
	document.getElementById('gift').disabled = true;
	document.getElementById('reset').style.visibility = "hidden";
	document.getElementById('reset').disabled = true;	
}

