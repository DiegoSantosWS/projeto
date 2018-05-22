package model

//DadosCep Recebe dados em formato json
type DadosCep struct {
	Bairro     string     `json:"bairro"`
	Cidade     string     `json:"cidade"`
	Logradouro string     `json:"logradouro"`
	UFINFO     EstadoInfo `json:"estado_info"`
	CEP        string     `json:"cep"`
	CidadeInfo CidadeInfo `json:"cidade_info"`
	Estado     string     `json:"estado"`
}

//EstadoInfo teste teste teste
type EstadoInfo struct {
	AreaKm2    string `json:"area_km2"`
	CodigoIbge string `json:"codigo_ibge"`
	Nome       string `json:"nome"`
}

//CidadeInfo teste teste teste
type CidadeInfo struct {
	AreaKm2    string `json:"area_km2"`
	CodigoIbge string `json:"codigo_ibge"`
}
