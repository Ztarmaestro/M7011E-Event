function getLocations(recData, status, stairID){
  var xmlHttp = null;

  xmlHttp = new XMLHttpRequest();
  xmlHttp.onreadystatechange=function() {
    if (xmlHttp.readyState==4 && xmlHttp.status==200) {
        var json = xmlHttp.responseText;
        var obj = JSON.parse(json);
        console.log(obj);
        loadMap(obj);  
        
    }else{
      return "Error";
    }
  };
  xmlHttp.open( "GET", "http://130.240.170.56:8000/stairs", false );   
  xmlHttp.send( null );
}

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
  console.log(id); 
  xmlHttp = new XMLHttpRequest();
  xmlHttp.onreadystatechange=function() {
    if (xmlHttp.readyState==4 && xmlHttp.status==200) {
        var json = xmlHttp.responseText;
        var obj = JSON.parse(json);
        console.log(obj);  
        console.log(obj.IdToken);
        setCookie("username", obj.IdToken);  
    }else{
      return "ERROR";
    }
  };
  xmlHttp.open( "GET", "http://130.240.170.56:8000/users/"+id, false );
  xmlHttp.send( null );
}

function uploadPhoto(form){

  var data = {};
  for (var i = 0, ii = form.length; i < ii; ++i) {
    var input = form[i];
    if (input.name) {
      data[input.name] = input.value;
    }
  }
  data.userID = parseInt(data.userID);
  data.idStair = parseInt(data.idStair);
  var photo = document.getElementById('modalImage');
  if(photo.files.length){
    var reader = new FileReader();
        function success(evt){
          data.photo = evt.target.result; 
          send(data);
          //alert(evt.target.result);

        };
        reader.onload = success;
        
     reader.readAsDataURL(photo.files[0]);
                
         var xmlHttp = null;
         xmlHttp = new XMLHttpRequest();
         xmlHttp.onreadystatechange=function() {
          if (xmlHttp.readyState==4 && xmlHttp.status==200) { 
          console.log('UPLOAD SUCCESS');     
          }else{
            return "ERROR";
          }
        };
        xmlHttp.open( "POST", "http://130.240.170.56:8000/picture", true );
        xmlHttp.send(JSON.stringify(data)); 

    }
  }
 
function send(data){
          
          console.log(data);

          xmlHttp = new XMLHttpRequest();
         xmlHttp.onreadystatechange=function() {
          if (xmlHttp.readyState==4 && xmlHttp.status==200) {

            getPreviewStair(data.idStair);
            var res = JSON.parse(xmlHttp.responseText); 
            console.log(res);
            console.log('UPLOAD SUCCESS'); 

            var suc = document.getElementById('suc');
            suc.innerHTML =  "Upload successful!";
            setTimeout(function(){console.log('TIMOUT DONE');document.getElementById('suc').innerHTML ="Upload picture";},3000);   
            }else{
            return "ERROR";
          }
        };
        xmlHttp.open( "POST", "http://130.240.170.56:8000/picture", true );
        xmlHttp.send(JSON.stringify(data));
}


function getPreviewEvent(eventID){
  var xmlHttp = null;

  xmlHttp = new XMLHttpRequest();
   xmlHttp.onreadystatechange=function() {
    if (xmlHttp.readyState==4 && xmlHttp.status==200) {
        var json = xmlHttp.responseText;
        console.log(json);
	var obj = JSON.parse(json);
        console.log(obj);
	//createPhotos(obj);  
        
    }else{
      return "Error";
    }
  };
  xmlHttp.open( "GET", "http://130.240.170.56:8000/event/picture/preview/"+eventID, false );   
  xmlHttp.send( null );
}

function getPreviewUser(userID){
  var xmlHttp = null;

  xmlHttp = new XMLHttpRequest();
  xmlHttp.onreadystatechange=function() {
    if (xmlHttp.readyState==4 && xmlHttp.status==200) {
        var json = xmlHttp.responseText;
        var obj = JSON.parse(json);
        createPhotos(obj);  
        
    }else{
      return "Error";
    }
  };
  xmlHttp.open( "GET", "http://130.240.170.56:8000/users/picture/preview/"+userID, false );   
  xmlHttp.send( null );
}

function getOriginal(img){
  var xmlHttp = null;

  xmlHttp = new XMLHttpRequest();
  xmlHttp.onreadystatechange=function() {
    if (xmlHttp.readyState==4 && xmlHttp.status==200) {
        var json = xmlHttp.responseText;
        var obj = JSON.parse(json);
        showOriginal(obj); 
        
    }else{
      return "Error";
    }
  };
  xmlHttp.open( "GET", "http://130.240.170.56:8000/picture/"+parseInt(img.id), false );   
  xmlHttp.send( null );
}

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

            //data.user = parseInt(data.user);
            xhr.open('POST',"http://130.240.170.56:8000/event" , true);
            
            xhr.onreadystatechange=function() {
              if (xhr.readyState==4 && xhr.status==200) {
                console.log("SUCCESSFULLY UPLOADED");
                //document.getElementById('newLoc').reset();
                //markerPlaced = 0;
                //$('#main-map-canvas').empty;
                //initialize();
                return false;
              }
            }

    xhr.send(JSON.stringify(data));

  };

  reader.onload = success;
  reader.readAsDataURL(photo.files[0]);
  
  }
}

function redirect(data){

  console.log("nu körs jag "+data);
  setEventid_Cookie(data);
  var x = getCookie("username");
  console.log(x);
  setCookie("username")
  //window.location.replace("/show_event");

}

