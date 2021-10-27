package user

import(
	"strconv"
	"spewebinar/database"
)

func (u *Users) GetUser() error{
	db, err := database.ConnectDB()
	if err != nil{
		return err
	}
	db.Find(&u)
	return nil
}

func (u *User) PostUser() error{
	db, err := database.ConnectDB()
	if err != nil{
		return err
	}
	db.Create(&u)
	return nil
}

func (u *User) UpdateUser(id string) error{
	db, err := database.ConnectDB()
	if err != nil{
		return err
	}
	i,_ := strconv.Atoi(id)
	db.Model(&User{}).Where("id = ?", i).Updates(&u)
	return nil
}

func (u *User) DeleteUser(id string) error{
	db, err := database.ConnectDB()
	if err != nil{
		return err
	}
	i,_ := strconv.Atoi(id)
	db.Delete(&u, i)
	return nil
}
