package main

import "errors"

type Store struct {
	data map[int]string
}

func (s *Store) AddUser(model *UserModel) {
	s.data[model.Id] = model.Email
}

func (s *Store) GetById(id int) (string, error) {
	email, ok := s.data[id]
	if !ok {
		return "", errors.New("does not exist")
	}
	return email, nil
}

func (s *Store) DeleteUser(id int) error {
	_, ok := s.data[id]
	if !ok {
		return errors.New("does not exist")
	}
	delete(s.data, id)
	return nil
}

func (s *Store) GetAll() map[int]string {
	return s.data
}
