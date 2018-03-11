#ifndef PS_SIP_HEADER_KEY_CMP_2_H
#define PS_SIP_HEADER_KEY_CMP_2_H

/*---------------- index definition ----------------*/
#define ABNF_SIP_HDR_UNKNOWN                ((unsigned int)(0))
#define ABNF_SIP_HDR_FROM                   ((unsigned int)(1))
#define ABNF_SIP_HDR_TO                     ((unsigned int)(2))
#define ABNF_SIP_HDR_VIA                    ((unsigned int)(3))
#define ABNF_SIP_HDR_CALL_ID                ((unsigned int)(4))
#define ABNF_SIP_HDR_CSEQ                   ((unsigned int)(5))
#define ABNF_SIP_HDR_CONTENT_LENGTH         ((unsigned int)(6))
#define ABNF_SIP_HDR_CONTENT_TYPE           ((unsigned int)(7))
#define ABNF_SIP_HDR_CONTACT                ((unsigned int)(8))
#define ABNF_SIP_HDR_MAX_FORWARDS           ((unsigned int)(9))
#define ABNF_SIP_HDR_ROUTE                  ((unsigned int)(10))
#define ABNF_SIP_HDR_RECORD_ROUTE           ((unsigned int)(11))
#define ABNF_SIP_HDR_CONTENT_DISPOSITION    ((unsigned int)(12))
#define ABNF_SIP_HDR_ALLOW                  ((unsigned int)(13))
#define ABNF_SIP_HDR_CONTENT_ENCODING       ((unsigned int)(14))
#define ABNF_SIP_HDR_DATE                   ((unsigned int)(15))
#define ABNF_SIP_HDR_SUBJECT                ((unsigned int)(16))
#define ABNF_SIP_HDR_SUPPORTED              ((unsigned int)(17))
#define ABNF_SIP_HDR_ALLOW_EVENTS           ((unsigned int)(18))
#define ABNF_SIP_HDR_EVENT                  ((unsigned int)(19))
#define ABNF_SIP_HDR_REFER_TO               ((unsigned int)(20))
#define ABNF_SIP_HDR_ACCEPT_CONTACT         ((unsigned int)(21))
#define ABNF_SIP_HDR_REJECT_CONTACT         ((unsigned int)(22))
#define ABNF_SIP_HDR_REQUEST_DISPOSITION    ((unsigned int)(23))
#define ABNF_SIP_HDR_REFERRED_BY            ((unsigned int)(24))
#define ABNF_SIP_HDR_SESSION_EXPIRES        ((unsigned int)(25))
#define ABNF_SIP_HDR_MIME_VERSION           ((unsigned int)(26))

/*---------------- action declaration ----------------*/
unsigned int GetSipHeaderIndex(char const** src, char const* end);

#endif /* PS_SIP_HEADER_KEY_CMP_2_H */
