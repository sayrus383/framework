package application

import (
	"github.com/namsral/flag"
	"os"
	"time"
)

type config struct {
	fl   *flag.FlagSet
	args []string
}

func newFlagConfig() Config {
	return &config{
		fl:   flag.NewFlagSet("framework", flag.PanicOnError),
		args: os.Args[1:],
	}
}

func (c *config) BoolVar(p *bool, name string, value bool, usage string) {
	c.fl.BoolVar(p, name, value, usage)
	c.parse()
}

func (c *config) IntVar(p *int, name string, value int, usage string) {
	c.fl.IntVar(p, name, value, usage)
	c.parse()
}

func (c *config) Int64Var(p *int64, name string, value int64, usage string) {
	c.fl.Int64Var(p, name, value, usage)
	c.parse()
}

func (c *config) UintVar(p *uint, name string, value uint, usage string) {
	c.fl.UintVar(p, name, value, usage)
	c.parse()
}

func (c *config) Uint64Var(p *uint64, name string, value uint64, usage string) {
	c.fl.Uint64Var(p, name, value, usage)
	c.parse()
}

func (c *config) StringVar(p *string, name string, value string, usage string) {
	c.fl.StringVar(p, name, value, usage)
	c.parse()
}

func (c *config) Float64Var(p *float64, name string, value float64, usage string) {
	c.fl.Float64Var(p, name, value, usage)
	c.parse()
}

func (c *config) DurationVar(p *time.Duration, name string, value time.Duration, usage string) {
	c.fl.DurationVar(p, name, value, usage)
	c.parse()
}

func (c *config) parse() {
	_ = c.fl.Parse(c.args)
}
