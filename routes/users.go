package routes

import (
	c "whatapp/controller"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	user := r.Group("User")
	{
		user.POST("/Register", c.Register)
		user.GET("/LoginUser", c.LoginUser)

	}

	contacts := r.Group("contacts")
	{
		contacts.GET("/ListContacts", c.ListContacts)
		contacts.GET("/ListGroups", c.ListGroups)
	}

	messages := r.Group("Messages")
	{
		messages.POST("/SendMessage", c.SendMessage)
		messages.POST("/DeleteMessage", c.DeleteMessage)
		messages.POST("/SendGroupMessage", c.SendGroupMessage)
		messages.POST("/DeletegroupMessage", c.DeletegroupMessage)
	}
	groups := r.Group("Groups")
	{
		groups.POST("/CreateGroup", c.CreateGroup)
		groups.POST("/Addtogroup", c.Addtogroup)
		groups.POST("/RemoveFromGroup", c.RemoveFromGroup)
	}
}
