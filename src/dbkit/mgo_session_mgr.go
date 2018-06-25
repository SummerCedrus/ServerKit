package dbkit
type SessionPool struct {
	Pool map[string]*MgoDBO
}

var SessionMgr SessionPool

func init() {
	SessionMgr = SessionPool{
		Pool:make(map[string]*MgoDBO, 0),
	}
}

func(sp *SessionPool) Add(sessionName string, dbo *MgoDBO){
	pool := sp.Pool
	if nil == pool{
		sp.Pool = make(map[string]*MgoDBO, 0)
	}

	pool[sessionName] = dbo
}

func (sp *SessionPool) Del(sessionName string) bool{
	pool := sp.Pool
	if nil == pool{
		return false
	}
	delete(pool, sessionName)

	return true
}
