package dto

type CreateItemRequestDto struct {
	Itemcode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}