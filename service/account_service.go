package service

import (
	"bytes"
	"encoding/json"
	"golang-demo/data/request"
	"golang-demo/helper"
	"io/ioutil"
	"net/http"
)

var baseUrl = "https://devapi.lrinternal.com/"

func Sott(apiSecret string, apiKey string) (map[string]interface{}, int) {
	// LR API call start
	url := baseUrl + "/identity/v2/manage/account/sott"
	req, err := http.NewRequest("GET", url, nil)
	helper.ErrorPanic(err)
	req.Header.Add("X-LoginRadius-ApiKey", apiKey)
	req.Header.Add("X-LoginRadius-ApiSecret", apiSecret)
	res, err_1 := http.DefaultClient.Do(req)
	helper.ErrorPanic(err_1)
	defer res.Body.Close()
	// LR API call end

	body, readErr := ioutil.ReadAll(res.Body)
	helper.ErrorPanic(readErr)
	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(string(body)), &jsonMap)

	return jsonMap, res.StatusCode
}

func Register(apiKey string, registerBody *request.RegisterRequest, sott string) (map[string]interface{}, int) {
	// LR API call start
	url := baseUrl + "/identity/v2/auth/register"
	a, err_1 := json.Marshal(registerBody)
	helper.ErrorPanic(err_1)
	req, err := http.NewRequest("POST", url, bytes.NewReader(a))
	helper.ErrorPanic(err)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-LoginRadius-Sott", sott)
	q := req.URL.Query()
	q.Add("apiKey", apiKey)
	req.URL.RawQuery = q.Encode()
	res, err_2 := http.DefaultClient.Do(req)
	helper.ErrorPanic(err_2)
	defer res.Body.Close()
	// LR API call end

	body, readErr := ioutil.ReadAll(res.Body)
	helper.ErrorPanic(readErr)
	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(string(body)), &jsonMap)
	return jsonMap, res.StatusCode
}

func EmailVerify(apiKey string, verificationtoken string) (map[string]interface{}, int) {
	// LR API call start
	url := baseUrl + "/identity/v2/auth/email"
	req, err := http.NewRequest("GET", url, nil)
	helper.ErrorPanic(err)
	q := req.URL.Query()
	q.Add("apiKey", apiKey)
	q.Add("verificationtoken", verificationtoken)
	req.URL.RawQuery = q.Encode()
	res, err_1 := http.DefaultClient.Do(req)
	helper.ErrorPanic(err_1)
	defer res.Body.Close()
	// LR API call end

	body, readErr := ioutil.ReadAll(res.Body)
	helper.ErrorPanic(readErr)
	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(string(body)), &jsonMap)
	return jsonMap, res.StatusCode
}

func Login(apiKey string, LoginBody *request.LoginRequest) (map[string]interface{}, int) {
	// LR API call start
	url := baseUrl + "/identity/v2/auth/login"
	a, err_1 := json.Marshal(LoginBody)
	helper.ErrorPanic(err_1)
	req, err := http.NewRequest("POST", url, bytes.NewReader(a))
	helper.ErrorPanic(err)
	req.Header.Add("Content-Type", "application/json")
	q := req.URL.Query()
	q.Add("apiKey", apiKey)
	req.URL.RawQuery = q.Encode()
	res, err_2 := http.DefaultClient.Do(req)
	helper.ErrorPanic(err_2)
	defer res.Body.Close()
	// LR API call end

	body, readErr := ioutil.ReadAll(res.Body)
	helper.ErrorPanic(readErr)
	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(string(body)), &jsonMap)
	return jsonMap, res.StatusCode
}

func ProfileUpdate(apiKey string, ProfileBody *request.ProfileUpdateRequest, token string) (map[string]interface{}, int) {
	// LR API call start
	url := baseUrl + "/identity/v2/auth/account"
	a, err_1 := json.Marshal(ProfileBody)
	helper.ErrorPanic(err_1)
	req, err := http.NewRequest("PUT", url, bytes.NewReader(a))
	helper.ErrorPanic(err)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token)
	q := req.URL.Query()
	q.Add("apiKey", apiKey)
	req.URL.RawQuery = q.Encode()
	res, err_2 := http.DefaultClient.Do(req)
	helper.ErrorPanic(err_2)
	defer res.Body.Close()
	// LR API call end

	body, readErr := ioutil.ReadAll(res.Body)
	helper.ErrorPanic(readErr)
	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(string(body)), &jsonMap)
	return jsonMap, res.StatusCode
}
