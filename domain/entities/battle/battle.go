package battle

// count rounds
// set turns based on speed
// each turn each entity have to choose actions
// (Attack, Run)
// IF attack: Mostra uma tela com as skills do usuário e oq ele pode ou não utilizar
// IF run: Calcula a change do usuário fugir baseando-se na velocidade tbm

type IBattle interface{}

type Battle struct{}

func NewBattle() IBattle {
	return &Battle{}
}
