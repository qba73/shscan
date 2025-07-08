package shscan_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/qba73/shscan"
)

func TestGenerateHostsFromNetAddress(t *testing.T) {
	t.Parallel()

	ipaddresses, err := shscan.GenerateHostRange("192.168.1.0/24")
	if err != nil {
		t.Fatal(err)
	}

	want := 254
	got := len(ipaddresses)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGenerateHostsFailsOnInvalidNetAddress(t *testing.T) {
	t.Parallel()

	netAddress := "192.168.1.0/45"
	_, err := shscan.GenerateHostRange(netAddress)
	if err == nil {
		t.Errorf("want err on invalid address %s, got nil", netAddress)
	}
}
