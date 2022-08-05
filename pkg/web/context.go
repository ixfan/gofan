package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/ixfan/gofan/pkg/database/orm"
	"github.com/ixfan/gofan/pkg/global"
	"github.com/ixfan/gofan/pkg/web/auth"
	"reflect"
	"strings"
)

type Context struct {
	*gin.Context
}

func NewContext(context *gin.Context) *Context {
	return &Context{context}
}

//Transaction 获取事务
func (context *Context) Transaction() *orm.Transaction {
	transaction, exists := context.Get(global.TransactionKey)
	if !exists {
		return nil
	}
	return transaction.(*orm.Transaction)
}

func (context *Context) Auth() *auth.User {
	user, exists := context.Get("AuthUser")
	if !exists {
		return nil
	}
	return user.(*auth.User)
}

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

func (context *Context) Validate(params interface{}) error {
	validate = validator.New()
	lang := zh.New()
	uni = ut.New(lang, lang)
	trans, _ := uni.GetTranslator("zh")
	_ = zhTranslations.RegisterDefaultTranslations(validate, trans)
	err := validate.Struct(params)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}
		elem := reflect.TypeOf(params).Elem()
		message := make([]string, 0)
		for _, err := range err.(validator.ValidationErrors) {
			fieldName := err.Field()
			field, isExist := elem.FieldByName(err.Field())
			if isExist {
				fieldName = field.Tag.Get("label")
			}
			message = append(message, fieldName+strings.ReplaceAll(err.Translate(trans), err.Field(), ""))
		}
		return fmt.Errorf(strings.Join(message, ";"))
	}
	return nil
}
