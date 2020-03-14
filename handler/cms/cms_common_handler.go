package cms

import (
	"github.com/gorilla/mux"
	"net/http"
	"wechat-mall-backend/defs"
	"wechat-mall-backend/utils"
)

// 生成-OSSPolicyToken
func (h *Handler) GetOSSPolicyToken(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dir := vars["dir"]

	ossConf := h.conf.Oss
	oss := utils.OSSPolicyToken{
		AccessKeyId:     ossConf.AccessKeyId,
		AccessKeySecret: ossConf.AccessKeySecret,
		Host:            "https://" + ossConf.BucketName + ".oss-cn-hangzhou.aliyuncs.com",
		UploadDir:       dir,
		ExpireTime:      30,
	}
	response := oss.GetPolicyToken()
	defs.SendNormalResponse(w, response)
}
