package data

type UserStatus int
const (
	UserStatusOnline  = 0x01
	UserStatusOffLine = 0x02
	UserStatusBusy    = 0x03
)

type User struct {
	Account string
	Passwd  string
	Id      string
	Frients []User
	Status  UserStatus
}