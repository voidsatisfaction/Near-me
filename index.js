var express = require('express');

var bodyParser = require('body-parser');
var morgan = require('morgan');

var moment = require('moment');

var rp = require('request-promise');

var app = express();

var PORT = process.env.PORT || 3003;

app.use(morgan('dev'));
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: false }));
app.use(express.static(__dirname + '/build'));

app.get('/', function(req, res) {
  res.send('worked');
});

app.get('/myplace', function(req, res) {
  var lat = req.query.lat;
  var lon = req.query.lng;

  getUserLocation({ lat, lon })
    .then(getEvents)
    .then((data) => {
      res.status(200).json(data);
    })
    .catch((error) => {
      console.error(error);
      res.status(500);
    })
});

function getUserLocation({ lat, lon }) {
  var yahooMapHost = 'https://map.yahooapis.jp/placeinfo/V1/get';
  var yahooAppId = 'dj00aiZpPUFuRWxDTE5rZFBBayZzPWNvbnN1bWVyc2VjcmV0Jng9MmE-';

  var uri = `${yahooMapHost}?lat=${lat}&lon=${lon}&appid=${yahooAppId}`;
  var options = {
    method: 'GET',
    uri,
    json: true
  }

  return rp(options)
    .then((data) => {
      return data.ResultSet.Address[0];
    })
}

function getEvents(place) {
  var connpassHost = 'https://connpass.com/api/v1/event/';

  var place = encodeURIComponent(place);
  var uri = `${connpassHost}?keyword=${place}&count=100`;
  var options = {
    method: 'GET',
    uri,
    json: true,
    headers: {
      'User-Agent': 'Request-Promise'
    }
  };
  return rp(options)
    .then((data) => {
      var now = moment();
      var filteredEvents = data.events.filter((event) => now.isSameOrBefore(event.started_at));
      data.events = filteredEvents;
      return data;
    })
}

app.listen(PORT, function() {
  console.log('Server listening on ' + PORT);
});