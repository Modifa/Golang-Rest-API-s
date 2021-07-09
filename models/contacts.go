package models

// type ListContacts struct {
// 	UserId   int64  `db:"user_id"`
// 	UserName string `db:"username_"`
// 	CellNum  string `db:"cell_num_"`
// }

type Contacts struct {
	UserId int64 `json:"userid_"`
}

type ListContacts struct {
	ContactList string `db:"fn_getcontactsjson"`
}
