function connected() {
	// body...
	var connectInternet = window.navigator.onLine;
	console.log(connectInternet);

	if (connectInternet) {
		document.getElementById('wifi').className = "fa fa-wifi";
	}else{
		alert('Check your connection before proceed!');
	}

	setTimeout(connected, 1000);
}


	
