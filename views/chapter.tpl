{{with .Data}}
<!DOCTYPE html>
<html>
<head>
{{template "head" .}}
</head>
{{template "topbar2" .}}
{{template "chapter" .}}
{{template "footerbar" .}}
</html>
{{end}}