package domain

type Captcha interface {
	GetCaptcha() (id, b64s string, err error)
}
