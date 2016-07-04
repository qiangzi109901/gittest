
define(['jquery'], function($){

    function Infinite(opts){
        var rules = $.extend({
            container : "#page_list",
            pagination : "#pagination",
            item : ".page_item",
            autoTimes : 2,
            nCurrentPage : 1, //当前页数
            nPageSize : 15, //每页条数
            loadMsg : "正在加载下一页数据...",
            endMsg : "完",
            fnLoadData : null,
            cb : null
        },opts || {});
        this.rules = rules;
        this.container = rules.container;
        this.pagination = rules.pagination;
        this.item = rules.item;
        this.loadMsg = rules.loadMsg;
        this.endMsg = rules.endMsg;
        this.fnLoadData = rules.fnLoadData;
        this.cb = rules.cb;
        this.isComplete = false;
        this.htmlData = "";  //html格式的列表数据 数组格式
        this.tempTimes = 0;
        this.bFlag = true;
        this.autoTimes = rules.autoTimes;
        this.pageSize = rules.nPageSize;
        this.currentPage = rules.nCurrentPage;

        this.init();
    }

    Infinite.prototype = {

        //初始化分页组件
        init : function(){
            var __this = this;
            while(__this.needLoad()){
                this.tempTimes ++;
                this.load();
            }

            $(this.pagination).click(function(){
                __this.load();
                __this.tempTimes = 1;
            });
            if(this.autoTimes > 0 && !this.isComplete){
                //当前页小于总页数
                $(window).scroll(function(){
                    if(__this.needLoad()){
                        if(__this.bFlag && __this.tempTimes<__this.autoTimes && !__this.isComplete){
                            __this.bFlag = false; //正在获取服务器数据，将下拉加载数据功能关闭
                            __this.tempTimes++;
                            __this.load();
                        }
                    }
                });
            }
        },
        needLoad : function(){
            var atop = $(this.pagination).offset().top;
            var stop = $(window).scrollTop();
            var hei = $(this.container).height();
            if(hei<window.innerHeight && this.tempTimes < this.autoTimes){
                return true;
            }
            if(stop + window.innerHeight > atop - 200 && this.tempTimes < this.autoTimes){
                return true;
            }
            return false;
        },
        load : function(){
            this.htmlData = this.fnLoadData(this.currentPage,this);    //同步方法
            var cstr = '"'+this.item.slice(1)+'"';   //必须要保证item被双引号包围，而不是单引号
            var times = this.getTimes(this.htmlData,cstr); //获取记录数
            if(times<this.pageSize){
                this.isComplete = true;
                this.setEnd();
            }else{
                $(this.pagination).css("visibility","hidden");
            }
            $(this.container).append(this.htmlData);
            this.currentPage ++;
            this.bFlag = true;

            if(this.tempTimes >= this.autoTimes){
                $(this.pagination).css("visibility","visible");
            }
            this.cb && this.cb();
        },
        getTimes : function(str,cstr){
            var count = 0;
            var index = 0;
            while((index = str.indexOf(cstr,index))!= -1){
                index ++;
                count ++;
            }
            return count;
        },
        setEnd : function(){
            $(this.pagination).css("visibility","visible").text(this.endMsg).unbind("click");
        }
    }


    return Infinite;

});





