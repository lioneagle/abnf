#include "ps_sip_header_key_cmp.h"

unsigned int GetSipHeaderIndex(char const** src, char const* end)
{
    char const*  p = *src;
    
    if (p == NULL || p >= end)
    {
        return ABNF_SIP_HDR_UNKNOWN;
    }
    
    switch (*(p++))
    {
        case 'F':
            if ((p < end) && (*(p++) == 'r')
            {
              if ((p < end) && (*(p++) == 'o')
              {
                if ((p < end) && (*(p++) == 'm')
                {
                  if (p >= end)
                  {
                      *src = p;
                      return ABNF_SIP_HDR_FROM;
                  }
                }
              }
            }

            {
                *src = p;
                return ABNF_SIP_HDR_UNKNOWN;
            }
        case 'f':
            if (p >= end)
            {
                *src = p;
                return ABNF_SIP_HDR_FROM;
            }

            {
                *src = p;
                return ABNF_SIP_HDR_UNKNOWN;
            }
    }
    
    return ABNF_SIP_HDR_UNKNOWN;
}
