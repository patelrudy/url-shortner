package server

import (
	"fmt"
	"urlshortner/model"
	"urlshortner/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func redirect(c *fiber.Ctx) error {
	urlshortnerUrl := c.Params("redirect")
	urlshortner, err := model.FindByUrlshortnerUrl(urlshortnerUrl)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "could not find urlshortner in DB " + err.Error(),
		})
	}

	urlshortner.Clicked += 1
	err = model.UpdateUrlshortner(urlshortner)
	if err != nil {
		fmt.Printf("error updating: %v\n", err)
	}


	return c.Redirect(urlshortner.Redirect, fiber.StatusTemporaryRedirect)
}

func getAllUrlshortners(c *fiber.Ctx) error {
	golies, err := model.GetAllUrlshortners()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "error getting all urlshortner links " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(golies)
}

func getUrlshortner(c *fiber.Ctx) error {
	id, err := strconv.ParseUint( c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "error could not parse id " + err.Error(),
		})
	}

	urlshortner, err := model.GetUrlshortner(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "error could not retreive urlshortner from db " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(urlshortner)
}

func createUrlshortner(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var urlshortner model.Urlshortner
	err := c.BodyParser(&urlshortner)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "error parsing JSON " + err.Error(),
		})
	}

	if urlshortner.Random {
		urlshortner.Urlshortner = utils.GetRandomUrl(8)
	}

	err = model.CreateUrlshortner(urlshortner)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "could not create urlshortner in db " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(urlshortner)

}

func updateUrlshortner(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var urlshortner model.Urlshortner

	err := c.BodyParser(&urlshortner)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "could not parse json " + err.Error(),
		})
	}

	err = model.UpdateUrlshortner(urlshortner)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "could not update urlshortner link in DB " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(urlshortner)
}

func deleteUrlshortner(c *fiber.Ctx) error {
	id, err := strconv.ParseUint( c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "could not parse id from url " + err.Error(),
		})
	}

	err = model.DeleteUrlshortner(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "could not delete from db " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map {
		"message": "urlshortner deleted.",
	})
}


func SetupAndListen() {

	router := fiber.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	router.Get("/r/:redirect", redirect)

	router.Get("/urlshortner", getAllUrlshortners)
	router.Get("/urlshortner/:id", getUrlshortner)
	router.Post("/urlshortner", createUrlshortner)
	router.Patch("/urlshortner", updateUrlshortner)
	router.Delete("/urlshortner/:id", deleteUrlshortner)

	router.Listen(":3000")
	
}