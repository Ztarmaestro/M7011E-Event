

function setCookie(cname, idToken, cevent, eventID) {

    document.cookie = cname + "=" + idToken + "; " cevent "=" + eventID + "; ";

}

function makeEventcookie (eventID) {

	var user = getCookie("username");
	setCookie ("username", user, "event", eventID);

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




