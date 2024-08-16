function swap(nameId) {
    var p = document.getElementById(nameId);
    if (p.getAttribute("src") == "../static/img/eye_close.png") p.setAttribute("src", "../static/img/eye_open.png");
    else p.setAttribute("src", "../static/img/eye_close.png");
}
