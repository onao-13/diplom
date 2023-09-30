let locationCard = `<div class="card"> 
<img src="" alt="" class="cover">
<div class="box">
  <div class="title">Название 1</div>
  <div class="position">Позиция 1</div>
  <div class="description">Описание 1</div>
</div>
</div>`;

let articleCard = `<div class="card"> 
<img src="" alt="" class="cover">
<div class="box">
  <div class="title">Название 1</div>
  <div class="description">Описание 1</div>
</div>
</div>`;

let locationList = document.querySelector(".locations__list");
let articleList = document.querySelector(".articles__list");
for (let i = 0; i < 6; i++) {
    locationList.insertAdjacentHTML('afterBegin', locationCard);
    articleList.insertAdjacentHTML('afterBegin', articleCard);
}

