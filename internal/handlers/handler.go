package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Overmindsas/gmphsite-backend/internal/model"
	"github.com/Overmindsas/gmphsite-backend/internal/services"
	"github.com/gin-gonic/gin"
)

type DataHandler struct {
	UserService services.DataService
	router      *gin.Engine
	db          *sql.DB
}

func NewUserHandler(userService services.DataService) DataHandler {
	return DataHandler{
		UserService: userService,
		router:      gin.New(),
		db:          &sql.DB{},
	}
}

func (u *DataHandler) GetData(c *gin.Context) {
	var data *model.Data
	r, err := http.Get("https://favqs.com/api/qotd")
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal([]byte(body), &data)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(string(body))
	if err := c.BindJSON(&body); err != nil {
		fmt.Println(err)
	}

	//fmt.Printf("%T", data.Quote.Body)
	u.UserService.GetData(
		data.QotdDate,
		data.Quote.Id,
		data.Quote.Dialoge,
		data.Quote.Private,
		data.Quote.Tags,
		data.Quote.URL,
		data.Quote.FavoritesCount,
		data.Quote.UpvotesCount,
		data.Quote.DownvotesCount,
		data.Quote.Author,
		data.Quote.AuthorPermalink,
		data.Quote.Body,
	)
	//u.UserService.GetData()
	c.JSON(http.StatusOK, &data)
	fmt.Println(data)

}

func (u *DataHandler) GetAllData(c *gin.Context) {

	a := u.UserService.GetAllData()
	c.JSON(200, a)
}
