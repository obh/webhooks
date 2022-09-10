package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.POST("/webhook", webhook)
	fmt.Println("Start server ....")
	e.Logger.Fatal(e.Start(":8080"))

}

func webhook(c echo.Context) error {
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
	genSignature, err := VerifySignature(signature, timestamp, req)
	if err != nil {
		return c.String(http.StatusOK, "Failure in verifying signature")
	}
	fmt.Println("generated signature: ", genSignature)
	fmt.Println("expected signature: ", signature)
	matched := signature == genSignature
	fmt.Println("match? ", matched)
	return c.String(http.StatusOK, "Request completed")
}

func VerifySignature(expectedSig string, ts int64, body string) (string, error) {
	t := time.Now()
	currentTS := t.Unix()
	fmt.Println(currentTS, ts, currentTS-ts)
	//if(currentTS - ts > 1000 * 300){
	//  return false
	//}
	signStr := strconv.FormatInt(ts, 10) + body
	fmt.Println("signing String: ", signStr)
	key := "b85928b6c8f941c0ff8b8252ce040280305a3f3c"
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(signStr))
	b := h.Sum(nil)
	return base64.StdEncoding.EncodeToString(b), nil
}
