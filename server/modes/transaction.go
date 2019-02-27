package modes

import (
	"card_public/server/db"
	"time"
	"fmt"
	"sort"
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
	MerchantName string  `json:"merchant_name" xorm:"merchant_name"` // 商家名称
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
	_, err := db.GetDBHand(0).Table( strName ).
		Where("tran_id = ?", inPara.TranId).
		Get(outPara)
	return err
}

/*
 * desc: 添加交易
 *
 *************************************************************************************/
func (this *TransactionFoot)Add(inPara, outPara *TransactionFoot) error {
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
	Id    string // 商家ID或用户ID
	Count int    // 单页的数量
	Page  int    // 页码
	Where string //自定义查询条件
}

/*
 * desc: 获取所有交易记录( 商家 )
 *
 *************************************************************************************/
func (this *TransactionFoot)MerchantGetAll(inPara *TransactionInfo, outPara *TransactionList )error{
	var val TransactionFoot
	val.MerchantId = inPara.Id
	strName, _ := val.name()
	return db.GetDBHand(0).Table(strName).Where("merchant_id = ?", inPara.Id).
		Desc("create_at").
		Limit(inPara.Count, inPara.Page).
		Find(outPara)
}

/*
 * desc: 获取所有交易记录( 商家 )
 *
 *************************************************************************************/
func (this *TransactionFoot) MerchantQuery(inPara *TransactionInfo, outPara *TransactionList) error {
	var val TransactionFoot
	val.MerchantId = inPara.Id
	strName, _ := val.name()
	return db.GetDBHand(0).Table(strName).Where("merchant_id = ?", inPara.Id).And(inPara.Where).
		Desc("create_at").
		Limit(inPara.Count, inPara.Page).
		Find(outPara)
}

type CashierFoot struct {
        MerchantId	string  // 商家ID、收银员ID
	CashierId	string  // 收银员ID
        Count		int     // 单页的数量
        Page		int     // 页码
}

/*
 * desc: 获取指定收银员两天的收款记录
 *
 *************************************************************************************/
func (this *TransactionFoot)CashierGetList(inPara *CashierFoot, outPara *TransactionList )error{
        var val TransactionFoot
        val.MerchantId = inPara.MerchantId
        strName, _ := val.name()
	nNowTime := time.Now().Unix()

	nNowTime -= ((nNowTime + 28800) % 86400) - 86400
	return db.GetDBHand(0).Table(strName).Where("merchant_id = ? AND cashier_id = ?", inPara.MerchantId, inPara.CashierId).
		Where("create_at > ?", nNowTime).
		Desc("create_at").
		Limit(inPara.Count, inPara.Page).
		Find(outPara)
}

/*
 * desc: 获取指定收银员两天的收款记录和
 *
 *************************************************************************************/
func (this *TransactionFoot) GetUserCash(inPara *CashierFoot, outPara *float64) error {
	var val TransactionFoot
	val.MerchantId = inPara.MerchantId
	strName, _ := val.name()
	nNowTime := time.Now().Unix()
	*outPara = -1
	nNowTime -= ((nNowTime + 28800) % 86400) - 86400
	f, err := db.GetDBHand(0).Table(strName).Where("merchant_id = ? AND cashier_id = ?", inPara.MerchantId, inPara.CashierId).
		Where("create_at > ?", nNowTime).
		Desc("create_at").Sum(new(struct{ Trust float64 }), "amount")
	if err != nil {
		return err
	}
	r, err := db.GetDBHand(0).Table(strName).Where("merchant_id = ? AND cashier_id = ?", inPara.MerchantId, inPara.CashierId).
		Where("create_at > ?", nNowTime).
		Desc("create_at").Sum(new(struct{ Trust float64 }), "refund_amount")
	*outPara = f - r
	return err
}



/*
 * desc: 获取指定收银员两天的收款总和
 *
 *************************************************************************************/
func (this *TransactionFoot)CashierGetSum(inPara *CashierFoot, outPara *float64 )error{
        var val TransactionFoot
	var err error
        val.MerchantId = inPara.MerchantId
        strName, _ := val.name()
        nNowTime := time.Now().Unix()
        //checkTime -= (checkTime + 28800) % 86400
        nNowTime -= (( nNowTime + 28800 ) % 86400 ) - 86400
        *outPara,err = db.GetDBHand(0).Table(strName).Where("merchant_id = ? AND cashier_id = ?",inPara.MerchantId,inPara.CashierId ).
                                        Where( "create_at > ?", nNowTime ).
					Sum( *inPara ,"amount")
                                        //Limit( inPara.Count, inPara.Page ).
	return err
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

type WhereFind struct {
	UserId string	// 用户ID
	Where  string	// SQL 的Where 条件
	Count  int	// 单页的数量
	Page   int	// 页码
}
/*
 * desc: 自定义Where条件,获取所有交易记录( 用户 )
 *
 *************************************************************************************/
func (this *TransactionFoot) UserWhereFind(inPara *WhereFind, outPara *TransactionList) error {
	var user UserFoot
	user.UserId = inPara.UserId
	silAreaNumber, sErr := user.getAll()
	if nil == sErr {
		for _, v := range silAreaNumber {
			strTableName := fmt.Sprintf("car_transaction_%s", v)
			db.GetDBHand(0).Table(strTableName).
					Where( inPara.Where ).
					Limit( inPara.Count, inPara.Page ).
					Find(outPara)
		}
	}
	sort.Sort(outPara)
	return sErr
}


// =========================================================================================================
/*
查询个人所有的交易记录和提现记录
提现记录
交易记录
叫鸡记录
*/
func (this *TransactionFoot) All(inPara *string, outPara *[]map[string][]byte) error {

	var err error
	strSqlHand := fmt.Sprintf("SELECT a.create_at, a.amount, a.oper_type, a.state, a.target_name, a.buy_id FROM ")
	strWithdrawal := fmt.Sprintf("SELECT 1 AS oper_type, submission_at AS create_at, wit_id AS buy_id, '云握' AS target_name, state, amount FROM chi_withdrawal_foot WHERE user_id = '%s'", *inPara )

	var user UserFoot
        user.UserId = *inPara
        silAreaNumber, _ := user.getAll()
	var strTransaction string
	for _, v := range silAreaNumber {
		strTransaction = fmt.Sprintf("%sSELECT 3 AS oper_type, tran_id AS buy_id, status AS state, merchant_name AS target_name, amount, create_at FROM car_transaction_%.3d WHERE user_id = '%s' UNION ALL ",strTransaction, v, *inPara )
	}

	strExchange := fmt.Sprintf("SELECT 2 AS oper_type, 0 AS state, trust_count AS amount, '云握' AS target_name, 'null' AS buy_id, create_at FROM chi_exchange_foot WHERE user_id = '%s'", *inPara )

	strSqlEnd  := "AS a ORDER BY create_at DESC;"

	strSql := fmt.Sprintf("%s ( %s UNION ALL %s%s ) %s", strSqlHand, strWithdrawal, strTransaction, strExchange , strSqlEnd )
	fmt.Println("-------------------------------------------------------------------------------------------------")
	fmt.Println( strSql )
	fmt.Println("-------------------------------------------------------------------------------------------------")
        *outPara, err = db.GetDBHand(0).Query( strSql )
	fmt.Println( outPara )
        return err
}


