$(document).ready(function(){
	$("#sign_in").click(function(){
		$.get("/sign_in",function(data,status){
			$("#main_content").html(data);
		});
		showAlert("我一点都不欢迎你登陆");
	});
	$("#order_manage").click(function(){
		$.get("/order_manage",function(data,status){
			$("#main_content").html(data);
		});
		showAlert("赶紧点饭吧");
	});
	$("#account_detail").click(function(){
		$.get("/account_detail",function(data,status){
			$("#main_content").html(data);
		})
		showAlert("年轻人，记得充值哦。 ————马化腾")
	})
	$("#order_own_list").click(function(){
		$.get("/order_own_list",function(data,status){
			$("#main_content").html(data);
		})
		showAlert("黑历史有什么好看的")
	})
	$("#fuli_list").click(function(){
		$.get("/fuli_list",function(data,status){
			$("#main_content").html(data);
		})
		showAlert("被你发现了。。。啊，不，我什么也不知道")
	})
	// $("#order_list").click(function(){
	// 	$.get("/order_list",function(data,status){
	// 		$("#right_list").html(data);
	// 	})
	// })
	/*Some Test*/
	$("#tts-submit").click(function(){
		var t = $("#tts-text").val()
		$("#tts-content").text(t)
		if (t != ""){
			if (t.length > 100)
				t = t.substring(0,99);
			var url ="http://api.voicerss.org/?key=fe1f1719cf7b47f69c3baa1e785a12b5&src="+t+"&hl=zh-tw&r=0"
			//var audio = document.getElementById("tts-audio")
			//audio.src = url
			$("#tts-audio")[0].src = url
			// $("#tts-audio")[0].play();
		};
	});
	/*主页直接刷新列表*/
	$.get("/order_list",function(data,status){
			$("#right_list").html(data);
	})
});


function showAlert(text){
	$('#alert').hide();
	$('#alert').fadeIn(1500);
	$("#alert_content").text(text);
	//$('#alert').fadeOut(2000);
}