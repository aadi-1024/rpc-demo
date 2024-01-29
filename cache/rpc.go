package main

type CacheRPC struct{}

type UserModel struct {
	Id    string
	Email string
}

func (r *CacheRPC) AddUser(user *UserModel, resp *string) error {
	userStore.AddUser(user)
	*resp = "success"
	return nil
}

func (r *CacheRPC) DeleteUser(id string, resp *string) error {
	err := userStore.DeleteUser(id)
	if err != nil {
		*resp = err.Error()
		return err
	}
	*resp = "success"
	return nil
}

func (r *CacheRPC) GetAll(_ int, resp *[]string) error {
	*resp = userStore.GetAll()
	return nil
}
