package modes

import (
	"card_public/lib"
	"card_public/server/db"
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

const MERCHANT = "MERCHANT_"  // 商家哈西列表
const AREA_LIST = "AREA_LIST" // 商家地理位置坐标
const BRANCH = "BRANCH_SET_"  // 商家分店列表

type MerchantInfo struct {
	MerchantId  string  // 商家ID
	UserId      string  // 用户ID
	ShopName    string  // 商家名称
	AreaNumber  int64   // 商家所在区的ID
	BucklePoint float64 // 本商当前费率
	Count       int64   // 商家交易次数
	Amount      float64 // 所有交易金额
	NowAmount   float64 // 当前金额
	//----------------------------------------------------------------------------
	UnixTime  int64   // 时间标志
	TarNumber int64   // 商家交易编号标志量( 从1递增,步长 1 )
	DayAmount float64 // 今日交易金额
}

func (this *MerchantInfo) Name() string {
	return fmt.Sprintf("%s%s", MERCHANT, this.MerchantId)
}

func (this *MerchantInfo) BranchAreaNumber() string {
	return fmt.Sprintf("%s%s", BRANCH, this.MerchantId)
}

/*
 * 描述: 获取本单的编号
 *
 *************************************************************/
func (this *MerchantInfo) GetTarNumber() error {
	err := this.Get()
	if nil == err {
		if lib.IsToday(this.UnixTime) {
			this.TarNumber += 1
			return db.GetRedis().HIncrBy(this.Name(), "TarNumber", 1).Err()
		} else {
			this.UnixTime = time.Now().Unix()
			this.TarNumber = 1
			this.DayAmount = 0
		}
		err = this.Set()
	}
	return err
}

/*
 * 描述: 获取本商家的区号
 *
 *************************************************************/
func (this *MerchantInfo) GetAreaNumber() error {
	var err error
	sCmd := db.GetRedis().HGet(this.Name(), "AreaNumber")
	this.AreaNumber, err = sCmd.Int64()
	fmt.Println("AreaNumber", this.AreaNumber, "Error:", err)
	return err
}

/*
 * 描述: 获取本商家的区号
 *
 *************************************************************/
func (this *MerchantInfo) Delete() {
	db.GetRedis().Del(this.Name())
	db.GetRedis().Del(this.BranchAreaNumber())
}

/*
 * 描述: 商家交易
 *
 *	fAmount : 交易金额
 *
 *************************************************************/
func (this *MerchantInfo) Transaction(fAmount float64) error {
	err := this.Get()
	if nil == err {
		this.Count++
		this.Amount += fAmount
		this.NowAmount += fAmount
		if lib.IsToday(this.UnixTime) {
			this.DayAmount += fAmount
		} else {
			this.UnixTime = time.Now().Unix()
			this.TarNumber = 1
			this.DayAmount = fAmount
		}
		err = this.Set()
	}
	return err
}

/*
 * 描述: 获取本商家的当前费率
 *
 *************************************************************/
func (this *MerchantInfo) GetBucklePoint() error {
	var err error
	sCmd := db.GetRedis().HGet(this.Name(), "BucklePoint")
	this.AreaNumber, err = sCmd.Int64()
	fmt.Println("AreaNumber", this.AreaNumber, "Error:", err)
	return err
}

/*
 * 描述: 设置本商家的当前费率
 *
 *************************************************************/
func (this *MerchantInfo) SetBucklePoint() error {
	return db.GetRedis().HSet(this.Name(), "BucklePoint", this.BucklePoint).Err()
}

/*
 * 描述: 添加商家
 *
 *******************************************************/
func (this *MerchantInfo) Add(merchant *Merchant) error {
	client := db.GetRedis()
	var staff StaffInfo
	staff.UserId = merchant.UserId
	if nFage, _ := staff.Exists(); 1 == nFage {
		return errors.New("此用户已存在")
	}
	this.MerchantId = merchant.MerchantId
	this.UserId = merchant.UserId
	this.ShopName = merchant.UserName
	this.AreaNumber = merchant.AreaNumber
	this.BucklePoint = merchant.MerchantRate
	mapMerchant := lib.ToMap(*this)
	_, err := client.HMSet(this.Name(), mapMerchant).Result()
	if nil == err {
		strKey := fmt.Sprintf("%d_%s", this.AreaNumber, this.MerchantId)
		geol := &redis.GeoLocation{Name: strKey, Longitude: merchant.Longitude, Latitude: merchant.Latitude}
		_, err = client.GeoAdd(AREA_LIST, geol).Result()
	} else {
		return err
	}
	return err
}

/*
 * 描述: 添加分店的区号
 *
 *	strLower 分店的MerchantId
 *
 *************************************************************/
func (this *MerchantInfo) AddBranch(strLower string) error {
	var err error
	var val MerchantInfo
	val.MerchantId = strLower
	if err = val.GetAreaNumber(); nil == err {
		_, err = db.GetRedis().SAdd(this.BranchAreaNumber(), val.AreaNumber).Result()
	}
	return err
}

/*
 * 描述: 获取所有分店的区号
 *
 *************************************************************/
func (this *MerchantInfo) GetAllAreaNumber() ([]string, error) {
	return db.GetRedis().SMembers(this.BranchAreaNumber()).Result()
}

/*
 * 描述: 获取商家
 *
 *************************************************************/
func (this *MerchantInfo) Get() error {
	client := db.GetRedis()
	sKey, sErr := client.HGetAll(this.Name()).Result()
	if nil == sErr {
		this.UserId, _ = sKey["UserId"]
		this.MerchantId, _ = sKey["MerchantId"]
		this.AreaNumber, _ = strconv.ParseInt(sKey["AreaNumber"], 10, 64)
		this.Count, _ = strconv.ParseInt(sKey["Count"], 10, 64)
		this.UnixTime, _ = strconv.ParseInt(sKey["UnixTime"], 10, 64)
		this.TarNumber, _ = strconv.ParseInt(sKey["TarNumber"], 10, 64)
		this.BucklePoint, _ = strconv.ParseFloat(sKey["BucklePoint"], 64) // 本商当前费率
		this.Amount, _ = strconv.ParseFloat(sKey["Amount"], 64)           // 所有交易金额
		this.NowAmount, _ = strconv.ParseFloat(sKey["NowAmount"], 64)     // 当前金额
		this.DayAmount, _ = strconv.ParseFloat(sKey["DayAmount"], 64)     // 今日交易金额
	}
	if !lib.IsToday(this.UnixTime) {
		this.TarNumber = 0
		this.DayAmount = 0
	}
	return sErr
}

/*
 * 描述: 设置商家
 *
 *************************************************************/
func (this *MerchantInfo) Set() error {
	mapMerchant := lib.ToMap(*this)
	_, err := db.GetRedis().HMSet(this.Name(), mapMerchant).Result()
	return err
}

//*****************以下是提供外部rpc调用**************************************
/*
 * 描述: 商家信息表
 *
 *  type_id     : 商家所属行业ID
 *  status      : 1 认证中 2 认证未通过 3 认证通过， 5 删除
 */
type Merchant struct {
	Id           int64   `json:"id" xorm:"id"`                       //商家表ID
	FID          string  `json:"fid" xorm:"fid"`                     //父商家id
	MerchantId   string  `json:"merchant_id" xorm:"merchant_id"`     //本商家Id
	UserId       string  `json:"user_id" xorm:"user_id"`             //用户分享ID
	InviteCode   string  `json:"invite_code" xorm:"invite_code"`     //商家邀请码
	MerchantType int64   `json:"merchant_type" xorm:"merchant_type"` //商家 行业 类型
	MerchantRate float64 `json:"rate" xorm:"rate"`                   //商家 行业 利率
	//TrustStatus  bool    `json:"trust_status" xorm:"trust_status"`   //是否诺商家 0 否 1 是
	AreaNumber  int64   `json:"area_number" xorm:"area_number"` //市 I   D
	AreaId      int64   `json:"area_id" xorm:"area_id"`         //县 I   D
	CreateAt    int64   `json:"-" xorm:"create_at"`             //创 建 时 间
	CreateAtStr string  `json:"create_at" xorm:"-"`             //创 建 时 间
	Describea   string  `json:"describea" xorm:"describea"`     //描       述
	Address     string  `json:"address" xorm:"address"`         //地       址
	UserName    string  `json:"name" xorm:"name"`               //店铺名称
	Status      int64   `json:"status" xorm:"status"`           //状       态 //1 认证中 2 认证未通过 3 认证通过， 5 删除
	Phone       string  `json:"phone" xorm:"phone"`             //手  机   号
	Icon        string  `json:"icon" xorm:"icon"`               //商 家 头 像
	LoopImg     string  `json:"loopimg" xorm:"loopimg"`         //商家轮播图
	InfoImg     string  `json:"infoimg" xorm:"infoimg"`         //商家详情图
	Video       string  `json:"video" xorm:"video"`             //商家视频介绍
	CheckDesc   string  `json:"checkdesc" xorm:"checkdesc"`     //认证失败描述
	Business    string  `json:"business" xorm:"business"`       //营 业  执 照
	NumberId    string  `json:"number_id" xorm:"number_id"`     //身   份   证
	Industry    string  `json:"industry" xorm:"industry"`       //行 业许可 证
	Longitude   float64 `json:"longitude" xorm:"longitude"`     //经        度
	Latitude    float64 `json:"latitude" xorm:"latitude"`       //纬        度
	Cash        float64 `json:"cash" xorm:"cash"`               //现        金
	Trust       float64 `json:"trust" xorm:"trust"`             //鍩        分
	Credits     float64 `json:"credits" xorm:"credits"`         //积        分
	Distance    float64 `json:"distance" xorm:"-"`              //商家与用户的距离
	lock sync.Mutex `json:"-" xorm:"-"`
}

type MerchantList []Merchant

func (this MerchantList) Len() int {
	return len(this)
}
func (this MerchantList) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
func (this MerchantList) Less(i, j int) bool {
	return this[i].Distance < this[j].Distance
}

func (this *Merchant) name() string {
	var val MerchantInfo
	val.MerchantId = this.MerchantId
	fmt.Println("this.MerchantId", this.MerchantId)
	val.GetAreaNumber()
	return fmt.Sprintf("car_merchant_%d", val.AreaNumber)
}

/*
 * desc: 获取商家
 *
 *************************************************************************************/
func (this *Merchant) Get(inPara, outPara *Merchant) error {
	_, err := db.GetDBHand(0).Table(inPara.name()).Get(inPara)
	*outPara = *inPara
	return err
}

type TarData struct {
	MerchantId string // 入参-商家ID
	Amount     int64  // 入参-交易金额
}

/*
 * desc: 获取本单交易单号
 *
 *************************************************************************************/
func (this *Merchant) GetTarNumber(inPara *TarData, outPara *string) error {
	var val MerchantInfo
	val.MerchantId = inPara.MerchantId
	//nY, nM, nD := time.Now().Date()
	err := val.GetTarNumber()
	if nil == err {
		//vas := C.EncrData(C.int(inPara.Amount))
		//fmt.Println(vas)
		//*outPara = fmt.Sprintf("%d%.2d%.2d%.8d%.12d", nY, nM, nD, val.TarNumber, uint64(vas))
	}
	return err
}

type MerchantAmount struct {
	Count     int64   // 商家交易次数
	Amount    float64 // 所有交易金额
	NowAmount float64 // 当前金额
	TarNumber int64   // 今日交易次数
	DayAmount float64 // 今日交易金额
}

/*
 * desc: 获取本商家金额信息
 *
 *************************************************************************************/
func (this *Merchant) GetMerchantAmount(inPara *string, outPara *MerchantAmount) error {
	var val MerchantInfo
	val.MerchantId = *inPara
	err := val.Get()
	if nil == err {
		outPara.Count = val.Count
		outPara.Amount = val.Amount
		outPara.NowAmount = val.NowAmount
		outPara.TarNumber = val.TarNumber
		outPara.DayAmount = val.DayAmount
	}
	return err
}

/*
 * desc: 更新费率
 *
 *************************************************************************************/
func (this *Merchant) UpdateRate(inPara, outPara *Merchant) error {
	_, err := db.GetDBHand(0).Table(inPara.name()).
		Where("merchant_id = ?", inPara.MerchantId).
		Cols("rate").
		Update(inPara)
	if err == nil {
		var val MerchantInfo
		val.MerchantId = inPara.MerchantId
		val.BucklePoint = inPara.MerchantRate
		err = val.SetBucklePoint()
	}
	return err
}

/*
 * desc: 获取所有分店
 *
 *************************************************************************************/
func (this *Merchant) GetAllBranch(inPara *Merchant, outPara *MerchantList) error {
	if inPara.MerchantId != "" {
		var merchant MerchantInfo
		merchant.MerchantId = inPara.MerchantId
		sData, sErr := merchant.GetAllAreaNumber()
		if sErr == nil {
			for _, v := range sData {
				val := make([]Merchant, 0)
				db.GetDBHand(0).Table(fmt.Sprintf("car_merchant_%s", v)).
					Where("fid = ?", inPara.MerchantId).Find(&val)
				*outPara = append(*outPara, val...)
			}
		}
		return sErr
	}
	return errors.New("成员属性 MerchantId 不可以为空")
}

type MerchantAdd struct{
	UserPhone string 	`json:"user_phone"`
	MercInfo  Merchant 	`json:"merc_info"`
}
/*
 * desc: 添加商家
 *
 *************************************************************************************/
func (this *Merchant) Add( inPara *MerchantAdd, outPara *int64) error {
	var val MerchantInfo
	var err error
	if era := val.Add( &inPara.MercInfo ); nil != era {
		return era
	}
	inPara.MercInfo.InviteCode = GetInvitecode()
	//outPara.InviteCode = "001122"
	if inPara.MercInfo.InviteCode == "" {
		return errors.New("邀请码获取失败")
	}
	//fmt.Println("费率", outPara.MerchantRate)
	//fmt.Println("费率", outPara.UserName)
	*outPara, err = db.GetDBHand(0).Table(inPara.MercInfo.name()).Insert( inPara.MercInfo )
	if nil == err {
		var add AddStaff
		var staff Staff
		staff.Name = inPara.MercInfo.UserName          // 员工姓名
		staff.MerchantId = inPara.MercInfo.MerchantId  // 商 家 ID
		staff.Phone = inPara.UserPhone                 // 员工手机号
		staff.UserId = inPara.MercInfo.UserId          // 员 工 ID
		staff.CreateAt = inPara.MercInfo.CreateAt      // 创建时间
		staff.State = 1                       // 状    态
		staff.NumberFage = 1                  // 身份标识
		staff.Authority = 9223372036854775807 // 权    限
		add.PStaff = staff
		add.AreaNumber = inPara.MercInfo.AreaNumber
		fmt.Println("商家入驻管理员信息:", staff)
		err = staff.Add(&add, &staff)
	}
	return err
}

/*
 * desc: 更新商家所有信息
 *
 *************************************************************************************/
func (this *Merchant) Update(inPara, outPara *Merchant) error {
	if inPara.MerchantId != "" {
		_, err := db.GetDBHand(0).Table(inPara.name()).Where("merchant_id = ?", inPara.MerchantId).Update(inPara)
		var val MerchantInfo
		val.Add(inPara)
		return err
	}
	return errors.New("成员属性 MerchantId 不可以为空")
}

/*
 * desc: 更新商家手机号
 *
 *************************************************************************************/
func (this *Merchant) UpdatePhone(inPara, outPara *Merchant) error {
	if inPara.MerchantId != "" {
		_, err := db.GetDBHand(0).Table(inPara.name()).
			Where("merchant_id = ?", inPara.MerchantId).
			Cols("phone").
			Update(inPara)
		outPara = inPara
		return err
	}
	return errors.New("成员属性 MerchantId 不可以为空")
}

func (this *Merchant) BackUpdateStatus(inPara, outPara *Merchant) error {
	if inPara.MerchantId != "" {
		_, err := db.GetDBHand(0).Table(inPara.name()).
			Where("merchant_id = ?", inPara.MerchantId).Cols("checkdesc").Cols("status").Update(inPara)
		return err
	}
	return errors.New("成员属性 MerchantId 不可以为空")
}

/*
 * desc: 更新商家状态,与认证信息
 *
 *************************************************************************************/
func (this *Merchant) UpdateStatus(inPara, outPara *Merchant) error {
	if inPara.MerchantId != "" {
		_, err := db.GetDBHand(0).Table(inPara.name()).
			Where("merchant_id = ?", inPara.MerchantId).
			Cols("checkdesc").
			Cols("status").
			Cols("business").
			Cols("number_id").
			Cols("industry").
			Update(inPara)
		outPara = inPara
		return err
	}
	return errors.New("成员属性 MerchantId 不可以为空")
}

/*
 * desc: 删除商家
 *
 *************************************************************************************/
func (this *Merchant) Delete(inPara, outPara *Merchant) error {
	if inPara.MerchantId != "" {
		var merchant MerchantInfo
		merchant.MerchantId = inPara.MerchantId
		_, err := db.GetDBHand(0).Table(inPara.name()).Where("merchant_id = ?", inPara.MerchantId).Delete(inPara)
		merchant.Delete()
		return err
	}
	return errors.New("成员属性 MerchantId 不可以为空")
}

type MerchantAddBranch struct {
	Superior string `json:"superior"` // 上级ID, 总店
	Lower    string `json:"lower"`    // 下级ID, 本店
}

/*
 * desc: 添加分店
 *
 *************************************************************************************/
func (this *Merchant) AddBranch(inPara *MerchantAddBranch, outPara *Merchant) error {
	if inPara.Superior != "" && inPara.Lower != "" {
		var err error
		var val MerchantInfo
		val.MerchantId = inPara.Superior
		if err = val.AddBranch(inPara.Lower); nil == err {
			outPara.MerchantId = inPara.Lower
			outPara.FID = inPara.Superior
			_, err = db.GetDBHand(0).Table(outPara.name()).
				Where("merchant_id = ?", outPara.MerchantId).
				Cols("fid").
				Update(outPara)
		}
		return err
	}
	return errors.New("成员属性 MerchantId 不可以为空")
}

/*
 * desc: 更新诺数量
 *
 *************************************************************************************/
func (this *Merchant) UpdateTrust(inPara, outPara *Merchant) error {
	if inPara.MerchantId != "" {
		_, err := db.GetDBHand(0).Table(inPara.name()).Where("merchant_id=?", inPara.MerchantId).Cols("trust").Update(inPara)
		return err
	}
	return errors.New("成员属性 MerchantId 不可以为空")
}

/*
 * desc: 获取本所有员工
 *
 *************************************************************************************/
func (this *Merchant) GetStaff(inPara *Merchant, outPara *StaffList) error {
	if inPara.MerchantId != "" {
		var err error
		var val MerchantInfo
		val.MerchantId = inPara.MerchantId
		if err = val.Get(); err == nil {
			strStaffTableName := fmt.Sprintf("chi_staff_%d", val.AreaNumber)
			db.GetDBHand(0).Table(strStaffTableName).Where("merchant_id = ?", val.MerchantId).Find(outPara)
		}
		return err
	}
	return errors.New("成员属性 MerchantId 不可以为空")
}

type CoordinatesPoint struct {
	Longitude float64 `json:"longitude"` //经  度
	Latitude  float64 `json:"latitude"`  //纬  度
	Page      int     `json:"page"`      //页  码
	OfferSet  int     `json:"offer_set"` //偏移量
}

var redClient *redis.Client
var result = make(map[string][]string, 0)
/*
 * desc: 获取附近商家列表
 *
 *************************************************************************************/
func (this *Merchant) GetNearMerchant(inPara *CoordinatesPoint, outPara *MerchantList) error {
	this.lock.Lock()
	defer func() {
		this.lock.Unlock()
	}()
	result = make(map[string][]string, 0)
	radius, err := findGeoRadius(inPara.Longitude, inPara.Latitude)
	if err != nil {
		return err
	}
	start := (inPara.Page - 1) * inPara.OfferSet
	end := start + inPara.OfferSet
	if start >= len(radius) {
		return errors.New("**** result is empty")
	}
	if end > len(radius) {
		end = len(radius)
	}
	radius = radius[start: end] //分页处理
	//表名对应所有的shareid取出
	for _,radiuv:= range radius {
		//数组第一个元素是表名   后面是share_id
		//addRadiueSelice(radius[i][0], radius[i][1])
		if len(radiuv)<2{
			continue
		}
		if len(result[radiuv[0]]) == 0 {
			result[radiuv[0]] = make([]string, 0)
		}
		result[radiuv[0]] = append(result[radiuv[0]], radiuv[1])
	}
	//查询mysql 表名对应的所有shareid 集合存储在result中
	res := make([]*Merchant, 0)
	for key, value := range result {
		m := make([]*Merchant, 0)
		name := fmt.Sprintf("car_merchant_%v", key)
		e := db.GetDBHand(0).Table(name).In("merchant_id", value).Find(&m)
		if e!=nil{
			fmt.Printf("%+v",value)
			continue
		}
		res = append(res, m...)
	}

	//对所有的商家添加距离
	//表名对应所有的shareid取出
	for i := 0; i < len(radius); i++ {
		//数组第一个元素是表名   后面是share_id
		for _, value := range res {
			if strings.EqualFold(radius[i][1], value.MerchantId) {
				value.Distance, _ = strconv.ParseFloat(radius[i][2], 64)
				*outPara = append(*outPara, *value)
				break
			}
		}
	}
	sort.Sort(outPara)
	result = nil
	return nil
}

/*根据表名获得对应的数组*/
func addRadiueSelice(tableName, share string) {
	if len(result[tableName]) == 0 {
		result[tableName] = make([]string, 0)
	}
	result[tableName] = append(result[tableName], share)
	return
}

const (
	Radius = 20 //半径多少
)

/*
	//  表名  shareId  距离(km)
	//[[310   123     0.0001] [311 123 2.4993]]
*/
func findGeoRadius(longitude, latitude float64) ([][3]string, error) {
	radius := &redis.GeoRadiusQuery{
		Radius:   Radius,
		Unit:     "km",
		WithDist: true,  //距离差
		Sort:     "asc", //排序
	}
	locations, err := db.GetRedis().GeoRadius(AREA_LIST, longitude, latitude, radius).Result()
	if err != nil {
		fmt.Println("redis距离错误", err)
		return nil, err
	}
	resv := make([][3]string, 0)
	func(loc []redis.GeoLocation) error {
		for _, value := range loc {
			var s [3]string
			str := strings.Split(value.Name, "_")
			if len(str) != 2 {
				continue
			}
			s[0], s[1] = str[0], str[1]
			s[2] = strconv.FormatFloat(value.Dist, 'f', -1, 64)
			resv = append(resv, s)
		}
		return nil
	}(locations)
	return resv, nil
}

/*
 * desc: 商家交易( 收入 )
 *
 *************************************************************************************/
func (this *Merchant) Trading(inPara, outPara *Merchant) error {
	if inPara.MerchantId != "" {
		_, err := db.GetDBHand(0).Table(inPara.name()).Cols("trust").Update(inPara)
		return err
	}
	return errors.New("成员属性 MerchantId 不可以为空")
}

/*
 * desc: TEST
 *
 *************************************************************************************/
func (this *Merchant) Test(inPara, outPara *Merchant) error {
	RedisParam := redis.GeoRadiusQuery{
		Radius: 20,
		Unit:   "km",
	}
	GeoLocationList, err := db.GetRedis().GeoRadius(AREA_LIST, inPara.Longitude, inPara.Latitude, &RedisParam).Result()
	fmt.Println(GeoLocationList)
	fmt.Println(err)
	return err
}
