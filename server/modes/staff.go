package modes

import (
	"public/server/db"
	"public/lib"
	"errors"
	"fmt"
	"strconv"
	"math"
	//	"strings"
	//	"time"
)

const STAFF_HASH  = "STAFF_HASH_"
const STAFF_PHONE = "PHONE_SET_"

type StaffInfo struct {
	UserId     string // 员工ID
	Name       string // 员工姓名
	MerchantId string // 商家ID
	Phone      string // 员工手机号
	AreaNumber int64  // 商家所在地的区号
	Identity   int64  // 身份类型 1: 店主 2: 员工
}

func (this *StaffInfo) name() string {
	return fmt.Sprintf("%s%s", STAFF_HASH, this.UserId)
}

/*
 * 描述: 获取员工在Redis的信息
 *
 *************************************************************************/
func (this *StaffInfo) getAll() error {
	client := db.GetRedis()
	sKey, sErr := client.HGetAll(this.name()).Result()
	if nil == sErr {
		this.UserId, _ = sKey["UserId"]
		this.MerchantId, _ = sKey["MerchantId"]
		this.Phone, _ = sKey["Phone"]
		this.Name = sKey["Name"]
		this.AreaNumber, _ = strconv.ParseInt(sKey["AreaNumber"], 10, 64)
		this.Identity, _ = strconv.ParseInt(sKey["Identity"], 10, 64)
	}
	return sErr
}

/*
 * 描述: 查看员工存不存在
 *
 *************************************************************************/
func (this *StaffInfo) Exists()( int64,error ) {
	return db.GetRedis().Exists(this.name()).Result()
}

/*
 * 描述: 删除员工
 *
 *	前置条件: UserId 与 Phone 不可以为空
 *
 *************************************************************************/
func (this *StaffInfo) Del() {
	client := db.GetRedis()
	strKey := fmt.Sprintf("%s%s", STAFF_PHONE, this.Phone)
	client.Del(this.name())
	client.Del(strKey)
}

/*
 * 描述: 添加员工
 *
 *************************************************************************/
func (this *StaffInfo) addStaff(strPhone string) error {
	mapStaff := lib.ToMap(*this)
	bn, _ := db.GetRedis().Exists(this.name()).Result()
	if 1 == bn {
		return errors.New(fmt.Sprintf("%s : 此用户已经存在!", this.UserId))
	}
	_, err := db.GetRedis().HMSet(this.name(), mapStaff).Result()
	if nil == err {
		strKey := fmt.Sprintf("%s%s", STAFF_PHONE, strPhone)
		db.GetRedis().Set(strKey, mapStaff["UserId"], 0)
	}
	return err
}

/*
 * 描述: 获取用户ID 
 *
 *************************************************************************/
func (this *StaffInfo) getUserId(strPhone string) error {
	var err error
	strKey := fmt.Sprintf("%s%s", STAFF_PHONE, strPhone)
	this.UserId, err = db.GetRedis().Get(strKey).Result()
	return err
}

/*
 * 描述: 获取员工身份标志
 *
 *************************************************************************/
func (this *StaffInfo) getIdentity() int64 {
	nIdentity, _ := db.GetRedis().HGet(this.name(), "Identity").Int64()
	return nIdentity
}

/*
 * 描述: 根据用户手机号获取用户所属的商家ID
 *
 *************************************************************************/
func (this *StaffInfo) phoneToMerchantId( strPhone string )error{
	var err error
	if err = this.getUserId( strPhone ); nil == err {
		this.MerchantId, err = db.GetRedis().HGet(this.name(), "MerchantId").Result()
	}
        return err
}


/*
 * 描述: 获取员工身份标志
 *
 *************************************************************************/
func (this *StaffInfo) getAreaNumber() int64 {
	nIdentity, _ := db.GetRedis().HGet(this.name(), "AreaNumber").Int64()
	return nIdentity
}

/*
 * 描述：商家员工表字段说明
 *
 *	商家员工信息已所在的市区号为标记,例如: 邯郸:310, 邢台:319, 石家庄:311
 *
 * =======================================================================================
 * authority: 
 *      0:1 收银权限
 *      如果员工存在收银权限: 收款码的生成规则
 *              {
 *                      "mid": "12345678901234567890123456789012",
 *                      "uid": "12345678901234567890123456789012"
 *              }
 *      M: 商家ID
 *      U: 用户ID
 *-----------------------------------------------------------------------------------------
 * sex: true 男, false 女
 *-----------------------------------------------------------------------------------------
 * number_fage: 1: 店长
 *
 ********************************************************************************************/
type Staff struct {
	Id         int    `json:"id" xorm:"id"`                   // id主键
	Name       string `json:"name" xorm:"name"`               // 员工姓名
	MerchantId string `json:"merchant_id" xorm:"merchant_id"` // 商 家 ID
	Phone      string `json:"phone" xorm:"phone"`             // 员工手机号
	UserId     string `json:"user_id" xorm:"user_id"`         // 员 工 ID
	NumberId   string `json:"number_id" xorm:"number_id"`     // 身份证号
	Sex        int    `json:"sex" xorm:"sex"`                 // 性    别
	CreateAt   int64  `json:"-" xorm:"create_at"`             // 创建时间
	State      int64  `json:"state" xorm:"state"`             // 状    态
	NumberFage int64  `json:"number_fage" xorm:"number_fage"` // 身份标识
	Authority  uint64 `json:"authority" xorm:"authority"`     // 权    限
	CreateStr  string `json:"create_at" xorm:"-"`             // 更新时间供前端展示
}

type StaffList []Staff

func (this *Staff) name() string {
	var val StaffInfo
	val.UserId = this.UserId
	val.getAll()
	return fmt.Sprintf("chi_staff_%d", val.AreaNumber)
}

func (this *Staff) getInfo() (StaffInfo, error) {
	var val StaffInfo
	val.UserId = this.UserId
	err := val.getAll()
	return val, err
}

/*
 * 描述: 获取本用户的所属区ID
 *
 ****************************************************************************/
func ( this *Staff )GetAreaNumber( strUserId *string, nAreaNumber *int64 )error{
	var staff StaffInfo
	staff.UserId = *strUserId
	*nAreaNumber = staff.getAreaNumber()
	if 0 == *nAreaNumber {
		return errors.New("此用户不存在!")
	}
	return nil
}

/*
 * 描述: 使用手机号换取用户ID
 *
 ****************************************************************************/
func (this *Staff) GetUserId(strPhone *string, strUserId *string) error {
	var staff StaffInfo
	err := staff.getUserId(*strPhone)
	if nil == err {
		*strUserId = staff.UserId
	}
	return err
}

/*
 *
 * 描述: 根据用户手机号获取用户所属的商家ID
 *
 ****************************************************************************/
func (this *Staff)PhoneToMerchantId(strPhone *string, strMerchantId *string) error {
	var val StaffInfo
	err := val.phoneToMerchantId( *strPhone )
	*strMerchantId = val.MerchantId
	return err
}

/*
 * 描述: 生成此用户的收款码
 *
 *************************************************************************/
func (this *Staff) GetQRCode(inPara *Staff, strQRCode *string) error {
	val, err := inPara.getInfo()
	*strQRCode = fmt.Sprintf("{\"mid\":\"%s\",\"uid\":\"%s\"}", val.MerchantId, val.UserId)
	return err
}

/*
 * 描述: 询问是不是店主
 *
 *************************************************************************/
func (this *Staff) AskIdentity(inPara *Staff, nFage *bool) error {
	var val StaffInfo
	val.UserId = inPara.UserId
	*nFage = false
	if 1 == val.getIdentity() {
		*nFage = true
	}
	return nil
}

/*
 * 描述: 获取员工信息表
 *
 *************************************************************************/
func (this *Staff) Get(inPara, outPara *Staff) error {
	outPara.UserId = inPara.UserId
	_, err := db.GetDBHand(0).Table(inPara.name()).Get(outPara)
	return err
}

type StaffAuthority struct {
	UserId	string   `json:"user_id"`
	Fage	int `json:"fage"`
}

/*
 * 描述: 添加权限
 *
 *	StaffAuthority.Fage : - 1 收银
 *	StaffAuthority.Fage : - 2 PC后台登登陆
 *	StaffAuthority.Fage : - 3 退款
 *	StaffAuthority.Fage : - 4 提现
 *
 *************************************************************************/
func (this *Staff) SetAuthority(inPara *StaffAuthority, outPara *Staff) error {
	outPara.UserId = inPara.UserId
	_, err := db.GetDBHand(0).Table(outPara.name()).Get(outPara)
	if nil == err {
		outPara.Authority = outPara.Authority | uint64( math.Pow( float64(2), float64( inPara.Fage - 1 )))
		_, err = db.GetDBHand(0).Table(outPara.name()).
					 Where("user_id = ?", outPara.UserId ).
					 Cols("authority").
				         Update( outPara )
	}
        return err
}

/*
 * 描述: 取消权限
 *
 *************************************************************************/
func (this *Staff)CancelAuthority( inPara *StaffAuthority , outPara *Staff) error {
        outPara.UserId = inPara.UserId
        _, err := db.GetDBHand(0).Table( outPara.name() ).Get( outPara )
        if nil == err {
                outPara.Authority = outPara.Authority &^ uint64( math.Pow( float64(2), float64( inPara.Fage - 1 )))
                _, err = db.GetDBHand(0).Table(outPara.name()).
                                         Where("user_id = ?", outPara.UserId ).
                                         Cols("authority").
                                         Update( outPara )
        }
        return err
}

/*
 * 描述: 查看权限
 *
 *************************************************************************/
func (this *Staff)ShowAuthority( inPara *StaffAuthority , outPara *bool ) error {
	var val Staff
        val.UserId = inPara.UserId
	*outPara = false
        _, err := db.GetDBHand(0).Table( val.name() ).Get( &val )
        if nil == err {
		n := uint64(math.Pow( float64(2), float64( inPara.Fage - 1) ))
		if n == val.Authority & n {
			*outPara = true
		}
        }
        fmt.Println("权限值",val.Authority)
        fmt.Println("权限值",val.Authority&2==2)
        return err
}

/*
 * 描述: 修改员工权限
 *
 *************************************************************************/
func (this *Staff) UpdateAuthority(inPara, outPara *Staff) error {
	_, err := db.GetDBHand(0).Table(inPara.name()).
		Where("merchant_id = ?", inPara.MerchantId).
		Cols("authority").
		Update(inPara)
	outPara = inPara
	return err
}

/*
 * 描述: 获取本用户所属商家Id
 *
 *************************************************************************/
func (this *Staff) GetMerchantId(inPara *Staff, pMerchantId *string) error {
	val, err := inPara.getInfo()
	*pMerchantId = val.MerchantId
	return err
}

type AddStaff struct {
	PStaff     Staff // 员工信息
	AreaNumber int64 // 区号ID
}

/*
 * 描述: 添加员工
 *
 *	nAreaNumber: 所在的区号 例如: 邯郸:310, 邢台:319, 北京:10
 *
 *************************************************************************/
func (this *Staff) Add(inPara *AddStaff, outPara *Staff) error {
	var val StaffInfo
	var err error
	val.MerchantId = inPara.PStaff.MerchantId
	val.UserId = inPara.PStaff.UserId
	val.AreaNumber = inPara.AreaNumber
	val.Identity = inPara.PStaff.NumberFage
	val.Phone = inPara.PStaff.Phone
	if err = val.addStaff(inPara.PStaff.Phone); nil == err {
		_, err = db.GetDBHand(0).Table(inPara.PStaff.name()).Insert(&inPara.PStaff)
		if nil != err {
			val.Del()
		}
	}
	return err
}

/**
 * @desc   : 更新员工
 */
func (this *Staff) Update(inPara *AddStaff, outPara *Staff) error {
	name := inPara.PStaff.name()
	_, err := db.GetDBHand(0).Table(name).Where("user_id=? ", inPara.PStaff.UserId).Update(&inPara.PStaff)
	return err
}

/*
 * desc: 删除管理员信息  支持根据商家ID或者员工ID组合删除
 * 
 **************************************************************************/
func (this *Staff) Del(inPara, outPara *Staff) error {
	var val StaffInfo
	val.UserId = inPara.UserId
	err := val.getAll()
	if nil == err {
		_, err = db.GetDBHand(0).Table(inPara.name()).
			Where("user_id = ?", inPara.UserId).
			Delete(inPara)
		if nil == err {
			val.Del()
		}
	}
	return err
}
/*
 * @desc   : 更新收款权限
 * @author : Ipencil
 * @date   : 2019/1/23
 *
func (this *Staff) UpdAuthority(inPara *Staff, outPara *Staff) error {
	name := inPara.name()
	_, err := db.GetDBHand(0).Table(name).Where("merchant_id = ?", inPara.MerchantId).Cols("authority").Update(inPara)
	return err
}
*/
