package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"time"
)

type Userinfo struct {
	Action     string `json:"action"`     //类型
	Username   string `json:"username"`   //账号
	Password   string `json:"password"`   //密码
	Questionid int    `json:"questionid"` //登录问题
	Answer     string `json:"answer"`     //登录答案
	Serverkey  string `json:"serverkey"`
	Signsubmit string `json:"signsubmit"`
	Formhash   string `json:"formhash"`
	Cookie     string
}

func T00ls_Go(u *Userinfo) *Userinfo {
	jar, err := cookiejar.New(nil)
	client := &http.Client{
		Jar: jar,
	}
	skey := u.Serverkey
	md5passwd := md5.Sum([]byte(u.Password)) //将密码进行md5加密
	u.Password = fmt.Sprintf("%x", md5passwd)
	want := url.Values{
		"action":   {u.Action},
		"username": {u.Username},
		"password": {u.Password},
		"answer":   {u.Answer},
	}
	//fmt.Println(want)
	want.Add("questionid", fmt.Sprintf("%v", u.Questionid))
	request, err := client.PostForm("https://www.t00ls.net/login.json",
		want)
	if err != nil || request.Status != "200 OK" {
		fmt.Println("[info] 土司好像挂了")
		Sendmsg(u.Serverkey, "土司好像挂了,一小时后会尝试重新打卡."+fmt.Sprintf("%d", time.Now().Unix()))
		time.Sleep(1 * time.Hour) //延时一小时，再执行签到。
		T00ls_Go(u)
	}
	//fmt.Println(request.Body.Read(b))
	defer request.Body.Close()
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		fmt.Println(err)
	}
	c := request.Cookies()
	var sid *http.Cookie = c[0]
	var ck *http.Cookie = c[2]
	status := fmt.Sprintf("%s", gjson.Get(string(body), "status"))
	formhash := fmt.Sprintf("%s", gjson.Get(string(body), "formhash"))
	if status != "success" {
		fmt.Println("[info] 登录失败！")
	}
	fmt.Println("[info] 登录成功!")
	h := &Userinfo{}
	h.Formhash = formhash
	h.Signsubmit = "true"
	h.Cookie = fmt.Sprintf("%s", ck) + fmt.Sprintf("%s", sid)
	h.Serverkey = skey
	Sign(h, client) //签到
	return h        //登录成功返回hash
}
func Sign(h *Userinfo, client *http.Client) { //签到函数
	want := url.Values{
		"signsubmit": {h.Signsubmit},
		"formhash":   {h.Formhash},
	}
	request, err := client.PostForm("https://www.t00ls.net/ajax-sign.json", want)
	if err != nil || request.Status != "200 OK" {
		fmt.Println("[info] 土司好像挂了")
	}
	//fmt.Println(request.Body.Read(b))
	defer request.Body.Close()
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(string(body))
	status := fmt.Sprintf("%s", gjson.Get(string(body), "status"))
	message := fmt.Sprintf("%s", gjson.Get(string(body), "message"))
	if status == "success" {
		fmt.Println("[info] 签到成功！")
		Sendmsg(h.Serverkey, "签到成功！"+fmt.Sprintf("%d", time.Now().Unix()))
	} else if message == "alreadysign" {
		fmt.Println("[info] 当前账号已签到！")
		Sendmsg(h.Serverkey, "当前账号已签到！"+fmt.Sprintf("%d", time.Now().Unix()))
	} else {
		fmt.Println("[info] 签到失败~")
	}
}
func Sendmsg(key string, msg string) {
	//server酱发送信息
	url1 := "https://sc.ftqq.com/" + key + ".send?text=" + url.QueryEscape(msg)

	resp, err := http.Get(url1)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	errmsg := fmt.Sprintf("%s", gjson.Get(string(body), "errmsg"))
	//fmt.Println(string(body))
	if errmsg == "success" {
		fmt.Println("[info] 信息推送成功！")
	} else {
		fmt.Println("[info] 信息推送失败！")
	}

}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("[info] 请输入配置文件地址，运行案例:Tubi.exe config.json (不行的话，请加绝对路径)")
	} else {
		filename := os.Args[1]
		file, _ := os.Open(filename)
		defer file.Close()
		decoder := json.NewDecoder(file)
		conf := &Userinfo{}
		err := decoder.Decode(&conf)
		if err != nil {
			fmt.Println("Error:", err)
		}
		if conf.Action != "login" {
			fmt.Println("[info] Action配置必须是login!")
			os.Exit(1)
		} else if conf.Username == "" || conf.Password == "" || conf.Answer == "" {
			fmt.Println("[info] 请认证填写config.json配置文件~")
		} else {
			fmt.Println("[info] 配置读取成功！ by: Try")
			fmt.Println("[info] By T00ls.Net；")
			T00ls_Go(conf) //开始执行
		}
	}
}
