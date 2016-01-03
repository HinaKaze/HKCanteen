package content

import (
	"blog/tools/iniparse"
)

var serverContent struct {
	WebPort int
}

func LoadServerContent() {
	iniparse.DefaultParse("./content/config.ini")
	s, ok := iniparse.GetSection("WEB")
	if ok {
		serverContent.WebPort = s.GetIntValue("webPort")
	}
}

func GetWebPort() int {
	return serverContent.WebPort
}
