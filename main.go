// package main

// import (
// 	dbmanager "gauth.com/httpd/DBManager"
// 	"gauth.com/httpd/handler"
// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	r := gin.Default()
// 	c := dbmanager.StartConnection()
// 	c.Query("CREATE TABLE GODRIVER (ID INT PRIMARY KEY, NAME VARCHAR(50), AGE INT);")
// 	r.GET("/ping", handler.PingGet())
// 	r.Run("localhost:8024")
// }
package main

import routes "gauth.com/Routes"

func main() {
	db := InitializeDB()
	routes.SetupRoutes(db)

}
