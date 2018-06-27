<div style="background:url('/static/img/img-43.jpg' ) no-repeat center" class="log-in" id="loginTpl" v-cloak>
    <div class="w1190">
        <div class="log-in-box">
            <div class="log-in-boxn">
                <h2>会员登录</h2>
                <form>
                    <i-input placeholder="请输入用户名" style="background:url('/static/img/img-21.png' )  no-repeat 10px center;border: none;" v-model="username" size="large" class="input-a" type="text"></i-input>
                    <i-input placeholder="请输入密码" style="background:url('/static/img/img-22.png') no-repeat 12px center;border: none;" v-model="password"  size="large" class="input-a" type="password"></i-input>
                    <i-button type="success" long @click="doLogin()" style="font-size: 18px;">登录</i-button>
                    <div class="two-a">
                        <a href="/account/registerTpl">免费注册</a>
                        <a href="">忘记密码</a>
                    </div>
                </form>
            </div>
        </div>
    </div>
    <Alert type="error" show-icon v-show="is_error">{{errMsg}}</Alert>
</div>
<script>
    var loginTpl = new Vue({
        el: '#loginTpl',
        data: {
            username:'',
            password:'',
            is_error:false,
            errMsg:''
        },
        methods: {
            init: function () {

            },
            doLogin:function () {
                var url='/account/doLogin';

            }
        },
        watch:{
            username:function () {
                
            },
            password:function () {
                
            }
        }
    });
    loginTpl.init();
</script>