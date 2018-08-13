package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	_ "time"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

type User struct {
	Id                     	int       `orm:"column(id);auto"`
	GradeId                 int       `orm:"column(grade_id)"`
	GroupId                 int       `orm:"column(group_id)"`
	ParentId                int       `orm:"column(parent_id)"`
	AreaId                	int       `orm:"column(area_id)"`
	Role                   	string    `orm:"column(role)"`
	AreaText                string    `orm:"column(area_text)"`
	AgentAid                int    	  `orm:"column(agent_aid)"`
	AgentAtext              string    `orm:"column(agent_atext);size(100)"`
	Destination             string    `orm:"column(destination);size(255)"`
	Username               	string    `orm:"column(username);size(50)"`
	Password           		string    `orm:"column(password);size(50)"`
	Salt           			string    `orm:"column(salt);size(8)"`
	PayPass           		string    `orm:"column(pay_pass);size(50)"`
	PaySalt           		string    `orm:"column(pay_salt);size(50)"`
	Nickname           		string    `orm:"column(nickname);size(50)"`
	Realname           		string    `orm:"column(realname);size(50)"`
	Mobile        			string    `orm:"column(mobile);size(100)"`
	Idcard        			string    `orm:"column(idcard);size(100)"`
	Question        		string    `orm:"column(question);size(255)"`
	Answer        			string    `orm:"column(answer);size(255)"`
	RefUrl        			string    `orm:"column(ref_url);size(255)"`
	Avatar        			string    `orm:"column(avatar);size(100)"`
	Sign        			string    `orm:"column(sign);size(100)"`
	LoginNum        		int       `orm:"column(login_num);size(11)"`
	IsDeleted        		int       `orm:"column(is_deleted);size(1)"`
	IsEnabled        		int       `orm:"column(is_enabled);size(1)"`
	IsVip        			int       `orm:"column(is_vip);size(1)"`
	IsAdmin        			int       `orm:"column(is_admin);size(1)"`
	Exp        				int       `orm:"column(exp);size(11)"`
	ResaleGrade        		int       `orm:"column(resale_grade);size(11)"`
	Credit        			int       `orm:"column(credit);size(11)"`
	CreditHappy        		int       `orm:"column(credit_happy);size(11)"`
	CreditCoin        		int       `orm:"column(credit_coin);size(11)"`
	Vouchers        		int       `orm:"column(vouchers);size(11)"`
	WorthGold        		float64   `orm:"column(worth_gold);digits(12);decimals(4)"`
	Balance        			float64   `orm:"column(balance);digits(12);decimals(4)"`
	Unusable        		float64   `orm:"column(unusable);digits(12);decimals(4)"`
	Income        			float64   `orm:"column(income);digits(12);decimals(4)"`
	Expend        			float64   `orm:"column(expend);digits(12);decimals(4)"`
	ReferralsId        		int       `orm:"column(referrals_id);size(11)"`
	LastOnlineTime        	int       `orm:"column(last_online_time);size(11)"`
	LastLoginTime        	int       `orm:"column(last_login_time);size(11)"`
	LastLoginIp        		int       `orm:"column(last_login_ip);size(11)"`
	CreateTime        		int       `orm:"column(create_time);size(11)"`
	UpdateTime        		int       `orm:"column(update_time);size(11)"`
	ExpriyTime        		int       `orm:"column(expriy_time);size(11)"`
	Email                  	string    `orm:"column(email);size(255)"`
}

func (t *User) TableName() string {
	return "scn_user"
}

func init() {
	orm.RegisterModel(new(User))
}


func AddUser(m *User) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}


func GetUserById(id int) (v *User, err error) {
	o := orm.NewOrm()
	v = &User{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}


func GetAllUser(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(User))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []User
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}


func UpdateUserById(m *User) (err error) {
	o := orm.NewOrm()
	v := User{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}


func DeleteUser(id int) (err error) {
	o := orm.NewOrm()
	v := User{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&User{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
func GetUserByMobile(mobile string) (v User, err error) {
	o := orm.NewOrm()
	err = o.QueryTable(new(User)).Filter("mobile", mobile).One(&v,"Id")
	beego.Info(err)
	if err == orm.ErrMultiRows {
		// 多条的时候报错
		fmt.Printf("Returned Multi Rows Not One")
		return v, err
	}
	if err == orm.ErrNoRows {
		// 没有找到记录
		fmt.Printf("Not row found")
		return v, err
	}
	return v, nil
}