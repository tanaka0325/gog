package sgogen

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_New(t *testing.T) {
	type args struct {
		cmd      string
		isDryRun bool
		paths    []string
	}

	tests := map[string]struct {
		args args
		want *Generator
	}{
		"normal": {
			args: args{
				cmd:      "go",
				isDryRun: true,
				paths:    []string{"app.go"},
			},
			want: &Generator{
				cmd:      "go",
				isDryRun: true,
				paths:    []string{"app.go"},
			},
		},
		"empty path": {
			args: args{
				cmd:   "ls",
				paths: []string{},
			},
			want: &Generator{
				cmd:   "ls",
				paths: []string{"."},
			},
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			got := New(tt.args.cmd, tt.args.isDryRun, tt.args.paths)

			opt := cmp.AllowUnexported(Generator{})
			if diff := cmp.Diff(tt.want, got, opt); diff != "" {
				t.Errorf("onegen.New failed (-want +got):%s", diff)
			}
		})
	}
}

func ExampleGenerator_Gen_not_executed() {
	g := &Generator{
		cmd:   "",
		paths: []string{"docs/sample1.go"},
	}

	g.Gen()
	// Output:
}

func ExampleGenerator_Gen_dryrun() {
	g := &Generator{
		cmd:      "echo",
		isDryRun: true,
		paths:    []string{"docs/sample1.go"},
	}

	g.Gen()
	// Output:
	// echo sample1
}

func ExampleGenerator_Gen_executed_and_print_results() {
	g := &Generator{
		cmd:   "echo",
		paths: []string{"docs/sample1.go"},
	}

	g.Gen()
	// Output:
	// echo sample1
	// sample1
}

func ExampleGenerator_Gen_executed_multi_files() {
	g := &Generator{
		cmd:   "echo",
		paths: []string{"docs/sample1.go", "docs/sample2.go"},
	}

	g.Gen()
	// Output:
	// echo sample1
	// sample1
	// echo sample2
	// sample2
}
