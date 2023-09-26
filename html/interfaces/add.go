package interfaces

import (
	"add/models"

)

type Add interface{
	Add(m *models.Add)(string)
}