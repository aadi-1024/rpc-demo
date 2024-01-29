package main

import "errors"

type Store struct {
	data map[string]string
}

func (s *Store) AddUser(model *UserModel) {
	s.data[model.Id] = model.Email
}

func (s *Store) GetById(id string) (string, error) {
	email, ok := s.data[id]
	if !ok {
		return "", errors.New("does not exist")
	}
	return email, nil
}

func (s *Store) DeleteUser(id string) error {
	_, ok := s.data[id]
	if !ok {
		return errors.New("does not exist")
	}
	delete(s.data, id)
	return nil
}

func (s *Store) GetAll() []string {
	x := make([]string, len(s.data))
	for _, i := range s.data {
		x = append(x, i)
	}
	return x
}
