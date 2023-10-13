package tipos

type Carros struct {
	ID         int64  `json:"id"`
	Carro      string `json:"carro"`
	Fabricante string `json:"fabricante"`
}
type CarData struct {
	Carros `json:"carros"`
}

type Compradores struct {
	ID           int64  `json:"id"`
	PrimeiroNome string `json:"primeiro_nome"`
	UltimoNome   string `json:"ultimo_nome"`
}

type CompradoresData struct {
	Compradores `json:"compradores"`
}
