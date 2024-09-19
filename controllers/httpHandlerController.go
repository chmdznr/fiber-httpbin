package controllers

import (
	"fiber-httpbin/models"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"time"
)

// GetHandler godoc
// @Summary Returns the GET data
// @Description Echoes the GET request data
// @Tags GET
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{}
// @Router /get [get]
func GetHandler(c *fiber.Ctx) error {
	response := map[string]interface{}{
		"args":    c.Queries(),
		"headers": c.GetReqHeaders(),
		"origin":  c.IP(),
		"url":     c.OriginalURL(),
	}
	return c.JSON(response)
}

// PostHandler godoc
// @Summary Returns the POST data
// @Description Echoes the POST request data
// @Tags POST
// @Accept  json
// @Produce  json
// @Param data body map[string]interface{} true "POST Data"
// @Success 200 {object} map[string]interface{}
// @Router /post [post]
func PostHandler(c *fiber.Ctx) error {
	var body map[string]interface{}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	response := map[string]interface{}{
		"args":    c.Queries(),
		"data":    body,
		"headers": c.GetReqHeaders(),
		"origin":  c.IP(),
		"url":     c.OriginalURL(),
	}
	return c.JSON(response)
}

// HeadersHandler godoc
// @Summary Returns the request headers
// @Description Echoes the request headers
// @Tags Headers
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{}
// @Router /headers [get]
func HeadersHandler(c *fiber.Ctx) error {
	headers := c.GetReqHeaders()
	response := map[string]interface{}{
		"headers": headers,
	}
	return c.JSON(response)
}

// IPHandler godoc
// @Summary Returns the origin IP
// @Description Returns the IP address of the requester
// @Tags IP
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string
// @Router /ip [get]
func IPHandler(c *fiber.Ctx) error {
	ip := c.IP()
	return c.JSON(fiber.Map{
		"origin": ip,
	})
}

// DelayHandler godoc
// @Summary Delays the response
// @Description Delays the response by a specified number of seconds
// @Tags Delay
// @Accept  json
// @Produce  json
// @Param seconds query int true "Delay in seconds"
// @Success 200 {object} map[string]string
// @Router /delay [get]
func DelayHandler(c *fiber.Ctx) error {
	seconds := c.QueryInt("seconds")

	time.Sleep(time.Duration(seconds) * time.Second)
	return c.JSON(fiber.Map{
		"message": "Delayed response",
		"delay":   seconds,
	})
}

// StatusHandler godoc
// @Summary Returns specified HTTP status
// @Description Returns a response with the given HTTP status code
// @Tags Status
// @Accept  json
// @Produce  json
// @Param code query int true "HTTP Status Code"
// @Success 200 {object} map[string]string
// @Router /status [get]
func StatusHandler(c *fiber.Ctx) error {
	code := c.QueryInt("code")

	c.Status(code)
	return c.JSON(fiber.Map{
		"status": code,
	})
}

// AnythingHandler godoc
// @Summary Echoes the request data
// @Description Returns all details about the request
// @Tags Anything
// @Accept  json
// @Produce  json
// @Param method path string true "HTTP Method"
// @Param path path string true "Request Path"
// @Success 200 {object} models.AnythingResponse
// @Router /anything [get]
// @Router /anything [post]
// @Router /anything [put]
// @Router /anything [delete]
// @Router /anything [patch]
// @Router /anything [options]
// @Router /anything [head]
func AnythingHandler(c *fiber.Ctx) error {
	// Initialize response structure
	response := models.AnythingResponse{
		Method:  c.Method(),
		URL:     c.OriginalURL(),
		Headers: c.GetReqHeaders(),
		Args:    c.Queries(),
	}

	// Handle different content types
	contentType := c.Get("Content-Type")
	if contentType == "application/json" {
		var jsonData interface{}
		if err := c.BodyParser(&jsonData); err == nil {
			response.JSON = jsonData
		}
	} else if contentType == "application/x-www-form-urlencoded" || contentType == "multipart/form-data" {
		formData, err := c.Request().MultipartForm()
		if err == nil && formData != nil {
			response.Form = formData.Value
			// Handle files if any
			if files := formData.File; len(files) > 0 {
				response.Files = make(map[string][]models.FileInfo)
				for key, fileHeaders := range files {
					for _, fileHeader := range fileHeaders {
						fileInfo := models.FileInfo{
							Filename: fileHeader.Filename,
							Size:     fileHeader.Size,
							Header:   fileHeader.Header,
						}
						response.Files[key] = append(response.Files[key], fileInfo)
					}
				}
			}
		}
	} else {
		// For other content types, return the raw body
		response.Data = string(c.Body())
	}

	return c.JSON(response)
}
