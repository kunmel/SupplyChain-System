package routers
import (
	"github.com/togettoyou/blockchain-real-estate/chaincode/blockchain-real-estate/lib"
	"github.com/togettoyou/blockchain-real-estate/chaincode/blockchain-real-estate/utils"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"strconv"
)
func StoreOrder (stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var order lib.Order

	order.ID=args[0]
	order.Buyer=args[1]
	order.Seller=args[2]
	order.BuyerID =args[3]
	order.SellerID =args[4]
	order.Address=args[5]
	order.GoodsID =args[6]
	order.Goodsname =args[7]
	order.Goodsnum, _ =strconv.Atoi(args[8])
	order.Orderamount, _ =strconv.Atoi(args[9])
	order.Createtime =args[10]
	order.Deadline =args[11]
	order.Remark =args[12]
	order.Transferstatus, _ =strconv.Atoi(args[13])
	order.TransferID =args[14]
	order.Status, _ =strconv.Atoi(args[15])

	if err:=utils.WriteLedger(order,stub,lib.OrderKey,[]string{order.ID});err!=nil{
		return shim.Error(fmt.Sprintf("写入Order失败%s", err))
	}

	return shim.Success(nil)
}

func QueryOrder(stub shim.ChaincodeStubInterface  ) pb.Response {
	var orderList []lib.Order
	results, err := utils.GetStateByPartialCompositeKeys2(stub, lib.OrderKey, nil)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var order lib.Order
			err := json.Unmarshal(v, &order)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryOrder-反序列化出错: %s", err))
			}
			orderList = append(orderList, order)
		}
	}
	orderListByte, err := json.Marshal(orderList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryOrder-序列化出错: %s", err))
	}
	return shim.Success(orderListByte)
}


func QueryOrderByID(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	results, err := utils.GetStateByPartialCompositeKeys2(stub, lib.OrderKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	return shim.Success(results[0])
}

func QueryOrderByBuyer (stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args)!=1 {
		return shim.Error(fmt.Sprintf("参数个数!=1"))
	}
	var orderList []lib.Order
	results, err := utils.GetStateByPartialCompositeKeys2(stub, lib.OrderKey, nil)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var order lib.Order
			err := json.Unmarshal(v, &order)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryOrder-反序列化出错: %s", err))
			}
			if order.Buyer==args[0] {
				orderList = append(orderList, order)
			}
		}
	}
	orderListByte, err := json.Marshal(orderList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryOrder-序列化出错: %s", err))
	}
	return shim.Success(orderListByte)
}

func QueryOrderBySeller (stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var orderList []lib.Order
	results, err := utils.GetStateByPartialCompositeKeys2(stub, lib.OrderKey, nil)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var order lib.Order
			err := json.Unmarshal(v, &order)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryOrder-反序列化出错: %s", err))
			}
			if order.Seller==args[0] {
				orderList = append(orderList, order)
			}
		}
	}
	orderListByte, err := json.Marshal(orderList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryOrder-序列化出错: %s", err))
	}
	return shim.Success(orderListByte)
}

func UpdateStatus (stub shim.ChaincodeStubInterface, args []string) pb.Response {//arg[0]:id arg[1]:status
	results, err := utils.GetStateByPartialCompositeKeys2(stub, lib.OrderKey,[]string{args[0]})
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	var order lib.Order
	_=json.Unmarshal(results[0],&order)

	order.Status, _ =strconv.Atoi(args[1])

	if err:=utils.WriteLedger(order,stub,lib.OrderKey,[]string{order.ID});err!=nil{
		return shim.Error(fmt.Sprintf("写入order失败%s", err))
	}

	return shim.Success(nil)
}

func UpdateTransferStatus (stub shim.ChaincodeStubInterface, args []string) pb.Response {//arg[0]:id arg[1]:TransferStatus
	results, err := utils.GetStateByPartialCompositeKeys2(stub, lib.OrderKey,[]string{args[0]})
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	var order lib.Order
	_=json.Unmarshal(results[0],&order)

	order.Transferstatus, _ =strconv.Atoi(args[2])
	order.TransferID =args[1]

	if err:=utils.WriteLedger(order,stub,lib.OrderKey,[]string{order.ID});err!=nil{
		return shim.Error(fmt.Sprintf("写入order失败%s", err))
	}

	return shim.Success(nil)
}

func DeleteOrder(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	utils.DelLedger(stub,lib.OrderKey,args)
	return shim.Success(nil)
}
