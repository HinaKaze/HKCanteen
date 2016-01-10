$(document).ready(function(){
	$("#sign_in").click(function(){
		$.get("/sign_in",function(data,status){
			$("#main_content").html(data);
		});
		showAlert("我一点都不欢迎你登陆");
	});
	$("#order_meal").click(function(){
		$.get("/order_meal",function(data,status){
			$("#main_content").html(data);
		});
		showAlert("赶紧点饭吧");
	});
});


function showAlert(text){
	$('#alert').hide();
	$('#alert').fadeIn(1500);
	$("#alert_content").text(text);
	//$('#alert').fadeOut(2000);
}