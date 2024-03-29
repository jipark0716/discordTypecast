// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/jipark0716/discordTypecast/ent/predicate"
	"github.com/jipark0716/discordTypecast/ent/typecastmessage"
)

// TypecastMessageUpdate is the builder for updating TypecastMessage entities.
type TypecastMessageUpdate struct {
	config
	hooks    []Hook
	mutation *TypecastMessageMutation
}

// Where appends a list predicates to the TypecastMessageUpdate builder.
func (tmu *TypecastMessageUpdate) Where(ps ...predicate.TypecastMessage) *TypecastMessageUpdate {
	tmu.mutation.Where(ps...)
	return tmu
}

// SetUserID sets the "user_id" field.
func (tmu *TypecastMessageUpdate) SetUserID(s string) *TypecastMessageUpdate {
	tmu.mutation.SetUserID(s)
	return tmu
}

// SetText sets the "text" field.
func (tmu *TypecastMessageUpdate) SetText(s string) *TypecastMessageUpdate {
	tmu.mutation.SetText(s)
	return tmu
}

// SetGuildID sets the "guild_id" field.
func (tmu *TypecastMessageUpdate) SetGuildID(s string) *TypecastMessageUpdate {
	tmu.mutation.SetGuildID(s)
	return tmu
}

// SetChannelID sets the "channel_id" field.
func (tmu *TypecastMessageUpdate) SetChannelID(s string) *TypecastMessageUpdate {
	tmu.mutation.SetChannelID(s)
	return tmu
}

// SetActorID sets the "actor_id" field.
func (tmu *TypecastMessageUpdate) SetActorID(s string) *TypecastMessageUpdate {
	tmu.mutation.SetActorID(s)
	return tmu
}

// SetLang sets the "lang" field.
func (tmu *TypecastMessageUpdate) SetLang(s string) *TypecastMessageUpdate {
	tmu.mutation.SetLang(s)
	return tmu
}

// SetMaxSeconds sets the "max_seconds" field.
func (tmu *TypecastMessageUpdate) SetMaxSeconds(i int) *TypecastMessageUpdate {
	tmu.mutation.ResetMaxSeconds()
	tmu.mutation.SetMaxSeconds(i)
	return tmu
}

// AddMaxSeconds adds i to the "max_seconds" field.
func (tmu *TypecastMessageUpdate) AddMaxSeconds(i int) *TypecastMessageUpdate {
	tmu.mutation.AddMaxSeconds(i)
	return tmu
}

// SetNaturalness sets the "naturalness" field.
func (tmu *TypecastMessageUpdate) SetNaturalness(f float64) *TypecastMessageUpdate {
	tmu.mutation.ResetNaturalness()
	tmu.mutation.SetNaturalness(f)
	return tmu
}

// AddNaturalness adds f to the "naturalness" field.
func (tmu *TypecastMessageUpdate) AddNaturalness(f float64) *TypecastMessageUpdate {
	tmu.mutation.AddNaturalness(f)
	return tmu
}

// SetSpeedX sets the "speed_x" field.
func (tmu *TypecastMessageUpdate) SetSpeedX(i int) *TypecastMessageUpdate {
	tmu.mutation.ResetSpeedX()
	tmu.mutation.SetSpeedX(i)
	return tmu
}

// AddSpeedX adds i to the "speed_x" field.
func (tmu *TypecastMessageUpdate) AddSpeedX(i int) *TypecastMessageUpdate {
	tmu.mutation.AddSpeedX(i)
	return tmu
}

// SetGid sets the "gid" field.
func (tmu *TypecastMessageUpdate) SetGid(s string) *TypecastMessageUpdate {
	tmu.mutation.SetGid(s)
	return tmu
}

// SetStyleIdx sets the "style_idx" field.
func (tmu *TypecastMessageUpdate) SetStyleIdx(i int) *TypecastMessageUpdate {
	tmu.mutation.ResetStyleIdx()
	tmu.mutation.SetStyleIdx(i)
	return tmu
}

// AddStyleIdx adds i to the "style_idx" field.
func (tmu *TypecastMessageUpdate) AddStyleIdx(i int) *TypecastMessageUpdate {
	tmu.mutation.AddStyleIdx(i)
	return tmu
}

// SetLastPitch sets the "last_pitch" field.
func (tmu *TypecastMessageUpdate) SetLastPitch(s string) *TypecastMessageUpdate {
	tmu.mutation.SetLastPitch(s)
	return tmu
}

// SetNillableLastPitch sets the "last_pitch" field if the given value is not nil.
func (tmu *TypecastMessageUpdate) SetNillableLastPitch(s *string) *TypecastMessageUpdate {
	if s != nil {
		tmu.SetLastPitch(*s)
	}
	return tmu
}

// ClearLastPitch clears the value of the "last_pitch" field.
func (tmu *TypecastMessageUpdate) ClearLastPitch() *TypecastMessageUpdate {
	tmu.mutation.ClearLastPitch()
	return tmu
}

// SetMode sets the "mode" field.
func (tmu *TypecastMessageUpdate) SetMode(s string) *TypecastMessageUpdate {
	tmu.mutation.SetMode(s)
	return tmu
}

// SetNillableMode sets the "mode" field if the given value is not nil.
func (tmu *TypecastMessageUpdate) SetNillableMode(s *string) *TypecastMessageUpdate {
	if s != nil {
		tmu.SetMode(*s)
	}
	return tmu
}

// ClearMode clears the value of the "mode" field.
func (tmu *TypecastMessageUpdate) ClearMode() *TypecastMessageUpdate {
	tmu.mutation.ClearMode()
	return tmu
}

// SetPitch sets the "pitch" field.
func (tmu *TypecastMessageUpdate) SetPitch(i int) *TypecastMessageUpdate {
	tmu.mutation.ResetPitch()
	tmu.mutation.SetPitch(i)
	return tmu
}

// AddPitch adds i to the "pitch" field.
func (tmu *TypecastMessageUpdate) AddPitch(i int) *TypecastMessageUpdate {
	tmu.mutation.AddPitch(i)
	return tmu
}

// SetStyleLabel sets the "style_label" field.
func (tmu *TypecastMessageUpdate) SetStyleLabel(s string) *TypecastMessageUpdate {
	tmu.mutation.SetStyleLabel(s)
	return tmu
}

// SetStyleLabelVersion sets the "style_label_version" field.
func (tmu *TypecastMessageUpdate) SetStyleLabelVersion(s string) *TypecastMessageUpdate {
	tmu.mutation.SetStyleLabelVersion(s)
	return tmu
}

// SetTempo sets the "tempo" field.
func (tmu *TypecastMessageUpdate) SetTempo(i int) *TypecastMessageUpdate {
	tmu.mutation.ResetTempo()
	tmu.mutation.SetTempo(i)
	return tmu
}

// AddTempo adds i to the "tempo" field.
func (tmu *TypecastMessageUpdate) AddTempo(i int) *TypecastMessageUpdate {
	tmu.mutation.AddTempo(i)
	return tmu
}

// SetStatus sets the "status" field.
func (tmu *TypecastMessageUpdate) SetStatus(i int8) *TypecastMessageUpdate {
	tmu.mutation.ResetStatus()
	tmu.mutation.SetStatus(i)
	return tmu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tmu *TypecastMessageUpdate) SetNillableStatus(i *int8) *TypecastMessageUpdate {
	if i != nil {
		tmu.SetStatus(*i)
	}
	return tmu
}

// AddStatus adds i to the "status" field.
func (tmu *TypecastMessageUpdate) AddStatus(i int8) *TypecastMessageUpdate {
	tmu.mutation.AddStatus(i)
	return tmu
}

// SetSendAt sets the "send_at" field.
func (tmu *TypecastMessageUpdate) SetSendAt(t time.Time) *TypecastMessageUpdate {
	tmu.mutation.SetSendAt(t)
	return tmu
}

// SetNillableSendAt sets the "send_at" field if the given value is not nil.
func (tmu *TypecastMessageUpdate) SetNillableSendAt(t *time.Time) *TypecastMessageUpdate {
	if t != nil {
		tmu.SetSendAt(*t)
	}
	return tmu
}

// ClearSendAt clears the value of the "send_at" field.
func (tmu *TypecastMessageUpdate) ClearSendAt() *TypecastMessageUpdate {
	tmu.mutation.ClearSendAt()
	return tmu
}

// Mutation returns the TypecastMessageMutation object of the builder.
func (tmu *TypecastMessageUpdate) Mutation() *TypecastMessageMutation {
	return tmu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tmu *TypecastMessageUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tmu.hooks) == 0 {
		affected, err = tmu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TypecastMessageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tmu.mutation = mutation
			affected, err = tmu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tmu.hooks) - 1; i >= 0; i-- {
			if tmu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tmu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tmu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tmu *TypecastMessageUpdate) SaveX(ctx context.Context) int {
	affected, err := tmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tmu *TypecastMessageUpdate) Exec(ctx context.Context) error {
	_, err := tmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tmu *TypecastMessageUpdate) ExecX(ctx context.Context) {
	if err := tmu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tmu *TypecastMessageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   typecastmessage.Table,
			Columns: typecastmessage.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: typecastmessage.FieldID,
			},
		},
	}
	if ps := tmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tmu.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: typecastmessage.FieldUserID,
		})
	}
	if value, ok := tmu.mutation.Text(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: typecastmessage.FieldText,
		})
	}
	if value, ok := tmu.mutation.GuildID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: typecastmessage.FieldGuildID,
		})
	}
	if value, ok := tmu.mutation.ChannelID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: typecastmessage.FieldChannelID,
		})
	}
	if value, ok := tmu.mutation.ActorID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: typecastmessage.FieldActorID,
		})
	}
	if value, ok := tmu.mutation.Lang(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: typecastmessage.FieldLang,
		})
	}
	if value, ok := tmu.mutation.MaxSeconds(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: typecastmessage.FieldMaxSeconds,
		})
	}
	if value, ok := tmu.mutation.AddedMaxSeconds(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: typecastmessage.FieldMaxSeconds,
		})
	}
	if value, ok := tmu.mutation.Naturalness(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: typecastmessage.FieldNaturalness,
		})
	}
	if value, ok := tmu.mutation.AddedNaturalness(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: typecastmessage.FieldNaturalness,
		})
	}
	if value, ok := tmu.mutation.SpeedX(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: typecastmessage.FieldSpeedX,
		})
	}
	if value, ok := tmu.mutation.AddedSpeedX(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: typecastmessage.FieldSpeedX,
		})
	}
	if value, ok := tmu.mutation.Gid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: typecastmessage.FieldGid,
		})
	}
	if value, ok := tmu.mutation.StyleIdx(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: typecastmessage.FieldStyleIdx,
		})
	}
	if value, ok := tmu.mutation.AddedStyleIdx(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: typecastmessage.FieldStyleIdx,
		})
	}
	if value, ok := tmu.mutation.LastPitch(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: typecastmessage.FieldLastPitch,
		})
	}
	if tmu.mutation.LastPitchCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: typecastmessage.FieldLastPitch,
		})
	}
	if value, ok := tmu.mutation.Mode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: typecastmessage.FieldMode,
		})
	}
	if tmu.mutation.ModeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: typecastmessage.FieldMode,
		})
	}
	if value, ok := tmu.mutation.Pitch(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: typecastmessage.FieldPitch,
		})
	}
	if value, ok := tmu.mutation.AddedPitch(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: typecastmessage.FieldPitch,
		})
	}
	if value, ok := tmu.mutation.StyleLabel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: typecastmessage.FieldStyleLabel,
		})
	}
	if value, ok := tmu.mutation.StyleLabelVersion(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: typecastmessage.FieldStyleLabelVersion,
		})
	}
	if value, ok := tmu.mutation.Tempo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: typecastmessage.FieldTempo,
		})
	}
	if value, ok := tmu.mutation.AddedTempo(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: typecastmessage.FieldTempo,
		})
	}
	if value, ok := tmu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: typecastmessage.FieldStatus,
		})
	}
	if value, ok := tmu.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: typecastmessage.FieldStatus,
		})
	}
	if value, ok := tmu.mutation.SendAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: typecastmessage.FieldSendAt,
		})
	}
	if tmu.mutation.SendAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: typecastmessage.FieldSendAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{typecastmessage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// TypecastMessageUpdateOne is the builder for updating a single TypecastMessage entity.
type TypecastMessageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TypecastMessageMutation
}

// SetUserID sets the "user_id" field.
func (tmuo *TypecastMessageUpdateOne) SetUserID(s string) *TypecastMessageUpdateOne {
	tmuo.mutation.SetUserID(s)
	return tmuo
}

// SetText sets the "text" field.
func (tmuo *TypecastMessageUpdateOne) SetText(s string) *TypecastMessageUpdateOne {
	tmuo.mutation.SetText(s)
	return tmuo
}

// SetGuildID sets the "guild_id" field.
func (tmuo *TypecastMessageUpdateOne) SetGuildID(s string) *TypecastMessageUpdateOne {
	tmuo.mutation.SetGuildID(s)
	return tmuo
}

// SetChannelID sets the "channel_id" field.
func (tmuo *TypecastMessageUpdateOne) SetChannelID(s string) *TypecastMessageUpdateOne {
	tmuo.mutation.SetChannelID(s)
	return tmuo
}

// SetActorID sets the "actor_id" field.
func (tmuo *TypecastMessageUpdateOne) SetActorID(s string) *TypecastMessageUpdateOne {
	tmuo.mutation.SetActorID(s)
	return tmuo
}

// SetLang sets the "lang" field.
func (tmuo *TypecastMessageUpdateOne) SetLang(s string) *TypecastMessageUpdateOne {
	tmuo.mutation.SetLang(s)
	return tmuo
}

// SetMaxSeconds sets the "max_seconds" field.
func (tmuo *TypecastMessageUpdateOne) SetMaxSeconds(i int) *TypecastMessageUpdateOne {
	tmuo.mutation.ResetMaxSeconds()
	tmuo.mutation.SetMaxSeconds(i)
	return tmuo
}

// AddMaxSeconds adds i to the "max_seconds" field.
func (tmuo *TypecastMessageUpdateOne) AddMaxSeconds(i int) *TypecastMessageUpdateOne {
	tmuo.mutation.AddMaxSeconds(i)
	return tmuo
}

// SetNaturalness sets the "naturalness" field.
func (tmuo *TypecastMessageUpdateOne) SetNaturalness(f float64) *TypecastMessageUpdateOne {
	tmuo.mutation.ResetNaturalness()
	tmuo.mutation.SetNaturalness(f)
	return tmuo
}

// AddNaturalness adds f to the "naturalness" field.
func (tmuo *TypecastMessageUpdateOne) AddNaturalness(f float64) *TypecastMessageUpdateOne {
	tmuo.mutation.AddNaturalness(f)
	return tmuo
}

// SetSpeedX sets the "speed_x" field.
func (tmuo *TypecastMessageUpdateOne) SetSpeedX(i int) *TypecastMessageUpdateOne {
	tmuo.mutation.ResetSpeedX()
	tmuo.mutation.SetSpeedX(i)
	return tmuo
}

// AddSpeedX adds i to the "speed_x" field.
func (tmuo *TypecastMessageUpdateOne) AddSpeedX(i int) *TypecastMessageUpdateOne {
	tmuo.mutation.AddSpeedX(i)
	return tmuo
}

// SetGid sets the "gid" field.
func (tmuo *TypecastMessageUpdateOne) SetGid(s string) *TypecastMessageUpdateOne {
	tmuo.mutation.SetGid(s)
	return tmuo
}

// SetStyleIdx sets the "style_idx" field.
func (tmuo *TypecastMessageUpdateOne) SetStyleIdx(i int) *TypecastMessageUpdateOne {
	tmuo.mutation.ResetStyleIdx()
	tmuo.mutation.SetStyleIdx(i)
	return tmuo
}

// AddStyleIdx adds i to the "style_idx" field.
func (tmuo *TypecastMessageUpdateOne) AddStyleIdx(i int) *TypecastMessageUpdateOne {
	tmuo.mutation.AddStyleIdx(i)
	return tmuo
}

// SetLastPitch sets the "last_pitch" field.
func (tmuo *TypecastMessageUpdateOne) SetLastPitch(s string) *TypecastMessageUpdateOne {
	tmuo.mutation.SetLastPitch(s)
	return tmuo
}

// SetNillableLastPitch sets the "last_pitch" field if the given value is not nil.
func (tmuo *TypecastMessageUpdateOne) SetNillableLastPitch(s *string) *TypecastMessageUpdateOne {
	if s != nil {
		tmuo.SetLastPitch(*s)
	}
	return tmuo
}

// ClearLastPitch clears the value of the "last_pitch" field.
func (tmuo *TypecastMessageUpdateOne) ClearLastPitch() *TypecastMessageUpdateOne {
	tmuo.mutation.ClearLastPitch()
	return tmuo
}

// SetMode sets the "mode" field.
func (tmuo *TypecastMessageUpdateOne) SetMode(s string) *TypecastMessageUpdateOne {
	tmuo.mutation.SetMode(s)
	return tmuo
}

// SetNillableMode sets the "mode" field if the given value is not nil.
func (tmuo *TypecastMessageUpdateOne) SetNillableMode(s *string) *TypecastMessageUpdateOne {
	if s != nil {
		tmuo.SetMode(*s)
	}
	return tmuo
}

// ClearMode clears the value of the "mode" field.
func (tmuo *TypecastMessageUpdateOne) ClearMode() *TypecastMessageUpdateOne {
	tmuo.mutation.ClearMode()
	return tmuo
}

// SetPitch sets the "pitch" field.
func (tmuo *TypecastMessageUpdateOne) SetPitch(i int) *TypecastMessageUpdateOne {
	tmuo.mutation.ResetPitch()
	tmuo.mutation.SetPitch(i)
	return tmuo
}

// AddPitch adds i to the "pitch" field.
func (tmuo *TypecastMessageUpdateOne) AddPitch(i int) *TypecastMessageUpdateOne {
	tmuo.mutation.AddPitch(i)
	return tmuo
}

// SetStyleLabel sets the "style_label" field.
func (tmuo *TypecastMessageUpdateOne) SetStyleLabel(s string) *TypecastMessageUpdateOne {
	tmuo.mutation.SetStyleLabel(s)
	return tmuo
}

// SetStyleLabelVersion sets the "style_label_version" field.
func (tmuo *TypecastMessageUpdateOne) SetStyleLabelVersion(s string) *TypecastMessageUpdateOne {
	tmuo.mutation.SetStyleLabelVersion(s)
	return tmuo
}

// SetTempo sets the "tempo" field.
func (tmuo *TypecastMessageUpdateOne) SetTempo(i int) *TypecastMessageUpdateOne {
	tmuo.mutation.ResetTempo()
	tmuo.mutation.SetTempo(i)
	return tmuo
}

// AddTempo adds i to the "tempo" field.
func (tmuo *TypecastMessageUpdateOne) AddTempo(i int) *TypecastMessageUpdateOne {
	tmuo.mutation.AddTempo(i)
	return tmuo
}

// SetStatus sets the "status" field.
func (tmuo *TypecastMessageUpdateOne) SetStatus(i int8) *TypecastMessageUpdateOne {
	tmuo.mutation.ResetStatus()
	tmuo.mutation.SetStatus(i)
	return tmuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tmuo *TypecastMessageUpdateOne) SetNillableStatus(i *int8) *TypecastMessageUpdateOne {
	if i != nil {
		tmuo.SetStatus(*i)
	}
	return tmuo
}

// AddStatus adds i to the "status" field.
func (tmuo *TypecastMessageUpdateOne) AddStatus(i int8) *TypecastMessageUpdateOne {
	tmuo.mutation.AddStatus(i)
	return tmuo
}

// SetSendAt sets the "send_at" field.
func (tmuo *TypecastMessageUpdateOne) SetSendAt(t time.Time) *TypecastMessageUpdateOne {
	tmuo.mutation.SetSendAt(t)
	return tmuo
}

// SetNillableSendAt sets the "send_at" field if the given value is not nil.
func (tmuo *TypecastMessageUpdateOne) SetNillableSendAt(t *time.Time) *TypecastMessageUpdateOne {
	if t != nil {
		tmuo.SetSendAt(*t)
	}
	return tmuo
}

// ClearSendAt clears the value of the "send_at" field.
func (tmuo *TypecastMessageUpdateOne) ClearSendAt() *TypecastMessageUpdateOne {
	tmuo.mutation.ClearSendAt()
	return tmuo
}

// Mutation returns the TypecastMessageMutation object of the builder.
func (tmuo *TypecastMessageUpdateOne) Mutation() *TypecastMessageMutation {
	return tmuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tmuo *TypecastMessageUpdateOne) Select(field string, fields ...string) *TypecastMessageUpdateOne {
	tmuo.fields = append([]string{field}, fields...)
	return tmuo
}

// Save executes the query and returns the updated TypecastMessage entity.
func (tmuo *TypecastMessageUpdateOne) Save(ctx context.Context) (*TypecastMessage, error) {
	var (
		err  error
		node *TypecastMessage
	)
	if len(tmuo.hooks) == 0 {
		node, err = tmuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TypecastMessageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tmuo.mutation = mutation
			node, err = tmuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tmuo.hooks) - 1; i >= 0; i-- {
			if tmuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tmuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tmuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*TypecastMessage)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TypecastMessageMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tmuo *TypecastMessageUpdateOne) SaveX(ctx context.Context) *TypecastMessage {
	node, err := tmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tmuo *TypecastMessageUpdateOne) Exec(ctx context.Context) error {
	_, err := tmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tmuo *TypecastMessageUpdateOne) ExecX(ctx context.Context) {
	if err := tmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tmuo *TypecastMessageUpdateOne) sqlSave(ctx context.Context) (_node *TypecastMessage, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   typecastmessage.Table,
			Columns: typecastmessage.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: typecastmessage.FieldID,
			},
		},
	}
	id, ok := tmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TypecastMessage.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, typecastmessage.FieldID)
		for _, f := range fields {
			if !typecastmessage.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != typecastmessage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tmuo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: typecastmessage.FieldUserID,
		})
	}
	if value, ok := tmuo.mutation.Text(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: typecastmessage.FieldText,
		})
	}
	if value, ok := tmuo.mutation.GuildID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: typecastmessage.FieldGuildID,
		})
	}
	if value, ok := tmuo.mutation.ChannelID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: typecastmessage.FieldChannelID,
		})
	}
	if value, ok := tmuo.mutation.ActorID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: typecastmessage.FieldActorID,
		})
	}
	if value, ok := tmuo.mutation.Lang(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: typecastmessage.FieldLang,
		})
	}
	if value, ok := tmuo.mutation.MaxSeconds(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: typecastmessage.FieldMaxSeconds,
		})
	}
	if value, ok := tmuo.mutation.AddedMaxSeconds(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: typecastmessage.FieldMaxSeconds,
		})
	}
	if value, ok := tmuo.mutation.Naturalness(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: typecastmessage.FieldNaturalness,
		})
	}
	if value, ok := tmuo.mutation.AddedNaturalness(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: typecastmessage.FieldNaturalness,
		})
	}
	if value, ok := tmuo.mutation.SpeedX(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: typecastmessage.FieldSpeedX,
		})
	}
	if value, ok := tmuo.mutation.AddedSpeedX(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: typecastmessage.FieldSpeedX,
		})
	}
	if value, ok := tmuo.mutation.Gid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: typecastmessage.FieldGid,
		})
	}
	if value, ok := tmuo.mutation.StyleIdx(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: typecastmessage.FieldStyleIdx,
		})
	}
	if value, ok := tmuo.mutation.AddedStyleIdx(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: typecastmessage.FieldStyleIdx,
		})
	}
	if value, ok := tmuo.mutation.LastPitch(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: typecastmessage.FieldLastPitch,
		})
	}
	if tmuo.mutation.LastPitchCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: typecastmessage.FieldLastPitch,
		})
	}
	if value, ok := tmuo.mutation.Mode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: typecastmessage.FieldMode,
		})
	}
	if tmuo.mutation.ModeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: typecastmessage.FieldMode,
		})
	}
	if value, ok := tmuo.mutation.Pitch(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: typecastmessage.FieldPitch,
		})
	}
	if value, ok := tmuo.mutation.AddedPitch(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: typecastmessage.FieldPitch,
		})
	}
	if value, ok := tmuo.mutation.StyleLabel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: typecastmessage.FieldStyleLabel,
		})
	}
	if value, ok := tmuo.mutation.StyleLabelVersion(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: typecastmessage.FieldStyleLabelVersion,
		})
	}
	if value, ok := tmuo.mutation.Tempo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: typecastmessage.FieldTempo,
		})
	}
	if value, ok := tmuo.mutation.AddedTempo(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: typecastmessage.FieldTempo,
		})
	}
	if value, ok := tmuo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: typecastmessage.FieldStatus,
		})
	}
	if value, ok := tmuo.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: typecastmessage.FieldStatus,
		})
	}
	if value, ok := tmuo.mutation.SendAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: typecastmessage.FieldSendAt,
		})
	}
	if tmuo.mutation.SendAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: typecastmessage.FieldSendAt,
		})
	}
	_node = &TypecastMessage{config: tmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{typecastmessage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
