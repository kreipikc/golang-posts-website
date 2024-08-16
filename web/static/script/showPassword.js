function show(name) {
    var p = document.getElementById(name);
    p.setAttribute('type', 'text');
}

function hide(name) {
    var p = document.getElementById(name);
    p.setAttribute('type', 'password');
}

var pwShown1 = 0;
var pwShown2 = 0;

document.getElementById("eye1").addEventListener("click", function () {
    if (pwShown1 == 0) {
        pwShown1 = 1;
        show('password1');
    } else {
        pwShown1 = 0;
        hide('password1');
    }
}, false);

document.getElementById("eye2").addEventListener("click", function () {
    if (pwShown2 == 0) {
        pwShown2 = 1;
        show('password2');
    } else {
        pwShown2 = 0;
        hide('password2');
    }
}, false);