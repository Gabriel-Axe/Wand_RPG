package main

var (
	// NOTE: Physical
	AttackHeadbutt = Attack{
		Name: "headbutt",
		Damage: 5,
	}
	AttackSlash = Attack{
		Name: "slash",
		Damage: 7,
	}
	AttackRoundkick = Attack{
		Name: "round kick",
		Damage: 4,
	}
	AttackPoisonSting = Attack {
		Name: "Poison Sting",
		Damage: 5,
		Effect: PoisonEffect{DamagePerTurn: 3, Turns: 7},
	}

	// NOTE: Magic
	AttackFireball = Attack {
		Name: "Fireball",
		Damage: 10,
		Effect: FireEffect{DamagePerTurn: 5, Turns: 3},
	}
)
