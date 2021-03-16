package rbac

import (
	"github.com/casbin/casbin/v2"
)

var enforcer *casbin.Enforcer

func InitRbacEnforcer(confPath string, csvPath string) {
	var err error
	enforcer, err = casbin.NewEnforcer(confPath, csvPath)
	if err != nil {
		panic(err)
	}
}

func CheckAccess(sub,obj,act string) bool {
	ok, err := enforcer.Enforce(sub, obj, act)
	if err != nil {
		panic(err)
	}
	return ok
}