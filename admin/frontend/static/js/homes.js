const cityId = sessionStorage.getItem("cityId");
const HOST = "http://77.105.174.83:8085/api";
const API = `${HOST}/cities/${cityId}/homes`;

let res = await fetch(API);

let body = await res.json();

if (!res.ok) {
    console.error(body.error);
}

let btnSave = document.querySelector("#save");
btnSave.addEventListener('click', () => {
    let cityName = document.querySelector("#name").value;
    saveCityName(cityName);
    location.reload();
})

let btnDelete = document.querySelector("#delete");
btnDelete.addEventListener('click', () => {
    deleteCity();
    location.replace("/cities");
});

let cityName = document.querySelector("#name");
cityName.value = sessionStorage.getItem("cityName");

function saveCityName(name) {
    const API = `${HOST}/cities/${cityId}`;

    let body = JSON.stringify({
        "name": name
    });

    fetch(API, {
        method: "PUT",
        body: body,
        credentials: 'include'
    });
    
    sessionStorage.setItem("cityName", name);
}

function deleteCity() {
    const API = `${HOST}/cities/${cityId}`;
    fetch(API, {
        method: 'DELETE',
        credentials: 'include'
    });
}

function createCity() {

}

function editHome(home) {
    sessionStorage.setItem("homeId", home.id);
    location.replace("/home");
}

const homesGrid = document.querySelector("#homes-grid");
body.homes.forEach(home => {
    homesGrid.appendChild(showCityCard(home));
});

function showCityCard(home) {
    const homeDiv = document.createElement('div');
    homeDiv.className = 'home';
    
    const previewImg = document.createElement('img');
    homeDiv.appendChild(previewImg);


    const cityData = document.createElement('div');
    cityData.className = 'city-data';
    homeDiv.appendChild(cityData);

    const name = document.createElement('p');
    name.className = 'name';
    name.innerText = 'Название: ' + home.name;
    cityData.appendChild(name);

    const street = document.createElement('p');
    street.className = 'price';
    street.innerText = 'Улица: ' + home.street;
    cityData.appendChild(street);

    const price = document.createElement('p');
    price.className = 'price';
    price.innerText = 'Цена: ' + home.price;
    cityData.appendChild(price);

    const btnEdit = document.createElement('button');
    btnEdit.innerHTML = 'Редактировать';
    btnEdit.id = 'edit-home';
    btnEdit.addEventListener('click', () => {
        editHome(home);
    });
    homeDiv.appendChild(btnEdit);

    return homeDiv;
}
