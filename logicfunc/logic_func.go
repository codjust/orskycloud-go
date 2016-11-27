package logicfunc

import (
	"orskycloud-go/models"
)

func GetHomePage(username, password string) string {
	last_login_time := models.ReturnHomePage(username, password)
	return last_login_time
}
