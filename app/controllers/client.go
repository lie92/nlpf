package controllers

import (
	"database/sql"
	"fmt"
	"github.com/revel/revel"
	"nlpf/app"
	"nlpf/app/models"
	"nlpf/app/routes"
	"time"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type Client struct {
	*revel.Controller
}

func (c Client) Index() revel.Result {


	if isAuth() && !isAdmin() {

	sqlStatement := `SELECT * FROM tags WHERE userId=$1`

	rows, err := app.Db.Query(sqlStatement, 2)
	checkErr(err)
	var total int64 = 0

	var tags []models.Tag
	for rows.Next() {
		var tag models.Tag

		err = rows.Scan(&tag.Id, &tag.UserId, &tag.Time, &tag.Place, &tag.Pending, &tag.Accepted, &tag.Reason, &tag.Price, &tag.Phone,
			&tag.Motif)
		checkErr(err)
		total += tag.Price.Int64
		tags = append(tags, tag)
	}

	return c.Render(tags, total)} else {
		return c.Redirect(routes.App.HTTP403())
	}
}

func (c Client) Facture() revel.Result {
	sqlStatement := `SELECT * FROM tags WHERE userId=$1`

	rows, err := app.Db.Query(sqlStatement, 2)
	checkErr(err)
	var total int64 = 0

	
	var tags []models.Tag
	for rows.Next() {
		var tag models.Tag

		err = rows.Scan(&tag.Id, &tag.UserId, &tag.Time, &tag.Place, &tag.Pending, &tag.Accepted, &tag.Reason, &tag.Price, &tag.Phone,
			&tag.Motif)
		checkErr(err)
		total += tag.Price.Int64
		tags = append(tags, tag)
	}

	return c.Render(tags, total)
}

func (c Client) Modify(id int) revel.Result {


	//TODO => check si le mec à le droit (si le tag existe et qu'il lui appartient, qu'il est pas dj accepté/refusé, etc...)
	sqlStatement := `SELECT * FROM tags WHERE userId=$1 AND id=$2`

	rows, err := app.Db.Query(sqlStatement, 2, id)
	checkErr(err)

	var tag models.Tag
	for rows.Next() {

		err = rows.Scan(&tag.Id, &tag.UserId, &tag.Time, &tag.Place, &tag.Pending, &tag.Accepted, &tag.Reason, &tag.Price, &tag.Phone,
			&tag.Motif)
		checkErr(err)
		fmt.Println(tag)
		fmt.Println(tag.Id)
		fmt.Println(tag.UserId)
		fmt.Println(tag.Time)
		fmt.Println(tag.Place)
		fmt.Println(tag.Pending)
		fmt.Println(tag.Accepted)
		fmt.Println(tag.Reason)
		fmt.Println(tag.Price)
		fmt.Println(tag.Phone)
		fmt.Println(tag.Motif)
		fmt.Println(tag.Pending)
		fmt.Println(".....................")

	}





	return c.Render(tag)
}

func (c Client) ModifyDemande(address, motif, phone string, id int) revel.Result {
	sqlStatement := `UPDATE public.tags
	SET place=$1, phone=$2, motif=$3
	WHERE id = $4`

	_, err := app.Db.Exec(sqlStatement, address, phone, motif, id)
	if err != nil {
		panic(err)
	}

	return c.Redirect(routes.Client.Index())

}

func (c Client) ProcessDemande(address, motif, phone string) revel.Result {

	sqlStatement := `INSERT INTO tags (userId, place, pending, accepted, motif, phone, time)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id`

	var id int
	err := app.Db.QueryRow(sqlStatement, 2, address, true, sql.NullBool{false, false}, motif, phone, "01/01/01").Scan(&id)
	if err != nil {
		panic(err)
	}

	return c.Redirect(routes.Client.Index())
	}

func (c Client) DeleteDemande(id int) revel.Result {

	sqlStatement := `DELETE FROM tags WHERE id = $1`

	_, err := app.Db.Exec(sqlStatement, id)
	if err != nil {
		panic(err)
	}

	return c.Redirect(routes.Client.Index())
}

func (c Client) Demande() revel.Result {
	today := time.Now()

	y := today.Year()
	var m int = int (today.Month())
	d := today.Day()
	return c.Render(y, m, d)
}