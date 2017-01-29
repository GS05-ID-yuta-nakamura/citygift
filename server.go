package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}
func main() {
	bot, err := linebot.New(
		"476d25dd83bd1f41c82952d0aa01919f",
		"XtiilwvVjxg7wHxVITFdPeb+yuC3zrzRUH+8YdgYu96vnkk7sqrLYKZ8CENpGi15Bls/s8GHCJHhEwyQUuSn09XzxaJAifCALrCyBiTMOCYm280ltAC0gXRuP/znjtvYcGUpcOiyJT9qJtGY0cyZzgdB04t89/1O/w1cDnyilFU=",
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
						left := linebot.NewMessageTemplateAction("Yes", "Yes!")
						fmt.Printf("left")
						fmt.Printf("%v", left)
						right := linebot.NewMessageTemplateAction("No", "No!")
						fmt.Printf("right")
						fmt.Printf("%v", right)
						fmt.Printf("templete")
						template := linebot.NewConfirmTemplate(
							"Do it?",
							left,
							right,
						)
						fmt.Printf("%v", template)
						fmt.Printf("replycontent")
						replycontent := linebot.NewTemplateMessage("Confirm alt text", template)
						fmt.Printf("%v", replycontent)
						if _, err := bot.ReplyMessage(
							event.ReplyToken,
							replycontent,
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
					if _, err := bot.ReplyMessage(
						event.ReplyToken,
						linebot.NewTemplateMessage(
							"Confirm alt text",
							linebot.NewConfirmTemplate(
								"Do it?",
								linebot.NewPostbackTemplateAction("Yes", "Yes", "Yes"),
								linebot.NewPostbackTemplateAction("No", "No", "No"),
							),
						),
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

	// This is just a sample code.
	// For actually use, you must support HTTPS by using `ListenAndServeTLS`, reverse proxy or etc.
	fmt.Printf("サーバーを起動しています...")

	if err := http.ListenAndServe(":1337", nil); err != nil {
		log.Fatal(err)
	}
}

//check
