package middlewares

import (
	"net/http"

	"Final/utils/token"

	"github.com/gin-gonic/gin"
)

func UserCheckMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		// Setelah token divalidasi, kita dapat memeriksa role dari user yang terautentikasi.
		role, err := token.ExtractUserRole(c)
		if err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		// Jika role tidak sesuai, berikan pesan error dan hentikan proses.
		if role != "user" && role != "admin" {
			c.String(http.StatusForbidden, "Access denied. Require Writer role.")
			c.Abort()
			return
		}

		// Jika role sesuai, lanjutkan proses ke handler berikutnya.
		c.Next()
	}
}

func AdminCheckMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		// Setelah token divalidasi, kita dapat memeriksa role dari user yang terautentikasi.
		role, err := token.ExtractUserRole(c)
		if err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		if role != "admin" {
			c.String(http.StatusForbidden, "Access denied. Insufficient role.")
			c.Abort()
			return
		}

		// Jika role sesuai, lanjutkan proses ke handler berikutnya.
		c.Next()
	}
}
