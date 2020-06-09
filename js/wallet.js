function connected() {
	// body...
	var connectInternet = window.navigator.onLine;
	console.log(connectInternet);

	if (connectInternet) {
		document.getElementById('wifi').className = "fa fa-wifi";
		alert("You're connected");
	}else{
		alert('Check your connection before proceed!');
	}
}


	
