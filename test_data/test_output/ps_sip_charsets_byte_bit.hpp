#ifndef PS_SIP_CHARSETS_BYTE_BIT_HPP
#define PS_SIP_CHARSETS_BYTE_BIT_HPP

/*---------------- mask definition ----------------*/
const PS_BYTE  PS_SIP_CHARSETS_MASK_DIGIT               = 0x01;
const PS_BYTE  PS_SIP_CHARSETS_MASK_ALPHA               = 0x02;
const PS_BYTE  PS_SIP_CHARSETS_MASK_LOWER               = 0x04;
const PS_BYTE  PS_SIP_CHARSETS_MASK_UPPER               = 0x08;
const PS_BYTE  PS_SIP_CHARSETS_MASK_ALPHANUM            = 0x10;
const PS_BYTE  PS_SIP_CHARSETS_MASK_HEX                 = 0x20;
const PS_BYTE  PS_SIP_CHARSETS_MASK_LOWER_HEX_ALPHA     = 0x40;
const PS_BYTE  PS_SIP_CHARSETS_MASK_UPPER_HEX_ALPHA     = 0x80;
const PS_BYTE  PS_SIP_CHARSETS_MASK_WSP                 = 0x01;

/*---------------- action declaration ----------------*/
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

#endif /* PS_SIP_CHARSETS_BYTE_BIT_HPP */

