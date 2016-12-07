var index = angular.module('index', []);

index.controller('indexController', ['$scope','$rootScope', function($scope, $rootScope){
	/*$scope.title = "esto es una prueba del funcionamiento";
	$rootScope.isVisibleMenu = false;
	$rootScope.isVisibleMenuDashboard = true;
	$rootScope.isVisibleMenuBack = true;*/
	alert('esta ingresando aqui indexController');
}]);