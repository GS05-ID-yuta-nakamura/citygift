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
					if userRequest := message.Text; userRequest == "プラン終了" {
						imageURL := "https://citygifttest.azurewebsites.net/static/top.jpg"
						phrase := "いかがでしたでしょうか？近くにあるオススメスポットを紹介いたします。"
						template := linebot.NewButtonsTemplate(
							imageURL, "Thank you for your coming", phrase,
							linebot.NewURITemplateAction("G's academy", "https://citygift-04.herokuapp.com/"),
							linebot.NewURITemplateAction("citygift公式", "https://citygift-04.herokuapp.com/"),
						)
						if _, err := bot.ReplyMessage(
							event.ReplyToken,
							linebot.NewTemplateMessage("Button template", template),
						).Do(); err != nil {
							log.Print(err)
						}
					} else {
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
				imageURL := "https://citygifttest.azurewebsites.net/static/top.jpg"
				phrase := "友達追加ありがとうございます。citygiftは対話型サービスとなっています。"
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
							imageURL, "渋谷エリア", "渋谷・表参道・原宿・代々木上原",
							linebot.NewPostbackTemplateAction("選択", postdata+"a", ""),
						),
						linebot.NewCarouselColumn(
							imageURL, "練馬エリア", "石神井公園・練馬・江古田",
							linebot.NewPostbackTemplateAction("選択", postdata+"b", ""),
						),
						linebot.NewCarouselColumn(
							imageURL, "鎌倉エリア", "鎌倉..",
							linebot.NewPostbackTemplateAction("選択", postdata+"c", ""),
						),
					)
					message1 := linebot.NewTextMessage("以下のareaからお好きな場所を選択するか位置情報をお送りください")
					message2 := linebot.NewTemplateMessage("carousel template", template)
					fmt.Printf("%v", template)
					if _, err := bot.ReplyMessage(
						event.ReplyToken,
						message1,
						message2,
					).Do(); err != nil {
						log.Print(err)
					}
				} else if postdata == "aa" {
					imageURL := "https://citygifttest.azurewebsites.net/static/top.jpg"
					// phrase := "連絡ありがとうございます。citygiftは対話型サービスとなっています。"
					phrase := "時間を選択してください"
					template := linebot.NewButtonsTemplate(
						imageURL, "Welcome to citygift", phrase,
						linebot.NewPostbackTemplateAction("1時間", postdata+"a", ""),
						linebot.NewPostbackTemplateAction("1.5時間", postdata+"b", ""),
						linebot.NewPostbackTemplateAction("2時間", postdata+"c", ""),
						linebot.NewPostbackTemplateAction("3時間", postdata+"d", ""),
					)
					message := linebot.NewTemplateMessage("carousel template", template)
					fmt.Printf("%v", template)
					if _, err := bot.ReplyMessage(
						event.ReplyToken,
						message,
					).Do(); err != nil {
						log.Print(err)
					}
				} else if postdata == "aad" {
					imageURL := "https://citygifttest.azurewebsites.net/static/top.jpg"
					// phrase := "連絡ありがとうございます。citygiftは対話型サービスとなっています。"
					phrase := "表参道エリア3時間満喫コース"
					template := linebot.NewButtonsTemplate(
						imageURL, "Welcome to citygift", phrase,
						linebot.NewURITemplateAction("webでみる", "https://citygift-04.herokuapp.com/"),
					)
					template2 := linebot.NewConfirmTemplate(
						"こちらのプランでよろしいですか?",
						linebot.NewPostbackTemplateAction("yes", "yes", ""),
						linebot.NewPostbackTemplateAction("No", "No", ""),
					)
					message1 := linebot.NewTextMessage("おすすめのプランを探して参りました。")
					message2 := linebot.NewTemplateMessage("carousel template", template)
					message3 := linebot.NewTemplateMessage("carousel template", template2)
					fmt.Printf("%v", template)
					if _, err := bot.ReplyMessage(
						event.ReplyToken,
						message1,
						message2,
						message3,
					).Do(); err != nil {
						log.Print(err)
					}
				} else if postdata := event.Postback.Data; postdata == "yes" {
					message1 := linebot.NewTextMessage("ぜひお楽しみください。")
					message2 := linebot.NewTextMessage("終了の際は『プラン終了』とご入力ください。")
					if _, err := bot.ReplyMessage(
						event.ReplyToken,
						message1,
						message2,
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
