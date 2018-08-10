<style>
    .reg-input-a{
        border: none !important;
    }
    [v-cloak] {
        display: none !important;
    }
</style>
<div class="bg" id="rgTpl" v-cloak>
    <div class="reg">
        <h2 class="h2">会员注册</h2>
        <div v-show="errMsg !== ''">
            <template>
                <Alert type="error" show-icon>{{errMsg}}</Alert>
            </template>
        </div>
        <div class="fl">
            <table>
                <tr>
                    <td><span class="tit">手机号码：</span></td>
                    <td>
                        <i-input placeholder="请输入手机号" v-model="mobile" size="large" class="reg-input-a" type="number" style="border: none"></i-input>
                    </td>
                </tr>
                <tr>
                    <td><span class="tit">图形验证码：</span></td>
                    <td>
                        <i-input placeholder="请输入验证码" v-model="captcha" size="large" class="reg-input-b" type="number" style="border: none"></i-input>
                        <img :src="captchaInfo.img" style="margin-left: 23px;height: 36px;cursor:pointer" @click="getCaptcha()">
                    </td>
                </tr>
                <tr>
                    <td><span class="tit">验证码：</span></td>
                    <td>
                        <i-input placeholder="请输入验证码" v-model="verfiCode" size="large" class="reg-input-b" type="number" style="border: none"></i-input>
                        <i-button type="primary" style="margin-left: 23px;height: 36px;">发送验证码</i-button>
                    </td>
                </tr>
                <tr>
                    <td><span class="tit">密码：</span></td>
                    <td>
                        <i-input placeholder="请输入6-20位包含数字和字母的密码" v-model="password" size="large" class="reg-input-a" type="password" style="border: none"></i-input>
                    </td>
                </tr>
                <tr>
                    <td><span class="tit">确认密码：</span></td>
                    <td>
                        <i-input placeholder="请再次输入密码" v-model="repassword" size="large" class="reg-input-a" type="password" style="border: none"></i-input>
                    </td>
                </tr>
                <tr>
                    <td><span class="tit">邀请码：</span></td>
                    <td>
                        <i-input placeholder="请输入邀请人号码（填写获得20会员积分）" v-model="invitationCode" size="large" class="reg-input-a" type="password" style="border: none"></i-input>
                    </td>
                </tr>
                <tr>
                    <td></td>
                    <td>
                        <i-button type="success" long @click="doRegister()" style="font-size: 18px;">注册</i-button>
                    </td>
                </tr>
            </table>
        </div>
        <div class="fr">
            <p><font style="font-size:14px;">邀请码说明</font>(<font style="color:#b30000;">分2种情况</font>)：</p>
            <p>A；无邀请码注册，会员激活后获<font style="color:#b30000;">10积分。</font></p>
            <p>B；有邀请码注册，会员激活后获<font style="color:#b30000;">20积分</font>，同时邀请码推荐人也获20积分（邀请码就是你推荐人的手机号，你注册成功后用你的邀请码推荐朋友，注册成功激活后，你就可以得到<font style="color:#b30000;">20积分</font>，你朋友也得<font style="color:#b30000;">20积分</font>，以此类推）。</p>
        </div>
    </div>

</div>
<script>
    var rgTpl = new Vue({
        el: '#rgTpl',
        data: {
            username:'',
            mobile:'',
            verfiCode:'',//手机验证码
            password:'',
            repassword:'',
            captcha:'',
            captchaInfo:{},
            invitationCode:'',//邀请码
            is_error:false,
            errMsg:'',
            is_captcha:false,
        },
        methods: {
            init: function () {
                this.getCaptcha();
            },
            getCaptcha:function () {
                sms.fpost("/captcha/createCaptcha", {}, function (data) {
                    rgTpl.captchaInfo=data
                }, function (code, msg) {
                    this.showError("获取图形验证码错误，请刷新")
                    return false
                });
            },
            doRegister:function () {
                var url='/account/doRegister';
                var p = {};
                p.username = this.username;
                p.password = this.password;
                sms.fpost(url, p, function (data) {
                    window.location.href='/account/SuccessTpl'
                }, function (code, msg) {
                    this.showError(msg);
                    return false
                });
            },
            showError:function (msg,timeout) {
                rgTpl.is_error = true;
                rgTpl.errMsg = msg;
                var time = timeout ? timeout : 1500;
                setTimeout(function () {
                    rgTpl.is_error = false;
                    rgTpl.errMsg = '';
                }, time);
            }
        },
        watch:{
            mobile:function (val) {
                if(!(/^1[34578]\d{9}$/.test(val))){
                    //this.showError("手机号码格式有误，请重填")
                    return false
                }
            },
            password:function (val) {
                if (!/^(?![0-9]+$)(?![a-zA-Z]+$)[0-9A-Za-z]{6,20}$/.test(val)) {
                    //this.showError("请输入6-20位包含数字和字母的密码")
                    return false;
                }
            },
            repassword:function (val) {
                if(val !== this.password) {
                    //this.showError("两次密码不一致，请重新输入")
                    return false;
                }
            },
            captcha:function (val) {
                var url='/captcha/verfiyCaptcha';
                var p = {};
                p.val = val;
                p.idkey = this.captchaInfo.captchaId
                sms.fpost(url, p, function (data) {
                    rgTpl.is_captcha = data
                }, function (code, msg) {
                    this.showError("请刷新重试")
                    return false
                });
            },
            verfiCode:function () {

            },
        }
    });
    rgTpl.init();
</script>