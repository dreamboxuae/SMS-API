var request = require('request');

var dataString = 'From=19876543212&To=13216549878&Body=Test SMS from Restcomm&StatusCallback=http://status.callback.url';

var options = {
    url: 'https://<accountSid>:<authToken>@cloud.restcomm.com/restcomm/2012-04-24/Accounts/<accountSid>/SMS/Messages',
    method: 'POST',
    body: dataString
};

function callback(error, response, body) {
    if (!error && response.statusCode == 200) {
        console.log(body);
    }
}

request(options, callback);
