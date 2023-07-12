package uuid

import (
	uuid "github.com/nu7hatch/gouuid"
	"log"
)

// Uuid 生成uuid.
func (tu *helper.TsUuid) Uuid() (newUuid string) {
	fileId, uuidErr := uuid.NewV4()
	if uuidErr != nil {
		log.Println("uuid 生成Err:", uuidErr)
		newUuid = "false"
		return
	}
	newUuid = fileId.String()
	return
}
