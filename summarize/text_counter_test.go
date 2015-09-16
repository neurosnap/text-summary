package summarize

import "testing"

func TestTextCounter(t *testing.T) {
	tc := TextCounter{}

	tc.Add("foo", 5.0)
	tc.Add("bar")
	tc.Add("baz", 2.0)
	tc.Add("baz", 9.0)

	expInt := 5.0
	if tc["foo"] != expInt {
		t.Fatalf("Expected count of %d, got %d\n", expInt, tc["foo"])
	}

	expInt = 1.0
	if tc["bar"] != expInt {
		t.Fatalf("Expected count of %d, got %d\n", expInt, tc["foo"])
	}

	expInt = 11.0
	if tc["baz"] != expInt {
		t.Fatalf("Expected count of %d, got %d\n", expInt, tc["foo"])
	}

	common := tc.MostCommon(2)
	expInt = 2.0
	if len(common) != int(expInt) {
		t.Fatalf("Expected common length of %d, got %d\n", expInt, len(common))
	}

	expStr := "baz"
	if common[0].Text != expStr {
		t.Fatalf("Expected most common text to be '%s', got '%s'\n", expStr, common[0].Text)
	}

	expInt = 11.0
	if common[0].Count != expInt {
		t.Fatalf("Expected most common count to be '%d', got '%d'\n", expInt, common[0].Count)
	}

	expStr = "foo"
	if common[1].Text != expStr {
		t.Fatalf("Expected second most common text to be '%s', got '%s'\n", expStr, common[1].Text)
	}

	expInt = 5.0
	if common[1].Count != expInt {
		t.Fatalf("Expected second most common count to be '%d', got '%d'\n", expInt, common[1].Count)
	}
}
