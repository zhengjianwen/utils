package context

import (
	"github.com/zhengjianwen/utils/http/cookie"
	"net/http"
)

const (
	CONTEXT_USER_ID   = "_userid_"
	CONTEXT_USER_NAME = "_username_"
)


func User(r *http.Request) (int64, string) {
	return cookie.ReadUser(r)
}
