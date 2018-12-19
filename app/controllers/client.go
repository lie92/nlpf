package controllers

import (
	"database/sql"
	"fmt"
	"github.com/revel/revel"
	"github.com/revel/revel/cache"
	"io"
	"nlpf/app"
	"nlpf/app/models"
	"nlpf/app/routes"
	"os"
	"strconv"
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

		var stored int

		_ = cache.Get("id", &stored)

		sqlStatement := `SELECT * FROM tags WHERE userId=$1`


		fmt.Printf("the is is is is : " + strconv.Itoa(stored) + "\n")

		rows, err := app.Db.Query(sqlStatement, stored)
		checkErr(err)
		var total int64 = 0

		var tags []models.Tag
		for rows.Next() {
			var tag models.Tag

			err = rows.Scan(&tag.Id, &tag.UserId, &tag.Time, &tag.Place, &tag.Pending, &tag.Accepted, &tag.Reason, &tag.Price, &tag.Phone,
				&tag.Motif, &tag.Orientation)
			checkErr(err)
			total += tag.Price.Int64
			tags = append(tags, tag)
		}

		return c.Render(tags, total)
	} else {
		return c.Redirect(routes.App.HTTP403())
	}
}

func (c Client) Facture() revel.Result {
	sqlStatement := `SELECT * FROM tags WHERE userId=$1`

	var stored int

	_ = cache.Get("id", &stored)

	rows, err := app.Db.Query(sqlStatement, stored)
	checkErr(err)
	var total int64 = 0

	var tags []models.Tag
	for rows.Next() {
		var tag models.Tag

		err = rows.Scan(&tag.Id, &tag.UserId, &tag.Time, &tag.Place, &tag.Pending, &tag.Accepted, &tag.Reason, &tag.Price, &tag.Phone,
			&tag.Motif, &tag.Orientation)
		checkErr(err)
		total += tag.Price.Int64
		tags = append(tags, tag)
	}

	return c.Render(tags, total)
}

func (c Client) Modify(id int) revel.Result {

	//TODO => check si le mec à le droit (si le tag existe et qu'il lui appartient, qu'il est pas dj accepté/refusé, etc...)

	var stored int

	_ = cache.Get("id", &stored)

	sqlStatement := `SELECT * FROM tags WHERE userId=$1 AND id=$2`

	rows, err := app.Db.Query(sqlStatement, stored, id)
	checkErr(err)

	var tag models.Tag
	for rows.Next() {

		err = rows.Scan(&tag.Id, &tag.UserId, &tag.Time, &tag.Place, &tag.Pending, &tag.Accepted, &tag.Reason, &tag.Price, &tag.Phone,
			&tag.Motif, &tag.Orientation)
		checkErr(err)

	}

	return c.Render(tag)
}

func (c Client) ModifyDemande(address, motif, phone, orientation string, id int) revel.Result {
	sqlStatement := `UPDATE public.tags
	SET place=$1, phone=$2, motif=$3, orientation=$4
	WHERE id = $5`

	_, err := app.Db.Exec(sqlStatement, address, phone, motif, orientation, id)
	if err != nil {
		panic(err)
	}

	return c.Redirect(routes.Client.Index())

}

func (c Client) ProcessDemande(address, motif, phone, orientation string) revel.Result {

	//TODO ==> Ici pour l'upload de file.
	//TODO ==> On doit upload le file, save le file (surement en local) puis save le tag dans la bdd.
	//TODO ==> Faudra donc modifier la bdd avec un champ photo
	//TODO ==> Faudra aussi mettre la photo dans le "tagprofile" et mettre l'upload dans modifier (ModifyDemande)
	//TODO ==> pour l'instant occupe toi que de ça, le front passe après (même si il est pas très beau)
	//TODO ==> Gl :)
	//TODO ==> Ps: exemple sur: https://github.com/revel/examples/tree/master/upload
	//TODO ==> mais il marche pas donc si t'arrive à le faire marcher on a gagné.

	var stored int

	_ = cache.Get("id", &stored)

	sqlStatement := `INSERT INTO tags (userId, place, pending, accepted, motif, phone, time, orientation)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id`

	var id int
	err := app.Db.QueryRow(sqlStatement, stored, address, true, sql.NullBool{false, false}, motif, phone, "01/01/01", orientation).Scan(&id)
	if err != nil {
		panic(err)
	}

	file := c.Params.Files["pic"][0]

	f, err := os.OpenFile("./public/img/"+strconv.Itoa(id)+".png", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return c.Redirect("/")
	}

	defer f.Close()

	f2, err := file.Open()
	if err != nil {
		fmt.Println(err)
		return c.Redirect("/")
	}
	defer f2.Close()

	io.Copy(f, f2)

	defer f2.Close()

	return c.Redirect(routes.Client.Index())
}

func (c Client) DeleteDemande(id int) revel.Result {
	//TODO => check si il a les droits
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
	var m int = int(today.Month())
	d := today.Day()
	return c.Render(y, m, d)
}

func (c Client) Tag(id int) revel.Result {
	//TODO => check si le mec à le droit (si le tag existe et qu'il lui appartient)

	var stored int

	_ = cache.Get("id", &stored)

	sqlStatement := `SELECT * FROM tags WHERE userId=$1 AND id=$2`

	rows, err := app.Db.Query(sqlStatement, stored, id)
	checkErr(err)

	var tag models.Tag
	for rows.Next() {

		err = rows.Scan(&tag.Id, &tag.UserId, &tag.Time, &tag.Place, &tag.Pending, &tag.Accepted, &tag.Reason, &tag.Price, &tag.Phone,
			&tag.Motif, &tag.Orientation)
		checkErr(err)

	}

	return c.Render(tag)
}
