package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	bc "github.com/togettoyou/blockchain-real-estate/application/blockchain"
	"github.com/togettoyou/blockchain-real-estate/application/pkg/app"
	"net/http"
)

type FinanceRequestBody struct {
	Supplier		string  `json:"Supplier"` //申请人
	Amount 			string `json:"Amount"` //申请金额
	Period        	string  `json:"Period"`       //申请日期
	Desc  			string  `json:"Desc"` //申请描述
	OrderID			string  `json:"OrderID"` //对应订单号
	Status       	string  `json:"Status"`       //融资状态
	Feedback   		string  `json:"Feedback"`   //融资反馈
}

type UpdateStatusRequestBody struct {
   Period         string  `json:"Period"`       //申请日期
   OrderID          string  `json:"OrderID"` //对应订单号
   Status         string  `json:"Status"`       //融资状态
   Feedback          string  `json:"Feedback"`   //融资反馈
}

type FinanceBySupplierQueryRequestBody struct {
	Supplier		string  `json:"Supplier"` //申请人
}

type FinanceByStatusQueryRequestBody struct {
	Status       	string  `json:"Status"`       //融资状态
}

type FinanceByOrderIDQueryRequestBody struct {
	OrderID			string  `json:"OrderID"` //对应订单号
}

// @Summary 存储到数据库
// @Param Finance body FinanceRequestBody true "Finance"
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/StoreFinance [post]
func StoreFinance(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(FinanceRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.Supplier))
	bodyBytes = append(bodyBytes, []byte(body.Amount))
	bodyBytes = append(bodyBytes, []byte(body.Period))
	bodyBytes = append(bodyBytes, []byte(body.Desc))
	bodyBytes = append(bodyBytes, []byte(body.OrderID))
	bodyBytes = append(bodyBytes, []byte(body.Status))
	bodyBytes = append(bodyBytes, []byte(body.Feedback))

	//调用智能合约
	_, err := bc.ChannelExecute("StoreFinance", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", fmt.Sprintf("存储成功"))
}

// @Summary 更新融资状态
// @Param Finance body UpdateStatusRequestBody true "Finance"
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/UpdateFinanceStatus [post]
func UpdateFinanceStatus(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(UpdateStatusRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.OrderID))
	bodyBytes = append(bodyBytes, []byte(body.Period))
	bodyBytes = append(bodyBytes, []byte(body.Status))
	bodyBytes = append(bodyBytes, []byte(body.Feedback))

	//调用智能合约
	_, err := bc.ChannelExecute("UpdateFinanceStatus", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", fmt.Sprintf("更新融资状态成功"))
}

// @Summary 查询所有
// @Param Finance body ProductByNameQueryRequestBody true "Finance"
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/QueryFinance [post]
func QueryFinance(c *gin.Context) {
	appG := app.Gin{C: c}

	var bodyBytes [][]byte

	//调用智能合约
	resp, err := bc.ChannelQuery("QueryFinance", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	// 反序列化json
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

// @Summary 查询所有
// @Param Finance body FinanceBySupplierQueryRequestBody true "Finance"
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/QueryFinanceBySupplier [post]
func QueryFinanceBySupplier(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(FinanceBySupplierQueryRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte

	bodyBytes = append(bodyBytes, []byte(body.Supplier))

	//调用智能合约
	resp, err := bc.ChannelQuery("QueryFinanceBySupplier", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	// 反序列化json
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

// @Summary 按融资状态查询
// @Param Finance body FinanceByStatusQueryRequestBody true "Finance"
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/QueryFinanceByStatus [post]
func QueryFinanceByStatus(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(FinanceByStatusQueryRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	if body.Status != "" {
		bodyBytes = append(bodyBytes, []byte(body.Status))
	}
	//调用智能合约
	resp, err := bc.ChannelQuery("QueryFinanceByStatus", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	// 反序列化json
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

// @Summary 按订单号查询
// @Param Finance body FinanceByOrderIDQueryRequestBody true "Finance"
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/QueryFinanceByOrderID [post]
func QueryFinanceByOrderID(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(FinanceByOrderIDQueryRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	if body.OrderID != "" {
		bodyBytes = append(bodyBytes, []byte(body.OrderID))
	}
	//调用智能合约
	resp, err := bc.ChannelQuery("QueryFinanceByOrderID", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	// 反序列化json
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

