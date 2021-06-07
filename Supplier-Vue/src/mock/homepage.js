import Mock from 'mockjs'

export function getHomeTotal() {
  return [
    {
      title: '订单总量',
      value: Mock.mock('@natural(0, 9999999)'),
      color: Mock.mock('@hex'),
      data: [Mock.mock('@integer(1, 120)'), Mock.mock('@integer(1, 120)'), Mock.mock('@integer(1, 120)'), Mock.mock('@integer(1, 120)'), Mock.mock('@integer(1, 120)'), Mock.mock('@integer(1, 120)'), Mock.mock('@integer(1, 120)')]
    },
    {
      title: '待确认订单量',
      value: Mock.mock('@natural(0, 9999999)'),
      color: Mock.mock('@hex'),
      data: [Mock.mock('@integer(1, 120)'), Mock.mock('@integer(1, 120)'), Mock.mock('@integer(1, 120)'), Mock.mock('@integer(1, 120)'), Mock.mock('@integer(1, 120)'), Mock.mock('@integer(1, 120)'), Mock.mock('@integer(1, 120)')]
    },
    {
      title: '商品总量',
      value: Mock.mock('@natural(0, 9999999)'),
      color: Mock.mock('@hex'),
      data: [Mock.mock('@integer(1, 120)'), Mock.mock('@integer(1, 120)'), Mock.mock('@integer(1, 120)'), Mock.mock('@integer(1, 120)'), Mock.mock('@integer(1, 120)'), Mock.mock('@integer(1, 120)'), Mock.mock('@integer(1, 120)')]
    },
    {
      title: '总营收',
      value: Mock.mock('@natural(0, 9999999)'),
      color: Mock.mock('@hex'),
      data: [Mock.mock('@integer(1, 120)'), Mock.mock('@integer(1, 120)'), Mock.mock('@integer(1, 120)'), Mock.mock('@integer(1, 120)'), Mock.mock('@integer(1, 120)'), Mock.mock('@integer(1, 120)'), Mock.mock('@integer(1, 120)')]
    }
  ]
}

export function getrefreshDate() {
  return [
    {
      date: '2021',
      color: '#ec407a'
    }
  ]
}

const rankList = []
const count = 30
for (let i = 0; i < count; i++) {
  rankList.push(Mock.mock({
    name: '@cname',
    value: '@natural(0, 9999999)',
    avatar: 'https://avatars.githubusercontent.com/u/' + '@natural(1000000, 9999999)'
  }))
}
export function getHomeInvestmentRank() {
  return rankList
}
