package swagger

import (
	"mceasy/cmd/docs"

	"github.com/spf13/viper"
)

func InitSwagger() {
	docs.SwaggerInfo.Title = "Mceasy Service"
	docs.SwaggerInfo.Description = "Ex nihilo nihil fit"
	docs.SwaggerInfo.Version = viper.GetString("application.version")
	docs.SwaggerInfo.Host = viper.GetString("swagger.host")
	docs.SwaggerInfo.BasePath = "/" + viper.GetString("application.name")
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}
