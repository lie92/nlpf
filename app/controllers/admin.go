package controllers

import (
	"fmt"
	"github.com/kataras/go-mailer"
	"github.com/revel/revel"
	"nlpf/app"
	"nlpf/app/models"
	"nlpf/app/routes"
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


func (c Admin) BanAction(id_account int) revel.Result {
	fmt.Println("banning")
	sqlStatement := `UPDATE public.users
	SET blacklist=true
	WHERE id = $1`

	_, err := app.Db.Exec(sqlStatement, id_account)
	if err != nil {
		panic(err)
	}

	return c.Redirect(routes.Admin.Ban())
}


func (c Admin) UnbanAction(id_account int) revel.Result {
	fmt.Println("unbanning")
	sqlStatement := `UPDATE public.users
	SET blacklist=false
	WHERE id = $1`

	_, err := app.Db.Exec(sqlStatement, id_account)
	if err != nil {
		panic(err)
	}

	return c.Redirect(routes.Admin.Ban())
}



func (c Admin) Ban() revel.Result { //aux
	//acceptOffer(tag, "", "")
	if (!isAuth() || !isAdmin()) {
		return c.Redirect(routes.App.HTTP403())
	}

	sqlStatement := `SELECT * FROM users` /*WHERE time>$1`*/

	rows, err := app.Db.Query(sqlStatement)//, time.Now)

	var users []models.User

	for rows.Next() {
		var user models.User

		err = rows.Scan(&user.UID, &user.Firstname, &user.Lastname, &user.Email, &user.Password, &user.Admin, &user.Phone, &user.Blacklist)
		checkErr(err)
		users = append(users, user)
	}
	return c.Render(users)
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
	booking := date + " " + hour;//+ strings.Split(hour, " ")[0] + ":00"
	fmt.Println(booking)
	fmt.Println("is printed")
	fmt.Println(price_rdv);
	sqlStatement := `
	UPDATE tags 
	SET accepted = true, pending = false, time=$2, price=$3
	WHERE id = $1; `

	_, err := app.Db.Exec(sqlStatement, id, booking, price_rdv)

	if err != nil {
		panic(err)
	}

	config := mailer.Config{
		Host:     "smtp.gmail.com",
		Username: "nlpfcorp@gmail.com",
		Password: "hahahaha11.",
		FromAddr: "nlpfcorp@gmail.com",
		Port:     587,
		UseCommand: false,
	}

	sender := mailer.New(config)
	subject := "Hello subject"
	content := `<h1>Hello</h1> <br/><br/> <span style="color:black"> Votre demande à été accepté </span>`
	to := []string{"mohamed.bennis@epita.fr"}
	err = sender.Send(subject, content, to...)

	if err != nil {
		println("error while sending the e-mail: " + err.Error())
	}

	fmt.Println("acceptation demande tag")
}

func refuseOffer(id int, reason string) {
	fmt.Println(id)
	fmt.Println(" and reason is ")
	fmt.Println(reason)
	sqlStatement := `
	UPDATE tags 
	SET pending = false, accepted = false, reason = $2
	WHERE id = $1; `

	_, err := app.Db.Exec(sqlStatement, id, reason)

	if err != nil {
		panic(err)
	}

	config := mailer.Config{
		Host:     "smtp.gmail.com",
		Username: "nlpfcorp@gmail.com",
		Password: "hahahaha11.",
		FromAddr: "nlpfcorp@gmail.com",
		Port:     587,
		UseCommand: false,
	}

	sender := mailer.New(config)
	subject := "Hello subject"
	content := `<h1>Hello</h1> <br/><br/> <span style="color:black"> Votre demande à été refusé </span>`
	to := []string{"mohamed.bennis@epita.fr"}
	err = sender.Send(subject, content, to...)

	if err != nil {
		println("error while sending the e-mail: " + err.Error())
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