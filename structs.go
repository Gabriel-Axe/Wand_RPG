package main

type unit struct {
	id int
	name string
	health int
	damage int
	is_defending bool
}

type player struct {
	id int
	name string
	is_attacker bool
	team []*unit
}

type game struct {
	attacker *player 
	defender *player 
	attacker_turn bool
	turn int
	// NOTE: I still dont know why i need to put these Uppercased
}

type AttackRequest struct {
	Unit_id *int  `json:"unit_id"`  
	Target_id *int `json:"target_id"`
}

type DefenseRequest struct {
	Unit_id *int  `json:"unit_id"`  
}
