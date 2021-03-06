// Code generated by entc, DO NOT EDIT.

package privacy

import (
	"context"
	"errors"
	"fmt"

	"github.com/arpb2/C-3PO/third_party/ent"
)

var (
	// Allow may be returned by rules to indicate that the policy
	// evaluation should terminate with an allow decision.
	Allow = errors.New("ent/privacy: allow rule")

	// Deny may be returned by rules to indicate that the policy
	// evaluation should terminate with an deny decision.
	Deny = errors.New("ent/privacy: deny rule")

	// Skip may be returned by rules to indicate that the policy
	// evaluation should continue to the next rule.
	Skip = errors.New("ent/privacy: skip rule")
)

// Allowf returns an formatted wrapped Allow decision.
func Allowf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Allow)...)
}

// Denyf returns an formatted wrapped Deny decision.
func Denyf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Deny)...)
}

// Skipf returns an formatted wrapped Skip decision.
func Skipf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Skip)...)
}

type (
	// QueryPolicy combines multiple query rules into a single policy.
	QueryPolicy []QueryRule

	// QueryRule defines the interface deciding whether a
	// query is allowed and optionally modify it.
	QueryRule interface {
		EvalQuery(context.Context, ent.Query) error
	}
)

// EvalQuery evaluates a query against a query policy.
func (policy QueryPolicy) EvalQuery(ctx context.Context, q ent.Query) error {
	for _, rule := range policy {
		switch err := rule.EvalQuery(ctx, q); {
		case err == nil || errors.Is(err, Skip):
		case errors.Is(err, Allow):
			return nil
		default:
			return err
		}
	}
	return nil
}

// QueryRuleFunc type is an adapter to allow the use of
// ordinary functions as query rules.
type QueryRuleFunc func(context.Context, ent.Query) error

// Eval returns f(ctx, q).
func (f QueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	return f(ctx, q)
}

type (
	// MutationPolicy combines multiple mutation rules into a single policy.
	MutationPolicy []MutationRule

	// MutationRule defines the interface deciding whether a
	// mutation is allowed and optionally modify it.
	MutationRule interface {
		EvalMutation(context.Context, ent.Mutation) error
	}
)

// EvalMutation evaluates a mutation against a mutation policy.
func (policy MutationPolicy) EvalMutation(ctx context.Context, m ent.Mutation) error {
	for _, rule := range policy {
		switch err := rule.EvalMutation(ctx, m); {
		case err == nil || errors.Is(err, Skip):
		case errors.Is(err, Allow):
			return nil
		default:
			return err
		}
	}
	return nil
}

// MutationRuleFunc type is an adapter to allow the use of
// ordinary functions as mutation rules.
type MutationRuleFunc func(context.Context, ent.Mutation) error

// EvalMutation returns f(ctx, m).
func (f MutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	return f(ctx, m)
}

// Policy groups query and mutation policies.
type Policy struct {
	Query    QueryPolicy
	Mutation MutationPolicy
}

// EvalQuery forwards evaluation to query policy.
func (policy Policy) EvalQuery(ctx context.Context, q ent.Query) error {
	return policy.Query.EvalQuery(ctx, q)
}

// EvalMutation forwards evaluation to mutation policy.
func (policy Policy) EvalMutation(ctx context.Context, m ent.Mutation) error {
	return policy.Mutation.EvalMutation(ctx, m)
}

// QueryMutationRule is the interface that groups query and mutation rules.
type QueryMutationRule interface {
	QueryRule
	MutationRule
}

// AlwaysAllowRule returns a rule that returns an allow decision.
func AlwaysAllowRule() QueryMutationRule {
	return fixedDecisionRule{Allow}
}

// AlwaysDenyRule returns a rule that returns a deny decision.
func AlwaysDenyRule() QueryMutationRule {
	return fixedDecisionRule{Deny}
}

type fixedDecisionRule struct{ err error }

func (f fixedDecisionRule) EvalQuery(context.Context, ent.Query) error       { return f.err }
func (f fixedDecisionRule) EvalMutation(context.Context, ent.Mutation) error { return f.err }

// The ClassroomQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type ClassroomQueryRuleFunc func(context.Context, *ent.ClassroomQuery) error

// EvalQuery return f(ctx, q).
func (f ClassroomQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ClassroomQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.ClassroomQuery", q)
}

// The ClassroomMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type ClassroomMutationRuleFunc func(context.Context, *ent.ClassroomMutation) error

// EvalMutation calls f(ctx, m).
func (f ClassroomMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.ClassroomMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.ClassroomMutation", m)
}

// The CredentialQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type CredentialQueryRuleFunc func(context.Context, *ent.CredentialQuery) error

// EvalQuery return f(ctx, q).
func (f CredentialQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.CredentialQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.CredentialQuery", q)
}

// The CredentialMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type CredentialMutationRuleFunc func(context.Context, *ent.CredentialMutation) error

// EvalMutation calls f(ctx, m).
func (f CredentialMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.CredentialMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.CredentialMutation", m)
}

// The LevelQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type LevelQueryRuleFunc func(context.Context, *ent.LevelQuery) error

// EvalQuery return f(ctx, q).
func (f LevelQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.LevelQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.LevelQuery", q)
}

// The LevelMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type LevelMutationRuleFunc func(context.Context, *ent.LevelMutation) error

// EvalMutation calls f(ctx, m).
func (f LevelMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.LevelMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.LevelMutation", m)
}

// The UserQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type UserQueryRuleFunc func(context.Context, *ent.UserQuery) error

// EvalQuery return f(ctx, q).
func (f UserQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.UserQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.UserQuery", q)
}

// The UserMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type UserMutationRuleFunc func(context.Context, *ent.UserMutation) error

// EvalMutation calls f(ctx, m).
func (f UserMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.UserMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.UserMutation", m)
}

// The UserLevelQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type UserLevelQueryRuleFunc func(context.Context, *ent.UserLevelQuery) error

// EvalQuery return f(ctx, q).
func (f UserLevelQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.UserLevelQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.UserLevelQuery", q)
}

// The UserLevelMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type UserLevelMutationRuleFunc func(context.Context, *ent.UserLevelMutation) error

// EvalMutation calls f(ctx, m).
func (f UserLevelMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.UserLevelMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.UserLevelMutation", m)
}
