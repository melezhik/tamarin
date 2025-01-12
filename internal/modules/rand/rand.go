package rand

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/cloudcmds/tamarin/internal/arg"
	"github.com/cloudcmds/tamarin/internal/scope"
	"github.com/cloudcmds/tamarin/object"
)

// Name of this module
const Name = "rand"

func Float(ctx context.Context, args ...object.Object) object.Object {
	if err := arg.Require("rand.float", 0, args); err != nil {
		return err
	}
	return object.NewFloat(rand.Float64())
}

func Int(ctx context.Context, args ...object.Object) object.Object {
	if err := arg.Require("rand.int", 0, args); err != nil {
		return err
	}
	return object.NewInteger(rand.Int63())
}

func IntN(ctx context.Context, args ...object.Object) object.Object {
	if err := arg.Require("rand.intn", 1, args); err != nil {
		return err
	}
	n, err := object.AsInteger(args[0])
	if err != nil {
		return err
	}
	return object.NewInteger(rand.Int63n(n))
}

func NormFloat(ctx context.Context, args ...object.Object) object.Object {
	if err := arg.Require("rand.norm_float", 0, args); err != nil {
		return err
	}
	return object.NewFloat(rand.NormFloat64())
}

func ExpFloat(ctx context.Context, args ...object.Object) object.Object {
	if err := arg.Require("rand.exp_float", 0, args); err != nil {
		return err
	}
	return object.NewFloat(rand.ExpFloat64())
}

func Shuffle(ctx context.Context, args ...object.Object) object.Object {
	if err := arg.Require("rand.shuffle", 1, args); err != nil {
		return err
	}
	arr, err := object.AsArray(args[0])
	if err != nil {
		return err
	}
	elements := arr.Elements
	rand.Shuffle(len(elements), func(i, j int) {
		elements[i], elements[j] = elements[j], elements[i]
	})
	return arr
}

func Module(parentScope *scope.Scope) (*object.Module, error) {
	s := scope.New(scope.Opts{
		Name:   fmt.Sprintf("module:%s", Name),
		Parent: parentScope,
	})

	m := &object.Module{Name: Name, Scope: s}

	if err := s.AddBuiltins([]*object.Builtin{
		{Module: m, Name: "float", Fn: Float},
		{Module: m, Name: "int", Fn: Int},
		{Module: m, Name: "intn", Fn: IntN},
		{Module: m, Name: "norm_float", Fn: NormFloat},
		{Module: m, Name: "exp_float", Fn: ExpFloat},
		{Module: m, Name: "shuffle", Fn: Shuffle},
	}); err != nil {
		return nil, err
	}
	return m, nil
}
