#include "ps_sip_header_key_cmp_2.h"

unsigned int GetSipHeaderIndex(char const** src, char const* end) {
    char const*  p = *src;
    
    if (p == NULL || p >= end) {
        return ABNF_SIP_HDR_UNKNOWN;
    }
    
    switch (*(p++)) {
    }
    
    return ABNF_SIP_HDR_UNKNOWN;
}
