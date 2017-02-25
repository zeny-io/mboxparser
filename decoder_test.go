package mboxparser

import (
	"testing"
)

func assertStrings(t *testing.T, actual, expected []string) {
	if len(actual) != len(expected) {
		t.Fatalf("Length missmatched: %d / %d", len(actual), len(expected))
	}

	for i, a := range actual {
		if a != expected[i] {
			t.Fatalf("String missmatched: %d: %s / %s", i, a, expected[i])
		}
	}
}

func Test_decodeHeaders(t *testing.T) {
	src := []string{
		"plain text",
		"=?UTF-8?B?5pel5pys6Kqe44OG44K544OI?=",
		"=?ISO-2022-JP?B?GyRCRnxLXDhsJUYlOSVIGyhC?=",
		"=?EUC-JP?B?xvzL3LjspcaluaXI?=",
		"=?SHIFT_JIS?B?k/qWe4zqg2WDWINn?=",
		"=?ISO-2022-JP?Q?=1B=24BF=7CK=5C8l=25F=259=25H=1B=28B?=",
	}

	dst := decodeHeaders(src)

	assertStrings(t, dst, []string{
		"plain text",
		"日本語テスト",
		"日本語テスト",
		"日本語テスト",
		"日本語テスト",
		"日本語テスト",
	})
}
