#ifndef PS_SIP_CHARSETS_BYTE_NO_BIT_H
#define PS_SIP_CHARSETS_BYTE_NO_BIT_H

#ifdef __cplusplus
extern "C"
{
#endif

/*---------------- action declaration ----------------*/
#define PS_SIP_isDigit(ch)            (g_sipCharsets0[(unsigned char)(ch)])
#define PS_SIP_isAlpha(ch)            (g_sipCharsets1[(unsigned char)(ch)])
#define PS_SIP_isLower(ch)            (g_sipCharsets2[(unsigned char)(ch)])
#define PS_SIP_isUpper(ch)            (g_sipCharsets3[(unsigned char)(ch)])
#define PS_SIP_isAlphanum(ch)         (g_sipCharsets4[(unsigned char)(ch)])
#define PS_SIP_isHex(ch)              (g_sipCharsets5[(unsigned char)(ch)])
#define PS_SIP_isLowerHexAlpha(ch)    (g_sipCharsets6[(unsigned char)(ch)])
#define PS_SIP_isUpperHexAlpha(ch)    (g_sipCharsets7[(unsigned char)(ch)])
#define PS_SIP_isWsp(ch)              (g_sipCharsets8[(unsigned char)(ch)])

/*---------------- var declaration ----------------*/
extern PS_BYTE const g_sipCharsets0[256];
extern PS_BYTE const g_sipCharsets1[256];
extern PS_BYTE const g_sipCharsets2[256];
extern PS_BYTE const g_sipCharsets3[256];
extern PS_BYTE const g_sipCharsets4[256];
extern PS_BYTE const g_sipCharsets5[256];
extern PS_BYTE const g_sipCharsets6[256];
extern PS_BYTE const g_sipCharsets7[256];
extern PS_BYTE const g_sipCharsets8[256];

#ifdef __cplusplus
}
#endif

#endif /* PS_SIP_CHARSETS_BYTE_NO_BIT_H */

