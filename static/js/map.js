// Note: This example requires that you consent to location sharing when
// prompted by your browser. If you see the error "The Geolocation service
// failed.", it means you probably did not give permission for the browser to
// locate you.
var map;
function initMap() {
    map = new google.maps.Map(document.getElementById('map-canvas-0'), {
        center: {lat: 65.6177455, lng: 22.137957}, //The start coordinates is the classroom where we have our lectures!
        zoom: 16
    });
    var infoWindow = new google.maps.InfoWindow({map: map});
    var marker = new google.maps.Marker({
        map: map,
        animation: google.maps.Animation.DROP,
        title: 'You are here'
    });

    // Try HTML5 geolocation.
    if (navigator.geolocation) {
        navigator.geolocation.getCurrentPosition(function(position) {
            var pos = {
                lat: position.coords.latitude,
                lng: position.coords.longitude
            };

            marker.setPosition(pos);
            map.setCenter(pos);
        }, function() {
            handleLocationError(true, infoWindow, map.getCenter());
        });
    } else {
        // Browser doesn't support Geolocation
        handleLocationError(false, infoWindow, map.getCenter());
    }
}

function handleLocationError(browserHasGeolocation, infoWindow, pos) {
    infoWindow.setPosition(pos);
    infoWindow.setContent(browserHasGeolocation ?
        'Error: The Geolocation service failed.' :
        'Error: Your browser doesn\'t support geolocation.');
}