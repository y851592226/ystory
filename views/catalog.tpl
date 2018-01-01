{{with .Data}}
<!DOCTYPE html>
<html>
<head>
{{template "head" .}}
</head>
{{template "topbar2" .}}
{{template "searchbar"}}
{{template "catalog" .}}
{{template "footerbar" .}}
</html>
{{end}}