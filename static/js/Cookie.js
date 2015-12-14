

function setCookie(cname, idToken) {
    var cookiename = getCookie("username");
    if (cookiename != idToken) {
        document.cookie = cname + "=" + idToken + "; ";
        console.log(document.cookie);
    }else{
        console.log("Cookie already exist")
    }
}

function setEventCookie(cname, idToken) {
    console.log(cname);
    console.log(idToken);

    document.cookie = cname + "-" + idToken + "; ";
    console.log(document.cookie);
}

function getCookie(cname) {
    var name = cname + "=";
    var ca = document.cookie.split(';');
    for(var i=0; i<ca.length; i++) {
        var c = ca[i];
        while (c.charAt(0)==' ') c = c.substring(1);
        if (c.indexOf(name) == 0) return c.substring(name.length, c.length);
    }
    return "";
}

function getEventid_Cookie(){
    var name = cname + "-";
    var ca = document.cookie.split(';');
    for(var i=0; i<ca.length; i++) {
        var c = ca[i];
        while (c.charAt(0)==' ') c = c.substring(1);
        if (c.indexOf(name) == 0) return c.substring(name.length, c.length);
    }
    return "";

}

function setEventid_Cookie(aid){
    var cookiename = getCookie("username");
    console.log(cookiename);

    setEventCookie(cookiename, aid);

}





