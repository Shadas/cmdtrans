package domain

type TransBody struct {
	From        string       `json:"from"`
	To          string       `json:"to"`
	TransResult TransResults `json:"trans_result"`
}

type TransResults []TransResult

type TransResult struct {
	Src string `json:"src"`
	Dst string `json:"dst"`
}
