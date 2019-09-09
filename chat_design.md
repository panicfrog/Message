# Chat Design



## UserService

```go
func AddUser(userId string, name string, passHash string)
func RemoveUser(userId string)
func AddFriendRequest(fromUserId string, toUserId string)
func ApproveFriengRequest(fromUserId string, toUserId string)
func RejectFriendRequest(fromUserId string, toUserId string)
```

## User

```go
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
  Friends []User
  Status  UserStatus
}

func MessageUser(friendId string, message string)
func MessageGroup(groupId string, message string)
func SendFriendRequest(friendId string)
func ReciveFriendRequest(frientId string)
func ApproveFriendRequest(friendId string)
func RejectFriendRequest(friendId string)
```

## Message

```go
type MessageType int
const (
  TextMessage    MessageType = 0x01
  PictureMessage MessageType = 0x02
  VoiceMessage   MessageType = 0x03
  FileMessage    MessageType = 0x04
)
type Message struct {
  Id          string
  From        string
  To          string
  CreateTime  time.TimeStamp
  Type        MessageType
  Content     string
}
func EncodeMessage(msg string) (Message, error)
func DecodeMessage(msg Message) (string, error)
```

```go
type AddRequestStatus int 

const (
  Add_UNREAD    AddRequestStatus = 0x01
  Add_READ      AddRequestStatus = 0x02
  Add_ACCETPED  AddRequestStatus = 0x03
  Add_REJECTED  AddRequestStatus = 0x04
)

type AddRequest struct {
  From        string
  To          string
  Id          string
  CreateTime  time.TimeStamp
  Status      AddRequestStatus
}
```