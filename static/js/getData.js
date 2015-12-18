
//Different function that we call for getting/creating events/user

function getAllEvent(){
  var xmlHttp = null;

  xmlHttp = new XMLHttpRequest();
  xmlHttp.onreadystatechange=function() {
    if (xmlHttp.readyState==4 && xmlHttp.status==200) {
        var json = xmlHttp.responseText;
        //console.log(json); 
        var obj = JSON.parse(json); 
        //console.log(obj); 
        filterAllEvent(obj);
    }
    else{
      return "Error";
    }
  };
  
    xmlHttp.open( "GET", "http://130.240.170.56:8000/event", false );
    xmlHttp.send( null ); 
}

function getEvent(id){
  var xmlHttp = null;

  xmlHttp = new XMLHttpRequest();
  xmlHttp.onreadystatechange=function() {
    if (xmlHttp.readyState==4 && xmlHttp.status==200) {
        var json = xmlHttp.responseText;
        //console.log(json);
        var obj = JSON.parse(json);  
        //console.log(obj);
        filterEvent(obj);
        geoaddress(obj);
    }    
    else{
      return "Error";
    }
  };
  xmlHttp.open( "GET", "http://130.240.170.56:8000/event/"+id, false );
  xmlHttp.send( null ); 
}

function getUser(id){
  var xmlHttp = null;
  xmlHttp = new XMLHttpRequest();
  xmlHttp.onreadystatechange=function() {
    if (xmlHttp.readyState==4 && xmlHttp.status==200) {
        var json = xmlHttp.responseText;
        var obj = JSON.parse(json);
        setCookie("username", id);  
    }else{
      return "ERROR";
    }
  };
  xmlHttp.open( "GET", "http://130.240.170.56:8000/users/"+id, false );
  xmlHttp.send( null );
}
 
function sendForm(form) {
  console.log(form);
  var data = {};
  for (var i = 0, ii = form.length; i < ii; ++i) {
    var input = form[i];
    if (input.name) {
      data[input.name] = input.value;
    }
  }
  
  var user = getCookie("username");
  console.log(user);
  data["User"] = user;

  var photo = document.getElementById('Photo');
  if(photo.files.length){
    var reader = new FileReader();
        function success(evt){
          data.Photo = evt.target.result; 

            console.log(data);
            var xhr = new XMLHttpRequest();
            xhr.onreadystatechange=function() {
              if (xhr.readyState==4 && xhr.status==200) {
                console.log("SUCCESSFULLY UPLOADED");
                return false;
              }
            }

    xhr.open('POST',"http://130.240.170.56:8000/event" , true);

    xhr.send(JSON.stringify(data));

  };

  reader.onload = success;
  reader.readAsDataURL(photo.files[0]);
  
  }
}

function redirect(data){
  addEventCookie(data); 
  window.location.replace("/show_event");
}

function homeredirect(){
  window.location.replace("/");
}

/*
Function not yet implemented
function getEventUser(userID){
  var xmlHttp = null;

  xmlHttp = new XMLHttpRequest();
  xmlHttp.onreadystatechange=function() {
    if (xmlHttp.readyState==4 && xmlHttp.status==200) {
        var json = xmlHttp.responseText;
        var obj = JSON.parse(json);
        console.log(obj);
        createUserLocations(obj);  
        
    }else{
      return "Error";
    }
  };
  xmlHttp.open( "GET", "http://130.240.170.56:8000/users/stairs/"+userID, false );   
  xmlHttp.send( null );
}
*/

/*
Function not yet implemented
function getUserEvent(user_id){
   var xmlHttp = null;

  xmlHttp = new XMLHttpRequest();
  xmlHttp.onreadystatechange=function() {
    if (xmlHttp.readyState==4 && xmlHttp.status==200) {
        var json = xmlHttp.responseText;
        var obj = JSON.parse(json);  
    }
    else{
      return "Error";
    }
  };
  
  xmlHttp.open( "GET", "http://130.240.170.56:8000/users/event/"+user_id, false );
  xmlHttp.send( null ); 
}
*/