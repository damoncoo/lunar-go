package main

import(
	"fmt"
	rt "github.com/damoncoo/lunar-go/router"
)

func main(){
	
	router := rt.Router()
	router.Run(":4000")
	fmt.Println("Server Started...")
}
