// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/jipark0716/discordTypecast/ent/typecastmessage"
)

// TypecastMessage is the model entity for the TypecastMessage schema.
type TypecastMessage struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID string `json:"user_id,omitempty"`
	// Text holds the value of the "text" field.
	Text string `json:"text,omitempty"`
	// GuildID holds the value of the "guild_id" field.
	GuildID string `json:"guild_id,omitempty"`
	// ChannelID holds the value of the "channel_id" field.
	ChannelID string `json:"channel_id,omitempty"`
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
	// Status holds the value of the "status" field.
	Status int8 `json:"status,omitempty"`
	// SendAt holds the value of the "send_at" field.
	SendAt *time.Time `json:"send_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TypecastMessage) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case typecastmessage.FieldNaturalness:
			values[i] = new(sql.NullFloat64)
		case typecastmessage.FieldID, typecastmessage.FieldMaxSeconds, typecastmessage.FieldSpeedX, typecastmessage.FieldStyleIdx, typecastmessage.FieldPitch, typecastmessage.FieldTempo, typecastmessage.FieldStatus:
			values[i] = new(sql.NullInt64)
		case typecastmessage.FieldUserID, typecastmessage.FieldText, typecastmessage.FieldGuildID, typecastmessage.FieldChannelID, typecastmessage.FieldActorID, typecastmessage.FieldLang, typecastmessage.FieldGid, typecastmessage.FieldLastPitch, typecastmessage.FieldMode, typecastmessage.FieldStyleLabel, typecastmessage.FieldStyleLabelVersion:
			values[i] = new(sql.NullString)
		case typecastmessage.FieldSendAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type TypecastMessage", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TypecastMessage fields.
func (tm *TypecastMessage) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case typecastmessage.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			tm.ID = int(value.Int64)
		case typecastmessage.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				tm.UserID = value.String
			}
		case typecastmessage.FieldText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field text", values[i])
			} else if value.Valid {
				tm.Text = value.String
			}
		case typecastmessage.FieldGuildID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field guild_id", values[i])
			} else if value.Valid {
				tm.GuildID = value.String
			}
		case typecastmessage.FieldChannelID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field channel_id", values[i])
			} else if value.Valid {
				tm.ChannelID = value.String
			}
		case typecastmessage.FieldActorID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field actor_id", values[i])
			} else if value.Valid {
				tm.ActorID = value.String
			}
		case typecastmessage.FieldLang:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field lang", values[i])
			} else if value.Valid {
				tm.Lang = value.String
			}
		case typecastmessage.FieldMaxSeconds:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field max_seconds", values[i])
			} else if value.Valid {
				tm.MaxSeconds = int(value.Int64)
			}
		case typecastmessage.FieldNaturalness:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field naturalness", values[i])
			} else if value.Valid {
				tm.Naturalness = value.Float64
			}
		case typecastmessage.FieldSpeedX:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field speed_x", values[i])
			} else if value.Valid {
				tm.SpeedX = int(value.Int64)
			}
		case typecastmessage.FieldGid:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gid", values[i])
			} else if value.Valid {
				tm.Gid = value.String
			}
		case typecastmessage.FieldStyleIdx:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field style_idx", values[i])
			} else if value.Valid {
				tm.StyleIdx = int(value.Int64)
			}
		case typecastmessage.FieldLastPitch:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_pitch", values[i])
			} else if value.Valid {
				tm.LastPitch = new(string)
				*tm.LastPitch = value.String
			}
		case typecastmessage.FieldMode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mode", values[i])
			} else if value.Valid {
				tm.Mode = new(string)
				*tm.Mode = value.String
			}
		case typecastmessage.FieldPitch:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field pitch", values[i])
			} else if value.Valid {
				tm.Pitch = int(value.Int64)
			}
		case typecastmessage.FieldStyleLabel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field style_label", values[i])
			} else if value.Valid {
				tm.StyleLabel = value.String
			}
		case typecastmessage.FieldStyleLabelVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field style_label_version", values[i])
			} else if value.Valid {
				tm.StyleLabelVersion = value.String
			}
		case typecastmessage.FieldTempo:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tempo", values[i])
			} else if value.Valid {
				tm.Tempo = int(value.Int64)
			}
		case typecastmessage.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				tm.Status = int8(value.Int64)
			}
		case typecastmessage.FieldSendAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field send_at", values[i])
			} else if value.Valid {
				tm.SendAt = new(time.Time)
				*tm.SendAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this TypecastMessage.
// Note that you need to call TypecastMessage.Unwrap() before calling this method if this TypecastMessage
// was returned from a transaction, and the transaction was committed or rolled back.
func (tm *TypecastMessage) Update() *TypecastMessageUpdateOne {
	return (&TypecastMessageClient{config: tm.config}).UpdateOne(tm)
}

// Unwrap unwraps the TypecastMessage entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tm *TypecastMessage) Unwrap() *TypecastMessage {
	_tx, ok := tm.config.driver.(*txDriver)
	if !ok {
		panic("ent: TypecastMessage is not a transactional entity")
	}
	tm.config.driver = _tx.drv
	return tm
}

// String implements the fmt.Stringer.
func (tm *TypecastMessage) String() string {
	var builder strings.Builder
	builder.WriteString("TypecastMessage(")
	builder.WriteString(fmt.Sprintf("id=%v, ", tm.ID))
	builder.WriteString("user_id=")
	builder.WriteString(tm.UserID)
	builder.WriteString(", ")
	builder.WriteString("text=")
	builder.WriteString(tm.Text)
	builder.WriteString(", ")
	builder.WriteString("guild_id=")
	builder.WriteString(tm.GuildID)
	builder.WriteString(", ")
	builder.WriteString("channel_id=")
	builder.WriteString(tm.ChannelID)
	builder.WriteString(", ")
	builder.WriteString("actor_id=")
	builder.WriteString(tm.ActorID)
	builder.WriteString(", ")
	builder.WriteString("lang=")
	builder.WriteString(tm.Lang)
	builder.WriteString(", ")
	builder.WriteString("max_seconds=")
	builder.WriteString(fmt.Sprintf("%v", tm.MaxSeconds))
	builder.WriteString(", ")
	builder.WriteString("naturalness=")
	builder.WriteString(fmt.Sprintf("%v", tm.Naturalness))
	builder.WriteString(", ")
	builder.WriteString("speed_x=")
	builder.WriteString(fmt.Sprintf("%v", tm.SpeedX))
	builder.WriteString(", ")
	builder.WriteString("gid=")
	builder.WriteString(tm.Gid)
	builder.WriteString(", ")
	builder.WriteString("style_idx=")
	builder.WriteString(fmt.Sprintf("%v", tm.StyleIdx))
	builder.WriteString(", ")
	if v := tm.LastPitch; v != nil {
		builder.WriteString("last_pitch=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := tm.Mode; v != nil {
		builder.WriteString("mode=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("pitch=")
	builder.WriteString(fmt.Sprintf("%v", tm.Pitch))
	builder.WriteString(", ")
	builder.WriteString("style_label=")
	builder.WriteString(tm.StyleLabel)
	builder.WriteString(", ")
	builder.WriteString("style_label_version=")
	builder.WriteString(tm.StyleLabelVersion)
	builder.WriteString(", ")
	builder.WriteString("tempo=")
	builder.WriteString(fmt.Sprintf("%v", tm.Tempo))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", tm.Status))
	builder.WriteString(", ")
	if v := tm.SendAt; v != nil {
		builder.WriteString("send_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// TypecastMessages is a parsable slice of TypecastMessage.
type TypecastMessages []*TypecastMessage

func (tm TypecastMessages) config(cfg config) {
	for _i := range tm {
		tm[_i].config = cfg
	}
}
