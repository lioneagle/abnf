#ifndef PS_SIP_CHARSETS_DWORD_BIT_H
#define PS_SIP_CHARSETS_DWORD_BIT_H

#ifdef __cplusplus
extern "C"
{
#endif

/*---------------- mask definition ----------------*/
#define PS_SIP_CHARSETS_MASK_DIGIT              ((PS_DWORD)(0x00000001))
#define PS_SIP_CHARSETS_MASK_ALPHA              ((PS_DWORD)(0x00000002))
#define PS_SIP_CHARSETS_MASK_LOWER              ((PS_DWORD)(0x00000004))
#define PS_SIP_CHARSETS_MASK_UPPER              ((PS_DWORD)(0x00000008))
#define PS_SIP_CHARSETS_MASK_ALPHANUM           ((PS_DWORD)(0x00000010))
#define PS_SIP_CHARSETS_MASK_HEX                ((PS_DWORD)(0x00000020))
#define PS_SIP_CHARSETS_MASK_LOWER_HEX_ALPHA    ((PS_DWORD)(0x00000040))
#define PS_SIP_CHARSETS_MASK_UPPER_HEX_ALPHA    ((PS_DWORD)(0x00000080))
#define PS_SIP_CHARSETS_MASK_WSP                ((PS_DWORD)(0x00000100))

/*---------------- action declaration ----------------*/
#define PS_SIP_isDigit(ch)            (g_sipCharsets0[(unsigned char)(ch)] & PS_SIP_CHARSETS_MASK_DIGIT)
#define PS_SIP_isAlpha(ch)            (g_sipCharsets0[(unsigned char)(ch)] & PS_SIP_CHARSETS_MASK_ALPHA)
#define PS_SIP_isLower(ch)            (g_sipCharsets0[(unsigned char)(ch)] & PS_SIP_CHARSETS_MASK_LOWER)
#define PS_SIP_isUpper(ch)            (g_sipCharsets0[(unsigned char)(ch)] & PS_SIP_CHARSETS_MASK_UPPER)
#define PS_SIP_isAlphanum(ch)         (g_sipCharsets0[(unsigned char)(ch)] & PS_SIP_CHARSETS_MASK_ALPHANUM)
#define PS_SIP_isHex(ch)              (g_sipCharsets0[(unsigned char)(ch)] & PS_SIP_CHARSETS_MASK_HEX)
#define PS_SIP_isLowerHexAlpha(ch)    (g_sipCharsets0[(unsigned char)(ch)] & PS_SIP_CHARSETS_MASK_LOWER_HEX_ALPHA)
#define PS_SIP_isUpperHexAlpha(ch)    (g_sipCharsets0[(unsigned char)(ch)] & PS_SIP_CHARSETS_MASK_UPPER_HEX_ALPHA)
#define PS_SIP_isWsp(ch)              (g_sipCharsets0[(unsigned char)(ch)] & PS_SIP_CHARSETS_MASK_WSP)

/*---------------- var declaration ----------------*/
extern PS_DWORD const g_sipCharsets0[256];

#ifdef __cplusplus
}
#endif

#endif /* PS_SIP_CHARSETS_DWORD_BIT_H */

