package modes

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"public/server/db"

	"github.com/go-redis/redis"

	//	"strings"
)

const MERCHANT = "MERCHANT_"
const AREA_LIST = "AREA_LIST"

type MerchantInfo struct {
	MerchantId string // 商家ID
	UserId     string // 用户ID
	AreaNumber int64  // 商家所在区的ID
}

func (this *MerchantInfo) Name() string {
	return fmt.Sprintf("%s%s", MERCHANT, this.MerchantId)
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
 * 描述: 添加商家
 *
 *************************************************************/
func (this *MerchantInfo) Add(fLongitude, fLatitude float64) error {
	client := db.GetRedis()
	strName := this.Name()
	client.HSet(strName, "MerchantId", this.MerchantId)
	client.HSet(strName, "AreaNumber", this.AreaNumber)
	client.HSet(strName, "UserId", this.UserId)

	strKey := fmt.Sprintf("%d_%s", this.AreaNumber, this.MerchantId)
	geol := &redis.GeoLocation{Name: strKey, Longitude: fLongitude, Latitude: fLatitude}
	_, err := client.GeoAdd(AREA_LIST, geol).Result()
	return err
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
	}
	return sErr
}

//*****************以下是提供外部rpc调用**************************************
/*
 * 描述: 商家信息表
 *
 *  type_id     : 商家所属行业ID
 *  status      : 1 认证中 2 认证未通过 3 认证通过  4 删除
 */
type Merchant struct {
	Id           int     `json:"id" xorm:"id"`                       //商家表ID,
	FID          string  `json:"fid" xorm:"fid"`                     //父商家id
	MerchantId   string  `json:"merchant_id" xorm:"merchant_id"`     //本商家Id
	UserId       string  `json:"user_id" xorm:"user_id"`             //用户分享ID,
	MerchantType int64   `json:"merchant_type" xorm:"merchant_type"` //商家 行业 类型
	MerchantRate float64 `json:"merchant_rate" xorm:"-"`             //商家 行业 利率
	TrustStatus  bool    `json:"trust_status" xorm:"trust_status"`   //是否诺商家, 0 否 1 是
	AreaNumber   int64   `json:"area_number" xorm:"area_number"`     //地 区 I   D,
	CreateAt     int64   `json:"-" xorm:"create_at"`                 //创 建 时 间,
	CreateAtStr  string  `json:"create_at" xorm:"-"`                 //创 建 时 间,
	Describea    string  `json:"describea" xorm:"describea"`         //描       述,
	Address      string  `json:"address" xorm:"address"`             //地       址,
	UserName     string  `json:"name" xorm:"name"`                   //名       称,
	Status       int64   `json:"status" xorm:"status"`               //状       态, 未认证过的 0  已经审核通过 1
	Phone        string  `json:"phone" xorm:"phone"`                 //手  机   号,
	Icon         string  `json:"icon" xorm:"icon"`                   //商 家 头 像,
	LoopImg      string  `json:"loopimg" xorm:"loopimg"`             //商家轮播图,
	InfoImg      string  `json:"infoimg" xorm:"infoimg"`             //商家详情图,
	Video        string  `json:"video" xorm:"video"`                 //商家视频介绍,
	CheckDesc    string  `json:"checkdesc" xorm:"checkdesc"`         //认证未失败描述
	CheckImg     string  `json:"checkimg" xorm:"checkimg"`           //认       证
	Longitude    float64 `json:"longitude" xorm:"longitude"`         //经       度
	Latitude     float64 `json:"latitude" xorm:"latitude"`           //纬       度
	Cash         float64 `json:"cash" xorm:"cash"`                   //现       金
	Trust        float64 `json:"trust" xorm:"trust"`                 //鍩       分
	Credits      float64 `json:"credits" xorm:"credits"`             //积       分
	Distance     float64 `json:"distance" xorm:"-"`                  //商家与用户的距离
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

/*
 * desc: 添加商家
 * 
 ***************/
func (this *Merchant) Add(inPara, outPara *Merchant) error {
	var val MerchantInfo
	val.MerchantId = inPara.MerchantId // 商家ID
	val.UserId = inPara.UserId         // 用户ID
	val.AreaNumber = inPara.AreaNumber // 商家所在区的ID
	val.Add(inPara.Longitude, inPara.Latitude)
	outPara = inPara
	_, err := db.GetDBHand(0).Table(inPara.name()).Insert(outPara)
	return err
}

func (this *Merchant)FindBranch(inPara *Merchant, outPara *MerchantList)error{
	err := db.GetDBHand(0).Table(inPara.name()).Where("merchant_id=? ",inPara.FID).Find(outPara)
	fmt.Println("err:",err)
	return err
}

/*
 * desc: 更新商家状态
 * 
 **************************************************************************************/
func (this *Merchant) UpdateStatus(inPara, outPara *Merchant) error {
	if inPara.MerchantId != "" {
		_, err := db.GetDBHand(0).Table(inPara.name()).Where("merchant_id=?", inPara.MerchantId).Cols("checkdesc").Cols("status").Update(inPara)
		outPara = inPara
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
	Longitude float64 //经       度
	Latitude  float64 //纬       度
	Page      int
	OfferSet  int
}

var redClient *redis.Client

/*
 * desc: 获取附近商家列表
 *
 *************************************************************************************/
func (this *Merchant) GetNearMerchant(inPara *CoordinatesPoint, outPara *MerchantList) error {
	radius, err := findGeoRadius(inPara.Longitude, inPara.Latitude)
	if err != nil {
		return err
	}
	fmt.Println("分页:", inPara.Page, inPara.OfferSet)

	radius = radius[(inPara.Page-1)*inPara.OfferSet:(inPara.Page-1)*inPara.OfferSet+inPara.OfferSet] //分页处理
	//表名对应所有的shareid取出
	for i := 0; i < len(radius); i++ {
		//数组第一个元素是表名   后面是share_id
		addRadiueSelice(radius[i][0], radius[i][1])
	}

	//查询mysql 表名对应的所有shareid 集合存储在result中
	m := make(MerchantList, 0)
	res := make(MerchantList, 0)
	for key, value := range result {
		name := fmt.Sprintf("car_merchant_%v", key)
		e := db.GetDBHand(0).Table(name).In("merchant_id", value).Find(&m)
		fmt.Printf("err:%+v  result:%+v \n", e, m)
		res = append(res, m...)
		//*outPara=append(*outPara,m...)
	}

	//对所有的商家添加距离
	//*outPara=append(*outPara,)
	//表名对应所有的shareid取出
	for i := 0; i < len(radius); i++ {
		//数组第一个元素是表名   后面是share_id
		for _, value := range res {
			if strings.EqualFold(radius[i][1], value.MerchantId) {
				value.Distance, _ = strconv.ParseFloat(radius[i][2], 64)
				*outPara = append(*outPara, value)
				break
			}
		}
	}
	sort.Sort(outPara)
	return nil
}

var result = make(map[string][]string, 0)

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
func findGeoRadius(longitude, latitude float64) (result [][3]string, err error) {
	radius := &redis.GeoRadiusQuery{
		Radius:   Radius,
		Unit:     "km",
		WithDist: true,  //距离差
		Sort:     "asc", //排序
	}

	locations, err := db.GetRedis().GeoRadius(AREA_LIST, longitude, latitude, radius).Result()
	if err != nil {
		return nil, err
	}
	result = make([][3]string, 0)
	func(loc []redis.GeoLocation) error {
		for _, value := range loc {
			var s [3]string
			str := strings.Split(value.Name, "_")
			if len(str) != 2 {
				continue
			}
			s[0], s[1] = str[0], str[1]
			s[2] = strconv.FormatFloat(value.Dist, 'f', -1, 64)
			result = append(result, s)
		}
		return nil
	}(locations)
	return
}

/*
 * desc: 商家管理首页  查询所有审核状态为空的商家  0:未审核  1:审核通过  2:微名片添加过的
 * @create: 2018/11/22
 *
func (this *Merchant) GetMerchants() ([]*Merchant, error) {
	//自己店铺
	merchants := make([]*Merchant, 0)
	err := db.GetDBHand(0).Table(lib.TableNameMerchant).Where("phone =?", this.Phone).Desc("create_at").Find(&merchants)
	if err != nil {
		return nil, err
	}
	//下级店铺
	merchantSon := make([]*Merchant, 0)
	for _, value := range merchants {
		err := db.GetDBHand(0).Table(lib.TableNameMerchant).Where("fid =?", value.Id).Desc("create_at").Find(&merchantSon)
		if err != nil {
			return nil, err
		}
	}
	merchants = append(merchants, merchantSon...)
	return merchants, nil
}

微名片 新增,编辑
func (this *Merchant) UpdateMerchant() error {
	engine := db.GetDBHand(0).Table(lib.TableNameMerchant)
	m := &Merchant{}
	b, err := engine.Where("phone=? and id=?", this.Phone, this.Id).Get(m)
	if err != nil {
		return err
	}
	if !b {
		_, err = engine.Table(lib.TableNameMerchant).Insert(this)
		return err
	}
	if (this.FID == "" && m.FID != "") || (this.SID == "" && m.SID != "") {
		return errors.New("fid or sid only submit.")
	}

	_, err = engine.Table(lib.TableNameMerchant).Update(this)
	return err
}

func (this *Merchant) UpdateMerchantStatus(queryName, queryId string) error {
	//如果同意,status改为1 认证通过
	var err error
	sql := fmt.Sprintf("update %v set %v = ? where id = ?", lib.TableNameMerchant, "status")
	engine := db.GetDBHand(0).Table(lib.TableNameMerchant) //获得test数据库
	if strings.EqualFold(queryName, "true") { //同意
		_, err = engine.Exec(sql, 1, queryId) //认证通过
	} else {
		_, err = engine.Exec(sql, 0, queryId) //认证失败
	}
	if err != nil {
		return err
	}
	return nil
}

func (this *Merchant) GetContainMerchants(start, end int64) ([]*Merchant, error) {
	merchant := make([]*Merchant, 0)
	engine := db.GetDBHand(0).Table(lib.TableNameMerchant) //获得test数据库
	err := engine.Where("create_at>= ? and create_at< ?", start, end).Desc("create_at").Find(&merchant)
	if err != nil {
		return nil, err
	}
	return merchant, nil
}

func (this *Merchant) GetMerchantByPhoneNumOrName(str string) []*Merchant {
	merchant := make([]*Merchant, 0)
	engine := db.GetDBHand(0).Table(lib.TableNameMerchant) //获得test数据库
	check := lib.IsPhone(str)
	if check {
		engine.Where("phone=?", str).Desc("create_at").Find(&merchant)
	} else {
		engine.Where("name=?", str).Desc("create_at").Find(&merchant)
	}
	return merchant
}

 *
 * desc:  获得id对应的用户 share_id  然后关联商家表,返回所有商家对应信息
 * @create: 2018/11/23
 *
func (this *Merchant) GetMerchant(share_id string) []*Merchant {
	//sql := `select * from car_merchant c where c.id in (select b.cid from user_relation b  where b.uid=?)`
	merchant := make([]*Merchant, 0)
	db.GetDBHand(0).Table(lib.TableNameMerchant).Where("user_id =?", share_id).Find(&merchant)
	return merchant
}

func (this *Merchant) GetMerchantByCardId() {
	db.GetDBHand(0).Table(lib.TableNameMerchant).Get(this)
	return
}
*/
