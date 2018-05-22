package models

type Sizer struct {
	Tamanho []string `json:"sizer"`
}

type Teste struct {
	Tamanho []struct{}
}

type Price struct {
	Promotion float64 `json:"promotion"`
	Price     float64 `json:"price"`
}

type Produtos struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Details     string  `json:"details"`
	Stock       float64 `json:"stock"`
	Type        string  `json:"type"`
	Sizer       Sizer
	Price       Price
}

type Pedidos struct {
	ID            int      `json:"id"`
	ClientID      int      `json:"clientID"`
	StatusPayment string   `json:"statusPayment"`
	StatusOrder   string   `json:"statusOrder"`
	Shopping      []string `json:"shopping"`
}
