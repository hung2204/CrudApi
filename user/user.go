package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/astaxie/beego/orm"
	"github.com/golang/glog"
	"github.com/labstack/echo/v4"
)

func init() {
	orm.RegisterModel(new(User))
}

type User struct {
	Id    int    `orm:"auto" json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Phone string `json:"phone"`
}

func CreateUser(c echo.Context) error {
	user := &User{}
	if err := c.Bind(user); err != nil {
		glog.Errorf("bind user error: %v", err)
		return err
	}
	o := orm.NewOrm()
	id, err := o.Insert(user)
	if err != nil {
		glog.Errorf("insert user error: %v", err)
		return err
	}
	glog.Infof("insert at row: %v", id)
	return c.JSON(http.StatusOK, user)
}

func ReadUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.QueryParam("id"))
	// name := c.QueryParam("name")
	o := orm.NewOrm()
	user := &User{
		Id: id,
	}
	err := o.Read(user, "id")
	if err != nil {
		glog.Errorf("get user %s error: %v", id, err)
		return err
	}
	return c.JSON(http.StatusOK, user)
}

func UpdateUser(c echo.Context) error {
	user := &User{}
	if err := c.Bind(user); err != nil {
		glog.Errorf("bind user error: %v", err)
		return err
	}
	glog.Infof("req update user: %+v", user)
	o := orm.NewOrm()
	_, err := o.Update(user, "Name", "Age", "Phone")
	if err != nil {
		glog.Errorf("update user %s error: %v", user, err)
		return err
	}
	userUpdate := &User{
		Name: user.Name,
	}
	o.Read(userUpdate, "Name")
	return c.JSON(http.StatusOK, userUpdate)
}

func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.FormValue("id"))
	glog.Infof("Deleting user %d", id)
	user := &User{
		Id: id,
	}
	o := orm.NewOrm()
	_, err := o.Delete(user)
	if err != nil {
		glog.Errorf("delete user %d error %v\n", id, err)
		return err
	}
	return c.String(http.StatusOK, fmt.Sprintf("delete complete user id: %d", id))
}
