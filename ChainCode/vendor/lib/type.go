package lib
type Order struct {
	ID             string
	Buyer          string
	Seller         string
	BuyerID        string
	SellerID       string
	Address        string
	GoodsID        string
	Goodsname      string
	Goodsnum       int
	Orderamount    int
	Createtime     string
	Deadline       string
	Remark         string
	Transferstatus int
	TransferID     string
	Status         int
}

type Product struct {
	goodsID string
	name string
	goodtype string
	amount int
	price int
	standards string
	material string
	workmanship string
	supplier string
	supplierID string
	adddate string
	tel string
	shippingarea string
	remark string
}
type User struct {
	ID string
	account string
	password string
	identity int
	name string
}

type Financing struct {
	supplier string
	amount string
	period string
	desc string//申请描述
	orderID string
	status string// ("0":待审批 "1":审批通过 "2":审批拒绝)
	feedback int
}

