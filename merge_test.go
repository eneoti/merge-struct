package structs

import (
	// "reflect"
	"testing"
	"time"

	"github.com/AdrianLungu/decimal"
)

type testStruct struct {
	String   string
	Bool     bool
	Integer  int
	Float    float64
	Time     time.Time
	LargeNum testSubStruct
	Decimal  decimal.Decimal
}

type otherStruct struct {
	String string
}

type testSubStruct struct {
	IntPart int
	IsZero  bool
}

func TestMergeOverwrite(t *testing.T) {
	to := testStruct{
		String:  "default string",
		Bool:    true,
		Integer: 100,
		Float:   float64(1.001112),
		Time:    time.Now(),
		LargeNum: testSubStruct{
			IntPart: 1000,
			IsZero:  false,
		},
		Decimal: decimal.NewFromFloat(1000),
	}
	from := otherStruct{
		String: "another string",
	}
	var target testStruct
	if err := MergeOverwrite(to, from, &target); err != nil {
		t.Fatal(err)
	}
	// t.Logf("target: %+v", to)
	// t.Logf("to: %+v", target)

	if expected := from.String; target.String != expected {
		t.Errorf("want %s got %s", expected, target.String)
	}
}
