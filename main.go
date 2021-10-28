package main

import (
	"strings"
	"github.com/jipark0716/discordTypecast/model"
	"fmt"
	// "time"
	"os"
	"io/ioutil"

	// "github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
	"github.com/jipark0716/typecastGo"
)

const (
	HELP_MAIN = "- &목소리 목소리 설정\n- &{할말} TTS가 읽어줌\n참고 https://github.com/jipark0716/discordTypecast"
	HELP_목소리= "- &목소리 {이름} 하면 설정\n목소리 지원 \n찬구, 주하, 카밀라, 우주, 혜정, 한유격, 영희, 소영, 지영 경완, 성욱, 애란, 줄리아, 민지, 신혜, 준상, 현경, 성규 윤성, 관용, 정순, 다희, 경숙, 동우, 연길, 지우, 나진, 쥬비 정희, 연우, 용호, 재훈, 정섭, 상도, 진혁, 민상, 찬혁, 지철 수진, 국희, 제임스, 선영, 금희, 명희, 미오, 보노, 덕춘 강수정, 이승주, 주원, 의찬, 마크, 트럼프, 김정은, 문재인 신디, 현진, 정원, 성배, 리춘희, 하은, 범수"
	ERROR_목소리 = "지원하지 않는 목소리 입니다. 지원하는 목소리를 확인하려면 \"&목소리 도움\"을 입력하세요"
	COMMENT_NOTFOUND = "지원하지 않는 명령어입니다. 도움말이 확인하려면 \"&도움\"을 입력하세요."
)

var VoiceConnections map[string]*VoiceConnection

var typecast typecastGo.TypeCast

func init() {
	VoiceConnections = make(map[string]*VoiceConnections)
}

func main() {
	var err error
	typecast, err = typecastGo.NewTypeCast("jipark0716@gmail.com", "h4CZ2gDgHrydinP")
	if err != nil {
		fmt.Printf("%#v\n", err)
		return
	}

	discord, err := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))
	if err != nil {
		fmt.Println(err)
		return
	}

	discord.AddHandler(onCreateMessage)

	err = discord.Open()
	if err != nil {
		fmt.Println(err)
		return
	}

	select{}
}

func onCreateMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == discord.State.User.ID {
		return
	}

	if !strings.HasPrefix(message.Content, "&") {
		return
	} else if strings.HasPrefix(message.Content, "&도움") {
		discord.ChannelMessageSend(message.ChannelID, HELP_MAIN)
	} else if strings.HasPrefix(message.Content, "&목소리 도움") {
		discord.ChannelMessageSend(message.ChannelID, HELP_목소리)
	} else if strings.HasPrefix(message.Content, "&목소리") {
		command := strings.Split(message.Content, " ")
		if len(command) != 2 {
			discord.ChannelMessageSend(message.ChannelID, HELP_목소리)
		} else {
			settingActor(discord, message, command[1])
		}
	} else {
		queueTts(discord, message, strings.Replace(message.Content, "&", "", 1))
	}
}

func settingActor(discord *discordgo.Session, message *discordgo.MessageCreate, actor string) {
	if _, ok := typecastGo.Actor[actor]; !ok {
		discord.ChannelMessageSend(message.ChannelID, ERROR_목소리)
		return
	}
	setting := model.GetTypecastSetting(message.Author.ID)
	setting.ActorName = actor
	err := setting.Save()
	if err == nil {
		discord.ChannelMessageSend(message.ChannelID, "변경완료")
	} else {
		discord.ChannelMessageSend(message.ChannelID, "변경실패")
	}
}

func queueTts(discord *discordgo.Session, message *discordgo.MessageCreate, text string) {
	if len(text) > 50 {
		discord.ChannelMessageSend(message.ChannelID, "50글자 미만으로 작성해주세요.")
		return
	}
	setting := model.GetTypecastSetting(message.Author.ID)
	if setting.ActorName == "" {
		discord.ChannelMessageSend(message.ChannelID, "목소리 설정을 먼저 해주세요. 도움말을 확인하려면 \"$목소리 도움\"을 입력하세요.")
		return
	}

	state, err := discord.State.VoiceState(message.GuildID, message.Author.ID)

	if err != nil {
		discord.ChannelMessageSend(message.ChannelID, "음성채널에 먼저 입장해주세요.")
		return
	}

	blob, err := typecast.Exec(text, setting.ActorName)
	if err != nil {
		discord.ChannelMessageSend(message.ChannelID, "뭔가 잘못됨 1")
		return
	}
	os.Mkdir(fmt.Sprintf("queue/%s", message.GuildID), 0707)

	chat := model.NewChat()
	chat.MemberID = message.Author.ID
	chat.GuildID = message.GuildID
	chat.ChannelID = message.ChannelID
	chat.Content = text
	chat.Save()

	err = ioutil.WriteFile(fmt.Sprintf("queue/%s/%d.wav", message.GuildID, chat.ID), blob, 0707)
	if err != nil {
		discord.ChannelMessageSend(message.ChannelID, "뭔가 잘못됨 2")
		return
	}

	dgv, ok := VoiceConnections[message.GuildID]
	if !ok || dgv.ChannelID != message.ChannelID {
		dgv, err = discord.ChannelVoiceJoin(message.GuildID, *ChannelID, false, true)
		if err != nil {
			fmt.Println(err)
			return
		}

		VoiceConnections[message.GuildID] = dvg
	}
}
