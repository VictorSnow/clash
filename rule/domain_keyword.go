package rules

import (
	"strings"

	C "github.com/Dreamacro/clash/constant"
)

type DomainKeyword struct {
	keyword string
	adapter string

	isMulti    bool
	arrKeyword []string
}

func (dk *DomainKeyword) RuleType() C.RuleType {
	return C.DomainKeyword
}

func (dk *DomainKeyword) Match(metadata *C.Metadata) bool {
	// fast fallover
	if !dk.isMulti {
		return strings.Contains(metadata.Host, dk.keyword)
	} else {
		for _, keyword := range dk.arrKeyword {
			if keyword != "" && !strings.Contains(metadata.Host, keyword) {
				return false
			}
		}

		return true
	}
}

func (dk *DomainKeyword) Adapter() string {
	return dk.adapter
}

func (dk *DomainKeyword) Payload() string {
	return dk.keyword
}

func (dk *DomainKeyword) ShouldResolveIP() bool {
	return false
}

func (dk *DomainKeyword) ShouldFindProcess() bool {
	return false
}

func NewDomainKeyword(keyword string, adapter string) *DomainKeyword {
	if strings.Index(keyword, " ") >= 0 {

		keywords := strings.Split(keyword, " ")

		return &DomainKeyword{
			keyword:    strings.ToLower(keyword),
			adapter:    adapter,
			isMulti:    true,
			arrKeyword: keywords,
		}
	} else {
		return &DomainKeyword{
			keyword:    strings.ToLower(keyword),
			adapter:    adapter,
			isMulti:    false,
			arrKeyword: nil,
		}
	}
}
