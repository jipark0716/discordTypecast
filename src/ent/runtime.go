// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/jipark0716/discordTypecast/ent/schema"
	"github.com/jipark0716/discordTypecast/ent/typecastmessage"
	"github.com/jipark0716/discordTypecast/ent/usertypecastsetting"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	typecastmessageFields := schema.TypecastMessage{}.Fields()
	_ = typecastmessageFields
	// typecastmessageDescStatus is the schema descriptor for status field.
	typecastmessageDescStatus := typecastmessageFields[17].Descriptor()
	// typecastmessage.DefaultStatus holds the default value on creation for the status field.
	typecastmessage.DefaultStatus = typecastmessageDescStatus.Default.(int8)
	usertypecastsettingFields := schema.UserTypecastSetting{}.Fields()
	_ = usertypecastsettingFields
	// usertypecastsettingDescLang is the schema descriptor for lang field.
	usertypecastsettingDescLang := usertypecastsettingFields[2].Descriptor()
	// usertypecastsetting.DefaultLang holds the default value on creation for the lang field.
	usertypecastsetting.DefaultLang = usertypecastsettingDescLang.Default.(string)
	// usertypecastsettingDescMaxSeconds is the schema descriptor for max_seconds field.
	usertypecastsettingDescMaxSeconds := usertypecastsettingFields[3].Descriptor()
	// usertypecastsetting.DefaultMaxSeconds holds the default value on creation for the max_seconds field.
	usertypecastsetting.DefaultMaxSeconds = usertypecastsettingDescMaxSeconds.Default.(int)
	// usertypecastsettingDescNaturalness is the schema descriptor for naturalness field.
	usertypecastsettingDescNaturalness := usertypecastsettingFields[4].Descriptor()
	// usertypecastsetting.DefaultNaturalness holds the default value on creation for the naturalness field.
	usertypecastsetting.DefaultNaturalness = usertypecastsettingDescNaturalness.Default.(float64)
	// usertypecastsettingDescSpeedX is the schema descriptor for speed_x field.
	usertypecastsettingDescSpeedX := usertypecastsettingFields[5].Descriptor()
	// usertypecastsetting.DefaultSpeedX holds the default value on creation for the speed_x field.
	usertypecastsetting.DefaultSpeedX = usertypecastsettingDescSpeedX.Default.(int)
	// usertypecastsettingDescGid is the schema descriptor for gid field.
	usertypecastsettingDescGid := usertypecastsettingFields[6].Descriptor()
	// usertypecastsetting.DefaultGid holds the default value on creation for the gid field.
	usertypecastsetting.DefaultGid = usertypecastsettingDescGid.Default.(string)
	// usertypecastsettingDescStyleIdx is the schema descriptor for style_idx field.
	usertypecastsettingDescStyleIdx := usertypecastsettingFields[7].Descriptor()
	// usertypecastsetting.DefaultStyleIdx holds the default value on creation for the style_idx field.
	usertypecastsetting.DefaultStyleIdx = usertypecastsettingDescStyleIdx.Default.(int)
	// usertypecastsettingDescPitch is the schema descriptor for pitch field.
	usertypecastsettingDescPitch := usertypecastsettingFields[10].Descriptor()
	// usertypecastsetting.DefaultPitch holds the default value on creation for the pitch field.
	usertypecastsetting.DefaultPitch = usertypecastsettingDescPitch.Default.(int)
	// usertypecastsettingDescStyleLabel is the schema descriptor for style_label field.
	usertypecastsettingDescStyleLabel := usertypecastsettingFields[11].Descriptor()
	// usertypecastsetting.DefaultStyleLabel holds the default value on creation for the style_label field.
	usertypecastsetting.DefaultStyleLabel = usertypecastsettingDescStyleLabel.Default.(string)
	// usertypecastsettingDescStyleLabelVersion is the schema descriptor for style_label_version field.
	usertypecastsettingDescStyleLabelVersion := usertypecastsettingFields[12].Descriptor()
	// usertypecastsetting.DefaultStyleLabelVersion holds the default value on creation for the style_label_version field.
	usertypecastsetting.DefaultStyleLabelVersion = usertypecastsettingDescStyleLabelVersion.Default.(string)
	// usertypecastsettingDescTempo is the schema descriptor for tempo field.
	usertypecastsettingDescTempo := usertypecastsettingFields[13].Descriptor()
	// usertypecastsetting.DefaultTempo holds the default value on creation for the tempo field.
	usertypecastsetting.DefaultTempo = usertypecastsettingDescTempo.Default.(int)
}
