package service

import (
	"fmt"
	"github.com/goccy/go-json"
	"log"
	"net/http"
	"wintervacation/model"
)

var config = model.Conf{
	ClientID:     "Iv1.d11b6181de5afc29",
	ClientSecret: "996a32e264ae2072d527b9560b05c9689ffbcd25",
	RedirectURL:  "http://localhost:8080/OAuth/redirect",
}

//要获得 token，必须访问Github 提供 token 的接口（一个url）

func GetTokenAuthUrl(code string) string {

	return fmt.Sprintf(
		"https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s",
		config.ClientID, config.ClientSecret, code,
	)

}

//直接访问这个url得到请求

func GetToken(url string) (*model.TokenString, error) {
	//形成请求
	r, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Printf("when new request error:%v", err)
		return nil, err
	}

	r.Header.Set("accept", "application/json")

	//发送请求并得到响应
	Client := http.Client{}
	res, err := Client.Do(r)
	if err != nil {
		log.Printf("when client error:%v", err)
		return nil, err
	}

	//将响应体解析为token并返回
	var token model.TokenString
	err = json.NewDecoder(res.Body).Decode(&token)
	if err != nil {
		log.Printf("when new decode error:%v", err)
		return nil, err
	}
	return &token, nil
}

func GetUserInfo(token *model.TokenString) (userInfo map[string]interface{}, err error) {

	//形成请求
	userInfoUrl := "https://api.github.com/user" //Github的api的地址
	req, err := http.NewRequest(http.MethodGet, userInfoUrl, nil)
	if err != nil {
		log.Printf("when new request in getting user info:%v", err)
		return
	}
	req.Header.Set("accept", "application/json")
	//请求时必须在HTTP信息里面带上令牌Authorization，格式：token 啥啥啥
	req.Header.Set("Authorization", fmt.Sprintf("token %s", token.AccessToken))

	//发送请求并获得GitHub响应
	Client := http.Client{}
	res, err := Client.Do(req)
	if err != nil {
		log.Printf("when client in getting user info:%v", err)
		return
	}

	//GitHub会响应一段数据，通过这段数据得到用户身份。
	userInfo = make(map[string]interface{}) //将响应的用户信息写入userInfo中
	err = json.NewDecoder(res.Body).Decode(&userInfo)
	if err != nil {
		log.Printf("when get user info decode error:%v ", err)
		return nil, err
	}
	return userInfo, nil

}
