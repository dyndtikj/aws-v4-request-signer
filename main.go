package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
	"io"
	"log"
	"net/http"
	"time"
)

var (
	// URL url for signing
	URL = "https://your_s3_domain/example_file.txt"
	// S3Region of yours
	S3Region = "ru-msk"
	// S3AccessKey S3Secret key - are creds, that are used for signing
	S3AccessKey = "YOUR_ACCESS"
	S3SecretKey = "YOUR_SECRET"
)

func main() {
	creds := credentials.NewStaticCredentials(S3AccessKey, S3SecretKey, "")
	awsSess, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config: aws.Config{
			Credentials: creds,
		},
	})

	if err != nil {
		log.Fatalf("failed creating session: %s", err)
	}
	signer := v4.NewSigner(awsSess.Config.Credentials)

	if err = DebugSign(signer); err != nil {
		log.Fatal("failed sign:", err)
	}
}

func DebugSign(signer *v4.Signer) error {
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return fmt.Errorf("failed create request, %w", err)
	}

	_, err = signer.Sign(req, nil, "s3", S3Region, time.Now())
	if err != nil {
		return fmt.Errorf("error signing request: %w", err)
	}

	fmt.Println("Signed request headers:")
	for k, v := range req.Header {
		fmt.Println(k, ":", v)
	}

	httpClient := http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed do request: %w", err)
	}

	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			_ = fmt.Errorf("failed close response body, %w", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed read response body: %w", err)
	}

	fmt.Printf("Response Body:\n %s", string(body))
	return nil
}
