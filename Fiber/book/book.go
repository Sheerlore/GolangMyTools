package book

import (
	"github.com/Sheerlore/GolngMytools/Fiber/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)
type Book struct {
	gorm.Model
	Title	string	`json:"name"`
	Author	string	`json:"author"`
	Rating	int		`json:"rating"`
}

func GetBooks(c *fiber.Ctx) {
	//c.Send("All Books")
	db := database.DBconn
	var books []Book
	db.Find(&books)
	c.JSON(books)
	
}

func GetBook(c *fiber.Ctx) {
	//c.Send("Single Book")
	id := c.Params("id")
	db := database.DBconn
	var book Book
	db.Find(&book, id)
	c.JSON(book)

}

func NewBook(c *fiber.Ctx) {
	//c.Send("New Book")
	db := database.DBconn
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&book)
	c.JSON(book)
}

func DeleteBook(c *fiber.Ctx){
	//c.Send("Delete Book")
	id := c.Params("id")
	db := database.DBconn

	var book Book
	db.First(&book, id)
	if book.Title == "" {
		c.Status(500).Send("No Book Founc with ID")
		return
	}
	db.Delete(&book)
	c.Send("Book Successfully deleted")
}




