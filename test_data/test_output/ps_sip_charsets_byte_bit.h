#ifndef PS_SIP_CHARSETS_BYTE_BIT_H
#define PS_SIP_CHARSETS_BYTE_BIT_H

/*---------------- mask definition ----------------*/
#define PS_SIP_CHARSETS_MASK_DIGIT              ((PS_BYTE)(0x01))
#define PS_SIP_CHARSETS_MASK_ALPHA              ((PS_BYTE)(0x02))
#define PS_SIP_CHARSETS_MASK_LOWER              ((PS_BYTE)(0x04))
#define PS_SIP_CHARSETS_MASK_UPPER              ((PS_BYTE)(0x08))
#define PS_SIP_CHARSETS_MASK_ALPHANUM           ((PS_BYTE)(0x10))
#define PS_SIP_CHARSETS_MASK_HEX                ((PS_BYTE)(0x20))
#define PS_SIP_CHARSETS_MASK_LOWER_HEX_ALPHA    ((PS_BYTE)(0x40))
#define PS_SIP_CHARSETS_MASK_UPPER_HEX_ALPHA    ((PS_BYTE)(0x80))
#define PS_SIP_CHARSETS_MASK_WSP                ((PS_BYTE)(0x01))

/*---------------- action definition ----------------*/
#define PS_SIP_isDigit(ch)            (g_sipCharsets0[(unsigned char)(ch)] & PS_SIP_CHARSETS_MASK_DIGIT)
#define PS_SIP_isAlpha(ch)            (g_sipCharsets0[(unsigned char)(ch)] & PS_SIP_CHARSETS_MASK_ALPHA)
#define PS_SIP_isLower(ch)            (g_sipCharsets0[(unsigned char)(ch)] & PS_SIP_CHARSETS_MASK_LOWER)
#define PS_SIP_isUpper(ch)            (g_sipCharsets0[(unsigned char)(ch)] & PS_SIP_CHARSETS_MASK_UPPER)
#define PS_SIP_isAlphanum(ch)         (g_sipCharsets0[(unsigned char)(ch)] & PS_SIP_CHARSETS_MASK_ALPHANUM)
#define PS_SIP_isHex(ch)              (g_sipCharsets0[(unsigned char)(ch)] & PS_SIP_CHARSETS_MASK_HEX)
#define PS_SIP_isLowerHexAlpha(ch)    (g_sipCharsets0[(unsigned char)(ch)] & PS_SIP_CHARSETS_MASK_LOWER_HEX_ALPHA)
#define PS_SIP_isUpperHexAlpha(ch)    (g_sipCharsets0[(unsigned char)(ch)] & PS_SIP_CHARSETS_MASK_UPPER_HEX_ALPHA)
#define PS_SIP_isWsp(ch)              (g_sipCharsets1[(unsigned char)(ch)] & PS_SIP_CHARSETS_MASK_WSP)

/*---------------- var declaration ----------------*/
extern PS_BYTE const g_sipCharsets0[256];
extern PS_BYTE const g_sipCharsets1[256];

#endif /* PS_SIP_CHARSETS_BYTE_BIT_H */
