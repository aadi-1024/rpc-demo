package main

type CacheRPC struct{}

type UserModel struct {
	Id    int
	Email string
}

func (r *CacheRPC) AddUser(user *UserModel, resp *string) error {
	userStore.AddUser(user)
	*resp = "success"
	return nil
}

func (r *CacheRPC) DeleteUser(id int, resp *string) error {
	err := userStore.DeleteUser(id)
	if err != nil {
		*resp = err.Error()
		return err
	}
	*resp = "success"
	return nil
}

func (r *CacheRPC) GetAll(_ int, resp *map[int]string) error {
	*resp = userStore.GetAll()
	return nil
}
