// Раскрытие
function show(name) {
    var p = document.getElementById(name);
    p.setAttribute('type', 'text');
}

// Скрытие
function hide(name) {
    var p = document.getElementById(name);
    p.setAttribute('type', 'password');
}

// Тумблеры
var pwShown1 = false;
var pwShown2 = false;

// Скрывает/раскрывает пароль для кнопки id="eye1"
document.getElementById("eye1").addEventListener("click", function () {
    if (!pwShown1) {
        pwShown1 = true;
        show('password1');
    } else {
        pwShown1 = false;
        hide('password1');
    }
}, false);

// Скрывает/раскрывает пароль для кнопки id="eye2"
document.getElementById("eye2").addEventListener("click", function () {
    if (!pwShown2) {
        pwShown2 = true;
        show('password2');
    } else {
        pwShown2 = false;
        hide('password2');
    }
}, false);