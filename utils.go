func list_unit_stats(u unit) {
	fmt.Println("\tname:",  u.name)
	fmt.Println("\thealth:", u.health)
	fmt.Println("\tdamage:", u.damage)
}

func is_pressing_key(expected string, player_input string) bool {
	if player_input == expected { return true }
	return  false
}

func action_menu_show() {
	fmt.Println("Ação: ")
	fmt.Println("Sair: q")
	fmt.Println("Listar time inimigo: l")
	fmt.Println("Listar time: m")
	fmt.Println("Atacar o time inimigo: a")
}

func setup_game() game {
	// p1 := &player{id: 1, name: "Alexander"}
	// p2 := &player{id: 2, name: "Oliver"}

	// WARN: Making the mother of all temporary solutions Jack, cant fret over every egg

	var team_1 [3] *unit
	team_1[0] = make_goblin()
	team_1[1] = make_elven()
	team_1[2] = make_werewolf()

	team_1[0].id = 1
	team_1[1].id = 2
	team_1[2].id = 3

	var team_2 [3] *unit
	team_2[0] = make_goblin()
	team_2[1] = make_elven()
	team_2[2] = make_werewolf()

	team_2[0].id = 1
	team_2[1].id = 2
	team_2[2].id = 3

	p1 := &player{id: 1, name: "Alexander", team: team_1[:]}
	p2 := &player{id: 2, name: "Oliver", team: team_2[:]}

	g := game{players: []*player{p1, p2}, turn: 1}

	return g
}

func get_input() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func get_player_team(p *player) {
	fmt.Println("Time do jogador", p.id, ":")
	for _, unit := range (p.team) {
		list_unit_stats(*unit)
	}
}

func game_loop(g game) {

	// choose_team(g.players[0])
	// choose_team(g.players[1])

	var player_input string

	for {
		for _, player := range g.players {
			for {
				action_menu_show()

				fmt.Scan(&player_input)
				if is_pressing_key("q", player_input) {
					os.Exit(0) 
				} else if is_pressing_key("l", player_input) { 
					list_opposit_player_team(player, g)
				} else if is_pressing_key("m", player_input) { 
					list_player_team(player, g)
				} else if is_pressing_key("a", player_input) {
					attack_opposite_player_team(player, g)
					break
				} else {
					fmt.Printf("Unknow key: %s", player_input)
				}

				continue
			}
		}
		g.turn += 1
		fmt.Println("Current turn:", g.turn)
	}
}
