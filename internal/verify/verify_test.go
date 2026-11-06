package verify

import (
	"strings"
	"testing"
)

func TestParseBuildErrors(t *testing.T) {
	res := &Result{}
	parse(`{"Action":"start","Package":"x"}
{"Action":"build-output","Output":"# x\n"}
{"Action":"build-output","Output":"./exercise.go:8:2: undefined: Message\n"}
{"Action":"build-output","Output":"./exercise.go:9:14: too many arguments in call to fmt.Println\n"}
{"Action":"build-fail","Package":"x"}
`, res)

	if res.OK {
		t.Fatal("expected failure")
	}
	if len(res.BuildErrors) != 2 {
		t.Fatalf("want 2 build errors, got %d: %+v", len(res.BuildErrors), res.BuildErrors)
	}
	e := res.BuildErrors[0]
	if e.File != "exercise.go" || e.Line != "8" || e.Col != "2" || e.Msg != "undefined: Message" {
		t.Errorf("parsed error wrong: %+v", e)
	}
}

func TestParseFailingTest(t *testing.T) {
	res := &Result{}
	parse(`{"Action":"start","Package":"x"}
{"Action":"run","Test":"TestAdd"}
{"Action":"output","Test":"TestAdd","Output":"=== RUN   TestAdd\n"}
{"Action":"output","Test":"TestAdd","Output":"    exercise_test.go:10: add(2, 3) = 0, want 5\n"}
{"Action":"output","Test":"TestAdd","Output":"--- FAIL: TestAdd (0.00s)\n"}
{"Action":"fail","Test":"TestAdd"}
{"Action":"output","Output":"FAIL\n"}
{"Action":"fail","Package":"x"}
`, res)

	if res.OK {
		t.Fatal("expected failure")
	}
	if res.TotalTests != 1 {
		t.Errorf("TotalTests = %d, want 1", res.TotalTests)
	}
	if len(res.FailedTests) != 1 || res.FailedTests[0].Test != "TestAdd" {
		t.Fatalf("failed tests wrong: %+v", res.FailedTests)
	}
	out := res.FailedTests[0].Output
	if !strings.Contains(out, "add(2, 3) = 0, want 5") {
		t.Errorf("useful output lost: %q", out)
	}
	if strings.Contains(out, "=== RUN") || strings.Contains(out, "--- FAIL") {
		t.Errorf("noise not filtered: %q", out)
	}
}

func TestParsePanic(t *testing.T) {
	res := &Result{}
	parse(`{"Action":"start","Package":"x"}
{"Action":"output","Output":"panic: runtime error: index out of range [5] with length 3\n"}
{"Action":"output","Output":"\ngoroutine 1 [running]:\nmain.main()\n\texercise.go:10 +0x1a\n"}
{"Action":"fail","Package":"x"}
`, res)

	if res.OK {
		t.Fatal("expected failure")
	}
	if !strings.Contains(res.PanicOutput, "index out of range") {
		t.Errorf("panic output missing: %q", res.PanicOutput)
	}
}

func TestParsePass(t *testing.T) {
	res := &Result{}
	parse(`{"Action":"start","Package":"x"}
{"Action":"run","Test":"TestAdd"}
{"Action":"output","Test":"TestAdd","Output":"=== RUN   TestAdd\n"}
{"Action":"pass","Test":"TestAdd"}
{"Action":"output","Output":"PASS\n"}
{"Action":"pass","Package":"x"}
`, res)

	if !res.OK {
		t.Fatal("expected success")
	}
	if res.TotalTests != 1 || len(res.FailedTests) != 0 {
		t.Errorf("wrong counts: total=%d failed=%d", res.TotalTests, len(res.FailedTests))
	}
}
