package model

func WaitOpen() (routerAddress string, postData map[string]interface{}) {

	routerAddress = "/tax-api/invoice/wait-open/v1"

	postData = map[string]interface{}{
		"taxpayer_num":      "20181010000000865115391398242493",
		"machine_no":        "125523523523",
		"invoice_type_code": "026",
	}

	return routerAddress, postData
}
