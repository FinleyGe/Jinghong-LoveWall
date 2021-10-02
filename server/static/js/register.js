/*
 * @Author: F1nley
 * @Date: 2021-10-02 11:10:17
 * @LastEditors: F1nley
 * @LastEditTime: 2021-10-02 13:52:40
 * @Description: 
 */
$(function() {
    //注册表单验证
    $("#register-form").validate({
        submitHandler: function(form) {
            alert("Hello");
            $(form).ajaxSubmit({
                url: "/register",
                type: "post",
                dataType: "json",

                success: function(data, status) {
                    alert("data:" + data.message)
                },
                err: function(data, status) {
                    alert("err:" + data.message + ":" + status)
                }
            })
        }
    })
})