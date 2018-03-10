#include "ps_sip_header_key_cmp_2.h"

unsigned int GetSipHeaderIndex(char const** src, char const* end) {
    char const*  p = *src;
    
    if (p == NULL || p >= end) {
        return ABNF_SIP_HDR_UNKNOWN;
    }
    
    if ((src < end) && ((*(src++) | 0x20) == 'f') {
      if (p >= end) {
          *src = p;
          return ABNF_SIP_HDR_FROM;
      }
      if ((src < end) && ((*(src++) | 0x20) == 'r') {
        if ((src < end) && ((*(src++) | 0x20) == 'o') {
          if ((src < end) && ((*(src++) | 0x20) == 'm') {
            if (p >= end) {
                *src = p;
                return ABNF_SIP_HDR_FROM;
            }
          }
        }
      }
    }
    
    return ABNF_SIP_HDR_UNKNOWN;
}
