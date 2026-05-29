package main

type Effect interface {
	Apply(target *Unit, attacker *Unit, attack *Attack) error
	ToStatusEffect() StatusEffect
}

type ItemHolder interface {
    GetName() string
    GetType() string
		Use(target *Unit) error
}
