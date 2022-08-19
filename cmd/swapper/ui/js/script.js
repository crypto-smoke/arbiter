var form = document.getElementById('swapform');
var inputs = form.getElementsByTagName('input');
var balance0 = document.getElementById('balance0');
let lastCalc;

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

const updateData = function (getPrice) {
    fetch('/update', {
        method: 'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            get_price: getPrice,
            get_balance: true,
            address: document.getElementById("address").value,
            quote: document.getElementById("quoteTokenAddress").value,
            base: document.getElementById("basetoken").value,
            router: document.getElementById("routerAddress").value,
        })
    })
        .then(response => response.json())
        .then(data => {
            if (data.price != null) {
                document.getElementById("price").value = data.price
            }

            if (data.base_balance != null) {
                document.getElementById("baseBalance").value = data.base_balance

            }
            if (data.quote_balance != null) {
                document.getElementById("quoteBalance").value = data.quote_balance
            }
            quoteTokenName.innerHTML = data.quote_name+ "["+data.quote_symbol+"]"
            amountQuoteTicker.innerHTML = data.quote_symbol
            balanceQuoteTicker.innerHTML = data.quote_symbol

            baseTokenName.innerHTML = data.base_name + "["+data.base_symbol+"]"
            priceBaseTicker.innerHTML = data.base_symbol
            totalBaseTicker.innerHTML = data.base_symbol
        })
        .catch(err => {
            //error block
        });
};

function loadToken(id, address) {
    document.getElementById("token" + id + "name").innerHTML = address;
}

function change(id, useValue, divisor) {
    document.getElementById(id).value = document.getElementById(useValue).value / divisor;
    if (id === "amount") {
        calcTotal()
    }
    if (id === "total") {
        calcAmount()
    }
}

function clearAmt(id) {
    document.getElementById(id).value = "";
}


function postform(isBuy) {
    let swapPath
    let amountIn
    let amountOut
    if (isBuy) {
        swapPath = [
            document.getElementById("basetoken").value,
            document.getElementById("quoteTokenAddress").value,
        ]
        amountIn = total.value
        amountOut = amount.value
    } else {
        swapPath = [
            document.getElementById("quoteTokenAddress").value,
            document.getElementById("basetoken").value,
        ]
        amountIn = amount.value
        amountOut = total.value
    }

    fetch('/swap', {
        method: 'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            rpc_url: rpcAddresss.value,
            router_address: routerAddress.value,
            address: address.value,
            swap_path: swapPath,
            amount_in: parseFloat(amountIn),
            amount_out: parseFloat(amountOut),
            gas_gwei: parseInt(gwei.value),
            slippage: parseFloat(slippage.value),
            is_buy: isBuy,
        })
    })
        .then((res) => {
            return res.text();
        })
        .then((txt) => {
            console.log(txt);
            log.value += "\n" + "Txn submitted"
        })
        .catch((err) => {
            console.log(err);
            log.value += "\n" + err
        });
    return false;
}

const logMessage = (message) => {
    var line = document.createElement('div');
    line.innerHTML=message
    log.appendChild(line)
   log.scrollTop = log.scrollHeight
}
window.setInterval(function () {
    // balance0.value++;
      logMessage("lel")
}, 1000);

updateData(autoPrice.checked);
window.setInterval(function () {
    updateData(autoPrice.checked);
}, 5000)

const calcTotal = function () {
    document.getElementById("total").value = document.getElementById("amount").value * document.getElementById("price").value
    lastCalc = calcTotal
}

function calcAmount() {
    document.getElementById("amount").value = document.getElementById("total").value / document.getElementById("price").value
    lastCalc = calcAmount

}