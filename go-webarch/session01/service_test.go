package architecture

import (
	"github.com/golang/mock/gomock"
	"testing"
)

// using gomock
// mockgen -source service.go -destination mock_test.go -package architecture -self_package github.com/eblind39/gowebcf/go-webarch/session01
func TestPut(t *testing.T) {
	ctl := gomock.NewController(t)
	acc := NewMockAccesor(ctl)

	p := Persona{
		First: "Steven",
	}

	acc.EXPECT().Save(1, p).MinTimes(1).MaxTimes(1)

	// Put(acc, 2, p) // Fail
	Put(acc, 1, p) // Pass

	ctl.Finish()
}

/*
func ExamplePut() {
	mdb := Db{}
	p := Persona{
		First: "Steven",
	}
	psm := NewPersonService(mdb)
	psm.Put(1, p) // Pass
	got := mdb.Retrieve(1)
	fmt.Println(got)
	//Output: {Steven}
}
*/
