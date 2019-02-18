
package modes

import (
	"public/server/db"
	"errors"
)

const YOAWORATE = "yoawo_rate"
/**
 * 行业以及费率模型
 */
type YoawoRate struct {
	Id      int64   `xorm:"id"`
	ClasId  int64   `xorm:"class_id"`
	LevelId int64   `xorm:"level_id"`
	Name    string  `xorm:"name"`
	Rate1   float64 `xorm:"rate_1"`
	Rate2   float64 `xorm:"rate_2"`
	Rate3   float64 `xorm:"rate_3"`
	Rate4   float64 `xorm:"rate_4"`
	Rate5   float64 `xorm:"rate_5"`
	Rate6   float64 `xorm:"rate_6"`
	Rate7   float64 `xorm:"rate_7"`
}

func (this *YoawoRate)GetOne( inPara, outPara *YoawoRate )error{
	_, e := db.GetDBHand(0).Table( YOAWORATE ).Get( inPara )
	return e
}

/**
获取当前分类下的子类信息
*/
func (this YoawoRate)GetList( inPara *YoawoRate, outPara *[]YoawoRate )error{
	e := db.GetDBHand(0).Table( YOAWORATE ).Where("class_id=?", this.ClasId).Find( outPara )
	return e
}

/**
获取当前分类下的子类信息
*/
func (this YoawoRate)GetTopList( inPara *YoawoRate, outPara *[]YoawoRate )error{
	e := db.GetDBHand(0).Table( YOAWORATE ).Where("level_id= 0 ").Find( outPara )
	return e
}

/**
根据名称搜索
*/
func (this YoawoRate) GetListByName( inPara *YoawoRate, outPara *[]YoawoRate )error{
	e := db.GetDBHand(0).SQL("SELECT * FROM " + YOAWORATE + " WHERE name like '%" + this.Name + "%' ").Find(outPara)
	return e
}

/*
  更新数据
*/
func (this YoawoRate)U( inPara, outPara *YoawoRate )error{
	var err error
	if inPara.Rate1 > 6 ||
	   inPara.Rate2 > 6 ||
	   inPara.Rate3 > 6 ||
	   inPara.Rate4 > 6 ||
	   inPara.Rate5 > 6 ||
	   inPara.Rate6 > 6 ||
	   inPara.Rate7 > 6 {
	   _, err = db.GetDBHand(0).Where("class_id = ? AND level_id = ?", inPara.ClasId, inPara.LevelId ).
			        Cols("rate_1").
			        Cols("rate_2").
			        Cols("rate_3").
			        Cols("rate_4").
			        Cols("rate_5").
			        Cols("rate_6").
			        Cols("rate_7").
				Update(inPara)
	}else{
		return errors.New("请查看费率，不可以低于6")
	}
        return err
}




