#include "ps_sip_header_key_cmp.hpp"

unsigned int GetSipHeaderIndex(char const** src, char const* end)
{
    char const*  p = *src;
    
    if (p == NULL || p >= end)
    {
        return ABNF_SIP_HDR_UNKNOWN;
    }

    switch (*(p++))
    {
        case 'A':
            switch (*(p++))
            {
                case 'c':
                    if ((p + 11) >= end)
                    {
                        *src = p;
                        return ABNF_SIP_HDR_UNKNOWN;
                    }
                    if ((*(p++) == 'c')
                        && (*(p++) == 'e')
                        && (*(p++) == 'p')
                        && (*(p++) == 't')
                        && (*(p++) == '-')
                        && (*(p++) == 'C')
                        && (*(p++) == 'o')
                        && (*(p++) == 'n')
                        && (*(p++) == 't')
                        && (*(p++) == 'a')
                        && (*(p++) == 'c')
                        && (*(p++) == 't'))
                    {
                        if (p >= end)
                        {
                            *src = p;
                            return ABNF_SIP_HDR_ACCEPT_CONTACT;
                        }
                    }
                    *src = p;
                    return ABNF_SIP_HDR_UNKNOWN;
                case 'l':
                    if ((p < end) && (*(p++) == 'l')
                    {
                        if ((p < end) && (*(p++) == 'o')
                        {
                            if ((p < end) && (*(p++) == 'w')
                            {
                                if (p >= end)
                                {
                                    *src = p;
                                    return ABNF_SIP_HDR_ALLOW;
                                }
                                if ((p + 6) >= end)
                                {
                                    *src = p;
                                    return ABNF_SIP_HDR_UNKNOWN;
                                }
                                if ((*(p++) == '-')
                                    && (*(p++) == 'E')
                                    && (*(p++) == 'v')
                                    && (*(p++) == 'e')
                                    && (*(p++) == 'n')
                                    && (*(p++) == 't')
                                    && (*(p++) == 's'))
                                {
                                    if (p >= end)
                                    {
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
        case 'C':
            switch (*(p++))
            {
                case 'S':
                    if ((p + 1) >= end)
                    {
                        *src = p;
                        return ABNF_SIP_HDR_UNKNOWN;
                    }
                    if ((*(p++) == 'e')
                        && (*(p++) == 'q'))
                    {
                        if (p >= end)
                        {
                            *src = p;
                            return ABNF_SIP_HDR_CSEQ;
                        }
                    }
                    *src = p;
                    return ABNF_SIP_HDR_UNKNOWN;
                case 'a':
                    if ((p + 4) >= end)
                    {
                        *src = p;
                        return ABNF_SIP_HDR_UNKNOWN;
                    }
                    if ((*(p++) == 'l')
                        && (*(p++) == 'l')
                        && (*(p++) == '-')
                        && (*(p++) == 'I')
                        && (*(p++) == 'D'))
                    {
                        if (p >= end)
                        {
                            *src = p;
                            return ABNF_SIP_HDR_CALL_ID;
                        }
                    }
                    *src = p;
                    return ABNF_SIP_HDR_UNKNOWN;
                case 'o':
                    if ((p < end) && (*(p++) == 'n')
                    {
                        if ((p < end) && (*(p++) == 't')
                        {
                            switch (*(p++))
                            {
                                case 'a':
                                    if ((p + 1) >= end)
                                    {
                                        *src = p;
                                        return ABNF_SIP_HDR_UNKNOWN;
                                    }
                                    if ((*(p++) == 'c')
                                        && (*(p++) == 't'))
                                    {
                                        if (p >= end)
                                        {
                                            *src = p;
                                            return ABNF_SIP_HDR_CONTACT;
                                        }
                                    }
                                    *src = p;
                                    return ABNF_SIP_HDR_UNKNOWN;
                                case 'e':
                                    if ((p < end) && (*(p++) == 'n')
                                    {
                                        if ((p < end) && (*(p++) == 't')
                                        {
                                            if ((p < end) && (*(p++) == '-')
                                            {
                                                switch (*(p++))
                                                {
                                                    case 'D':
                                                        if ((p + 9) >= end)
                                                        {
                                                            *src = p;
                                                            return ABNF_SIP_HDR_UNKNOWN;
                                                        }
                                                        if ((*(p++) == 'i')
                                                            && (*(p++) == 's')
                                                            && (*(p++) == 'p')
                                                            && (*(p++) == 'o')
                                                            && (*(p++) == 's')
                                                            && (*(p++) == 'i')
                                                            && (*(p++) == 't')
                                                            && (*(p++) == 'i')
                                                            && (*(p++) == 'o')
                                                            && (*(p++) == 'n'))
                                                        {
                                                            if (p >= end)
                                                            {
                                                                *src = p;
                                                                return ABNF_SIP_HDR_CONTENT_DISPOSITION;
                                                            }
                                                        }
                                                        *src = p;
                                                        return ABNF_SIP_HDR_UNKNOWN;
                                                    case 'E':
                                                        if ((p + 6) >= end)
                                                        {
                                                            *src = p;
                                                            return ABNF_SIP_HDR_UNKNOWN;
                                                        }
                                                        if ((*(p++) == 'n')
                                                            && (*(p++) == 'c')
                                                            && (*(p++) == 'o')
                                                            && (*(p++) == 'd')
                                                            && (*(p++) == 'i')
                                                            && (*(p++) == 'n')
                                                            && (*(p++) == 'g'))
                                                        {
                                                            if (p >= end)
                                                            {
                                                                *src = p;
                                                                return ABNF_SIP_HDR_CONTENT_ENCODING;
                                                            }
                                                        }
                                                        *src = p;
                                                        return ABNF_SIP_HDR_UNKNOWN;
                                                    case 'L':
                                                        if ((p + 4) >= end)
                                                        {
                                                            *src = p;
                                                            return ABNF_SIP_HDR_UNKNOWN;
                                                        }
                                                        if ((*(p++) == 'e')
                                                            && (*(p++) == 'n')
                                                            && (*(p++) == 'g')
                                                            && (*(p++) == 't')
                                                            && (*(p++) == 'h'))
                                                        {
                                                            if (p >= end)
                                                            {
                                                                *src = p;
                                                                return ABNF_SIP_HDR_CONTENT_LENGTH;
                                                            }
                                                        }
                                                        *src = p;
                                                        return ABNF_SIP_HDR_UNKNOWN;
                                                    case 'T':
                                                        if ((p + 2) >= end)
                                                        {
                                                            *src = p;
                                                            return ABNF_SIP_HDR_UNKNOWN;
                                                        }
                                                        if ((*(p++) == 'y')
                                                            && (*(p++) == 'p')
                                                            && (*(p++) == 'e'))
                                                        {
                                                            if (p >= end)
                                                            {
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
            }
            *src = p;
            return ABNF_SIP_HDR_UNKNOWN;
        case 'D':
            if ((p + 2) >= end)
            {
                *src = p;
                return ABNF_SIP_HDR_UNKNOWN;
            }
            if ((*(p++) == 'a')
                && (*(p++) == 't')
                && (*(p++) == 'e'))
            {
                if (p >= end)
                {
                    *src = p;
                    return ABNF_SIP_HDR_DATE;
                }
            }
            *src = p;
            return ABNF_SIP_HDR_UNKNOWN;
        case 'E':
            if ((p + 3) >= end)
            {
                *src = p;
                return ABNF_SIP_HDR_UNKNOWN;
            }
            if ((*(p++) == 'v')
                && (*(p++) == 'e')
                && (*(p++) == 'n')
                && (*(p++) == 't'))
            {
                if (p >= end)
                {
                    *src = p;
                    return ABNF_SIP_HDR_EVENT;
                }
            }
            *src = p;
            return ABNF_SIP_HDR_UNKNOWN;
        case 'F':
            if ((p + 2) >= end)
            {
                *src = p;
                return ABNF_SIP_HDR_UNKNOWN;
            }
            if ((*(p++) == 'r')
                && (*(p++) == 'o')
                && (*(p++) == 'm'))
            {
                if (p >= end)
                {
                    *src = p;
                    return ABNF_SIP_HDR_FROM;
                }
            }
            *src = p;
            return ABNF_SIP_HDR_UNKNOWN;
        case 'M':
            switch (*(p++))
            {
                case 'I':
                    if ((p + 9) >= end)
                    {
                        *src = p;
                        return ABNF_SIP_HDR_UNKNOWN;
                    }
                    if ((*(p++) == 'M')
                        && (*(p++) == 'E')
                        && (*(p++) == '-')
                        && (*(p++) == 'V')
                        && (*(p++) == 'e')
                        && (*(p++) == 'r')
                        && (*(p++) == 's')
                        && (*(p++) == 'i')
                        && (*(p++) == 'o')
                        && (*(p++) == 'n'))
                    {
                        if (p >= end)
                        {
                            *src = p;
                            return ABNF_SIP_HDR_MIME_VERSION;
                        }
                    }
                    *src = p;
                    return ABNF_SIP_HDR_UNKNOWN;
                case 'a':
                    if ((p + 9) >= end)
                    {
                        *src = p;
                        return ABNF_SIP_HDR_UNKNOWN;
                    }
                    if ((*(p++) == 'x')
                        && (*(p++) == '-')
                        && (*(p++) == 'F')
                        && (*(p++) == 'o')
                        && (*(p++) == 'r')
                        && (*(p++) == 'w')
                        && (*(p++) == 'a')
                        && (*(p++) == 'r')
                        && (*(p++) == 'd')
                        && (*(p++) == 's'))
                    {
                        if (p >= end)
                        {
                            *src = p;
                            return ABNF_SIP_HDR_MAX_FORWARDS;
                        }
                    }
                    *src = p;
                    return ABNF_SIP_HDR_UNKNOWN;
            }
            *src = p;
            return ABNF_SIP_HDR_UNKNOWN;
        case 'R':
            switch (*(p++))
            {
                case 'e':
                    switch (*(p++))
                    {
                        case 'c':
                            if ((p + 8) >= end)
                            {
                                *src = p;
                                return ABNF_SIP_HDR_UNKNOWN;
                            }
                            if ((*(p++) == 'o')
                                && (*(p++) == 'r')
                                && (*(p++) == 'd')
                                && (*(p++) == '-')
                                && (*(p++) == 'R')
                                && (*(p++) == 'o')
                                && (*(p++) == 'u')
                                && (*(p++) == 't')
                                && (*(p++) == 'e'))
                            {
                                if (p >= end)
                                {
                                    *src = p;
                                    return ABNF_SIP_HDR_RECORD_ROUTE;
                                }
                            }
                            *src = p;
                            return ABNF_SIP_HDR_UNKNOWN;
                        case 'f':
                            if ((p < end) && (*(p++) == 'e')
                            {
                                if ((p < end) && (*(p++) == 'r')
                                {
                                    switch (*(p++))
                                    {
                                        case '-':
                                            if ((p + 1) >= end)
                                            {
                                                *src = p;
                                                return ABNF_SIP_HDR_UNKNOWN;
                                            }
                                            if ((*(p++) == 'T')
                                                && (*(p++) == 'o'))
                                            {
                                                if (p >= end)
                                                {
                                                    *src = p;
                                                    return ABNF_SIP_HDR_REFER_TO;
                                                }
                                            }
                                            *src = p;
                                            return ABNF_SIP_HDR_UNKNOWN;
                                        case 'r':
                                            if ((p + 4) >= end)
                                            {
                                                *src = p;
                                                return ABNF_SIP_HDR_UNKNOWN;
                                            }
                                            if ((*(p++) == 'e')
                                                && (*(p++) == 'd')
                                                && (*(p++) == '-')
                                                && (*(p++) == 'B')
                                                && (*(p++) == 'y'))
                                            {
                                                if (p >= end)
                                                {
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
                            if ((p + 10) >= end)
                            {
                                *src = p;
                                return ABNF_SIP_HDR_UNKNOWN;
                            }
                            if ((*(p++) == 'e')
                                && (*(p++) == 'c')
                                && (*(p++) == 't')
                                && (*(p++) == '-')
                                && (*(p++) == 'C')
                                && (*(p++) == 'o')
                                && (*(p++) == 'n')
                                && (*(p++) == 't')
                                && (*(p++) == 'a')
                                && (*(p++) == 'c')
                                && (*(p++) == 't'))
                            {
                                if (p >= end)
                                {
                                    *src = p;
                                    return ABNF_SIP_HDR_REJECT_CONTACT;
                                }
                            }
                            *src = p;
                            return ABNF_SIP_HDR_UNKNOWN;
                        case 'q':
                            if ((p + 15) >= end)
                            {
                                *src = p;
                                return ABNF_SIP_HDR_UNKNOWN;
                            }
                            if ((*(p++) == 'u')
                                && (*(p++) == 'e')
                                && (*(p++) == 's')
                                && (*(p++) == 't')
                                && (*(p++) == '-')
                                && (*(p++) == 'D')
                                && (*(p++) == 'i')
                                && (*(p++) == 's')
                                && (*(p++) == 'p')
                                && (*(p++) == 'o')
                                && (*(p++) == 's')
                                && (*(p++) == 'i')
                                && (*(p++) == 't')
                                && (*(p++) == 'i')
                                && (*(p++) == 'o')
                                && (*(p++) == 'n'))
                            {
                                if (p >= end)
                                {
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
                    if ((p + 2) >= end)
                    {
                        *src = p;
                        return ABNF_SIP_HDR_UNKNOWN;
                    }
                    if ((*(p++) == 'u')
                        && (*(p++) == 't')
                        && (*(p++) == 'e'))
                    {
                        if (p >= end)
                        {
                            *src = p;
                            return ABNF_SIP_HDR_ROUTE;
                        }
                    }
                    *src = p;
                    return ABNF_SIP_HDR_UNKNOWN;
            }
            *src = p;
            return ABNF_SIP_HDR_UNKNOWN;
        case 'S':
            switch (*(p++))
            {
                case 'e':
                    if ((p + 12) >= end)
                    {
                        *src = p;
                        return ABNF_SIP_HDR_UNKNOWN;
                    }
                    if ((*(p++) == 's')
                        && (*(p++) == 's')
                        && (*(p++) == 'i')
                        && (*(p++) == 'o')
                        && (*(p++) == 'n')
                        && (*(p++) == '-')
                        && (*(p++) == 'E')
                        && (*(p++) == 'x')
                        && (*(p++) == 'p')
                        && (*(p++) == 'i')
                        && (*(p++) == 'r')
                        && (*(p++) == 'e')
                        && (*(p++) == 's'))
                    {
                        if (p >= end)
                        {
                            *src = p;
                            return ABNF_SIP_HDR_SESSION_EXPIRES;
                        }
                    }
                    *src = p;
                    return ABNF_SIP_HDR_UNKNOWN;
                case 'u':
                    switch (*(p++))
                    {
                        case 'b':
                            if ((p + 3) >= end)
                            {
                                *src = p;
                                return ABNF_SIP_HDR_UNKNOWN;
                            }
                            if ((*(p++) == 'j')
                                && (*(p++) == 'e')
                                && (*(p++) == 'c')
                                && (*(p++) == 't'))
                            {
                                if (p >= end)
                                {
                                    *src = p;
                                    return ABNF_SIP_HDR_SUBJECT;
                                }
                            }
                            *src = p;
                            return ABNF_SIP_HDR_UNKNOWN;
                        case 'p':
                            if ((p + 5) >= end)
                            {
                                *src = p;
                                return ABNF_SIP_HDR_UNKNOWN;
                            }
                            if ((*(p++) == 'p')
                                && (*(p++) == 'o')
                                && (*(p++) == 'r')
                                && (*(p++) == 't')
                                && (*(p++) == 'e')
                                && (*(p++) == 'd'))
                            {
                                if (p >= end)
                                {
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
        case 'T':
            if ((p < end) && (*(p++) == 'o')
            {
                if (p >= end)
                {
                    *src = p;
                    return ABNF_SIP_HDR_TO;
                }
            }
            *src = p;
            return ABNF_SIP_HDR_UNKNOWN;
        case 'V':
            if ((p + 1) >= end)
            {
                *src = p;
                return ABNF_SIP_HDR_UNKNOWN;
            }
            if ((*(p++) == 'i')
                && (*(p++) == 'a'))
            {
                if (p >= end)
                {
                    *src = p;
                    return ABNF_SIP_HDR_VIA;
                }
            }
            *src = p;
            return ABNF_SIP_HDR_UNKNOWN;
        case 'a':
            if (p >= end)
            {
                *src = p;
                return ABNF_SIP_HDR_ACCEPT_CONTACT;
            }
            *src = p;
            return ABNF_SIP_HDR_UNKNOWN;
        case 'b':
            if (p >= end)
            {
                *src = p;
                return ABNF_SIP_HDR_REFERRED_BY;
            }
            *src = p;
            return ABNF_SIP_HDR_UNKNOWN;
        case 'c':
            if (p >= end)
            {
                *src = p;
                return ABNF_SIP_HDR_CONTENT_TYPE;
            }
            *src = p;
            return ABNF_SIP_HDR_UNKNOWN;
        case 'e':
            if (p >= end)
            {
                *src = p;
                return ABNF_SIP_HDR_CONTENT_ENCODING;
            }
            *src = p;
            return ABNF_SIP_HDR_UNKNOWN;
        case 'f':
            if (p >= end)
            {
                *src = p;
                return ABNF_SIP_HDR_FROM;
            }
            *src = p;
            return ABNF_SIP_HDR_UNKNOWN;
        case 'i':
            if (p >= end)
            {
                *src = p;
                return ABNF_SIP_HDR_CALL_ID;
            }
            *src = p;
            return ABNF_SIP_HDR_UNKNOWN;
        case 'j':
            if (p >= end)
            {
                *src = p;
                return ABNF_SIP_HDR_REJECT_CONTACT;
            }
            *src = p;
            return ABNF_SIP_HDR_UNKNOWN;
        case 'k':
            if (p >= end)
            {
                *src = p;
                return ABNF_SIP_HDR_SUPPORTED;
            }
            *src = p;
            return ABNF_SIP_HDR_UNKNOWN;
        case 'l':
            if (p >= end)
            {
                *src = p;
                return ABNF_SIP_HDR_CONTENT_LENGTH;
            }
            *src = p;
            return ABNF_SIP_HDR_UNKNOWN;
        case 'm':
            if (p >= end)
            {
                *src = p;
                return ABNF_SIP_HDR_CONTACT;
            }
            *src = p;
            return ABNF_SIP_HDR_UNKNOWN;
        case 'o':
            if (p >= end)
            {
                *src = p;
                return ABNF_SIP_HDR_EVENT;
            }
            *src = p;
            return ABNF_SIP_HDR_UNKNOWN;
        case 'r':
            if (p >= end)
            {
                *src = p;
                return ABNF_SIP_HDR_REFER_TO;
            }
            *src = p;
            return ABNF_SIP_HDR_UNKNOWN;
        case 's':
            if (p >= end)
            {
                *src = p;
                return ABNF_SIP_HDR_SUBJECT;
            }
            *src = p;
            return ABNF_SIP_HDR_UNKNOWN;
        case 't':
            if (p >= end)
            {
                *src = p;
                return ABNF_SIP_HDR_TO;
            }
            *src = p;
            return ABNF_SIP_HDR_UNKNOWN;
        case 'u':
            if (p >= end)
            {
                *src = p;
                return ABNF_SIP_HDR_ALLOW_EVENTS;
            }
            *src = p;
            return ABNF_SIP_HDR_UNKNOWN;
        case 'v':
            if (p >= end)
            {
                *src = p;
                return ABNF_SIP_HDR_VIA;
            }
            *src = p;
            return ABNF_SIP_HDR_UNKNOWN;
        case 'x':
            if (p >= end)
            {
                *src = p;
                return ABNF_SIP_HDR_SESSION_EXPIRES;
            }
            *src = p;
            return ABNF_SIP_HDR_UNKNOWN;
    }

    return ABNF_SIP_HDR_UNKNOWN;
}
