#include "ps_sip_header_key_cmp_2.hpp"

unsigned int GetSipHeaderIndex(char const** src, char const* end)
{
    char const*  p = *src;
    
    if (p == NULL || p >= end) {
        return ABNF_SIP_HDR_UNKNOWN;
    }

    switch (*(p++) | 0x20) {
        case 'a':
            if (p >= end) {
                *src = p;
                return ABNF_SIP_HDR_ACCEPT_CONTACT;
            }
            switch (*(p++) | 0x20) {
                case 'c':
                    if ((p+11) >= end) {
                        *src = p;
                        return ABNF_SIP_HDR_UNKNOWN;
                    }
                    if (((*(p++) | 0x20) == 'c')
                        && ((*(p++) | 0x20) == 'e')
                        && ((*(p++) | 0x20) == 'p')
                        && ((*(p++) | 0x20) == 't')
                        && (*(p++) == '-')
                        && ((*(p++) | 0x20) == 'c')
                        && ((*(p++) | 0x20) == 'o')
                        && ((*(p++) | 0x20) == 'n')
                        && ((*(p++) | 0x20) == 't')
                        && ((*(p++) | 0x20) == 'a')
                        && ((*(p++) | 0x20) == 'c')
                        && ((*(p++) | 0x20) == 't')) {
                        if (p >= end) {
                            *src = p;
                            return ABNF_SIP_HDR_ACCEPT_CONTACT;
                        }
                    }
                    *src = p;
                    return ABNF_SIP_HDR_UNKNOWN;
                case 'l':
                    if ((src < end) && ((*(src++) | 0x20) == 'l') {
                        if ((src < end) && ((*(src++) | 0x20) == 'o') {
                            if ((src < end) && ((*(src++) | 0x20) == 'w') {
                                if (p >= end) {
                                    *src = p;
                                    return ABNF_SIP_HDR_ALLOW;
                                }
                                if ((p+6) >= end) {
                                    *src = p;
                                    return ABNF_SIP_HDR_UNKNOWN;
                                }
                                if ((*(p++) == '-')
                                    && ((*(p++) | 0x20) == 'e')
                                    && ((*(p++) | 0x20) == 'v')
                                    && ((*(p++) | 0x20) == 'e')
                                    && ((*(p++) | 0x20) == 'n')
                                    && ((*(p++) | 0x20) == 't')
                                    && ((*(p++) | 0x20) == 's')) {
                                    if (p >= end) {
                                        *src = p;
                                        return ABNF_SIP_HDR_ALLOW_EVENTS;
                                    }
                                }
                            }
                        }
                    }
                    *src = p;
                    return ABNF_SIP_HDR_UNKNOWN;
            }
            *src = p;
            return ABNF_SIP_HDR_UNKNOWN;
        case 'b':
            if (p >= end) {
                *src = p;
                return ABNF_SIP_HDR_REFERRED_BY;
            }
            *src = p;
            return ABNF_SIP_HDR_UNKNOWN;
        case 'c':
            if (p >= end) {
                *src = p;
                return ABNF_SIP_HDR_CONTENT_TYPE;
            }
            switch (*(p++) | 0x20) {
                case 'a':
                    if ((p+4) >= end) {
                        *src = p;
                        return ABNF_SIP_HDR_UNKNOWN;
                    }
                    if (((*(p++) | 0x20) == 'l')
                        && ((*(p++) | 0x20) == 'l')
                        && (*(p++) == '-')
                        && ((*(p++) | 0x20) == 'i')
                        && ((*(p++) | 0x20) == 'd')) {
                        if (p >= end) {
                            *src = p;
                            return ABNF_SIP_HDR_CALL_ID;
                        }
                    }
                    *src = p;
                    return ABNF_SIP_HDR_UNKNOWN;
                case 'o':
                    if ((src < end) && ((*(src++) | 0x20) == 'n') {
                        if ((src < end) && ((*(src++) | 0x20) == 't') {
                            switch (*(p++) | 0x20) {
                                case 'a':
                                    if ((p+1) >= end) {
                                        *src = p;
                                        return ABNF_SIP_HDR_UNKNOWN;
                                    }
                                    if (((*(p++) | 0x20) == 'c')
                                        && ((*(p++) | 0x20) == 't')) {
                                        if (p >= end) {
                                            *src = p;
                                            return ABNF_SIP_HDR_CONTACT;
                                        }
                                    }
                                    *src = p;
                                    return ABNF_SIP_HDR_UNKNOWN;
                                case 'e':
                                    if ((src < end) && ((*(src++) | 0x20) == 'n') {
                                        if ((src < end) && ((*(src++) | 0x20) == 't') {
                                            if ((p < end) && (*(p++) == '-') {
                                                switch (*(p++) | 0x20) {
                                                    case 'd':
                                                        if ((p+9) >= end) {
                                                            *src = p;
                                                            return ABNF_SIP_HDR_UNKNOWN;
                                                        }
                                                        if (((*(p++) | 0x20) == 'i')
                                                            && ((*(p++) | 0x20) == 's')
                                                            && ((*(p++) | 0x20) == 'p')
                                                            && ((*(p++) | 0x20) == 'o')
                                                            && ((*(p++) | 0x20) == 's')
                                                            && ((*(p++) | 0x20) == 'i')
                                                            && ((*(p++) | 0x20) == 't')
                                                            && ((*(p++) | 0x20) == 'i')
                                                            && ((*(p++) | 0x20) == 'o')
                                                            && ((*(p++) | 0x20) == 'n')) {
                                                            if (p >= end) {
                                                                *src = p;
                                                                return ABNF_SIP_HDR_CONTENT_DISPOSITION;
                                                            }
                                                        }
                                                        *src = p;
                                                        return ABNF_SIP_HDR_UNKNOWN;
                                                    case 'e':
                                                        if ((p+6) >= end) {
                                                            *src = p;
                                                            return ABNF_SIP_HDR_UNKNOWN;
                                                        }
                                                        if (((*(p++) | 0x20) == 'n')
                                                            && ((*(p++) | 0x20) == 'c')
                                                            && ((*(p++) | 0x20) == 'o')
                                                            && ((*(p++) | 0x20) == 'd')
                                                            && ((*(p++) | 0x20) == 'i')
                                                            && ((*(p++) | 0x20) == 'n')
                                                            && ((*(p++) | 0x20) == 'g')) {
                                                            if (p >= end) {
                                                                *src = p;
                                                                return ABNF_SIP_HDR_CONTENT_ENCODING;
                                                            }
                                                        }
                                                        *src = p;
                                                        return ABNF_SIP_HDR_UNKNOWN;
                                                    case 'l':
                                                        if ((p+4) >= end) {
                                                            *src = p;
                                                            return ABNF_SIP_HDR_UNKNOWN;
                                                        }
                                                        if (((*(p++) | 0x20) == 'e')
                                                            && ((*(p++) | 0x20) == 'n')
                                                            && ((*(p++) | 0x20) == 'g')
                                                            && ((*(p++) | 0x20) == 't')
                                                            && ((*(p++) | 0x20) == 'h')) {
                                                            if (p >= end) {
                                                                *src = p;
                                                                return ABNF_SIP_HDR_CONTENT_LENGTH;
                                                            }
                                                        }
                                                        *src = p;
                                                        return ABNF_SIP_HDR_UNKNOWN;
                                                    case 't':
                                                        if ((p+2) >= end) {
                                                            *src = p;
                                                            return ABNF_SIP_HDR_UNKNOWN;
                                                        }
                                                        if (((*(p++) | 0x20) == 'y')
                                                            && ((*(p++) | 0x20) == 'p')
                                                            && ((*(p++) | 0x20) == 'e')) {
                                                            if (p >= end) {
                                                                *src = p;
                                                                return ABNF_SIP_HDR_CONTENT_TYPE;
                                                            }
                                                        }
                                                        *src = p;
                                                        return ABNF_SIP_HDR_UNKNOWN;
                                                }
                                            }
                                        }
                                    }
                                    *src = p;
                                    return ABNF_SIP_HDR_UNKNOWN;
                            }
                        }
                    }
                    *src = p;
                    return ABNF_SIP_HDR_UNKNOWN;
                case 's':
                    if ((p+1) >= end) {
                        *src = p;
                        return ABNF_SIP_HDR_UNKNOWN;
                    }
                    if (((*(p++) | 0x20) == 'e')
                        && ((*(p++) | 0x20) == 'q')) {
                        if (p >= end) {
                            *src = p;
                            return ABNF_SIP_HDR_CSEQ;
                        }
                    }
                    *src = p;
                    return ABNF_SIP_HDR_UNKNOWN;
            }
            *src = p;
            return ABNF_SIP_HDR_UNKNOWN;
        case 'd':
            if ((p+2) >= end) {
                *src = p;
                return ABNF_SIP_HDR_UNKNOWN;
            }
            if (((*(p++) | 0x20) == 'a')
                && ((*(p++) | 0x20) == 't')
                && ((*(p++) | 0x20) == 'e')) {
                if (p >= end) {
                    *src = p;
                    return ABNF_SIP_HDR_DATE;
                }
            }
            *src = p;
            return ABNF_SIP_HDR_UNKNOWN;
        case 'e':
            if (p >= end) {
                *src = p;
                return ABNF_SIP_HDR_CONTENT_ENCODING;
            }
            if ((p+3) >= end) {
                *src = p;
                return ABNF_SIP_HDR_UNKNOWN;
            }
            if (((*(p++) | 0x20) == 'v')
                && ((*(p++) | 0x20) == 'e')
                && ((*(p++) | 0x20) == 'n')
                && ((*(p++) | 0x20) == 't')) {
                if (p >= end) {
                    *src = p;
                    return ABNF_SIP_HDR_EVENT;
                }
            }
            *src = p;
            return ABNF_SIP_HDR_UNKNOWN;
        case 'f':
            if (p >= end) {
                *src = p;
                return ABNF_SIP_HDR_FROM;
            }
            if ((p+2) >= end) {
                *src = p;
                return ABNF_SIP_HDR_UNKNOWN;
            }
            if (((*(p++) | 0x20) == 'r')
                && ((*(p++) | 0x20) == 'o')
                && ((*(p++) | 0x20) == 'm')) {
                if (p >= end) {
                    *src = p;
                    return ABNF_SIP_HDR_FROM;
                }
            }
            *src = p;
            return ABNF_SIP_HDR_UNKNOWN;
        case 'i':
            if (p >= end) {
                *src = p;
                return ABNF_SIP_HDR_CALL_ID;
            }
            *src = p;
            return ABNF_SIP_HDR_UNKNOWN;
        case 'j':
            if (p >= end) {
                *src = p;
                return ABNF_SIP_HDR_REJECT_CONTACT;
            }
            *src = p;
            return ABNF_SIP_HDR_UNKNOWN;
        case 'k':
            if (p >= end) {
                *src = p;
                return ABNF_SIP_HDR_SUPPORTED;
            }
            *src = p;
            return ABNF_SIP_HDR_UNKNOWN;
        case 'l':
            if (p >= end) {
                *src = p;
                return ABNF_SIP_HDR_CONTENT_LENGTH;
            }
            *src = p;
            return ABNF_SIP_HDR_UNKNOWN;
        case 'm':
            if (p >= end) {
                *src = p;
                return ABNF_SIP_HDR_CONTACT;
            }
            switch (*(p++) | 0x20) {
                case 'a':
                    if ((p+9) >= end) {
                        *src = p;
                        return ABNF_SIP_HDR_UNKNOWN;
                    }
                    if (((*(p++) | 0x20) == 'x')
                        && (*(p++) == '-')
                        && ((*(p++) | 0x20) == 'f')
                        && ((*(p++) | 0x20) == 'o')
                        && ((*(p++) | 0x20) == 'r')
                        && ((*(p++) | 0x20) == 'w')
                        && ((*(p++) | 0x20) == 'a')
                        && ((*(p++) | 0x20) == 'r')
                        && ((*(p++) | 0x20) == 'd')
                        && ((*(p++) | 0x20) == 's')) {
                        if (p >= end) {
                            *src = p;
                            return ABNF_SIP_HDR_MAX_FORWARDS;
                        }
                    }
                    *src = p;
                    return ABNF_SIP_HDR_UNKNOWN;
                case 'i':
                    if ((p+9) >= end) {
                        *src = p;
                        return ABNF_SIP_HDR_UNKNOWN;
                    }
                    if (((*(p++) | 0x20) == 'm')
                        && ((*(p++) | 0x20) == 'e')
                        && (*(p++) == '-')
                        && ((*(p++) | 0x20) == 'v')
                        && ((*(p++) | 0x20) == 'e')
                        && ((*(p++) | 0x20) == 'r')
                        && ((*(p++) | 0x20) == 's')
                        && ((*(p++) | 0x20) == 'i')
                        && ((*(p++) | 0x20) == 'o')
                        && ((*(p++) | 0x20) == 'n')) {
                        if (p >= end) {
                            *src = p;
                            return ABNF_SIP_HDR_MIME_VERSION;
                        }
                    }
                    *src = p;
                    return ABNF_SIP_HDR_UNKNOWN;
            }
            *src = p;
            return ABNF_SIP_HDR_UNKNOWN;
        case 'o':
            if (p >= end) {
                *src = p;
                return ABNF_SIP_HDR_EVENT;
            }
            *src = p;
            return ABNF_SIP_HDR_UNKNOWN;
        case 'r':
            if (p >= end) {
                *src = p;
                return ABNF_SIP_HDR_REFER_TO;
            }
            switch (*(p++) | 0x20) {
                case 'e':
                    switch (*(p++) | 0x20) {
                        case 'c':
                            if ((p+8) >= end) {
                                *src = p;
                                return ABNF_SIP_HDR_UNKNOWN;
                            }
                            if (((*(p++) | 0x20) == 'o')
                                && ((*(p++) | 0x20) == 'r')
                                && ((*(p++) | 0x20) == 'd')
                                && (*(p++) == '-')
                                && ((*(p++) | 0x20) == 'r')
                                && ((*(p++) | 0x20) == 'o')
                                && ((*(p++) | 0x20) == 'u')
                                && ((*(p++) | 0x20) == 't')
                                && ((*(p++) | 0x20) == 'e')) {
                                if (p >= end) {
                                    *src = p;
                                    return ABNF_SIP_HDR_RECORD_ROUTE;
                                }
                            }
                            *src = p;
                            return ABNF_SIP_HDR_UNKNOWN;
                        case 'f':
                            if ((src < end) && ((*(src++) | 0x20) == 'e') {
                                if ((src < end) && ((*(src++) | 0x20) == 'r') {
                                    switch (*(p++) | 0x20) {
                                        case '-':
                                            if ((p+1) >= end) {
                                                *src = p;
                                                return ABNF_SIP_HDR_UNKNOWN;
                                            }
                                            if (((*(p++) | 0x20) == 't')
                                                && ((*(p++) | 0x20) == 'o')) {
                                                if (p >= end) {
                                                    *src = p;
                                                    return ABNF_SIP_HDR_REFER_TO;
                                                }
                                            }
                                            *src = p;
                                            return ABNF_SIP_HDR_UNKNOWN;
                                        case 'r':
                                            if ((p+4) >= end) {
                                                *src = p;
                                                return ABNF_SIP_HDR_UNKNOWN;
                                            }
                                            if (((*(p++) | 0x20) == 'e')
                                                && ((*(p++) | 0x20) == 'd')
                                                && (*(p++) == '-')
                                                && ((*(p++) | 0x20) == 'b')
                                                && ((*(p++) | 0x20) == 'y')) {
                                                if (p >= end) {
                                                    *src = p;
                                                    return ABNF_SIP_HDR_REFERRED_BY;
                                                }
                                            }
                                            *src = p;
                                            return ABNF_SIP_HDR_UNKNOWN;
                                    }
                                }
                            }
                            *src = p;
                            return ABNF_SIP_HDR_UNKNOWN;
                        case 'j':
                            if ((p+10) >= end) {
                                *src = p;
                                return ABNF_SIP_HDR_UNKNOWN;
                            }
                            if (((*(p++) | 0x20) == 'e')
                                && ((*(p++) | 0x20) == 'c')
                                && ((*(p++) | 0x20) == 't')
                                && (*(p++) == '-')
                                && ((*(p++) | 0x20) == 'c')
                                && ((*(p++) | 0x20) == 'o')
                                && ((*(p++) | 0x20) == 'n')
                                && ((*(p++) | 0x20) == 't')
                                && ((*(p++) | 0x20) == 'a')
                                && ((*(p++) | 0x20) == 'c')
                                && ((*(p++) | 0x20) == 't')) {
                                if (p >= end) {
                                    *src = p;
                                    return ABNF_SIP_HDR_REJECT_CONTACT;
                                }
                            }
                            *src = p;
                            return ABNF_SIP_HDR_UNKNOWN;
                        case 'q':
                            if ((p+15) >= end) {
                                *src = p;
                                return ABNF_SIP_HDR_UNKNOWN;
                            }
                            if (((*(p++) | 0x20) == 'u')
                                && ((*(p++) | 0x20) == 'e')
                                && ((*(p++) | 0x20) == 's')
                                && ((*(p++) | 0x20) == 't')
                                && (*(p++) == '-')
                                && ((*(p++) | 0x20) == 'd')
                                && ((*(p++) | 0x20) == 'i')
                                && ((*(p++) | 0x20) == 's')
                                && ((*(p++) | 0x20) == 'p')
                                && ((*(p++) | 0x20) == 'o')
                                && ((*(p++) | 0x20) == 's')
                                && ((*(p++) | 0x20) == 'i')
                                && ((*(p++) | 0x20) == 't')
                                && ((*(p++) | 0x20) == 'i')
                                && ((*(p++) | 0x20) == 'o')
                                && ((*(p++) | 0x20) == 'n')) {
                                if (p >= end) {
                                    *src = p;
                                    return ABNF_SIP_HDR_REQUEST_DISPOSITION;
                                }
                            }
                            *src = p;
                            return ABNF_SIP_HDR_UNKNOWN;
                    }
                    *src = p;
                    return ABNF_SIP_HDR_UNKNOWN;
                case 'o':
                    if ((p+2) >= end) {
                        *src = p;
                        return ABNF_SIP_HDR_UNKNOWN;
                    }
                    if (((*(p++) | 0x20) == 'u')
                        && ((*(p++) | 0x20) == 't')
                        && ((*(p++) | 0x20) == 'e')) {
                        if (p >= end) {
                            *src = p;
                            return ABNF_SIP_HDR_ROUTE;
                        }
                    }
                    *src = p;
                    return ABNF_SIP_HDR_UNKNOWN;
            }
            *src = p;
            return ABNF_SIP_HDR_UNKNOWN;
        case 's':
            if (p >= end) {
                *src = p;
                return ABNF_SIP_HDR_SUBJECT;
            }
            switch (*(p++) | 0x20) {
                case 'e':
                    if ((p+12) >= end) {
                        *src = p;
                        return ABNF_SIP_HDR_UNKNOWN;
                    }
                    if (((*(p++) | 0x20) == 's')
                        && ((*(p++) | 0x20) == 's')
                        && ((*(p++) | 0x20) == 'i')
                        && ((*(p++) | 0x20) == 'o')
                        && ((*(p++) | 0x20) == 'n')
                        && (*(p++) == '-')
                        && ((*(p++) | 0x20) == 'e')
                        && ((*(p++) | 0x20) == 'x')
                        && ((*(p++) | 0x20) == 'p')
                        && ((*(p++) | 0x20) == 'i')
                        && ((*(p++) | 0x20) == 'r')
                        && ((*(p++) | 0x20) == 'e')
                        && ((*(p++) | 0x20) == 's')) {
                        if (p >= end) {
                            *src = p;
                            return ABNF_SIP_HDR_SESSION_EXPIRES;
                        }
                    }
                    *src = p;
                    return ABNF_SIP_HDR_UNKNOWN;
                case 'u':
                    switch (*(p++) | 0x20) {
                        case 'b':
                            if ((p+3) >= end) {
                                *src = p;
                                return ABNF_SIP_HDR_UNKNOWN;
                            }
                            if (((*(p++) | 0x20) == 'j')
                                && ((*(p++) | 0x20) == 'e')
                                && ((*(p++) | 0x20) == 'c')
                                && ((*(p++) | 0x20) == 't')) {
                                if (p >= end) {
                                    *src = p;
                                    return ABNF_SIP_HDR_SUBJECT;
                                }
                            }
                            *src = p;
                            return ABNF_SIP_HDR_UNKNOWN;
                        case 'p':
                            if ((p+5) >= end) {
                                *src = p;
                                return ABNF_SIP_HDR_UNKNOWN;
                            }
                            if (((*(p++) | 0x20) == 'p')
                                && ((*(p++) | 0x20) == 'o')
                                && ((*(p++) | 0x20) == 'r')
                                && ((*(p++) | 0x20) == 't')
                                && ((*(p++) | 0x20) == 'e')
                                && ((*(p++) | 0x20) == 'd')) {
                                if (p >= end) {
                                    *src = p;
                                    return ABNF_SIP_HDR_SUPPORTED;
                                }
                            }
                            *src = p;
                            return ABNF_SIP_HDR_UNKNOWN;
                    }
                    *src = p;
                    return ABNF_SIP_HDR_UNKNOWN;
            }
            *src = p;
            return ABNF_SIP_HDR_UNKNOWN;
        case 't':
            if (p >= end) {
                *src = p;
                return ABNF_SIP_HDR_TO;
            }
            if ((src < end) && ((*(src++) | 0x20) == 'o') {
                if (p >= end) {
                    *src = p;
                    return ABNF_SIP_HDR_TO;
                }
            }
            *src = p;
            return ABNF_SIP_HDR_UNKNOWN;
        case 'u':
            if (p >= end) {
                *src = p;
                return ABNF_SIP_HDR_ALLOW_EVENTS;
            }
            *src = p;
            return ABNF_SIP_HDR_UNKNOWN;
        case 'v':
            if (p >= end) {
                *src = p;
                return ABNF_SIP_HDR_VIA;
            }
            if ((p+1) >= end) {
                *src = p;
                return ABNF_SIP_HDR_UNKNOWN;
            }
            if (((*(p++) | 0x20) == 'i')
                && ((*(p++) | 0x20) == 'a')) {
                if (p >= end) {
                    *src = p;
                    return ABNF_SIP_HDR_VIA;
                }
            }
            *src = p;
            return ABNF_SIP_HDR_UNKNOWN;
        case 'x':
            if (p >= end) {
                *src = p;
                return ABNF_SIP_HDR_SESSION_EXPIRES;
            }
            *src = p;
            return ABNF_SIP_HDR_UNKNOWN;
    }

    return ABNF_SIP_HDR_UNKNOWN;
}
