package modes

import (
	"card_public/lib"
	"card_public/server/db"
	"fmt"
	"time"
)

/*从mysql分组(县)拿出所有县
拿出所有县的所有广告位(是否存在已上架,存在,跳过,不存在,是否存在等待中,存在,设置),
*/
var t = time.NewTimer(time.Second * 1)

func ReviewBanner() {
	go reviewRate()
}

func reviewRate() {
	fmt.Println("当前时间", lib.TimeToString(time.Now().Unix()))
	review()
	for {
		select {
		case <-t.C:
			t = time.NewTimer(time.Second * 60)
			review()
		default:
			time.Sleep(time.Second * 30)
		}
	}
}

func review() {
	var result = make(chan *Banner, 1)
	var clost = make(chan struct{})
	defer func(){
		close(clost)
		close(result)
	}()
	var areaList []int64
	fmt.Println("区域",areaList)
	db.GetDBHand(0).Table(BANNERTABLE).Cols("area_id").GroupBy("area_id").Find(&areaList)
	go upload(result, clost)
	//县
	for _, areaId := range areaList {
		var siteList []string
		db.GetDBHand(0).Table(BANNERTABLE).Cols("banner_site").Where("area_id=?",areaId).GroupBy("banner_site").Find(&siteList)
		//广告位
		for _, site := range siteList {
			ban := &Banner{AreaId: areaId, BannerSite: site, BannerStatus: 2}
			b, _ := db.GetDBHand(0).Table(BANNERTABLE).Get(ban)
			if b {
				continue
			}
			ban.BannerStatus = 1
			b, _ = db.GetDBHand(0).Table(BANNERTABLE).Asc("pay_time").Get(ban)
			if b {
				fmt.Println("一个任务")
				result <- ban
			}
		}
	}
	clost <- struct{}{}
}

func upload(banner chan *Banner, clost chan struct{}) {
	select {
	case <-clost:
		return
	case ban := <-banner:
		if err:=ban.upShow();err!=nil{
			fmt.Println("上传有误:",err)
		}
		fmt.Println("上架:success")
	}
}
