
function filterEvent(data){

	console.log(data);
	printDiv = document.getElementById('Event_ID');
	headline = document.getElementById('Event_Name');

	printDiv.innerHTML = data.Event_ID;
	headline.innerHTML = data.Event_Name;
	
	console.log(data.Event_ID);
	console.log(data.Event_Name);
}