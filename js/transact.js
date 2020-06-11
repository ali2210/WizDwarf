
// set google loader api....


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
	document.getElementById('more').style.visibility = "visible";
	document.getElementById('more').disabled = false;
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
	document.getElementById('amount').style.visibility = "hidden";
	document.getElementById('amount').disabled = true;
	document.getElementById('nonce').style.visibility = "hidden";
	document.getElementById('nonce').disabled = true;
	document.getElementById('gas').style.visibility = "hidden";
	document.getElementById('gas').disabled = true;

	document.getElementById('cancel').style.visibility = "hidden";
	document.getElementById('cancel').disabled = true;
	document.getElementById('more').style.visibility = "hidden";
	document.getElementById('more').disabled = true;		
}

function CompleteTransact(){
	document.getElementById('amount').style.visibility = "visible";
	document.getElementById('amount').disabled = false;
	document.getElementById('nonce').style.visibility = "visible";
	document.getElementById('nonce').disabled = false;
	document.getElementById('gas').style.visibility = "visible";
	document.getElementById('gas').disabled = false;
}

