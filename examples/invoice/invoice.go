package main

import (
	"fmt"

	"github.com/gaodengyun/goldencloud-sdk-go/examples/invoice/model"
	"github.com/gaodengyun/goldencloud-sdk-go/goldencloud/common"
	"github.com/gaodengyun/goldencloud-sdk-go/goldencloud/invoice"
)

func main() {

	sdk := invoice.NewSdk(common.HMAC_SHA256, "2ecde22f58f23f0b4e11", "a1c8484c9af809bc42d322a715d097d1", "", "test")

	routerAddress, postData := model.Blue() //发票开具
	//routerAddress, postData := model.Red() //发票红冲
	//routerAddress, postData := model.Invalid() //发票作废
	//routerAddress, postData := model.Print() //发票打印
	//routerAddress, postData := model.Query() //发票查询
	//routerAddress, postData := model.WaitOpen() //查询下一张待开发票

	r, err := sdk.HttpPost("https://apigw-test.goldentec.com", routerAddress, postData)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(r))
	}
}
