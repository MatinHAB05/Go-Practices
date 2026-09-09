package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type InfoHandler struct {
}

func (ih *InfoHandler) New() {

}

// Inspiration godoc
//
//	@Summary		Redirect to inspiration source
//	@Description	Redirects the client to the Quera problem page that inspired this project
//	@Tags			Info
//	@Produce		json
//	@Success		301	{string}	string	"Moved Permanently"
//	@Router			/inspiration [get]
func (ih *InfoHandler) Inspiration(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "https://quera.org/problemset/220672")
}

// Info godoc
//
//	@Summary		Get developer information
//	@Description	Returns basic contact and professional info about the developer
//	@Tags			Info
//	@Produce		json
//	@Success		200	{object}	map[string]string	"Developer info details"
//	@Router			/info [get]
func (ih *InfoHandler) Info(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"1-Author":   "Matin Hasanali Baki",
		"2-NickName": "Matin HAB",
		"3-Job":      "None of Your Business!",
		"4-Telegram": "blah-blah",
		"5-Else?":    "Not Else",
		"6-Bye?":     "Bye!",
	})
}
