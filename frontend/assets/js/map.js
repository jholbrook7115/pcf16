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

app.controller("MapCtrl", function($scope, $mdSidenav){
	$scope.searchText = "";

	


 });

function initMap() {
	var map;
    map = new google.maps.Map(document.getElementById('map'), {

        center: {lat: -34.397, lng: 150.644},
        zoom: 8
    });
}