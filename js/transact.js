(function() {
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

        function handleChainID(_chainId) {}

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
            setInterval(function() {
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





// const bodyAlertSys = document.getElementsByClassName("container-alert")[0];
// const childLeft = bodyAlertSys.children[0];
// const childRight = bodyAlertSys.children[1];
// const closeFailBtn = childLeft.children[2];
// const closeSuccessBtn = childRight.children[1];

// function onrequestaction() {
//     childLeft.style.visibility = "hidden";
// }

// closeFailBtn.addEventListener('click', onrequestaction, false);

// function onrequestsuccess() {
//     childRight.style.visibility = "hidden";
// }

// closeSuccessBtn.addEventListener('click', onrequestsuccess, false);