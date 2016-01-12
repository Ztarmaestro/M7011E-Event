
// Sets the users FB idtoken to cookie
function setCookie(cname, idToken) {
    document.cookie = cname + "=" + idToken + "; ";  
    homeredirect();
}

// Adds the event id you clicked on to the cookie
function addEventCookie(eID) {   
    var cookiename = getCookie("username");
    var userID = "username" + "-" + cookiename;
    var eventID = "event" + "+" + eID;
    document.cookie = userID + "=" + eventID + "=";
}

// Gets the FB idtoken from cookie
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

// Gets the Eventid from cookie
function getEventid_Cookie(cname){
    var name = cname + "+";
    var ca = document.cookie.split('=');
    for(var i=0; i<ca.length; i++) {
        var c = ca[i];
        while (c.charAt(0)==' ') c = c.substring(1);
        if (c.indexOf(name) == 0) return c.substring(name.length, c.length);
    }
    return "";

}

function getCookie2(cname) {
    var name = cname + "-";
    var ca = document.cookie.split('=');
    for(var i=0; i<ca.length; i++) {
        var c = ca[i];
        while (c.charAt(0)==' ') c = c.substring(1);
        if (c.indexOf(name) == 0) return c.substring(name.length, c.length);
    }
    return "";
}




