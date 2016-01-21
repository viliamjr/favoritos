"use strict";

// necessário para alterar os delimitadores do template do angular
// para evitar conflito com os templates do Go
var myApp = angular.module('minhaApp', [], function($interpolateProvider) {
  $interpolateProvider.startSymbol('[[');
  $interpolateProvider.endSymbol(']]');
});

myApp.controller('MeuController', function($scope, $http) {

  $http.get('/dados')
    .success(function(data) {
      $scope.msg = data.msg;
    })
    .error(function(data, status) {
      $scope.msg = "Ops, erro [" + status + "]: " + data;
    });
});
