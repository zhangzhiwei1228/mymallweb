(function ($) {
    window.CONFIG = {

        USER_STATE_NORMAL: 1,
        USER_STATE_DISABLE: 0,
        getUserState: function () {
            var t = {};
            t[this.USER_STATE_NORMAL] = '正常';
            t[this.USER_STATE_DISABLE] = '禁用';
            return t;
        },


        PAY_TYPE_YF: 0,
        PAY_TYPE_HF: 1,
        getPayType: function () {
            var t = {};
            t[this.PAY_TYPE_YF] = '预付费';
            t[this.PAY_TYPE_HF] = '后付费';
            return t;
        },

        CHECK_STATUS_READY: 1,
        CHECK_STATUS_SUCCESS: 2,
        CHECK_STATUS_FAILURE: 4,
        CHECK_STATUS_ING: 3,
        CHECK_STATUS_ING6: 6,
        getCheckState: function () {
            var t = {};
            t[this.CHECK_STATUS_READY] = '待审核';
            t[this.CHECK_STATUS_ING6] = '审核中';
            t[this.CHECK_STATUS_SUCCESS] = '审批通过';
            t[this.CHECK_STATUS_FAILURE] = '审核未通过';
            t[this.CHECK_STATUS_ING] = '审核中';
            return t;
        },

        // 短信发送状态  2=同步下单成功，3=同步下单失败，4=订单成功，5=订单失败, 6=订单超时成功',
        SEND_STATUS_CREATE_SUCEESS: 2,
        SEND_STATUS_CREATE_FAIL: 3,
        SEND_STATUS_ORDER_SUCEESS: 4,
        SEND_STATUS_ORDER_FAIL: 5,
        SEND_STATUS_TIMEOUT_SUCCESS: 6,
        getSendState: function () {
            var t = {};
            t[this.SEND_STATUS_CREATE_SUCEESS] = '发送中';
            t[this.SEND_STATUS_CREATE_FAIL] = '失败';
            t[this.SEND_STATUS_ORDER_FAIL] = '失败';
            t[this.SEND_STATUS_ORDER_SUCEESS] = '成功';
            t[this.SEND_STATUS_TIMEOUT_SUCCESS] = '成功';
            return t;
        },

        // 订单状态 0为创建，1为已提交，4为充值成功，5为充值失败'
        PAY_STATUS_READY: 0,
        PAY_STATUS_SUBMIT: 1,
        PAY_STATUS_SUCCESS: 4,
        PAY_STATUS_FAILURE: 5,
        getOrderState: function () {
            var t = {};
            t[this.PAY_STATUS_READY] = '未支付';
            t[this.PAY_STATUS_SUBMIT] = '支付确认中...';
            t[this.PAY_STATUS_SUCCESS] = '支付成功';
            t[this.PAY_STATUS_FAILURE] = '支付失败';
            return t;
        },

        ////短信类型 1=行业短信，2=会员营销短信，3=彩信，4=闪信，5=视频短信，6=上行短信，7=语音验证码，8=精准营销标签文字短信，9=精准营销标签视频短信，10=精准营销场景文字短信，11=精准营销场景视频短信
        SMS_TYPE_PT: 1,
        SMS_TYPE_HYYX: 2,
        SMS_TYPE_MMS: 3,
        SMS_TYPE_SX: 4,
        SMS_TYPE_VEDIO: 5,
        SMS_TYPE_MO: 6,
        SMS_TYPE_VOICE: 7,
        SMS_TYPE_TAGSMS: 8,
        SMS_TYPE_TAGVADIO: 9,
        SMS_TYPE_TAGJZSMS: 10,
        SMS_TYPE_TAGJZVADIO: 11,
        SMS_TYPE_TAGJZYXSMS: 12,
        SMS_TYPE_TAGJZYXVADIO: 13,
        getSmsType: function () {
            var t = {};
            t[this.SMS_TYPE_PT] = '行业短信';
            t[this.SMS_TYPE_HYYX] = '营销短信';
            t[this.SMS_TYPE_MMS] = '彩信';
            t[this.SMS_TYPE_SX] = '闪信';
            t[this.SMS_TYPE_VEDIO] = '视频短信';
            t[this.SMS_TYPE_MO] = '上行短信';
            t[this.SMS_TYPE_VOICE] = '语音验证码';
            t[this.SMS_TYPE_TAGSMS] = '精准营销标签文字短信';
            t[this.SMS_TYPE_TAGVADIO] = '精准营销标签视频短信';
            t[this.SMS_TYPE_TAGJZSMS] = '精准营销场景文字短信';
            t[this.SMS_TYPE_TAGJZVADIO] = '精准营销场景视频短信';
            t[this.SMS_TYPE_TAGJZYXSMS] = '精准营销文字短信';
            t[this.SMS_TYPE_TAGJZYXVADIO] = '精准营销视频短信';
            return t;
        },
        SMS_LEFT_PT: 1,
        SMS_LEFT_HYYX: 2,
        SMS_LEFT_MMS: 3,
        SMS_LEFT_SX: 4,
        SMS_LEFT_VEDIO: 5,
        SMS_LEFT_MO: 6,
        SMS_LEFT_VOICE: 7,
        SMS_LEFT_TAGSMS: 8,
        SMS_LEFT_TAGVADIO: 9,
        SMS_LEFT_TAGJZSMS: 10,
        SMS_LEFT_TAGJZVADIO: 11,
        getSmsLeftPre: function () {
            var t = {};
            t[this.SMS_LEFT_PT] = '行业短信剩余';
            t[this.SMS_LEFT_HYYX] = '营销短信剩余';
            t[this.SMS_LEFT_MMS] = '彩信剩余';
            t[this.SMS_LEFT_SX] = '闪信剩余';
            t[this.SMS_LEFT_VEDIO] = '视频短信剩余';
            t[this.SMS_LEFT_MO] = '上行短信剩余';
            t[this.SMS_LEFT_VOICE] = '语音验证码剩余';
            t[this.SMS_LEFT_TAGSMS] = '精准营销标签文字短信剩余';
            t[this.SMS_LEFT_TAGVADIO] = '精准营销标签视频短信剩余';
            t[this.SMS_LEFT_TAGJZSMS] = '精准营销场景文字短信剩余';
            t[this.SMS_LEFT_TAGJZVADIO] = '精准营销场景视频短信剩余';
            return t;
        },
        getSmsLeft: function () {
            var t = {};
            t[this.SMS_LEFT_PT] = '已发送行业短信';
            t[this.SMS_LEFT_HYYX] = '已发送营销短信';
            t[this.SMS_LEFT_MMS] = '已发送彩信';
            t[this.SMS_LEFT_SX] = '已发送闪信';
            t[this.SMS_LEFT_VEDIO] = '已发送视频短信';
            t[this.SMS_LEFT_MO] = '已发送上行短信';
            t[this.SMS_LEFT_VOICE] = '已发送语音验证码';
            t[this.SMS_LEFT_TAGSMS] = '已发送精准营销标签文字短信';
            t[this.SMS_LEFT_TAGVADIO] = '已发送精准营销标签视频短信';
            t[this.SMS_LEFT_TAGJZSMS] = '已发送精准营销场景文字短信';
            t[this.SMS_LEFT_TAGJZVADIO] = '已发送精准营销场景视频短信';
            return t;
        },
        PROFILE_TYPE_ZD: 0,
        PROFILE_TYPE_CJ: 1,
        getProfileType: function () {
            var t = {};
            t[this.PROFILE_TYPE_ZD] = '自助式标签画像';
            t[this.PROFILE_TYPE_CJ] = '场景画像';
            return t;
        },
        SEND_STATUS_INI: 0, //待发送
        SEND_STATUS_ING: 1, //发送中
        SEND_STATUS_FINISH: 2, //已发送
        SEND_STATUS_ERROR:3,
        getSendStatuss: function () {
            var t = {};
            t[this.SEND_STATUS_INI] = '待发送';
            t[this.SEND_STATUS_ING] = '发送中';
            t[this.SEND_STATUS_FINISH] = '已发送';
            t[this.SEND_STATUS_ERROR] = '发送错误';
            return t;
        },

    };

    Vue.filter('parValue', function (val, prodType) {
        if (!val) {
            return val;
        }
        if (prodType == CONFIG.PROD_TYPE_HF) {
            return sms.f2y(val) + '元';
        } else {
            if (val / 1024 >= 1 && val % 1024 == 0) {
                return val / 1024 + 'G';
            } else {
                return val + 'M';
            }
        }
    });
    Vue.filter('f2y', function (fen) {
        return sms.f2y(fen);
    });

    Vue.filter('userState', function (val) {
        return CONFIG.getUserState()[val] || '';
    });

    Vue.filter('payType', function (val) {
        return CONFIG.getPayType()[val] || '';
    });
    Vue.filter('profileType', function (val) {
        return CONFIG.getProfileType()[val] || '';
    });
    Vue.filter('sendStatus', function (val) {
        return CONFIG.getSendStatuss()[val] || '';
    });

    Vue.filter('checkState', function (val) {
        return CONFIG.getCheckState()[val] || '';
    });

    Vue.filter('orderStatus', function (val) {
        return CONFIG.getOrderState()[val] || '';
    });
    Vue.filter('smsSendStatus', function (val) {
        return CONFIG.getSendState()[val] || '';
    });
    Vue.filter('smsType', function (val) {
        return CONFIG.getSmsType()[val] || '';
    });
    Vue.filter('smsLeft', function (val) {
        return CONFIG.getSmsLeft()[val] || '';
    });
    Vue.filter('smsLeftPre', function (val) {
        return CONFIG.getSmsLeftPre()[val] || '';
    });

    window.sms = {
        showConfirm:function(msg, cb){
            var $tmpl = $('<div class="alert alert-info alert-confirm" style="top:30%; background: #fff;animation: show ease 1s,moveTop1 ease 1s; -webkit-animation: show ease 1s,moveTop1 ease 1s;" role="alert"><span style="margin-right: 10px" class="glyphicon"></span>' + msg + '' +
                '<br/><br/><button class="btn btn_1">取消</button><button style="margin-left:20px;" class="btn btn-info btn-info_1">确定</button></div>');
            if ($('body').find('.alert-confirm').length == 0) {
                $('body').append($tmpl);
                $(".btn_1").click(function(){
                    $tmpl.remove();
                })
                $(".btn-info_1").click(function(){
                    $tmpl.remove();
                    cb && cb();
                })
            }

        },
        showTip: function (type, msg, stayTime, cb, icon) {
            if (stayTime !== 0 && !stayTime) {
                stayTime = msg.length * 300;
                stayTime = stayTime < 5000 ? 5000 : stayTime;
            }
            var $tmpl = $('<div class="alert alert-' + type + '" role="alert"><span style="margin-right: 10px" class="glyphicon ' + icon + '"></span>' + msg + '</div>');
            $('body').append($tmpl);

            setTimeout(function () {
                $tmpl.remove();
                cb && cb();
            }, stayTime);
        },
        showTipErr: function (msg, stayTime) {
            sms.showTip('danger', msg, stayTime, '', 'glyphicon-remove');
        },
        showTipSuc: function (msg, stayTime) {
            sms.showTip('success', msg, stayTime, '', 'glyphicon-ok');
        },
        showTipInfo: function (msg, stayTime) {
            sms.showTip('info', msg, stayTime, '', 'glyphicon-info-sign');
        },
        showLoading: function () {
            _globalLoading = $('#loading');
            if (_globalLoading.length <= 0) {
                _globalLoading = $('<div id="loading"><div class="spinner"><div class="spinner-container container1">' +
                    '<div class="circle1"></div><div class="circle2"></div> <div class="circle3"></div> <div class="circle4"></div></div> ' +
                    '<div class="spinner-container container2"><div class="circle1"></div> <div class="circle2"></div> <div class="circle3"></div> <div class="circle4"></div></div> ' +
                    '<div class="spinner-container container3"> <div class="circle1"></div> <div class="circle2"></div> <div class="circle3"></div> <div class="circle4"></div> </div> ' +
                    '</div></div>');
                $('body').append(_globalLoading);
            }
        },
        hideLoading: function () {
            $('#loading').remove();
        },
        download: function (url, params) {
            var sp = url.indexOf('?') != -1 ? '&' : '?';
            url = url + sp + $.param(params);
            $("body").append("<iframe src='" + url + "' style='display:none;'></iframe>");
            sms.showTipInfo('请等待自动提示下载');
        },
        f2y: function (fen) {
            var yuan = '';
            fen = '' + fen;
            switch (fen.length) {
                case 1:
                    yuan = '0.0' + fen;
                    break;
                case 2:
                    yuan = '0.' + fen;
                    break;
                default:
                    if (fen.indexOf(".") > 0) {
                        yuan = (fen / 100).toFixed(2);
                        break;
                    } else {
                        var l = fen.length;
                        yuan = fen.substr(0, l - 2) + "." + fen.substr(-2);
                        break;
                    }

            }
            return yuan;
        },
        fajax: function (url, p, settings) {
            var defaults = {
                type: settings.type,
                error: function (code, msg) {
                    sms.showTipErr(msg);
                },
                success: function () {
                },
                complete: function (jqXHR, textStatus) {
                }
            };
            settings = $.extend(defaults, settings);
            sms.showLoading();
            var xsrf, xsrflist;

            try {
                xsrf = $.cookie("_xsrf");
                xsrflist = xsrf.split("|");
                p._xsrf = $.base64.decode(xsrflist[0]);
            } catch(err) {
                window.location.reload()
                return false;
            }
            /*finally {

            }*/

            $.ajax({
                cache: false,
                url: url,
                type: settings.type,
                contentType: 'application/json',
                data: JSON.stringify(p),
                success: function (result) {
                    sms.hideLoading();
                    if (result.code == 0) {
                        settings.success(result.data, result);
                    } else {
                        settings.error(result.code, result.msg, result);
                    }
                },
                complete: function (jqXHR, textStatus) {
                    settings.complete(jqXHR, textStatus);
                    sms.hideLoading();
                },
                error: function (jqXHR, textStatus, errorThrown) {
                    settings.error(-1, textStatus);
                    sms.hideLoading();
                }
            });
        },
        fget: function (url, p, success, error) {
            sms.fajax(url, p, {
                type: 'get',
                success: success,
                error: error
            });
        },
        fpost: function (url, p, success, error) {
            sms.fajax(url, p, {
                type: 'post',
                success: success,
                error: error
            });
        },
        getQuery: function (name, url) {
            var _s = location.search;
            if (url) {
                var aa = document.createElement('a');
                aa.href = url;
                _s = aa.search;
            }
            var result = _s.match(new RegExp("[\?\&][^\?\&]+=[^\?\&]+", "g"))

            if (result == null) {
                return '';
            }
            for (var i = 0; i < result.length; i++) {
                var keyValue = result[i].substring(1).split("=");
                result[keyValue[0]] = keyValue[1];
            }
            return name ? result[name] : result;
        },
        setCookie: function (cname, cvalue, exday) {
            var d = new Date();
            exday = exday || 1;
            d.setTime(d.getTime() + (exday * 24 * 60 * 60 * 1000));
            var expires = "expires=" + d.toUTCString();
            document.cookie = cname + "=" + cvalue + "; path=/;" + expires;
        },

        getCookie: function (cname) {
            var name = cname + "=";
            var ca = document.cookie.split(';');
            for (var i = 0; i < ca.length; i++) {
                var c = ca[i];
                while (c.charAt(0) == ' ') c = c.substring(1);
                if (c.indexOf(name) != -1) return decodeURIComponent(c.substring(name.length, c.length));
            }
            return null;
        }
    };
})(jQuery);

(function ($) {
    $.extend($.fn, {
        datepicker: function (options, callback) {
            if (typeof options == 'function') {
                callback = options;
                options = {};
            }
            options = $.extend({}, {
                autoApply: true,
                singleDatePicker: false,
                autoUpdateInput: true,
                maxDate: new Date(),
                //minDate:moment('2016-11-01'),
                locale: {
                    "format": "YYYY-MM-DD",
                    "applyLabel": "应用",
                    "cancelLabel": "取消",
                    "customRangeLabel": "Custom",
                    "daysOfWeek": ["日", "一", "二", "三", "四", "五", "六"],
                    "monthNames": ["一月", "二月", "三月", "四月", "五月", "六月", "七月", "八月", "九月", "十月", "十一月", "十二月"]
                }
            }, options || {});
            return $(this).daterangepicker(options, callback);
        },
        fajax: function (type, url, p, success, error) {
            var $this = $(this);
            var nv = $this.data('ltext');
            var fillMethod = $this[0].tagName.toLowerCase() == 'button' ? 'text' : 'val';
            var ov = $this[fillMethod]();

            if ($this.data('loading')) {
                return;
            }
            if (!error) {
                error = function (code, msg, res) {
                    sms.showTipErr(msg);
                }
            }

            $this[fillMethod](nv).data('loading', 1);
            sms.showLoading();
            sms[type](url, p, function (data, res) {
                sms.hideLoading();
                $this[fillMethod](ov);
                $this.removeData('loading');
                success && success(data, res);
            }, function (code, msg, res) {
                sms.hideLoading();
                $this[fillMethod](ov);
                $this.removeData('loading');
                error && error(code, msg, res);
            });
        },
        fget: function (url, p, success, error) {
            $(this).fajax('fget', url, p, success, error);
        },
        fpost: function (url, p, success, error) {
            $(this).fajax('fpost', url, p, success, error);
        }
    });
    window.SimplePager = function (curPage, recordCount, pageSize, onchange) {

        var THIS = this;
        pageSize = pageSize || 15;
        var totalPage = Math.ceil(recordCount / pageSize);

        var pageStr = parsePage(curPage, totalPage);

        var $dom = $(
            '<div class="pagination_pager"><div class="pagination">' + pageStr + '</div>' +
            '<span class="totaoPage" style="padding-left: 10px">共<i class="showTotaoPage">' + totalPage + '</i></span><span>页</span>' +
            '<span style="padding: 0px 3px;">/<i class="showTotaoNum">' + recordCount + '</i></span><span>条数据</span>' +
            '<div class="setNum"><span style="margin-left: 10px;">每页&nbsp;</span><select class="form-control input-sm" style="width: 50px;height: 24px;">' +
            '<option selected="selected" value="'+pageSize+'">'+pageSize+'</option><option value="50">50</option><option' +
            ' value="100">100</option></select>&nbsp;条记录 ' +
            '<span style="margin-left: 10px;">转到</span> <input class="form-control input-sm"  type="number" name="pageNum" style="width: 50px; height: 24px;padding: 0;text-align: center"> ' +
            '<span class="pageBtn btn-primary btn-sm" style="cursor: pointer">确定</span> </div> </div>');

        var $totalPage = $dom.find('.showTotaoPage');
        var $recordCount = $dom.find('.showTotaoNum');
        var $inputPage = $dom.find('input[name=pageNum]');
        var $pagination = $dom.find('.pagination');

        $dom.find('select').val(pageSize);

        $dom.find('select').change(function () {
            pageSize = $(this).val() * 1;
            THIS.setRecordCount(recordCount, 1);
            THIS.goto(curPage);
        });
        $dom.find('.pageBtn').click(function () {
            THIS.goto($inputPage.val() * 1)
        });
        $dom.on('click', '.turnLeft', function () {
            THIS.prev();
        });
        $dom.on('click', '.turnRight', function () {
            THIS.next();
        });
        $dom.on('click', 'a.item', function () {
            var tp = $(this).text() * 1;
            !isNaN(tp) && THIS.goto(tp);
        });

        function parsePage(curPage, totalPage) {
            var prev = curPage > 1 ? "<li><a class='turnLeft'>&laquo;</a></li>" : "";
            var next = curPage < totalPage ? "<li><a class='turnRight'>&raquo;</a></li>" : "";
            var dot = "<li class='disabled item'><a>...</a></li>";
            var m, str = "", begin, end;

            if (totalPage <= 9) {
                for (m = 1; m <= totalPage; m += 1) {
                    if (m === curPage) {
                        str += "<li class='active'><a class='item'>" + m + "</a></li>";
                    } else {
                        str += "<li><a class='item'>" + m + "</a></li>";
                    }
                }
            } else {
                if (curPage <= 5) {
                    for (m = 1; m <= 7; m += 1) {
                        if (curPage === m) {
                            str += "<li class='active'><a class='item'>" + m + "</a></li>";
                        } else {
                            str += "<li><a class='item'>" + m + "</a></li>";
                        }
                    }
                    str += dot;
                } else {
                    begin = curPage - 2;
                    end = curPage + 2;
                    str += "<li><a class='item'>1</a><a class='item'>2</a></li>";
                    str += dot;
                    if (end > totalPage) {
                        end = totalPage;
                        begin = end - 4;
                        if (begin + 2 < curPage) {
                            begin += 1;
                        }
                    } else if (end + 1 === totalPage) {
                        end = totalPage;
                    }
                    for (m = begin; m <= end; m += 1) {
                        if (curPage === m) {
                            str += "<li class='active'><a class='item'>" + m + "</a></li>";
                        } else {
                            str += "<li><a class='item'>" + m + "</a></li>";
                        }
                    }
                    if (end !== totalPage) {
                        str += dot;
                    }
                }
            }

            return prev + str + next;
        }

        this.renderTo = function ($obj) {
            $obj.html($dom);
            $dom[totalPage < 1 ? 'hide' : 'show']();
            return this;
        };

        this.setRecordCount = function (record, _page) {
            totalPage = Math.ceil(record / pageSize);
            if (_page) {
                curPage = _page * 1;
            }

            recordCount = record;
            $totalPage.html(totalPage);
            $recordCount.html(record);
            $pagination.html(parsePage(curPage, totalPage));
            $dom[totalPage < 1 ? 'hide' : 'show']();
        };

        this.getPageSize = function () {
            return pageSize;
        };
        this.getTotalPage = function () {
            return totalPage;
        };

        this.getCurPage = function () {
            return curPage;
        };

        this.goto = function (_page) {
            _page = isNaN(_page) ? curPage : _page;
            if (_page > totalPage) {
                _page = totalPage;
            } else if (_page < 1) {
                _page = 1;
            }
            /*else if(_page==curPage){
             return this;
             }*/
            curPage = _page;
            $pagination.html(parsePage(curPage, totalPage));
            onchange && onchange(curPage, pageSize);
            return this;
        };
        this.prev = function () {
            this.goto(curPage - 1);
        };
        this.next = function () {
            this.goto(curPage + 1);
        };

        return this;
    };
    window.cmodal = function () {
        var html = '<div id="[Id]" class="modal fade" role="dialog" aria-labelledby="modalLabel">' +
            '<div class="modal-dialog modal-sm">' +
            '<div class="modal-content">' +
            '<div class="modal-header">' +
            '<button type="button" class="close" data-dismiss="modal"></button>' +
            '<h4 class="modal-title" id="modalLabel">[Title]</h4>' +
            '</div>' +
            '<div class="modal-body">' +
            '<p>[Message]</p>' +
            '</div>' +
            '<div class="modal-footer">' +
            '<button type="button" class="btn btn-default cancel" data-dismiss="modal">[BtnCancel]</button>' +
            '<button type="button" class="btn btn-primary ok" data-dismiss="modal">[BtnOk]</button>' +
            '</div>' +
            '</div>' +
            '</div>' +
            '</div>';


        var dialogdHtml = '<div id="[Id]" class="modal fade" role="dialog" aria-labelledby="modalLabel">' +
            '<div class="modal-dialog">' +
            '<div class="modal-content">' +
            '<div class="modal-header">' +
            '<button type="button" class="close" data-dismiss="modal"></button>' +
            '<h4 class="modal-title" id="modalLabel">[Title]</h4>' +
            '</div>' +
            '<div class="modal-body">' +
            '</div>' +
            '</div>' +
            '</div>' +
            '</div>';
        var reg = new RegExp("\\[([^\\[\\]]*?)\\]", 'igm');
        var generateId = function () {
            var date = new Date();
            return 'mdl' + date.valueOf();
        };
        var init = function (options) {
            options = typeof options == 'string' ? {messate: options} : options;
            options = $.extend({}, {
                title: "提示",
                message: "提示内容",
                btnok: "确定",
                btncl: "取消",
                width: 200,
                auto: false
            }, options || {});
            var modalId = generateId();
            var content = html.replace(reg, function (node, key) {
                return {
                    Id: modalId,
                    Title: options.title,
                    Message: options.message,
                    BtnOk: options.btnok,
                    BtnCancel: options.btncl
                }[key];
            });
            $('body').append(content);
            $('#' + modalId).modal({
                width: options.width,
                backdrop: 'static'
            });
            $('#' + modalId).on('hide.bs.modal', function (e) {
                $('body').find('#' + modalId).remove();
            });
            return modalId;
        };

        return {
            alert: function (options) {
                if (typeof options == 'string') {
                    options = {
                        message: options
                    };
                }
                var id = init(options);
                var modal = $('#' + id);
                modal.find('.ok').removeClass('btn-success').addClass('btn-primary');
                modal.find('.cancel').hide();

                return {
                    id: id,
                    on: function (callback) {
                        if (callback && callback instanceof Function) {
                            modal.find('.ok').click(function () {
                                callback(true);
                            });
                        }
                    },
                    hide: function (callback) {
                        if (callback && callback instanceof Function) {
                            modal.on('hide.bs.modal', function (e) {
                                callback(e);
                            });
                        }
                    }
                };
            },
            confirm: function (options) {
                if (typeof options == 'string') {
                    options = {
                        message: options
                    };
                }
                var id = init(options);
                var modal = $('#' + id);
                //modal.find('.ok').removeClass('btn-primary').addClass('btn-success');
                modal.find('.cancel').show();
                return {
                    id: id,
                    on: function (callback) {
                        if (callback && callback instanceof Function) {
                            modal.find('.ok').click(function () {
                                callback(true);
                            });
                            modal.find('.cancel').click(function () {
                                callback(false);
                            });
                        }
                    },
                    hide: function (callback) {
                        if (callback && callback instanceof Function) {
                            modal.on('hide.bs.modal', function (e) {
                                callback(e);
                            });
                        }
                    }
                };
            },
            dialog: function (options) {
                options = $.extend({}, {
                    title: 'title',
                    url: '',
                    content: '',
                    width: 800,
                    height: 550,
                    onReady: function () {
                    },
                    onShown: function (e) {
                    }
                }, options || {});
                var modalId = generateId();

                var content = dialogdHtml.replace(reg, function (node, key) {
                    return {
                        Id: modalId,
                        Title: options.title
                    }[key];
                });
                $('body').append(content);
                var target = $('#' + modalId);
                if (options.url) {
                    target.find('.modal-body').load(options.url);
                } else {
                    target.find('.modal-body').html(options.content);
                }
                if (options.onReady())
                    options.onReady.call(target);
                target.modal();
                target.on('shown.bs.modal', function (e) {
                    if (options.onReady(e))
                        options.onReady.call(target, e);
                });
                target.on('hide.bs.modal', function (e) {
                    $('body').find(target).remove();
                });
            }
        }
    }();
})(jQuery);

