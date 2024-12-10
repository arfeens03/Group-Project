package main
import (
	"fmt"
	"flag"
	"log"
)

func main(){
	t := mustToken()
	//token = flags.Get(token)

	//tgClient=telegram.New(token)

	//fetcher = fetcher.New()

	//processor = processor.New(tgClient)

	//consumer.Start(fetcher. processor)

}

func mustToken() string {
token := flag.String(
	name:"token-bot-token", 
	value:"",
	usage:"token for access to telegram bot",
)

flag.Parse()
if *token == "" {
	log.Fatal(v...:"token is not specified")
}
return *token
}
