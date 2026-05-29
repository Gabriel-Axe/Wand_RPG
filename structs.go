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

type Unit struct {
    ID          int    `json:"id"`
    Name        string `json:"name"`
    Health      int    `json:"health"`
		ManaPool int
    IsDefending bool   `json:"is_defending"`
		Items []ItemHolder
		Types []UnitType `json: "unit_types"`
		Attacks []Attack `json:"attacks"`
		Effects []StatusEffect `json:"effects"`
}

type UnitType struct {
	Name string             `json:"name"`
	StrongAgainst *UnitType `json:"-"`
	WeakAgainst *UnitType   `json:"-"`
}

type Attack struct {
	Name string
	Damage int
	ManaUsage int
	Effect Effect `json:"-"`
}

type Item struct {
	Name string
	Type string // wand, potion, weapon, poison...
}

type Wand struct {
	Item
	Spells []Attack
	ManaPool int
	RechargeRate int
}

type Potion struct {
	Item
	Effects []StatusEffect
	Count int
}

type StatusEffect struct {
    Type     string `json:"type"`
    Damage   int    `json:"damage"`
		Slowdown int
    Duration int    `json:"duration"`
}

type PoisonEffect struct {
	DamagePerTurn int
	Turns int
}

type FireEffect struct {
	DamagePerTurn int
	Turns int
}

type FreezeEffect struct {
	Slowdown int
	Turns int
}
