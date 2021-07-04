import Mock from 'mockjs'

const count = 30
const mergeTableData = []
const GoodsTable = []

const userFiannceData = []

for(let i = 0; i < 10; i++) {
  userFiannceData.push(Mock.mock({
    "Supplier|1" :[
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
    Amount:Mock.mock('@integer(1, 120)') + '万',
    Period:Mock.mock('@date("2021-MM-dd")') + '-' + Mock.mock('@date("2022-MM-dd")') ,
    "Desc|1" : [
      "本人"+Mock.mock('@cname')+"经营一铁件加工厂，3年来经营状况良好。2020年营业额"+Mock.mock('@integer(1, 120)')+"万元。 本人库存钢材价值"+Mock.mock('@integer(1, 30)')+"万元，加工成品"+Mock.mock('@integer(1, 120)')+"万元。现因资金周转紧张，特申请贷款，请给予批准。",

    ],
    OrderID:Mock.mock('@natural'),
    "Status|1" :[
        "1",
        "2"
      ],
    Feedback:''
  }
  ))
}

for(let i = 0; i < count; i++) {
  mergeTableData.push(Mock.mock({

    ID:Mock.mock('@natural'),

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

    Address:Mock.mock('@county(true)'),

    GoodsID: Mock.mock('@natural'),

    Goodsname:'钢材',

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

    Goodsnum: Mock.mock('@integer(1, 120)') + '吨',

    Orderamount:Mock.mock('@integer(1, 30)') + '万',

    Createtime:Mock.mock('@date("2021-MM-dd")') ,

    Deadline:Mock.mock('@date("2022-MM-dd")') ,

    Remark:'备注',

    "Status|1" :[
      "0",
      "1",
      "2",
      "3",
      "4"
    ],
    "Transferstatus|1" : [
      "0",
      "1",
      "2"
    ],
    TransferId : Mock.mock('@natural'),
  }))
}

for(let i = 0; i < count; i++) {
  GoodsTable.push(Mock.mock({
    GoodsID:Mock.mock('@natural'),
    "Name|1": ['钢材'],
    "Goodtype|1" :[
      "1",
      "2",
      "3",
    ],
    
    Amount: Mock.mock('@integer(1,100)') + '吨',
    Price:Mock.mock('@float(0.1, 5,1,2)')+"万",
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
    "Material|1": ['精钢','塑料','木头'],
    "Workmanship|1":['锻造','合成'],
    Supplier:'鞍山钢铁集团',
    Adddate:'2020-6-6',
    Tel:'13838383838',
    Shippingarea:Mock.mock('@county(true)'),
    Remark:'无'
  }))
}

mergeTableData.forEach(data => {
  data.compRatio = `${data.passcount}/${data.total}`
  data.implementRatio = `${data.getover}/${data.total}`
})

export function delFiles(ids) {
  let length = ids.length || 0
  fileCount -= length
  return true
}

export {  mergeTableData,GoodsTable,userFiannceData}
