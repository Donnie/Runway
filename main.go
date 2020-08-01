package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func init() {
	if _, err := os.Stat(".env.local"); os.IsNotExist(err) {
		godotenv.Load(".env")
	} else {
		godotenv.Load(".env.local")
	}
	fmt.Println("Running for " + os.Getenv("ENV"))
}

func main() {
	teleToken, exists := os.LookupEnv("TELEGRAM_TOKEN")
	if !exists {
		fmt.Println("Add TELEGRAM_TOKEN to .env file")
		os.Exit(1)
	}
	filename, exists := os.LookupEnv("DBFILE")
	if !exists {
		fmt.Println("Add DBFILE to .env file")
		os.Exit(1)
	}

	bot, err := tgbotapi.NewBotAPI(teleToken)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))

	global := Global{
		Bot:  bot,
		File: filename,
	}

	r := gin.Default()
	r.POST("/hook", global.handleHook)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, nil)
	})
	r.Run()
}
