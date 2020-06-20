package api

import (
	"bytes"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"makespace-remaster/serializer"
	"net/http"
	"time"
)

func Captcha(c *gin.Context, length ...int) {
	l := captcha.DefaultLen
	w, h := 107, 36
	if len(length) == 1 {
		l = length[0]
	}
	if len(length) == 2 {
		w = length[1]
	}
	if len(length) == 3 {
		h = length[2]
	}
	captchaId := captcha.NewLen(l)
	_ = serve(c.Writer, c.Request, captchaId, ".png", "zh", false, w, h)
}
func CaptchaVerify(c *gin.Context) {
	type verify struct {
		captchaId string `json:"captcha_id" validate:"required"`
		code      string `json:"code" validate:"required"`
	}
	var v verify
	if err := c.ShouldBindJSON(v); err == nil {
		if captcha.VerifyString(v.captchaId, v.code) {
			c.JSON(http.StatusOK,serializer.PureErrorResponse{
				Status: 0,
				Msg:    "success",
			})
		} else {
			c.JSON(http.StatusOK,serializer.PureErrorResponse{
				Status: -1,
				Msg:    "failed",
			})
		}
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}

}

func serve(w http.ResponseWriter, r *http.Request, id, ext, lang string, download bool, width, height int) error {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	var content bytes.Buffer
	switch ext {
	case ".png":
		w.Header().Set("Content-Type", "image/png")
		_ = captcha.WriteImage(&content, id, width, height)
	case ".wav":
		w.Header().Set("Content-Type", "audio/x-wav")
		_ = captcha.WriteAudio(&content, id, lang)
	default:
		return captcha.ErrNotFound
	}

	if download {
		w.Header().Set("Content-Type", "application/octet-stream")
	}
	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))
	return nil
}