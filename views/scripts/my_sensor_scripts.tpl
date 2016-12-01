<script>
	$(document).ready(function() {
		$('.dropdown-menu li a').hover(
		function() {
			$(this).children('i').addClass('icon-white');
		},
		function() {
			$(this).children('i').removeClass('icon-white');
		});

		if($(window).width() > 760)
		{
			$('tr.list-users td div ul').addClass('pull-right');
		}
	});

 $(function () {
    $("#pagination0").bootstrapPaginator({
      currentPage: '{{.Page.PageNo}}',
      totalPages: '{{.Page.TotalPage}}',
      bootstrapMajorVersion: 3,
      size: "small",
      onPageClicked: function(e,originalEvent,type,page){
        window.location.href = "/mysensor/" + page
      }
    });
  });

 // $.jqPaginator('#pagination0', {
 //        totalPages:  '{{.Page.TotalPage}}',
 //        visiblePages: 5,
 //        currentPage: '{{.Page.PageNo}}',
 //        prev: '<li class="prev"><a href="javascript:;">上一页</a></li>',
 //        next: '<li class="next"><a href="javascript:;">下一页</a></li>',
 //        page: '<li class="page"><a href="javascript:;">{{.Page.PageNo}}</a></li>',
 //        onPageChange: function (page, type) {
 //            // alert(type + '：' + page);
 //            // if(if_firstime){
 //            //     if_firstime = false;
 //            // }else if(!if_firstime){
 //            //     changePage(page);
 //            // }
 //            window.location.href = "/mydevice/" + page
 //        }
 //    });
</script>