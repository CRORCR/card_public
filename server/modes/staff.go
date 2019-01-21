package modes

import (
	"public/server/db"
	"time"

	//	"errors"
	"fmt"
	"strconv"
	//	"strings"
	//	"time"
)

const STAFF_HASH = "STAFF_HASH_"

type StaffInfo struct {
	UserId     string // 员工ID
	MerchantId string // 商家ID
	AreaNumber int64  // 商家所在地的区号
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
		this.AreaNumber, _ = strconv.ParseInt(sKey["AreaNumber"], 10, 64)
	}
	fmt.Printf("sKey:%+v\n",sKey)
	return sErr
}

/*
 * 描述: 添加员工
 *
 *************************************************************************/
func (this *StaffInfo) addStaff() {
	client := db.GetRedis()
	strName := this.name()
	client.HSet(strName, "UserId", this.UserId)
	client.HSet(strName, "MerchantId", this.MerchantId)
	client.HSet(strName, "AreaNumber", this.AreaNumber)
}

func (this *StaffInfo) delStaff() {
	client := db.GetRedis()
	strName := this.name()
	b, e := client.Expire(strName, 1*time.Second).Result()
	fmt.Println("删除redis",b,e)
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
	Sex        bool   `json:"sex" xorm:"sex"`                 // 性    别
	CreateAt   int64  `json:"-" xorm:"create_at"`             // 创建时间
	State      int64  `json:"state" xorm:"state"`             // 状    态
	NumberFage int64  `json:"number_fage" xorm:"number_fage"` // 身份标识
	Authority  int64  `json:"authority" xorm:"authority"`     // 权    限
	CreateStr  string `json:"create_at" xorm:"-"`             // 更新时间供前端展示
}

type StaffList []Staff

func (this *Staff) name() (string,int64) {
	var val StaffInfo
	val.UserId = this.UserId
	val.getAll()
	fmt.Printf("area_number:%v\n",val.AreaNumber)
	return fmt.Sprintf("chi_staff_%d", val.AreaNumber),val.AreaNumber
}

func (this *Staff) getInfo() (StaffInfo, error) {
	var val StaffInfo
	val.UserId = this.UserId
	err := val.getAll()
	return val, err
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
 * 描述: 获取员工信息表
 *
 *************************************************************************/
func (this *Staff) Get(inPara *Staff, outPara *AddStaff) error {
	name, area := inPara.name()
	_, err := db.GetDBHand(0).Table(name).Get(inPara)
	outPara.PStaff=*inPara
	outPara.AreaNumber=area
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
	PStaff     Staff `json:"p_staff"`// 员工信息
	AreaNumber int64 `json:"area_number"`// 区号ID
}

/*
 * 描述: 添加员工
 *
 *	nAreaNumber: 所在的区号 例如: 邯郸:310, 邢台:319, 北京:10
 *
 *************************************************************************/
func (this *Staff) Add(inPara *AddStaff, outPara *Staff) error {
	var val StaffInfo
	val.MerchantId = inPara.PStaff.MerchantId
	val.UserId = inPara.PStaff.UserId
	val.AreaNumber = inPara.AreaNumber
	val.addStaff()
	name, _ := inPara.PStaff.name()
	_, err := db.GetDBHand(0).Table(name).Insert(&inPara.PStaff)
	return err
}

/**
 * @desc   : 更新员工
 * @author : Ipencil
 * @date   : 2019/1/21
 */
func (this *Staff) Update(inPara *AddStaff, outPara *Staff) error {
	name, _ := inPara.PStaff.name()
	_, err := db.GetDBHand(0).Table(name).Where("phone=? ",inPara.PStaff.Phone).Update(&inPara.PStaff)
	return err
}

/**
 * @desc   : 删除员工
 * @author : Ipencil
 * @date   : 2019/1/21
 */
func (this *Staff) Delete(inPara *Staff, outPara *Staff) error {
	var val StaffInfo
	val.UserId = inPara.UserId
	val.delStaff()
	name, _ := inPara.name()
	_, err := db.GetDBHand(0).Table(name).Get(inPara)
	if err!=nil{
		return err
	}
	_, err = db.GetDBHand(0).Table(name).Where("id=?",inPara.Id).Delete(inPara)
	if err!=nil{
		return err
	}
	return err
}