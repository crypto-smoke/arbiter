var form = document.getElementById('swapform');
var inputs = form.getElementsByTagName('input');
var balance0 = document.getElementById('balance0');

window.setInterval(function() {
    balance0.value++;
}, 1000);

function loadToken(id, address) {
    document.getElementById("token" + id + "name").innerHTML = address;
}

function change(id, divisor) {
    document.getElementById("amount" + id).value = document.getElementById("balance" + id).value / divisor;
}

function clearAmt(id) {
    document.getElementById("amount" + id).value = "";
}



function postform () {
    var form = document.getElementById("swapform");
    var data = new FormData(form);

    fetch("swap", {
        method: "post",
        body: data
    })
        .then((res) => { return res.text(); })
        .then((txt) => { console.log(txt); })
        .catch((err) => { console.log(err); });
    return false;
}