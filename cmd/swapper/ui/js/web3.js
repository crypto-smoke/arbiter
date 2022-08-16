
if (typeof window.ethereum !== 'undefined') {
    console.log('MetaMask is installed!');
}
const ethereumButton = document.querySelector('.enableEthereumButton');
const showAccount = document.querySelector('.showAccount');

ethereumButton.addEventListener('click', () => {
    getAccount();
});
if (!window.ethereum) {
    console.log("no wallet installed")
}

async function getAccount() {
    const accounts = await ethereum.request({ method: 'eth_requestAccounts' });
    const account = accounts[0];
    showAccount.innerHTML = account;
    window.onload = (event) => {
        console.log('page is fully loaded');

    };
}