package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	bc "github.com/togettoyou/blockchain-real-estate/application/blockchain"
	"github.com/togettoyou/blockchain-real-estate/application/pkg/app"
	"net/http"
	"strconv"
)

type OrderRequestBody struct {
	ID              string	`json:"buyer"` //订单ID
	Buyer 			string  `json:"buyer"` //发起企业
	Seller          string  `json:"seller"`       //供货商
	BuyerID			string  `json:"seller"`       //企业ID
	SellerID		string  `json:"seller"`       //供应商ID
	Address			string  `json:"seller"`       //收货地址
	GoodsID       	string `json:"price"`        //货品ID
	GoodsName		string `json:"price"`        //货品名称
	GoodsNum   		int     `json:"salePeriod"`   //货品数量
	OrderAmount 	float64  `json:"buyer"` //订单总金额
	CreateTime      string `json:"price"`        //创建时间
	Deadline  	    int     `json:"salePeriod"`   //截止日期
	Remark 			int     `json:"salePeriod"`   //备注
	TransferStatus	string `json:"price"`        //物流状态
	TransferID		string `json:"price"`        //物流ID
	Status       	int  `json:"seller"`       //订单状态
}

type OrderByIDQueryRequestBody struct {
	ID 	string  `json:"buyer"` //订单ID
}

type OrderByBuyerQueryRequestBody struct {
	Buyer 			string  `json:"buyer"` //发起企业
}

type OrderBySellerQueryRequestBody struct {
	Seller   string  `json:"seller"`       //供货商
}

type UpdateOrderStatusRequestBody struct {
	ID 		  string    `json:"buyer"` //订单ID
	Status    int  		`json:"seller"`       //订单状态
}

type UpdateTransferStatusRequestBody struct {
	ID 		  		string    `json:"buyer"` //订单ID
	TransferStatus	string `json:"price"`        //物流状态
	TransferID		string `json:"price"`        //物流ID
}

// @Summary 存储到数据库
// @Param Order body OrderRequestBody true "Order"
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/StoreOrder [post]
func StoreOrder(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(OrderRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	if body.ID == "" {
		appG.Response(http.StatusBadRequest, "失败", "订单ID不能为空")
		return
	}

	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.ID))
	bodyBytes = append(bodyBytes, []byte(body.Buyer))
	bodyBytes = append(bodyBytes, []byte(body.Seller))
	bodyBytes = append(bodyBytes, []byte(body.BuyerID))
	bodyBytes = append(bodyBytes, []byte(body.SellerID))
	bodyBytes = append(bodyBytes, []byte(body.Address))
	bodyBytes = append(bodyBytes, []byte(body.GoodsID))
	bodyBytes = append(bodyBytes, []byte(body.GoodsName))
	bodyBytes = append(bodyBytes, []byte(strconv.Itoa(body.GoodsNum)))
	bodyBytes = append(bodyBytes, []byte(strconv.FormatFloat(body.OrderAmount, 'E', -1, 64)))
	bodyBytes = append(bodyBytes, []byte(body.CreateTime))
	bodyBytes = append(bodyBytes, []byte(strconv.Itoa(body.Deadline)))
	bodyBytes = append(bodyBytes, []byte(strconv.Itoa(body.Remark)))
	bodyBytes = append(bodyBytes, []byte(body.TransferStatus))
	bodyBytes = append(bodyBytes, []byte(body.TransferID))
	bodyBytes = append(bodyBytes, []byte(strconv.Itoa(body.Status)))
	//调用智能合约
	_, err := bc.ChannelExecute("StoreOrder", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", fmt.Sprintf("保存成功"))
}

// @Summary 按订单ID查询，具体查询某一个订单
// @Param Order body OrderByIDQueryRequestBody true "Order"
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/QueryOrderByID [post]
func QueryOrderByID(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(OrderByIDQueryRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	if body.ID != "" {
		bodyBytes = append(bodyBytes, []byte(body.ID))
	}
	//调用智能合约
	resp, err := bc.ChannelQuery("QueryOrderByID", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	// 反序列化json
	var data map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "反序列化失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

// @Summary 测试：按订单ID查询，具体查询某一个订单
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/QueryOrderByIDTest [post]
func QueryOrderByIDTest(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(OrderByIDQueryRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	if body.ID != "" {
		bodyBytes = append(bodyBytes, []byte(body.ID))
	}

	//var resp []byte
	resp := []byte("{\"ID\":\"123\","+
		"\"Buyer\":\"北京奔驰汽车有限公司顺义工厂\","+
		"\"Seller\":\"苏州快钢供应链信息技术服务有限公司\","+
		"\"BuyerID\":\"134155651531\","+
		"\"sellerID\":\"2347641535\","+
		"\"address\":\"北京市顺义区赵全营镇东盈路19号\","+
		"\"GoodsID\":\"lkdjklfjdkjfkl\","+
		"\"GoodsName\":\" 碳结圆钢 45# Φ85 鲁丽\","+
		"\"GoodsNum\":\"1000kg\","+
		"\"OrderAmount\":\"￥1000000\","+
		"\"CreateTime\":\"2021/6/11\","+
		"\"Deadline\":\"2021/6/14\","+
		"\"Remark\":\"无\","+
		"\"TransferStatus\":\"未发货\","+
		"\"TransferID\":\"231315213\","+
		"\"Status\":\"待确认\"}")

	// 反序列化json
	var data []map[string]interface{}
	if err := json.Unmarshal(bytes.NewBuffer(resp).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

// @Summary 按发起企业查询订单内容
// @Param Order body OrderByBuyerQueryRequestBody true "Order"
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/QueryOrderByBuyer [post]
func QueryOrderByBuyer(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(OrderByBuyerQueryRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	if body.Buyer != "" {
		bodyBytes = append(bodyBytes, []byte(body.Buyer))
	}
	//调用智能合约
	resp, err := bc.ChannelQuery("QueryOrderByBuyer", bodyBytes)
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

// @Summary 按供货商查询订单内容
// @Param Order body OrderBySellerQueryRequestBody true "Order"
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/QueryOrderBySeller [post]
func QueryOrderBySeller(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(OrderBySellerQueryRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	if body.Seller != "" {
		bodyBytes = append(bodyBytes, []byte(body.Seller))
	}
	//调用智能合约
	resp, err := bc.ChannelQuery("QueryOrderBySeller", bodyBytes)
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

// @Summary 更新是否接收状态
// @Param Order body UpdateOrderStatusRequestBody true "Order"
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/UpdateStatus [post]
func UpdateStatus(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(UpdateOrderStatusRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.ID))
	bodyBytes = append(bodyBytes, []byte(strconv.Itoa(body.Status)))

	//调用智能合约
	_, err := bc.ChannelExecute("UpdateStatus", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", fmt.Sprintf("更新是否接收状态成功"))
}

// @Summary 更新订单物流状态状态
// @Param Order body UpdateTransferStatusRequestBody true "Order"
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/UpdateTransferStatus [post]
func UpdateTransferStatus(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(UpdateTransferStatusRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.ID))
	bodyBytes = append(bodyBytes, []byte(body.TransferID))
	bodyBytes = append(bodyBytes, []byte(body.TransferStatus))

	//调用智能合约
	_, err := bc.ChannelExecute("UpdateStatus", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", fmt.Sprintf("更新订单物流状态成功"))
}

// @Summary 更改订单状态
// @Param Order body OrderByIDQueryRequestBody true "Order"
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/DeleteOrder [post]
func DeleteOrder(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(OrderByIDQueryRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.ID))

	//调用智能合约
	_, err := bc.ChannelExecute("DeleteOrder", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", fmt.Sprintf("删除成功"))
}
