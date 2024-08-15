function show() {
    var p = document.getElementById('password');
    p.setAttribute('type', 'text');
}

function hide() {
    var p = document.getElementById('password');
    p.setAttribute('type', 'password');
}

function swap() {
    var p = document.getElementById('img_eye');
    console.log(p.src)
    if (p.getAttribute("src") == "../static/img/eye_close.png") p.setAttribute("src", "../static/img/eye_open.png");
    else p.setAttribute("src", "../static/img/eye_close.png");
}

var pwShown = 0;

document.getElementById("eye").addEventListener("click", function () {
    if (pwShown == 0) {
        pwShown = 1;
        show();
    } else {
        pwShown = 0;
        hide();
    }
}, false);