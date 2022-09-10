# Introduction
This project contains code to handle Cashfree Webhooks in the popular Express framework. 
The steps to verify the webhook remain same. 

```bash
npm install
node server.js
```

1. Fetch Raw JSON and the headers
```javascript
app.use(
  express.json({
    limit: '2mb',
    verify: (req, res, buf) => {
      req.rawBody = buf.toString();
    },
  })
);
```

2. Compute signature and verify
```javascript
function verify(ts, rawBody){
  const body = ts + rawBody
  let test = crypto.createHmac('sha256', "").update(body).digest("base64");
  return test
}
```