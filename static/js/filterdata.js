
function filterEvent(data){

	console.log(data);
	var printDiv = document.getElementById('Event_ID');
	var headline = document.getElementById('Event_Name');
	console.log(Event_ID);
	console.log(Event_Name);
	printDiv.innerHTML = data;
	headline.innerHTML = data.Event_Name;

}