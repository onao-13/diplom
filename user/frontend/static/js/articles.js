let articles = document.querySelector(".articles");

var json;

let response = await fetch("http://77.105.174.83:8120");
if (response.ok) {
    json = await response.json();
}

json.articles.forEach(article => {
    let category = createArticleCategory(article);
    articles.insertAdjacentHTML('afterbegin', category);
});

function createArticleCategory(category) {
    var list = [];
    for (let cat of category.list) {
        let data = createArticle(cat);
        list.push(data);
    }

    return `<div class="category-title">
        <h1>${category.category_name}</h1>
        <div class="category__list">
            ${list.join('')}
        </div>
    </div>`
}

function createArticle(article) {
    return `<div class="card">
        <img src="" alt="">
        <div class="name">${article.name}</div>
        <div class="description">${article.description}</div>
    </div>`
}

/* 
<div class="category-title">
            <h1></h1>
            <div class="article">
                <img src="" alt="">
                <div class="title"></div>
                <div class="description"></div>
            </div>
        </div>
*/
