# Wand RPG

## Sobre

Este é um projeto de um backend para um RPG baseado em [Noita](https://store.steampowered.com/app/881100/Noita/).

## Objetivo

O principal objetivo do projeto é implementar a complexa e inteligente mecânica de varinhas do jogo em um cenário de RPG em turnos para efeitos de combate e suporte.

## Estrutura do Projeto

O projeto é organizado nestes arquivos:

```txt
.
├── attacks.go: definições de ataques
├── combat.go: funções e structs relacionadas ao combate
├── combat_test.go: testes de combate
├── effects.go: interfaces e structs relacianadas a efeitos de turno
├── go.mod: definição do módulo
├── http.go: coisas relacianadas aos endpoints do servidor http
├── items.go: funções e structs relacianadas aos itens e varinhas
├── main.go: entry point do projeto
├── README.md: "documentação"
├── state.go: structs utilizadas para guardar o atual estado do jogo e dos jogadores
├── structs.go: structs que ainda não tem casa própria (como `combat.go`)
└── utils.go: funções utilizadas somente durante o desenvolvimento
```

Esta estrutura ainda esta em organização, visto que muitas coisas são ambíguas atualmente (`Unit` pode cair tanto em `combat.go` quanto em `state.go`).

## Utilização

Primeiramente, garanta que você tem [Go](https://go.dev/) instalado em seu computador.

```bash
git clone git@github.com:Gabriel-Axe/Wand_RPG.git
cd Wand_RPG
go run .
```
