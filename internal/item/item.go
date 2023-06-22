package item

import (
	"regexp"
	"strings"
)

type Item struct {
	Class     string `json:"class"`
	Service   string `json:"service"`
	Account   string `json:"account"`
	Attribute string `json:"attribute"`
	Label     string `json:"label"`
	Comment   string `json:"comment"`
}

const (
	GenericPasswordClass  = "genp"
	InternetPasswordClass = "inet"
	CertificateClass      = "cert"
	KeyClass              = "key"
)

var (
	serviceRegex    = regexp.MustCompile(`^\s+"svce"<blob>="(?P<Service>.*)"`)
	accountRegex    = regexp.MustCompile(`^\s+"acct"<blob>="(?P<Account>.*)"`)
	attributeRegex  = regexp.MustCompile(`^\s+"gena"<blob>="(?P<Attribute>.*)"`)
	commentRegex    = regexp.MustCompile(`^\s+"icmt"<blob>="(?P<Comment>.*)"`)
	labelRegex      = regexp.MustCompile(`^\s+0x00000007 <blob>="(?P<Label>.*)"`)
	classRegex      = regexp.MustCompile(`^\s*class:\s+"(?P<Class>.*)"`)
	serviceParser   = lineParser(serviceRegex, "Service")
	accountParser   = lineParser(accountRegex, "Account")
	attributeParser = lineParser(attributeRegex, "Attribute")
	commentParser   = lineParser(commentRegex, "Comment")
	labelParser     = lineParser(labelRegex, "Label")
	classParser     = lineParser(classRegex, "Class")
	knownClasses    = map[string]struct{}{
		GenericPasswordClass:  {},
		InternetPasswordClass: {},
		CertificateClass:      {},
		KeyClass:              {},
	}
)

func lineParser(r *regexp.Regexp, capturedGroupName string) func(string) (string, bool) {
	return func(line string) (string, bool) {
		matches := r.FindStringSubmatch(line)
		if len(matches) == 0 {
			return "", false
		}

		return matches[r.SubexpIndex(capturedGroupName)], true
	}
}

func From(lines []string) []Item {
	var items []Item
	var current *Item
	for _, l := range lines {
		if strings.HasPrefix(l, "keychain:") {
			if current != nil {
				if _, known := knownClasses[current.Class]; known {
					items = append(items, *current)
				}
			}

			current = &Item{}
		} else {
			if current == nil {
				continue
			}

			if class, match := classParser(l); match {
				current.Class = class
				continue
			}

			if service, match := serviceParser(l); match {
				current.Service = service
				continue
			}

			if account, match := accountParser(l); match {
				current.Account = account
				continue
			}

			if attribute, match := attributeParser(l); match {
				current.Attribute = attribute
				continue
			}

			if comment, match := commentParser(l); match {
				current.Comment = comment
				continue
			}

			if label, match := labelParser(l); match {
				current.Label = label
			}
		}
	}

	if current != nil {
		if _, known := knownClasses[current.Class]; known {
			items = append(items, *current)
		}
	}

	return items
}
