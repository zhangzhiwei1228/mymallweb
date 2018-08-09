package controllers

import (
	"github.com/mojocn/base64Captcha"
	"github.com/astaxie/beego"
	"encoding/json"
)

// CaptchaController operations for Captcha
type CaptchaController struct {
	BaseController
}

func (c *CaptchaController) CreateCaptcha() {
	//config struct for digits
	//数字验证码配置
	/*var configD = base64Captcha.ConfigDigit{
		Height:     80,
		Width:      240,
		MaxSkew:    0.7,
		DotCount:   80,
		CaptchaLen: 5,
	}*/

	//config struct for audio
	//声音验证码配置
	/*var configA = base64Captcha.ConfigAudio{
		CaptchaLen: 6,
		Language:   "zh",
	}*/
	//config struct for Character
	//字符,公式,验证码配置
	var configC = base64Captcha.ConfigCharacter{
		Height:             60,
		Width:              240,
		//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
		Mode:               base64Captcha.CaptchaModeArithmetic,
		ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
		ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
		IsShowHollowLine:   false,
		IsShowNoiseDot:     false,
		IsShowNoiseText:    false,
		IsShowSlimeLine:    false,
		IsShowSineLine:     false,
		CaptchaLen:         6,
	}
	//创建声音验证码
	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	//idKeyA, capA := base64Captcha.GenerateCaptcha("", configA)
	//以base64编码
	//base64stringA := base64Captcha.CaptchaWriteToBase64Encoding(capA)


	//创建字符公式验证码.
	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	idKeyC, capC := base64Captcha.GenerateCaptcha("", configC)
	//以base64编码
	base64stringC := base64Captcha.CaptchaWriteToBase64Encoding(capC)


	//创建数字验证码.
	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	//idKeyD, capD := base64Captcha.GenerateCaptcha("", configD)
	//以base64编码
	//base64stringD := base64Captcha.CaptchaWriteToBase64Encoding(capD)


	body := map[string]interface{}{"img": base64stringC, "captchaId": idKeyC}
	c.SetJson(0,body,"")
	beego.Info(idKeyC)
	return
}

func (c *CaptchaController) CheckCaptcha()  {

}

func (c *CaptchaController) VerfiyCaptcha(){
	beego.Info(string(c.Ctx.Input.RequestBody))
	postData := map[string]string{"val": "", "idkey": ""}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &postData)
	if err != nil {
		c.SetJson(1, nil, "数据格式错误")
		return
	}
	idkey := postData["idkey"]
	verifyValue := postData["val"]
	verifyResult := base64Captcha.VerifyCaptcha(idkey, verifyValue)
	beego.Info(verifyResult)
	c.SetJson(0,verifyResult,"")
	return
}
