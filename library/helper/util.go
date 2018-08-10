package helper

import (
	"strings"
	"crypto/md5"
	"io"
	"fmt"
	"regexp"
	"github.com/mojocn/base64Captcha"
	"time"
	"math/rand"
	"net/http"
	"io/ioutil"
)
type Util map[string]interface{}
/**
	如果索引!=-1则表示包含该字符串。空字符串""在任何字符串中均存在
 */
const (
	regularMobile = "^(13[0-9]|14[57]|15[0-35-9]|18[07-9])\\d{8}$"
)
/**
	密码加密
 */
func PwdEncrypt(pwd string,salt string) string {
	if strings.Index("$.", pwd) != -1 {
		return pwd
	} else {
		w := md5.New()
		io.WriteString(w, pwd)   //将pass写入到w中
		md5str1 := fmt.Sprintf("%x", w.Sum(nil)) + salt  //w.Sum(nil)将w的hash转成[]byte格式
		w1 := md5.New()
		io.WriteString(w1, md5str1)
		md5str2 := "$." + fmt.Sprintf("%x", w1.Sum(nil))
		return md5str2
	}
}
/**
	手机号正则
 */
func ValidateMobile(mobileNum string) bool {
	reg := regexp.MustCompile(regularMobile)
	return reg.MatchString(mobileNum)
}
/***
	验证图形验证码
 */
func ValidateCaptcha(idkey string,verifyValue string) bool  {
	return base64Captcha.VerifyCaptcha(idkey, verifyValue)
}

/**
	生成六位随机验证码
 */
func RandCaptchaCode() string{
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	return vcode
}

/**
	模拟http get 请求
 */
func HttpGet(geturl string) string{
	client := &http.Client{}
	//向服务端发送get请求
	request, _ := http.NewRequest("GET", geturl, nil)
	//接收服务端返回给客户端的信息
	response, _ := client.Do(request)
	if response.StatusCode == 200 {
	 	str, _ := ioutil.ReadAll(response.Body)
		return string(str)
	} else {
		return ""
	}
}