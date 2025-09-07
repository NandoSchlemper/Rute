package tests

import (
	"testing"

	"github.com/NandoSchlemper/rute/domain/entities"
)

func TestGet(t *testing.T) {
	AttributeSystem := entities.NewAttributeSystem()

	if AttributeSystem == nil {
		t.Errorf("Sistema de atributos não renderizado")
	}
	expected := 0
	AttributeValue := AttributeSystem.Get(entities.CHARISMA)

	if AttributeValue != expected {
		t.Errorf("Valor do atributo diferente do que deveria ser %d", AttributeValue)
	}
}
