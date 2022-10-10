// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/jipark0716/discordTypecast/ent/predicate"
	"github.com/jipark0716/discordTypecast/ent/usertypecastsetting"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeUserTypecastSetting = "UserTypecastSetting"
)

// UserTypecastSettingMutation represents an operation that mutates the UserTypecastSetting nodes in the graph.
type UserTypecastSettingMutation struct {
	config
	op                  Op
	typ                 string
	id                  *int
	actor_id            *string
	text                *string
	lang                *string
	max_seconds         *int
	addmax_seconds      *int
	naturalness         *float64
	addnaturalness      *float64
	speed_x             *int
	addspeed_x          *int
	gid                 *string
	style_idx           *int
	addstyle_idx        *int
	last_pitch          *string
	mode                *string
	pitch               *int
	addpitch            *int
	style_label         *string
	style_label_version *string
	tempo               *int
	addtempo            *int
	clearedFields       map[string]struct{}
	done                bool
	oldValue            func(context.Context) (*UserTypecastSetting, error)
	predicates          []predicate.UserTypecastSetting
}

var _ ent.Mutation = (*UserTypecastSettingMutation)(nil)

// usertypecastsettingOption allows management of the mutation configuration using functional options.
type usertypecastsettingOption func(*UserTypecastSettingMutation)

// newUserTypecastSettingMutation creates new mutation for the UserTypecastSetting entity.
func newUserTypecastSettingMutation(c config, op Op, opts ...usertypecastsettingOption) *UserTypecastSettingMutation {
	m := &UserTypecastSettingMutation{
		config:        c,
		op:            op,
		typ:           TypeUserTypecastSetting,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserTypecastSettingID sets the ID field of the mutation.
func withUserTypecastSettingID(id int) usertypecastsettingOption {
	return func(m *UserTypecastSettingMutation) {
		var (
			err   error
			once  sync.Once
			value *UserTypecastSetting
		)
		m.oldValue = func(ctx context.Context) (*UserTypecastSetting, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().UserTypecastSetting.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUserTypecastSetting sets the old UserTypecastSetting of the mutation.
func withUserTypecastSetting(node *UserTypecastSetting) usertypecastsettingOption {
	return func(m *UserTypecastSettingMutation) {
		m.oldValue = func(context.Context) (*UserTypecastSetting, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserTypecastSettingMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserTypecastSettingMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserTypecastSettingMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserTypecastSettingMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().UserTypecastSetting.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetActorID sets the "actor_id" field.
func (m *UserTypecastSettingMutation) SetActorID(s string) {
	m.actor_id = &s
}

// ActorID returns the value of the "actor_id" field in the mutation.
func (m *UserTypecastSettingMutation) ActorID() (r string, exists bool) {
	v := m.actor_id
	if v == nil {
		return
	}
	return *v, true
}

// OldActorID returns the old "actor_id" field's value of the UserTypecastSetting entity.
// If the UserTypecastSetting object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserTypecastSettingMutation) OldActorID(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldActorID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldActorID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldActorID: %w", err)
	}
	return oldValue.ActorID, nil
}

// ResetActorID resets all changes to the "actor_id" field.
func (m *UserTypecastSettingMutation) ResetActorID() {
	m.actor_id = nil
}

// SetText sets the "text" field.
func (m *UserTypecastSettingMutation) SetText(s string) {
	m.text = &s
}

// Text returns the value of the "text" field in the mutation.
func (m *UserTypecastSettingMutation) Text() (r string, exists bool) {
	v := m.text
	if v == nil {
		return
	}
	return *v, true
}

// OldText returns the old "text" field's value of the UserTypecastSetting entity.
// If the UserTypecastSetting object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserTypecastSettingMutation) OldText(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldText is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldText requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldText: %w", err)
	}
	return oldValue.Text, nil
}

// ResetText resets all changes to the "text" field.
func (m *UserTypecastSettingMutation) ResetText() {
	m.text = nil
}

// SetLang sets the "lang" field.
func (m *UserTypecastSettingMutation) SetLang(s string) {
	m.lang = &s
}

// Lang returns the value of the "lang" field in the mutation.
func (m *UserTypecastSettingMutation) Lang() (r string, exists bool) {
	v := m.lang
	if v == nil {
		return
	}
	return *v, true
}

// OldLang returns the old "lang" field's value of the UserTypecastSetting entity.
// If the UserTypecastSetting object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserTypecastSettingMutation) OldLang(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldLang is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldLang requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLang: %w", err)
	}
	return oldValue.Lang, nil
}

// ResetLang resets all changes to the "lang" field.
func (m *UserTypecastSettingMutation) ResetLang() {
	m.lang = nil
}

// SetMaxSeconds sets the "max_seconds" field.
func (m *UserTypecastSettingMutation) SetMaxSeconds(i int) {
	m.max_seconds = &i
	m.addmax_seconds = nil
}

// MaxSeconds returns the value of the "max_seconds" field in the mutation.
func (m *UserTypecastSettingMutation) MaxSeconds() (r int, exists bool) {
	v := m.max_seconds
	if v == nil {
		return
	}
	return *v, true
}

// OldMaxSeconds returns the old "max_seconds" field's value of the UserTypecastSetting entity.
// If the UserTypecastSetting object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserTypecastSettingMutation) OldMaxSeconds(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldMaxSeconds is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldMaxSeconds requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMaxSeconds: %w", err)
	}
	return oldValue.MaxSeconds, nil
}

// AddMaxSeconds adds i to the "max_seconds" field.
func (m *UserTypecastSettingMutation) AddMaxSeconds(i int) {
	if m.addmax_seconds != nil {
		*m.addmax_seconds += i
	} else {
		m.addmax_seconds = &i
	}
}

// AddedMaxSeconds returns the value that was added to the "max_seconds" field in this mutation.
func (m *UserTypecastSettingMutation) AddedMaxSeconds() (r int, exists bool) {
	v := m.addmax_seconds
	if v == nil {
		return
	}
	return *v, true
}

// ResetMaxSeconds resets all changes to the "max_seconds" field.
func (m *UserTypecastSettingMutation) ResetMaxSeconds() {
	m.max_seconds = nil
	m.addmax_seconds = nil
}

// SetNaturalness sets the "naturalness" field.
func (m *UserTypecastSettingMutation) SetNaturalness(f float64) {
	m.naturalness = &f
	m.addnaturalness = nil
}

// Naturalness returns the value of the "naturalness" field in the mutation.
func (m *UserTypecastSettingMutation) Naturalness() (r float64, exists bool) {
	v := m.naturalness
	if v == nil {
		return
	}
	return *v, true
}

// OldNaturalness returns the old "naturalness" field's value of the UserTypecastSetting entity.
// If the UserTypecastSetting object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserTypecastSettingMutation) OldNaturalness(ctx context.Context) (v float64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldNaturalness is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldNaturalness requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldNaturalness: %w", err)
	}
	return oldValue.Naturalness, nil
}

// AddNaturalness adds f to the "naturalness" field.
func (m *UserTypecastSettingMutation) AddNaturalness(f float64) {
	if m.addnaturalness != nil {
		*m.addnaturalness += f
	} else {
		m.addnaturalness = &f
	}
}

// AddedNaturalness returns the value that was added to the "naturalness" field in this mutation.
func (m *UserTypecastSettingMutation) AddedNaturalness() (r float64, exists bool) {
	v := m.addnaturalness
	if v == nil {
		return
	}
	return *v, true
}

// ResetNaturalness resets all changes to the "naturalness" field.
func (m *UserTypecastSettingMutation) ResetNaturalness() {
	m.naturalness = nil
	m.addnaturalness = nil
}

// SetSpeedX sets the "speed_x" field.
func (m *UserTypecastSettingMutation) SetSpeedX(i int) {
	m.speed_x = &i
	m.addspeed_x = nil
}

// SpeedX returns the value of the "speed_x" field in the mutation.
func (m *UserTypecastSettingMutation) SpeedX() (r int, exists bool) {
	v := m.speed_x
	if v == nil {
		return
	}
	return *v, true
}

// OldSpeedX returns the old "speed_x" field's value of the UserTypecastSetting entity.
// If the UserTypecastSetting object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserTypecastSettingMutation) OldSpeedX(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldSpeedX is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldSpeedX requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSpeedX: %w", err)
	}
	return oldValue.SpeedX, nil
}

// AddSpeedX adds i to the "speed_x" field.
func (m *UserTypecastSettingMutation) AddSpeedX(i int) {
	if m.addspeed_x != nil {
		*m.addspeed_x += i
	} else {
		m.addspeed_x = &i
	}
}

// AddedSpeedX returns the value that was added to the "speed_x" field in this mutation.
func (m *UserTypecastSettingMutation) AddedSpeedX() (r int, exists bool) {
	v := m.addspeed_x
	if v == nil {
		return
	}
	return *v, true
}

// ResetSpeedX resets all changes to the "speed_x" field.
func (m *UserTypecastSettingMutation) ResetSpeedX() {
	m.speed_x = nil
	m.addspeed_x = nil
}

// SetGid sets the "gid" field.
func (m *UserTypecastSettingMutation) SetGid(s string) {
	m.gid = &s
}

// Gid returns the value of the "gid" field in the mutation.
func (m *UserTypecastSettingMutation) Gid() (r string, exists bool) {
	v := m.gid
	if v == nil {
		return
	}
	return *v, true
}

// OldGid returns the old "gid" field's value of the UserTypecastSetting entity.
// If the UserTypecastSetting object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserTypecastSettingMutation) OldGid(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldGid is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldGid requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldGid: %w", err)
	}
	return oldValue.Gid, nil
}

// ResetGid resets all changes to the "gid" field.
func (m *UserTypecastSettingMutation) ResetGid() {
	m.gid = nil
}

// SetStyleIdx sets the "style_idx" field.
func (m *UserTypecastSettingMutation) SetStyleIdx(i int) {
	m.style_idx = &i
	m.addstyle_idx = nil
}

// StyleIdx returns the value of the "style_idx" field in the mutation.
func (m *UserTypecastSettingMutation) StyleIdx() (r int, exists bool) {
	v := m.style_idx
	if v == nil {
		return
	}
	return *v, true
}

// OldStyleIdx returns the old "style_idx" field's value of the UserTypecastSetting entity.
// If the UserTypecastSetting object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserTypecastSettingMutation) OldStyleIdx(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldStyleIdx is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldStyleIdx requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStyleIdx: %w", err)
	}
	return oldValue.StyleIdx, nil
}

// AddStyleIdx adds i to the "style_idx" field.
func (m *UserTypecastSettingMutation) AddStyleIdx(i int) {
	if m.addstyle_idx != nil {
		*m.addstyle_idx += i
	} else {
		m.addstyle_idx = &i
	}
}

// AddedStyleIdx returns the value that was added to the "style_idx" field in this mutation.
func (m *UserTypecastSettingMutation) AddedStyleIdx() (r int, exists bool) {
	v := m.addstyle_idx
	if v == nil {
		return
	}
	return *v, true
}

// ResetStyleIdx resets all changes to the "style_idx" field.
func (m *UserTypecastSettingMutation) ResetStyleIdx() {
	m.style_idx = nil
	m.addstyle_idx = nil
}

// SetLastPitch sets the "last_pitch" field.
func (m *UserTypecastSettingMutation) SetLastPitch(s string) {
	m.last_pitch = &s
}

// LastPitch returns the value of the "last_pitch" field in the mutation.
func (m *UserTypecastSettingMutation) LastPitch() (r string, exists bool) {
	v := m.last_pitch
	if v == nil {
		return
	}
	return *v, true
}

// OldLastPitch returns the old "last_pitch" field's value of the UserTypecastSetting entity.
// If the UserTypecastSetting object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserTypecastSettingMutation) OldLastPitch(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldLastPitch is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldLastPitch requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLastPitch: %w", err)
	}
	return oldValue.LastPitch, nil
}

// ResetLastPitch resets all changes to the "last_pitch" field.
func (m *UserTypecastSettingMutation) ResetLastPitch() {
	m.last_pitch = nil
}

// SetMode sets the "mode" field.
func (m *UserTypecastSettingMutation) SetMode(s string) {
	m.mode = &s
}

// Mode returns the value of the "mode" field in the mutation.
func (m *UserTypecastSettingMutation) Mode() (r string, exists bool) {
	v := m.mode
	if v == nil {
		return
	}
	return *v, true
}

// OldMode returns the old "mode" field's value of the UserTypecastSetting entity.
// If the UserTypecastSetting object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserTypecastSettingMutation) OldMode(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldMode is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldMode requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMode: %w", err)
	}
	return oldValue.Mode, nil
}

// ResetMode resets all changes to the "mode" field.
func (m *UserTypecastSettingMutation) ResetMode() {
	m.mode = nil
}

// SetPitch sets the "pitch" field.
func (m *UserTypecastSettingMutation) SetPitch(i int) {
	m.pitch = &i
	m.addpitch = nil
}

// Pitch returns the value of the "pitch" field in the mutation.
func (m *UserTypecastSettingMutation) Pitch() (r int, exists bool) {
	v := m.pitch
	if v == nil {
		return
	}
	return *v, true
}

// OldPitch returns the old "pitch" field's value of the UserTypecastSetting entity.
// If the UserTypecastSetting object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserTypecastSettingMutation) OldPitch(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPitch is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPitch requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPitch: %w", err)
	}
	return oldValue.Pitch, nil
}

// AddPitch adds i to the "pitch" field.
func (m *UserTypecastSettingMutation) AddPitch(i int) {
	if m.addpitch != nil {
		*m.addpitch += i
	} else {
		m.addpitch = &i
	}
}

// AddedPitch returns the value that was added to the "pitch" field in this mutation.
func (m *UserTypecastSettingMutation) AddedPitch() (r int, exists bool) {
	v := m.addpitch
	if v == nil {
		return
	}
	return *v, true
}

// ResetPitch resets all changes to the "pitch" field.
func (m *UserTypecastSettingMutation) ResetPitch() {
	m.pitch = nil
	m.addpitch = nil
}

// SetStyleLabel sets the "style_label" field.
func (m *UserTypecastSettingMutation) SetStyleLabel(s string) {
	m.style_label = &s
}

// StyleLabel returns the value of the "style_label" field in the mutation.
func (m *UserTypecastSettingMutation) StyleLabel() (r string, exists bool) {
	v := m.style_label
	if v == nil {
		return
	}
	return *v, true
}

// OldStyleLabel returns the old "style_label" field's value of the UserTypecastSetting entity.
// If the UserTypecastSetting object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserTypecastSettingMutation) OldStyleLabel(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldStyleLabel is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldStyleLabel requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStyleLabel: %w", err)
	}
	return oldValue.StyleLabel, nil
}

// ResetStyleLabel resets all changes to the "style_label" field.
func (m *UserTypecastSettingMutation) ResetStyleLabel() {
	m.style_label = nil
}

// SetStyleLabelVersion sets the "style_label_version" field.
func (m *UserTypecastSettingMutation) SetStyleLabelVersion(s string) {
	m.style_label_version = &s
}

// StyleLabelVersion returns the value of the "style_label_version" field in the mutation.
func (m *UserTypecastSettingMutation) StyleLabelVersion() (r string, exists bool) {
	v := m.style_label_version
	if v == nil {
		return
	}
	return *v, true
}

// OldStyleLabelVersion returns the old "style_label_version" field's value of the UserTypecastSetting entity.
// If the UserTypecastSetting object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserTypecastSettingMutation) OldStyleLabelVersion(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldStyleLabelVersion is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldStyleLabelVersion requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStyleLabelVersion: %w", err)
	}
	return oldValue.StyleLabelVersion, nil
}

// ResetStyleLabelVersion resets all changes to the "style_label_version" field.
func (m *UserTypecastSettingMutation) ResetStyleLabelVersion() {
	m.style_label_version = nil
}

// SetTempo sets the "tempo" field.
func (m *UserTypecastSettingMutation) SetTempo(i int) {
	m.tempo = &i
	m.addtempo = nil
}

// Tempo returns the value of the "tempo" field in the mutation.
func (m *UserTypecastSettingMutation) Tempo() (r int, exists bool) {
	v := m.tempo
	if v == nil {
		return
	}
	return *v, true
}

// OldTempo returns the old "tempo" field's value of the UserTypecastSetting entity.
// If the UserTypecastSetting object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserTypecastSettingMutation) OldTempo(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTempo is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTempo requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTempo: %w", err)
	}
	return oldValue.Tempo, nil
}

// AddTempo adds i to the "tempo" field.
func (m *UserTypecastSettingMutation) AddTempo(i int) {
	if m.addtempo != nil {
		*m.addtempo += i
	} else {
		m.addtempo = &i
	}
}

// AddedTempo returns the value that was added to the "tempo" field in this mutation.
func (m *UserTypecastSettingMutation) AddedTempo() (r int, exists bool) {
	v := m.addtempo
	if v == nil {
		return
	}
	return *v, true
}

// ResetTempo resets all changes to the "tempo" field.
func (m *UserTypecastSettingMutation) ResetTempo() {
	m.tempo = nil
	m.addtempo = nil
}

// Where appends a list predicates to the UserTypecastSettingMutation builder.
func (m *UserTypecastSettingMutation) Where(ps ...predicate.UserTypecastSetting) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *UserTypecastSettingMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (UserTypecastSetting).
func (m *UserTypecastSettingMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserTypecastSettingMutation) Fields() []string {
	fields := make([]string, 0, 14)
	if m.actor_id != nil {
		fields = append(fields, usertypecastsetting.FieldActorID)
	}
	if m.text != nil {
		fields = append(fields, usertypecastsetting.FieldText)
	}
	if m.lang != nil {
		fields = append(fields, usertypecastsetting.FieldLang)
	}
	if m.max_seconds != nil {
		fields = append(fields, usertypecastsetting.FieldMaxSeconds)
	}
	if m.naturalness != nil {
		fields = append(fields, usertypecastsetting.FieldNaturalness)
	}
	if m.speed_x != nil {
		fields = append(fields, usertypecastsetting.FieldSpeedX)
	}
	if m.gid != nil {
		fields = append(fields, usertypecastsetting.FieldGid)
	}
	if m.style_idx != nil {
		fields = append(fields, usertypecastsetting.FieldStyleIdx)
	}
	if m.last_pitch != nil {
		fields = append(fields, usertypecastsetting.FieldLastPitch)
	}
	if m.mode != nil {
		fields = append(fields, usertypecastsetting.FieldMode)
	}
	if m.pitch != nil {
		fields = append(fields, usertypecastsetting.FieldPitch)
	}
	if m.style_label != nil {
		fields = append(fields, usertypecastsetting.FieldStyleLabel)
	}
	if m.style_label_version != nil {
		fields = append(fields, usertypecastsetting.FieldStyleLabelVersion)
	}
	if m.tempo != nil {
		fields = append(fields, usertypecastsetting.FieldTempo)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserTypecastSettingMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case usertypecastsetting.FieldActorID:
		return m.ActorID()
	case usertypecastsetting.FieldText:
		return m.Text()
	case usertypecastsetting.FieldLang:
		return m.Lang()
	case usertypecastsetting.FieldMaxSeconds:
		return m.MaxSeconds()
	case usertypecastsetting.FieldNaturalness:
		return m.Naturalness()
	case usertypecastsetting.FieldSpeedX:
		return m.SpeedX()
	case usertypecastsetting.FieldGid:
		return m.Gid()
	case usertypecastsetting.FieldStyleIdx:
		return m.StyleIdx()
	case usertypecastsetting.FieldLastPitch:
		return m.LastPitch()
	case usertypecastsetting.FieldMode:
		return m.Mode()
	case usertypecastsetting.FieldPitch:
		return m.Pitch()
	case usertypecastsetting.FieldStyleLabel:
		return m.StyleLabel()
	case usertypecastsetting.FieldStyleLabelVersion:
		return m.StyleLabelVersion()
	case usertypecastsetting.FieldTempo:
		return m.Tempo()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserTypecastSettingMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case usertypecastsetting.FieldActorID:
		return m.OldActorID(ctx)
	case usertypecastsetting.FieldText:
		return m.OldText(ctx)
	case usertypecastsetting.FieldLang:
		return m.OldLang(ctx)
	case usertypecastsetting.FieldMaxSeconds:
		return m.OldMaxSeconds(ctx)
	case usertypecastsetting.FieldNaturalness:
		return m.OldNaturalness(ctx)
	case usertypecastsetting.FieldSpeedX:
		return m.OldSpeedX(ctx)
	case usertypecastsetting.FieldGid:
		return m.OldGid(ctx)
	case usertypecastsetting.FieldStyleIdx:
		return m.OldStyleIdx(ctx)
	case usertypecastsetting.FieldLastPitch:
		return m.OldLastPitch(ctx)
	case usertypecastsetting.FieldMode:
		return m.OldMode(ctx)
	case usertypecastsetting.FieldPitch:
		return m.OldPitch(ctx)
	case usertypecastsetting.FieldStyleLabel:
		return m.OldStyleLabel(ctx)
	case usertypecastsetting.FieldStyleLabelVersion:
		return m.OldStyleLabelVersion(ctx)
	case usertypecastsetting.FieldTempo:
		return m.OldTempo(ctx)
	}
	return nil, fmt.Errorf("unknown UserTypecastSetting field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserTypecastSettingMutation) SetField(name string, value ent.Value) error {
	switch name {
	case usertypecastsetting.FieldActorID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetActorID(v)
		return nil
	case usertypecastsetting.FieldText:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetText(v)
		return nil
	case usertypecastsetting.FieldLang:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLang(v)
		return nil
	case usertypecastsetting.FieldMaxSeconds:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMaxSeconds(v)
		return nil
	case usertypecastsetting.FieldNaturalness:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetNaturalness(v)
		return nil
	case usertypecastsetting.FieldSpeedX:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSpeedX(v)
		return nil
	case usertypecastsetting.FieldGid:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetGid(v)
		return nil
	case usertypecastsetting.FieldStyleIdx:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStyleIdx(v)
		return nil
	case usertypecastsetting.FieldLastPitch:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLastPitch(v)
		return nil
	case usertypecastsetting.FieldMode:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMode(v)
		return nil
	case usertypecastsetting.FieldPitch:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPitch(v)
		return nil
	case usertypecastsetting.FieldStyleLabel:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStyleLabel(v)
		return nil
	case usertypecastsetting.FieldStyleLabelVersion:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStyleLabelVersion(v)
		return nil
	case usertypecastsetting.FieldTempo:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTempo(v)
		return nil
	}
	return fmt.Errorf("unknown UserTypecastSetting field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserTypecastSettingMutation) AddedFields() []string {
	var fields []string
	if m.addmax_seconds != nil {
		fields = append(fields, usertypecastsetting.FieldMaxSeconds)
	}
	if m.addnaturalness != nil {
		fields = append(fields, usertypecastsetting.FieldNaturalness)
	}
	if m.addspeed_x != nil {
		fields = append(fields, usertypecastsetting.FieldSpeedX)
	}
	if m.addstyle_idx != nil {
		fields = append(fields, usertypecastsetting.FieldStyleIdx)
	}
	if m.addpitch != nil {
		fields = append(fields, usertypecastsetting.FieldPitch)
	}
	if m.addtempo != nil {
		fields = append(fields, usertypecastsetting.FieldTempo)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserTypecastSettingMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case usertypecastsetting.FieldMaxSeconds:
		return m.AddedMaxSeconds()
	case usertypecastsetting.FieldNaturalness:
		return m.AddedNaturalness()
	case usertypecastsetting.FieldSpeedX:
		return m.AddedSpeedX()
	case usertypecastsetting.FieldStyleIdx:
		return m.AddedStyleIdx()
	case usertypecastsetting.FieldPitch:
		return m.AddedPitch()
	case usertypecastsetting.FieldTempo:
		return m.AddedTempo()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserTypecastSettingMutation) AddField(name string, value ent.Value) error {
	switch name {
	case usertypecastsetting.FieldMaxSeconds:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddMaxSeconds(v)
		return nil
	case usertypecastsetting.FieldNaturalness:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddNaturalness(v)
		return nil
	case usertypecastsetting.FieldSpeedX:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddSpeedX(v)
		return nil
	case usertypecastsetting.FieldStyleIdx:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddStyleIdx(v)
		return nil
	case usertypecastsetting.FieldPitch:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddPitch(v)
		return nil
	case usertypecastsetting.FieldTempo:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddTempo(v)
		return nil
	}
	return fmt.Errorf("unknown UserTypecastSetting numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserTypecastSettingMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserTypecastSettingMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserTypecastSettingMutation) ClearField(name string) error {
	return fmt.Errorf("unknown UserTypecastSetting nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserTypecastSettingMutation) ResetField(name string) error {
	switch name {
	case usertypecastsetting.FieldActorID:
		m.ResetActorID()
		return nil
	case usertypecastsetting.FieldText:
		m.ResetText()
		return nil
	case usertypecastsetting.FieldLang:
		m.ResetLang()
		return nil
	case usertypecastsetting.FieldMaxSeconds:
		m.ResetMaxSeconds()
		return nil
	case usertypecastsetting.FieldNaturalness:
		m.ResetNaturalness()
		return nil
	case usertypecastsetting.FieldSpeedX:
		m.ResetSpeedX()
		return nil
	case usertypecastsetting.FieldGid:
		m.ResetGid()
		return nil
	case usertypecastsetting.FieldStyleIdx:
		m.ResetStyleIdx()
		return nil
	case usertypecastsetting.FieldLastPitch:
		m.ResetLastPitch()
		return nil
	case usertypecastsetting.FieldMode:
		m.ResetMode()
		return nil
	case usertypecastsetting.FieldPitch:
		m.ResetPitch()
		return nil
	case usertypecastsetting.FieldStyleLabel:
		m.ResetStyleLabel()
		return nil
	case usertypecastsetting.FieldStyleLabelVersion:
		m.ResetStyleLabelVersion()
		return nil
	case usertypecastsetting.FieldTempo:
		m.ResetTempo()
		return nil
	}
	return fmt.Errorf("unknown UserTypecastSetting field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserTypecastSettingMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserTypecastSettingMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserTypecastSettingMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserTypecastSettingMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserTypecastSettingMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserTypecastSettingMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserTypecastSettingMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown UserTypecastSetting unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserTypecastSettingMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown UserTypecastSetting edge %s", name)
}