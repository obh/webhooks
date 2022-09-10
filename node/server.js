var express = require('express')
var bodyParser = require('body-parser');
var crypto = require('crypto');
var app = express()

app.use(
  express.json({
    limit: '2mb',
    verify: (req, res, buf) => {
      req.rawBody = buf.toString();
    },
  })
);
app.use(bodyParser.json());

app.post('/webhook', function(req, res) {
  console.log(req.rawBody);
  const ts = req.headers["x-webhook-timestamp"]  
  const signature = req.headers["x-webhook-signature"]
  const currTs = Math.floor(new Date().getTime() / 1000)
  if(currTs - ts > 30000){
    res.send("Failed")
  }  
  genSign = verify(ts, req.rawBody)
  matched = genSign === signature
  console.log(genSign, signature, matched)
  res.send(matched)
})

function verify(ts, rawBody){
  const body = ts + rawBody
  let test = crypto.createHmac('sha256', "").update(body).digest("base64");
  return test
}

app.listen(8080, function () {
    console.log('listening on port 8080')
})
