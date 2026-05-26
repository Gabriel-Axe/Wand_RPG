package main

func QuickGameSetup() *Game {

	team_1 := MakeGoblinTeam()

	team_1[0].ID = 1
	team_1[1].ID = 2
	team_1[2].ID = 3

	team_2 := MakeElvenTeam()

	team_2[0].ID = 1
	team_2[1].ID = 2
	team_2[2].ID = 3

	p1 := &player{ID: 1, Name: "Alexander", Team: team_1[:]}
	p2 := &player{ID: 2, Name: "Oliver", Team: team_2[:]}

	g := Game{Attacker: p1, Defender: p2, Turn: 1}

	return &g
}

func ManaTestSetup() *Game {

	team_1 := MakeElvenTeam()

	team_1[0].ID = 1
	team_1[1].ID = 2
	team_1[2].ID = 3

	team_2 := MakeElvenTeam()

	team_2[0].ID = 1
	team_2[1].ID = 2
	team_2[2].ID = 3

	p1 := &player{ID: 1, Name: "Alexander", Team: team_1[:]}
	p2 := &player{ID: 2, Name: "Oliver", Team: team_2[:]}

	g := Game{Attacker: p1, Defender: p2, Turn: 1}

	return &g
}

func MakeGoblinTeam() [3]*Unit {
	var team [3] *Unit
	team[0] = MakeGoblin()
	team[1] = MakeGoblin()
	team[2] = MakeGoblin()

	return team
}

// NOTE: I ain believin i was making a Elvis team this whole time, what a great day for short kings
func MakeElvenTeam() [3]*Unit {
	var team [3] *Unit
	team[0] = MakeElven()
	team[1] = MakeElven()
	team[2] = MakeElven()

	return team
}
func MakeWerewolfTeam() [3]*Unit {
	var team [3] *Unit
	team[0] = MakeWerewolf()
	team[1] = MakeWerewolf()
	team[2] = MakeWerewolf()

	return team
}

func GetGameStatus(g *Game) map[string]interface{} {
	if g == nil {
		return map[string]interface{}{"error": "No game started"}
	}

	players := make([]map[string]interface{}, 2)
	players[0] = map[string]interface{}{
		"id": g.Attacker.ID,
	}
	players[1] = map[string]interface{}{
		"id": g.Defender.ID,
	}

	return map[string]interface{}{
		"defender": players,
		"turn": g.Turn,
	}
}

func MakeGoblin() *Unit {
	return &Unit{
		Name: "Goblin",
		Health: 70,
		ManaPool: 50,
		Types: []UnitType{
			*Flesh,
		},
		Attacks: []Attack{
			AttackHeadbutt,
		},
	}
}

func MakeWerewolf() *Unit {
	return &Unit{
		Name: "Werewolf",
		Health: 130,
		ManaPool: 15,
		Attacks: []Attack{
			AttackSlash,
			AttackHeadbutt,
		},
	}
}

func MakeElven() *Unit {

	cool_wand := 	NewWand(
		"Wand of Instant Barbcue",
		[]Attack{
			AttackFireball,
		},
		100, 5)

		return &Unit{
			Name: "Elven",
			Health: 100,
			ManaPool: 100,
			Items: []ItemHolder{
				cool_wand,
			},
			Types: []UnitType{
				*Flesh,
			},
			Attacks: []Attack{
				AttackSlash,
				AttackRoundkick,
			},
		}
}
