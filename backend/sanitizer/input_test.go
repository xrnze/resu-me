package sanitizer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSanitize(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{name: "clean text", input: "experienced Go developer with 5 years", wantErr: false},
		{name: "legitimate angle brackets", input: "C++ > Java and 3 < 5", wantErr: false},
		{name: "normal punctuation", input: "Worked on REST APIs, SQL databases, and AWS", wantErr: false},
		{name: "common tech terms", input: "SELECT and INSERT are SQL commands used daily", wantErr: false},
		{name: "script tag", input: "<script>alert('x')</script>", wantErr: true},
		{name: "img with onerror", input: "<img src=x onerror=alert(1)>", wantErr: true},
		{name: "iframe tag", input: "<iframe src='http://evil.com'></iframe>", wantErr: true},
		{name: "div with script", input: "<div><script>evil()</script></div>", wantErr: true},
		{name: "mixed case script", input: "<ScRiPt>alert(1)</ScRiPt>", wantErr: true},
		{name: "javascript URL", input: "javascript:void(0)", wantErr: true},
		{name: "data text/html", input: "data:text/html,<script>alert(1)</script>", wantErr: true},
		{name: "javascript with eval", input: "javascript:eval('alert(1)')", wantErr: true},
		{name: "onerror attribute", input: "<img src=x onerror='alert(1)'>", wantErr: true},
		{name: "onclick attribute", input: "<button onclick='stealData()'>click</button>", wantErr: true},
		{name: "onload attribute", input: "<body onload='init()'>", wantErr: true},
		{name: "onmouseover", input: "<div onmouseover='evil()'>hover</div>", wantErr: true},
		{name: "DROP TABLE", input: "'; DROP TABLE users; --", wantErr: true},
		{name: "UNION SELECT", input: "UNION SELECT * FROM passwords", wantErr: true},
		{name: "INSERT INTO", input: "INSERT INTO users VALUES ('admin','pass')", wantErr: true},
		{name: "DELETE FROM", input: "DELETE FROM accounts WHERE 1=1", wantErr: true},
		{name: "OR 1=1", input: "' OR '1'='1", wantErr: true},
		{name: "URL-encoded script", input: "%3Cscript%3Ealert(1)%3C/script%3E", wantErr: true},
		{name: "HTML-encoded script", input: "&#60;script&#62;alert(1)&#60;/script&#62;", wantErr: true},
		{name: "hex encoded", input: "&#x3C;script&#x3E;alert(1)&#x3C;/script&#x3E;", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Sanitize(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
