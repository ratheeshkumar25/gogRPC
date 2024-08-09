package config

import (
	"errors"
	"os"

	"github.com/twilio/twilio-go"
	verify "github.com/twilio/twilio-go/rest/verify/v2"
)

// TwilioService represents teh twilio client and config files.
type TwilioService struct {
	Client *twilio.RestClient
}

// SetupTwilio will initialise the connection to twilio client.
func SetupTwilio() *TwilioService {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: os.Getenv("TWILIO_ACCOUNT_SID"),
		Password: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	return &TwilioService{
		Client: client,
	}
}

// SendTwilioOTP manage to sent otp to a desired phone number.
func (t *TwilioService) SendTwilioOTP(phone string) (*verify.VerifyV2Verification, error) {
	params := &verify.CreateVerificationParams{}
	params.SetTo("+919353306805")
	params.SetChannel("sms")

	resp, err := t.Client.VerifyV2.CreateVerification(os.Getenv("SERVICE_TOKEN"), params)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// VerifyTwilioOTP will check the provided otp with the phone number as key.
func (t *TwilioService) VerifyTwilioOTP(phone, otp string) error {
	params := verify.CreateVerificationCheckParams{}
	params.SetTo("+919353306805")
	params.SetCode(otp)

	resp, err := t.Client.VerifyV2.CreateVerificationCheck(os.Getenv("SERVICE_TOKEN"), &params)
	if err != nil {
		return err
	} else if *resp.Status == "approved" {
		return nil
	}
	return errors.New("incorrect code")
}
