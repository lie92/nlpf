package controllers

import (
	"fmt"
	"github.com/revel/revel"
	"nlpf/app"
	"nlpf/app/models"
	"nlpf/app/routes"
	//	"nlpf/app/routes"
	"time"
)

type Admin struct {
	*revel.Controller
}


func (c Admin) Administration(begin_date_input time.Time, end_date_input time.Time, motifrejet string, currentofferrefused int, currentofferaccepted int, date string, hour string, price_rdv float32) revel.Result {

	if (!isAuth() || !isAdmin()) {
		return c.Redirect(routes.App.HTTP403())
	}

	sqlStatement := `SELECT * FROM tags` /*WHERE time>$1`*/

	rows, err := app.Db.Query(sqlStatement)//, time.Now)

	var tags []models.Tag

	for rows.Next() {
		var tag models.Tag

		err = rows.Scan(&tag.Id, &tag.UserId, &tag.Time, &tag.Place, &tag.Pending, &tag.Accepted, &tag.Reason, &tag.Price, &tag.Phone,
			&tag.Motif, &tag.Orientation)
		checkErr(err)
		const longForm = "Jan 2, 2006 at 3:04pm (MST)"
		t, _ := time.Parse(longForm, "Dec 29, 2012 at 7:54pm (PST)")
		t2, _ := time.Parse(longForm, "Dec 29, 2099 at 7:54pm (PST)")
		if (begin_date_input != t && end_date_input != t2) {
			///if (tag.Time > begin_date_input && tag.Time < end_date_input) {
			if (end_date_input.Sub(begin_date_input) > 0) {
				tags = append(tags, tag)
			}
		}
		if (begin_date_input == t || end_date_input == t2 || begin_date_input == time.Time{} || end_date_input == time.Time{}) {
			tags = append(tags, tag)
		}
		/*if (begin_date_input != nil) {
			if (end_date_input != nil) {
				//if (begin_date_input > tag.Time && end_date_input < tag.Time) { // marche pas faut conv en string et recoder la conversion
				//}
			}
		} else if (end_date_input != nil) {

		} else {
		tags = append(tags, tag)
		}*/
	}

	checkErr(err)

	fmt.Println(tags);
	fmt.Println("entering admin")
	fmt.Println(begin_date_input);
	fmt.Println(end_date_input);
	if (motifrejet != "") {
		fmt.Println("curr offer is")
		fmt.Println(currentofferrefused);
		refuseOffer(currentofferrefused, motifrejet)
		var rep = "Une demande a bien été rejetée";
		return c.Render(tags, rep);
	}
	if (&currentofferaccepted != nil && currentofferaccepted != 0) {
		//autre alternative : on ne veut plus se servir du currentoffer mais juste du champ date et du champ prix après
		acceptOffer(currentofferaccepted, date, hour, price_rdv)
		var rep = "Une demande a bien été acceptée"
		return c.Render(tags, rep);
	}
	return c.Render(tags)
}

func (c Admin) AcceptOffer(tag int) revel.Result { //aux
	//acceptOffer(tag, "", "")
	return c.Render()
}

func (c Admin) RefuseOffer(tag int) revel.Result { //aux
	//refuseOffer(tag, "")
	return c.Render()
}


/*func (c Admin) Administration(begin_date_input time.Time, end_date_input time.Time) revel.Result {

/*

	if isAuth() && isAdmin() {
		sqlStatement := `SELECT * FROM tags` //WHERE time>$1

		rows, err := app.Db.Query(sqlStatement) //, time.Now)

		var tags []models.Tag

		for rows.Next() {
			var tag models.Tag

			err = rows.Scan(&tag.Id, &tag.UserId, &tag.Time, &tag.Place, &tag.Pending, &tag.Accepted, &tag.Reason, &tag.Price, &tag.Phone,
				&tag.Motif)
			checkErr(err)
			tags = append(tags, tag)
		}

		checkErr(err)

		fmt.Println(err);
		fmt.Println(rows);
		fmt.Println(tags);
		fmt.Println("rows")
		fmt.Println("entering admin")
		fmt.Println(begin_date_input);
		fmt.Println(end_date_input);
		return c.Render()

	} else {
		return c.Redirect(routes.App.HTTP403())
	}
}*/

func acceptOffer(id int, date string, hour string, price_rdv float32) {
	booking := date + " " + hour
	fmt.Println(booking)
	fmt.Println(price_rdv);
	sqlStatement := `
	UPDATE tags 
	SET accepted = true, pending = false, time= $2, price=$3
	WHERE id = $1; `

	_, err := app.Db.Exec(sqlStatement, id, booking, price_rdv)

	if err != nil {
		panic(err)
	}
	fmt.Println("acceptation demande tag")
}

func refuseOffer(id int, reason string) {
	sqlStatement := `
	UPDATE tags 
	SET pending = false, accepted = false, reason = $2
	WHERE id = $1; `

	_, err := app.Db.Exec(sqlStatement, id, reason)

	if err != nil {
		panic(err)
	}
	fmt.Println("refus demande tag")
}

func blacklist(id int) {
	sqlStatement := `
	UPDATE users 
	SET blacklist=true
	WHERE id = $1; `

	_, err := app.Db.Exec(sqlStatement, id)

	if err != nil {
		panic(err)
	}
	fmt.Println("blacklist")
}

func (c Admin) Demandes () revel.Result {

	sqlStatement := `SELECT * FROM tags WHERE pending=$1`

	rows, err := app.Db.Query(sqlStatement, true)
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

func (c Admin) Details (id int) revel.Result {

	sqlStatement := `SELECT * FROM tags WHERE id=$1`

	rows, err := app.Db.Query(sqlStatement, id)
	checkErr(err)

	var tag models.Tag
	for rows.Next() {

		err = rows.Scan(&tag.Id, &tag.UserId, &tag.Time, &tag.Place, &tag.Pending, &tag.Accepted, &tag.Reason, &tag.Price, &tag.Phone,
			&tag.Motif, &tag.Orientation)
		checkErr(err)

	}

	return c.Render(tag)

}