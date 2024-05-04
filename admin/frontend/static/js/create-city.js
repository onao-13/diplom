const API = "http://176.123.164.135:8085/api/cities";

let btnCreateCity = document.querySelector('#create-city');
btnCreateCity.addEventListener('click', () => {
    let cityName = document.querySelector('#city-name');
    send(cityName.value);
});

async function send(cityName) {
    let body = JSON.stringify({
        "name": cityName
    });

    let res = await fetch(API, {
        method: "POST",
        body: body,
        credentials: 'include'
    });

    if (!res.ok) {
        console.error(res.error);
    }

    location.reload();
}