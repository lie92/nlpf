package models

type User struct {
	Firstname   string
	Lastname    string
	Email	    string
	Password	string
	admin 		bool
	UID       	int
	Phone		string
}