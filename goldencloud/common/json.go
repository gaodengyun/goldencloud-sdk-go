package common

import jsoniter "github.com/json-iterator/go"

func GetPostDataWithMap(post map[string]interface{}) (string, error) {
	json := jsoniter.Config{
		MarshalFloatWith6Digits: true,
		EscapeHTML:              false,
		SortMapKeys:             true, //本身高灯平台仅要求对最外层json key进行asci码升序排序，但map是无序且随机的，所以签名和post数据均排序以保持一致
		UseNumber:               true,
		DisallowUnknownFields:   false,
		CaseSensitive:           true,
	}.Froze()

	s, err := json.Marshal(&post)
	if err != nil {
		return "", err
	}

	return string(s), nil
}
