const API = "http://77.105.174.83:8085/api/cities";

let res = await fetch(API, {
    mode: 'cors',
    credentials: 'include'
});

if (!res.ok) {
    console.log(res.error);
}

let citiesDiv = document.querySelector(".cities");

let body = await res.json();
body.cities.forEach(city => {
    let cityDiv = createCity(city);
    citiesDiv.prepend(cityDiv);
});

function createCity(city) {
    let cityCard = document.createElement('div');
    cityCard.className = 'card-city city';

    let name = document.createElement('p');
    name.innerText = city.name;
    cityCard.appendChild(name);

    let btnEdit = document.createElement('button');
    btnEdit.innerHTML = 'Редактировать';
    btnEdit.addEventListener('click', () => {
        sessionStorage.setItem("cityId", city.id);
        sessionStorage.setItem("cityName", city.name);
        location.replace("/homes");
    });
    cityCard.appendChild(btnEdit);

    return cityCard;
}
