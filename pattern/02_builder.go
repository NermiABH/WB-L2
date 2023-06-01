package main

import "fmt"

type KnifeBuilder interface {
	setLength()
	setAttackDamage()
	getKnife() Knife
}

func getKnifeBuilder(knifeType string) KnifeBuilder {
	if knifeType == "cheese" {
		return newCheeseKnife()
	}
	if knifeType == "meat" {
		return newMeatKnife()
	}
	return nil
}

type CheeseKnife struct {
	length       int
	attackDamage int
}

func newCheeseKnife() *CheeseKnife {
	return &CheeseKnife{}
}

func (c *CheeseKnife) setLength() {
	c.length = 10
}

func (c *CheeseKnife) setAttackDamage() {
	c.attackDamage = 5
}

func (c *CheeseKnife) getKnife() Knife {
	return Knife{
		Length:       c.length,
		AttackDamage: c.attackDamage,
	}
}

type MeatKnife struct {
	length       int
	attackDamage int
}

func newMeatKnife() *MeatKnife {
	return &MeatKnife{}
}

func (m *MeatKnife) setLength() {
	m.length = 16
}

func (m *MeatKnife) setAttackDamage() {
	m.attackDamage = 20
}

func (m *MeatKnife) getKnife() Knife {
	return Knife{
		Length:       m.length,
		AttackDamage: m.attackDamage,
	}
}

type Knife struct {
	Length       int
	AttackDamage int
}

type Director struct {
	builder KnifeBuilder
}

func newDirector(builder KnifeBuilder) *Director {
	return &Director{builder: builder}
}

func (d *Director) setBuilder(builder KnifeBuilder) {
	d.builder = builder
}

func (d *Director) buildKnife() Knife {
	d.builder.setLength()
	d.builder.setAttackDamage()
	return d.builder.getKnife()
}

func main() {
	cheeseKnifeBuilder := getKnifeBuilder("cheese")
	meatKnifeBuilder := getKnifeBuilder("meat")

	director := newDirector(cheeseKnifeBuilder)
	cheeseKnife := director.buildKnife()
	fmt.Printf("Cheese Knife: Length: %v\n", cheeseKnife.Length)
	fmt.Printf("Cheese Knife: Damage Attack: %v\n", cheeseKnife.AttackDamage)

	director.setBuilder(meatKnifeBuilder)
	meatKnife := director.buildKnife()
	fmt.Printf("Meat Knife: Length: %v\n", meatKnife.Length)
	fmt.Printf("Meat Knife: Damage Attack: %v\n", meatKnife.AttackDamage)

}
