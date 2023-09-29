package main

import "strategy-as1/soldier_strategy"

func main() {
	soldier := soldier_strategy.GetConfig()
	singletonSoldier := soldier_strategy.GetConfig()
	basicSoldier := soldier
	BasicSoldier1 := singletonSoldier
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
	soldier5 := soldier_strategy.SoldierBehavior{SB: BasicSoldier1}

	soldier1.DisplayStats()
	soldier2.DisplayStats()
	soldier3.DisplayStats()
	soldier4.DisplayStats()
	soldier5.DisplayStats()
}
