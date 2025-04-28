package main

import (
    "go_test/src/db"
    "go_test/src/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    db.Connect()

    r := gin.Default()
    routes.UserSetup(r)
	routes.SaleSetup(r)


    r.Run(":8080")
}
