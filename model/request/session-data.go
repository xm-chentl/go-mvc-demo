package request

type SessionData struct {
	MobileDevice

	UID string
}

func (s *SessionData) SetSession(uid string) {
	s.UID = uid
}
