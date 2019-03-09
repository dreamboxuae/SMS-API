<?php

$ch = curl_init();

curl_setopt($ch, CURLOPT_URL, 'https://<accountSid>:<authToken>@cloud.restcomm.com/restcomm/2012-04-24/Accounts/<accountSid>/SMS/Messages');
curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
curl_setopt($ch, CURLOPT_POSTFIELDS, "From=19876543212&To=13216549878&Body=Test SMS from Restcomm&StatusCallback=http://status.callback.url");
curl_setopt($ch, CURLOPT_POST, 1);

$headers = array();
$headers[] = 'Content-Type: application/x-www-form-urlencoded';
curl_setopt($ch, CURLOPT_HTTPHEADER, $headers);

$result = curl_exec($ch);
if (curl_errno($ch)) {
    echo 'Error:' . curl_error($ch);
}
curl_close ($ch);
