

// function ChangePic() {
// 	// body...

// 	var img = document.getElementById('change').src = "/images/Female.png";
// 	console.log(img);
// 	document.getElementById('female').disabled = false;
// 	var female = document.getElementById('female').style.visibility = "visible";
// 	var male = document.getElementById('male').style.visibility = "hidden";


// }

// function Reverse() {
// 	var reset = document.getElementById('change').src = "/images/myAvatar.png";
// 	console.log(reset);
// 	var male = document.getElementById('male').style.visibility = "visible";
// 	var female = document.getElementById('female').style.visibility = "hidden";
// }

// function PayByEth() {
// 	document.getElementById('btn').style.visibility = "hidden";
// 	document.getElementById('add').style.visibility = "visible";
// 	document.getElementById('your').style.visibility = "visible";
// 	document.getElementById('time').style.visibility = "visible";
// 	document.getElementById('transact').style.visibility = "visible";
// 	document.getElementById('transact').disabled = false;
// 	document.getElementById('ok').style.visibility = "visible";
// 	document.getElementById('ok').disabled = false;
// 	document.getElementById('cancel').style.visibility = "visible";
// 	document.getElementById('cancel').disabled = false;
// 	document.getElementById('more').style.visibility = "visible";
// 	document.getElementById('more').disabled = false;
// }


// function Cancel() {
// 	document.getElementById('btn').style.visibility = "visible";
// 	document.getElementById('add').style.visibility = "hidden";
// 	document.getElementById('your').style.visibility = "hidden";
// 	document.getElementById('time').style.visibility = "hidden";
// 	document.getElementById('transact').style.visibility = "hidden";
// 	document.getElementById('transact').disabled = true;
// 	document.getElementById('ok').style.visibility = "hidden";
// 	document.getElementById('ok').disabled = true;
// 	document.getElementById('amount').style.visibility = "hidden";
// 	document.getElementById('amount').disabled = true;

// 	document.getElementById('cancel').style.visibility = "hidden";
// 	document.getElementById('cancel').disabled = true;
// 	document.getElementById('more').style.visibility = "hidden";
// 	document.getElementById('more').disabled = true;
// }

// function CompleteTransact() {
// 	document.getElementById('amount').style.visibility = "visible";
// 	document.getElementById('amount').disabled = false;

(function () {
	var divBox = document.getElementsByClassName('div-meta');
	var network = "";
	async function connectEthereum() {
		network = await detectEthereumProvider();

		if (network === window.ethereum && window.ethereum.isMetaMask && window.ethereum.isConnected()) {
			divBox[0].children[1].children[0].style.visibility = "hidden";
			divBox[0].children[1].style.visibility = "hidden";
		}
		const _chainId = await ethereum.request({
			method: 'eth_chainId'
		});

		handleChainID(_chainId);
		ethereum.on('chainChanged', handleChainID);

		function handleChainID(_chainId) { }

		ethereum
			.request({ method: 'eth_accounts' })
			.then(handleAccountsChanged)
			.catch((err) => {
				console.error("error got", err);
			});
		ethereum.on('accountsChanged', handleAccountsChanged)

		var accPromise = await window.ethereum.request({ method: 'eth_requestAccounts' })
		function handleDivHTml(divBox) {
			divBox.children[1].children[0].style.visibility = "hidden";
			divBox.children[0].children[0].style.visibility = "visible";
			divBox.children[0].children[0].style.marginLeft = "50px";
			divBox.children[0].children[0].style.marginTop = "20px";
			setInterval(function () {
				window.location.reload();
			}, 7000);
		}
		function handleAccountsChanged(accounts) {
			var metamaskAcc = null;
			if (accounts.length === 0 && window.ethereum.isConnected()) {
				divBox[0].children[1].children[0].style.visibility = "visible";
				divBox[0].children[1].style.visibility = "visible";
				divBox[0].children[1].children[0].style.marginLeft = "50px";
				divBox[0].children[1].children[0].style.marginTop = "20px";
				divBox[0].children[0].children[0].style.visibility = "hidden";
			} else {
				if (metamaskAcc != accounts[0]) {
					metamaskAcc = accounts[0];
					handleDivHTml(divBox[0]);
				} else {
					handleDivHTml(divBox[0]);
				}
			}
		}

	}
	window.addEventListener('load', connectEthereum, false);
})();

const inputCheck = document.getElementById("checkouts");
const divCheck = document.getElementById("describe");
const spanInput = document.getElementById("square");
const spanCross = document.getElementById("cross");


const instance = document.getElementById("cloudInstance");
const divBody = document.getElementById("data");
const checkIcon = document.getElementById("sqr");
const checkCross = document.getElementById("crss");


const vmMac = document.getElementById("vm");
const body = document.getElementById("multivm");
const checkPos = document.getElementById("pos");
const checkNeg = document.getElementById("neg");
const bodyAlertSys = document.getElementsByClassName("container-alert")[0];
const childLeft = bodyAlertSys.children[0];
const childRight = bodyAlertSys.children[1];
const closeFailBtn = childLeft.children[2];
const closeSuccessBtn = childRight.children[1];

function input() {
	if (inputCheck.checked) {
		spanInput.style.visibility = "visible";
		spanCross.style.visibility = "hidden";
		divCheck.classList.add(spanCross);
		divCheck.classList.remove(spanCross);
	} else {
		spanInput.style.visibility = "hidden";
		spanCross.style.visibility = "visible";
		divCheck.classList.add(spanInput);
		divCheck.classList.remove(spanInput);
	}


}
inputCheck.addEventListener('change', input, false);


function clusterForm() {
	if (instance.checked) {
		checkIcon.style.visibility = "visible";
		checkCross.style.visibility = "hidden";
		divBody.classList.add(checkCross);
		divBody.classList.remove(checkCross);
	} else {
		checkIcon.style.visibility = "hidden";
		checkCross.style.visibility = "visible";
		divBody.classList.add(checkIcon);
		divBody.classList.remove(checkIcon);
	}
}

instance.addEventListener('change', clusterForm, false);

function vmForm() {
	if (vmMac.checked) {
		checkPos.style.visibility = "visible";
		checkNeg.style.visibility = "hidden";
		body.classList.add(checkNeg);
		body.classList.remove(checkNeg);
	} else {
		checkPos.style.visibility = "hidden";
		checkNeg.style.visibility = "visible";
		body.classList.add(checkPos);
		body.classList.remove(checkPos);
	}
}

vmMac.addEventListener('change', vmForm, false);

function onrequestaction() {
	childLeft.style.visibility = "hidden";
}

closeFailBtn.addEventListener('click', onrequestaction, false);

function onrequestsuccess() {
	childRight.style.visibility = "hidden";
}

closeSuccessBtn.addEventListener('click', onrequestsuccess, false);


