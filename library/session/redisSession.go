package session

import "github.com/astaxie/beego/session"

var globalSessions *session.Manager

func init() {
	sessionConfig := &session.ManagerConfig{
		CookieName:"gosessionid",
		EnableSetCookie: true,
		Gclifetime:3600,
		Maxlifetime: 3600,
		Secure: false,
		CookieLifeTime: 3600,
		ProviderConfig: "127.0.0.1:6379,100,astaxie",
	}
	globalSessions, _ = session.NewManager("redis",sessionConfig)
	go globalSessions.GC()
}
func setSessionRedis()  {

}