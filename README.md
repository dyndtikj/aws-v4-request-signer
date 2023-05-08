# AWS V4 Signer

### This util can sign your GET request without body. 
### After run you will get signed request headers

## Example 1
```
import (
    "github.com/dyndtikj/aws-v4-request-signer/signer"
)
func main() {
    headers, err := signer.SignGetRequest(URL, S3Region, S3AccessKey, S3SecretKey)
    if err != nil {
    	fmt.Printf("faield to sign request: %s", err)
    }
       
    fmt.Println("Headers:")
    for k, v := range headers {
    	fmt.Println(k, ":", v)
    }
}	
```
### Output headers

```
Headers:
X-Amz-Content-Sha256 : [___some-hash]
Authorization : [AWS4-HMAC-SHA256 Credential=___some-hash/20230508/eu-west-2/s3/aws4_request, SignedHeaders=host;x-amz-content-sha256;x-amz-date, Signature=___some-hash]
X-Amz-Date : [20230508T182440Z]
```

## Example 2:
### Use main.go file, set vars there
```
URL = "https://your_s3_domain/example_file.txt"
S3Region = "eu-west-2"
S3AccessKey = "YOUR_ACCESS"
S3SecretKey = "YOUR_SECRET"
```
### go run main.go
```
Signed request headers:
X-Amz-Content-Sha256 : [___some-hash]
Authorization : [AWS4-HMAC-SHA256 Credential=___some-hash/20230508/eu-west-2/s3/aws4_request, SignedHeaders=host;x-amz-content-sha256;x-amz-date, Signature=___some-hash]
X-Amz-Date : [20230508T182440Z]
Response Body:
____bytes
```

