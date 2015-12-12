
function filterEvent(data){

	console.log(data);
	
	Event_name = document.getElementById('Event_name2');
	Address = document.getElementById('Address2');
	Zipcode = document.getElementById('Zipcode2');
	date = document.getElementById('Date2');
	Info = document.getElementById('Description2');
	Photo= document.getElementById('Photo2');
	//Preview = document.getElementById('Preview');

	if(document.getElementById("Event_name2") != null){
    	Event_name.innerHTML = Event_name.innerHTML + data.Event_name;
	}
	if(document.getElementById("Address2") != null){
    	Address.innerHTML = Address.innerHTML + data.Address;
	}
	if(document.getElementById("Zipcode2") != null){
    	Zipcode.innerHTML = Zipcode.innerHTML + data.Zipcode;
	}
	if(document.getElementById("Date2") != null){
    	date.innerHTML = date.innerHTML + data.Date;
	}
	if(document.getElementById("Description2") != null){
    	Info.innerHTML = Info.innerHTML + data.Info;
	}
	if(document.getElementById("Photo2") != null){
    	Photo.src = data.Photo;
	}

	//Preview.innerHTML = Preview.innerHTML + data.Preview;

}

function filterAllEvent(data){

	console.log(data);

	Event_name = document.getElementById('Event_name2');
	Address = document.getElementById('Address2');
	Zipcode = document.getElementById('Zipcode2');
	date = document.getElementById('Date2');
	Info = document.getElementById('Description2');
	Photo= document.getElementById('Photo2');
	//Preview = document.getElementById('Preview');

/*
	for( var i=0, l=data.length; i<l; i++ ) {
    	console.log( data[i] );
	    document.write("
	    	<nav class="Event-flow">
			    <div class="col-xs-10 col-sm-7" id="flow">
			        <a href="/show_event" id="" class="list-group-item">

			        	<h1 id="Event_name2"> </h1>
			    		<p id="Date2">Date: </p>
			    		<img src="" id="Photo2">
			      
			    	</a>       
			    </div>
			</nav>
");
*/

		if(document.getElementById("Event_name2") != null){
	    	Event_name.innerHTML = Event_name.innerHTML + data[i].Event_name;
		}
		if(document.getElementById("Address2") != null){
	    	Address.innerHTML = Address.innerHTML + data[i].Address;
		}
		if(document.getElementById("Zipcode2") != null){
	    	Zipcode.innerHTML = Zipcode.innerHTML + data[i].Zipcode;
		}
		if(document.getElementById("Date2") != null){
	    	date.innerHTML = date.innerHTML + data[i].Date;
		}
		if(document.getElementById("Description2") != null){
	    	Info.innerHTML = Info.innerHTML + data[i].Info;
		}
		if(document.getElementById("Photo2") != null){
	    	Photo.src = data[i].Photo;
		}

		//Preview.innerHTML = Preview.innerHTML + data.Preview;
	}
	
	/*
	Event_name = document.getElementById('Event_name2');
	Address = document.getElementById('Address2');
	Zipcode = document.getElementById('Zipcode2');
	date = document.getElementById('Date2');
	Info = document.getElementById('Description2');
	Photo= document.getElementById('Photo2');
	//Preview = document.getElementById('Preview');

	if(document.getElementById("Event_name2") != null){
    	Event_name.innerHTML = Event_name.innerHTML + data.Event_name;
	}
	if(document.getElementById("Address2") != null){
    	Address.innerHTML = Address.innerHTML + data.Address;
	}
	if(document.getElementById("Zipcode2") != null){
    	Zipcode.innerHTML = Zipcode.innerHTML + data.Zipcode;
	}
	if(document.getElementById("Date2") != null){
    	date.innerHTML = date.innerHTML + data.Date;
	}
	if(document.getElementById("Description2") != null){
    	Info.innerHTML = Info.innerHTML + data.Info;
	}
	if(document.getElementById("Photo2") != null){
    	Photo.src = data.Photo;
	}

	//Preview.innerHTML = Preview.innerHTML + data.Preview;

	*/

}
