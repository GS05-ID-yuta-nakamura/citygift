package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}
func main() {
	bot, err := linebot.New(
		os.Getenv("LINE_CHANNEL_SECRET"),
		os.Getenv("LINE_CHANNEL_ACCESS_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", handler)
	// Setup HTTP Server for receiving requests from LINE platform
	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		fmt.Printf("ping\n")
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
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
					} else if userRequest == "confirm" {
						imageURL := "https://citygifttest.azurewebsites.net/static/top.jpg"
						phrase := "連絡ありがとうございます。citygiftは対話型サービスとなっています。"
						template := linebot.NewButtonsTemplate(
							imageURL, "Welcome to citygift", phrase,
							linebot.NewURITemplateAction("citygiftとは？", "https://citygift-04.herokuapp.com/"),
							linebot.NewPostbackTemplateAction("プランスタート", "a", ""),
							linebot.NewPostbackTemplateAction("プラン投稿", "b", ""),
						)
						fmt.Printf("%v", template)
						if _, err := bot.ReplyMessage(
							event.ReplyToken,
							linebot.NewTemplateMessage("Button template", template),
						).Do(); err != nil {
							log.Print(err)
						}
					} else {
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("連絡ありがとうございます。citygiftは対話型サービスとなっています。 citygiftについてもっと知りたい方は、『citygiftとは？』と入力ください プランをお探しの方は、『プランスタート』と入力ください プランを投稿される方は、『プラン投稿』と入力ください。")).Do(); err != nil {
							log.Print(err)
						}
					}
				//位置情報
				case *linebot.LocationMessage:
					template := linebot.NewConfirmTemplate(
						"Do it?",
						linebot.NewPostbackTemplateAction("yes", "yes", ""),
						linebot.NewPostbackTemplateAction("No", "No", ""),
					)
					if _, err := bot.ReplyMessage(
						event.ReplyToken,
						linebot.NewTemplateMessage(
							"Confirm alt text",
							template,
						),
					).Do(); err != nil {
						log.Print(err)
					}
				}
			} else if event.Type == linebot.EventTypeFollow {
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("友達追加ありがとうございます。街歩き体験サービス『citygift』の公式アカウントです！citygiftについてもっと知りたい方は、『citygiftとは？』と入力ください(wink)プランをお探しの方は、『プランスタート』と入力くださいプランを投稿される方は、『プラン投稿』と入力ください。")).Do(); err != nil {
					log.Print(err)
				}
			} else if event.Type == linebot.EventTypePostback {
				if postdata := event.Postback.Data; postdata == "b" {
					sorry := "プラン投稿機能はまだ実装できておりません。今しばらくお待ち下さい。"
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(sorry)).Do(); err != nil {
						log.Print(err)
					}
				} else if postdata == "a" {
					imageURL := "https://citygifttest.azurewebsites.net/static/top.jpg"
					// phrase := "連絡ありがとうございます。citygiftは対話型サービスとなっています。"
					template := linebot.NewCarouselTemplate(
						linebot.NewCarouselColumn(
							imageURL, "hoge", "fuga",
							linebot.NewURITemplateAction("Go to line.me", "https://line.me"),
							linebot.NewPostbackTemplateAction("Say hello1", "hello こんにちは", ""),
						),
						linebot.NewCarouselColumn(
							imageURL, "hoge", "fuga",
							linebot.NewPostbackTemplateAction("言 hello2", "hello こんにちは", "hello こんにちは"),
							linebot.NewMessageTemplateAction("Say message", "Rice=米"),
						),
					)
					fmt.Printf("%v", template)
					if _, err := bot.ReplyMessage(
						event.ReplyToken,
						linebot.NewTemplateMessage("Button template", template),
					).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	})

	// This is just a sample code.
	// For actually use, you must support HTTPS by using `ListenAndServeTLS`, reverse proxy or etc.
	fmt.Printf("server")

	if err := http.ListenAndServe(":"+os.Getenv("HTTP_PLATFORM_PORT"), nil); err != nil {
		log.Fatal(err)
	}
}

//check
