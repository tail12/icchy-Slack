package main

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/lvgs/tool-SlackBotAdmin/logging"
	"github.com/nlopes/slack"
	"log"
	"net/http"
	"os"
)

type envConfig struct {
	BotAccessToken    string `envconfig:"BOT_ACCESS_TOKEN" required:"true"`
	BotID             string `envconfig:"BOT_ID" required:"true"`
	ChannelID         string `envconfig:"CHANNEL_ID" required:"true"`
	OauthAccessToken  string `envconfig:"OAUTH_ACCESS_TOKEN" required:"true"`
	VerificationToken string `envconfig:"VERIFICATION_TOKEN" required:"true"`
}

var env envConfig

const Port string = "80"

func init() {
	loadEnvironment()
}

func main() {
	logging.LogSettings()
	os.Exit(_main())
}

func _main() int {
	log.Printf("info: Start Slack event listening")
	client := slack.New(env.BotAccessToken)
	slackListener := &slackListener{
		client:    client,
		botID:     env.BotID,
		channelID: env.ChannelID,
	}

	go slackListener.ListenAndResponse()

	http.HandleFunc("/interaction", interaction)

	log.Printf("info: Server listening on :%s", Port)
	if err := http.ListenAndServe(":"+Port, nil); err != nil {
		log.Printf("error: %s", err)
		return 1
	}

	return 0

}

func loadEnvironment() {
	if err := godotenv.Load(); err != nil {
		log.Printf("error: Faild to load env file")
	}
	loadEnvConfig()
}

func loadEnvConfig() {
	if err := envconfig.Process("", &env); err != nil {
		log.Printf("error: Failed to process env var: %s", err)
	}
}
