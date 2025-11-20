package session

import (
	"sync"
	"time"
)

type Session struct {
	UserLogin string
	ExpiresAt time.Time
}

type Store struct { // TODO - потом мб переделать с подключением в redis
	mu       sync.RWMutex
	sessions map[string]Session
}

func NewStore() *Store {
	return &Store{
		sessions: make(map[string]Session),
	}
}

func (s *Store) Set(sessionID string, userID string, duration time.Duration) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.sessions[sessionID] = Session{
		UserLogin: userID,
		ExpiresAt: time.Now().Add(duration),
	}
}

func (s *Store) Get(sessionID string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	session, exists := s.sessions[sessionID]
	if !exists {
		return "", false
	}

	if time.Now().After(session.ExpiresAt) {
		return "", false
	}

	return session.UserLogin, true
}

func (s *Store) Delete(sessionID string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.sessions, sessionID)
}

func (s *Store) CleanupExpired() {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now()
	for id, session := range s.sessions {
		if now.After(session.ExpiresAt) {
			delete(s.sessions, id)
		}
	}
}
