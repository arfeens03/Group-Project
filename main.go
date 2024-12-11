package main
import (
	"flag"
	"log"
)

func main(){
	t := mustToken()
	

	tgClient=telegram.New(musttoken())

	//fetcher = fetcher.New()

	//processor = processor.New(tgClient)

	//consumer.Start(fetcher. processor)

}

func mustToken() string {
token := flag.String(
	name:"token-bot-token", 
	value:" ", 
	usage: "token for access to telegram bot",
)

flag.Parse()
if *token == "" {
	log.Fatal(v...:"token is not specified")
}
return *token
}
