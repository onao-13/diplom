var homeId = sessionStorage.getItem("homeId");

let home;
let res = await fetch(`http://185.187.91.14:8120/api/homes/${homeId}`)
if (res.ok) {
    home = await res.json();
} else {
    console.error(home.error)
}

// set home name
let categotyTitle = document.querySelector(".category-title");
categotyTitle.innerHTML = `<h1>/${home.name}</h1>`;

document.querySelector('.home__description').querySelector('.content').innerHTML = home.description;

let layoutDiv = document.querySelector(".layout");
layoutDiv.querySelector(".content").innerHTML = home.layout;

let greenZoneDiv = document.querySelector(".green-zone");
greenZoneDiv.querySelector(".content").innerHTML = home.green_zone;

let eventDiv = document.querySelector(".events");
eventDiv.querySelector(".content").innerHTML = home.events;

let infrastructure = document.querySelector(".infrastructure");
infrastructure.querySelector(".content").innerHTML = home.infrastructure;

let schoolDiv = document.querySelector(".school");
schoolDiv.querySelector(".content").innerHTML = home.schools;

let transportDiv = document.querySelector(".transport");
transportDiv.querySelector(".content").innerHTML = home.transports;
