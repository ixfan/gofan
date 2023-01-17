package web

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/arrutil"
	"github.com/ixfan/gofan/pkg/tools"
	"os"
	"path"
	"time"
)

type UploadController struct {
	Controller
}

func Upload() {
	upload := Default().Group("/api/v1/upload")
	{
		upload.POST("/file", func(context *gin.Context) {
			UploadController{}.UploadFile(NewContext(context))
		})
		upload.POST("/image", func(context *gin.Context) {
			UploadController{}.UploadImage(NewContext(context))
		})
	}
}

// UploadFile 上传文件
// @Summary 上传文件
// @Tags 公共模块
// @Accept json
// @Produce json
// @Router /api/v1/common/upload/file [post]
// @Success 200 {object} object{code=int,data={path=string}}
func (controller UploadController) UploadFile(ctx *Context) {
	file, err := ctx.Context.FormFile("file")
	if err != nil {
		controller.Response(ctx, nil, errors.New("请上传文件"))
		return
	}
	ext := path.Ext(file.Filename)
	newId, _ := tools.NewSnowflakeId()
	folder := "/" + StaticFolder + "/files/" + time.Now().Format("20060102")
	_ = os.MkdirAll("."+folder, os.ModePerm)
	dst := folder + "/" + fmt.Sprintf("%v", newId) + ext
	err = ctx.SaveUploadedFile(file, "."+dst)
	if err != nil {
		controller.Response(ctx, nil, errors.New("请确认上传目录是否有写入权限"))
		return
	}
	controller.Response(ctx, map[string]string{
		"path": dst,
	}, nil)
}

// UploadImage 上传图片
// @Summary 上传图片
// @Tags 公共模块
// @Accept json
// @Produce json
// @Router /api/v1/common/upload/image [post]
// @Success 200 {object} object{code=int,data={path=string}}
func (controller UploadController) UploadImage(ctx *Context) {
	file, err := ctx.Context.FormFile("image")
	if err != nil {
		controller.Response(ctx, nil, ThrowError(ArgError, "请上传文件"))
		return
	}
	ext := path.Ext(file.Filename)
	if !arrutil.InStrings(ext, []string{".jpg", ".png", ".gif", ".jpeg"}) {
		controller.Response(ctx, nil, ThrowError(ArgError, "请确认上传.jpg,.png,.gif,.jpeg后缀格式的图片"))
	}
	newId, _ := tools.NewSnowflakeId()
	folder := "/" + StaticFolder + "/images/" + time.Now().Format("20060102")
	dst := folder + "/" + fmt.Sprintf("%v", newId) + ext
	_ = os.MkdirAll("."+folder, os.ModePerm)
	err = ctx.SaveUploadedFile(file, "."+dst)
	if err != nil {
		controller.Response(ctx, nil, ThrowError(ArgError, "请确认上传目录是否有写入权限"))
		return
	}
	controller.Response(ctx, map[string]string{
		"path": dst,
	}, nil)
}
