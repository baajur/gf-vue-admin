package system

import (
	v1 "gf-server/app/api/v1"
	"gf-server/global"
	"gf-server/middleware"
)

// InitMenuRouter 注册menu路由
func InitMenuRouter() {
	// TODO 缺少CasbinHandler中间件
	MenuRouter := global.GFVA_SERVER.Group("menu").Middleware(middleware.JwtAuth)
	{
		MenuRouter.POST("getMenu", v1.GetMenu)          // 获取菜单树
		MenuRouter.POST("getMenuList", v1.GetMenuList)      // 分页获取基础menu列表
		MenuRouter.POST("addBaseMenu", v1.AddBaseMenu)      // 新增菜单
		MenuRouter.POST("getBaseMenuTree", v1.GetBaseMenuTree)  // 获取用户动态路由
		MenuRouter.POST("addMenuAuthority", v1.AddMenuAuthority) //	增加menu和角色关联关系
		MenuRouter.POST("getMenuAuthority", v1.GetMenuAuthority) // 获取指定角色menu
		MenuRouter.POST("deleteBaseMenu", v1.DeleteBaseMenu)   // 删除菜单
		MenuRouter.POST("updateBaseMenu", v1.UpdateBaseMenu)   // 更新菜单
		MenuRouter.POST("getBaseMenuById", v1.GetBaseMenuTree)  // 根据id获取菜单
	}
}