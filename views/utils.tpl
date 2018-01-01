{{define "head"}}
<meta charset="utf8" />
<title>{{ .Title }}</title>
<meta name="MobileOptimized" content="240"/>
<meta name="applicable-device" content="mobile"/>
<meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0" />
<meta http-equiv="Content-Type" content="application/vnd.wap.xhtml+xml;charset=gbk"/>
<link rel="shortcut icon" href="/favicon.ico" />
<link href="/file/css/style.css" rel="stylesheet">
<link href="/file/css/layer.css" rel="stylesheet">
<link rel="stylesheet" href="/file/css/jquery-ui.css">
<script src="/file/js/jquery-1.9.1.js"></script>
<script src="/file/js/jquery-ui.js"></script>
<script src="/file/js/style.js" type="text/javascript"></script>
<script src="/file/js/layer.js?" type="text/javascript"></script>
{{end}}

{{define "line"}}
&nbsp;&nbsp;&nbsp;&nbsp;{{ str2html .}}<br><br>
{{end}}

{{define "topbar1"}}
<div class="topbar">
<span class="logo"><a href="/">小说阅读网</a></span><span class="user"><a href="/login">登录</a><a href="/register">注册</a></span>
</div>
{{end}}

{{define "topbar2"}}
<div class="topbar">
<span class="c1"><a href="javascript:history.go(-1)" >返回</a></span><span class="c3">{{ .BookName }}</span><span class="c4"><a href="/" >首页</a></span>
</div>
{{end}}

{{define "subnav"}}
<ul><li><a href="/">首页</a></li><li><a href="/classify">分类</a></li><li><a href="/rank">排行</a></li><li><a href="/shujia">书架</a></li></ul>
{{end}}

{{define "subnavbar"}}
<div class="subnav">
{{template "subnav" .}}
</div>
{{end}}

{{define "footerbar"}}
<footer id="footer">
{{template "subnav" .}}
</footer>
{{end}}

{{define "chaptertitle"}}
<h1 id="BookTitle">{{ .ChapterName }}</h1>
{{end}}

{{define "link"}}
<a href={{ .ChapterUrlBefore }}>上一章</a><a href={{ .CatalogUrl }}>回目录</a><a href={{ .ChapterUrlAfter }}>下一章</a>
{{end}}

{{define "content"}}
<div id="content">
{{range .Content}}
{{template "line" .}}
{{end}}
</div>
{{end}}

{{define "chapter"}}
<body id="BookBody">
<div id="main">
  <section id="container">
  <div class="bookset clrfix"><script>bookset();</script></div>
    <article>
      {{template "chaptertitle" .}}
      <p class="link bx">{{template "link" .}}</p>
      {{template "content" .}}
      <p class="link">{{template "link" .}}</p>
    </article>
  </section>
</div>
<script>getset()</script>
<script type="text/javascript">var preview_page={{ .ChapterUrlBefore }};var next_page={{ .ChapterUrlAfter }};var index_page={{ .CatalogUrl }};</script>
</body>
{{end}}

{{define "searchbar"}}
<div class="searchbar">
    <form class="searchwrap" name="ssform" action="/book/search" method="get"><input type="text" name="searchkey" id="searchkey"   onblur="if(this.value==''){this.value='请输入小说名或作者...';this.style.color='#aaa'}" onfocus="if(this.value=='请输入小说名或作者...'){this.value='';this.style.color='#666'}" value="请输入小说名或作者..." class="search-text" /><span><input class="search-submit" type="submit" value="" /></span></form>
  </div>
{{end}}

"layer.open({content: '你已成功将本书加入书架',btn: '我知道了',})"

{{define "catalog"}}
<div id="main">
<script type="text/javascript">
submit = function(book_id){
$.ajax({  
                type: "GET",  
                url:"/addbookcase/book/" + book_id.toString(),
                async: false,  
                success: function(data) {
                    if (data == 'success'){
                        layer.open({
                            content: "你已经成功将本书加入书架",
                            btn: '我知道了', 
                        });
                    }
                    else{
                        layer.open({
                            content: data,
                            btn: '我知道了', 
                        });
                    }   
                }  
            });  
}
</script>
    <div class="book-about clrfix">
      <div class="book-info clrfix">
      <div class="l"><h1>{{.BookName}}</h1><em>作者：{{.BookAuthor}}</em></div>
      <div class="r"><a href="javascript:void(0)" onclick=submit({{.BookId}})>加入书架</a><a href="javascript:void(0)" onclick = "layer.open({content: '你已成功为本书投了一张推荐票',btn: '我知道了',})">投推荐票</a><a href="#footer">直达底部</a></div>
    </div>
    <p class="stats"><span class="l"><b>最新章节：</b><a href={{.LastChapterUrl}}>{{.LastChapterName}}</a></span><span class="l"><b>上次阅读：</b><a href={{.LastReadChapterUrl}}>{{.LastReadChapterName}}</a></span><span class="r">状态：<i>{{.BookStatus}}</i>章节数：<i>{{.BookChapterNum}}</i>更新时间：<i>{{.UpdateTime}}</i></span></p>
      <p class="intro"><b>内容简介：{{.BookIntro}}
    </p>
    </div>
    <dl class="chapter-list clrfix">
    {{range .Chapters}}
    {{if .New }}
    <dd><a href={{.ChapterUrl}}>{{.ChapterName}}<p style="float:right;color:red">最新章节</p></a></dd>
    {{else}}
    <dd><a href={{.ChapterUrl}}>{{.ChapterName}}</a></dd>
    {{end}}
    {{end}}
</div>
{{end}}

{{define "indexitem"}}
<section class="inner">
<div class="title"><h3>{{ .Title }}
{{if .More}}
<a href="{{.MoreUrl}}" style="float:right;color:#129151">更多&gt;</a>
{{end}}
</h3></div>
<div class="details clrfix">
<ul class="item-list">
{{$itemLen := len .BookItems}}
{{if gt $itemLen 0 }}
{{range .BookItems}}
{{if .New}}
<li>[{{.BookType1}}]&nbsp;&nbsp;<a href={{.BookUrl}}>{{.BookName}}</a>&nbsp;&nbsp;{{.BookAuthor}}<p style="float:right;color:red">有更新</p></li>
{{else}}
<li>[{{.BookType1}}]&nbsp;&nbsp;<a href={{.BookUrl}}>{{.BookName}}</a>&nbsp;&nbsp;{{.BookAuthor}}</li>
{{end}}
{{end}}
{{else}}
查无结果
{{end}}
</ul>
</div>
</section>
{{end}}

{{define "index"}}
{{range .IndexItems}}
{{template "indexitem" .}}
{{end}}
{{end}}

{{define "shujiaitem"}}
<section class="inner">
<div class="title"><h3>{{ .Title }}
{{if .More}}
<a href="{{.MoreUrl}}" style="float:right;color:#129151">更多&gt;</a>
{{end}}
</h3></div>
<div class="details clrfix">
<ul class="item-list">
{{$itemLen := len .ShujiaItems}}
{{if gt $itemLen 0 }}
{{range .ShujiaItems}}
{{if .New}}
<li><span>书名：</span><a href={{.BookUrl}}><font color="#129151">{{ .BookName}}</font></a><em><a href="{{.RemoveUrl}}"><font color="black">[移除]</font></a></em><em><a href="{{.BookUrl}}"><font color="red">有更新&nbsp&nbsp&nbsp</font></a></em></li>
{{else}}
<li><span>书名：</span><a href={{.BookUrl}}><font color="#129151">{{ .BookName}}</font></a><em><a href="{{.RemoveUrl}}"><font color="black">[移除]</font></a></em></li>
{{end}}
{{end}}
{{else}}
查无结果
{{end}}
</ul>
</div>
</section>
{{end}}
