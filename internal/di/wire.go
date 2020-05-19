// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"kratos_frame/internal/dao"
	"kratos_frame/internal/server/http"
	"kratos_frame/internal/service"

	"github.com/google/wire"
)

var daoProvider = wire.NewSet(dao.New, dao.NewDB, dao.NewRedis, dao.NewMC, dao.NewCasBin)
var serviceProvider = wire.NewSet(service.New)

func InitApp() (*App, func(), error) {
	panic(wire.Build(daoProvider, serviceProvider, http.New, NewApp))
}
