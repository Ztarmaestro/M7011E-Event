
function filterEvent(data){

	console.log(data);
	
	Event_name = document.getElementById('Event_name');
	Address = document.getElementById('Address');
	Zipcode = document.getElementById('Zipcode');
	date = document.getElementById('Date');
	Info = document.getElementById('Info');
	Preview = document.getElementById('Preview');

	console.log(data.Event_name);
	Event_name.innerHTML = Event_name.innerHTML + data.Event_name;
	console.log(data.Address);
	Address.innerHTML = Address.innerHTML + data.Address;
	console.log(data.Zipcode);
	Zipcode.innerHTML = Zipcode.innerHTML + data.Zipcode;
	console.log(data.Date);
	date.innerHTML = date.innerHTML + data.Date;
	console.log(data.Info);
	Info.innerHTML = Info.innerHTML + data.Info;
	console.log(data.Preview);
	Preview.innerHTML = Preview.innerHTML + data.Preview;

}