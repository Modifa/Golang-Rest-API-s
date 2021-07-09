package models

//
type AddtoGroups struct {
	Userid    int64  `json:"userid_"`
	Addedby   int64  `json:"addedby_"`
	Groupname string `json:"groupname"`
}

//
type RemoveFromGroup struct {
	Deleted_user int64 `json:"deleted_user"`
	Deletedby_   int64 `json:"deletedby_"`
}

//
type CreateGroup struct {
	Groupname string `json:"groupname"`
}

//
type Groups struct {
	UserId int64 `json:"userid_"`
}

type ListGroups struct {
	Userid int64 `db:"fn_getgroupsjson"`
}
