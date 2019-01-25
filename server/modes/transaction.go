package modes

import (
	"public/server/db"
	"fmt"
)

type TransactionFoot struct {
	Id           int64   `json:"id" xorm:"id"`                       // 表    ID
	TranTd       string  `json:"tran_id" xorm:"tran_id"`             // 交 易 ID
	UserId       string  `json:"user_id" xorm:"user_id"`             // 用 户 ID
	UserPhone    string  `json:"user_phone" xorm:"user_phone"`       // 用户手机号
	UserName     string  `json:"user_name" xorm:"user_name"`         // 用户姓名
	CashierId    string  `json:"cashier_id" xorm:"cashier_id"`       // 收 银 ID
	MerchantId   string  `json:"merchant_id" xorm:"merchant_id"`     // 商 家 ID
	Rate         float64 `json:"rate" xorm:"rate"`                   // 本单费率金额
	Amount       float64 `json:"amount" xorm:"amount"`               // 交易金额
	RefundAmount float64 `json:"refund_amount" xorm:"refund_amount"` // 退款金额
	MerBalance   float64 `json:"mer_balance" xorm:"mer_balance"`     // 商户余额
	TranType     int64   `json:"tran_type" xorm:"tran_type"`         // 交易类型
	NoteTest     string  `json:"note_test" xorm:"note_test"`         // 备    注
	Status       int64   `json:"status" xorm:"status"`               // 当前状态
	CreateAt     int64   `json:"create_at" xorm:"create_at"`         // 创建时间
	UpdateAt     int64   `json:"update_at" xorm:"update_at"`         // 更新时间
}

type TransactionList []TransactionFoot

func (this TransactionList) Len() int {
	return len(this)
}
func (this TransactionList) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
func (this TransactionList) Less(i, j int) bool {
	return this[i].UpdateAt < this[j].UpdateAt
}

func (this *TransactionFoot) name() string {
	var val MerchantInfo
	val.MerchantId = this.MerchantId
	val.GetAreaNumber()
	return fmt.Sprintf("car_merchant_%d", val.AreaNumber)
}

/*
 * desc: 获本条记录
 *
 *************************************************************************************/
func (this *TransactionFoot) Get(inPara, outPara *TransactionFoot) error {
	_, err := db.GetDBHand(0).Table(inPara.name()).
		Where("tran_id = ?", inPara.TranTd).
		Get(outPara)
	return err
}

/*
 * desc: 添加交易
 *
 *************************************************************************************/
func (this *TransactionFoot) Add(inPara, outPara *TransactionFoot) error {
	//_, err := db.GetDBHand(0).Table( inPara.name() ).In
	return nil
}
