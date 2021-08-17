package main

import (
	"github.com/astaxie/beego/orm"
	"github.com/crud/user"
	"github.com/golang/glog"
	"github.com/labstack/echo/v4"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	err := orm.RegisterDataBase("default", "mysql", "root:123456@/crud?charset=utf8")
	if err != nil {
		glog.Fatal("Fail to register database %v", err)
	}
	name := "default"
	force := true
	verbose := true
	err = orm.RunSyncdb(name, force, verbose)
	if err != nil {
		glog.Fatal("Fail to run sync, error: %v", err)
	}
}

func main() {
	server := echo.New()

	server.PUT("/create", user.CreateUser)
	server.GET("/read", user.ReadUser)
	server.POST("/update", user.UpdateUser)
	server.DELETE("/delete", user.DeleteUser)

	server.Logger.Fatal(server.Start(":2204"))
}
