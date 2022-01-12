package param

type Room struct {
	RoomID     string `json:"room_id"`
	RoomPasswd string `json:"room_passwd"`
	RoomOwener string `json:"room_owener"`
}
