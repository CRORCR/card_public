package modes

import (
	"public/server/db"
	//	"../../lib"
	//	"strconv"
	//	"errors"
	"fmt"
	"sort"
	//	"github.com/go-redis/redis"
	//	"strings"
)

const USERFOOT = "USER_FOOT_SET_"	// 此用户消费的区的纪录

type UserFoot struct {
	UserId	string		// 用 户 ID
	//AreaId []string	用户所消费的区ID
}

func ( this *UserFoot )name()string{
	return fmt.Sprintf("%s%s", USERFOOT, this.UserId )
}

/*
 * desc: 向此用户消费的区的纪录添加记录
 *
 *	前置条件: 如果 strAreaId 不存在添加
 *
 *************************************************************************************/
func ( this *UserFoot )isAdd( nAreaNumber int64 )error{
	return db.GetRedis().SAdd( this.name() ,nAreaNumber ).Err()
}


/*
 * desc: 获取用户所有的消费区的ID
 *
 *************************************************************************************/
func ( this *UserFoot )getAll()( []string, error ){
	fmt.Println("用户所有的消费区的ID:", this.name())
	return db.GetRedis().SMembers( this.name() ).Result()
}

type TransactionFoot struct {
	Id           int64   `json:"id" xorm:"id"`                       // 表    ID
	TranId       string  `json:"tran_id" xorm:"tran_id"`             // 交 易 ID
	UserId       string  `json:"user_id" xorm:"user_id"`             // 用 户 ID
	UserPhone    string  `json:"user_phone" xorm:"user_phone"`       // 用户手机号
	UserName     string  `json:"user_name" xorm:"user_name"`         // 用户姓名
	CashierId    string  `json:"cashier_id" xorm:"cashier_id"`       // 收 银 ID
	MerchantId   string  `json:"merchant_id" xorm:"merchant_id"`     // 商 家 ID
	Rate         float64 `json:"rate" xorm:"rate"`                   // 本单费率金额
	Amount       float64 `json:"amount" xorm:"amount"`               // 交易金额
	RefundAmount float64 `json:"refund_amount" xorm:"refund_amount"` // 退款金额
	MerBalance   float64 `json:"mer_balance" xorm:"mer_balance"`     // 商户余额
	TranType     int64   `json:"tran_type" xorm:"tran_type"`         // 交易类型  0:现金  1:诺
	NoteTest     string  `json:"note_test" xorm:"note_test"`         // 备    注
	Status       int64   `json:"status" xorm:"status"`               // 当前状态  0:支付  1:退款
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

func (this *TransactionFoot)name()(string, int64 ) {
	var val MerchantInfo
	val.MerchantId = this.MerchantId
	val.GetAreaNumber()
	return fmt.Sprintf("car_transaction_%d", val.AreaNumber), val.AreaNumber
}

/*
 * desc: 获本条记录
 *
 *************************************************************************************/
func (this *TransactionFoot) Get(inPara, outPara *TransactionFoot) error {
	strName, _ := inPara.name()
	_, err := db.GetDBHand(0).Table(strName).Where("tran_id = ?", inPara.TranId).Get(outPara)
	return err
}

/*
 * desc: 添加交易
 *
 *************************************************************************************/
func (this *TransactionFoot) Add(inPara, outPara *TransactionFoot) error {
	strName, nAreaNumber := inPara.name()
	var err error
	var user UserFoot
	user.UserId = inPara.UserId
	if err = user.isAdd( nAreaNumber ); err != nil {
		return err
	}
	var merc MerchantInfo
	merc.MerchantId = inPara.MerchantId
	if err = merc.Transaction( inPara.Amount ); err == nil {
		_, err = db.GetDBHand(0).Table( strName ).Insert(inPara)
	}
	return err
}

type TransactionInfo struct {
	Id	string	// 商家或用户ID
	Count	int	// 单页的数量
	Page	int	// 页码
}

/*
 * desc: 获取所有交易记录( 商家 )
 *
 *************************************************************************************/
func (this *TransactionFoot)MerchantGetAll(inPara *TransactionInfo, outPara *TransactionList )error{
	var val TransactionFoot
	val.MerchantId = inPara.Id
	strName, _ := val.name()
	db.GetDBHand(0).Table( strName ).Where("merchant_id = ?", inPara.Id ).
					Desc("create_at").
					Limit( inPara.Count, inPara.Page ).
					Find( outPara )
	return nil
}

/*
 * desc: 获取所有交易记录( 用户 )
 *
 *************************************************************************************/
func (this *TransactionFoot)UserGetAll(inPara *string, outPara *TransactionList )error{
	var user UserFoot
	user.UserId = *inPara
	silAreaNumber, sErr := user.getAll()
	if nil == sErr {
		for _, v := range silAreaNumber {
			strTableName := fmt.Sprintf("car_transaction_%s", v )
			db.GetDBHand(0).Table( strTableName ).Where("user_id = ?", *inPara).Find( outPara )
		}
	}
	sort.Sort(outPara)
	return sErr
}

