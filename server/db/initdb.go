package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"../../lib"
)

var g_dbHand []*xorm.Engine

func GetDBHand( nIndex int ) *xorm.Engine {
    if len(g_dbHand) >= nIndex{
	    return g_dbHand[nIndex]
    }
    return nil
}

func InitDB() error {

    dblist := lib.ReadDBConfig("./config/db.json")
    for _,db := range dblist{
        strConn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", db.DBUser,
                                                      db.DBPass,
                                                      db.DBHome,
                                                      db.DBPort,
                                                      db.DBName)
	dbHand, err := xorm.NewEngine("mysql", strConn )
	if err != nil {
		fmt.Println("Web database link failed", err)
		return err
	}
	if err := dbHand.Ping(); err != nil {
		fmt.Println("Test link failed",err)
		return err
	}
	dbHand.SetTableMapper(core.SameMapper{})
	dbHand.SetColumnMapper(core.SameMapper{})
	dbHand.ShowSQL(true)
	dbHand.SetMaxIdleConns(5)
	dbHand.SetMaxOpenConns(5)
	g_dbHand = append(g_dbHand, dbHand)
    }
    fmt.Println("DB:", dblist)
    return nil
}

