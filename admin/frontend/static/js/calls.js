const API = "http://77.105.174.83:8085/api/manager-calls";

let res = await fetch(API);

let body = await res.json();

if (!res.ok) {
    console.error(body.error);
}

let callCardsGrid = document.querySelector('.manager_calls');

body.manager_calls.forEach(call => {
    let callCard = createCallCard(call);
    callCardsGrid.prepend(callCard);
});

function createCallCard(call) {
    let card = document.createElement('div');
    card.className = 'call-card';

    let application = document.createElement('p');
    application.className = 'application';
    application.innerText = `Заявка #${call.id}`;
    card.appendChild(application);

    let nameSurname = document.createElement('p');
    nameSurname.innerText = `Имя: ${call.name}`;
    card.appendChild(nameSurname);

    let number = document.createElement('p');
    number.innerText = `Номер: ${call.number}`;
    card.appendChild(number);

    let homeName = document.createElement('strong');
    homeName.innerText = `Дом: ${call.home_name}`;
    card.appendChild(homeName);

    return card;
}
