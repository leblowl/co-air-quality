var express = require('express');
var control = require('./control.js');

function loadRoutes(app) {
  app.get('/', control.getIndex);
}

module.exports = {
  loadRoutes
};
