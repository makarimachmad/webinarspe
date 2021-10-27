package user

import(
	"github.com/gofiber/fiber/v2"
)

func UserGet(c *fiber.Ctx) error{
	user := new(Users)
	err := user.GetUser()
	if err != nil{
		return c.Status(404).JSON(fiber.Map{"status":"not found", "message":"data tidak ada", "data":nil})
	}
	return c.Status(200).JSON(fiber.Map{"status":"berhasil", "message":"data ditemukan", "data":user})
}

func UserPost(c *fiber.Ctx) error{
	user := new(User)
	err := c.BodyParser(user)
	if err != nil{
		return c.Status(500).JSON(fiber.Map{"status":"internal server error", "message":"tolong cek kembali isiannya", "data":nil})
	}
	err = user.PostUser()
	if err != nil{
		return c.Status(500).JSON(fiber.Map{"status":"internal server erro", "message":"gagal dimasukkan ke database", "data":nil})
	}
	return c.Status(200).JSON(fiber.Map{"status":"berhasil", "message":"berhasil dimasukkan ke database", "data":user})
}

func UserUpdate(c *fiber.Ctx) error{
	user := new(User)
	err := c.BodyParser(user)
	if err != nil{
		return c.Status(500).JSON(fiber.Map{"status":"internal server error", "message":"tolong cek kembali isiannya", "data":nil})
	}
	id := c.Params("idx")
	err = user.UpdateUser(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status":"internal server erro", "message":"gagal dimasukkan ke database", "data":nil})
	}
	return c.Status(200).JSON(fiber.Map{"status":"berhasil", "message":"berhasil dimasukkan ke database", "data":user})
}

func UserDelete(c *fiber.Ctx) error{
	user := new(User)
	id := c.Params("idx")
	err := user.DeleteUser(id)
	if err != nil{
		return c.Status(500).JSON(fiber.Map{"status":"internal server error", "message":"tolong cek kembali isiannya", "data":nil})
	}
	return c.Status(200).JSON(fiber.Map{"status":"berhasil", "message":"berhasil dimasukkan ke database"})
}
