/**
*  Module
*
* Description
*/

var app = angular.module('mainApp', []);

app.controller('main', ['$rootScope', function($rootScope){
    $rootScope.nombre = "ESTO FUNCIONA";
}]);

/*app.config(['$routeProvider', '$httpProvider', function($routeProvider) {
  console.log('esta ingresando aqui');
  $routeProvider.
      when('/', {'ngRoute','oc.lazyLoad'
        templateUrl: 'public/app/views/info.html',
        controller: 'indexController'
        /*resolve: {
            lazy: ['$ocLazyLoad', function($ocLazyLoad) {
                console.log('ingreso aqyu en la cagar de dependencias');
                return $ocLazyLoad.load([{
                    name: 'index',
                    files: ['app/controllers/index_controller.js']
                }]);
            }]
        }
      }).
      otherwise({
        redirectTo: '/'
      })
}]);*/