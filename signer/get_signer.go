package signer

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
	"net/http"
	"time"
)

// SignGetRequest returns authorization headers for get request
func SignGetRequest(url, region, aKey, sKey string) (requestHeaders map[string][]string, err error) {
	creds := credentials.NewStaticCredentials(aKey, sKey, "")
	awsSess, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config: aws.Config{
			Credentials: creds,
		},
	})

	if err != nil {
		return nil, fmt.Errorf("failed creating session: %w", err)
	}

	v4Signer := v4.NewSigner(awsSess.Config.Credentials)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed create request: %w", err)
	}

	_, err = v4Signer.Sign(req, nil, "s3", region, time.Now())
	if err != nil {
		return nil, fmt.Errorf("error signing request: %w", err)
	}

	return req.Header, nil
}
