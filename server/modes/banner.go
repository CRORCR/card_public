package modes

//对于历史数据,只能查询,不能编辑,如果想根据历史数据重新生成,只能拿到历史数据,重新形成一个广告
//父id  表id    查询广告位,最后按照时间排序,最后一个,作为自己父id(其实可以在上架的时候检查,
// 但是为了后期可能需要判断自己第几次上架,还是多加一步)

import (
	"card_public/lib"
	"card_public/server/db"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"sync"
	"time"
)

/**
 * @desc    模板 五个广告位,增量
 * @author Ipencil
 * @create 2019/3/7
 */
type BannerShow struct {
	AreaID     int64 `json:"area_id"`     //县id
	Site       string `json:"site"`        //位置 轮播1,轮播2,轮播3,轮播4,轮播5
	Count      int64 `json:"count"`       //当前一共需要点击多少次
	UnixTime   int64 `json:"-"`           // 时间标志
	TodayTimes int64 `json:"today_times"` //今日点击次数     每日递增
	TickOuts   int64 `json:"tick_outs"`   //累计点击次数     每次递增
	Remains    int64 `json:"remains"`     //剩余点击次数     总-累计   次数(原子)
	TotalTimes int64 `json:"total_times"` //总共点击次数    通过广告位+价格
}

//多少广告 目前是5个
var bannerSum = [5]string{"轮播1", "轮播2", "轮播3", "轮播4", "轮播5"}

var BannerShowList []BannerShow

const (
	BANNERTEMPLIST = "banner_list" //正在展示的本县广告, 县areaID_展示位
	BANNERTABLE    = "banner"
)

type Banner struct {
	ID           int64  `json:"id" xorm:"id"`
	AreaId       int64  `json:"area_id" xorm:"area_id"`             //县id
	MerchantID   string `json:"merchant_id" xorm:"merchant_id"`     //商家id
	DadID        int64  `json:"dad" xorm:"dad"`                     //表id
	BannerSite   string `json:"banner_site" xorm:"banner_site"`     //广告位 (轮播1,轮播2,轮播3,轮播4...)
	BannerPrice  int64  `json:"banner_price" xorm:"banner_price"`   //广告位价格(100,200,500)    计算出总次数
	TodayTimes   int64  `json:"today_times" xorm:"today_times"`     //今日点击次数     每日递增
	TickOuts     int64  `json:"tick_outs" xorm:"tick_outs"`         //累计点击次数     每次递增
	Remains      int64  `json:"remains" xorm:"remains"`             //剩余点击次数     总-累计   次数(原子)
	TotalTimes   int64  `json:"total_times" xorm:"total_times"`     //总共点击次数    通过广告位+价格
	PayTime      int64  `json:"pay_time" xorm:"pay_time"`           //付款时间
	PayTimes     string `json:"pay_times" xorm:"-"`                 //付款时间
	ShowTime     int64  `json:"-" xorm:"show_time"`                 //上架时间
	ShowTimes    string `json:"show_time" xorm:"-"`                 //上架时间 string
	BannerStatus int64  `json:"banner_status" xorm:"banner_status"` //状态  1:等待中 2:上架中 3:已下架 4:删除
	BannerUrl    string `json:"banner_url" xorm:"banner_url"`       //图片 地址
	Title        string `json:"title" xorm:"title"`                 //标题
	TitleSec     string `json:"title_sec" xorm:"title_sec"`         //副标题
	Content      string `json:"content" xorm:"content"`             //内容
	ShowEnd      int64  `json:"-" xorm:"show_end"`                  //结束时间
	ShowEnds     string `json:"show_end" xorm:"-"`                  //结束时间
}

/*任务列表提取设置成已上架 -- 广告展示*/
func (this *Banner) upShow() error {
	this.ShowTime = time.Now().Unix()
	this.BannerStatus = 2
	_, err := db.GetDBHand(0).Table(BANNERTABLE).Where("id=?", this.ID).Update(this)
	if err != nil {
		return fmt.Errorf("上架失败,err:%v", err.Error())
	}
	//2.数据提交到线上  先从线上拿下来,然后进行替换
	var key = fmt.Sprintf("%v_%v", this.AreaId, this.BannerSite)

	db.GetRedis().SAdd(BANNERTEMPLIST, key)
	db.GetRedis().HSet(key, "unix_time", time.Now().Unix())
	db.GetRedis().HSet(key, "area_id", this.AreaId)
	db.GetRedis().HSet(key, "site", this.BannerSite)
	db.GetRedis().HSet(key, "today_times", 0)           //今日点击次数     每日递增
	db.GetRedis().HSet(key, "tick_outs", 0)             //累计点击次数     每次递增
	db.GetRedis().HSet(key, "remains", this.TotalTimes) //剩余点击次数     总-累计   次数(原子)
	db.GetRedis().HSet(key, "total_times", this.TotalTimes)
	return nil
}

/*查看商家广告*/
func (this *Banner) DownShow(input, output *Banner) error {
	return downShow(input.AreaId, input.BannerSite)
}

//下架  每次递减,如果变成0,需要再次选出一个广告
func downShow(areaId int64, site string) error {
	var key = fmt.Sprintf("%v_%v", areaId, site)
	npids, _ := db.GetRedis().HGetAll(key).Result()
	if len(npids) == 0 {
		return fmt.Errorf("查看上架广告递减失败")
	}
	i, _ := strconv.ParseInt(npids["remains"], 10, 64)
	unix_time, _ := strconv.ParseInt(npids["unix_time"], 10, 64)
	db.GetRedis().HIncrBy(key, "count", -1)
	//今日时间戳满足,加1,否则置为1,清空今日时间戳
	if lib.IsToday(unix_time) {
		db.GetRedis().HIncrBy(key, "today_times", 1)
	} else {
		db.GetRedis().HSet(key, "unix_time", time.Now().Unix())
		db.GetRedis().HSet(key, "today_times", 1)
	}
	db.GetRedis().HIncrBy(key, "tick_outs", 1)
	db.GetRedis().HIncrBy(key, "remains", -1)
	if i <= 1 { //当前广告无效,选举下一个广告
		//指定县,指定公告位置,指定公告状态为上架中
		area, _ := strconv.ParseInt(npids["area_id"], 10, 64)
		ban := &Banner{BannerSite: npids["site"], AreaId: area, BannerStatus: 2}
		db.GetDBHand(0).Table(BANNERTABLE).Get(ban)
		//修改为已下架
		ban.BannerStatus = 3
		ban.ShowEnd = time.Now().Unix() //修改结束时间
		//下架之后,更新mysql 次数信息
		ban.TodayTimes, _ = strconv.ParseInt(npids["today_times"], 10, 64) //今日点击次数     每日递增
		ban.TodayTimes+=1
		ban.TickOuts, _ = strconv.ParseInt(npids["tick_outs"], 10, 64)     //累计点击次数     每次递增
		ban.TickOuts+=1
		ban.Remains, _ = strconv.ParseInt(npids["remains"], 10, 64)        //剩余点击次数     总-累计   次数(原子)
		ban.Remains-=1
		db.GetDBHand(0).Table(BANNERTABLE).Where("id=?", ban.ID).Update(ban)
		//获取下一个需要展示的数据
		ba := &Banner{DadID: ban.ID}
		db.GetDBHand(0).Table(BANNERTABLE).Get(ba)
		ba.upShow() //提交redis
	}
	return nil
}

//************************************************
//创建一条广告,先存储数据库,更新redis
func (this *Banner) AddBanner(banner, outPut *Banner) error {
	var ban = &Banner{BannerSite: banner.BannerSite, AreaId: banner.AreaId}
	db.GetDBHand(0).Table(BANNERTABLE).Desc("pay_time").Get(ban)
	banner.DadID = ban.ID
	_, err := db.GetDBHand(0).Table(BANNERTABLE).Insert(banner)
	if err != nil {
		return fmt.Errorf("创建广告失败:err:%v", err)
	}
	//更新总数
	var key = fmt.Sprintf("%v_%v", banner.AreaId, banner.BannerSite)
	db.GetRedis().HIncrBy(key, "count", banner.TotalTimes)
	return err
}

type Where struct {
	SQL    string
	OffSet int
	Sum    int
}

type ResultBanner struct {
	BannerResultList []*Banner
	Total            int64
	Error            error
}

/**
 * @desc   : 根据条件查询历史记录(不展示当天,剩余次数)
 * @author : Ipencil
 * @date   : 2019/3/8
 */
//查询本店所有广告,待上架,上架中,已下架 也走这个查询
func (this *Banner) FindBanner(where *Where, outPut *ResultBanner)error {
	ban := make([]*Banner, 0)
	s := new(Banner)
	db.GetDBHand(0).Table(BANNERTABLE).Where(where.SQL).Limit(where.Sum, where.OffSet).Desc("pay_time").Iterate(s, func(idx int, bean interface{}) error {
		value := bean.(*Banner)
		value.PayTimes = lib.TimeToString(value.PayTime)
		if value.ShowTime!=0{
			value.ShowTimes = lib.TimeToString(value.ShowTime)
		}
		if value.ShowEnd!=0{
			value.ShowEnds = lib.TimeToString(value.ShowEnd)
		}
		ban = append(ban, value)
		return nil
	})
	total, err := db.GetDBHand(0).Table(BANNERTABLE).Where(where.SQL).Count(s)
	outPut.Error = err
	outPut.Total = total
	fmt.Println("total",total)
	outPut.BannerResultList = rec(ban)
	return nil
}

func rec(ban []*Banner) []*Banner {
	for i := 0; i < len(ban); i++ {
		switch ban[i].BannerStatus {
		case 2: //上架中 去redis查看
			var key = fmt.Sprintf("%v_%v", ban[i].AreaId, ban[i].BannerSite)
			npids, _ := db.GetRedis().HGetAll(key).Result()
			ban[i].TodayTimes, _ = strconv.ParseInt(npids["today_times"], 10, 64) //今日点击次数     每日递增
			ban[i].TickOuts, _ = strconv.ParseInt(npids["tick_outs"], 10, 64)     //累计点击次数     每次递增
			ban[i].Remains, _ = strconv.ParseInt(npids["remains"], 10, 64)        //剩余点击次数     总-累计   次数(原子)
		case 3: //状态为3 查询mysql,需要检查今日时间
			if !lib.IsToday(ban[i].ShowEnd) {
				ban[i].TodayTimes = 0
			}
		default:
			continue
		}
	}
	return ban
}

/*更新广告*/
func (this *Banner) UpdateBanner(input, outPut *Banner) error {
	i, err := db.GetDBHand(0).Table(BANNERTABLE).Where("banner_status<3 and id=?", input.ID).Update(input)
	if err!=nil{
		return err
	}
	if i==0{
		return fmt.Errorf("更新失败,不得更新历史订单")
	}
	return nil
}

/** todo test
 * @desc   : 查询指定县 指定广告  县必须指定,广告位可以随意
 * @author : Ipencil
 * @date   : 2019/3/8
 */
func (this *Banner) QueryBannerShowInfo(banner *BannerShow, bannerSite *[]BannerShow)error{
	if banner.AreaID==0{
		return fmt.Errorf("区域必须指定")
	}
	var siteList = make([]string, 0)
	if len(banner.Site) == 0 {
		var list []string
		db.GetDBHand(0).Table(BANNERTABLE).Cols("banner_site").Where("area_id=?",banner.AreaID).GroupBy("banner_site").Find(&list)
		siteList = append(siteList, list...)
	} else {
		siteList = append(siteList, banner.Site)
	}
	fmt.Println("集合都是啥",siteList)
	for _, value := range siteList {
		var key = fmt.Sprintf("%v_%v", banner.AreaID, value)
		if isMember, _ := db.GetRedis().SIsMember(BANNERTEMPLIST, key).Result(); isMember {
			fmt.Println("广告存在记录查询参数")
			npids, _ := db.GetRedis().HGetAll(key).Result()
			var ban = BannerShow{}
			ban.AreaID, _ = strconv.ParseInt(npids["area_id"], 10, 64)
			ban.Site = npids["site"]
			ban.Count, _ = strconv.ParseInt(npids["count"], 10, 64)
			ban.TodayTimes, _ = strconv.ParseInt(npids["today_times"], 10, 64)
			ban.TickOuts, _ = strconv.ParseInt(npids["tick_outs"], 10, 64)
			ban.Remains, _ = strconv.ParseInt(npids["remains"], 10, 64)
			ban.TotalTimes, _ = strconv.ParseInt(npids["total_times"], 10, 64)
			*bannerSite = append(*bannerSite, ban)
		}
	}
	return nil
}

/**
 * @desc   : 模板控制
 * @author : Ipencil
 * @date   : 2019/3/7
 */
type TemplateBanner struct {
	Name string      `json:"name"` //模板名称
	Pri  []TempPrice `json:"pri"`
	lock sync.Mutex  `json:"-"`
}

type TempPrice struct {
	Price int64 `json:"price"` //价格
	Count int64 `json:"count"` //次数
}

type Temp struct {
	Url   string `json:"url"` //图片地址
	Temps []TemplateBanner
}

var Template Temp

var fileName = "./config/banner_temp.json"

func tempInit(strFileName string) error {
	Template = Temp{Temps: make([]TemplateBanner, 0)}
	jsonFile, err := os.Open(strFileName)
	if err != nil {
		panic("打开文件错误，请查看:" + strFileName)
	}
	defer jsonFile.Close()

	jsonData, era := ioutil.ReadAll(jsonFile)
	if era != nil {
		panic("读取文件错误:" + strFileName)
	}
	json.Unmarshal(jsonData, &Template)
	return era
}

func (this *TemplateBanner) Get(temp2, temp *Temp) error {
	err := tempInit(fileName)
	if err != nil {
		return err
	}
	*temp = Template
	return nil
}

func (this *TemplateBanner) Set(tb, tb2 *Temp) error {
	this.lock.Lock()
	defer func() { this.lock.Unlock() }()
	if len(Template.Temps) == 0 {
		err := tempInit(fileName)
		if err != nil {
			return err
		}
	}
	Template = *tb
	buff, err := json.MarshalIndent(Template, "", " ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(fileName, buff, 0644)
	return err
}
