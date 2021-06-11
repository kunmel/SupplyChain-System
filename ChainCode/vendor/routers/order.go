package routers
import (
	"blockchain-real-estate-master/chaincode/blockchain-real-estate/lib"
	"encoding/json"
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

	key:=order.ID
	bytes, _ := json.Marshal(order)
	stub.PutState(key, bytes)
}

func QueryOrder(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	bytes,_:=stub.GetState(args[0])
	return shim.Success(bytes)
}

func QueryOrderByBuyer (stub shim.ChaincodeStubInterface, args []string) pb.Response {
}

func QueryOrderBySeller (stub shim.ChaincodeStubInterface, args []string) pb.Response {

}

func UpdateStatus (stub shim.ChaincodeStubInterface, args []string) pb.Response {
	bytes,_:=stub.GetState(args[0])
	var order lib.Order
	_=json.Unmarshal(bytes,&order)
	order.Status, _ =strconv.Atoi(args[1])
	key:=args[0]
	bytes, _ = json.Marshal(order)
	stub.PutState(key, bytes)
	return shim.Success(nil)
}

func UpdateTransferStatus (stub shim.ChaincodeStubInterface, args []string) pb.Response {
	bytes,_:=stub.GetState(args[0])
	var order lib.Order
	_=json.Unmarshal(bytes,&order)
	order.Transferstatus, _ =strconv.Atoi(args[1])
	key:=args[0]
	bytes, _ = json.Marshal(order)
	stub.PutState(key, bytes)
	return shim.Success(nil)
}

func DeleteOrder(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	key:=args[0]
	stub.DelState(key)
	return shim.Success(nil)
}

