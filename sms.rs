# Sample - Send SMS via Restcomm API in RuST programming langugage.

extern crate reqwest;

fn main() -> Result<(), reqwest::Error> {

    let res = reqwest::Client::new()
        .post("https://<accountSid>:<authToken>@cloud.restcomm.com/restcomm/2012-04-24/Accounts/<accountSid>/SMS/Messages")
        .body("From=19876543212&To=13216549878&Body=TestSMSfromRestcomm&StatusCallback=http://status.callback.url")
        .send()?
        .text()?;
    println!("{}", res);

    Ok(())
}
