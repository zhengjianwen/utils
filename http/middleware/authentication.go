package middleware

import (
	"net/http"

	"context"
	ctx "github.com/zhengjianwen/utils/http/context"
	"github.com/zhengjianwen/utils/http/cookie"
	"github.com/zhengjianwen/utils/http/render"
	"github.com/toolkits/web/errors"
)

func Authentication(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	userid, username := cookie.ReadUser(r)
	if userid <= 0 {
		panic(errors.NotLoginError())
	}

	r = r.WithContext(context.WithValue(r.Context(), ctx.CONTEXT_USER_ID, userid))
	r = r.WithContext(context.WithValue(r.Context(), ctx.CONTEXT_USER_NAME, username))

	render.Put(r, "userid", userid)
	render.Put(r, "username", username)

	next(rw, r)
}
