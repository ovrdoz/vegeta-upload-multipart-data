package main
import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
	vegeta "github.com/tsenart/vegeta/lib"
)

func main() {
	// set header
	header := http.Header{}
	header.Set("apiKey","<your-api-key>")
	header.Set("Content-Type", "multipart/form-data; boundary=vegetaboundary")
	// get image
	file, err := os.Open("body.txt")
  	if err != nil {
    		panic(err)
  	}
  	defer file.Close()
  	body := &bytes.Buffer{}
  	writer := multipart.NewWriter(body)
  	writer.SetBoundary("vegetaboundary")
  	writer.WriteField("payload", "content-of-payload")
  	part, err := writer.CreateFormFile("file", filepath.Base(file.Name()))
  	if err != nil {
    		panic(err)
  	}
  	io.Copy(part, file)
  	writer.Close()

  	rate := vegeta.Rate{Freq: 2, Per: time.Second}
  	duration := 3 * time.Second
  	targeter := vegeta.NewStaticTargeter(vegeta.Target{
    		Method: "POST",
    		URL:    "http://localhost:8080/api/path",
    		Header: header,
    		Body:   body.Bytes(),
  	})
  	attacker := vegeta.NewAttacker()
  	var metrics vegeta.Metrics
  	for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
    		metrics.Add(res)
  	}
  	metrics.Close()
  	fmt.Printf("%v", metrics.StatusCodes) // get status
}
