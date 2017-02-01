package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

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
	smart := "citygift template mode"
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
						title := "G's Academy"
						phrase := "セカイを変えるGEEKを輩出するエンジニア養成学校"
						template := linebot.NewButtonsTemplate(
							imageURL, title, phrase,
							linebot.NewURITemplateAction(title+" Web", "https://citygift-04.herokuapp.com/"),
							linebot.NewURITemplateAction("citygift公式", "https://citygift-04.herokuapp.com/"),
						)
						messaget := "プランはいかがでしたでしょうか？近くにあるオススメスポットも紹介いたします。"
						message1 := linebot.NewTextMessage(messaget)
						message2 := linebot.NewTemplateMessage(smart, template)
						if _, err := bot.ReplyMessage(
							event.ReplyToken,
							message1,
							message2,
						).Do(); err != nil {
							log.Print(err)
						}
					} else {
						imageURL := "https://citygifttest.azurewebsites.net/static/top.jpg"
						phrase := "連絡ありがとうございます。citygiftは対話型サービスとなっています。"
						template := linebot.NewButtonsTemplate(
							imageURL, "Welcome to citygift", phrase,
							linebot.NewURITemplateAction("citygiftとは？", "https://citygift-04.herokuapp.com/"),
							linebot.NewPostbackTemplateAction("プランスタート", "getplan,", ""),
							linebot.NewPostbackTemplateAction("プラン投稿", "pushplan,", ""),
						)
						fmt.Printf("%v", template)
						if _, err := bot.ReplyMessage(
							event.ReplyToken,
							linebot.NewTemplateMessage(smart, template),
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
							smart,
							template,
						),
					).Do(); err != nil {
						log.Print(err)
					}
				}
			} else if event.Type == linebot.EventTypeFollow {
				imageURL := "https://citygifttest.azurewebsites.net/citygift/static/top.jpg"
				phrase := "友達追加ありがとうございます。citygiftは対話型の街歩きプラン紹介サービスです。"
				template := linebot.NewButtonsTemplate(
					imageURL, "Welcome to citygift", phrase,
					linebot.NewURITemplateAction("citygiftとは？", "https://citygift-04.herokuapp.com/"),
					linebot.NewPostbackTemplateAction("プランスタート", "getplan,", ""),
					linebot.NewPostbackTemplateAction("プラン投稿", "pushplan,", ""),
				)
				smart := "本サービスはスマートホン推奨になっております。ボタンが表示されない場合は、Lineのversionが7以上かもお確かめください。"
				message1 := linebot.NewTextMessage(smart)
				message2 := linebot.NewTemplateMessage(smart, template)
				if _, err := bot.ReplyMessage(
					event.ReplyToken,
					message1,
					message2,
				).Do(); err != nil {
					log.Print(err)
				}
			} else if event.Type == linebot.EventTypePostback {

				if postdata := event.Postback.Data; postdata == "pushplan," {
					sorry := "プラン投稿機能はまだ実装できておりません。今しばらくお待ち下さい。"
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(sorry)).Do(); err != nil {
						log.Print(err)
					}
				} else if postdata == "getplan," {
					imageURL := "https://citygifttest.azurewebsites.net/github.com/citygift/static/top.jpg"
					// phrase := "連絡ありがとうございます。citygiftは対話型サービスとなっています。"
					shibuya := postdata + "a_shibuya,"
					fmt.Printf("%v", shibuya)
					template := linebot.NewCarouselTemplate(
						linebot.NewCarouselColumn(
							imageURL, "渋谷エリア", "渋谷・表参道・原宿・代々木上原",
							linebot.NewPostbackTemplateAction("選択", shibuya, "a_shibuya,"),
						),
						linebot.NewCarouselColumn(
							imageURL, "練馬エリア", "石神井公園・練馬・江古田",
							linebot.NewPostbackTemplateAction("選択", postdata+"a_nerima,", ""),
						),
						linebot.NewCarouselColumn(
							imageURL, "鎌倉エリア", "鎌倉..",
							linebot.NewPostbackTemplateAction("選択", postdata+"a_kamakura,", ""),
						),
					)
					message1 := linebot.NewTextMessage("以下のareaからお好きな場所を選択するか位置情報をお送りください")
					message2 := linebot.NewTemplateMessage(smart, template)
					fmt.Printf("%v", template)
					if _, err := bot.ReplyMessage(
						event.ReplyToken,
						message1,
						message2,
					).Do(); err != nil {
						log.Print(err)
					}
				} else if postdata == "getplan,a_shibuya,d" {
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
					message2 := linebot.NewTemplateMessage(smart, template)
					message3 := linebot.NewTemplateMessage(smart, template2)
					fmt.Printf("%v", template)
					if _, err := bot.ReplyMessage(
						event.ReplyToken,
						message1,
						message2,
						message3,
					).Do(); err != nil {
						log.Print(err)
					}
				} else if strings.LastIndexAny(postdata, "getplan") > 0 && strings.LastIndexAny(postdata, ",t_") < 0 {
					imageURL := "https://citygifttest.azurewebsites.net/static/top.jpg"
					// phrase := "連絡ありがとうございます。citygiftは対話型サービスとなっています。"
					phrase := "時間を選択してください"
					template := linebot.NewButtonsTemplate(
						imageURL, "Welcome to citygift", phrase,
						linebot.NewPostbackTemplateAction("1時間", postdata+"t_a", ""),
						linebot.NewPostbackTemplateAction("1.5時間", postdata+"t_b", ""),
						linebot.NewPostbackTemplateAction("2時間", postdata+"t_c", ""),
						linebot.NewPostbackTemplateAction("3時間", postdata+"t_d", ""),
					)
					message := linebot.NewTemplateMessage(smart, template)
					fmt.Printf("%v", template)
					if _, err := bot.ReplyMessage(
						event.ReplyToken,
						message,
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
