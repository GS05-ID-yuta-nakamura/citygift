package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}
func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	bot, err := linebot.New(
		os.Getenv("LINE_CHANNEL_SECRET"),
		os.Getenv("LINE_CHANNEL_ACCESS_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.POST("/callback", func(c *gin.Context) {
		events, err := bot.ParseRequest(c.Request)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				log.Print(err)
			}
			return
		}

		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					fmt.Printf("%v", message)
					//textmessage
					if userRequest := message.Text; userRequest == "citygiftとは？" {
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("aa")).Do(); err != nil {
							log.Print(err)
						}
					} else if userRequest == "プランスタート" {
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("bb")).Do(); err != nil {
							log.Print(err)
						}
					} else if userRequest == "プラン投稿" {
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("cc")).Do(); err != nil {
							log.Print(err)
						}
					} else {
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("連絡ありがとうございます。citygiftは対話型サービスとなっています。 citygiftについてもっと知りたい方は、『citygiftとは？』と入力ください プランをお探しの方は、『プランスタート』と入力ください プランを投稿される方は、『プラン投稿』と入力ください。")).Do(); err != nil {
							log.Print(err)
						}
					}
				//位置情報
				case *linebot.LocationMessage:
					// fmt.Printf("%v", message)
					// imageURL := "https://citygifttest.azurewebsites.net/static/top.jpg"
					// template := linebot.NewButtonsTemplate(
					// 	imageURL, "My button sample", "Hello, my button",
					// 	linebot.NewURITemplateAction("Go to line.me", "https://line.me"),
					// 	linebot.NewPostbackTemplateAction("Say hello1", "hello こんにちは", ""),
					// 	linebot.NewPostbackTemplateAction("言 hello2", "hello こんにちは", "hello こんにちは"),
					// 	linebot.NewMessageTemplateAction("Say message", "Rice=米"),
					// )
					// if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTemplateMessage("Buttons alt text", template)).Do(); err != nil {
					// 	log.Print(err)
					// }
					fmt.Printf("%v", message)
					template := linebot.NewConfirmTemplate(
						"Do it?",
						linebot.NewMessageTemplateAction("Yes", "Yes!"),
						linebot.NewMessageTemplateAction("No", "No!"),
					)
					if _, err := bot.ReplyMessage(
						event.ReplyToken,
						linebot.NewTemplateMessage("Confirm alt text", template),
					).Do(); err != nil {
						log.Print(err)
					}
				}
			} else if event.Type == linebot.EventTypeFollow {
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("友達追加ありがとうございます。街歩き体験サービス『citygift』の公式アカウントです！citygiftについてもっと知りたい方は、『citygiftとは？』と入力ください(wink)プランをお探しの方は、『プランスタート』と入力くださいプランを投稿される方は、『プラン投稿』と入力ください。")).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	})

	router.Run(":" + port)
}

//check
