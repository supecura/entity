package entity

import "github.com/bwmarrin/discordgo"

type Role int

const (
	Survivor	Role = iota
	Killer
)

func (role Role) Name() string {
	switch role {
	case Survivor:
		return "Survivor"
	case Killer:
		return "Killer"
	default:
		return "Unknown"
	}
}

func (role Role)Value(name string) Role {
	switch name {
		case "killer":
			return Killer
		default:
			return Survivor
	}
}

func (role Role)Roles() []string{
	return []string{"survivor","killer"}
}


//Player ...
type Player struct {
	ID          string
	Role        Role
	Park        []Park
	Item        Item
	addon       Addon
	Offering    Offering
	PickPattern PickPattern
}


func (player Player) ShowPlayerInfo(s *discordgo.Session,c *discordgo.Channel){
	var message string
	message += "ID"+"\t"+ player.ID
	//SendMessage(s,c,message)
}

func (player Player) GetPickPattern() PickPattern {
	return player.PickPattern
}

func (player Player) SetPickPattern(p PickPattern) {
	player.PickPattern = p
}

func (player Player) GetRole() Role {
	return player.Role
}

func (player Player) SetRole(role Role) {
	player.Role = role
}

//NewPlayer
func NewPlayer(id string, role Role) Player {
	var p PickPattern
	player := Player{ID: id, Role: role,PickPattern:p.Value("all")}
	return player
}

func (player Player) Equipment() string {
	var role Role
	var equipment string
	equipment += "Park"
	for i := 0; i < 4; i++ {
		equipment += "\n"+ player.Park[i].JapaneseName+"\t"+ player.Park[i].EnglishName
	}
	equipment += "\n"
	equipment += "\n"
	if player.Role == role.Value("survivor"){
		equipment += "Item"
		equipment += "\n"+ *player.Item.JapaneseName+"\t"+ *player.Item.EnglishName
		equipment += "\n"
		equipment += "\n"
	}
	equipment += "Offering"
	equipment += "\n"+ *player.Offering.JapaneseName+"\t"+ *player.Offering.EnglishName
	return equipment
}

