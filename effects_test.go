package main

import "testing"

func TestApplyEffect(t *testing.T) {
	g := QuickGameSetup()
	d := g.Defender.Team[0]
	a := g.Attacker.Team[0]

	before := d.Health
	logGameState(t, g)

	AttackFireball.Effect.Apply(d, a)

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
