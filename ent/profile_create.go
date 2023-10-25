// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"go-onion-arch-sample/ent/profile"
	"go-onion-arch-sample/ent/task"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProfileCreate is the builder for creating a Profile entity.
type ProfileCreate struct {
	config
	mutation *ProfileMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (pc *ProfileCreate) SetName(s string) *ProfileCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetEmail sets the "email" field.
func (pc *ProfileCreate) SetEmail(s string) *ProfileCreate {
	pc.mutation.SetEmail(s)
	return pc
}

// SetPhotoURL sets the "photo_url" field.
func (pc *ProfileCreate) SetPhotoURL(s string) *ProfileCreate {
	pc.mutation.SetPhotoURL(s)
	return pc
}

// SetNillablePhotoURL sets the "photo_url" field if the given value is not nil.
func (pc *ProfileCreate) SetNillablePhotoURL(s *string) *ProfileCreate {
	if s != nil {
		pc.SetPhotoURL(*s)
	}
	return pc
}

// SetCreateAt sets the "create_at" field.
func (pc *ProfileCreate) SetCreateAt(t time.Time) *ProfileCreate {
	pc.mutation.SetCreateAt(t)
	return pc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (pc *ProfileCreate) SetNillableCreateAt(t *time.Time) *ProfileCreate {
	if t != nil {
		pc.SetCreateAt(*t)
	}
	return pc
}

// SetUpdateAt sets the "update_at" field.
func (pc *ProfileCreate) SetUpdateAt(t time.Time) *ProfileCreate {
	pc.mutation.SetUpdateAt(t)
	return pc
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (pc *ProfileCreate) SetNillableUpdateAt(t *time.Time) *ProfileCreate {
	if t != nil {
		pc.SetUpdateAt(*t)
	}
	return pc
}

// SetDeleteAt sets the "delete_at" field.
func (pc *ProfileCreate) SetDeleteAt(t time.Time) *ProfileCreate {
	pc.mutation.SetDeleteAt(t)
	return pc
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (pc *ProfileCreate) SetNillableDeleteAt(t *time.Time) *ProfileCreate {
	if t != nil {
		pc.SetDeleteAt(*t)
	}
	return pc
}

// SetID sets the "id" field.
func (pc *ProfileCreate) SetID(s string) *ProfileCreate {
	pc.mutation.SetID(s)
	return pc
}

// AddTaskIDs adds the "Tasks" edge to the Task entity by IDs.
func (pc *ProfileCreate) AddTaskIDs(ids ...int) *ProfileCreate {
	pc.mutation.AddTaskIDs(ids...)
	return pc
}

// AddTasks adds the "Tasks" edges to the Task entity.
func (pc *ProfileCreate) AddTasks(t ...*Task) *ProfileCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return pc.AddTaskIDs(ids...)
}

// Mutation returns the ProfileMutation object of the builder.
func (pc *ProfileCreate) Mutation() *ProfileMutation {
	return pc.mutation
}

// Save creates the Profile in the database.
func (pc *ProfileCreate) Save(ctx context.Context) (*Profile, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProfileCreate) SaveX(ctx context.Context) *Profile {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *ProfileCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *ProfileCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *ProfileCreate) defaults() {
	if _, ok := pc.mutation.CreateAt(); !ok {
		v := profile.DefaultCreateAt()
		pc.mutation.SetCreateAt(v)
	}
	if _, ok := pc.mutation.UpdateAt(); !ok {
		v := profile.DefaultUpdateAt()
		pc.mutation.SetUpdateAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *ProfileCreate) check() error {
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Profile.name"`)}
	}
	if v, ok := pc.mutation.Name(); ok {
		if err := profile.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Profile.name": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "Profile.email"`)}
	}
	if v, ok := pc.mutation.Email(); ok {
		if err := profile.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "Profile.email": %w`, err)}
		}
	}
	if _, ok := pc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "Profile.create_at"`)}
	}
	if _, ok := pc.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "Profile.update_at"`)}
	}
	if v, ok := pc.mutation.ID(); ok {
		if err := profile.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Profile.id": %w`, err)}
		}
	}
	return nil
}

func (pc *ProfileCreate) sqlSave(ctx context.Context) (*Profile, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Profile.ID type: %T", _spec.ID.Value)
		}
	}
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *ProfileCreate) createSpec() (*Profile, *sqlgraph.CreateSpec) {
	var (
		_node = &Profile{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(profile.Table, sqlgraph.NewFieldSpec(profile.FieldID, field.TypeString))
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(profile.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pc.mutation.Email(); ok {
		_spec.SetField(profile.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := pc.mutation.PhotoURL(); ok {
		_spec.SetField(profile.FieldPhotoURL, field.TypeString, value)
		_node.PhotoURL = &value
	}
	if value, ok := pc.mutation.CreateAt(); ok {
		_spec.SetField(profile.FieldCreateAt, field.TypeTime, value)
		_node.CreateAt = value
	}
	if value, ok := pc.mutation.UpdateAt(); ok {
		_spec.SetField(profile.FieldUpdateAt, field.TypeTime, value)
		_node.UpdateAt = value
	}
	if value, ok := pc.mutation.DeleteAt(); ok {
		_spec.SetField(profile.FieldDeleteAt, field.TypeTime, value)
		_node.DeleteAt = value
	}
	if nodes := pc.mutation.TasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   profile.TasksTable,
			Columns: []string{profile.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ProfileCreateBulk is the builder for creating many Profile entities in bulk.
type ProfileCreateBulk struct {
	config
	err      error
	builders []*ProfileCreate
}

// Save creates the Profile entities in the database.
func (pcb *ProfileCreateBulk) Save(ctx context.Context) ([]*Profile, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Profile, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProfileMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *ProfileCreateBulk) SaveX(ctx context.Context) []*Profile {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *ProfileCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *ProfileCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}