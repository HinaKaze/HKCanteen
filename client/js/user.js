$(document).ready(function(){
	$("#test-ajax").click(function(){
		$.get("http://sina.com.cn",function(data,status){
			$("#blog-content").text(data)
		});
	});
});