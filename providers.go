package gocaptcha

import "context"

type IProvider interface {
	// SolveImageCaptcha is the implementation of getting the response of an image captcha
	SolveImageCaptcha(ctx context.Context, settings *Settings, payload *ImageCaptchaPayload) (ICaptchaResponse, error)

	// SolveRecaptchaV2 is the implementation of getting the response of a version 2 recaptcha
	SolveRecaptchaV2(ctx context.Context, settings *Settings, payload *RecaptchaV2Payload) (ICaptchaResponse, error)

	// SolveRecaptchaV3 is the implementation of getting the response of a version 3 recaptcha
	SolveRecaptchaV3(ctx context.Context, settings *Settings, payload *RecaptchaV3Payload) (ICaptchaResponse, error)

	// SolveHCaptcha is the implementation of getting the response of an HCaptcha captcha
	SolveHCaptcha(ctx context.Context, settings *Settings, payload *HCaptchaPayload) (ICaptchaResponse, error)
}

type ICaptchaResponse interface {
	// Solution will return the solution of the captcha as a string
	Solution() string

	// ReportBad reports the captcha to be invalid if the provider and captcha type support it.
	ReportBad(ctx context.Context) error

	// ReportGood reports the captcha to be valid if the provider and captcha type support it.
	ReportGood(ctx context.Context) error
}