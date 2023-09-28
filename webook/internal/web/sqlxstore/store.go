package sqlxstore

import (
	GinSession "github.com/gin-contrib/sessions"
	"github.com/gorilla/sessions"
	"net/http"
)

// @program:   go-basic-exercises
// @file:      store.go
// @author:    bowen
// @time:      2023-09-27 11:33
// @description:

type Store struct {
}

func (s *Store) Get(r *http.Request, name string) (*sessions.Session, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Store) New(r *http.Request, name string) (*sessions.Session, error) {
	//TODO implement me
	panic("implement me")
}

func (st *Store) Save(r *http.Request, w http.ResponseWriter, s *sessions.Session) error {
	//TODO implement me
	panic("implement me")
}

func (s Store) Options(options GinSession.Options) {
	//TODO implement me
	panic("implement me")
}
