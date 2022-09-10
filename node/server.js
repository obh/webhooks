var express = require('express')
var bodyParser = require('body-parser');
var crypto = require('crypto');
var app = express()

app.use(
  express.json({
    limit: '5mb',
    verify: (req, res, buf) => {
      req.rawBody = buf.toString();
    },
  })
);
app.use(bodyParser.json());

app.post('/', function (req, res) {
    console.log(JSON.stringify(req.headers));
    console.log(req.body)
    res.send('OK')
})

app.post('/webhook', function(req, res) {
  console.log(req.rawBody);
  const ts = req.headers["x-webhook-timestamp"]
  const signature = req.headers["x-webhook-signature"]
  console.log("expected sign --> ", signature);
  verify(ts, req.rawBody)
  res.send('OK')
})

function verify(ts, rawBody){
  const body = ts + rawBody
  let test = crypto.createHmac('sha256', "b85928b6c8f941c0ff8b8252ce040280305a3f3c").update("json").digest("base64");
  console.log(hash);
}

app.listen(8080, function () {
    console.log('listening on port 8080')
})
