// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/jipark0716/discordTypecast/ent/usertypecastsetting"
)

// UserTypecastSetting is the model entity for the UserTypecastSetting schema.
type UserTypecastSetting struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID string `json:"user_id,omitempty"`
	// ActorID holds the value of the "actor_id" field.
	ActorID string `json:"actor_id,omitempty"`
	// Lang holds the value of the "lang" field.
	Lang string `json:"lang,omitempty"`
	// MaxSeconds holds the value of the "max_seconds" field.
	MaxSeconds int `json:"max_seconds,omitempty"`
	// Naturalness holds the value of the "naturalness" field.
	Naturalness float64 `json:"naturalness,omitempty"`
	// SpeedX holds the value of the "speed_x" field.
	SpeedX int `json:"speed_x,omitempty"`
	// Gid holds the value of the "gid" field.
	Gid string `json:"gid,omitempty"`
	// StyleIdx holds the value of the "style_idx" field.
	StyleIdx int `json:"style_idx,omitempty"`
	// LastPitch holds the value of the "last_pitch" field.
	LastPitch *string `json:"last_pitch,omitempty"`
	// Mode holds the value of the "mode" field.
	Mode *string `json:"mode,omitempty"`
	// Pitch holds the value of the "pitch" field.
	Pitch int `json:"pitch,omitempty"`
	// StyleLabel holds the value of the "style_label" field.
	StyleLabel string `json:"style_label,omitempty"`
	// StyleLabelVersion holds the value of the "style_label_version" field.
	StyleLabelVersion string `json:"style_label_version,omitempty"`
	// Tempo holds the value of the "tempo" field.
	Tempo int `json:"tempo,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserTypecastSetting) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case usertypecastsetting.FieldNaturalness:
			values[i] = new(sql.NullFloat64)
		case usertypecastsetting.FieldID, usertypecastsetting.FieldMaxSeconds, usertypecastsetting.FieldSpeedX, usertypecastsetting.FieldStyleIdx, usertypecastsetting.FieldPitch, usertypecastsetting.FieldTempo:
			values[i] = new(sql.NullInt64)
		case usertypecastsetting.FieldUserID, usertypecastsetting.FieldActorID, usertypecastsetting.FieldLang, usertypecastsetting.FieldGid, usertypecastsetting.FieldLastPitch, usertypecastsetting.FieldMode, usertypecastsetting.FieldStyleLabel, usertypecastsetting.FieldStyleLabelVersion:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type UserTypecastSetting", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserTypecastSetting fields.
func (uts *UserTypecastSetting) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case usertypecastsetting.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			uts.ID = int(value.Int64)
		case usertypecastsetting.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				uts.UserID = value.String
			}
		case usertypecastsetting.FieldActorID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field actor_id", values[i])
			} else if value.Valid {
				uts.ActorID = value.String
			}
		case usertypecastsetting.FieldLang:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field lang", values[i])
			} else if value.Valid {
				uts.Lang = value.String
			}
		case usertypecastsetting.FieldMaxSeconds:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field max_seconds", values[i])
			} else if value.Valid {
				uts.MaxSeconds = int(value.Int64)
			}
		case usertypecastsetting.FieldNaturalness:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field naturalness", values[i])
			} else if value.Valid {
				uts.Naturalness = value.Float64
			}
		case usertypecastsetting.FieldSpeedX:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field speed_x", values[i])
			} else if value.Valid {
				uts.SpeedX = int(value.Int64)
			}
		case usertypecastsetting.FieldGid:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gid", values[i])
			} else if value.Valid {
				uts.Gid = value.String
			}
		case usertypecastsetting.FieldStyleIdx:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field style_idx", values[i])
			} else if value.Valid {
				uts.StyleIdx = int(value.Int64)
			}
		case usertypecastsetting.FieldLastPitch:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_pitch", values[i])
			} else if value.Valid {
				uts.LastPitch = new(string)
				*uts.LastPitch = value.String
			}
		case usertypecastsetting.FieldMode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mode", values[i])
			} else if value.Valid {
				uts.Mode = new(string)
				*uts.Mode = value.String
			}
		case usertypecastsetting.FieldPitch:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field pitch", values[i])
			} else if value.Valid {
				uts.Pitch = int(value.Int64)
			}
		case usertypecastsetting.FieldStyleLabel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field style_label", values[i])
			} else if value.Valid {
				uts.StyleLabel = value.String
			}
		case usertypecastsetting.FieldStyleLabelVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field style_label_version", values[i])
			} else if value.Valid {
				uts.StyleLabelVersion = value.String
			}
		case usertypecastsetting.FieldTempo:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tempo", values[i])
			} else if value.Valid {
				uts.Tempo = int(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this UserTypecastSetting.
// Note that you need to call UserTypecastSetting.Unwrap() before calling this method if this UserTypecastSetting
// was returned from a transaction, and the transaction was committed or rolled back.
func (uts *UserTypecastSetting) Update() *UserTypecastSettingUpdateOne {
	return (&UserTypecastSettingClient{config: uts.config}).UpdateOne(uts)
}

// Unwrap unwraps the UserTypecastSetting entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (uts *UserTypecastSetting) Unwrap() *UserTypecastSetting {
	_tx, ok := uts.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserTypecastSetting is not a transactional entity")
	}
	uts.config.driver = _tx.drv
	return uts
}

// String implements the fmt.Stringer.
func (uts *UserTypecastSetting) String() string {
	var builder strings.Builder
	builder.WriteString("UserTypecastSetting(")
	builder.WriteString(fmt.Sprintf("id=%v, ", uts.ID))
	builder.WriteString("user_id=")
	builder.WriteString(uts.UserID)
	builder.WriteString(", ")
	builder.WriteString("actor_id=")
	builder.WriteString(uts.ActorID)
	builder.WriteString(", ")
	builder.WriteString("lang=")
	builder.WriteString(uts.Lang)
	builder.WriteString(", ")
	builder.WriteString("max_seconds=")
	builder.WriteString(fmt.Sprintf("%v", uts.MaxSeconds))
	builder.WriteString(", ")
	builder.WriteString("naturalness=")
	builder.WriteString(fmt.Sprintf("%v", uts.Naturalness))
	builder.WriteString(", ")
	builder.WriteString("speed_x=")
	builder.WriteString(fmt.Sprintf("%v", uts.SpeedX))
	builder.WriteString(", ")
	builder.WriteString("gid=")
	builder.WriteString(uts.Gid)
	builder.WriteString(", ")
	builder.WriteString("style_idx=")
	builder.WriteString(fmt.Sprintf("%v", uts.StyleIdx))
	builder.WriteString(", ")
	if v := uts.LastPitch; v != nil {
		builder.WriteString("last_pitch=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := uts.Mode; v != nil {
		builder.WriteString("mode=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("pitch=")
	builder.WriteString(fmt.Sprintf("%v", uts.Pitch))
	builder.WriteString(", ")
	builder.WriteString("style_label=")
	builder.WriteString(uts.StyleLabel)
	builder.WriteString(", ")
	builder.WriteString("style_label_version=")
	builder.WriteString(uts.StyleLabelVersion)
	builder.WriteString(", ")
	builder.WriteString("tempo=")
	builder.WriteString(fmt.Sprintf("%v", uts.Tempo))
	builder.WriteByte(')')
	return builder.String()
}

// UserTypecastSettings is a parsable slice of UserTypecastSetting.
type UserTypecastSettings []*UserTypecastSetting

func (uts UserTypecastSettings) config(cfg config) {
	for _i := range uts {
		uts[_i].config = cfg
	}
}
