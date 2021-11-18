package main

import (
	"errors"
	"fmt"
	"time"
)

// blog：https://zhuanlan.zhihu.com/p/110453784

type Monster struct {
	Name string
}

func NewMonster() Monster {
	return Monster{Name: "zhangsan"}
}


type Player struct {
	Name string
}

func NewPlayer(name string) (Player, error) {
	if time.Now().Unix()%2 == 0 {
		return Player{}, errors.New("player dead")
	}
	return Player{Name: name}, nil
}


type Mission struct {
	Player  Player
	Monster Monster
}

func NewMission(p Player, m Monster) Mission {
	return Mission{p, m}
}

func (m Mission) Start() {
	fmt.Printf("%s defeats %s, world peace!\n", m.Player.Name, m.Monster.Name)
}

func main() {
	/*monster := NewMonster()
	player := NewPlayer("yabingo")
	mission := NewMission(player, monster)*/

	mission, _ := InitMission("yabingo")
	mission.Start()
}