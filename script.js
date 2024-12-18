let bitcoin = 0;
let farms = 0;
let animals = 0;

function plantBitcoin() {
    bitcoin += 10;
    updateStats();
}

function buyFarm() {
    if (bitcoin >= 50) {
        farms++;
        bitcoin -= 50;
        updateStats();
    } else {
        alert("Not enough Bitcoin to buy a farm!");
    }
}

function buyAnimals() {
    if (bitcoin >= 100) {
        animals++;
        bitcoin -= 100;
        updateStats();
    } else {
        alert("Not enough Bitcoin to buy animals!");
    }
}

function updateStats() {
    document.getElementById("bitcoin").innerText = bitcoin;
    document.getElementById("farms").innerText = farms;
    document.getElementById("animals").innerText = animals;
}
