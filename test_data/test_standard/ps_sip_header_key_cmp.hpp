#ifndef PS_SIP_HEADER_KEY_CMP_HPP
#define PS_SIP_HEADER_KEY_CMP_HPP

/*---------------- index definition ----------------*/
const unsigned int  ABNF_SIP_HDR_UNKNOWN                = 0;
const unsigned int  ABNF_SIP_HDR_FROM                   = 1;
const unsigned int  ABNF_SIP_HDR_TO                     = 2;
const unsigned int  ABNF_SIP_HDR_VIA                    = 3;
const unsigned int  ABNF_SIP_HDR_CALL_ID                = 4;
const unsigned int  ABNF_SIP_HDR_CSEQ                   = 5;
const unsigned int  ABNF_SIP_HDR_CONTENT_LENGTH         = 6;
const unsigned int  ABNF_SIP_HDR_CONTENT_TYPE           = 7;
const unsigned int  ABNF_SIP_HDR_CONTACT                = 8;
const unsigned int  ABNF_SIP_HDR_MAX_FORWARDS           = 9;
const unsigned int  ABNF_SIP_HDR_ROUTE                  = 10;
const unsigned int  ABNF_SIP_HDR_RECORD_ROUTE           = 11;
const unsigned int  ABNF_SIP_HDR_CONTENT_DISPOSITION    = 12;
const unsigned int  ABNF_SIP_HDR_ALLOW                  = 13;
const unsigned int  ABNF_SIP_HDR_CONTENT_ENCODING       = 14;
const unsigned int  ABNF_SIP_HDR_DATE                   = 15;
const unsigned int  ABNF_SIP_HDR_SUBJECT                = 16;
const unsigned int  ABNF_SIP_HDR_SUPPORTED              = 17;
const unsigned int  ABNF_SIP_HDR_ALLOW_EVENTS           = 18;
const unsigned int  ABNF_SIP_HDR_EVENT                  = 19;
const unsigned int  ABNF_SIP_HDR_REFER_TO               = 20;
const unsigned int  ABNF_SIP_HDR_ACCEPT_CONTACT         = 21;
const unsigned int  ABNF_SIP_HDR_REJECT_CONTACT         = 22;
const unsigned int  ABNF_SIP_HDR_REQUEST_DISPOSITION    = 23;
const unsigned int  ABNF_SIP_HDR_REFERRED_BY            = 24;
const unsigned int  ABNF_SIP_HDR_SESSION_EXPIRES        = 25;
const unsigned int  ABNF_SIP_HDR_MIME_VERSION           = 26;

/*---------------- action declaration ----------------*/
unsigned int GetSipHeaderIndex(char const** src, char const* end);

#endif /* PS_SIP_HEADER_KEY_CMP_HPP */
