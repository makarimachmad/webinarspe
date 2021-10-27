package user

type(
	User struct{
		ID	int	`json:"id"`
		Nama	string	`json:"nama"`
		Alamat	string	`json:"alamat"`
		NoHP	string	`json:"no_hp"`
		Sekolahan Sekolah
	}
	Users []User

	Sekolah struct{
		
	}
)
