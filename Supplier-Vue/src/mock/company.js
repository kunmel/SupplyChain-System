import Mock from 'mockjs'

const count = 30
const allOrder = []
const allProduct = []
for(let i = 0; i < count; i++) {
    allOrder.push(Mock.mock({
        ID:Mock.mock('@id'),
        "Buyer|1": [
            "上海国金金属材料有限公司",
            "南昌市赣达金属材料有限公司",
            "江苏景同钢铁有限公司",
            "东莞市金力盛精密五金有限公司",
            "福州九川商贸有限公司",
            "苏州铭越金属材料有限公司",
          ],
          "Seller|1":[
            "苏州东锜精密模具",
            "南昌市赣达金属材料有限公司",
            "冉海钢材",
            "包头中旺钢材有限公司",
            "天津市中鼎凯达贸易有限公司",
            "聊城乾盛钢材有限公司",
            "唐山市丰润区达江商贸有限公司",
            "云南虎鑫经贸有限公司",
            "东莞市台康金属材料有限公司",
            "上海国金金属材料有限公司",
          ],
        BuyerID: Mock.mock('@id'),
        SellerID:Mock.mock('@id'),
        Address:Mock.mock('@county(true)'),
        GoodsID: Mock.mock('@id'),
        Goodsname:'钢材',
        Goodsnum: Mock.mock('@integer(50, 2000)')+'吨',
        Orderamount:Mock.mock('@integer(50, 200)')+'万元',
        Createtime:Mock.mock('@date("2021-MM-dd")'),
        Deadline:Mock.mock('@date("2021-MM-dd")'),
        Remark:'无',
        "Transferstatus|1":[
            '0',
            '1',
            '2'
        ],
        TransferID:Mock.mock('@id'),
        "Status|1":[
            '0',
            '1',
            '2',
            '3',
            '4'
        ],
    }))
    allProduct.push(Mock.mock({
        GoodsID:Mock.mock('@id'),
        "Name|1":['圆钢','圆坯','模具钢','优特圆钢'],
        Goodtype:'钢材',
        Amount:Mock.mock('@integer(50, 2000)')+'吨',
        Price:'100',
        "Standards|1":[
            "Φ6.5",
            "Φ48",
            "Φ5.5",
            "Φ120",
            "Φ210",
            "Φ18",
            "Φ300",
            "Φ180",
            "Φ28",
            "Φ200",
            "Φ27"
          ],
        Material:'精钢',
        Workmanship:'冷拉',
        "Supplier|1":[
            "苏州东锜精密模具",
            "南昌市赣达金属材料有限公司",
            "冉海钢材",
            "包头中旺钢材有限公司",
            "天津市中鼎凯达贸易有限公司",
            "聊城乾盛钢材有限公司",
            "唐山市丰润区达江商贸有限公司",
            "云南虎鑫经贸有限公司",
            "东莞市台康金属材料有限公司",
            "上海国金金属材料有限公司",
          ],
        SupplierID:Mock.mock('@id'),
        Adddate:Mock.mock('@date("2021-MM-dd")'),
        Tel:Mock.mock('@integer(13503203793, 15698190726)'),
        Shippingarea:Mock.mock('@county(true)'),
        Remark:'无'
}))
  }

  function StoreOrder(orderData) {
  }

  export { allOrder,allProduct,StoreOrder}