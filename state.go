package main

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
