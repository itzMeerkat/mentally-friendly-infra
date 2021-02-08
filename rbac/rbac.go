package rbac

import (
	"github.com/casbin/casbin/v2"
)

type AccessControl interface {
	GetAction() (string, string, string)
}

var enforcer *casbin.Enforcer

func InitRbacEnforcer(confPath string, csvPath string) {
	var err error
	enforcer, err = casbin.NewEnforcer(confPath, csvPath)
	if err != nil {
		panic(err)
	}
}

func CheckAccess(d AccessControl) bool {
	sub, obj, act := d.GetAction()
	ok, err := enforcer.Enforce(sub, obj, act)
	if err != nil {
		panic(err)
	}
	return ok
}