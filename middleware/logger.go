package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func Logger(c *fiber.Ctx) error {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.DebugLevel)

	nameFile := fmt.Sprintf("logger/%d-%d-%d-%d.log", time.Now().Year(), time.Now().Month(), time.Now().Day())
	file, _ := os.OpenFile(nameFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(file)

	start := time.Now()
	c.Next()
	end := time.Now()
	logger.WithFields(logrus.Fields{
		"status":     c.Response().StatusCode(),
		"method":     c.Method(),
		"path":       c.Path(),
		"ip":         c.IP(),
		"latency":    end.Sub(start),
		"user-agent": c.Get("User-Agent"),
		"request":    c.Body(),
	}).Info()
	return nil

}
