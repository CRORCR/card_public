package modes

import (
	"card_public/server/db"
	"time"
)

/*
 * 描述：提现记录表
 *
 *      wit_type   : 提现类型 1 现金提现， 2 微信提现， 3 支付宝提现
 *      state	   : 状    态 0 提交，1 审核， 2 审核通过，3 审核未通过， 100 到帐
 *
 ********************************************************************************************/
type MWithdrawalFoot struct {
	Id           int64   `json:"id" xorm:"not null 'id'"`                       // 表    ID
	WitId        string  `json:"wit_id" xorm:"not null 'wit_id'"`               // 提现单号
	WitType      int64   `json:"wit_type" xorm:"not null 'wit_type'"`           // 提现类型
	Amount       float64 `json:"amount" xorm:"not null 'amount'"`               // 提现金额
	Poundage     float64 `json:"poundage" xorm:"not null 'poundage'"`           // 手 续 费
	SubmissionAt int64   `json:"submission_at" xorm:"not null 'submission_at'"` // 提交时间
	State        int64   `json:"state" xorm:"not null 'state'"`                 // 状    态
	ArriveAt     int64   `json:"arrive_at" xorm:"not null 'arrive_at'"`         // 到帐时间
	DescInfo     string  `json:"desc_info" xorm:"not null 'desc_info'"`         // 描    述
	UserId       string  `json:"user_id" xorm:"not null 'user_id'"`             // 用 户 ID
	UserName     string  `json:"user_name" xorm:"not null 'user_name'"`         // 用户姓名
	MerchantId   string  `json:"merchant_id" xorm:"not null 'merchant_id'"`     // 商 户 ID
	MerchantName string  `json:"merchant_name" xorm:"not null 'merchant_name'"` // 商户名称
	UserPhone    string  `json:"user_phone" xorm:"not null 'user_phone'"`       // 手 机 号
	TargetId     string  `json:"target_id" xorm:"not null 'target_id'"`         // 目 标 ID
}

/*
 * 描述：添加提现记录
 *
 *******************************************************************************/
func (this *MWithdrawalFoot) Save(inPara *MWithdrawalFoot, outPara *int64) error {
	var err error
	*outPara, err = db.GetDBHand(0).Table("car_withdrawal_foot").Insert(inPara)
	return err
}

/*
 * 描述：获取本条记录
 *
 *******************************************************************************/
func (this *MWithdrawalFoot) Get(inPara *string, outPara *MWithdrawalFoot) error {
	_, err := db.GetDBHand(0).Table("car_withdrawal_foot").
		Where("wit_id = ?", *inPara).
		Get(outPara)
	return err
}

type WithdrawalWhere struct {
	Where string	`json:"where"`
	Page  int	`json:"page"`
	Count int	`json:"count"`
}

/*
 * 描述：根据inPara 的Where的条件获取记录
 *
 *******************************************************************************/
func (this *MWithdrawalFoot) QueryWhere(inPara *WithdrawalWhere, outPara *[]MWithdrawalFoot) error {
	return db.GetDBHand(0).Table("car_withdrawal_foot").
		Where(inPara.Where).
		Limit(inPara.Count, inPara.Page).
		Desc("submission_at").
		Find(outPara)
}

/*
 * 描述：更新
 *
 *******************************************************************************/
func (this *MWithdrawalFoot) Update(inPara *MWithdrawalFoot, outPara *int64) error {
	var err error
	*outPara, err = db.GetDBHand(0).Table("car_withdrawal_foot").
		Where("wit_id = ?", inPara.WitId).
		Update(inPara)
	return err
}

/*
 * 描述：获取指定用户已经冻结的金额
 *
 ******************************************************************************/
func (this *MWithdrawalFoot)GetNotUse( inPara *MWithdrawalFoot, outPara *float64 )error{
	var err error
	*outPara, err = db.GetDBHand(0).Table("car_withdrawal_foot").
                                Where("merchant_id = ? AND state < ?", inPara.MerchantId, 3 ).
                                Sum( new( MWithdrawalFoot ) ,"amount")
	return err
}

/*
 * 描述：获取指定用户已提现的金额
 *
 ******************************************************************************/
func (this *MWithdrawalFoot)GetSuccess( inPara *MWithdrawalFoot, outPara *float64 )error{
        var err error
        *outPara, err = db.GetDBHand(0).Table("car_withdrawal_foot").
                                Where("merchant_id = ? AND state = ?", inPara.MerchantId, 100 ).
                                Sum( new( MWithdrawalFoot ) ,"amount")
        return err
}

/*
 * 描述：获取指定条件下的总条数
 *
 ******************************************************************************/
func (this *MWithdrawalFoot)GetWhereCount( inPara *string, outPara *int64 )error{
        var err error
        *outPara, err = db.GetDBHand(0).Table("car_withdrawal_foot").
                                Where( *inPara ).
                                Count( new(MWithdrawalFoot) )
        return err
}

/*
 * 描述：获取本月已经提现的金额
 *  
 ******************************************************************************/
func (this *MWithdrawalFoot)GetMonthWithdrawal( inPara *MWithdrawalFoot, outPara *float64 )error {
	var err error
	nY,nM ,_ := time.Now().Date()
        nMTime := time.Date(nY, nM, 1, 0, 0, 0, 0, time.Local).Unix()
	*outPara, err = db.GetDBHand(0).Table("car_withdrawal_foot").
		Where("merchant_id = ? AND submission_at > ? AND state = ?", inPara.MerchantId, nMTime, 100).
		Sum(new(MWithdrawalFoot), "amount")
	return err
}

/*
 * 描述：获取本月已经提现的次数
 *
 ******************************************************************************/
func (this *MWithdrawalFoot) GetMonthWithdrawalCount(inPara *MWithdrawalFoot, outPara *int64) error {
	var err error
	nY, nM, _ := time.Now().Date()
	nMTime := time.Date(nY, nM, 1, 0, 0, 0, 0, time.Local).Unix()
	*outPara, err = db.GetDBHand(0).Table("car_withdrawal_foot").
		Where("Merchant_id = ? AND submission_at > ?", inPara.MerchantId, nMTime).
		Count(new(MWithdrawalFoot))
	return err
}

/*
 * 描述：所有记录
 *
 *******************************************************************************/
func (this *MWithdrawalFoot) List(inPara *string, outPara *[]MWithdrawalFoot) error {
	return db.GetDBHand(0).Table("car_withdrawal_foot").
		Where("merchant_id = ?", *inPara).
		Desc("submission_at").
		Limit(50, 0).
		Find(outPara)
}

/*
type WithdrawalPage struct {
	MerchantId string `json:"merchant_id"`
	Page       int    `json:"page"`
	Count      int    `json:"count"`
}
 *
 * 描述：所有记录
 *
 *******************************************************************************
func (this *MWithdrawalFoot)ListPage( inPara *WithdrawalPage, outPara *[]MWithdrawalFoot)error{
        return db.GetDBHand(0).Table("car_withdrawal_foot").
                Where("merchant_id = ?", inPara.MerchantId ).
                Desc("submission_at").
                Limit( inPara.Page, inPara.Count ).
                Find( outPara )
}
*/
