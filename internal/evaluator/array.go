package evaluator

import (
	"context"

	"github.com/myzie/tamarin/internal/ast"
	"github.com/myzie/tamarin/internal/object"
	"github.com/myzie/tamarin/internal/scope"
)

func (e *Evaluator) evalArrayLiteral(
	ctx context.Context,
	node *ast.ArrayLiteral,
	s *scope.Scope,
) object.Object {
	elements := e.evalExpressions(ctx, node.Elements, s)
	if len(elements) == 1 && isError(elements[0]) {
		return elements[0]
	}
	return &object.Array{Elements: elements}
}