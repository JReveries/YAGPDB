{{/* 抓取所有參數 */}}
{{ $a := .CmdArgs }}

{{/* 用迴圈輸出所有參數 */}}
{{ range $a }}
	{{/* "." 為 $a 中的每個值，前後加上 "-" 輸出不會換行 */}}
	{{- . -}}
	{{- " " -}}
{{ end }}

{{/* 刪除觸發指令的訊息 */}}
{{ deleteTrigger 0 }}
