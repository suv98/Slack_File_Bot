package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	os.Setenv("SLACK-BOT-TOKEN", "xoxb-3467133186801-3460082958484-DrnGkjsVmHQRCSzAB89HLmU0")
	os.Setenv("CHANNEL-ID", "C03CXSGE9K9")
	api := slack.New(os.Getenv("SLACK-BOT-TOKEN"))
	ChannelArr := []string{os.Getenv("CHANNEL-ID")}
	fileArr := []string{"text.pdf"} //we can add more number of files to it

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: ChannelArr,
			File:     fileArr[i],
		}

		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
		}
		fmt.Printf("Name=%s\n,URL=%s\n", file.Name, file.URL)
	}
}
