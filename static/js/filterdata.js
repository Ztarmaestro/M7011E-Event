
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

	for( var i=0, l=data.length; i<l; i++ ) {

		var a = document.createElement("a");
    	a.className = list-group-item;
    	a.id = linkevent;
		a.href = "/show_event";
    	document.getElementById("flow").appendChild(a);

    	var head = document.createElement("h1");
    	head.id = Event_name2;
    	document.getElementById("linkevent").appendChild(head);

    	var para = document.createElement("p");
    	para.id = Date2;
    	document.getElementById("linkevent").appendChild(para);

    	var pic = document.createElement("img");
    	pic.id = Photo2;
    	pic.src = "";
    	document.getElementById("linkevent").appendChild(pic);

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
	
}
