package main

var (
	Fire = &UnitType{Name: "File"}
	Water = &UnitType{Name: "Water"}
	Flesh = &UnitType{Name: "Flesh"}
)

func init() {
	Fire.WeakAgainst = Water
	Fire.StrongAgainst = Flesh
	Flesh.WeakAgainst = Fire
	Water.StrongAgainst = Fire
}
