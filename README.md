# Vegeta upload multipart/form-data

To reduce the complexity of the testing process using vegeta and also the scarcity of good examples from the https://github.com/tsenart/vegeta project, I created here a detailed example for sending files via multipart/form-data.

When submitting a multipart/form-data request in the API server benchmark, the vegeta worker could not be sent without you coding in golang, but Vegeta could easily send the request without writing code.

By specifying the text describing the request header and request body with the <code>-targets</code> option of the vegeta attack, you can flexibly send HTTP requests.

The following is an example of the <code>targets.tgt</code> file specified for <code>-targets</code>.

```
POST http://localhost:8080/api/path
Content-Type: multipart/form-data; boundary=vegetaboundary
@body.txt
```
<code>@body.txt</code> is the path to the file that recorded the request body and is written as follows.
```
--vegetaboundary
Content-Disposition: form-data; name="file"; filename="dummy"
Content-Type: application/octet-stream

12345-body-of-binary
--vegetaboundary--
```
When the files are ready, perform the vegeta attack.
```
./vegeta attack -rate=1/1s -duration=1m -targets=targets.tgt -timeout=60s | ./vegeta report
```
You can submit by writing the request body and the request header in text like this.
The result of the execution will look something like this:
```
Requests      [total, rate, throughput]  60, 1.02, 0.99
Duration      [total, attack, wait]      1m0.569015152s, 58.999960073s, 1.569055079s
Latencies     [mean, 50, 95, 99, max]    1.921592061s, 1.369102872s, 4.41999714s, 5.395822318s, 5.465820907s
Bytes In      [total, mean]              0, 0.00
Bytes Out     [total, mean]              307208280, 5120138.00
Success       [ratio]                    100.00%
Status Codes  [code:count]               200:60
```
