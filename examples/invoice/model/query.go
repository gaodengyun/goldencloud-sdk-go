package model

func Query() (routerAddress string, postData map[string]interface{}) {

	routerAddress = "/tax-api/invoice/query/v1"

	postData = map[string]interface{}{
		"seller_taxpayer_num": "500102010004038",
		"order_sn":            "",
		"order_id":            "gp_20190828115454_742_11296_31131",
		"is_red":              0,
	}

	return routerAddress, postData
}
