package internel

import "errors"

var (
	EncodeMessageError = errors.New("encode message error")
	DecodeMessageError = errors.New("decode message error")
)

// DB
var (
	DBErrorExited = errors.New("data is exits")
)

var (
	RedisTokenNotExited = errors.New("token not exits")
	RedisTokenExpire    = errors.New("token expire")
)

var (
	RoomMemeberNotExited          = errors.New("room member not exited")
	RoomAdministratorNotExited    = errors.New("room administrator not exited")
	RoomAdministratorUnremoveable = errors.New("can remove administrator")
	RoomIsAlreadyAdministrator    = errors.New("already administrator")
	RoomCantCancelOwner           = errors.New("cant cancel owner")
)

var (
	UserCannotAddFriendWithSelf   = errors.New("can't add friend with self")
	UserYouAreAleadyFriend        = errors.New("your are aleady friends")
)