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


func (c Admin) Administration(begin_date_input time.Time, end_date_input time.Time, motifrejet string, currentoffer int, isaccepted string) revel.Result {

	if (!isAuth() || !isAdmin()) {
		return c.Redirect(routes.App.HTTP403())
	}

	sqlStatement := `SELECT * FROM tags` /*WHERE time>$1`*/

	rows, err := app.Db.Query(sqlStatement)//, time.Now)

	var tags []models.Tag

	for rows.Next() {
		var tag models.Tag

		err = rows.Scan(&tag.Id, &tag.UserId, &tag.Time, &tag.Place, &tag.Pending, &tag.Accepted, &tag.Reason, &tag.Price, &tag.Phone,
			&tag.Motif)
		checkErr(err)
		tags = append(tags, tag)
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
		fmt.Println(currentoffer);
		refuseOffer(currentoffer, motifrejet)
		var rep = "Une demande a bien été rejetée";
		return c.Render(tags, rep);
	}
	if (isaccepted != "") {
		acceptOffer(currentoffer)
		var rep = "Une demande a bien été acceptée"
		return c.Render(tags, rep);
	}
	return c.Render(tags)
}

func (c Admin) AcceptOffer(tag int) revel.Result {
	acceptOffer(tag)
	return c.Render()
}

func (c Admin) RefuseOffer(tag int) revel.Result {
	//refuseOffer(tag, "FUCK YOU")
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

func acceptOffer(id int) {
	sqlStatement := `
	UPDATE tags 
	SET accepted = true, pending = false
	WHERE id = $1; `

	_, err := app.Db.Exec(sqlStatement, id)

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
