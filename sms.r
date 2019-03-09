require(httr)

data = list(
  `From` = '19876543212',
  `To` = '13216549878',
  `Body` = 'Test SMS from Restcomm',
  `StatusCallback` = 'http://status.callback.url'
)

res <- httr::POST(url = 'https://%3CaccountSid%3E:%3CauthToken%3E@cloud.restcomm.com/restcomm/2012-04-24/Accounts/%3CaccountSid%3E/SMS/Messages', body = data)
