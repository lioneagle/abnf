package charset_gen

import (
	"fmt"
	"testing"

	"github.com/lioneagle/goutil/src/test"
)

func TestCharsetInfoGetMaskNameAndGetActionName(t *testing.T) {
	testdata := []struct {
		name             string
		maskPrefix       string
		actionPrefix     string
		actionFirstLower bool
		maskName         string
		actionName       string
	}{
		{"digit", "", "", false, "MASK_DIGIT", "IsDigit"},
		{"digit", "sip_chars", "SIP", true, "SIP_CHARS_MASK_DIGIT", "SIP_isDigit"},
		{"alpha-digit", "sip_chars", "SIP", true, "SIP_CHARS_MASK_ALPHA_DIGIT", "SIP_isAlphaDigit"},
	}

	for i, v := range testdata {
		v := v
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			config := NewConfig()
			config.SetMaskPrefix(v.maskPrefix)
			config.SetActionPrefix(v.actionPrefix)
			config.SetActionFirstLower(v.actionFirstLower)
			c := NewCharsetInfo("")
			c.Name = v.name

			test.EXPECT_EQ(t, c.GetMaskName(config), v.maskName, "")
			test.EXPECT_EQ(t, c.GetActionName(config), v.actionName, "")
		})
	}

}
