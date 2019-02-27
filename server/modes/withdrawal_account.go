package modes

import (
	"card_public/server/db"
)

/*
 * 描述：商家提现管理表
 *
 *	car_type: 证件类型, 身份证：IDENTITY_CARD、护照：PASSPORT、军官证：OFFICER_CARD、士兵证：SOLDIER_CARD、户口本：HOKOU等。
 *	is_main : 主 账 号， 'TRUE' 是， 'FALSE', 否
 *	account_type: 微信- WECHAT 支付宝 - ALIPAY
 *	state : 状态: 0 正常 1 删除
 *
 ********************************************************************************************/
type WithdrawalAccount struct {
	Id          int64  `json:"id" xorm:"not null 'id'"`                     // 表    ID
	MerchantId  string `json:"merchant_id" xorm:"not null 'merchant_id'"`   // 商 户 ID
	AccountType string `json:"account_type" xorm:"not null 'account_type'"` // 账号类型
	Account     string `json:"account" xorm:"not null 'account'"`           // 账    号
	Phone       string `json:"phone" xorm:"not null 'phone'"`               // 手 机 号
	UserName    string `json:"user_name" xorm:"not null 'user_name'"`       // 用户姓名
	CarType     string `json:"car_type" xorm:"not null 'car_type'"`         // 证件类型
	CarNumber   string `json:"car_number" xorm:"not null 'car_number'"`     // 证件号码
	CreateAt    int64 `json:"create_at" xorm:"not null 'create_at'"`       // 创建时间
	State       int64  `json:"state" xorm:"not null 'state'"`               // 当前状态
	IsMain      string `json:"is_main" xorm:"not null 'is_main'"`           // 主 账 号
}

/*
 * 描述：添加提现账号
 *
 *******************************************************************************/
func (this *WithdrawalAccount) Save(inPara *WithdrawalAccount, outPara *int64) error {
	var err error
	*outPara, err = db.GetDBHand(0).Table("car_withdrawal_account").Insert(inPara)
	return err
}

/*
 * 描述：根据inPara 的Where的条件获取记录
 *
 *******************************************************************************/
func (this *WithdrawalAccount) QueryWhere(inPara *WithdrawalWhere, outPara *[]WithdrawalAccount) error {
	return db.GetDBHand(0).Table("car_withdrawal_account").
		Where(inPara.Where).
		Limit(inPara.Count, inPara.Page).
		Desc("id").
		Find(outPara)
}

/*
 * 描述：获取指定条件下的总条数
 *
 ******************************************************************************/
func (this *WithdrawalAccount)GetWhereCount( inPara *string, outPara *int64 )error{
	var err error
	*outPara, err = db.GetDBHand(0).Table("car_withdrawal_account").
		Where( *inPara ).
		Count( new(WithdrawalAccount) )
	return err
}

/*
 * 描述：查询一条记录
 *
 *******************************************************************************/
func (this *WithdrawalAccount) GetWithdrawalAccount(inPara *WithdrawalAccount, outPara *WithdrawalAccount) error {
	_,err:=db.GetDBHand(0).Table("car_withdrawal_account").Get(outPara)
	return err
}



/*
 * 描述：所有账号
 *
 *******************************************************************************/
func (this *WithdrawalAccount) List(inPara *string, outPara *[]WithdrawalAccount) error {
	return db.GetDBHand(0).Table("car_withdrawal_account").
		Where("merchant_id = ?", *inPara).
		Find(outPara)
}
