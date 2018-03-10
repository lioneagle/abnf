#ifndef PS_SIP_HEADER_KEY_CMP_H
#define PS_SIP_HEADER_KEY_CMP_H

/*---------------- index definition ----------------*/
#define ABNF_SIP_HDR_UNKNOWN ((unsigned int)(0))
#define ABNF_SIP_HDR_FROM    ((unsigned int)(1))

/*---------------- action declaration ----------------*/
unsigned int GetSipHeaderIndex(char const** src, char const* end);

#endif /* PS_SIP_HEADER_KEY_CMP_H */
