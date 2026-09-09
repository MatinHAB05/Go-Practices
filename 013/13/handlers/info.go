package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type InfoHandler struct {
}

func (i_f *InfoHandler) Inspiration(c *gin.Context) {
	c.Redirect(301, "https://quera.org/problemset/220672")
}

func (i_f *InfoHandler) Info(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"1-Author":   "Matin Hasanali Baki",
		"2-NickName": "Matin HAB",
		"3-Job":      "None of Your Business!",
		"4-Telegram": "blah-blah",
		"5-Else?":    "Not Else",
		"6-Bye?":     "Bye!",
	})
}
