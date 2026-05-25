package main

type Effect interface {
	Apply(target *Unit, attacker *Unit)
}

type StatusEffect struct {
    Type     string `json:"type"`
    Damage   int    `json:"damage"`
    Duration int    `json:"duration"`
}

type PoisonEffect struct {
	DamagePerTurn int
	Turns int
}

var AttackPoisonSting = Attack {
	Name: "Poison Sting",
	Damage: 10,
	Effect: PoisonEffect{DamagePerTurn: 5, Turns: 3},
}

func (p PoisonEffect) Apply(target *Unit, attacker *Unit) {
	target.Effects = append(target.Effects, StatusEffect{
		Type: "poison",
		Damage: p.DamagePerTurn,
		Duration: p.Turns,
	})
}
