const HOST = "176.123.164.135:8120"
const SEND_API_URL = `http://${HOST}/api/send-manager-call`

var showManagerFormButtom = document.querySelector('#manager_form_button');
showManagerFormButtom.addEventListener('click', () => {
    showManagerForm(sessionStorage.getItem("homeId"));
});


function showManagerForm(homeId) {
    const form = createForm(homeId);
    const window = createWindow();
    window.appendChild(form);

    document.body.prepend(window);
}

function createWindow() {
    const windowDiv = document.createElement('div');
    windowDiv.className = 'windowDiv';

    window.onclick = function(event) {
        if (event.target == windowDiv) {
            document.body.removeChild(windowDiv);
        }
    }

    return windowDiv;
}

function createForm(homeId) {
    const form = document.createElement('div');
    form.className = 'manager_call_form';
    
    const title = document.createElement('strong');
    title.innerText = "Связаться с менеджером";
    form.appendChild(title);

    const error = document.createElement('p');
    error.id = 'error';
    form.appendChild(error);

    const name = document.createElement('input');
    name.placeholder = 'Имя фамилия';
    form.appendChild(name);

    const telephoneNumber = document.createElement('input');
    telephoneNumber.placeholder = 'Номер телефона';
    form.appendChild(telephoneNumber);

    const send = document.createElement("button");
    send.innerText = 'Оставить обратный звонок';
    send.className = 'button__main';
    send.addEventListener('click', () => {
        if (validate(name.value, telephoneNumber.value)) {
            sendCall(homeId, name.value, telephoneNumber.value);
        }
    });
    form.appendChild(send);

    return form;
}

function validate(name, telephoneNumber) {
    if (name.length == 0) {
        showError("Имя и фамилия не указана");
        return false;
    }

    if (telephoneNumber.length == 0) {
        showError("Номер телефона не указан");
        return false;
    }

    return true;
}

async function sendCall(homeId, name, telephoneNumber) {
    var body = JSON.stringify({
        "name": name,
        "number": telephoneNumber,
        "home_id": homeId
    });

    var response = await fetch(SEND_API_URL, {
        body: body,
        method: "POST"
    });

    if (!response.ok) {
        var responseJson = await response.json();
        showError(responseJson.error);
        return
    }
    
    location.reload();
}

function showError(msg) {
    var error = document.querySelector("#error");
    error.innerHTML = msg;
}
