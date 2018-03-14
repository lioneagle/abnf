#ifndef PS_SIP_CHARSETS_DWORD_BIT_HPP
#define PS_SIP_CHARSETS_DWORD_BIT_HPP

/*---------------- mask definition ----------------*/
const PS_DWORD  PS_SIP_CHARSETS_MASK_DIGIT               = 0x00000001;
const PS_DWORD  PS_SIP_CHARSETS_MASK_ALPHA               = 0x00000002;
const PS_DWORD  PS_SIP_CHARSETS_MASK_LOWER               = 0x00000004;
const PS_DWORD  PS_SIP_CHARSETS_MASK_UPPER               = 0x00000008;
const PS_DWORD  PS_SIP_CHARSETS_MASK_ALPHANUM            = 0x00000010;
const PS_DWORD  PS_SIP_CHARSETS_MASK_HEX                 = 0x00000020;
const PS_DWORD  PS_SIP_CHARSETS_MASK_LOWER_HEX_ALPHA     = 0x00000040;
const PS_DWORD  PS_SIP_CHARSETS_MASK_UPPER_HEX_ALPHA     = 0x00000080;
const PS_DWORD  PS_SIP_CHARSETS_MASK_WSP                 = 0x00000100;

/*---------------- action declaration ----------------*/
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

#endif /* PS_SIP_CHARSETS_DWORD_BIT_HPP */

