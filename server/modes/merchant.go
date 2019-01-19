package modes

import (
	"public/server/db"
	"strconv"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
//	"strings"
)

const MERCHANT  = "MERCHANT_"
const AREA_LIST = "AREA_LIST"

type MerchantInfo struct {
	MerchantId string	// 商家ID
	UserId	   string	// 用户ID
	AreaNumber int64	// 商家所在区的ID
}

func( this *MerchantInfo )Name()string{
	return fmt.Sprintf( "%s%s",MERCHANT, this.MerchantId )
}

/*
 * 描述: 获取本商家的区号
 *
 *************************************************************/
func( this *MerchantInfo )GetAreaNumber()error{
	var err error
	sCmd := db.GetRedis().HGet(this.Name(), "AreaNumber")
	this.AreaNumber, err = sCmd.Int64()
	fmt.Print("%+v\n",sCmd)
	fmt.Println("AreaNumber", this.AreaNumber, "Error:", err)
	return err
}

/*
 * 描述: 添加商家
 *
 *************************************************************/
func( this *MerchantInfo )Add( fLongitude, fLatitude float64 ) error {
	client := db.GetRedis()
        strName := this.Name()
        client.HSet( strName, "MerchantId", this.MerchantId )
        client.HSet( strName, "AreaNumber", this.AreaNumber )
        client.HSet( strName, "UserId",     this.UserId )

	strKey := fmt.Sprintf( "%d_%s", this.AreaNumber, this.MerchantId )
	geol := &redis.GeoLocation{Name:strKey, Longitude:fLongitude, Latitude:fLatitude }
	_, err := client.GeoAdd( AREA_LIST, geol ).Result()
	return err

}

/*
 * 描述: 添加商家
 *
 *************************************************************/
func( this *MerchantInfo )Get() error {
        client := db.GetRedis()
        sKey, sErr := client.HGetAll( this.Name() ).Result()
        if nil == sErr {
                this.UserId, _       = sKey["UserId"]
                this.MerchantId, _   = sKey["MerchantId"]
                this.AreaNumber, _   = strconv.ParseInt( sKey["AreaNumber"] , 10 , 64 )
        }
        return sErr
}


/*
 * 描述: 商家信息表
 *
 *  type_id     : 商家所属行业ID
 *  status      : 1 认证中 2 认证未通过 3 认证通过， 4 删除
 */
type Merchant struct {
	Id           string  `json:"id" xorm:"id"`                       //商家表ID,
	FID          string  `json:"fid" xorm:"fid"`                     //父商家id
	MerchantId   string  `json:"merchant_id" xorm:"merchant_id"`     //本商家Id
	UserId       string  `json:"user_id" xorm:"user_id"`             //用户分享ID,
	MerchantType string  `json:"merchant_type" xorm:"merchant_type"` //商家 行业 类型
	MerchantRate float64 `json:"merchant_rate" xorm:"-"`             //商家 行业 利率
	TrustStatus  bool    `json:"trust_status" xorm:"trust_status"`   //是否诺商家, 0 否 1 是
	AreaNumber   int64   `json:"area_number" xorm:"area_number"`     //地 区 I   D,
	CreateAt     int64   `josn:"-" xorm:"create_at"`                 //创 建 时 间,
	CreateAtStr  string  `josn:"create_at" xorm:"-"`                 //创 建 时 间,
	Describea    string  `josn:"describea" xorm:"describea"`         //描       述,
	Address      string  `josn:"address" xorm:"address"`             //地       址,
	UserName     string  `josn:"name" xorm:"name"`                   //名       称,
	Status       int64   `josn:"status" xorm:"status"`               //状       态, 未认证过的 0  已经审核通过 1
	Phone        string  `josn:"phone" xorm:"phone"`                 //手  机   号,
	Icon         string  `josn:"icon" xorm:"icon"`                   //商 家 头 像,
	LoopImg      string  `josn:"loopimg" xorm:"loopimg"`             //商家轮播图,
	InfoImg      string  `josn:"infoimg" xorm:"infoimg"`             //商家详情图,
	Video        string  `josn:"video" xorm:"video"`                 //商家视频介绍,
	CheckDesc    string  `josn:"checkdesc" xorm:"checkdesc"`         //认证未失败描述
	CheckImg     string  `josn:"checkimg" xorm:"checkimg"`           //认       证
	Longitude    float64 `josn:"longitude" xorm:"longitude"`         //经       度
	Latitude     float64 `josn:"latitude" xorm:"latitude"`           //纬       度
	Cash         float64 `josn:"cash" xorm:"cash"`                   //现       金
	Trust        float64 `josn:"trust" xorm:"trust"`                 //鍩       分
	Credits      float64 `josn:"credits" xorm:"credits"`             //积       分
	Distance     float64 `json:"distance" xorm:"-"`		         //商家与用户的距离
}

func ( this *Merchant )name()string{
	var val MerchantInfo
	val.MerchantId = this.MerchantId
	fmt.Println("this.MerchantId", this.MerchantId)
	val.GetAreaNumber()
	return fmt.Sprintf("car_merchant_%d", val.AreaNumber )
}

/*
 * desc: 获取商家
 * 
 *************************************************************************************/
func (this *Merchant)Get( inPara, outPara *Merchant )error{
	outPara = inPara
	_, err := db.GetDBHand(0).Table( inPara.name() ).Get( outPara )
	return err
}

/*
 * desc: 添加商家
 * 
 *************************************************************************************/
func (this *Merchant)Add( inPara, outPara *Merchant )error{
	var val MerchantInfo
	val.MerchantId = inPara.MerchantId  // 商家ID
        val.UserId     = inPara.UserId      // 用户ID
        val.AreaNumber = inPara.AreaNumber  // 商家所在区的ID
	val.Add( inPara.Longitude, inPara.Latitude )
	outPara = inPara
	_, err := db.GetDBHand(0).Table( inPara.name() ).Insert( outPara )
	return err
}

/*
 * desc: 更新商家状态
 * 
 *************************************************************************************/
func (this *Merchant)UpdateStatus( inPara, outPara *Merchant )error{
	if inPara.MerchantId != "" {
		_, err := db.GetDBHand(0).Table( inPara.name() ).Cols("checkdesc").Cols("status").Update( inPara )
		outPara = inPara
		return err
	}
        return errors.New("成员属性 MerchantId 不可以为空")
}

/*
 * desc: 更新诺数量
 *
 *************************************************************************************/
func (this *Merchant)UpdateTrust( inPara, outPara *Merchant )error{
        if inPara.MerchantId != "" {
                _, err := db.GetDBHand(0).Table( inPara.name() ).Cols("trust").Update( inPara )
                return err
        }
        return errors.New("成员属性 MerchantId 不可以为空")
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
