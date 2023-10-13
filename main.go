package main

import (
	"strategy-as1/config"
	"strategy-as1/soldier_strategy"
)

func main() {
	//singletone
	instanceOfDB := config.GetInstanceOfDB()
	instanceOfDB.Info()
	configSingleton := config.GetInstanceOfDB()
	configSingleton.Info()
	basicSoldier := soldier_strategy.BasicSoldier{}
	bowSoldier := soldier_strategy.BowSoldier{Soldier: basicSoldier}
	shieldSoldier := soldier_strategy.ShieldSoldier{Soldier: basicSoldier}
	shieldBowSoldier := soldier_strategy.ShieldSoldier{
		Soldier: soldier_strategy.BowSoldier{
			Soldier: basicSoldier,
		},
	}

	soldier1 := soldier_strategy.SoldierBehavior{SB: basicSoldier}
	soldier2 := soldier_strategy.SoldierBehavior{SB: bowSoldier}
	soldier3 := soldier_strategy.SoldierBehavior{SB: shieldSoldier}
	soldier4 := soldier_strategy.SoldierBehavior{SB: shieldBowSoldier}

	soldier1.DisplayStats()
	soldier2.DisplayStats()
	soldier3.DisplayStats()
	soldier4.DisplayStats()
}
