package handler

import (
	"fmt"
	"net/http"
)

// 拦截器
func HTTPInterceptor(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			r.ParseForm()
			username := r.Form.Get("username")
			token := r.Form.Get("token")
			fmt.Println("HTTPInterceptor:" + token + " " + username)

			if len(username) < 3 || !IsTokenValid(token) {
				//w.WriteHeader(http.StatusForbidden)
				fmt.Println("HTTPInterceptor: 校验失败")
				http.Redirect(w, r, "/static/view/signin.html", http.StatusFound)
				return
			}
			h(w, r)
		})
}
