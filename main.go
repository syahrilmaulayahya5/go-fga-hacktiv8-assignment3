package main

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	ginEngine := gin.Default()

	ginEngine.LoadHTMLFiles("template/index.html")
	type DataPoint struct {
		WaterLevel  int    `json:"water"`
		WaterStatus string `json:"water_status"`
		WindLevel   int    `json:"wind"`
		WindStatus  string `json:"wind_status"`
	}
	data := []DataPoint{}
	ginEngine.GET("/index", func(c *gin.Context) {
		newData := DataPoint{
			WaterLevel: rand.Intn(100),
			WindLevel:  rand.Intn(100),
		}

		waterStatus := "AMAN"
		if newData.WaterLevel > 6 && newData.WaterLevel < 8 {
			waterStatus = "SIAGA"
		} else if newData.WaterLevel > 8 {
			waterStatus = "BAHAYA"
		}
		windStatus := "AMAN"
		if newData.WindLevel > 7 && newData.WindLevel < 15 {
			windStatus = "SIAGA"
		} else if newData.WindLevel > 15 {
			windStatus = "BAHAYA"
		}
		newData.WaterStatus = waterStatus
		newData.WindStatus = windStatus

		data = append(data, newData)

		file, _ := json.MarshalIndent(data, "", " ")
		_ = ioutil.WriteFile("test.json", file, 0644)

		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":        "Value website",
			"water_status": waterStatus,
			"wind_status":  windStatus,
			"data":         data,
		})
	})
	ginEngine.Run(":8080")
}
