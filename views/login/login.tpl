<div style="background:url('/static/img/img-43.jpg' ) no-repeat center" class="log-in" id="loginTpl" v-cloak>
    <div class="w1190">
        <div class="log-in-box">
            <div class="log-in-boxn">
                <h2>会员登录</h2>
                <form>
                    <i-input placeholder="请输入用户名" style="background:url('/static/img/img-21.png' )  no-repeat 10px center;border: none;" v-model="username" size="large" class="input-a" type="text"></i-input>
                    <i-input placeholder="请输入密码" style="background:url('/static/img/img-22.png') no-repeat 12px center;border: none;" v-model="password"  size="large" class="input-a" type="password"></i-input>
                    <Alert type="error" v-show="is_error">{{errMsg}}</Alert>
                    <i-button type="success" long @click="doLogin()" style="font-size: 18px;">登录</i-button>
                    <div class="two-a">
                        <a href="/account/registerTpl">免费注册</a>
                        <a href="">忘记密码</a>
                    </div>
                </form>
            </div>
        </div>
    </div>
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
                var p = {};
                p.username = this.username;
                p.password = this.password;
                sms.fpost(url, p, function (data) {
                    console.log(data)
                }, function (code, msg) {
                    loginTpl.is_error = true;
                    loginTpl.errMsg = msg;
                    setTimeout(function () {
                        loginTpl.is_error = false;
                        loginTpl.errMsg = '';
                    }, 1000);
                });
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