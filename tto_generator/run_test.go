/**
 * Copyright 2015 @ at3.net.
 * name : orm_test
 * author : jarryliu
 * date : 2016-11-11 15:26
 * description :
 * history :
 */
package tool

import (
	"github.com/ixre/gof/db"
	"github.com/ixre/gof/db/orm"
	"github.com/ixre/gof/shell"
	"github.com/ixre/tto"
	"os"
	"testing"
)

var (
	//connString = "root:123456@tcp(127.0.0.1:3306)/gcy_v3?charset=utf8"
	driver     = "postgresql"
	dbName     = ""
	dbPrefix   = "mch_express"
	connString = "postgres://postgres:123456@127.0.0.1:5432/go2o?sslmode=disable"
	genDir     = "output/"
)

// 生成数据库所有的代码文件
func TestGenAll(t *testing.T) {
	connString = "postgres://postgres:123456@www.dev.bitsxx.com:5432/go2o?sslmode=disable"
	// 初始化生成器
	conn, _ := db.NewConnector(driver, connString, nil, false)
	dialect := getDialect(driver)
	ds := orm.DialectSession(conn.Raw(), dialect)
	dg := tto.DBCodeGenerator()
	dg.IdUpper = false
	// 获取表格并转换
	raws, _ := ds.TablesByPrefix(dbName, "", dbPrefix)
	tables, err := dg.Parses(raws, false)
	if err != nil {
		t.Error(err)
		return
	}
	// 设置包名
	dg.Var(tto.PKG, "go2o/core")
	// 清理上次生成的代码
	os.RemoveAll(genDir)
	// 生成GoRepo代码
	dg.GenerateGoRepoCodes(tables, genDir)
	// 生成自定义代码
	opt := &tto.GenerateOptions{
		TplDir:          "./templates",
		AttachCopyright: true,
		OutputDir:       genDir,
		ExcludeFiles:    nil,
	}
	dg.WalkGenerateCodes(tables, opt)
	//格式化代码
	shell.Run("gofmt -w " + genDir)
	t.Log("生成成功, 输出目录", genDir)
}

func getDialect(driver string) orm.Dialect {
	switch driver {
	case "mysql":
		return &orm.MySqlDialect{}
	case "postgres", "postgresql":
		return &orm.PostgresqlDialect{}
	}
	return nil
}
