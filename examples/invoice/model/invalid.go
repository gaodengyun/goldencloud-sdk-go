package model

func Invalid() (routerAddress string, postData map[string]interface{}) {

	routerAddress = "/tax-api/invoice/invalid/v1"

	postData = map[string]interface{}{
		"seller_taxpayer_num": "500102010004038",
		"order_sn":            "",
		"order_id":            "gp_20190828115454_742_11296_31131",
		"name":                "fff",
		"is_red":              0,
	}

	return routerAddress, postData
}
