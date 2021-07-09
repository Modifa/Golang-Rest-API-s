package models

type UserRegisted struct {
	Username string `db:"username"`
	Cell_num string `db:"cell_num"`
	Password string `db:"password"`
	Addedby  int64  `db:"addedby"`
}

type LoggedUser struct {
	LoggedUser string `db:"fn_loginjson"`
}

type Login struct {
	Cell_num string `json:"cell_num"`
	Password string `json:"password"`
}
