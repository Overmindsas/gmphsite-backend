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

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"localhost", "5432", "postgres", "123456", "db")

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	var d *model.Data
	r, err := http.Get("https://favqs.com/api/qotd")
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal([]byte(body), &d)
	if err != nil {
		log.Fatal(err)
	}

	sqlInsert :=
		`INSERT INTO userdata
	(qotd_date,
	id,
	dialoge,
	private_field,
	tags,
	url,
	favorite_count,
	upvotes_count,
	downvotes_count,
	author,
	author_permalink,
	body)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12);`

	_, err = db.Exec(sqlInsert,
		d.QotdDate,
		d.Quote.Id,
		d.Quote.Dialoge,
		d.Quote.Private,
		d.Quote.Tags,
		d.Quote.URL,
		d.Quote.FavoritesCount,
		d.Quote.UpvotesCount,
		d.Quote.DownvotesCount,
		d.Quote.Author,
		d.Quote.AuthorPermalink,
		d.Quote.Body)

	if err != nil {
		log.Fatal(err)
	}
	//u.UserService.GetData()
	c.JSON(http.StatusOK, &d)
	fmt.Println(d)

}

func (u *DataHandler) GetAllData(c *gin.Context) {
	a := u.UserService.GetAllData()
	c.JSON(200, a)
}
