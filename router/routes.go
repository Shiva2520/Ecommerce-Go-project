package router

import (
	"ecommerce/constant"
	"ecommerce/controller"
	"net/http"
)

// health check service
var healthCheckRoutes = Routes{
	Route{"Health Check", http.MethodGet, constant.HealthCheckRoute, controller.HealthCheckRoute},
}
