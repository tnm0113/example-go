// package main

// import (
// 	"github.com/toannm/example-go/json"
// 	"github.com/toannm/example-go/jwt"
// )

// func main() {
// 	jwt.GenJwkSignAndParse()
// 	json.IterateObjectAndArray()
// }

//package main
//
//import (
//    "crypto/tls"
//    "fmt"
//    "golang.org/x/net/context"
//    cc "golang.org/x/oauth2/clientcredentials"
//    "io/ioutil"
//    "net/http"
//)
//
//type SpecialClient struct {
//    *http.Client
//}
//
//func main() {
//    client := NewClient(
//        "98faa314-3a61-46e4-b011-de7c427c39cf",
//        "SlS3PqHY7F5xzNr.wCNQW0wXkD",
//    )
//
//    // the client will update its token if it's expired
//    client.Transport =  &http.Transport{
//        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
//    }
//    resp, err := client.Get("https://backend.vht-iot.vn/api/groups?limit=10&offset=4")
//
//    if err != nil {
//       fmt.Printf("error %v", err)
//    }
//
//    body, _ := ioutil.ReadAll(resp.Body)
//    defer resp.Body.Close()
//
//    // If response code is 200 it was successful
//    if resp.StatusCode == 200 {
//        fmt.Println("The request was successful. Response below:")
//        fmt.Println(string(body))
//    } else {
//        fmt.Println("Could not perform request to the endpoint. Response below:")
//        fmt.Println(string(body))
//    }
//}
//
//func NewClient(cid, csec string) *SpecialClient {
//
//    // this should match whatever service has given you
//    // client credential access
//    config := &cc.Config{
//        ClientID:     cid,
//        ClientSecret: csec,
//        TokenURL: "http://10.55.123.115:30447/oauth2/token",
//        Scopes:   []string{""},
//    }
//
//    // you can modify the client (for example ignoring bad certs or otherwise)
//    // by modifying the context
//    ctx := context.Background()
//    client := config.Client(ctx)
//    return &SpecialClient{client}
//}

package main

import (
   "fmt"
   "github.com/turnage/graw/reddit"
)

func main() {
   bot, err := reddit.NewBotFromAgentFile("reminderbot.agent", 0)
   if err != nil {
       fmt.Println("Failed to create bot handle: ", err)
       return
   }

   harvest, err := bot.ListingWithParams("/api/info.json", map[string]string{"id": "t1_gxpmkx8"})
   if err != nil {
       fmt.Println("Failed to fetch /r/TestMyBotTip: ", err)
       return
   }
    fmt.Printf("harvets %v \n", harvest)
   for _, post := range harvest.Posts {
       fmt.Printf("[%s] posted [%s] link [%s] id [%s]\n", post.Author, post.Title, post.URL, post.ID)
   }
    for _, post := range harvest.Comments {
        fmt.Printf("[%s] posted [%s] link [%s] id [%s]\n", post.Author, post.Body, post.Name, post.ID)
    }
}

//package main
//
//import (
//"fmt"
//"strings"
//"time"
//
//"github.com/turnage/graw"
//"github.com/turnage/graw/reddit"
//)
//
//type reminderBot struct {
//bot reddit.Bot
//}
//
//func (r *reminderBot) Post(p *reddit.Post) error {
//fmt.Printf("new post: %s, self text %s", p.Title, p.SelfText)
// fmt.Println("remind bot")
//if strings.Contains(p.SelfText, "remind me of this post") {
//    <-time.After(10 * time.Second)
//    fmt.Printf("remind: %s", p.Author)
//    fmt.Println("send msg")
//    return r.bot.SendMessage(
//        p.Author,
//        fmt.Sprintf("Reminder: %s", p.Title),
//        "You've been reminded!",
//    )
//}
//return nil
//}
//
//func (r *reminderBot) Message(msg *reddit.Message) error {
//fmt.Printf("receive message from %s content %s", msg.Author, msg.Body)
//return nil
//}
//
//func (r *reminderBot) Mention(mention *reddit.Message) error {
//    fmt.Printf("receive mention from %s content %s link title %s\n", mention.Author, mention.Body, mention.LinkTitle)
//    fmt.Printf("mention subreddit %s mention parent id %s mention subject %s \n", mention.Subreddit, mention.ParentID, mention.Subject)
//    parentId := strings.Split(mention.ParentID, "_")[1]
//    sr := "/r/" + mention.Subreddit + "/comments/" + parentId + "/" + mention.LinkTitle
//    fmt.Printf("subreddit %s \n", sr)
//    harvest, err := r.bot.Listing(sr, "")
//
//    if err != nil {
//    fmt.Println("Failed to fetch /r/TestMyBotTip: ", err)
//    return err
//    }
//    fmt.Printf("harvets %v \n", harvest)
//    for _, post := range harvest.Posts {
//    fmt.Printf("[%s] posted [%s]\n", post.Author, post.Title)
//    }
//    return nil
//}
//
//func main() {
//    if bot, err := reddit.NewBotFromAgentFile("reminderbot.agent", 10); err != nil {
//        fmt.Println("Failed to create bot handle: ", err)
//    } else {
//        fmt.Println("start bot")
//        //err = bot.SendMessage("tnm_tip_bot", "aaaaa", "asddsadas")
//        //if err != nil {
//        //    fmt.Println("failed to send message ", err)
//        //}
//        cfg := graw.Config{Subreddits: []string{"TestMyBotTip"}, Messages: true, Mentions: true}
//        handler := &reminderBot{bot: bot}
//        if _, wait, err := graw.Run(handler, bot, cfg); err != nil {
//            fmt.Println("Failed to start graw run: ", err)
//        } else {
//            fmt.Println("graw run failed: ", wait())
//        }
//    }
//}