var citiesList = document.querySelector(".cities");

let json;

let response = await fetch("http://localhost:8080/api/locations/new")
if (response.ok) {
  json = await response.json();
} else {
  // console.log(response)
}

var cities = json.cities;

for (let i = cities.length-1; i >= 0 ; i--) {
    let cityData = cities[i];
    let cityHtml = createCity(cityData);
    citiesList.insertAdjacentHTML('afterBegin', cityHtml);
}

function createCity(cityData) {
    var list = [];
    for(let location of cityData.locations) {
        let data = createCityCardLocation(location);
        list.push(data);
    }

    return `
    <div class="city">
        <p class="name">${cityData.name}</p>
        <div class="cards">${list.join('')}</div>
    </div>
    `
}

function createCityCardLocation(location) {
    var transpotrs = [];
    for (let transport of location.features.transport) {
        let data = createCityTransportRow(transport);
        transpotrs.push(data);
    }

    var popularLocations = [];
    for (let popLoc of location.features.popular_locations) {
        let data = createPopularLocation(popLoc);
        popularLocations.push(data);
    }

    return `
    <div class="card" id="${location.id}">
        <img src="" alt="">
        <div class="data">
        <div class="main">
            <div class="name">${location.name}</div>
            <div class="price">Цена: ${location.price}</div>
        </div>
            <div class="details">
                <div class="list">
                    <p>Транспорт:</p>
                    ${transpotrs.join('')}
                </div>
                <div class="list">
                    <p>Популярные места:</p>
                    ${popularLocations.join('')}
                </div>
            </div>
        </div>
    </div>
    `
}

function createCityTransportRow(transport) {
    return `
    <div class="row">
        <p>${transport.name}</p>
    </div>`
}

function createPopularLocation(popularLocation) {
    return `
    <div class="row">
        <p>${popularLocation.name}</p>
    </div>
    `
} 

let cards = document.querySelectorAll(".card");
cards.forEach(card => {
    card.addEventListener('click', () => {
        window.location.replace(`http://localhost:8080/home?id=${card.id}`)
    });
})