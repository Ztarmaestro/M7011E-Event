
function filterEvent(data){

	console.log(data);
	
	Event_name = document.getElementById('Event_name');
	Address = document.getElementById('Address');
	Zipcode = document.getElementById('Zipcode');
	date = document.getElementById('Date');
	Info = document.getElementById('Description');
	//Preview = document.getElementById('Preview');
	Photo= document.getElementById('Photo');

	Event_name.innerHTML = Event_name.innerHTML + data.Event_name;
	date.innerHTML = date.innerHTML + data.Date;
	Photo.src = data.Photo;
	Zipcode.innerHTML = Zipcode.innerHTML + data.Zipcode;
	Address.innerHTML = Address.innerHTML + data.Address;
	Info.innerHTML = Info.innerHTML + data.Info;
	//Preview.innerHTML = Preview.innerHTML + data.Preview;


}