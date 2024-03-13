package result

import (
	"testing"
)

func TestResult(t *testing.T) {
	result := NewResult()
	result.AddVul("127.0.0.1", 80, &Vulnerability{"sql", "?id=1&name=a' or 1=1#"})
	result.AddVul("127.0.0.1", 80, &Vulnerability{"xss", "?name=<script>alert(1)</script>"})

	result.AddVul("127.0.0.1", 443, &Vulnerability{"sql", "?id=1&name=a' or 1=1#"})
	result.PrintMapValues()
}
