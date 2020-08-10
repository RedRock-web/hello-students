// @program: hello-students
// @author: edte
// @create: 2020-08-03 14:32
package handler

import (
	"github.com/gin-gonic/gin"

	"hello-students/config"
)

// setCookie 设置 cookie
func setCookie(c *gin.Context) {
	c.SetCookie(
		config.CookieConfig.Name,
		config.CookieConfig.Value,
		config.CookieConfig.MaxAge,
		config.CookieConfig.Path,
		config.CookieConfig.Domain,
		config.CookieConfig.Secure,
		config.CookieConfig.HttpOnly)
}
