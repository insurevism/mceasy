package controller

import (
	"mceasy/internal/applications/health/dto"
	"mceasy/internal/applications/health/service"
	"mceasy/internal/helper/response"
	"mceasy/internal/vars"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

const FlagDependency = "dependency"

type HealthController struct {
	service service.HealthService
}

func NewHealthController(service service.HealthService) *HealthController {
	return &HealthController{
		service: service,
	}
}

// Health status check for application
// flag query parameter with value `/health?flag=dependency` will use for health check READINESS application
// flag query parameter with NO value `/health` will use for health check LIVELINESS application
func (controller *HealthController) Health(c echo.Context) error {
	queryFlag := c.QueryParam("flag")
	if queryFlag != FlagDependency {
		var responseLivelinessDto = dto.HealthResponse{
			Status:     "UP",
			Message:    "",
			Timestamp:  time.Now(),
			Components: nil,
		}

		return response.Success(c, responseLivelinessDto)
	}

	msgController := "hello from controller layer "
	result, err := controller.service.Health(c.Request().Context(), msgController)

	var responseDto = dto.HealthResponse{
		Status:    "UP",
		Message:   result["final_msg"],
		Timestamp: time.Now(),
		Components: &dto.Components{
			Ctx: dto.Details{
				Status: result["ctx_status"],
				Name:   result["ctx_name"],
			},
			Db: dto.Details{
				Status: result["db_status"],
				Name:   result["db_name"],
			},
			Cache: dto.Details{
				Status: result["cache_status"],
				Name:   result["cache_name"],
			},
		},
	}

	if err != nil {
		errCode := http.StatusInternalServerError
		return response.Base(c, errCode, errCode, vars.ApplicationName()+" system not ready", responseDto, err)
	}

	return response.Success(c, responseDto)
}
