function connected() {
	// body...
	var connectInternet = window.navigator.onLine;
	console.log(connectInternet);

	if (connectInternet) {
		document.getElementById('wifi').className = "fa fa-wifi";
	}

	setTimeout(connected, 1000);
}


	
