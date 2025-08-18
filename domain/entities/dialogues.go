package entities

type Dialogue struct {
	Character  Character // personagem do dialogo
	Content    string    // conteudo do dialogo
	Conditions string    // Criar tipo pra condições
	Responses  string    // criar tipo pra responses
}
