

function setCookie(cname, idToken) {
        console.log("Create cookie");
        document.cookie = cname + "=" + idToken + "; ";
        console.log("Cookie already exist")
    
}

function addEventCookie(eID) {
    
    var cookiename = getCookie("username");
    var userID = "username"+ "=" + cookiename;
    var eventID = "event" + "-" + eID;

    document.cookie = userID + "; " + eventID + "/ ";
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

function getEventid_Cookie(cname){
    var name = cname + "-";
    var ca = document.cookie.split('/');
    for(var i=0; i<ca.length; i++) {
        var c = ca[i];
        while (c.charAt(0)==' ') c = c.substring(1);
        if (c.indexOf(name) == 0) return c.substring(name.length, c.length);
    }
    return "";

}

function setEventid_Cookie(aid){

    console.log("sets event id to cookie");

    addEventCookie(cookiename, aid);
    var x = getEventid_Cookie();
    console.log(x);
    
}





