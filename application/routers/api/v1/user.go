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

type User struct {
	ID 				string `json:"ID"`  //ID
	Account  		string  `json:"account"` //账号
	Password		string  `json:"password"` //密码
	Identity       	string  `json:"identity"`       //身份
	Name   			string  `json:"name"`   //具体名字
}

type UserTest struct {
	Account string `json:"account"`
	Password  string    `json:"password"`
}

type UserByIDQueryRequestBody struct {
	Account  		string  `json:"account"` //账号
	Password		string  `json:"password"` //密码
}

// @Summary 存储到数据库
// @Param user body User true "user"
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/StoreOrder [post]
func StoreUser(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(User)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	var bodyBytes1 [][]byte
	bodyBytes1 = append(bodyBytes1, []byte(body.Account))

	//调用智能合约
	resp1, err := bc.ChannelQuery("QueryUser", bodyBytes1)
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

	var lenIdent = 0
	for i:=0; i<len(data1); i++{
		for _, v := range data1[i]{
			if v == body.Identity {
				lenIdent++
			}
		}
		body.ID = "user" + body.Identity +strconv.Itoa(lenIdent)
		fmt.Printf("map值为%v\n", body.ID)
	}

	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.ID))
	bodyBytes = append(bodyBytes, []byte(body.Account))
	bodyBytes = append(bodyBytes, []byte(body.Password))
	bodyBytes = append(bodyBytes, []byte(body.Identity))
	bodyBytes = append(bodyBytes, []byte(body.Name))

	//调用智能合约
	_, err = bc.ChannelExecute("StoreUser", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}

	appG.Response(http.StatusOK, "成功", fmt.Sprintf("注册成功"))
}

// @Summary 根据账号查询
// @Param	user body UserByIDQueryRequestBody true "user"
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/QueryUser [post]
func QueryUser(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(UserByIDQueryRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.Account))

	//调用智能合约
	resp, err := bc.ChannelQuery("QueryUser", bodyBytes)
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

	var existAcc = 0
	for i:=0; i<len(data); i++{
		//fmt.Printf("%s\n", m[i])
		for _, v := range data[i]{
			if v == body.Account {
				existAcc++
				for _, va := range data[i]{
					if va == body.Password {
						delete(data[i], "Password")
						delete(data[i], "Name")
						appG.Response(http.StatusOK, "密码正确", data[i])
						fmt.Printf("map值为%v\n", data[i])
						return
					} else {
						appG.Response(http.StatusBadRequest, "错误", fmt.Sprintf("密码错误"))
						return
					}
				}
			}
		}
	}
	if existAcc == 0 {
		appG.Response(http.StatusBadRequest, "错误", fmt.Sprintf("不存在该用户"))
		return
	}
}

// @Summary 根据账号查询
// @Param	user body UserByIDQueryRequestBody true "user"
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/QueryUserTest [post]
func QueryUserTest(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(UserByIDQueryRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	fmt.Printf("传出的参数为：%s\n",body.Account)

	//var resp []byte
	u1 := UserTest{Account: "admin", Password: "123456"}
	resp, _ := json.Marshal(&u1)
	if body.Account == "admin" {
		// 反序列化json
		var data map[string]interface{}
		_ = json.Unmarshal(resp, &data)
		appG.Response(http.StatusOK, "成功", data)
	} else {
		appG.Response(http.StatusBadRequest, "失败", "不存在该用户")
		return
	}


}

// @Summary 根据账号查询
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/QueryUserTestGet [get]
func QueryUserTestGet(c *gin.Context) {
	appG := app.Gin{C: c}

	u1 := UserTest{Account: "admin", Password: "123456"}

	b, _ := json.Marshal(&u1)
	fmt.Printf("b=%s\n", string(b))
	var data map[string]interface{}
	_ = json.Unmarshal(b, &data)
	for k, v := range data{
		fmt.Printf("key:%v value:%v\n", k, v)
	}

	appG.Response(http.StatusOK, "成功", data)
}
