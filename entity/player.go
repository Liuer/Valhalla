package entity

import (
	"fmt"

	"github.com/Hucaru/Valhalla/constant"
	"github.com/Hucaru/Valhalla/mnet"
	"github.com/Hucaru/Valhalla/mpacket"
)

// Players type alias
type Players []*Player

// GetFromConn retrieve the player from the connection
func (p *Players) GetFromConn(conn mnet.Client) (*Player, error) {
	for _, v := range *p {
		if v.conn == conn {
			return v, nil
		}
	}

	return nil, fmt.Errorf("Could not retrieve player")
}

// GetFromName retrieve the player from the connection
func (p *Players) GetFromName(name string) (*Player, error) {
	for _, v := range *p {
		if v.char.name == name {
			return v, nil
		}
	}

	return nil, fmt.Errorf("Could not retrieve player")
}

// RemoveFromConn removes the player based on the connection
func (p *Players) RemoveFromConn(conn mnet.Client) error {
	i := -1

	for j, v := range *p {
		if v.conn == conn {
			i = j
			break
		}
	}

	if i == -1 {
		return fmt.Errorf("Could not find player")
	}

	(*p)[i] = (*p)[len(*p)-1]
	(*p)[len(*p)-1] = nil
	(*p) = (*p)[:len(*p)-1]

	return nil
}

// Player connected to server
type Player struct {
	conn       mnet.Client
	char       Character
	instanceID int
}

func NewPlayer(conn mnet.Client, char Character) *Player {
	return &Player{conn: conn, char: char, instanceID: 0}
}

func (p Player) Char() Character {
	return p.char
}

func (p Player) InstanceID() int {
	return p.instanceID
}

func (p *Player) SetInstance(id int) {
	p.instanceID = id
}

func (p Player) Send(packet mpacket.Packet) {
	p.conn.Send(packet)
}

func (p *Player) SetJob(id int16) {
	p.char.job = id
	p.conn.Send(PacketPlayerStatChange(true, constant.JobID, int32(id)))
}

func (p *Player) SetLevel(amount byte) {

}

func (p *Player) GiveLevel(amount byte) {

}

func (p *Player) SetStr(amount int16) {

}

func (p *Player) SetDex(amount int16) {

}

func (p *Player) SetInt(amount int16) {

}

func (p *Player) SetLuk(amount int16) {

}

func (p *Player) SetHP(amount int16) {

}

func (p *Player) SetMaxHP(amount int16) {

}

func (p *Player) SetMP(amount int16) {

}

func (p *Player) SetMaxMP(amount int16) {

}

func (p *Player) SetAP(amount int16) {

}

func (p *Player) SetSp(amount int16) {

}

func (p *Player) SetEXP(amount int32) {

}

func (p *Player) GiveEXP(amount int32) {

}

func (p *Player) SetFame(amount int16) {

}

func (p *Player) SetGuild(name string) {

}

func (p *Player) SetEquipSlotSize(size byte) {

}

func (p *Player) SetUseSlotSize(size byte) {

}

func (p *Player) SetEtcSlotSize(size byte) {

}

func (p *Player) SetCashSlotSize(size byte) {

}

func (p *Player) SetMesos(amount int32) {

}

func (p *Player) GiveMesos(amount int32) {

}

func (p *Player) SetMinigameWins(amount int32) {

}

func (p *Player) SetMinigameDraws(amount int32) {

}

func (p *Player) SetMinigameLoss(amount int32) {

}

func (p *Player) UpdateMovement(frag movementFrag) {
	p.char.pos.x = frag.x
	p.char.pos.y = frag.y
	p.char.foothold = frag.foothold
	p.char.stance = frag.stance
}

func (p *Player) SetPos(pos pos) {
	p.char.pos = pos
}

func (p Player) CheckPos(pos pos, xRange, yRange int16) bool {
	var xValid, yValid bool

	if xRange == 0 {
		xValid = p.char.pos.x == pos.x
	} else {
		xValid = (pos.x-xRange < p.char.pos.x && p.char.pos.x < pos.x+xRange)
	}

	if yRange == 0 {
		xValid = p.char.pos.y == pos.y
	} else {
		yValid = (pos.y-yRange < p.char.pos.y && p.char.pos.y < pos.y+yRange)
	}

	return xValid && yValid
}

func (p Player) Pos() pos {
	return p.char.pos
}

func (p *Player) SetFoothold(fh int16) {
	p.char.foothold = fh
}

func (p *Player) SetMapID(id int32) {
	p.char.mapID = id
}

func (p *Player) SetMapPosID(pos byte) {
	p.char.mapPos = pos
}
