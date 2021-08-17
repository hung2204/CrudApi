# crud
1. Câu lệnh docker để cài mysql trên máy local:
docker run -p 3306:3306 --name mysql -e MYSQL_ROOT_PASSWORD=123456 -d mysql:latest
2. Cài đặt các thư viện: go get -u
"github.com/astaxie/beego/orm"
"github.com/crud/user"
"github.com/golang/glog"
"github.com/labstack/echo/v4"

_ "github.com/go-sql-driver/mysql"
