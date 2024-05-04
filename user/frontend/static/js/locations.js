const HOST = "77.105.174.83:8120"

var citiesList = document.querySelector(".cities");

let json;

let response = await fetch(`http://${HOST}/api/cities`)
if (response.ok) {
  json = await response.json();
} else {
  // console.log(response)
}

const MAX_CITY_COUNT = 6;

var cities = json.cities;

for (let i = cities.length-1; i >= 0 ; i--) {
    let cityData = cities[i];
    let cityHtml = createCity(cityData);
    citiesList.insertAdjacentHTML('afterBegin', cityHtml);
}

function createCity(cityData) {
    var list = [];
    if (cityData.homes != null) {
        for(let location of cityData.homes) {
            let data = createCityCardLocation(location);
            list.push(data);
        }
    }

    const city = document.createElement('div');
    city.className = 'city';

    const name = document.createElement('p');
    name.className = 'name';
    name.innerText = cityData.name;
    city.appendChild(name);

    const cards = document.createElement('div');
    cards.className = 'cards';
    city.appendChild(cards);
    cards.insertAdjacentHTML('beforebegin', list.join(''));

    if (list.length > MAX_CITY_COUNT) {
        const showMoreBtn = document.createElement('button');
        showMoreBtn.className = 'button__main';
        showMoreBtn.innerText = 'Показать больше';
        showMoreBtn.addEventListener('click', () => {
            console.log("add change page");
        });
        city.appendChild(showMoreBtn);
    }

    return `
    <div class="city">
        <p class="name">${cityData.name}</p>
        <div class="cards">${list.join('')}</div>
    </div>
    `
}

function createCityCardLocation(location) {
    var imgCover = "";
    if (location.images != null) {
        imgCover = location.images[0].url;
    }
    return `
    <div class="card" id="${location.id}">
        <img src="${imgCover}" alt="">
        <div class="data">
        <div class="main">
            <div class="name">${location.name}, ${location.street}</div>
            <div class="price">Цена: ${location.price}</div>
        </div>
        <div class="details">
            <div class="list">
                <title>Транспорт</title>
                <p>${location.transports}</p>
            </div>
            <div class="list">
                <title>Популярные места</title>
                <p>${location.popular_locations}</p>
            </div>
        </div>
        </div>
    </div>
    `
}

let cards = document.querySelectorAll(".card");
cards.forEach(card => {
    card.addEventListener('click', () => {
        sessionStorage.setItem("homeId", card.id)
        window.location.replace(`http://${HOST}/home`)
    });
})