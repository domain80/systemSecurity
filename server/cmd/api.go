package main

import (
	"fmt"
	"log"
	"security418/internal/auth"
	"security418/internal/dbAdaptors/sqlte"
	"security418/internal/drug"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/mattn/go-sqlite3"
)



type server struct {
  auth auth.Service
  drug drug.Service
  api *fiber.App
}


func main() {
  fmt.Println("another banger")

  fmt.Println("another banger")
  db := sqlte.NewAdaptor()

  app := &server{
    auth: auth.NewAuthService(db),
    drug: drug.NewDrugService(db),
    api : fiber.New(),
  }

	app.api.Use(logger.New())

	app.api.Post("/login", app.login)
	app.api.Post("/logout", logout)
	app.api.Post("/change-role", app.changeRole)
	app.api.Post("/add-user", app.addUser)
	app.api.Get("/drugs", app.getAllDrugs)
	app.api.Post("/add-drug", app.addDrug)
	app.api.Put("/update-drug/:id", app.updateDrug)
	app.api.Put("/archive-drug/:id", app.archiveDrug)

  log.Println("Server starting on port 3000")
	log.Fatal(app.api.Listen(":3000"))
}

func (this server) login(c *fiber.Ctx) error {
  var user auth.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

  tokenString, err := this.auth.Login(user.ID, user.Password)
  if(err != nil) {
    return c.JSON(fiber.Map{"status": "error", "message": err.Error()})
  }


  // Set the token in the Authorization header of the response
	c.Set("Authorization", "Bearer "+tokenString)

	return c.JSON(fiber.Map{"message": "Login successful"})
}

func logout(c *fiber.Ctx) error {
  c.Set("Authorization", "") // Remove token from response header
	return c.JSON(fiber.Map{"message": "Logout successful"})
}

func (this *server) changeRole(c *fiber.Ctx) error {
	type ChangeRoleRequest struct {
		UserID int `json:"userID"`
		Role   string `json:"role"`
	}
	var request ChangeRoleRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

  err := this.auth.ChangeRole(request.UserID, request.Role)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Role updated successfully"})
}

func (this *server) addUser(c *fiber.Ctx) error {
	var user auth.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

  _user, err := this.auth.AddUser(&user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

  return c.JSON(fiber.Map{"message": "auth.User added successfully", "data": _user})
}

func (this *server) getAllDrugs(c *fiber.Ctx) error {

  drugs, err := this.drug.GetAll()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.JSON(drugs)
}

func (this *server) addDrug(c *fiber.Ctx) error {
	var drug drug.Drug
	if err := c.BodyParser(&drug); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

  err := this.drug.AddDrug(&drug)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Drug added successfully"})
}


func (this *server) updateDrug(c *fiber.Ctx) error {
	id, err :=  strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}
	var drug drug.Drug
	if err := c.BodyParser(&drug); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

  drug, err = this.drug.UpdateDrug(id, &drug)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

  return c.JSON(fiber.Map{"message": "Drug updated successfully", "data": drug})
}

func (this *server) archiveDrug(c *fiber.Ctx) error {
	id , err := strconv.Atoi( c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

  this.drug.Archive(id)

	return c.JSON(fiber.Map{"message": "Drug archived successfully"})
}
