
function filterEvent(data){

	console.log(data);
	printDiv = document.getElementById('Event_ID');
	headline = document.getElementById('Event_Name');
	printDiv.innerHTML = data;
	headline.innerHTML = data.Event_Name;

}