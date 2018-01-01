{{with .Data}}
<!DOCTYPE html>
<html>
<head>
{{template "head" .}}
</head>
{{template "topbar1" .}}
{{template "subnavbar" .}}
{{template "searchbar"}}
{{template "index" .}}
{{template "footerbar" .}}
</html>
{{end}}