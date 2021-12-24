package shamir

// @license: [MPL-2.0 License](https://github.com/hashicorp/vault/blob/main/LICENSE)
// @source:  https://github.com/hashicorp/vault/blob/main/shamir/tables_test.go

import "testing"

func TestTables(t *testing.T) {
	for i := 1; i < 256; i++ {
		logV := logTable[i]
		expV := expTable[logV]
		if expV != uint8(i) {
			t.Fatalf("bad: %d log: %d exp: %d", i, logV, expV)
		}
	}
}
