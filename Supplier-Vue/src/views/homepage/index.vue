<template>
  <div class="homepage-container">

    <el-row class="home-total">
      <el-col :xs="12" :sm="12" :md="12" :lg="6" :xl="6" class="home-total-item" v-for="(item, index) of homeTotalData" :key="'line-' + index">
        <div class="wrapper-item">
          <p class="title">{{item.title}}</p>
          <p class="value digital-number" ref="countup">{{item.value}}</p>
          <color-line :id='"main"+index' :color="item.color" :optionData="item.data" width="180px" height="70px"></color-line>
        </div>
      </el-col>
    </el-row>
    <el-row class="home-part1" :gutter="20">
        <el-col :span="6" >
        <div class="home-detail-item" style="background: #ec407a" @click="gotoUrll('/orderlist/allorder')">
          <div class="name">{{'查看订单'}}</div>
        </div>
        </el-col>
        <el-col :span="6">
        <div class="home-detail-item" style="background: #ab47bc"  @click="gotoUrll('/goods/allogoods')">
          <div class="name">{{'查看商品'}}</div>
        </div>
        </el-col>
        <el-col :span="6">
        <div class="home-detail-item" style="background: #2196f3"   @click="gotoUrll('/finance/all')" >
          <div style="padding: 30px 0 0px 0;text-align: center;font-size: 35px;color: white;">{{'融资管理'}}</div>
          <!-- <div style="padding: 10px 0 0px 0;text-align: center;font-size: 30px;color: white;">{{'待改'}}</div> -->
        </div>
        </el-col>
    </el-row>

    <el-row class="home-part2" :gutter="0">
      <el-col :span="12">
        <div class="financing-sprinkled">
          <div class="title">
            <p class="title-value">平台近6个月的交易记录</p>
          </div>
          <div class="content" ref="">
            <near-six-month width="100%" height="100%"></near-six-month>
          </div>
        </div>
      </el-col>
      <el-col :span="12">
        <div class="financing-sprinkled">
          <div class="title">
            <p class="title-value">借贷金额及融资期限分布图示</p>
          </div>
          <div class="content" ref="">
            <!-- 投资 -->
            <div class="investment">
              <span class="title">借贷金额比例</span>
              <investment-pie width="100%" height="50%"></investment-pie>
              <div class="detail">
                <span class="detail-item">
                  1万元以下
                  <br>
                  33.04%
                </span>
                <span class="detail-item">
                  1-10万
                  <br>
                  30.57%
                </span>
                <span class="detail-item">
                  10-40万
                  <br>
                  23.08%
                </span>
                <span class="detail-item">
                  40万以上
                  <br>
                  13.31%
                </span>
              </div>
            </div>
            <!-- 融资 -->
            <div class="financing">
              <span class="title">借贷期限</span>
              <financing-pie width="100%" height="50%"></financing-pie>
              <div class="detail">
                <span class="detail-item">
                  0-3个月
                  <br>
                  18.91%
                </span>
                <span class="detail-item">
                  3-6个月
                  <br>
                  29.41%
                </span>
                <span class="detail-item">
                  6-12个月
                  <br>
                  32.77%
                </span>
                <span class="detail-item">
                  12个月以上
                  <br>
                  18.91%
                </span>
              </div>
            </div>
          </div>
        </div>
      </el-col>
    </el-row>

  </div>
</template>
<script>
  import CountUp from 'countup.js'
  import {getHomeTotal, getrefreshDate, getRank} from '@/api/homepage'
  import ColorLine from '@/components/color-line'
  import NearSixMonth from '@/views/homepage/near-six-month'
  import BScroll from 'better-scroll'
  import InvestmentPie from '@/views/homepage/investment-pie'
  import FinancingPie from '@/views/homepage/financing-pie'
  export default {
    components: {
      ColorLine,
      NearSixMonth,
      InvestmentPie,
      FinancingPie
    },
    data() {
      return {
        homeTotalData: [],
        refreshDate: { },
        rankList: [],
        numAnim: null
      }
    },
    methods: {
      gotoUrll(url) {
        this.$router.push(url)
      },
      initCountUp() {
        this.$nextTick(() => {
          let countupLength = this.$refs.countup.length
          let i = 0
          for (i; i < countupLength; i++) {
            this.numAnim = new CountUp(this.$refs.countup[i], 0, this.$refs.countup[i].innerText, 2, 1.5)
            this.numAnim.start()
          }
        })
      },
      _initScroll() {
        if (!this.scroll) {
          this.scroll = new BScroll(this.$refs.rankContent, {
            scrollY: true,
            click: true,
            scrollbar: {
              fade: false,
              interactive: true // 1.8.0 新增
            },
            mouseWheel: {
              speed: 20,
              invert: false,
              easeTime: 300
            }
          })
        } else {
          this.scroll.refresh()
        }
      }
    },
    created() {
      // 获取头部hometotal
      getHomeTotal().then((resp) => {
        this.homeTotalData = resp.data
        this.initCountUp()
      }).catch(() => {
        console.log('获取home-total出现异常')
      })
      // 获取 detailItem
      // getrefreshDate().then(resp => {
      //   this.refreshDate = resp.data
      // }).catch(() => {
      //   console.log('获取detailItem出现异常')
      // })
      // // 获取投资榜
      // getRank().then(resp => {
      //   this.rankList = resp.data
      //   this._initScroll()
      // }).catch(() => {
      //   console.log('获取rankList出现异常')
      // })
    },
    mounted() {},
    updated() {
      // this.$nextTick(function() {
      //   this.initCountUp()
      // })
    }
  }
</script>
<style scoped lang="stylus">
.homepage-container
  min-width 580px

.home-total {
  width: 100%;
  height: 160px;
  border: 1px solid #ddd;
  border-radius: 4px;
  margin: 0 0 15px 0;
  .home-total-item {
    box-sizing: border-box;
    display: inline-block;
    height: 100%;
    padding: 15px 0;
    vertical-align: top;
    .wrapper-item {
      height: 100%;
      padding: 0 20px;
      border-right: 1px solid #ccc;
      text-align: center;
      .title {
        margin: 0px 0;
        font-size : 20px;
      }
      .value {
        margin 5px 0
        font-size 34px
        color: #ffc107
      }
    }
    &:last-child {
      .wrapper-item {
        border: none;
      }
    }
  }
}
.home-part1 {
  margin: 50px 0px 50px 0px;
  left: 200px;
}
.home-part2 {
  margin-top: 15px;
  .financing-sprinkled {
    border: 1px solid #eee;
    height: 350px;
    .title {
      background: #dde3ef;
      padding: 10px 0;
      .title-value {
        margin-left: 4px;
        text-indent: 4px;
        color: #666;
        &:before {
          display: inline-block;
          content: '';
          width: 4px;
          height: 16px;
          background: purple;
          margin-right: 4px;
          border-radius: 4px;
          vertical-align: middle;
        }
      }
    }
    .content {
      display: inline-flex;
      width: 100%;
      height: 310px;
      .investment {
        height: 310px;
        width: 50%;
        .title {
          display: inherit;
          text-align: center;
          background: transparent;
          padding-top: 20px;
        }
        .detail {
          margin-left: 10px;
          text-align: center;
          .detail-item {
            display: inline-block;
            width: 40%;
            margin: 5px;
            padding-left: 5px;
            border-left: 5px solid #ccc;
            color: #666;
          }
        }
      }
      .financing {
        height: 310px;
        width: 50%;
        .title {
          display: inherit;
          text-align: center;
          background: transparent;
          padding-top: 20px;
        }
        .detail {
          margin-left: 10px;
          text-align: center;
          .detail-item {
            display: inline-block;
            width: 40%;
            margin: 5px;
            padding-left: 5px;
            border-left: 5px solid #ccc;
            color: #666;
          }
        }
      }
    }
  }
  .bad-debt {
    height: 350px;
    margin-left: 10px;
    border: 1px solid #eee;

    .title {
      background: #dde3ef;
      padding: 10px 0;
      .title-value {
        margin-left: 4px;
        text-indent: 4px;
        color: #666;
        &:before {
          display: inline-block;
          content: '';
          width: 4px;
          height: 16px;
          background: purple;
          margin-right: 4px;
          border-radius: 4px;
          vertical-align: middle;
        }
      }
    }
    .content {
      height: inherit;
      .bad {
        height: 50%;
        padding: 20px 15px;
        .total {
          display: inline-block;
          width: 30%
          color: #666;
          vertical-align: top;
          .total1 {
            text-align: center;
            .num {
              font-size: 24px;
            }
          }
          .total2 {
            text-align: center;
            margin-top: 20px;
            .num {
              font-size: 24px;
            }
          }
        }
        .chart {
          display: inline-block;
          width: 68%;
          .title {
            background: none;
            border-bottom: 1px solid #ccc;
          }
          .line {
            border-bottom: 1px solid #ccc;
            padding-bottom: 30px;
            &:last-child {
              border-bottom-color: #000;
            }
          }
          &:after {
            content: '0';
            position: relative;
            font-size: 70px;
            left: 20px;
            top: -70px;
            color: #ddd;
          }
        }
      }
      .overdue {
        padding: 20px 15px;
        height: 50%;
        .total {
          display: inline-block;
          width: 30%
          color: #666;
          vertical-align: top;
          .total1 {
            text-align: center;
            .num {
              font-size: 24px;
            }
          }
          .total2 {
            text-align: center;
            margin-top: 20px;
            .num {
              font-size: 24px;
            }
          }
        }
        .chart {
          display: inline-block;
          width: 68%
          .title {
            background: none;
            border-bottom: 1px solid #ccc;
          }
          .line {
            border-bottom: 1px solid #ccc;
            padding-bottom: 30px;
            &:last-child {
              border-bottom-color: #000;
            }
          }
          &:after {
            content: '0';
            position: relative;
            font-size: 70px;
            left: 20px;
            top: -70px;
            color: #ddd;
          }
        }
      }
    }
  }
}
.el-col {
  border-radius: 4px;
}
.bg-purple-dark {
  background: #99a9bf;
}
.bg-purple {
  background: #d3dce6;
}
.bg-purple-light {
  background: #e5e9f2;
}
.grid-content {
  border-radius: 4px;
  min-height: 36px;
}
.home-detail-item {
      flex: 0 0 48%
      height: 145px
      border: 1px solid #eee
      background-image linear-gradient(rgba(255, 255, 255, .1), rgba(255, 255, 255, .3)) !important
      cursor pointer
}
.name {
        padding: 40px 0 0px 0;
        text-align: center;
        font-size: 35px;
        color: white;
}
.content /deep/ .bscroll-vertical-scrollbar
  z-index 2 !important

</style>
