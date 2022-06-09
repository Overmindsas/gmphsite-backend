package services

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Overmindsas/gmphsite-backend/internal/model"
)

type UserImplimented struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) DataService {
	return &UserImplimented{db: db}
}

func (u *UserImplimented) GetData(
	qotd string,
	id int,
	dialoge bool,
	private bool,
	tags string,
	url string,
	favorite int,
	upvotes int,
	downvotes int,
	author string,
	author_permalink string,
	body string,
) {
	var d model.Data
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
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`

	_, err := u.db.Exec(sqlInsert,
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

	fmt.Println(d.Quote.Body)
	if err != nil {
		log.Fatal(err)
	}
}

func (u *UserImplimented) GetAllData() []model.Data {
	rows, err := u.db.Query("SELECT * FROM userdata")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	data := []model.Data{}

	for rows.Next() {
		d := model.Data{}
		err := rows.Scan(
			&d.QotdDate,
			&d.Quote.Id,
			&d.Quote.Dialoge,
			&d.Quote.Private,
			&d.Quote.Tags,
			&d.Quote.URL,
			&d.Quote.FavoritesCount,
			&d.Quote.UpvotesCount,
			&d.Quote.DownvotesCount,
			&d.Quote.Author,
			&d.Quote.AuthorPermalink,
			&d.Quote.Body)
		if err != nil {
			log.Fatal(err)
			continue
		}
		data = append(data, d)

	}

	return data
}
