package main

import (
	"strings"
	"github.com/jipark0716/discordTypecast/model"
	"fmt"
	"time"
	"sync"
	"unicode/utf8"
	"os"
	"math/rand"
	"io/ioutil"

	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
	typecastGo "github.com/jipark0716/typecastGo"
)

const (
	HELP_MAIN = "- &목소리 목소리 설정\n- &{할말} TTS가 읽어줌\n참고 https://github.com/jipark0716/discordTypecast"
	HELP_목소리= "- &목소리 {이름} 하면 설정\n목소리 지원 확인 https://github.com/jipark0716/typecastGo#%EB%AA%A9%EC%86%8C%EB%A6%AC-%EC%A7%80%EC%9B%90"
	ERROR_목소리 = "지원하지 않는 목소리 입니다. 지원하는 목소리를 확인하려면 \"&목소리 도움\"을 입력하세요"
	COMMENT_NOTFOUND = "지원하지 않는 명령어입니다. 도움말이 확인하려면 \"&도움\"을 입력하세요."
)

var typecast typecastGo.TypeCast
var voiceMutex map[string]*sync.Mutex

func init() {
	rand.Seed(time.Now().UnixNano())
	voiceMutex = make(map[string]*sync.Mutex)
	dgvoice.OnError = func(str string, err error) {
		// println(str)
	}
}

func main() {
	var err error
	typecast, err = typecastGo.NewTypeCast(os.Getenv("TYPECAST_ID"), os.Getenv("TYPECAST_pw"))
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
	name := message.Author.Username
	if message.Member.Nick != "" {
		name = message.Member.Nick
	}
	fmt.Printf("[%s]%s/%s %s : %s\n", time.Now().Format("2006-01-02 15:04:05"), message.GuildID, message.ChannelID, name, message.Content)
	if message.Author.ID == discord.State.User.ID {
		return
	}

	if !strings.HasPrefix(message.Content, "&") {
		return
	} else if strings.HasPrefix(message.Content, "&도움") {
		discord.ChannelMessageSend(message.ChannelID, HELP_MAIN)
	} else if strings.HasPrefix(message.Content, "&룰렛") {
		roulette(discord, message)
	} else if strings.HasPrefix(message.Content, "&목소리 도움") {
		discord.ChannelMessageSend(message.ChannelID, HELP_목소리)
	} else if strings.HasPrefix(message.Content, "&벙어리") {
		muteToggle(discord, message)
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

func roulette(discord *discordgo.Session, message *discordgo.MessageCreate) {
	state, err := discord.State.VoiceState(message.GuildID, message.Author.ID)
	if err != nil {
		discord.ChannelMessageSend(message.ChannelID, "음성채널에 먼저 입장해주세요.")
		return
	}

	members := []*discordgo.Member{}

	guild, err := discord.State.Guild(message.GuildID)
	for _, s := range guild.VoiceStates {
		if s.ChannelID == state.ChannelID {
			for _, block := range message.Message.Mentions {
				if s.UserID == block.ID {
					goto ROULETTE_BLOCK
				}
			}
			member, err := discord.State.Member(message.GuildID, s.UserID)
			if err == nil && !member.User.Bot {
				members = append(members, member)
			}
		}
		ROULETTE_BLOCK:
	}
	if len(members) < 2 {
		discord.ChannelMessageSend(message.ChannelID, "참가자가 두명이상이여야 합니다.")
		return
	}
	prevMessage := "참가자 : "
	for _, member := range members {
		prevMessage += member.Mention()+" "
	}
	discord.ChannelMessageSend(message.ChannelID, prevMessage)

	rand.Shuffle(len(members), func(i int, j int) {
		members[i], members[j] = members[j], members[i]
	})

	discord.ChannelMessageSend(message.ChannelID, "당첨 : "+members[0].Mention())
}

func muteToggle(discord *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID != "429996274104926208" {
		discord.ChannelMessageSend(message.ChannelID, "권한이 없어요.")
		return
	}
	if len(message.Message.Mentions) == 0 {
		discord.ChannelMessageSend(message.ChannelID, "맨션 하세요.")
		return
	}
	setting := model.GetTypecastSetting(message.Message.Mentions[0].ID)
	setting.IsMute = !setting.IsMute
	setting.Save()

	discord.ChannelMessageSend(message.ChannelID, "변경 완료")
}

func queueTts(discord *discordgo.Session, message *discordgo.MessageCreate, text string) {
	if utf8.RuneCountInString(text) > 50 {
		discord.ChannelMessageSend(message.ChannelID, "50글자 미만으로 작성해주세요.")
		return
	}
	setting := model.GetTypecastSetting(message.Author.ID)
	if setting.IsMute {
		discord.ChannelMessageSend(message.ChannelID, "당신은 벙어리 입니다.")
		return
	}
	if setting.ActorName == "" {
		discord.ChannelMessageSend(message.ChannelID, "목소리 설정을 먼저 해주세요. 도움말을 확인하려면 \"&목소리 도움\"을 입력하세요.")
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

	fileName := fmt.Sprintf("queue/%s/%d.wav", message.GuildID, chat.ID)
	err = ioutil.WriteFile(fileName, blob, 0707)
	if err != nil {
		discord.ChannelMessageSend(message.ChannelID, "뭔가 잘못됨 2")
		return
	}

	mute, ok := voiceMutex[message.GuildID]
	if !ok {
		mute = &sync.Mutex{}
		voiceMutex[message.GuildID] = mute
	}
	mute.Lock()

	dgv, err := discord.ChannelVoiceJoin(message.GuildID, state.ChannelID, false, true)
	if err != nil {
		discord.ChannelMessageSend(message.ChannelID, "뭔가 잘못됨 3")
		return
	}

	dgvoice.PlayAudioFile(dgv, fileName, make(chan bool))
	os.Remove(fileName)

	mute.Unlock()
}
