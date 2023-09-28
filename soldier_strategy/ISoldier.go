package soldier_strategy

import (
	"fmt"
)

type ISoldier interface {
	Attack() int
	HealthPoints() int
	Info()
}

type SoldierBehavior struct {
	SB ISoldier
}

func (soldier SoldierBehavior) DisplayInfo() {
	soldier.SB.Info()
}

func (soldier SoldierBehavior) DisplayStats() {
	fmt.Printf("Soldier's stats: attack %d, health-points %d \n", soldier.SB.Attack(), soldier.SB.HealthPoints())
}
