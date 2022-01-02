package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"math"
	"time"
)

func main() {
	r := gin.Default()
	config := SlideRateLimitConfig{
		windowSize: 1000,
		splitNum:   100,
		startTime:  time.Now(),
		limit:      1,
	}

	r.Use(SlideRateLimit(&config))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	log.Fatal(r.Run())

}


func CounterRateLimit(maxCount int, windowSize int64) gin.HandlerFunc {
	counter := 1
	lastTime := time.Now()
	return func(c *gin.Context) {
		now := time.Now()
		if now.Sub(lastTime).Milliseconds() < windowSize {
			if counter < maxCount {
				counter++
				c.JSON(200, gin.H{"message": "success"})
				return
			} else {
				c.AbortWithStatusJSON(400, map[string]string{"message": "请求数量超过限制"})
				return
			}
		}
		lastTime = now
		counter = 0
		return
	}
}

type SlideRateLimitConfig struct {
	windowSize int64     
	splitNum   uint      
	startTime  time.Time 
	limit      int       
}


func SlideRateLimit(config *SlideRateLimitConfig) gin.HandlerFunc {
	counters := make([]int, config.splitNum)
	splitNum := config.splitNum
	index := 0

	slideWindow := func(windowsNum float64) {
		if windowsNum <= 0 {
			return
		}
		slideNum := int(math.Min(windowsNum, float64(config.splitNum)))
		for i := 0; i < slideNum; i++ {
			index = (index + 1) % int(config.splitNum)
			counters[index] = 0
		}

		addTime := int64(windowsNum) * (config.windowSize / int64(config.splitNum))
		config.startTime = config.startTime.Add(time.Duration(addTime) * time.Millisecond)
	}

	return func(c *gin.Context) {
		now := time.Now()
		diffTime := float64(now.Sub(config.startTime).Milliseconds() - config.windowSize)
		windowTime := float64(config.windowSize) / float64(config.splitNum)
		windowsNum := math.Max(diffTime, 0) / (windowTime)
		slideWindow(windowsNum)
		count := 0
		for i := 0; i < int(splitNum); i++ {
			count += counters[i]
		}

		if count > config.limit {
			c.AbortWithStatusJSON(400, map[string]string{"message": "请求数量超过限制"})
			return
		} else {
			counters[index]++
			c.JSON(200, gin.H{"message": "success"})
			return
		}
	}
}
