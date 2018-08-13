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
	"strconv"
	"errors"
)
type Util map[string]interface{}
/**
	如果索引!=-1则表示包含该字符串。空字符串""在任何字符串中均存在
 */
const (
	regularMobile = "^(13[0-9]|14[57]|15[0-35-9]|18[07-9])\\d{8}$"
)
type IntIP struct {
	IP    string
	Intip int
}

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
func ValidateCaptcha(idkey string,verifyValue string,isClear bool) bool  {
	return base64Captcha.VerifyCaptchaAndIsClear(idkey, verifyValue, isClear)
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

/**
	将ip转化成int
 */
func IpToInt(ip string) int {
	var x *IntIP = &IntIP{IP: ip}
	var ipIntVal, _ = x.ToIntIp()
	return ipIntVal
}
func (self *IntIP) String() string {
	return self.IP
}

func (self *IntIP) ToIntIp() (int, error) {
	Intip, err := ConvertToIntIP(self.IP)
	if err != nil {
		return 0, err
	}
	self.Intip = Intip
	return Intip, nil
}

func (self *IntIP) ToString() (string, error) {
	i4 := self.Intip & 255
	i3 := self.Intip >> 8 & 255
	i2 := self.Intip >> 16 & 255
	i1 := self.Intip >> 24 & 255
	if i1 > 255 || i2 > 255 || i3 > 255 || i4 > 255 {
		return "", errors.New("Isn't a IntIP Type.")
	}
	ipstring := fmt.Sprintf("%d.%d.%d.%d", i4, i3, i2, i1)
	self.IP = ipstring
	return ipstring, nil
}

func ConvertToIntIP(ip string) (int, error) {
	ips := strings.Split(ip, ".")
	E := errors.New("Not A IP.")
	if len(ips) != 4 {
		return 0, E
	}
	var intIP int
	for k, v := range ips {
		i, err := strconv.Atoi(v)
		if err != nil || i > 255 {
			return 0, E
		}
		intIP = intIP | i<<uint(8*(3-k))
	}
	return intIP, nil
}
