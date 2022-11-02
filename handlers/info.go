package handlers

import "github.com/gofiber/fiber/v2"

func Info(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"baseUrl":     c.BaseURL(),
		"host":        c.Hostname(),
		"ip":          c.IP(),
		"ips":         c.IPs(),
		"method":      c.Method(),
		"originalUrl": c.OriginalURL(),
		"path":        c.Path(),
		"protocol":    c.Protocol(),
		"query":       c.Query("name"),
		"route":       c.Route().Path,
		"subdomains":  c.Subdomains(),
		"xhr":         c.XHR(),
		"AllParams":   c.AllParams(),
	})
}
