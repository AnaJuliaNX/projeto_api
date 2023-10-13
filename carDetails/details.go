package cardetails

type CarDetails struct {
	ID_carro     int64  `json:"id_carro"`
	Carro        string `json:"carro"`
	Fabricante   string `json:"fabricante"`
	PrimeiroNome string `json:"primeiro_nome"`
	UltimoNome   string `json:"ultimo_nome"`
}
