package model

type CreateSocketRoomResponse struct {
	RoomId            string `json:"roomId,omitempty"`
	SocketStatus      string `json:"ss,omitempty"`
	GameStatus        string `json:"gs,omitempty"`
	Address           string `json:"ip,omitempty"`
	Port              int `json:"port,omitempty"`
	Protocol          string `json:"protocol,omitempty"`
	Audio             string `json:"audio,omitempty"`
	RegistrationCount int `json:"regcount,omitempty"`
}

func NewCreateSocketRoomResponse() (rcvr *CreateSocketRoomResponse) {
	rcvr = &CreateSocketRoomResponse{}
	return
}
func (rcvr *CreateSocketRoomResponse) GetAddress() (string) {
	return rcvr.Address
}
func (rcvr *CreateSocketRoomResponse) GetAudio() (string) {
	return rcvr.Audio
}
func (rcvr *CreateSocketRoomResponse) GetGameStatus() (string) {
	return rcvr.GameStatus
}
func (rcvr *CreateSocketRoomResponse) GetPort() (int) {
	return rcvr.Port
}
func (rcvr *CreateSocketRoomResponse) GetProtocol() (string) {
	return rcvr.Protocol
}
func (rcvr *CreateSocketRoomResponse) GetRegistrationCount() (int) {
	return rcvr.RegistrationCount
}
func (rcvr *CreateSocketRoomResponse) GetRoomId() (string) {
	return rcvr.RoomId
}
func (rcvr *CreateSocketRoomResponse) GetSocketStatus() (string) {
	return rcvr.SocketStatus
}
func (rcvr *CreateSocketRoomResponse) SetAddress(address string) {
	rcvr.Address = address
}
func (rcvr *CreateSocketRoomResponse) SetAudio(audio string) {
	rcvr.Audio = audio
}
func (rcvr *CreateSocketRoomResponse) SetGameStatus(gameStatus string) {
	rcvr.GameStatus = gameStatus
}
func (rcvr *CreateSocketRoomResponse) SetPort(port int) {
	rcvr.Port = port
}
func (rcvr *CreateSocketRoomResponse) SetProtocol(protocol string) {
	rcvr.Protocol = protocol
}
func (rcvr *CreateSocketRoomResponse) SetRegistrationCount(registrationCount int) {
	rcvr.RegistrationCount = registrationCount
}
func (rcvr *CreateSocketRoomResponse) SetRoomId(roomId string) {
	rcvr.RoomId = roomId
}
func (rcvr *CreateSocketRoomResponse) SetSocketStatus(socketStatus string) {
	rcvr.SocketStatus = socketStatus
}
