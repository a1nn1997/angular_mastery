package models
import(
	"time"
)
type User struct {
	Id       		uint   		`json:"id"`
	First_name		string				`json:"first_name" validate:"required,min=2,max=100"`
	Last_name		string				`json:"last_name" validate:"required,min=2,max=100"`
	Email			string				`json:"email" gorm:"unique" validate:"email,required"`
	Password		[]byte				`json:"-" validate:"required,min=8"`		   
	Phone			string				`json:"phone" gorm:"unique" validate:"required,min=10,max=10" `	
	Created_at		time.Time			`json:"created_at"`
	Updated_at		time.Time			`json:"updated_at"`
	Is_active		bool				`json:"active"`
}