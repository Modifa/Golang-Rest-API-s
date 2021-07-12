package controller

import (
	"fmt"
	"net/http"
	"os"

	// "strconv"
	"whatapp/models"
	s "whatapp/services"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	// "github.com/jackc/pgconn"
)

// Register New User
func Register(c *gin.Context) {
	db := s.DB{}
	//
	var rb models.Returnblock
	var u models.UserRegisted
	if err := c.ShouldBindBodyWith(&u, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	//use function

	_, err := db.SaveOnDB("public.fn_adduser", u)
	if err != nil {
		fmt.Println("QueryRow failed: ", err.Error())
		errormessage := fmt.Sprintf("%v\n", err)
		c.JSON(http.StatusBadRequest, errormessage)
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, rb.New(false, "", err))
	} else {
		c.JSON(http.StatusOK, rb.New(true, "", u))
		fmt.Println(models.Message{})
	}
}

//Logged in User
func LoginUser(c *gin.Context) {
	var rb models.Returnblock
	var l models.Login
	if err := c.ShouldBindJSON(&l); err != nil {
		c.JSON(http.StatusUnprocessableEntity, rb.New(false, "Invalid json provided", err))
	}
	db := *new(s.DB)
	rows, err := db.ReturnUser("public.fn_loginjson", l)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		errormessage := fmt.Sprintf("%v\n", err)
		c.JSON(http.StatusBadRequest, rb.New(false, errormessage, err))
	}
	fmt.Fprintf(os.Stdout, "QueryRow Success: %v\n", rows)

}

//Get all Contacts

func ListContacts(c *gin.Context) {
	var rb models.Returnblock

	db := s.DB{}
	var k models.Contacts
	if err := c.ShouldBindBodyWith(&k, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, rb.New(false, "Invalid json provided", err))
		return
	}

	contacts, err := db.ReturnContactList("public.fn_getcontactsjson", k)

	if err != nil {
		fmt.Println("QueryRow failed: ", err.Error())
		errormessage := fmt.Sprintf("%v\n", err)
		c.JSON(http.StatusBadRequest, errormessage)
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	} else {
		fmt.Fprintf(os.Stdout, "QueryRow Success: %v\n", contacts)

		c.JSON(http.StatusOK, rb.New(true, "", contacts))
	}
}

// Get All Groups

func ListGroups(c *gin.Context) {
	var rb models.Returnblock

	db := s.DB{}
	var k models.Groups
	if err := c.ShouldBindBodyWith(&k, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, rb.New(false, "Invalid json provided", err))
		return
	}
	groups, err := db.ReturnGroups("public.fn_getgroupsjson", k)
	if err != nil {
		fmt.Println("QueryRow failed: ", err.Error())
		errormessage := fmt.Sprintf("%v\n", err)
		c.JSON(http.StatusBadRequest, errormessage)
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	} else {
		fmt.Fprintf(os.Stdout, "QueryRow Success: %v\n", groups)

		c.JSON(http.StatusOK, rb.New(true, "", groups))
	}
}

//Send Message

func SendMessage(c *gin.Context) {
	//database
	db := s.DB{}
	//
	var rb models.Returnblock
	var m models.Message
	if err := c.ShouldBindBodyWith(&m, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, rb.New(false, "Invalid json provided", err))
		return
	}
	//use function
	_, err := db.SaveOnDB("public.fn_send_message", m)
	if err != nil {
		fmt.Println("QueryRow failed: ", err.Error())
		errormessage := fmt.Sprintf("%v\n", err)
		c.JSON(http.StatusBadRequest, errormessage)
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, rb.New(false, "", err))
	} else {
		c.JSON(http.StatusOK, rb.New(true, "", m))
		fmt.Fprintf(os.Stdout, "QueryRow Success: %v\n", m.Message)
	}
}

// Delete sent Message

func DeleteMessage(c *gin.Context) {
	//database
	db := s.DB{}
	//
	var rb models.Returnblock
	var d models.Deletedmessage
	if err := c.ShouldBindBodyWith(&d, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, rb.New(false, "Invalid json provided", err))
		return
	}
	//use function
	_, err := db.SaveOnDB("public.fn_deletedmessage", d)
	if err != nil {
		fmt.Println("QueryRow failed: ", err.Error())
		errormessage := fmt.Sprintf("%v\n", err)
		c.JSON(http.StatusBadRequest, errormessage)
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, rb.New(false, "", err))
	}
	c.JSON(http.StatusOK, rb.New(true, "", d))
	fmt.Fprintf(os.Stdout, "QueryRow Success: %v\n", d.Messageid)
}

//Create New Group

func CreateGroup(c *gin.Context) {
	//database
	db := s.DB{}
	//
	var rb models.Returnblock
	var g models.CreateGroup
	if err := c.ShouldBindBodyWith(&g, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, rb.New(false, "Invalid json provided", err))
		return
	}
	//use function
	_, err := db.SaveOnDB("public.fn_addnewgroup", g)
	if err != nil {
		fmt.Println("QueryRow failed: ", err.Error())
		errormessage := fmt.Sprintf("%v\n", err)
		c.JSON(http.StatusBadRequest, errormessage)
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, rb.New(false, "", err))
	}
	c.JSON(http.StatusOK, rb.New(true, "", g))
	fmt.Fprintf(os.Stdout, "QueryRow Success: %v\n", g.Groupname)

}

//Add contacts to group

func Addtogroup(c *gin.Context) {
	//database
	db := s.DB{}
	//
	var rb models.Returnblock
	var a models.AddtoGroups
	if err := c.ShouldBindBodyWith(&a, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, rb.New(false, "Invalid json provided", err))
		return
	}
	//use function
	_, err := db.SaveOnDB("public.fn_addmembers", a)
	if err != nil {
		fmt.Println("QueryRow failed: ", err.Error())
		errormessage := fmt.Sprintf("%v\n", err)
		c.JSON(http.StatusBadRequest, errormessage)
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, rb.New(false, "", err))
	}
	c.JSON(http.StatusOK, rb.New(true, "", a))
	fmt.Fprintf(os.Stdout, "QueryRow Success: %v\n", a.Userid)

}

// Remove contact from group

func RemoveFromGroup(c *gin.Context) {
	//database
	db := s.DB{}
	//
	var rb models.Returnblock
	var r models.RemoveFromGroup
	if err := c.ShouldBindBodyWith(&r, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, rb.New(false, "Invalid json provided", err))
		return
	}
	//use function
	_, err := db.SaveOnDB("public.fn_deletemembers", r)
	if err != nil {
		fmt.Println("QueryRow failed: ", err.Error())
		errormessage := fmt.Sprintf("%v\n", err)
		c.JSON(http.StatusBadRequest, errormessage)
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, rb.New(false, "", err))
	}
	c.JSON(http.StatusOK, rb.New(true, "", r))
	fmt.Fprintf(os.Stdout, "QueryRow Success: %v\n", r.Deleted_user)

}

// send group Message

func SendGroupMessage(c *gin.Context) {
	//database
	db := s.DB{}
	//
	var rb models.Returnblock
	var m models.GroupsMessage
	if err := c.ShouldBindBodyWith(&m, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, rb.New(false, "Invalid json provided", err))
		return
	}
	//use function
	_, err := db.SaveOnDB("public.fn_sendgroupmessage", m)
	if err != nil {
		fmt.Println("QueryRow failed: ", err.Error())
		errormessage := fmt.Sprintf("%v\n", err)
		c.JSON(http.StatusBadRequest, errormessage)
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, rb.New(false, "", err))
	}
	c.JSON(http.StatusOK, rb.New(true, "", m))
	fmt.Fprintf(os.Stdout, "QueryRow Success: %v\n", m.Groupmessage)

}

// Delete Group Message

func DeletegroupMessage(c *gin.Context) {
	//database
	db := s.DB{}
	//
	var rb models.Returnblock
	var d models.DeletedGroupmessage
	if err := c.ShouldBindBodyWith(&d, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, rb.New(false, "Invalid json provided", err))
		return
	}
	//use function
	_, err := db.SaveOnDB("public.fn_deleted_groupmessage", d)
	if err != nil {
		fmt.Println("QueryRow failed: ", err.Error())
		errormessage := fmt.Sprintf("%v\n", err)
		c.JSON(http.StatusBadRequest, errormessage)
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, rb.New(false, "", err))
	}
	c.JSON(http.StatusOK, rb.New(true, "", d))
	fmt.Fprintf(os.Stdout, "QueryRow Success: %v\n", d.DeleteMessageid)

}

//

func Getmessages(c *gin.Context) {
	db := s.DB{}
	var rb models.Returnblock
	var m models.GetMessage
	if err := c.ShouldBindBodyWith(&m, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, rb.New(false, "Invalid json provided", err))
		return
	}

	messages, err := db.ReturnMessage("fn_getMessagesjson", m)
	if err != nil {
		fmt.Println("QueryRow failed: ", err.Error())
		errormessage := fmt.Sprintf("%v\n", err)
		c.JSON(http.StatusBadRequest, errormessage)
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, rb.New(false, "", err))
	}
	c.JSON(http.StatusOK, rb.New(true, "", m))
	fmt.Fprintf(os.Stdout, "QueryRow Success: %v\n", messages)

}

//
func GetGroupmessages(c *gin.Context) {
	db := s.DB{}
	var rb models.Returnblock
	var gm models.GetGroupMessage
	if err := c.ShouldBindBodyWith(&gm, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, rb.New(false, "Invalid json provided", err))
		return
	}

	groupMessages, err := db.ReturnGroupMessage("", gm)
	if err != nil {
		fmt.Println("QueryRow failed: ", err.Error())
		errormessage := fmt.Sprintf("%v\n", err)
		c.JSON(http.StatusBadRequest, errormessage)
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, rb.New(false, "", err))
	}
	c.JSON(http.StatusOK, rb.New(true, "", gm))
	fmt.Fprintf(os.Stdout, "QueryRow Success: %v\n", groupMessages)

}
