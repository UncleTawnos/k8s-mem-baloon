package main

import (
	"fmt"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var memStats = new(runtime.MemStats)

var garbageArray []byte

func main() {
	router := gin.Default()

	router.GET("/alloc", alloc)
	router.GET("/dealloc", dealloc)
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	_ = router.Run(":8080")
}

func alloc(c *gin.Context) {
	garbageArray = make([]byte, 1024*1024*200)
	runtime.GC()

	runtime.ReadMemStats(memStats)
	fmt.Println(memStats.Alloc)
	c.JSON(200, gin.H{
		"message": "alloc",
		"alloc":   memStats.Alloc,
	})
}

func dealloc(c *gin.Context) {
	garbageArray = nil
	runtime.GC()
	runtime.ReadMemStats(memStats)
	fmt.Println(memStats.Alloc)
	c.JSON(200, gin.H{
		"message": "dealloc",
		"alloc":   memStats.Alloc,
	})
}
