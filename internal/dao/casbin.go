package dao

import (
	"os"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	_ "github.com/go-sql-driver/mysql"

	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/database/sql"
)

func NewCasBin() (e *casbin.Enforcer, err error) {
	var cfg struct {
		Client *sql.Config
	}
	if err = paladin.Get("db.toml").UnmarshalTOML(&cfg); err != nil {
		return
	}

	adapter, err := gormadapter.NewAdapter("mysql", cfg.Client.DSN)
	if err != nil {
		return
	}

	if e, err = casbin.NewEnforcer(os.Getenv("GOPATH")+"/src/kratos_frame/configs/rbac_model.conf", adapter); err != nil {
		return
	}

	// Load the policy from DB.
	err = e.LoadPolicy()
	return
}

// 添加权限
// args 角色， 访问路径， 访问动作
func (d *Dao) AddPolicy(name, path, method string) (bool, error) {
	return d.enforcer.AddPolicy(name, path, method)
}

// 删除某个角色的权限
func (d *Dao) DelPolicyByName(name string) (bool, error) {
	return d.enforcer.DeleteRole(name)
}

// 权限验证
func (d *Dao) EnforcePolicy(name, path, method string) (bool, error) {
	return d.enforcer.Enforce(name, path, method)
}

// 获取一个角色的所有授权
func (d *Dao) GetPolicyByName(name string) [][]string {
	return d.enforcer.GetFilteredPolicy(0, name)
}
