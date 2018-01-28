#ifndef PS_SIP_CHARSETS_BYTE_NO_BIT_H
#define PS_SIP_CHARSETS_BYTE_NO_BIT_H

/*---------------- action definition ----------------*/
inline bool PS_SIP_isDigit(unsigned char ch)          { return g_sipCharsets0[ch]; }
inline bool PS_SIP_isAlpha(unsigned char ch)          { return g_sipCharsets1[ch]; }
inline bool PS_SIP_isLower(unsigned char ch)          { return g_sipCharsets2[ch]; }
inline bool PS_SIP_isUpper(unsigned char ch)          { return g_sipCharsets3[ch]; }
inline bool PS_SIP_isAlphanum(unsigned char ch)       { return g_sipCharsets4[ch]; }
inline bool PS_SIP_isHex(unsigned char ch)            { return g_sipCharsets5[ch]; }
inline bool PS_SIP_isLowerHexAlpha(unsigned char ch)  { return g_sipCharsets6[ch]; }
inline bool PS_SIP_isUpperHexAlpha(unsigned char ch)  { return g_sipCharsets7[ch]; }
inline bool PS_SIP_isWsp(unsigned char ch)            { return g_sipCharsets8[ch]; }

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

#endif /* PS_SIP_CHARSETS_BYTE_NO_BIT_H */
