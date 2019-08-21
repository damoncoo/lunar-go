package main

import (
	"fmt"

	rt "github.com/damoncoo/lunar-go/router"
)

func main() {

	serverFileUpload()

	router := rt.Router()
	_ = router.Run(":4000")
	fmt.Println("Server Started...")
}
