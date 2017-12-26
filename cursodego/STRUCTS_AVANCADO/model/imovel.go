package model


type Imovel struct {
	X 	  int `json:"coordenada_x"`
	Y 	  int `json:"coordenada_y"`
	Nome  string `json:"nome"`
	valor int    
}
//Sete valor define o valor do imovel
func (i *Imovel) SetValor(valor int) {
	i.valor = valor
}
//Retorna o valor do imovel
func (i *Imovel) GetValor() int {
	return i.valor
}