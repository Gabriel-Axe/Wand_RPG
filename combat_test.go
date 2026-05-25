package main

import "testing"

func TestMakeAttack(t *testing.T) {
	g := QuickGameSetup()
	d := g.Defender.Team[0]

	before := d.Health
	logGameState(t, g)
	
	err := MakeAttack(g, 0, 0, 0)
	if err != nil {
		t.Fatalf("error: %s", err)
	}

	after := d.Health
	
	logGameState(t, g)
	assertHealthChanged(t, before, after)
	if before == after {
		t.Fatalf("The health of defender before (%d) and after (%d) is the same", before, after)
	}
}

func TestToggleDefend(t *testing.T) {
	g := QuickGameSetup()
	logGameState(t, g)

	d := g.Defender.Team[0]
	ToggleDefend(g, 0)

	before := d.Health
	
	err := MakeAttack(g, 0, 0, 0)
	if err != nil {
		t.Fatalf("error: %s", err)
	}

	after := d.Health
	
	t.Logf("Defender health after: %d", d.Health)
	assertHealthChanged(t, before, after)
}

func assertHealthChanged(t *testing.T, before, after int) {
    t.Helper() // marks this as a helper, so failures point to the caller
    if before == after {
        t.Fatalf("health didn't change: before=%d, after=%d", before, after)
    }
}

func logGameState(t *testing.T, g *Game) {
    t.Helper()
    t.Logf("Turn: %d", g.Turn)
    for i, u := range g.Attacker.Team {
        t.Logf("Attacker[%d]: health=%d, defending=%v", i, u.Health, u.IsDefending)
    }
    for i, u := range g.Defender.Team {
        t.Logf("Defender[%d]: health=%d, defending=%v", i, u.Health, u.IsDefending)
    }
}
