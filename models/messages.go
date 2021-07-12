package models

// Message Group
type Message struct {
	Message    string `db:"message"`
	Senderid   int64  `json:"senderid"`
	Recieverid int64  `json:"recieverid"`
}
type Deletedmessage struct {
	Messageid int64 `json:"messageid"`
}

//
type DeletedGroupmessage struct {
	DeleteMessageid int64 `db:"DeleteMessageid"`
	Deleter         int64 `json:"Deleter"`
}

//
type GroupsMessage struct {
	Group_id     int64  `json:"groupid"`
	Sender_id    int64  `json:"senderid"`
	Groupmessage string `db:"groupmessage"`
}

type ReturnMessages struct {
	Messages string `db:"fn_getMessagesjson"`
}
type GetMessage struct {
	UserId int64 `json:"userid_"`
}
