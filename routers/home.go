package routers

import (
	"gitlab.com/kanban/kanban/modules/setting"
	"github.com/Unknwon/macaron"
)

func Home(ctx *macaron.Context) {
	ctx.Data["Version"] = setting.App_Version
	ctx.Data["GitlabHost"] = setting.Cfg.Section("gitlab").Key("ROOT_URL").String()
	ctx.HTML(200, "templates/index")
}