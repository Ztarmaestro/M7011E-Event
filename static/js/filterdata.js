
function filterEvent(data){

	console.log(data);
	
	Event_name = document.getElementById('Event_name2');
	Address = document.getElementById('Address2');
	Zipcode = document.getElementById('Zipcode2');
	date = document.getElementById('Date2');
	Info = document.getElementById('Description2');
	//Preview = document.getElementById('Preview');
	Photo= document.getElementById('Photo2');

	Event_name.innerHTML = Event_name.innerHTML + data.Event_name;
	date.innerHTML = date.innerHTML + data.Date;
	Photo.src = data.Photo;
	Zipcode.innerHTML = Zipcode.innerHTML + data.Zipcode;
	Address.innerHTML = Address.innerHTML + data.Address;
	Info.innerHTML = Info.innerHTML + data.Info;
	//Preview.innerHTML = Preview.innerHTML + data.Preview;


}