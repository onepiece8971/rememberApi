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
	beego.Router("/remember/:uid", &controllers.RememberController{}, "*:GetMemoryUserBooksByUid")
	beego.Router("/posts", &controllers.PostsController{}, "*:GetPostsByUserBooksId")
	beego.Router("/recite/:id", &controllers.ReciteController{}, "*:GetRecitesByUserBooksId")
}
