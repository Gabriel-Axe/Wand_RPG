package main

type unit struct {
	id int
	name string
	health int
	damage int
}

type player struct {
	id int
	name string
	team []*unit
}

type game struct {
	players []*player // TODO: Make this hold exactly 2
	turn int
	// NOTE: I still dont know why i need to put these Uppercased
}

func make_goblin() *unit {
	return &unit{name: "Goblin", health: 70, damage: 10}
}

func make_werewolf() *unit {
	return &unit{name: "Werewolf", health: 130, damage: 30}
}

func make_elven() *unit {
	return &unit{name: "Elven", health: 100, damage: 20}
}
