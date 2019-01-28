package modes

import (
	"public/server/db"
	"os"
)

const INVITECODE = "INVITECODE"

var pf *os.File

func InviteInit(strName string) error {
	var err error
	pf, err = os.OpenFile(strName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	return err
}

/*
 * 描述: 发放邀请码
 *
 *************************************************************/
func GetInvitecode() string {
	count, err := db.GetRedis().IncrBy(INVITECODE, 1).Result()
	if nil == err {
		buf := make([]byte, 6)
		pf.Seek((count * 6), 0)
		if _, err := pf.Read(buf); err == nil {
			return string(buf)
		}
	}
	return ""
}

