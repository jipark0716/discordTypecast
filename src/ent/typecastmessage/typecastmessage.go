// Code generated by ent, DO NOT EDIT.

package typecastmessage

const (
	// Label holds the string label denoting the typecastmessage type in the database.
	Label = "typecast_message"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldText holds the string denoting the text field in the database.
	FieldText = "text"
	// FieldGuildID holds the string denoting the guild_id field in the database.
	FieldGuildID = "guild_id"
	// FieldChannelID holds the string denoting the channel_id field in the database.
	FieldChannelID = "channel_id"
	// FieldActorID holds the string denoting the actor_id field in the database.
	FieldActorID = "actor_id"
	// FieldLang holds the string denoting the lang field in the database.
	FieldLang = "lang"
	// FieldMaxSeconds holds the string denoting the max_seconds field in the database.
	FieldMaxSeconds = "max_seconds"
	// FieldNaturalness holds the string denoting the naturalness field in the database.
	FieldNaturalness = "naturalness"
	// FieldSpeedX holds the string denoting the speed_x field in the database.
	FieldSpeedX = "speed_x"
	// FieldGid holds the string denoting the gid field in the database.
	FieldGid = "gid"
	// FieldStyleIdx holds the string denoting the style_idx field in the database.
	FieldStyleIdx = "style_idx"
	// FieldLastPitch holds the string denoting the last_pitch field in the database.
	FieldLastPitch = "last_pitch"
	// FieldMode holds the string denoting the mode field in the database.
	FieldMode = "mode"
	// FieldPitch holds the string denoting the pitch field in the database.
	FieldPitch = "pitch"
	// FieldStyleLabel holds the string denoting the style_label field in the database.
	FieldStyleLabel = "style_label"
	// FieldStyleLabelVersion holds the string denoting the style_label_version field in the database.
	FieldStyleLabelVersion = "style_label_version"
	// FieldTempo holds the string denoting the tempo field in the database.
	FieldTempo = "tempo"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldSendAt holds the string denoting the send_at field in the database.
	FieldSendAt = "send_at"
	// Table holds the table name of the typecastmessage in the database.
	Table = "typecast_messages"
)

// Columns holds all SQL columns for typecastmessage fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldText,
	FieldGuildID,
	FieldChannelID,
	FieldActorID,
	FieldLang,
	FieldMaxSeconds,
	FieldNaturalness,
	FieldSpeedX,
	FieldGid,
	FieldStyleIdx,
	FieldLastPitch,
	FieldMode,
	FieldPitch,
	FieldStyleLabel,
	FieldStyleLabelVersion,
	FieldTempo,
	FieldStatus,
	FieldSendAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus int8
)
