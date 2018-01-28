#ifndef PS_SIP_CHARSETS_DWORD_BIT_H
#define PS_SIP_CHARSETS_DWORD_BIT_H

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

/*---------------- action definition ----------------*/
inline bool PS_SIP_isDigit(unsigned char ch)          { return g_sipCharsets0[ch] & PS_SIP_CHARSETS_MASK_DIGIT; }
inline bool PS_SIP_isAlpha(unsigned char ch)          { return g_sipCharsets0[ch] & PS_SIP_CHARSETS_MASK_ALPHA; }
inline bool PS_SIP_isLower(unsigned char ch)          { return g_sipCharsets0[ch] & PS_SIP_CHARSETS_MASK_LOWER; }
inline bool PS_SIP_isUpper(unsigned char ch)          { return g_sipCharsets0[ch] & PS_SIP_CHARSETS_MASK_UPPER; }
inline bool PS_SIP_isAlphanum(unsigned char ch)       { return g_sipCharsets0[ch] & PS_SIP_CHARSETS_MASK_ALPHANUM; }
inline bool PS_SIP_isHex(unsigned char ch)            { return g_sipCharsets0[ch] & PS_SIP_CHARSETS_MASK_HEX; }
inline bool PS_SIP_isLowerHexAlpha(unsigned char ch)  { return g_sipCharsets0[ch] & PS_SIP_CHARSETS_MASK_LOWER_HEX_ALPHA; }
inline bool PS_SIP_isUpperHexAlpha(unsigned char ch)  { return g_sipCharsets0[ch] & PS_SIP_CHARSETS_MASK_UPPER_HEX_ALPHA; }
inline bool PS_SIP_isWsp(unsigned char ch)            { return g_sipCharsets0[ch] & PS_SIP_CHARSETS_MASK_WSP; }

/*---------------- var declaration ----------------*/
extern PS_DWORD const g_sipCharsets0[256];

#endif /* PS_SIP_CHARSETS_DWORD_BIT_H */

