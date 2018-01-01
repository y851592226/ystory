{{with .Data}}
<!DOCTYPE html>
<html>
<head>
{{template "head" .}}
</head>
{{template "topbar2" .}}
{{template "subnavbar" .}}
{{template "searchbar"}}
<style>
.book-top{font-weight:bold;padding:10px 5px;}
.bookcase{padding:0 5px;overflow:hidden;}
.bookcase ul{overflow:hidden;border-top:1px solid #ececec;padding:0 5px;}
.bookcase ul li{overflow:hidden;clear:both;padding:2px 0;}
.bookcase ul li span{color:#999;}
.bookcase ul li em{float:right;}
</style>
<form action="" method="post" name="checkform">
<div class="bookcase">
{{template "shujiaitem" .}}
</div>
</form>
{{template "footerbar" .}}
</html>
{{end}}
