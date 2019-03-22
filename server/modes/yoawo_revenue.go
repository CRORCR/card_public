package modes

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L card_public/lib -lshane
#include "card_public/config/shane.h"
*/
import "C"
import (
	"card_public/server/db"
	"fmt"
	"time"
)

const REVENUETABLE = "yoawo_revenue"

/*
 * 描述：云握营收表
 *
 * account_type : 账 号  类 型: n < 100 全为真实数据 n > 100 全为机器人数据
 *
 ********************************************************************************************/
type YoawoRevenue struct {
	Id           int64  `xorm:"id"`            // 表        ID
	BillNo       string `xorm:"bill_no"`       // 账   单   号
	Source       int64  `xorm:"source"`        // 来        源
	PayType      int64  `xorm:"pay_type"`      // 支 付  类 型
	PayAmount    int64  `xorm:"pay_amount"`    // 支 付  金 额
	PayAccount   string `xorm:"pay_account"`   // 支付 者 账号
	PayRefund    int64  `xorm:"pay_refund"`    // 退 款  金 额
	AccountType  int64  `xorm:"account_type"`  // 账 号  类 型
	UserNickname string `xorm:"user_nickname"` // 支付者的昵称
	UserIcon     string `xorm:"user_icon"`     // 支付者的头像
	YoawoAccount string `xorm:"yoawo_account"` // 云握收款账号
	AreaId       int64  `xorm:"area_id"`       // 支 付  区 域
	CreateAt     int64  `xorm:"create_at"`     // 创 建  时 间
	UserName     string `xorm:"user_name"`     // 名        称
	Status       int64  `xorm:"status"`        // 状        态
}

func (this *YoawoRevenue) Save(inPara *YoawoRevenue, outPara *int64) error {
	var err error
	this.CreateAt = time.Now().Unix()
	*outPara, err = db.GetDBHand(0).Table(REVENUETABLE).Insert(inPara)
	return err
}

/*
 * desc: 获取本单交易单号
 *
 *************************************************************************************/
func (this *YoawoRevenue) GetTarNumber(Amount *int64, outPara *string) error {
	nY, nM, nD := time.Now().Date()
	vas := C.EncrData(C.int(*Amount))
	fmt.Println(vas)
	*outPara = fmt.Sprintf("%d%.2d%.2d%d%.12d", nY, nM, nD, time.Now().Unix, uint64(vas))
	return nil
}
