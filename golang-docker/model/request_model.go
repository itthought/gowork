package model

type CreateSocketRoomRequest struct {
	GameId      int			`json:"gid,omitempty"`
	SsoId       string		`json:"ssoId,omitempty"`
	RoomId      string      `json:"roomId,omitempty"`
	RoomSize    int         `json:"rsize,omitempty"`
	Address     string       `json:"address,omitempty"`
	Port        int          `json:"port,omitempty"`
	Protocol    string       `json:"protocol,omitempty"`
	AudioEnable bool         `json:"isAudio,omitempty"`
	MinPlayer   int          `json:"minPlayer,omitempty"`
	StopTime    int           `json:"stopTime,omitempty"`
}

func NewCreateSocketRoomRequest() (rcvr *CreateSocketRoomRequest) {
	rcvr = &CreateSocketRoomRequest{}
	return
}
func (rcvr *CreateSocketRoomRequest) GetAddress() (string) {
	return rcvr.Address
}
func (rcvr *CreateSocketRoomRequest) GetAudioEnable() (bool) {
	return rcvr.AudioEnable
}
func (rcvr *CreateSocketRoomRequest) GetGameId() (int) {
	return rcvr.GameId
}
func (rcvr *CreateSocketRoomRequest) GetMinPlayer() (int) {
	return rcvr.MinPlayer
}
func (rcvr *CreateSocketRoomRequest) GetPort() (int) {
	return rcvr.Port
}
func (rcvr *CreateSocketRoomRequest) GetProtocol() (string) {
	return rcvr.Protocol
}
func (rcvr *CreateSocketRoomRequest) GetRoomId() (string) {
	return rcvr.RoomId
}
func (rcvr *CreateSocketRoomRequest) GetRoomSize() (int) {
	return rcvr.RoomSize
}
func (rcvr *CreateSocketRoomRequest) GetSsoId() (string) {
	return rcvr.SsoId
}
func (rcvr *CreateSocketRoomRequest) GetStopTime() (int) {
	return rcvr.StopTime
}
func (rcvr *CreateSocketRoomRequest) SetAddress(address string) {
	rcvr.Address = address
}
func (rcvr *CreateSocketRoomRequest) SetAudioEnable(audioEnable bool) {
	rcvr.AudioEnable = audioEnable
}
func (rcvr *CreateSocketRoomRequest) SetGameId(gameId int) {
	rcvr.GameId = gameId
}
func (rcvr *CreateSocketRoomRequest) SetMinPlayer(minPlayer int) {
	rcvr.MinPlayer = minPlayer
}
func (rcvr *CreateSocketRoomRequest) SetPort(port int) {
	rcvr.Port = port
}
func (rcvr *CreateSocketRoomRequest) SetProtocol(protocol string) {
	rcvr.Protocol = protocol
}
func (rcvr *CreateSocketRoomRequest) SetRoomId(roomId string) {
	rcvr.RoomId = roomId
}
func (rcvr *CreateSocketRoomRequest) SetRoomSize(roomSize int) {
	rcvr.RoomSize = roomSize
}
func (rcvr *CreateSocketRoomRequest) SetSsoId(ssoId string) {
	rcvr.SsoId = ssoId
}
func (rcvr *CreateSocketRoomRequest) SetStopTime(stopTime int) {
	rcvr.StopTime = stopTime
}

type JoinSocketRoomRequest struct {
	GameId      string  `json:"gameId,omitempty"`
	SsoId       string  `json:"ssoId,omitempty"`
	RoomId      string  `json:"roomId,omitempty"`
	AudioEnable bool    `json:"isAudio,omitempty"`
}

type SocketRoomStatusRequest struct {
	GameId      string  `json:"gameId,omitempty"`
	SsoId       string  `json:"ssoid,omitempty"`
	RoomId      string  `json:"roomId,omitempty"`
}

