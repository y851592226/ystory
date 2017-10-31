{{define "head"}}
<head>
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
</head>
{{end}}


{{define "topbar1"}}
<div class="topbar">
{{with .Topbar}}
<span class="logo"><a href={{ .Item[0].Url }} >{{  .Item[0].Name }}</a></span><span class="user"><a href={{ .Item[1].Url }}>{{ .Item[1].Name }}</a><a href={{ .Item[2].Url}} >{{ .Item[2].Name }}</a></span>
{{end}}
</div>
{{end}}


{{define "topbar2"}}
<div class="topbar">
{{with .Topbar}}
<span class="c1"><a href={{ .Item[0].Url }} >{{  .Item[0].Name }}</a></span><span class="c3">{{ .Item[1].Name }}</span><span class="c4"><a href={{ .Item[2].Url }} >{{ .Item[2].Name }}</a></span>
{{end}}
</div>
{{end}}