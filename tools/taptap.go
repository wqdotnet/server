package tools

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

type TapTapJson struct {
	Data struct {
		Name    string `json:"name"`
		Avatar  string `json:"avatar"`
		Gender  string `json:"gender"`
		Openid  string `json:"openid"`
		Unionid string `json:"unionid"`
	} `json:"data"`
	Success bool `json:"success"`
}

func HmacSHA1(key string, data string) string {
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(data))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))

}

//客户端 ID 	 SDK 获取的 mac_key
func TapTapLogin(client_id, access_token, mac_key string) *TapTapJson {
	taptapjson := &TapTapJson{}

	url := fmt.Sprintf("https://openapi.taptap.com/account/profile/v1?client_id=%v", client_id)

	//当前时间戳
	ts := time.Now().Unix()
	// 随机数，正式上线请替换
	nonce := RandString(5)

	// # 请求方法
	METHOD := "GET"
	// # 请求地址 (带 query string)
	REQUEST_URI := fmt.Sprintf("/account/profile/v1?client_id=%v", client_id)
	// # 请求域名
	REQUEST_HOST := "openapi.taptap.com"
	mac := HmacSHA1(mac_key, fmt.Sprintf("%v\n%v\n%v\n%v\n%v\n443\n\n", ts, nonce, METHOD, REQUEST_URI, REQUEST_HOST))

	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("http :", err)
	}
	reqest.Header.Add("AUTHORIZATION", fmt.Sprintf("MAC id=\"%v\",ts=\"%v\",nonce=\"%v\",mac=\"%v\"", access_token, ts, nonce, mac))

	client := &http.Client{}
	response, err := client.Do(reqest)
	if err != nil {
		fmt.Println("http get:", err)
	}
	defer response.Body.Close()

	respbyte, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(respbyte, taptapjson)
	return taptapjson
}

func RandString(len int) string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}
