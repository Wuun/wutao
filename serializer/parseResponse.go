package serializer

//ParseResponse is parse response
type ParseResponse struct {
	Code    int
	Content interface{}
}

//CustomerInformation is the information we get from the page.
type CustomerInformation struct {
	ShopID string `json:"shop_id"`
	Status string `json:"status"`
}
