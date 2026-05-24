package main

type Unit struct {
    ID          int    `json:"id"`
    Name        string `json:"name"`
    Health      int    `json:"health"`
    IsDefending bool   `json:"is_defending"`
		Attacks []Attack `josn:"attacks"`
}

type player struct {
    ID   int     `json:"id"`
    Name string  `json:"name"`
    Team []*Unit `json:"team"`
}

type Game struct {
	Attacker *player 
	Defender *player 
	Turn int
}

type AttackRequest struct {
	Unit_id *int  `json:"unit_id"`  
	Target_id *int `json:"target_id"`
	Attack_id *int `json:"attack_id"`
}

type DefenseRequest struct {
	Unit_id *int  `json:"unit_id"`  
}

type PlayerResponse struct {
	id int
	name string
	team []*Unit
}

type Attack struct {
	Name string
	Damage int
}
