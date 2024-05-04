const HOST = "77.105.174.83:8120"

var homeId = sessionStorage.getItem("homeId");

let home;
let res = await fetch(`http://${HOST}/api/homes/${homeId}`)
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


// SLIDER 

let imagesLen = home.images.length;
var images = home.images;

function createSlide(name, i) {
    let slide = document.createElement('div');
    slide.className = name;
    slide.id = i;

    let slideName = document.createElement('div');
    slideName.innerText = "/Слайд " + i;

    slide.appendChild(slideName)

    return slide;
} 

var slides = document.querySelector(".slider__slides");
//  create selected
slides.appendChild(createSlide("slide-selected", 0))

// create unselected
for (let i = 1; i < imagesLen; i++) {
    slides.appendChild(createSlide("slide-unselected", i))
}

// константы классов в html, чтобы не запоминать как они пишутся
const classSlideSelected = "slide-selected";
const classSlideUnselected = "slide-unselected";

// получаем элементы слайдера по классу
let selectedSlide = document.getElementsByClassName(classSlideSelected)[0]; //берем 0-ой элемнент, потому что активный слайдер может быть только один
let unselectedSlides = document.getElementsByClassName(classSlideUnselected);

// объединям элементы в один массив
var slides = Array.from(unselectedSlides);
slides.push(selectedSlide)
// подключаем к слайдам слушатель кликов
slides.forEach(slide => {
    slide.addEventListener("click", () => selectSlide(slide.id))
});

// хранит айди активного слайдера. по умолчанию берется айди с названием класса slide-selected
let selectedSlideId = selectedSlide.id;

/**
 * функция для выбора слайда
 * @param {BigInteger} id - айди слайда
 */
function selectSlide(id) {
    console.log(id);
    // обновляет айди
    selectedSlideId = id;
    // обновляет состояние слайдера
    updateSlider();
}

/** 
 * функция обеновления состояния слайдера
*/
function updateSlider() {
    slides.forEach(slide => {
        // если айди слайдера равен выбранному на данный момент слайду, ему присваивается класс slide-selected, иначе - slide-unselected
        if (slide.id == selectedSlideId) {
            updateImage()
            slide.classList.remove(classSlideUnselected);
            slide.classList.add(classSlideSelected);
        } else {
            slide.classList.remove(classSlideSelected);
            slide.classList.add(classSlideUnselected);
        }
    });
}

// устанавливает картинку по умолчанию для активного слайда
let slideImage = document.getElementsByClassName("slide-image")[0];
slideImage.src = images[0].url;

/**
 *  функция обновления картинки. меняет путь изображения слайдера в зависимости от выбранного айди слайдера 
*/
function updateImage() {
    slideImage.src = images[parseInt(selectedSlideId)].url;
}

// функция, меняющая картинку слайдера каждые 5 секунд
setInterval(
    () => {
        selectSlide(Math.round(Math.random() * (slides.length - 1)))
    },
    5000
)
