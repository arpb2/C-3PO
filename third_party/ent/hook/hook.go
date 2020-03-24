// Code generated by entc, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/arpb2/C-3PO/third_party/ent"
)

// The ClassroomFunc type is an adapter to allow the use of ordinary
// function as Classroom mutator.
type ClassroomFunc func(context.Context, *ent.ClassroomMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ClassroomFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ClassroomMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ClassroomMutation", m)
	}
	return f(ctx, mv)
}

// The CredentialFunc type is an adapter to allow the use of ordinary
// function as Credential mutator.
type CredentialFunc func(context.Context, *ent.CredentialMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CredentialFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.CredentialMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CredentialMutation", m)
	}
	return f(ctx, mv)
}

// The LevelFunc type is an adapter to allow the use of ordinary
// function as Level mutator.
type LevelFunc func(context.Context, *ent.LevelMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f LevelFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.LevelMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.LevelMutation", m)
	}
	return f(ctx, mv)
}

// The UserFunc type is an adapter to allow the use of ordinary
// function as User mutator.
type UserFunc func(context.Context, *ent.UserMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.UserMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserMutation", m)
	}
	return f(ctx, mv)
}

// The UserLevelFunc type is an adapter to allow the use of ordinary
// function as UserLevel mutator.
type UserLevelFunc func(context.Context, *ent.UserLevelMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserLevelFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.UserLevelMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserLevelMutation", m)
	}
	return f(ctx, mv)
}

// On executes the given hook only of the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
//
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if m.Op().Is(op) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
//
func Reject(op ent.Op) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if m.Op().Is(op) {
				return nil, fmt.Errorf("%s operation is not allowed", m.Op())
			}
			return next.Mutate(ctx, m)
		})
	}
}
