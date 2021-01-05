package invoice

import (
	"encoding/base64"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gaodengyun/goldencloud-sdk-go/goldencloud/common"
)

func NewSdk(algorithm, appkey, appsecret, privateKey, env string) *Client {
	return &Client{
		env:        env,
		appkey:     appkey,
		appsecret:  appsecret,
		algorithm:  algorithm,
		privateKey: privateKey,
	}
}

type Client struct {
	env        string
	appkey     string
	appsecret  string
	privateKey string
	algorithm  string
}

func (this *Client) HttpPost(baseUrl string, routerAddress string, post map[string]interface{}) ([]byte, error) {

	// 获取转换后的post字符串
	postData, err := common.GetPostDataWithMap(post)
	if err != nil {
		return nil, err
	}

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)                //获取时间戳
	nonce := common.GetRand(6)                                           //6位随机数
	auth, err := this.getAuth(routerAddress, postData, timestamp, nonce) //认证串
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	reqest, err := http.NewRequest("POST", baseUrl+routerAddress, strings.NewReader((postData)))
	if err != nil {
		return nil, err
	}

	reqest.Header.Add("Authorization", auth)
	reqest.Header.Add("Content-Type", "application/json")

	r, err := client.Do(reqest)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	return ioutil.ReadAll(r.Body)
}

// getAuth 获取授权字符串
func (this *Client) getAuth(routerAddress string, postData string, timestamp string, nonce string) (string, error) {

	originStr := "algorithm=" + this.algorithm + "|appkey=" + this.appkey + "|nonce=" + nonce + "|timestamp=" + timestamp + "|" + routerAddress + "|" + postData

	var signs string
	var err error
	switch this.algorithm {
	case common.HMAC_SHA256:
		signs = base64.StdEncoding.EncodeToString(common.HMAC_Sha256(originStr, this.appsecret))
	default:
		panic(errors.New("algorithm not exists"))
	}

	if err != nil {
		return "", err
	}

	auth := "algorithm=" + this.algorithm + ",appkey=" + this.appkey + ",nonce=" + nonce + ",timestamp=" + timestamp + ",signature=" + signs

	return auth, nil
}
