let API = `http://77.105.174.83:8085/api/cities/${sessionStorage.getItem('cityId')}/homes`;
let method;

let homeId = sessionStorage.getItem('homeId');
if (homeId == null) {
    document.querySelector('.top').querySelector('h2').innerText = 'Создание дома';
    method = 'POST';
} else {
    let res = await fetch(`${API}/${homeId}`);
    method = 'PATCH';

    let body = await res.json();

    if (!res.ok) {
        console.error(body.error);
    }

    document.querySelector('#name').value = body.name;
    document.querySelector('#street').value = body.street;
    document.querySelector('#price').value = body.price;
    document.querySelector('#description').value = body.description;
    document.querySelector('#transports').value = body.transports;
    document.querySelector('#locations').value = body.popular_locations;
    document.querySelector('#greenZone').value = body.green_zone;
    document.querySelector('#infrastructure').value = body.infrastructure;
    document.querySelector('#schools').value = body.schools;
    document.querySelector('#events').value = body.events;
    document.querySelector('#layout').value = body.layout;
    var images = document.querySelectorAll('.images');

    if (body.images != null) {
        images.forEach((input, i) => {
            var imgUrl = body.images[i];
            if (imgUrl !== undefined) {
                input.id = imgUrl.id;
                input.value = imgUrl.url;
            } else {
                input.id = '0';
            }
        });
    }
}

async function saveHome(id, name, street, price, description, transports, locations, greenZone, infrastructure, schools, events, layout, images) {
    let body = JSON.stringify({
        "name": name,
        "street": street,
        "price": price,
        "images": images,
        "transports": transports,
        "popular_locations": locations,
        "description": description,
        "green_zone": greenZone,
        "infrastructure": infrastructure,
        "schools": schools,
        "events": events,
        "layout": layout
    });

    console.log(body);

    if (id != null) {
        API += `/${id}`;
    }

    let res = await fetch(API, {
        method: method,
        body: body
    });

    let bodyRes = await res.json();
    if (!res.ok) {
        console.error(bodyRes.error);
    }
}

async function deleteHome(id) {
    let res = await fetch(`${API}/${id}`, {
        method: 'DELETE',
    });

    let body = await res.json();

    if (!res.ok) {
        console.error(body.error);
    }
}

let saveBtn = document.querySelector('#save');
saveBtn.addEventListener('click', () => {
    let id = sessionStorage.getItem('homeId');
    let name = document.querySelector('#name').value;
    let street = document.querySelector('#street').value;
    let price = document.querySelector('#price').value;
    let description = document.querySelector('#description').value;
    let transports = document.querySelector('#transports').value;
    let locations = document.querySelector('#locations').value;
    let greenZone = document.querySelector('#greenZone').value;
    let infrastructure = document.querySelector('#infrastructure').value;
    let schools = document.querySelector('#schools').value;
    let events = document.querySelector('#events').value;
    let layout = document.querySelector('#layout').value;
    let images = [];
    document.querySelectorAll('.images').forEach((input) => {
        if (input.value.length != 0) {
            images.push({
                "id": input.id,
                "url": input.value
            });
        }
    });

    saveHome(id, name, street, price, description, transports, locations, greenZone, infrastructure, schools, events, layout, images);

    location.replace('/homes');
});


let deleteBtn = document.querySelector('#delete');
deleteBtn.addEventListener('click', () => {
    deleteHome(sessionStorage.getItem('homeId'));
    location.replace('/homes');
});
