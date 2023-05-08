# AWS V4 Signer

### This util can sign your GET request without body. 
### After run you will get signed request headers

## Example:
### Signed request headers:

### X-Amz-Content-Sha256 : [___some-hash]

### Authorization : [AWS4-HMAC-SHA256 Credential=___some-hash/20230508/ru-msk/s3/aws4_request, SignedHeaders=host;x-amz-content-sha256;x-amz-date, Signature=___some-hash]

### X-Amz-Date : [20230508T182440Z]

### Response Body:

### ____bytes