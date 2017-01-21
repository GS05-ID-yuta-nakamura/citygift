package main  
import (  
    "fmt"
    "net/http"
    "os" 
    "log"

    "github.com/line/line-bot-sdk-go/linebot"
)

func handler(w http.ResponseWriter, r *http.Request) {  
    fmt.Fprintf(w, "Hi there - this page was served using Go \\o/")
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
                    if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
                        log.Print(err)
                    }
                }
            }
        }
    })

    // This is just a sample code.
    // For actually use, you must support HTTPS by using `ListenAndServeTLS`, reverse proxy or etc.
    fmt.Printf("サーバーを起動しています...")

    if err := http.ListenAndServe(":"+os.Getenv("HTTP_PLATFORM_PORT"), nil); err != nil {
        log.Fatal(err)
    }
}