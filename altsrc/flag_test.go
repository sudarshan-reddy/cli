package altsrc

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/urfave/cli/v2"
)

type testApplyInputSource struct {
	Flag               FlagInputSourceExtension
	FlagName           string
	FlagSetName        string
	Expected           string
	ContextValueString string
	EnvVarValue        string
	EnvVarName         string
	SourcePath         string
	MapValue           interface{}
}

func TestGenericApplyInputSourceValue(t *testing.T) {
	v := &Parser{"abc", "def"}
	runTest(
		t,
		testApplyInputSource{
			Flag:     NewGenericFlag(&cli.GenericFlag{Name: "test", Value: &Parser{}}),
			FlagName: "test",
			MapValue: v,
		},
		func(c *cli.Context) {
			expect(t, v, c.Generic("test"))
		},
	)
}

func TestGenericApplyInputSourceMethodContextSet(t *testing.T) {
	p := &Parser{"abc", "def"}
	runTest(
		t,
		testApplyInputSource{
			Flag:               NewGenericFlag(&cli.GenericFlag{Name: "test", Value: &Parser{}}),
			FlagName:           "test",
			MapValue:           &Parser{"efg", "hig"},
			ContextValueString: p.String(),
		},
		func(c *cli.Context) {
			expect(t, p, c.Generic("test"))
		},
	)
}

func TestGenericApplyInputSourceMethodEnvVarSet(t *testing.T) {
	runTest(
		t,
		testApplyInputSource{
			Flag: NewGenericFlag(&cli.GenericFlag{
				Name:    "test",
				Value:   &Parser{},
				EnvVars: []string{"TEST"},
			}),
			FlagName:    "test",
			MapValue:    &Parser{"efg", "hij"},
			EnvVarName:  "TEST",
			EnvVarValue: "abc,def",
		},
		func(c *cli.Context) {
			expect(t, &Parser{"abc", "def"}, c.Generic("test"))
		},
	)
}

func TestStringSliceApplyInputSourceValue(t *testing.T) {
	runTest(
		t,
		testApplyInputSource{
			Flag:     NewStringSliceFlag(&cli.StringSliceFlag{Name: "test"}),
			FlagName: "test",
			MapValue: []interface{}{"hello", "world"},
		},
		func(c *cli.Context) {
			expect(t, []string{"hello", "world"}, c.StringSlice("test"))
		},
	)
}

func TestStringSliceApplyInputSourceMethodContextSet(t *testing.T) {
	runTest(
		t,
		testApplyInputSource{
			Flag:               NewStringSliceFlag(&cli.StringSliceFlag{Name: "test"}),
			FlagName:           "test",
			MapValue:           []interface{}{"hello", "world"},
			ContextValueString: "ohno",
		},
		func(c *cli.Context) {
			expect(t, []string{"ohno"}, c.StringSlice("test"))
		},
	)
}

func TestStringSliceApplyInputSourceMethodEnvVarSet(t *testing.T) {
	runTest(
		t,
		testApplyInputSource{
			Flag:        NewStringSliceFlag(&cli.StringSliceFlag{Name: "test", EnvVars: []string{"TEST"}}),
			FlagName:    "test",
			MapValue:    []interface{}{"hello", "world"},
			EnvVarName:  "TEST",
			EnvVarValue: "oh,no",
		},
		func(c *cli.Context) {
			expect(t, []string{"oh", "no"}, c.StringSlice("test"))
		},
	)
}

func TestIntSliceApplyInputSourceValue(t *testing.T) {
	runTest(
		t,
		testApplyInputSource{
			Flag:     NewIntSliceFlag(&cli.IntSliceFlag{Name: "test"}),
			FlagName: "test",
			MapValue: []interface{}{1, 2},
		},
		func(c *cli.Context) {
			expect(t, []int{1, 2}, c.IntSlice("test"))
		},
	)
}

func TestIntSliceApplyInputSourceMethodContextSet(t *testing.T) {
	runTest(
		t,
		testApplyInputSource{
			Flag:               NewIntSliceFlag(&cli.IntSliceFlag{Name: "test"}),
			FlagName:           "test",
			MapValue:           []interface{}{1, 2},
			ContextValueString: "3",
		},
		func(c *cli.Context) {
			expect(t, []int{3}, c.IntSlice("test"))
		},
	)
}

func TestIntSliceApplyInputSourceMethodEnvVarSet(t *testing.T) {
	runTest(
		t,
		testApplyInputSource{
			Flag:        NewIntSliceFlag(&cli.IntSliceFlag{Name: "test", EnvVars: []string{"TEST"}}),
			FlagName:    "test",
			MapValue:    []interface{}{1, 2},
			EnvVarName:  "TEST",
			EnvVarValue: "3,4",
		},
		func(c *cli.Context) {
			expect(t, []int{3, 4}, c.IntSlice("test"))
		},
	)
}

func TestBoolApplyInputSourceMethodSet(t *testing.T) {
	runTest(
		t,
		testApplyInputSource{
			Flag:     NewBoolFlag(&cli.BoolFlag{Name: "test"}),
			FlagName: "test",
			MapValue: true,
		},
		func(c *cli.Context) {
			expect(t, true, c.Bool("test"))
		},
	)
}

func TestBoolApplyInputSourceMethodContextSet(t *testing.T) {
	runTest(
		t,
		testApplyInputSource{
			Flag:               NewBoolFlag(&cli.BoolFlag{Name: "test"}),
			FlagName:           "test",
			MapValue:           false,
			ContextValueString: "true",
		},
		func(c *cli.Context) {
			expect(t, true, c.Bool("test"))
		},
	)
}

func TestBoolApplyInputSourceMethodEnvVarSet(t *testing.T) {
	runTest(
		t,
		testApplyInputSource{
			Flag:        NewBoolFlag(&cli.BoolFlag{Name: "test", EnvVars: []string{"TEST"}}),
			FlagName:    "test",
			MapValue:    false,
			EnvVarName:  "TEST",
			EnvVarValue: "true",
		},
		func(c *cli.Context) {
			expect(t, true, c.Bool("test"))
		},
	)
}

func TestStringApplyInputSourceMethodSet(t *testing.T) {
	runTest(
		t,
		testApplyInputSource{
			Flag:     NewStringFlag(&cli.StringFlag{Name: "test"}),
			FlagName: "test",
			MapValue: "hello",
		},
		func(c *cli.Context) {
			expect(t, "hello", c.String("test"))
		},
	)
}

func TestStringApplyInputSourceMethodContextSet(t *testing.T) {
	runTest(
		t,
		testApplyInputSource{
			Flag:               NewStringFlag(&cli.StringFlag{Name: "test"}),
			FlagName:           "test",
			MapValue:           "hello",
			ContextValueString: "goodbye",
		},
		func(c *cli.Context) {
			expect(t, "goodbye", c.String("test"))
		},
	)
}

func TestStringApplyInputSourceMethodEnvVarSet(t *testing.T) {
	runTest(
		t,
		testApplyInputSource{
			Flag:        NewStringFlag(&cli.StringFlag{Name: "test", EnvVars: []string{"TEST"}}),
			FlagName:    "test",
			MapValue:    "hello",
			EnvVarName:  "TEST",
			EnvVarValue: "goodbye",
		},
		func(c *cli.Context) {
			expect(t, "goodbye", c.String("test"))
		},
	)
}

func TestPathApplyInputSourceMethodSet(t *testing.T) {
	expected := "/path/to/source/hello"
	if runtime.GOOS == "windows" {
		var err error
		// Prepend the corresponding drive letter (or UNC path?), and change
		// to windows-style path:
		expected, err = filepath.Abs(expected)
		if err != nil {
			t.Fatal(err)
		}
	}

	runTest(
		t,
		testApplyInputSource{
			Flag:       NewPathFlag(&cli.PathFlag{Name: "test"}),
			FlagName:   "test",
			MapValue:   "hello",
			SourcePath: "/path/to/source/file",
		},
		func(c *cli.Context) {
			expect(t, expected, c.String("test"))
		},
	)
}

func TestPathApplyInputSourceMethodContextSet(t *testing.T) {
	runTest(
		t,
		testApplyInputSource{
			Flag:               NewPathFlag(&cli.PathFlag{Name: "test"}),
			FlagName:           "test",
			MapValue:           "hello",
			ContextValueString: "goodbye",
			SourcePath:         "/path/to/source/file",
		},
		func(c *cli.Context) {
			expect(t, "goodbye", c.String("test"))
		},
	)
}

func TestPathApplyInputSourceMethodEnvVarSet(t *testing.T) {
	runTest(
		t,
		testApplyInputSource{
			Flag:        NewPathFlag(&cli.PathFlag{Name: "test", EnvVars: []string{"TEST"}}),
			FlagName:    "test",
			MapValue:    "hello",
			EnvVarName:  "TEST",
			EnvVarValue: "goodbye",
			SourcePath:  "/path/to/source/file",
		},
		func(c *cli.Context) {
			expect(t, "goodbye", c.String("test"))
		},
	)
}

func TestIntApplyInputSourceMethodSet(t *testing.T) {
	runTest(
		t,
		testApplyInputSource{
			Flag:     NewIntFlag(&cli.IntFlag{Name: "test"}),
			FlagName: "test",
			MapValue: 15,
		},
		func(c *cli.Context) {
			expect(t, 15, c.Int("test"))
		},
	)
}
func TestIntApplyInputSourceMethodContextSet(t *testing.T) {
	runTest(
		t,
		testApplyInputSource{
			Flag:               NewIntFlag(&cli.IntFlag{Name: "test"}),
			FlagName:           "test",
			MapValue:           15,
			ContextValueString: "7",
		},
		func(c *cli.Context) {
			expect(t, 7, c.Int("test"))
		},
	)
}

func TestIntApplyInputSourceMethodEnvVarSet(t *testing.T) {
	runTest(
		t,
		testApplyInputSource{
			Flag:        NewIntFlag(&cli.IntFlag{Name: "test", EnvVars: []string{"TEST"}}),
			FlagName:    "test",
			MapValue:    15,
			EnvVarName:  "TEST",
			EnvVarValue: "12",
		},
		func(c *cli.Context) {
			expect(t, 12, c.Int("test"))
		},
	)
}

func TestDurationApplyInputSourceMethodSet(t *testing.T) {
	runTest(
		t,
		testApplyInputSource{
			Flag:     NewDurationFlag(&cli.DurationFlag{Name: "test"}),
			FlagName: "test",
			MapValue: 30 * time.Second,
		},
		func(c *cli.Context) {
			expect(t, 30*time.Second, c.Duration("test"))
		},
	)
}

func TestDurationApplyInputSourceMethodContextSet(t *testing.T) {
	runTest(
		t,
		testApplyInputSource{
			Flag:               NewDurationFlag(&cli.DurationFlag{Name: "test"}),
			FlagName:           "test",
			MapValue:           30 * time.Second,
			ContextValueString: (15 * time.Second).String(),
		},
		func(c *cli.Context) {
			expect(t, 15*time.Second, c.Duration("test"))
		},
	)
}

func TestDurationApplyInputSourceMethodEnvVarSet(t *testing.T) {
	runTest(
		t,
		testApplyInputSource{
			Flag:        NewDurationFlag(&cli.DurationFlag{Name: "test", EnvVars: []string{"TEST"}}),
			FlagName:    "test",
			MapValue:    30 * time.Second,
			EnvVarName:  "TEST",
			EnvVarValue: (15 * time.Second).String(),
		},
		func(c *cli.Context) {
			expect(t, 15*time.Second, c.Duration("test"))
		},
	)
}

func TestFloat64ApplyInputSourceMethodSet(t *testing.T) {
	runTest(
		t,
		testApplyInputSource{
			Flag:     NewFloat64Flag(&cli.Float64Flag{Name: "test"}),
			FlagName: "test",
			MapValue: 1.3,
		},
		func(c *cli.Context) {
			expect(t, 1.3, c.Float64("test"))
		},
	)
}

func TestFloat64ApplyInputSourceMethodContextSet(t *testing.T) {
	runTest(
		t,
		testApplyInputSource{
			Flag:               NewFloat64Flag(&cli.Float64Flag{Name: "test"}),
			FlagName:           "test",
			MapValue:           1.3,
			ContextValueString: fmt.Sprintf("%v", 1.4),
		},
		func(c *cli.Context) {
			expect(t, 1.4, c.Float64("test"))
		},
	)
}

func TestFloat64ApplyInputSourceMethodEnvVarSet(t *testing.T) {
	runTest(
		t,
		testApplyInputSource{
			Flag:        NewFloat64Flag(&cli.Float64Flag{Name: "test", EnvVars: []string{"TEST"}}),
			FlagName:    "test",
			MapValue:    1.3,
			EnvVarName:  "TEST",
			EnvVarValue: fmt.Sprintf("%v", 1.4),
		},
		func(c *cli.Context) {
			expect(t, 1.4, c.Float64("test"))
		},
	)
}

func runTest(t *testing.T, test testApplyInputSource, checkVal func(*cli.Context)) {
	inputSource := &MapInputSource{
		file:     test.SourcePath,
		valueMap: map[interface{}]interface{}{test.FlagName: test.MapValue},
	}

	app := cli.App{
		Flags: []cli.Flag{
			test.Flag,
		},
		Action: func(c *cli.Context) error {
			if test.ContextValueString != "" {
				_ = c.Set(test.FlagName, test.ContextValueString)
			}
			_ = test.Flag.ApplyInputSourceValue(c, inputSource)

			expect(t, c.IsSet(test.FlagName), true)
			checkVal(c)

			return nil
		},
	}

	if test.EnvVarName != "" && test.EnvVarValue != "" {
		_ = os.Setenv(test.EnvVarName, test.EnvVarValue)
		defer os.Setenv(test.EnvVarName, "")
	}

	_ = app.Run([]string{"the-app"})
}

type Parser [2]string

func (p *Parser) Set(value string) error {
	parts := strings.Split(value, ",")
	if len(parts) != 2 {
		return fmt.Errorf("invalid format")
	}

	(*p)[0] = parts[0]
	(*p)[1] = parts[1]

	return nil
}

func (p *Parser) String() string {
	return fmt.Sprintf("%s,%s", p[0], p[1])
}

// TestApplyInputSource verifies that ApplyInputSource correctly respects
// different hierarchical app/command/sub-command scenarios
func TestApplyInputSource(t *testing.T) {
	app := cli.App{
		Name: "the-app",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "simple",
			},
			NewStringFlag(&cli.StringFlag{
				Name: "loadable",
			}),
			NewStringFlag(&cli.StringFlag{
				Name: "loadable-shadow",
			}),
			&cli.StringFlag{
				Name: "simple-app",
			},
			NewStringFlag(&cli.StringFlag{
				Name: "loadable-app",
			}),
		},
		Commands: []*cli.Command{
			{
				Name: "one",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name: "simple",
					},
					NewStringFlag(&cli.StringFlag{
						Name: "loadable",
					}),
					&cli.StringFlag{
						Name: "loadable-shadow", // this make app's variable non-loadable
					},
					&cli.StringFlag{
						Name: "simple-one",
					},
					NewStringFlag(&cli.StringFlag{
						Name: "loadable-one",
					}),
				},
			},
			{
				Name: "two",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name: "simple",
					},
					NewStringFlag(&cli.StringFlag{
						Name: "loadable",
					}),
					&cli.StringFlag{
						Name: "override-two",
					},
					NewStringFlag(&cli.StringFlag{
						Name: "loadable-two",
					}),
				},
				Subcommands: []*cli.Command{
					{
						Name: "red",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name: "loadable-shadow", // this shadows loadable variable in the app
							},
							&cli.StringFlag{
								Name: "simple-red",
							},
							NewStringFlag(&cli.StringFlag{
								Name: "override-two", // this will make 'override-two' loadable
							}),
						},
					},
				},
			},
		},
	}

	genSource := func(kv ...string) InputSourceContext {
		m := make(map[interface{}]interface{})
		for i := 0; i < len(kv); i += 2 {
			m[kv[i]] = kv[i+1]
		}
		return NewMapInputSource("test-source", m)
	}

	tdata := []struct {
		testCase string
		source   InputSourceContext
		args     string
		check    []string
		expected map[string]string
	}{
		{
			testCase: "no args or file",
			source:   genSource(),
			args:     "app",
			check:    []string{"simple", "loadable", "loadable-shadow"},
			expected: make(map[string]string),
		},
		{
			testCase: "app only",
			source:   genSource("loadable", "APP", "loadable-shadow", "SHADOW", "simple", "SIMPLE"),
			args:     "app",
			check:    []string{"simple", "loadable", "loadable-shadow"},
			expected: map[string]string{
				"loadable":        "APP",
				"loadable-shadow": "SHADOW",
			},
		},
		{
			testCase: "cmd one",
			source:   genSource("loadable", "APP", "loadable-shadow", "SHADOW", "simple", "SIMPLE", "loadable-app", "APP"),
			args:     "app one",
			check:    []string{"simple", "loadable", "loadable-shadow", "loadable-app", "loadable-one"},
			expected: map[string]string{
				"loadable":        "APP",
				"loadable-shadow": "", // non-loadable flag hides loadable flag at higher level
				"loadable-app":    "APP",
			},
		},
		{
			testCase: "cmd two red",
			source:   genSource("loadable", "APP", "loadable-shadow", "SHADOW", "simple", "SIMPLE", "loadable-app", "APP", "loadable-two", "TWO", "override-two", "RED"),
			args:     "app two",
			check:    []string{"simple", "loadable", "loadable-shadow", "loadable-app", "loadable-two", "override-two"},
			expected: map[string]string{
				"loadable":        "APP",
				"loadable-shadow": "SHADOW", // subcommand "two" redefined loadable flag
				"loadable-app":    "APP",
				"loadable-two":    "TWO",
				"override-two":    "", // subcommand "two" defines as non-loadable
			},
		},
		{
			testCase: "cmd two",
			source:   genSource("loadable", "APP", "loadable-shadow", "SHADOW", "simple", "SIMPLE", "loadable-app", "APP", "override-two", "RED"),
			args:     "app two red",
			check:    []string{"simple", "loadable", "loadable-shadow", "loadable-app", "loadable-two", "override-two"},
			expected: map[string]string{
				"loadable":        "APP",
				"loadable-shadow": "", // subcommand "red" redefined as non-loadable
				"loadable-app":    "APP",
				"override-two":    "RED", // subcommand "red" redefined as loadable
			},
		},
	}

	for _, test := range tdata {
		t.Run(test.testCase, func(t *testing.T) {
			app.Action = func(context *cli.Context) error {
				if err := ApplyInputSource(context, test.source); err != nil {
					t.Fatal(test.testCase, err)
				}
				for _, name := range test.check {
					actual := context.String(name)
					if actual != test.expected[name] {
						t.Errorf("%s: expected %q got %q", name, test.expected[name], actual)
					}
				}

				return nil
			}
			for _, command := range app.Commands {
				command.Action = app.Action
				for _, sub := range command.Subcommands {
					sub.Action = app.Action
				}
			}

			args := strings.Split(test.args, " ")
			fmt.Println(args)
			err := app.Run(args)
			if err != nil {
				t.Errorf("%s action returned %v", test.testCase, err)
			}
		})
	}
}
