var app = angular.module("MapApp", ['ngMaterial']);

	// app.directive("GMaps", function(){
	// 	return {
	// 		resrict: 'E',
	// 		template: '<div></div>',
	// 		repalce: true,
	// 		link: function(scope, element, attrs){
	// 			var mapoptions = {
	// 				center: new google.maps.LatLng
	// 			};
	// 		}
	// 	};
	// });

app.controller("MapCtrl", function($scope, $mdSidenav, $log, $mdSidenav, $http, $window){

	$window.map;
	$window.initMap = function initMap() {
		// var map;
	    $scope.map = new google.maps.Map(document.getElementById('map'), {

	        center: {lat: 40.7127, lng: 74.0059},
	        zoom: 8
	    });

	    var geocoder = new google.maps.Geocoder;


	    if(navigator.geolocation){
	    	navigator.geolocation.getCurrentPosition(function(position){
					initialLocation = new google.maps.LatLng(position.coords.latitude, position.coords.longitude);
					$scope.map.setCenter(initialLocation);

					geocoder.geocode({'location': $scope.map.getCenter()}, function(results, status){
						if(status === google.maps.GeocoderStatus.OK){
							if (results[1]){
								$scope.map.setZoom(11);
								var marker = new google.maps.Marker({
									position: $scope.map.getCenter(),
									map: $scope.map
								});

								$log.debug(results[1].formatted_address);
								$scope.searchText=results[1].formatted_address;
							}

						}
					})
				}, function(){
					//alert that the location service is not available
				});

		} else {
			alert("geolocation services not available");

		}
	}	
	$scope.searchText = "";
	// $scope.geoLocationAvailable = navigator.geolocation;
	// if($scope.geoLocationAvailable){
	// 	$log.debug('Geolocation is available');
	// } else {
	// 	$log.debug('Geolocation NOT available');
	// }

	$scope.repositionMap = function(position){

		if(navigator.geolocation){
			var startPos;
			// 
		}
	};

	$scope.listResults = '';
	$scope.listStatus;
	

	$scope.search = function(useCurrentLocation){
		var geocoder;
		var initialLocation;
		$scope.url;
		// map = document.getElementById("map");
		var PP_url = "https://www.parkingpanda.com/api/v2/locations?search=";
		
		// var geoSuccess = function(position){
		// 	$log.debug(position);
		// 	//geocoder.geocode( {'location': position}, )
		// 	url = PP_url.concat(position);
		// };
		if(useCurrentLocation){

			url = PP_url.concat();
		} else {
			url = PP_url.concat($scope.searchText);
		}
		//url = "https://localhost:8082/api/lots";
		url = "https://pcf-2016.appspot.com/api/lots";
		$log.debug(url);
		$http({
			method: 'GET',
			url: url
		}).then(function(response){
			$log.debug(response.status);
			$log.debug(response.data);

			$scope.listResults = response.data;
			$scope.openRightSidenav();

			//add markers to the map

			for(var i = 0; i < response.data.length; i++){
				var location = response.data[i];
				var marker = new google.maps.Marker({
					position: {lat: location.lat, lng: location.lng},
					map: $scope.map
					
				});
			}

		}, function(response){
			$log.debug(response.data || "Request Failed");
			$log.debug(response.status);
		});

	};

	$scope.openRightSidenav = function(){
		if($mdSidenav('right-sidenav').isOpen()){
			$mdSidenav('right-sidenav').close();
		} else {
			$mdSidenav('right-sidenav').open();
		}
	}
 });

// function initMap() {
// 	var map;
//     map = new google.maps.Map(document.getElementById('map'), {

//         center: {lat: -34.397, lng: 150.644},
//         zoom: 8
//     });

//     if(navigator.geolocation){
//     	navigator.geolocation.getCurrentPosition(function(position){
// 				initialLocation = new google.maps.LatLng(position.coords.latitude, position.coords.longitude);
// 				map.setCenter(initialLocation);

// 			}, function(){
// 				//alert that the location service is not available
// 			});

// 	} else {
// 		alert("geolocation services not available");

// 	}
// }