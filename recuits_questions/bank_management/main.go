package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"strconv"
)

var (
	DB *gorm.DB
)

// User Model
type User struct {
	ID              int     `json:"id"`
	Name            string  `json:"name"`
	Password        string  `json:"password"`
	Money           float64 `json:"money"`
	Money_type       string  `json:"money_type"`
	X        		int  	`json:"x"`
	Y        		int  	`json:"Y"`
}

func initMySQL() (err error) {
	dsn := "root:zjfzjf@tcp(127.0.0.1)/bank?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

func main() {
	// 创建数据库
	// 连接数据库
	// 先在数据库创建一个名为bank的库
	err := initMySQL()
	if err != nil {
		panic(err)
	}
	defer DB.Close() //程序退出关闭数据库连接
	//模型绑定
	DB.AutoMigrate(&User{})

	r := gin.Default()
	//引用的静态文件
	r.Static("/statics", "./html")
	//模板文件
	r.LoadHTMLFiles("html/templates/index.html",
							"html/templates/login.html",
							"html/templates/register.html",
							"html/templates/zhuanzhang.html")

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", nil)
	})
	r.GET("/zhuanzhang", func(c *gin.Context) {
		c.HTML(http.StatusOK, "zhuanzhang.html", nil)
	})
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	// v1
	//银行 登录,查看,修改,转账,
	v1Group := r.Group("v1")
	{
		// 注册
		v1Group.POST("/sighup", func(c *gin.Context) {
			// 前端页面填写数据 点击提交 发请求到这里
			// 1. 从请求中把数据拿出来
			var user User
			_ = c.BindJSON(&user) //这个函数自动返回前端的数据
			// 2. 存入数据库
			err = DB.Create(&user).Error
			if err != nil {
			// 3. 反回响应
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, user)
			}
		})

		//登录判断
		//返回带有一个map[string]bool的json文件
		v1Group.GET("/login/: id",func(c *gin.Context){
			id, ok := c.Params.Get("id")

			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
				return
			}
			password,ok := c.Params.Get("password")
			if !ok{
				c.JSON(http.StatusOK,gin.H{"error": "无效的id"})
				return
			}

			var user User

			if err = DB.Where("id=?", id).First(&user).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
				return
			}else{
				if password == user.Password{
					c.JSON(http.StatusOK,gin.H{"login":true})
				}else{
					c.JSON(http.StatusOK,gin.H{"login":false})
				}
			}

		})

		//查询
		//返回包含信息的json文件
		v1Group.GET("/bank/:id", func(c *gin.Context) {

			id, ok := c.Params.Get("id")

			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
				return
			}

			var user User
			if err = DB.Where("id=?", id).First(&user).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
				} else {
					c.JSON(http.StatusOK, gin.H{
						"id" :    user.ID,
						"name" :  user.Name,
						"Money" : user.Money,
						"Money_type" :user.Money_type,
						"X" : user.X,
						"Y" : user.Y,
				})
			}
		})

		// 转账
		v1Group.PUT("/trans/:id/:id2/:password/:account", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "无效的本人id"})
				return
			}

			id2, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "无效的对方id"})
				return
			}

			Password, ok := c.Params.Get("password")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "无效的支付密码"})
				return
			}

			account, ok := c.Params.Get("account")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "无效的数值"})
				return
			}

			number, err := strconv.ParseFloat(account,64)

			var user User
			var target User

			if err = DB.Where("id=?", id).First(&user).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
				return
			}

			if user.Password == Password {

				_ = c.BindJSON(&user)
				if err = DB.Save(&user).Error; err != nil {
					c.JSON(http.StatusOK, gin.H{"error": err.Error()})
				} else {
					temp := user.Money
					temp = number - number
					c.JSON(http.StatusOK, gin.H{
						"id":    user.ID,
						"name":  user.Name,
						"Money": temp,
						"Money_type" :user.Money_type,
						"X" : user.X,
						"Y" : user.Y,
					})
				}

				if err = DB.Where("id=?", id2).First(&target).Error; err != nil {
					c.JSON(http.StatusOK, gin.H{"error": err.Error()})
					return
				} else {
					temp := user.Money
					temp = number + number
					c.JSON(http.StatusOK, gin.H{
						"id":    user.ID,
						"name":  user.Name,
						"Money": temp,
						"Money_type" :user.Money_type,
						"X" : user.X,
						"Y" : user.Y,
					})
				}

				_ = c.BindJSON(&target)
				if err = DB.Save(&target).Error; err != nil {
					c.JSON(http.StatusOK, gin.H{"error": err.Error()})
				} else {
					c.JSON(http.StatusOK, target)
				}
			}
		})

	}
	_ = r.Run(":9090")
	//运行后访问127.0.0.1:9090/index
}
