package main

import "testing"

func TestApplyEffect(t *testing.T) {
	g := QuickGameSetup()
	d := g.Defender.Team[0]
	a := g.Attacker.Team[0]

	before := d.Health
	logGameState(t, g)

	AttackFireball.Effect.Apply(d, a, &AttackFireball)

	NextTurn(g)
	after := d.Health
	assertHealthChanged(t, before, after)

	NextTurn(g)
	after = d.Health
	assertHealthChanged(t, before, after)

	NextTurn(g)
	if len(d.Effects) > 0 {
		t.Fatalf("Effect should have ended, have %d effects, effect 0 having %d turns left", len(d.Effects), d.Effects[0].Duration)
	}
}

func TestManaUsage(t *testing.T) {

	g := &Game{
		Attacker: &player{
			Name: "",
			Team: []*Unit{
				MakeElven(),
			},
		},
		Defender: &player{
			Name: "",
			Team: []*Unit{
				MakeGoblin(),
			},
		},
	}

	d := g.Defender.Team[0]
	a := g.Attacker.Team[0]

	if a.ManaPool == 0 {
		t.Fatalf("Unit %s mana pool is 0", a.Name)
	}

	before := a.ManaPool
	logGameState(t, g)

	elven := MakeElven()
	if a.ManaPool != elven.ManaPool {
		t.Fatalf("Mana is not the same: unit=%d, elven=%d", a.ManaPool, elven.ManaPool)
	}

	err := AttackFireball.Effect.Apply(d, a, &AttackFireball)
	if err != nil {
		t.Fatalf("error: %s", err)
	}

	after := a.ManaPool
	assertManaChanged(t, before, after)
}

func assertManaChanged(t *testing.T, before, after int) {
    t.Helper() // marks this as a helper, so failures point to the caller
    if before == after {
        t.Fatalf("Mana didn't change: before=%d, after=%d", before, after)
    }
}
