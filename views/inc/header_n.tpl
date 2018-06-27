<div class="header" id="header_n" v-cloak>
    <div class="header-t">
        <div class="box w1190">
            <img src="/static/img/img-01.png " alt="">
        </div>
    </div>
    <div class="header-m">
        <div class="box w1190">
            <p>您好，123消费网欢迎您的来到！</p>
            <h2>400-888-8888</h2>
            <ul>
                <li><a href="/account/loginTpl">登陆</a></li>
                <li><a href="/account/registerTpl">注册</a></li>
                <li><a href="">兑换商品</a></li>
                <li><a href="">合作商家</a></li>
                <li><a href="">帮助中心</a></li>
                <li><a href="">服务热线</a></li>
            </ul>
        </div>
    </div>
    <div class="header-d">
        <div class="box w1190">
            <div class="logo">
                <a href="/ "><img src="/static/img/img-01.jpg " alt=""></a>
            </div>
            <div class="search">
                <div class="top">
                    <div class="search-box-shop" style="display: block;">
                        <i-select size="large" v-model="selectType" style="width:100px;" class ='select'>
                            <i-option v-for="item in typeList" :value="item.value" :key="item.value" style="float: none;">{{ item.label }}</i-option>
                        </i-select>
                        <i-input style="width: 280px;margin-left: 10px;margin-right:10px" size="large" class="te" v-model="keyWord" type="text" placeholder="请输入搜索内容"></i-input>
                        <i-button type="ghost" shape="circle" icon="ios-search" @click="searchPro()" class="sub">Search</i-button>
                    </div>
                    <div class="wel-shop-car">
                        <a href="/shopcar/shopcar "><p>购物车</p></a>
                        <div class="rad">10</div>
                    </div>
                </div>
                <div class="down">
                    <ul>
                        <li v-for="item in selectHis"><a href="/product_list/info ">{{item.label}}</a></li>
                    </ul>
                </div>
            </div>
            <div class="two-img">
                <img src="/static/img/img-02.jpg" alt="">
            </div>
        </div>
    </div>
    <div class="header-nav">
        <div class="box w1190">
            <div class="box w1190">
                <ul>
                    <li class="li-sp">
                        <a href="/"><p>首页</p></a>
                    </li>
                    <li v-for="items in meunList">
                        <a :href="'/product_list/list/?menu_id='+ item.id"  v-for="item in items">{{item.label}}</a>
                    </li>
                    <li>
                        <a href="/product_list/list">今日特价</a>
                        <a href="/product_list/list ">新货上架</a>
                    </li>
                    <li>
                        <a href="/text">使用指南</a>
                        <a href="/shop_list/list">获取免费积分合作商家</a>
                    </li>
                </ul>
            </div>
        </div>
    </div>
</div>
<script>
    new Vue({
        el: '#header_n',
        data: {
            typeList: [
                {
                    value: 1,
                    label: '商品'
                },
                {
                    value: 2,
                    label: '积分'
                },
                {
                    value: 3,
                    label: '测试'
                }
            ],
            selectHis: [
                {
                    id: 1,
                    label: '扫地机器人'
                },
                {
                    id: 2,
                    label: '皮鞋'
                },
                {
                    id: 3,
                    label: '换鞋凳'
                }
            ],
            meunList: [
                [
                    {
                        id: 1,
                        label: '一、免费积分A商城'
                    },
                    {
                        id: 2,
                        label: '二、积分币A商城'
                    },
                ],
                [
                    {
                        id: 3,
                        label: '三、团购商城'
                    },
                    {
                        id: 4,
                        label: '四、抢购商城'
                    },
                ],
                [
                    {
                        id: 5,
                        label: '五、免费积分B商城'
                    },
                    {
                        id: 6,
                        label: '六、积分币B商城'
                    },
                ],
                [
                    {
                        id: 7,
                        label: '七、抢购商城'
                    },
                    {
                        id: 8,
                        label: '八、进口商城'
                    },
                ],
            ],
            keyWord: '',
            selectType: 1
        },
        methods: {
            searchPro:function () {
                console.log(this.keyWord)
                console.log(this.selectType)
            },
        }
    })
</script>