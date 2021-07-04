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

type ProductRequestBody struct {
	GoodsID			string  `json:"GoodsID"` //货品ID
	Name 			string  `json:"Name"` //货品名称
	GoodType		int  `json:"GoodType"` //货品类型
	Amount			int     `json:"Amount"`   //货品数量
	Price 			int		`json:"Price"`//价格
	Standards       string  `json:"Standards"`       //规格
	Material       	string  `json:"Material"`        //材料
	Workmanship   	string  `json:"Workmanship"`   //制作工艺
	Supplier 		string  `json:"Supplier"` //供应商
	SupplierID		string	`json:"SupplierID"`
	AddDate        	string  `json:"AddDate"`       //上架时间
	Tel      		string  `json:"Tel"`        //联系方式
	ShippingArea    string  `json:"ShippingArea"`   //发货地区
	Remark 			string  `json:"Remark"`   //备注
}

type ProductByNameQueryRequestBody struct {
	GoodsID			string  `json:"GoodsID"` //货品ID
}

type ProductByAmountUpdateRequestBody struct {
	GoodsID			string  `json:"GoodsID"` //货品ID
	Amount			int     `json:"Amount"`   //货品数量
}

type ProductByGoodTypeQueryRequestBody struct {
	GoodType		string  `json:"GoodType"` //货品类型
}

type ProductByMaterialQueryRequestBody struct {
	Material       	string  `json:"Material"`        //材料
}

type ProductByWorkmanshipQueryRequestBody struct {
	Workmanship   	string  `json:"Workmanship"`   //制作工艺
}

type ProductBySupplierQueryRequestBody struct {
	Supplier 		string  `json:"Supplier"` //供应商
}

// @Summary 存储到数据库
// @Param Product body ProductRequestBody true "Product"
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/StoreProduct [post]
func StoreProduct(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(ProductRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	if body.Name == "" || body.Supplier == "" {
		appG.Response(http.StatusBadRequest, "失败", "Product货品名称Name和供应商Supplier不能为空")
		return
	}
	var bodyBytes1 [][]byte

	//调用智能合约
	resp1, err := bc.ChannelQuery("QueryProduct", bodyBytes1)
	if err != nil {
	   appG.Response(http.StatusInternalServerError, "查询失败", err.Error())
	   return
	}
	// 反序列化json
	var data1 []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp1.Payload).Bytes(), &data1); err != nil {
	   appG.Response(http.StatusInternalServerError, "反序列化失败", err.Error())
	   return
	}

	body.GoodsID = "product" + strconv.Itoa(len(data1)+1)
	fmt.Printf("map值为%v\n", body.GoodsID)
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.GoodsID))
	bodyBytes = append(bodyBytes, []byte(body.Name))
	bodyBytes = append(bodyBytes, []byte(strconv.Itoa(body.GoodType)))
	bodyBytes = append(bodyBytes, []byte(strconv.Itoa(body.Amount)))
	bodyBytes = append(bodyBytes, []byte(strconv.Itoa(body.Price)))
	bodyBytes = append(bodyBytes, []byte(body.Standards))
	bodyBytes = append(bodyBytes, []byte(body.Material))
	bodyBytes = append(bodyBytes, []byte(body.Workmanship))
	bodyBytes = append(bodyBytes, []byte(body.Supplier))
	bodyBytes = append(bodyBytes, []byte(body.SupplierID))
	bodyBytes = append(bodyBytes, []byte(body.AddDate))
	bodyBytes = append(bodyBytes, []byte(body.Tel))
	bodyBytes = append(bodyBytes, []byte(body.ShippingArea))
	bodyBytes = append(bodyBytes, []byte(body.Remark))

	//调用智能合约
	_, err = bc.ChannelExecute("StoreProduct", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", fmt.Sprintf("保存成功"))
}

// @Summary 更新货品数量
// @Param Product body ProductByAmountUpdateRequestBody true "Product"
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/UpdateProductByAmount [post]
func UpdateProductByAmount(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(ProductByAmountUpdateRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.GoodsID))
	bodyBytes = append(bodyBytes, []byte(strconv.Itoa(body.Amount)))

	//调用智能合约
	_, err := bc.ChannelExecute("UpdateProductByAmount", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", fmt.Sprintf("更新货品数量成功"))
}

// @Summary 查询全部货品
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/QueryProduct [post]
func QueryProduct(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(ProductByNameQueryRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte

	//调用智能合约
	resp, err := bc.ChannelQuery("QueryProduct", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	// 反序列化json
	var data []map[string]interface{}
	var data1 map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data1); err != nil{
			appG.Response(http.StatusInternalServerError, "失败", err.Error())
			return
		}else {
			appG.Response(http.StatusOK, "成功", data1)
		}
	}else {
		appG.Response(http.StatusOK, "成功", data)
	}
}

// @Summary 按货品ID查询
// @Param Product body ProductByNameQueryRequestBody true "Product"
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/QueryProductByID [post]
func QueryProductByID(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(ProductByNameQueryRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	if body.GoodsID != "" {
		bodyBytes = append(bodyBytes, []byte(body.GoodsID))
	}
	//调用智能合约
	resp, err := bc.ChannelQuery("QueryProductByID", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	// 反序列化json
	var data map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "查询成功", data)
}

// @Summary 按货品类型查询
// @Param Product body ProductByGoodTypeQueryRequestBody true "Product"
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/QueryProductByGoodType [post]
func QueryProductByGoodType(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(ProductByGoodTypeQueryRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	if body.GoodType != "" {
		if body.GoodType == "钢材" || body.GoodType == "木方" || body.GoodType == "烧结砖" {
			bodyBytes = append(bodyBytes, []byte(body.GoodType))
		} else {
			appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("不存在该货品类型"))
		}
	}
	//调用智能合约
	resp, err := bc.ChannelQuery("QueryProductByGoodType", bodyBytes)
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

// @Summary 按材料查询
// @Param Product body ProductByMaterialQueryRequestBody true "Product"
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/QueryProductByMaterial [post]
func QueryProductByMaterial(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(ProductByMaterialQueryRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	if body.Material != "" {
		bodyBytes = append(bodyBytes, []byte(body.Material))
	}
	//调用智能合约
	resp, err := bc.ChannelQuery("QueryProductByMaterial", bodyBytes)
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

// @Summary 按制作工艺查询
// @Param Product body ProductByWorkmanshipQueryRequestBody true "Product"
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/QueryProductByWorkmanship [post]
func QueryProductByWorkmanship(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(ProductByWorkmanshipQueryRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	if body.Workmanship != "" {
		bodyBytes = append(bodyBytes, []byte(body.Workmanship))
	}
	//调用智能合约
	resp, err := bc.ChannelQuery("QueryProductByWorkmanship", bodyBytes)
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

// @Summary 按供应商查询
// @Param Product body ProductBySupplierQueryRequestBody true "Product"
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/QueryProductBySupplier [post]
func QueryProductBySupplier(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(ProductBySupplierQueryRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	if body.Supplier != "" {
		bodyBytes = append(bodyBytes, []byte(body.Supplier))
	}
	//调用智能合约
	resp, err := bc.ChannelQuery("QueryProductBySupplier", bodyBytes)
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

// @Summary 更改订单状态
// @Param Product body ProductByNameQueryRequestBody true "Product"
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/DeleteProduct [post]
func DeleteProduct(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(ProductByNameQueryRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.GoodsID))

	//调用智能合约
	_, err := bc.ChannelExecute("DeleteProduct", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", fmt.Sprintf("删除成功"))
}

