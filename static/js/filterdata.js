
function filterEvent(data){

	console.log(data);
	
	event_name = document.getElementById('Event_name');
	address = document.getElementById('Address');
	zipcode = document.getElementById('Zipcode');
	date = document.getElementById('Date');
	Info = document.getElementById('Info');
	user = document.getElementById('User');
	event_id = document.getElementById('Event_ID');
	preview = document.getElementById('Preview');
	

	event_name.innerHTML = data.Event_name;
	address.innerHTML = data.address;
	zipcode.innerHTML = data.zipcode;
	date.innerHTML = data.date;
	Info.innerHTML = data.Info;
	user.innerHTML = data.user;
	event_id.innerHTML = data.Event_ID;
	preview.innerHTML = data.preview

	console.log(data.Event_ID);
	console.log(data.Event_name);

}