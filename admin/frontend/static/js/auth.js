const SERVER_URL = "http://176.123.164.135:8085/api/auth";

let btnLogin = document.querySelector("#login");
btnLogin.addEventListener('click', () => {
    login();
});

async function login() {
    var username = document.querySelector("#username").value;
    var password = document.querySelector("#password").value;

    if (username.length == 0) {
        showError("Имя пользователя не указано");
        return;
    }

    if (password.length == 0) {
        showError("Пароль не указан");
        return;
    }

    var result = await sendRequest(username, password);

    if (!result.ok) {
        var response = await result.json();
        showError(response.error);
        return
    } else {
        location.replace("/panel");
    }
}

function showError(message) {
    var err = document.querySelector("#error")
    err.innerHTML = message;
}

async function sendRequest(username, password) {
    var body = JSON.stringify({
        "username": username,
        "password": password
    });

    var result = await fetch(SERVER_URL, {
        body: body,
        method: "POST"
    });

    return result;
}
