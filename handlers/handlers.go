package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/TaranovDmitry/DomainMicroservice/entity"
)

type Service interface {
	Ports() (entity.Ports, error)
	UpdatePorts(ports entity.Ports) error
}

type Handler struct {
	service Service
}

func NewHandler(services Service) *Handler {
	return &Handler{
		service: services,
	}
}

type err struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, err{message})
}

func (h *Handler) ports(c *gin.Context) {
	ports, err := h.service.Ports()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, ports)
}

func (h *Handler) uploadPorts(c *gin.Context) {
	formFile, err := c.FormFile("file")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	openedFile, err := formFile.Open()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	file, err := ioutil.ReadAll(openedFile)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var ports entity.Ports
	err = json.Unmarshal(file, &ports)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.UpdatePorts(ports)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.AbortWithStatus(http.StatusCreated)
}
