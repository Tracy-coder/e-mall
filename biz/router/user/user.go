// Code generated by hertz generator. DO NOT EDIT.

package user

import (
	user "github.com/Tracy-coder/e-mall/biz/handler/user"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_api := root.Group("/api", _apiMw()...)
		{
			_v1 := _api.Group("/v1", _v1Mw()...)
			_v1.GET("/captcha", append(_captchaMw(), user.Captcha)...)
			_v1.GET("/carousels", append(_listcarouselsMw(), user.ListCarousels)...)
			_v1.GET("/categories", append(_listcategoryMw(), user.ListCategory)...)
			_v1.GET("/product", append(_showproductMw(), user.ShowProduct)...)
			_product := _v1.Group("/product", _productMw()...)
			_product.GET("/img", append(_showproductimgMw(), user.ShowProductImg)...)
			_product.GET("/rankings", append(_showproductrankingsMw(), user.ShowProductRankings)...)
			_v1.GET("/products", append(_listproductMw(), user.ListProduct)...)
			_v1.POST("/register", append(_registerMw(), user.Register)...)
			_v1.GET("/verify_email", append(_verifyemailMw(), user.VerifyEmail)...)
			{
				_github := _v1.Group("/github", _githubMw()...)
				_github.GET("/login", append(_gtloginMw(), user.GTLogin)...)
				_login := _github.Group("/login", _loginMw()...)
				_login.GET("/callback", append(_gtlogincallbackMw(), user.GTLoginCallback)...)
			}
			{
				_user := _v1.Group("/user", _userMw()...)
				_user.GET("/info", append(_userinfoMw(), user.UserInfo)...)
				_user.POST("/upload_avatar", append(_uploadavatarMw(), user.UploadAvatar)...)
				_user.POST("/valid_email", append(_validemailMw(), user.ValidEmail)...)
			}
		}
		{
			_v2 := _api.Group("/v2", _v2Mw()...)
			_v2.POST("/carousels", append(_createcarouselMw(), user.CreateCarousel)...)
			_v2.POST("/categories", append(_createcategoryMw(), user.CreateCategory)...)
			_v2.DELETE("/product", append(_deleteproductMw(), user.DeleteProduct)...)
			_product0 := _v2.Group("/product", _product0Mw()...)
			_product0.POST("/img", append(_createproductimgMw(), user.CreateProductImg)...)
			_v2.PUT("/product", append(_updateproductMw(), user.UpdateProduct)...)
			_v2.POST("/products", append(_createproductMw(), user.CreateProduct)...)
		}
	}
}
