package main

import (
	"fmt"
	web2 "github.com/beego/beego/v2/server/web"
	_ "github.com/dgrijalva/jwt-go"
	"github.com/ecnepsnai/web"
	_ "github.com/gravitational/teleport/api/types"
	_ "github.com/mattermost/mattermost-server/v6/model"
	"github.com/uadmin/uadmin"
)

func main() {
	uadmin.StartServer()
	server := web.New("127.0.0.1:8080")
	fmt.Println(server)

	web2.Run()
}
