var form = document.getElementById('swapform');
var inputs = form.getElementsByTagName('input');
var balance0 = document.getElementById('balance0');


const updatePrice = function () {
    fetch('/tokenPrice?' + new URLSearchParams({
        base: '0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174',
        quote: "0x22a31bD4cB694433B6de19e0aCC2899E553e9481",
        router: "0x51aBA405De2b25E5506DeA32A6697F450cEB1a17",
    }), {
        method: 'get'
    })
        .then(response => response.json())
        .then(jsonData => {
            document.getElementById("price").value = jsonData
            console.log(jsonData)
        })
        .catch(err => {
            //error block
        });
};

const updateData = function () {
    fetch('/update', {
        method: 'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            get_price: true,
            get_balance: true,
            address: document.getElementById("address").value,
            quote: document.getElementById("token0address").value,
            base: document.getElementById("basetoken").value,
            router: document.getElementById("routerAddress").value,
        })
    })
        .then(response => response.json())
        .then(data => {
            document.getElementById("price").value = data.price
            document.getElementById("baseBalance").value = data.base_balance
            document.getElementById("quoteBalance").value = data.quote_balance
            console.log(data)
        })
        .catch(err => {
            //error block
        });
};

function loadToken(id, address) {
    document.getElementById("token" + id + "name").innerHTML = address;
}

function change(id,useValue, divisor) {
    document.getElementById(id).value = document.getElementById(useValue).value / divisor;
    if (id === "amount") {
    calcTotal()
    }
    if (id === "total") {
        calcAmount()
    }
}

function clearAmt(id) {
    document.getElementById( id).value = "";
}


function postform() {
    var form = document.getElementById("swapform");
    var data = new FormData(form);

    fetch("swap", {
        method: "post",
        body: data
    })
        .then((res) => {
            return res.text();
        })
        .then((txt) => {
            console.log(txt);
        })
        .catch((err) => {
            console.log(err);
        });
    return false;
}

window.setInterval(function () {
   // balance0.value++;
}, 1000);
updateData();
window.setInterval(function () {
    if (!document.getElementById("autoPrice").checked) {
        return
    }
    updatePrice()
}, 5000)
function calcTotal() {
    document.getElementById("total").value = document.getElementById("amount").value * document.getElementById("price").value
}
function calcAmount() {
    document.getElementById("amount").value = document.getElementById("total").value / document.getElementById("price").value

}