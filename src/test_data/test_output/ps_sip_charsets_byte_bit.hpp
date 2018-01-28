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
inline bool PS_SIP_isDigit(unsigned char ch)          { return g_sipCharsets0[ch] & PS_SIP_CHARSETS_MASK_DIGIT; }
inline bool PS_SIP_isAlpha(unsigned char ch)          { return g_sipCharsets0[ch] & PS_SIP_CHARSETS_MASK_ALPHA; }
inline bool PS_SIP_isLower(unsigned char ch)          { return g_sipCharsets0[ch] & PS_SIP_CHARSETS_MASK_LOWER; }
inline bool PS_SIP_isUpper(unsigned char ch)          { return g_sipCharsets0[ch] & PS_SIP_CHARSETS_MASK_UPPER; }
inline bool PS_SIP_isAlphanum(unsigned char ch)       { return g_sipCharsets0[ch] & PS_SIP_CHARSETS_MASK_ALPHANUM; }
inline bool PS_SIP_isHex(unsigned char ch)            { return g_sipCharsets0[ch] & PS_SIP_CHARSETS_MASK_HEX; }
inline bool PS_SIP_isLowerHexAlpha(unsigned char ch)  { return g_sipCharsets0[ch] & PS_SIP_CHARSETS_MASK_LOWER_HEX_ALPHA; }
inline bool PS_SIP_isUpperHexAlpha(unsigned char ch)  { return g_sipCharsets0[ch] & PS_SIP_CHARSETS_MASK_UPPER_HEX_ALPHA; }
inline bool PS_SIP_isWsp(unsigned char ch)            { return g_sipCharsets1[ch] & PS_SIP_CHARSETS_MASK_WSP; }

/*---------------- var declaration ----------------*/
extern PS_BYTE const g_sipCharsets0[256];
extern PS_BYTE const g_sipCharsets1[256];

#endif /* PS_SIP_CHARSETS_BYTE_BIT_H */
