{{/* user預設為發訊息者，若有tag則將user改為其他使用者 */}}
{{ $user := .User }}
{{ if .Message.Mentions }}
	{{ $user = .Message.Mentions }}
{{ end }}

{{/* 取得使用者的頭像url */}}
{{ range $user }}
	{{ $url := ( userArg .ID ).AvatarURL "1024" }}

	{{/* 設定embed訊息 */}}
	{{ $embed := cembed 
		"image" ( sdict "url" $url )
		"color" 4645612
	}}

	{{/* 傳送embed訊息 */}}
	{{ sendMessage nil $embed }}

{{ end }}
