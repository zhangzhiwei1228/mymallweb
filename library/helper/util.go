package helper

import (
	"strings"
	"crypto/md5"
	"io"
	"fmt"
)
type Util map[string]interface{}
/**
	如果索引!=-1则表示包含该字符串。空字符串""在任何字符串中均存在
 */
func (t *Util) encrypt(pwd string,salt string) string {
	if strings.Index("$.", pwd) != -1 {
		return pwd
	} else {
		w := md5.New()
		io.WriteString(w, pwd)   //将pass写入到w中
		md5str1 := fmt.Sprintf("%x", w.Sum(nil)) + salt  //w.Sum(nil)将w的hash转成[]byte格式
		io.WriteString(w, md5str1)
		md5str2 := "$." + fmt.Sprintf("%x", w.Sum(nil))
		return md5str2
	}
}