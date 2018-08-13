package models

import (
	_ "time"
	"github.com/astaxie/beego/orm"
)

type Captcha struct {
	Id                     	int       `orm:"column(id);auto"`
	UserId                  int       `orm:"column(user_id)"`
	Mobile        			string    `orm:"column(mobile);size(11)"`
	CaptchaKey        		string    `orm:"column(captcha_key);size(20)"`
	Captcha                 int       `orm:"column(captcha)"`
	VerfiCode        		int    	  `orm:"column(verfi_code)"`
	CreateDate        		int64       `orm:"column(create_date);size(11)"`
	ExpirationDate        	int64       `orm:"column(expiration_date);size(11)"`
	Ip                		int       `orm:"column(ip)"`
}

func (t *Captcha) TableName() string {
	return "scn_captcha"
}

func init() {
	orm.RegisterModel(new(Captcha))
}


func AddCaptcha(m *Captcha) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}




