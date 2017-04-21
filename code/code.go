package code

import (
	"time"
	"bytes"
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/dchest/captcha"
	"ckcaptcha/global"
)

type CaptchaInfo struct {
	CodeLen int
	Expiration time.Duration
	Width int
	Height int
}

func NewCpatchaInfo(len int, expire time.Duration, width int, height int) *CaptchaInfo{
	cp := &CaptchaInfo{
		CodeLen: len,
		Expiration: expire,
		Width: width,
		Height: height,
	}
	return cp
}

func (cp *CaptchaInfo) NewLen(length int) (id string) {
	id = cp.RandomId()
	res := captcha.RandomDigits(length)
	expire := time.Duration(900) * time.Second
	err := global.RedisClient.Set(id, res, expire).Err()
	if err != nil {
		log.Errorf("New Captcha Error: %s", err.Error())
	}
	return
}

func (cp *CaptchaInfo) GetCaptchaId() string {
	id := cp.NewLen(cp.CodeLen)
	return id
}

func (cp *CaptchaInfo) RandomId() string {
	idLen := 20
	var idChars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	b := captcha.RandomDigits(idLen)
	for i, c := range b {
		b[i] = idChars[c]
	}
	return string(b)
}

func (cp *CaptchaInfo) ReloadCaptcha(id string) bool {
	old, _ := global.RedisClient.Get(id).Bytes()
	if old == nil {
		return false
	}

	res := captcha.RandomDigits(len(old))
	expire := time.Duration(900) * time.Second
	err := global.RedisClient.Set(id, res, expire).Err()
	if err != nil {
		log.Errorf("Reload Captcha Error: %s", err.Error())
		return false
	}
	return true
}

func (cp *CaptchaInfo) GetCaptchaImage(buff bytes.Buffer, id string, w int, h int) ([]byte, error) {
	var err error
	var d []byte
	d, _ = global.RedisClient.Get(id).Bytes()
	if d == nil {
		return nil, fmt.Errorf("captcha id not match %s", id)
	}

	if w == 0 || h == 0 {
		_, err = captcha.NewImage(id, d, cp.Width, cp.Height).WriteTo(&buff)
	} else {
		_, err = captcha.NewImage(id, d, w, h).WriteTo(&buff)
	}
	if err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}

func (cp *CaptchaInfo) VerifyCaptcha(id string, digits string) bool {
	if digits == "" {
		return false
	}
	ns := make([]byte, len(digits))
	for i := range ns {
		d := digits[i]
		switch {
		case '0' <= d && d <= '9':
			ns[i] = d - '0'
		case d == ' ' || d == ',':
			// ignore
		default:
			return false
		}
	}
	return cp.Verify(id, ns)
}

func (cp *CaptchaInfo) Verify(id string, digits []byte) bool {
	if digits == nil || len(digits) == 0 {
		return false
	}
	reaId, _ := global.RedisClient.Get(id).Bytes()
	if reaId == nil {
		return false
	}
	return bytes.Equal(digits, reaId)
}