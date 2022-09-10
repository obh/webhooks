# Introduction
This project contains code to handle Cashfree Webhooks in the popular Echo framework. 
The steps to verify the webhook remain same

```
go mod download
go run main.go
```

1. Fetch Raw JSON and the headers
```go
signature := c.Request().Header.Get("x-webhook-signature")
	tsStr := c.Request().Header.Get("x-webhook-timestamp")
	timestamp, err := strconv.ParseInt(tsStr, 0, 64)
	if err != nil {
		fmt.Println("failed in getting proper timestamp")
		return c.String(http.StatusBadRequest, "Bad request")
	}
	slurp1, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}
	req := string(slurp1)
	fmt.Println("Raw request body:", req)
```

2. Compute signature and verify

```go
func VerifySignature(expectedSig string, ts int64, body string) (string, error) {
	t := time.Now()
	currentTS := t.Unix()
	if currentTS-ts > 1000*300 {
		return "", errors.New("webhook delivered too late")
	}
	signStr := strconv.FormatInt(ts, 10) + body
	fmt.Println("signing String: ", signStr)
	key := ""
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(signStr))
	b := h.Sum(nil)
	r
```