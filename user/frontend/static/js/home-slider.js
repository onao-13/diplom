// set home name
let categotyTitle = document.querySelector(".category-title");
categotyTitle.innerHTML = `<h1>/${home.name}</h1>`; 

// fill slider
var slider = document.querySelector(".slider__slides");
var slideNum = 0;

home.images.forEach(img => {
    if (slideNum == 0) {
        slider.insertAdjacentHTML("beforeend", slideActive(img.location, slideNum));
    } else {
        slider.insertAdjacentHTML("beforeend", slideInnactive(img.location, slideNum));
    }
    slideNum++;
});

let slides = document.querySelectorAll(".slide");
slides.forEach((slide, id) => {
    slide.addEventListener("click", () => setActive(id));
});

function slideActive(locName, num) {
    return `<div class="slide active" id="${num}"> 
    <div class="name">/${locName}</div>
        <div class="bubble">
            <img src="../icons/arrow-right.svg" alt="">
        </div>
    </div>`
} 

function slideInnactive(locName, num) {
    return `<div class="slide innactive" id="${num}">
        <div class="name">/${locName}</div>
        <div class="bubble">
            <img src="../icons/arrow-right.svg" alt="">
        </div>
    </div>`
}

function setActive(id) {
    let slide = document.getElementById(`${id}`);
    slide.classList.remove('innactive');
    slide.classList.add('active');

    slides.forEach((slide, i) => {
        if (i != id) {
            slide.classList.remove('active');
            slide.classList.add('innactive');
        }
    });
}