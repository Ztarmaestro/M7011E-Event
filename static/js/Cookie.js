

function setCookie(cname, idToken) {
        console.log("Create cookie");
        document.cookie = cname + "=" + idToken + "; ";
        console.log(document.cookie);   
}

function addEventCookie(eID) {
    
    var cookiename = getCookie("username");
    console.log(cookiename);
    var userID = "username"+ "+" + cookiename;
    console.log(userID);
    var eventID = "event" + "-" + eID;
    console.log(eventID);

    document.cookie = userID + "=" + eventID + "; ";
    console.log(document.cookie);
}

function getCookie(cname) {
    var name = cname + "=";
    var ca = document.cookie.split('-');
    for(var i=0; i<ca.length; i++) {
        var c = ca[i];
        while (c.charAt(0)==' ') c = c.substring(1);
        if (c.indexOf(name) == 0) return c.substring(name.length, c.length);
    }
    return "";
}

function getEventid_Cookie(cname){
    var name = cname + "=";
    var ca = document.cookie.split(';');
    for(var i=0; i<ca.length; i++) {
        var c = ca[i];
        console.log(c);
        while (c.charAt(0)==' ') c = c.substring(1);
        console.log(c);
        if (c.indexOf(name) == 0) return c.substring(name.length, c.length);
        console.log(c);
    }
    return "";

}





