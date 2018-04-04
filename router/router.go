package router

import (
	"fmt"
	"net/http"
	"time"
	"strconv"
	"github.com/gin-gonic/gin"
	calendar "github.com/damoncoo/lunar-go/calendar"
)

//Router 导出的路由
func Router() *gin.Engine {

	router := gin.Default()

	router.GET("/lunar/:timestamp/sex/:sex",func(c *gin.Context){

		timestamp := c.Param("timestamp")
		// nl := calendar.NewLunarNow()
		// fmt.Println("现在是" +  strconv.Itoa(nl.Year()) + "年" + strconv.Itoa(nl.Month()) + "月" + strconv.Itoa(nl.Day()) + "日" + strconv.Itoa(nl.Hour()) + "时" )

		// sx := calendar.AnimalYear(nl.Year())
		// fmt.Println("今年是" + sx)

		// yl := nl.Convert()
		// fmt.Println(
		// "现在是" +  strconv.Itoa(yl.Year()) + 
		// "年" + strconv.Itoa(int(yl.Month())) +
		// "月" + strconv.Itoa(yl.Day()) + 
		// "日" + strconv.Itoa(yl.Hour()) +
		// "时" )

		i, err := strconv.ParseInt(timestamp, 10, 64)
		if err != nil {
			panic(err)
		}
		t := time.Unix(i,0)
		yl := calendar.NewSolarTime(t)
		fmt.Println(yl)		

		nl := yl.Convert()
		fmt.Println(nl)				
		fmt.Println(calendar.LunarDayString(nl.Day()))

		rs := map[string]string {"Hello":"World" }
		c.JSON(http.StatusOK,rs)

	})
	return router
}