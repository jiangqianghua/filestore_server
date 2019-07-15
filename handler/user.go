package handler

import (
	"io/ioutil"
	"net/http"

	dblayer "filestore_server/db"
	"filestore_server/util"
	"fmt"
	"time"
)

const (
	pwd_salt = "*@#$#$#@"
)

// 处理用户注册请求
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		data, err := ioutil.ReadFile("./static/view/signup.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(data)
		return
	}
	r.ParseForm()
	username := r.Form.Get("username")
	passwd := r.Form.Get("password")

	if len(username) < 3 || len(passwd) < 5 {
		w.Write([]byte("Invalid parameter"))
		return
	}

	enc_passwd := util.Sha1([]byte(passwd + pwd_salt))
	suc := dblayer.UserSignup(username, enc_passwd)
	if suc {
		w.Write([]byte("SUCCESS"))
	} else {
		w.Write([]byte("FAILED"))
	}

}

// 登陆验证
func SigninHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		data, err := ioutil.ReadFile("./static/view/signin.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(data)
		return
	}
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	encPasswd := util.Sha1([]byte(password + pwd_salt))

	//1 校验用户名和密码
	pwdChecked := dblayer.UserSignin(username, encPasswd)
	if !pwdChecked {
		w.Write([]byte("FAILED"))
		return
	}

	// 生成访问的token
	token := GenToken(username)
	dblayer.UpdateToken(username, token)

	// 重定向到主页
	w.Write([]byte("http://" + r.Host + "/static/view/home.html"))
	// data, _ := ioutil.ReadFile("./static/view/home.html")
	// w.Write(data)
}

// 生成token
func GenToken(username string) string {
	// 40字符的token,规则是 md5(username + timestamp + token_salt) + timestamp[:8]
	ts := fmt.Sprintf("%x", time.Now().Unix())
	tokenPrefix := util.MD5([]byte(username + ts + "_tokensalt"))
	return tokenPrefix + ts[:8]
}
