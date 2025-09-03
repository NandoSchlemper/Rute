package unitary_test

import (
	"log"
	"testing"

	"github.com/NandoSchlemper/rute/domain/entities"
)

// 1. Criar player
// 2. Testar se tudo está funcionando como deveria
// 3. vida boa vida boa

func TestCriandoJogador(t *testing.T) {
	skillHuman1 := entities.NewSkill(
		1,
		1,
		entities.CHARISMA,
		0.2,
		entities.NONE,
		entities.ONE,
		"Charme",
		entities.RACE,
	)
	log.Println("Created skill:", skillHuman1)

	attributeHuman := entities.NewAttributeSystem(2, 0, 0, 0, 1, 0)
	log.Println("Created attribute system for human:", attributeHuman)

	var skillsHuman []entities.ISkill
	skillsHuman = append(skillsHuman, skillHuman1)
	log.Printf("Appended skill to skills slice: total %d skills\n", len(skillsHuman))

	skillSystem := entities.NewSkillSystem()
	lvlSystem := entities.NewLevel(1)
	attribute := entities.NewAttributeSystem(0, 0, 0, 0, 0, 0)
	race := entities.NewRaceSystem("Humano", skillsHuman, attributeHuman)
	player := entities.NewPlayer("P1", lvlSystem, attribute, race, skillSystem)
	log.Println("Created new player:", player)
}
