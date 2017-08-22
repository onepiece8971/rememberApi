// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"rememberApi/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/home", &controllers.HomeController{}, "*:GetBooks")
	beego.Router("/remember/:uid", &controllers.RememberController{}, "get:GetMemoryUserBooks")
	beego.Router("/allRemember/:uid", &controllers.RememberController{}, "get:GetUserBooks")
	beego.Router("/posts/:id/:page", &controllers.PostsController{}, "*:GetPostsByUserBooksId")
	beego.Router("/post/:ubId/:postId", &controllers.PostsController{}, "*:GetPostById")
	beego.Router("/recite/:id", &controllers.ReciteController{}, "*:GetRecitesByUserBooksId")
	beego.Router("/remember", &controllers.ReciteController{}, "post:Remember")
	beego.Router("/forget", &controllers.ReciteController{}, "post:Forget")
	beego.Router("/addrecite/:ubId/:postId", &controllers.ReciteController{}, "*:AddRecite")
}
