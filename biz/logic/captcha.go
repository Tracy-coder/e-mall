package logic

import (
	"github.com/Tracy-coder/e-mall/biz/domain"
	"github.com/Tracy-coder/e-mall/configs"
	"github.com/Tracy-coder/e-mall/data"
	"github.com/Tracy-coder/e-mall/pkg/captcha"
	"github.com/mojocn/base64Captcha"
)

var (
	captchaDriver *base64Captcha.DriverDigit
	CaptchaStore  *captcha.CacheStore
)

func init() {
	c := configs.Data()
	captchaDriver = base64Captcha.NewDriverDigit(c.Captcha.ImgHeight, c.Captcha.ImgWidth,
		c.Captcha.KeyLong, 0.7, 80)
	CaptchaStore = captcha.NewCacheStore(data.CacheDB("captcha"))
}

type Captcha struct {
	CaptchaDriver *base64Captcha.DriverDigit
	CaptchaStore  *captcha.CacheStore
}

func NewCaptcha() domain.Captcha {
	// if captchaDriver == nil || CaptchaStore == nil {
	// 	initCaptcha()
	// }
	return &Captcha{
		CaptchaDriver: captchaDriver,
		CaptchaStore:  CaptchaStore,
	}
}

func (c *Captcha) GetCaptcha() (id, b64s string, err error) {
	captchaGen := base64Captcha.NewCaptcha(c.CaptchaDriver, c.CaptchaStore)
	id, b64s, err = captchaGen.Generate()
	return
}
