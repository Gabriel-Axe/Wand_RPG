package main

import "testing"

func TestMakeAttack(t *testing.T) {
	g := quick_game_setup()
	a := g.Attacker.Team[0]
	d := g.Defender.Team[0]
	
	t.Logf("Attacker defending: %v, Defender defending: %v", a.IsDefending, d.IsDefending)
	t.Logf("Attack damage: %d", a.Attacks[0].Damage)
	t.Logf("Defender health before: %d", d.Health)

	before := d.Health
	
	err := make_attack(g, 0, 0, 0)
	if err != nil {
		t.Fatalf("error: %s", err)
	}

	after := d.Health
	
	t.Logf("Defender health after: %d", d.Health)
	if before == after {
		t.Fatalf("The health of defender before (%d) and after (%d) is the same", before, after)
	}
}
