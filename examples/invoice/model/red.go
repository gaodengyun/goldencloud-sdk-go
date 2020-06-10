package model

func Red() (routerAddress string, postData map[string]interface{}) {

	routerAddress = "/tax-api/invoice/red/v1"

	invoices := make([]map[string]interface{}, 1)

	invoice := make(map[string]interface{})
	invoice["seller_taxpayer_num"] = "111112222233333"
	invoice["callback_url"] = "http://test.feehi.com/sign/mock/invoice-callback.php"
	invoice["order_sn"] = "6555407740980870214"

	invoices[0] = invoice

	postData = map[string]interface{}{
		"invoices": invoices,
	}

	return routerAddress, postData
}
