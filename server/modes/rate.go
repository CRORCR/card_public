package modes

import (
	"fmt"
	"public/server/db"
)

const YOAWORATE = "yoawo_rate"

/**
 * 行业以及费率模型
 */
type YoawoRate struct {
	Id      int64   `json:"id" xorm:"id"`   //表    ID
	ClasId  int64   `json:"class_id" xorm:"class_id"` //一 类 ID
	LevelId int64   `json:"level_id" xorm:"level_id"` //二 类 ID
	Name    string  `json:"name" xorm:"name"`  //名 称 ID
	Rate1   float64 `json:"rate_1" xorm:"rate_1"`//费    率
	Rate2   float64 `json:"rate_2" xorm:"rate_2"`//费    率
	Rate3   float64 `json:"rate_3" xorm:"rate_3"`//费    率
	Rate4   float64 `json:"rate_4" xorm:"rate_4"`//费    率
	Rate5   float64 `json:"rate_5" xorm:"rate_5"`//费    率
	Rate6   float64 `json:"rate_6" xorm:"rate_6"`//费    率
	Rate7   float64 `json:"rate_7" xorm:"rate_7"`//费    率
}

func (this *YoawoRate)GetOne( inPara, outPara *YoawoRate )error{
	fmt.Println(outPara)
	_, e := db.GetDBHand(0).Table( YOAWORATE ).Where("id=?",inPara.Id).Get( outPara )
	fmt.Println(outPara)
	return e
}

/**
获取当前分类下的子类信息
*/
func (this *YoawoRate)GetList( inPara *YoawoRate, outPara *[]YoawoRate )error{
	e := db.GetDBHand(0).Table( YOAWORATE ).Where("class_id = ?", inPara.ClasId).Find( outPara )
	return e
}

/**
获取当前分类下的子类信息
*/
func (this *YoawoRate)GetTopList( inPara *YoawoRate, outPara *[]YoawoRate )error{
	e := db.GetDBHand(0).Table( YOAWORATE ).Where("level_id= 0 ").Find( outPara )
	return e
}

/**
根据名称搜索
*/
func (this *YoawoRate) GetListByName( inPara *YoawoRate, outPara *[]YoawoRate )error{
	e := db.GetDBHand(0).SQL("SELECT * FROM " + YOAWORATE + " WHERE name like '%" + inPara.Name + "%' ").Find(outPara)
	return e
}

/*
  更新数据
*/
func (this *YoawoRate)Update( inPara, outPara *YoawoRate )error{
	_, err := db.GetDBHand(0).Table( YOAWORATE ).
				  Where("id = ?", inPara.Id ).
				  Update(inPara)
        return err
}