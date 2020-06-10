package model

func Blue() (routerAddress string, postData map[string]interface{}) {

	routerAddress = "/tax-api/invoice/blue/v1"

	postData = map[string]interface{}{
		"seller_name":          "wkbb invoice blue test",
		"seller_taxpayer_num":  "your seller_taxpayer_num",
		"seller_address":       "",
		"seller_tel":           "",
		"seller_bank_name":     "",
		"seller_bank_account":  "",
		"title_type":           1,
		"buyer_title":          "海南高灯科技",
		"buyer_taxpayer_num":   "",
		"buyer_address":        "",
		"buyer_bank_name":      "",
		"buyer_bank_account":   "",
		"buyer_phone":          "",
		"buyer_email":          "",
		"taker_phone":          "",
		"order_id":             "fc5fec2d1b6942658c009a151d4cffa5",
		"invoice_type_code":    "032",
		"callback_url":         "http://www.test.com",
		"drawer":               "TEST",
		"payee":                "TEST",
		"checker":              "TEST",
		"terminal_code":        "661234567789",
		"user_openid":          "ba9ea0bdfa1f460993c990564caab18f",
		"special_invoice_kind": "",
		"zsfs":                 "",
		"deduction":            0,
		"amount_has_tax":       998000,
		"tax_amount":           9980,
		"amount_without_tax":   8644,
		"remark":               "readme",
	}

	//项目商品明细
	items := make([]map[string]interface{}, 1)
	item := make(map[string]interface{})
	item["name"] = "海鲜真划算"
	item["tax_code"] = "your tax_code"
	item["models"] = "zyx"
	item["unit"] = "个"
	item["total_price"] = 8846
	item["total"] = "5"
	item["price"] = "17.22"
	item["tax_rate"] = 100
	item["tax_amount"] = 846
	item["discount"] = 0
	item["preferential_policy_flag"] = "0"
	item["zero_tax_flag"] = "0"
	item["vat_special_management"] = ""
	items[0] = item

	postData["items"] = items

	return routerAddress, postData
}
