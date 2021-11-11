package model

import (
	"fmt"
)

type TypecastSetting struct {
	ID int `db:"id"`
	MemberID string `db:"member_id"`
	ActorName string `db:"actor_name"`
	IsMute bool `db:"mute"`
}

func GetTypecastSetting(memberID string) TypecastSetting {
	sql := fmt.Sprintf("SELECT * FROM `typecast_settings` WHERE `member_id` = '%s'", memberID)
	result := []TypecastSetting{}
	Get(&result, sql)
	if len(result) == 0 {
		return TypecastSetting{
			MemberID: memberID,
		}
	} else {
		return result[0]
	}
}

func (this *TypecastSetting) Save() (err error) {
	if this.ID != 0 {
		_, err = Update(map[string]interface{}{
			"actor_name": this.ActorName,
			"mute": BooltoInt(this.IsMute),
		}, "typecast_settings", fmt.Sprintf("`id` = %d", this.ID))
	} else {
		_, err = Insert(map[string]interface{}{
			"actor_name": this.ActorName,
			"member_id": this.MemberID,
			"mute": BooltoInt(this.IsMute),
		}, "typecast_settings")
	}
	return
}
