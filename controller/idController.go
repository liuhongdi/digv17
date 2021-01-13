package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/liuhongdi/digv17/pkg/result"
	"github.com/liuhongdi/digv17/service"
)

type IdController struct{}

func NewIdController() IdController {
	return IdController{}
}
//得到一个id
func (a *IdController) GetOne(c *gin.Context) {
	resultRes := result.NewResult(c)
    idType := "logId"
	idOne,err := service.GetOneId(idType);
	if err != nil {
		resultRes.Error(404,"数据查询错误")
	} else {
		fmt.Println(idOne)
		resultRes.Success(&idOne);
	}
	return
}
