var view = require('./view.js');

function initModel() {
}

const getIndex = function(req, res) {
  view.index(res);
}

module.exports = {
  initModel,
  getIndex
};
