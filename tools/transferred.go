package tools

import "strings"

func TransferredToHtml(str string) string {
	str = strings.Replace(str," ","&nbsp;",-1)
	str = strings.Replace(str,"\"","&quot;",-1)
	str = strings.Replace(str,"<","&lt;",-1)
	str = strings.Replace(str,">","&gt;",-1)
	return str
}

func UnTransferredToHtml(str string) string {
	str = strings.Replace(str,"&nbsp;"," ",-1)
	str = strings.Replace(str,"&quot;","\"",-1)
	str = strings.Replace(str,"&lt;","<",-1)
	str = strings.Replace(str,"&gt;",">",-1)
	return str
}

