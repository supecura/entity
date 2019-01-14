package entity

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func New(token string) (entity Entity,e error){
	discord, err := discordgo.New()
	entity = Entity{}

	if err != nil {
		fmt.Println(err)
		return
	}
	discord.Token = "Bot "+token
	discord.AddHandler(entity.OnMessageCreate)
	err = discord.Open()
	return entity , e
}