package model

func Print() (routerAddress string, postData map[string]interface{}) {

	routerAddress = "/tax-api/invoice/print/v1"

	postData = map[string]interface{}{
		"seller_taxpayer_num": "500102010004038",
		"order_sn":            "",
		"order_id":            "gp_20190828115454_742_11296_31131",
		"print_type":          "",
		"print_flag":          "",
		"print_mode":          "",
		"is_red":              0,
	}

	return routerAddress, postData
}
