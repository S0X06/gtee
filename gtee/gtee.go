// gt3 project main.go
package gtee

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
)

func tomd5(str string) string {
	h := md5.New()
	h.Write([]byte(str)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

func randint(from float64, to float64) float64 {
	return math.Floor(rand.Float64()*(to-from+1) + from)
}

//gt3 的类
type Geetest struct {
	geetest_id    string
	geetest_key   string
	PROTOCOL      string
	API_SERVER    string
	VALIDATE_PATH string
	REGISTER_PATH string
	TIMEOUT       int
	NEW_CAPTCHA   bool
	JSON_FORMAT   int
}

//注册返回的结构
type Register_result struct {
	Challenge   string `json "challenge"`
	Success     int    `json "success"`
	Gt          string `json "gt"`
	New_captcha bool   `json "new_captcha"`
}

//验证上传的结构
type validate_data struct {
	Gt          string `json "gt"`
	Seccode     string `json "seccode"`
	Json_format string `json "json_format"`
}

func (Geetest *Geetest) Register(client_type string, ip_address string, callback func(*Register_result)) {

	surl := Geetest.PROTOCOL + Geetest.API_SERVER + Geetest.REGISTER_PATH
	u, _ := url.Parse(surl)
	q := u.Query()
	q.Set("gt", Geetest.geetest_id)
	q.Set("json_format", strconv.Itoa(Geetest.JSON_FORMAT))
	q.Set("sdk", "Node_2.1.0")
	q.Set("client_type", client_type)
	q.Set("ip_address", ip_address)
	q.Set("new_captcha", "1")
	u.RawQuery = q.Encode()
	res, err := http.Get(u.String())
	if err != nil {
		fmt.Printf("error")
		return
	}
	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		fmt.Printf("error")
		return
	}

	p := &Register_result{}
	p.Gt = Geetest.geetest_id
	p.New_captcha = Geetest.NEW_CAPTCHA
	p.Success = 0
	err = json.Unmarshal([]byte(result), p)
	if err != nil {

	}
	callback(p)

}

func (Geetest *Geetest) Validate(fallback bool, challenge string, validate string, seccode string, callback func(bool)) {
	fmt.Print("fallback+", fallback)
	if fallback {
		if tomd5(challenge) == validate {
			callback(true)
		} else {
			callback(false)
		}
	} else {
		var hash = Geetest.geetest_key + "geetest" + challenge
		if tomd5(hash) == validate {
			datas := new(validate_data)
			datas.Gt = Geetest.geetest_id
			datas.Seccode = seccode
			datas.Json_format = strconv.Itoa(Geetest.JSON_FORMAT)
			b, err := json.Marshal(datas)
			if err != nil {
				callback(false)
			}
			body := bytes.NewBuffer([]byte(b))
			res, err := http.Post(Geetest.PROTOCOL+Geetest.API_SERVER+Geetest.VALIDATE_PATH, "application/json;charset=utf-8", body)
			if err != nil {
				callback(false)
				return
			}
			result, err := ioutil.ReadAll(res.Body)
			res.Body.Close()
			if err != nil {
				callback(false)
				return
			}
			fmt.Printf("@@@ %s", result)

		} else {
			callback(false)
		}
	}

}
func NewGeetest(geetest_id string, geetest_key string) Geetest {
	return Geetest{
		geetest_id:    geetest_id,
		geetest_key:   geetest_key,
		PROTOCOL:      "http://",
		API_SERVER:    "api.geetest.com",
		VALIDATE_PATH: "/validate.php",
		REGISTER_PATH: "/register.php",
		TIMEOUT:       2000,
		NEW_CAPTCHA:   true,
		JSON_FORMAT:   1,
	}
}
