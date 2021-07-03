package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"lib"
	"routers"
	"utils"
	//"github.com/togettoyou/blockchain-real-estate/chaincode/blockchain-real-estate/lib"
	//"github.com/togettoyou/blockchain-real-estate/chaincode/blockchain-real-estate/routers"
	//"github.com/togettoyou/blockchain-real-estate/chaincode/blockchain-real-estate/utils"
	"time"
)

type BlockChainRealEstate struct {
}

// Init 链码初始化
func (t *BlockChainRealEstate) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("链码初始化")
	//初始化默认数据
	fmt.Println("-----------初始化三种类型的用户-------------")
	user1:=lib.User{Account: "qiye",Password: "123",ID: "user1",Identity: 1,Name: "江苏景同钢铁有限公司"}
	if err:=utils.WriteLedger(user1,stub,lib.UserKey,[]string{user1.ID});err!=nil{
		return shim.Error(fmt.Sprintf("写入User失败%s", err))
	}

	user2:=lib.User{Account: "gongyingshang",Password: "123",ID: "user2",Identity: 2,Name:"唐山市丰润区达江商贸有限公司"}
	if err:=utils.WriteLedger(user2,stub,lib.UserKey,[]string{user2.ID});err!=nil{
		return shim.Error(fmt.Sprintf("写入User失败%s", err))
	}

	user3:=lib.User{Account: "yinhang",Password: "123",ID: "user3",Identity: 3,Name: "某家银行"}
	if err:=utils.WriteLedger(user3,stub,lib.UserKey,[]string{user3.ID});err!=nil{
		return shim.Error(fmt.Sprintf("写入User失败%s", err))
	}

	fmt.Println("-----------初始化货品信息-------------")
	product1:=lib.Product{GoodsID: "product1",Name: "钢材",Goodtype:1,Amount: 16543,Price: 534,Standards: "#210",Material: "木头",
		Workmanship: "锻造",Supplier:"唐山市丰润区达江商贸有限公司",SupplierID: "user2",Adddate: "2020-6-6",Tel: "13838383838",Shippingarea: "香港特别行政区 九龙 黄大仙区",Remark: "无"}
	if err:=utils.WriteLedger(product1,stub,lib.ProductKey,[]string{product1.GoodsID});err!=nil{
		return shim.Error(fmt.Sprintf("写入Product失败%s", err))
	}

	fmt.Println("-----------初始化融资信息-------------")
	financing1:=lib.Financing{Supplier: "唐山市丰润区达江商贸有限公司",Amount: 94,Period:"2021-02-14-2022-07-31",
		Desc: "本人丁磊经营一铁件加工厂，3年来经营状况良好。2020年营业额46万元。 本人库存钢材价值29万元，加工成品90万元。现因资金周转紧张，特申请贷款，请给予批准。",
		OrderID: "31000019860811337X",Status: 0,Feedback: "无"}
	if err:=utils.WriteLedger(financing1,stub,lib.FinancingKey,[]string{financing1.OrderID,financing1.Period});err!=nil{
		return shim.Error(fmt.Sprintf("写入Financing失败%s", err))
	}

	fmt.Println("-----------初始化订单-------------")
	order1:=lib.Order{ID: "31000019860811337X",Buyer: "江苏景同钢铁有限公司",Seller: "唐山市丰润区达江商贸有限公司",BuyerID:"user1",
		SellerID: "user2",Address: "湖北省 鄂州市 其它区",GoodsID:"product1",Goodsname: "钢材",Goodsnum: 1627,Orderamount: 69,
		Createtime: "2021-11-06",Deadline: "2021-07-24",Remark: "无",Transferstatus:2,TransferID: "360000200206083832",Status: 2}
	if err:=utils.WriteLedger(order1,stub,lib.OrderKey,[]string{order1.ID});err!=nil{
		return shim.Error(fmt.Sprintf("写入Order失败%s", err))
	}
	return shim.Success(nil)
}

// Invoke 实现Invoke接口调用智能合约
func (t *BlockChainRealEstate) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	funcName, args := stub.GetFunctionAndParameters()
	switch funcName {
	case "StoreOrder":
		return routers.StoreOrder(stub,args)
	case "QueryOrder":
		return routers.QueryOrder(stub)
	case "QueryOrderByID":
		return routers.QueryOrderByID(stub, args)
	case "QueryOrderByBuyer":
		return routers.QueryOrderByBuyer(stub, args)
	case "QueryOrderBySeller":
		return routers.QueryOrderBySeller(stub, args)
	case "UpdateStatus":
		return routers.UpdateStatus(stub, args)
	case "UpdateTransferStatus":
		return routers.UpdateTransferStatus(stub, args)
	case "DeleteOrder":
		return routers.DeleteOrder(stub, args)
	case "StoreProduct":
		return routers.StoreProduct(stub, args)
	case "UpdateProductByAmount":
		return routers.UpdateProductByAmount(stub, args)
	case "QueryProduct":
		return routers.QueryProduct(stub)
	case "QueryProductByID":
		return routers.QueryProductByID(stub, args)
	case "QueryProductByGoodType":
		return routers.QueryProductByGoodType(stub, args)
	case "QueryProductByMaterial":
		return routers.QueryProductByMaterial(stub, args)
	case "QueryProductByWorkmanship":
		return routers.QueryProductByWorkmanship(stub, args)
	case "QueryProductBySupplier":
		return routers.QueryProductBySupplier(stub, args)
	case "DeleteProduct":
		return routers.DeleteProduct(stub, args)
	case "StoreFinance":
		return routers.StoreFinance(stub, args)
	case "UpdateFinanceStatus":
		return routers.UpdateFinanceStatus(stub, args)
	case "QueryFinance":
		return routers.QueryFinance(stub)
	case "QueryFinanceByStatus":
		return routers.QueryFinanceByStatus(stub, args)
	case "QueryFinanceByOrderID":
		return routers.QueryFinanceByOrderID(stub, args)
	case "StoreUser":
		return routers.StoreUser(stub, args)
	case "QueryUser":
		return routers.QueryUser(stub)
	case "QueryUserByID":
		return routers.QueryUserByID(stub,args)
	default:
		return shim.Error(fmt.Sprintf("没有该功能: %s", funcName))
	}
}

func main() {
	err := shim.Start(new(BlockChainRealEstate))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
