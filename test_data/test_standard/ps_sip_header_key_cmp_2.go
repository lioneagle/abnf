package sip_header

/*---------------- index definition ----------------*/
const (
	ABNF_SIP_HDR_UNKNOWN             SipHeaderIndexType = 0
	ABNF_SIP_HDR_FROM                SipHeaderIndexType = 1
	ABNF_SIP_HDR_TO                  SipHeaderIndexType = 2
	ABNF_SIP_HDR_VIA                 SipHeaderIndexType = 3
	ABNF_SIP_HDR_CALL_ID             SipHeaderIndexType = 4
	ABNF_SIP_HDR_CSEQ                SipHeaderIndexType = 5
	ABNF_SIP_HDR_CONTENT_LENGTH      SipHeaderIndexType = 6
	ABNF_SIP_HDR_CONTENT_TYPE        SipHeaderIndexType = 7
	ABNF_SIP_HDR_CONTACT             SipHeaderIndexType = 8
	ABNF_SIP_HDR_MAX_FORWARDS        SipHeaderIndexType = 9
	ABNF_SIP_HDR_ROUTE               SipHeaderIndexType = 10
	ABNF_SIP_HDR_RECORD_ROUTE        SipHeaderIndexType = 11
	ABNF_SIP_HDR_CONTENT_DISPOSITION SipHeaderIndexType = 12
	ABNF_SIP_HDR_ALLOW               SipHeaderIndexType = 13
	ABNF_SIP_HDR_CONTENT_ENCODING    SipHeaderIndexType = 14
	ABNF_SIP_HDR_DATE                SipHeaderIndexType = 15
	ABNF_SIP_HDR_SUBJECT             SipHeaderIndexType = 16
	ABNF_SIP_HDR_SUPPORTED           SipHeaderIndexType = 17
	ABNF_SIP_HDR_ALLOW_EVENTS        SipHeaderIndexType = 18
	ABNF_SIP_HDR_EVENT               SipHeaderIndexType = 19
	ABNF_SIP_HDR_REFER_TO            SipHeaderIndexType = 20
	ABNF_SIP_HDR_ACCEPT_CONTACT      SipHeaderIndexType = 21
	ABNF_SIP_HDR_REJECT_CONTACT      SipHeaderIndexType = 22
	ABNF_SIP_HDR_REQUEST_DISPOSITION SipHeaderIndexType = 23
	ABNF_SIP_HDR_REFERRED_BY         SipHeaderIndexType = 24
	ABNF_SIP_HDR_SESSION_EXPIRES     SipHeaderIndexType = 25
	ABNF_SIP_HDR_MIME_VERSION        SipHeaderIndexType = 26
)

func GetSipHeaderIndex(src []byte) SipHeaderIndexType {
	pos := 0
	len1 := len(src)

	if pos >= len1 {
		return ABNF_SIP_HDR_UNKNOWN
	}

	switch src[pos] | 0x20 {
	case 'a':
		pos++
		if pos >= len1 {
			return ABNF_SIP_HDR_ACCEPT_CONTACT
		}
		switch src[pos] | 0x20 {
		case 'c':
			pos++
			if (pos + 11) >= len1 {
				return ABNF_SIP_HDR_UNKNOWN
			}
			if ((src[pos] | 0x20) == 'c') &&
				((src[pos+1] | 0x20) == 'e') &&
				((src[pos+2] | 0x20) == 'p') &&
				((src[pos+3] | 0x20) == 't') &&
				(src[pos+4] == '-') &&
				((src[pos+5] | 0x20) == 'c') &&
				((src[pos+6] | 0x20) == 'o') &&
				((src[pos+7] | 0x20) == 'n') &&
				((src[pos+8] | 0x20) == 't') &&
				((src[pos+9] | 0x20) == 'a') &&
				((src[pos+10] | 0x20) == 'c') &&
				((src[pos+11] | 0x20) == 't') {
				pos += 12
				if pos >= len1 {
					return ABNF_SIP_HDR_ACCEPT_CONTACT
				}
			}
			return ABNF_SIP_HDR_UNKNOWN
		case 'l':
			pos++
			if (pos < len1) && ((src[pos] | 0x20) == 'l') {
				pos++
				if (pos < len1) && ((src[pos] | 0x20) == 'o') {
					pos++
					if (pos < len1) && ((src[pos] | 0x20) == 'w') {
						pos++
						if pos >= len1 {
							return ABNF_SIP_HDR_ALLOW
						}
						if (pos + 6) >= len1 {
							return ABNF_SIP_HDR_UNKNOWN
						}
						if (src[pos] == '-') &&
							((src[pos+1] | 0x20) == 'e') &&
							((src[pos+2] | 0x20) == 'v') &&
							((src[pos+3] | 0x20) == 'e') &&
							((src[pos+4] | 0x20) == 'n') &&
							((src[pos+5] | 0x20) == 't') &&
							((src[pos+6] | 0x20) == 's') {
							pos += 7
							if pos >= len1 {
								return ABNF_SIP_HDR_ALLOW_EVENTS
							}
						}
					}
				}
			}
			return ABNF_SIP_HDR_UNKNOWN
		}
		return ABNF_SIP_HDR_UNKNOWN
	case 'b':
		pos++
		if pos >= len1 {
			return ABNF_SIP_HDR_REFERRED_BY
		}
		return ABNF_SIP_HDR_UNKNOWN
	case 'c':
		pos++
		if pos >= len1 {
			return ABNF_SIP_HDR_CONTENT_TYPE
		}
		switch src[pos] | 0x20 {
		case 'a':
			pos++
			if (pos + 4) >= len1 {
				return ABNF_SIP_HDR_UNKNOWN
			}
			if ((src[pos] | 0x20) == 'l') &&
				((src[pos+1] | 0x20) == 'l') &&
				(src[pos+2] == '-') &&
				((src[pos+3] | 0x20) == 'i') &&
				((src[pos+4] | 0x20) == 'd') {
				pos += 5
				if pos >= len1 {
					return ABNF_SIP_HDR_CALL_ID
				}
			}
			return ABNF_SIP_HDR_UNKNOWN
		case 'o':
			pos++
			if (pos < len1) && ((src[pos] | 0x20) == 'n') {
				pos++
				if (pos < len1) && ((src[pos] | 0x20) == 't') {
					pos++
					switch src[pos] | 0x20 {
					case 'a':
						pos++
						if (pos + 1) >= len1 {
							return ABNF_SIP_HDR_UNKNOWN
						}
						if ((src[pos] | 0x20) == 'c') &&
							((src[pos+1] | 0x20) == 't') {
							pos += 2
							if pos >= len1 {
								return ABNF_SIP_HDR_CONTACT
							}
						}
						return ABNF_SIP_HDR_UNKNOWN
					case 'e':
						pos++
						if (pos < len1) && ((src[pos] | 0x20) == 'n') {
							pos++
							if (pos < len1) && ((src[pos] | 0x20) == 't') {
								pos++
								if (pos < len1) && (src[pos] == '-') {
									pos++
									switch src[pos] | 0x20 {
									case 'd':
										pos++
										if (pos + 9) >= len1 {
											return ABNF_SIP_HDR_UNKNOWN
										}
										if ((src[pos] | 0x20) == 'i') &&
											((src[pos+1] | 0x20) == 's') &&
											((src[pos+2] | 0x20) == 'p') &&
											((src[pos+3] | 0x20) == 'o') &&
											((src[pos+4] | 0x20) == 's') &&
											((src[pos+5] | 0x20) == 'i') &&
											((src[pos+6] | 0x20) == 't') &&
											((src[pos+7] | 0x20) == 'i') &&
											((src[pos+8] | 0x20) == 'o') &&
											((src[pos+9] | 0x20) == 'n') {
											pos += 10
											if pos >= len1 {
												return ABNF_SIP_HDR_CONTENT_DISPOSITION
											}
										}
										return ABNF_SIP_HDR_UNKNOWN
									case 'e':
										pos++
										if (pos + 6) >= len1 {
											return ABNF_SIP_HDR_UNKNOWN
										}
										if ((src[pos] | 0x20) == 'n') &&
											((src[pos+1] | 0x20) == 'c') &&
											((src[pos+2] | 0x20) == 'o') &&
											((src[pos+3] | 0x20) == 'd') &&
											((src[pos+4] | 0x20) == 'i') &&
											((src[pos+5] | 0x20) == 'n') &&
											((src[pos+6] | 0x20) == 'g') {
											pos += 7
											if pos >= len1 {
												return ABNF_SIP_HDR_CONTENT_ENCODING
											}
										}
										return ABNF_SIP_HDR_UNKNOWN
									case 'l':
										pos++
										if (pos + 4) >= len1 {
											return ABNF_SIP_HDR_UNKNOWN
										}
										if ((src[pos] | 0x20) == 'e') &&
											((src[pos+1] | 0x20) == 'n') &&
											((src[pos+2] | 0x20) == 'g') &&
											((src[pos+3] | 0x20) == 't') &&
											((src[pos+4] | 0x20) == 'h') {
											pos += 5
											if pos >= len1 {
												return ABNF_SIP_HDR_CONTENT_LENGTH
											}
										}
										return ABNF_SIP_HDR_UNKNOWN
									case 't':
										pos++
										if (pos + 2) >= len1 {
											return ABNF_SIP_HDR_UNKNOWN
										}
										if ((src[pos] | 0x20) == 'y') &&
											((src[pos+1] | 0x20) == 'p') &&
											((src[pos+2] | 0x20) == 'e') {
											pos += 3
											if pos >= len1 {
												return ABNF_SIP_HDR_CONTENT_TYPE
											}
										}
										return ABNF_SIP_HDR_UNKNOWN
									}
								}
							}
						}
						return ABNF_SIP_HDR_UNKNOWN
					}
				}
			}
			return ABNF_SIP_HDR_UNKNOWN
		case 's':
			pos++
			if (pos + 1) >= len1 {
				return ABNF_SIP_HDR_UNKNOWN
			}
			if ((src[pos] | 0x20) == 'e') &&
				((src[pos+1] | 0x20) == 'q') {
				pos += 2
				if pos >= len1 {
					return ABNF_SIP_HDR_CSEQ
				}
			}
			return ABNF_SIP_HDR_UNKNOWN
		}
		return ABNF_SIP_HDR_UNKNOWN
	case 'd':
		pos++
		if (pos + 2) >= len1 {
			return ABNF_SIP_HDR_UNKNOWN
		}
		if ((src[pos] | 0x20) == 'a') &&
			((src[pos+1] | 0x20) == 't') &&
			((src[pos+2] | 0x20) == 'e') {
			pos += 3
			if pos >= len1 {
				return ABNF_SIP_HDR_DATE
			}
		}
		return ABNF_SIP_HDR_UNKNOWN
	case 'e':
		pos++
		if pos >= len1 {
			return ABNF_SIP_HDR_CONTENT_ENCODING
		}
		if (pos + 3) >= len1 {
			return ABNF_SIP_HDR_UNKNOWN
		}
		if ((src[pos] | 0x20) == 'v') &&
			((src[pos+1] | 0x20) == 'e') &&
			((src[pos+2] | 0x20) == 'n') &&
			((src[pos+3] | 0x20) == 't') {
			pos += 4
			if pos >= len1 {
				return ABNF_SIP_HDR_EVENT
			}
		}
		return ABNF_SIP_HDR_UNKNOWN
	case 'f':
		pos++
		if pos >= len1 {
			return ABNF_SIP_HDR_FROM
		}
		if (pos + 2) >= len1 {
			return ABNF_SIP_HDR_UNKNOWN
		}
		if ((src[pos] | 0x20) == 'r') &&
			((src[pos+1] | 0x20) == 'o') &&
			((src[pos+2] | 0x20) == 'm') {
			pos += 3
			if pos >= len1 {
				return ABNF_SIP_HDR_FROM
			}
		}
		return ABNF_SIP_HDR_UNKNOWN
	case 'i':
		pos++
		if pos >= len1 {
			return ABNF_SIP_HDR_CALL_ID
		}
		return ABNF_SIP_HDR_UNKNOWN
	case 'j':
		pos++
		if pos >= len1 {
			return ABNF_SIP_HDR_REJECT_CONTACT
		}
		return ABNF_SIP_HDR_UNKNOWN
	case 'k':
		pos++
		if pos >= len1 {
			return ABNF_SIP_HDR_SUPPORTED
		}
		return ABNF_SIP_HDR_UNKNOWN
	case 'l':
		pos++
		if pos >= len1 {
			return ABNF_SIP_HDR_CONTENT_LENGTH
		}
		return ABNF_SIP_HDR_UNKNOWN
	case 'm':
		pos++
		if pos >= len1 {
			return ABNF_SIP_HDR_CONTACT
		}
		switch src[pos] | 0x20 {
		case 'a':
			pos++
			if (pos + 9) >= len1 {
				return ABNF_SIP_HDR_UNKNOWN
			}
			if ((src[pos] | 0x20) == 'x') &&
				(src[pos+1] == '-') &&
				((src[pos+2] | 0x20) == 'f') &&
				((src[pos+3] | 0x20) == 'o') &&
				((src[pos+4] | 0x20) == 'r') &&
				((src[pos+5] | 0x20) == 'w') &&
				((src[pos+6] | 0x20) == 'a') &&
				((src[pos+7] | 0x20) == 'r') &&
				((src[pos+8] | 0x20) == 'd') &&
				((src[pos+9] | 0x20) == 's') {
				pos += 10
				if pos >= len1 {
					return ABNF_SIP_HDR_MAX_FORWARDS
				}
			}
			return ABNF_SIP_HDR_UNKNOWN
		case 'i':
			pos++
			if (pos + 9) >= len1 {
				return ABNF_SIP_HDR_UNKNOWN
			}
			if ((src[pos] | 0x20) == 'm') &&
				((src[pos+1] | 0x20) == 'e') &&
				(src[pos+2] == '-') &&
				((src[pos+3] | 0x20) == 'v') &&
				((src[pos+4] | 0x20) == 'e') &&
				((src[pos+5] | 0x20) == 'r') &&
				((src[pos+6] | 0x20) == 's') &&
				((src[pos+7] | 0x20) == 'i') &&
				((src[pos+8] | 0x20) == 'o') &&
				((src[pos+9] | 0x20) == 'n') {
				pos += 10
				if pos >= len1 {
					return ABNF_SIP_HDR_MIME_VERSION
				}
			}
			return ABNF_SIP_HDR_UNKNOWN
		}
		return ABNF_SIP_HDR_UNKNOWN
	case 'o':
		pos++
		if pos >= len1 {
			return ABNF_SIP_HDR_EVENT
		}
		return ABNF_SIP_HDR_UNKNOWN
	case 'r':
		pos++
		if pos >= len1 {
			return ABNF_SIP_HDR_REFER_TO
		}
		switch src[pos] | 0x20 {
		case 'e':
			pos++
			switch src[pos] | 0x20 {
			case 'c':
				pos++
				if (pos + 8) >= len1 {
					return ABNF_SIP_HDR_UNKNOWN
				}
				if ((src[pos] | 0x20) == 'o') &&
					((src[pos+1] | 0x20) == 'r') &&
					((src[pos+2] | 0x20) == 'd') &&
					(src[pos+3] == '-') &&
					((src[pos+4] | 0x20) == 'r') &&
					((src[pos+5] | 0x20) == 'o') &&
					((src[pos+6] | 0x20) == 'u') &&
					((src[pos+7] | 0x20) == 't') &&
					((src[pos+8] | 0x20) == 'e') {
					pos += 9
					if pos >= len1 {
						return ABNF_SIP_HDR_RECORD_ROUTE
					}
				}
				return ABNF_SIP_HDR_UNKNOWN
			case 'f':
				pos++
				if (pos < len1) && ((src[pos] | 0x20) == 'e') {
					pos++
					if (pos < len1) && ((src[pos] | 0x20) == 'r') {
						pos++
						switch src[pos] | 0x20 {
						case '-':
							pos++
							if (pos + 1) >= len1 {
								return ABNF_SIP_HDR_UNKNOWN
							}
							if ((src[pos] | 0x20) == 't') &&
								((src[pos+1] | 0x20) == 'o') {
								pos += 2
								if pos >= len1 {
									return ABNF_SIP_HDR_REFER_TO
								}
							}
							return ABNF_SIP_HDR_UNKNOWN
						case 'r':
							pos++
							if (pos + 4) >= len1 {
								return ABNF_SIP_HDR_UNKNOWN
							}
							if ((src[pos] | 0x20) == 'e') &&
								((src[pos+1] | 0x20) == 'd') &&
								(src[pos+2] == '-') &&
								((src[pos+3] | 0x20) == 'b') &&
								((src[pos+4] | 0x20) == 'y') {
								pos += 5
								if pos >= len1 {
									return ABNF_SIP_HDR_REFERRED_BY
								}
							}
							return ABNF_SIP_HDR_UNKNOWN
						}
					}
				}
				return ABNF_SIP_HDR_UNKNOWN
			case 'j':
				pos++
				if (pos + 10) >= len1 {
					return ABNF_SIP_HDR_UNKNOWN
				}
				if ((src[pos] | 0x20) == 'e') &&
					((src[pos+1] | 0x20) == 'c') &&
					((src[pos+2] | 0x20) == 't') &&
					(src[pos+3] == '-') &&
					((src[pos+4] | 0x20) == 'c') &&
					((src[pos+5] | 0x20) == 'o') &&
					((src[pos+6] | 0x20) == 'n') &&
					((src[pos+7] | 0x20) == 't') &&
					((src[pos+8] | 0x20) == 'a') &&
					((src[pos+9] | 0x20) == 'c') &&
					((src[pos+10] | 0x20) == 't') {
					pos += 11
					if pos >= len1 {
						return ABNF_SIP_HDR_REJECT_CONTACT
					}
				}
				return ABNF_SIP_HDR_UNKNOWN
			case 'q':
				pos++
				if (pos + 15) >= len1 {
					return ABNF_SIP_HDR_UNKNOWN
				}
				if ((src[pos] | 0x20) == 'u') &&
					((src[pos+1] | 0x20) == 'e') &&
					((src[pos+2] | 0x20) == 's') &&
					((src[pos+3] | 0x20) == 't') &&
					(src[pos+4] == '-') &&
					((src[pos+5] | 0x20) == 'd') &&
					((src[pos+6] | 0x20) == 'i') &&
					((src[pos+7] | 0x20) == 's') &&
					((src[pos+8] | 0x20) == 'p') &&
					((src[pos+9] | 0x20) == 'o') &&
					((src[pos+10] | 0x20) == 's') &&
					((src[pos+11] | 0x20) == 'i') &&
					((src[pos+12] | 0x20) == 't') &&
					((src[pos+13] | 0x20) == 'i') &&
					((src[pos+14] | 0x20) == 'o') &&
					((src[pos+15] | 0x20) == 'n') {
					pos += 16
					if pos >= len1 {
						return ABNF_SIP_HDR_REQUEST_DISPOSITION
					}
				}
				return ABNF_SIP_HDR_UNKNOWN
			}
			return ABNF_SIP_HDR_UNKNOWN
		case 'o':
			pos++
			if (pos + 2) >= len1 {
				return ABNF_SIP_HDR_UNKNOWN
			}
			if ((src[pos] | 0x20) == 'u') &&
				((src[pos+1] | 0x20) == 't') &&
				((src[pos+2] | 0x20) == 'e') {
				pos += 3
				if pos >= len1 {
					return ABNF_SIP_HDR_ROUTE
				}
			}
			return ABNF_SIP_HDR_UNKNOWN
		}
		return ABNF_SIP_HDR_UNKNOWN
	case 's':
		pos++
		if pos >= len1 {
			return ABNF_SIP_HDR_SUBJECT
		}
		switch src[pos] | 0x20 {
		case 'e':
			pos++
			if (pos + 12) >= len1 {
				return ABNF_SIP_HDR_UNKNOWN
			}
			if ((src[pos] | 0x20) == 's') &&
				((src[pos+1] | 0x20) == 's') &&
				((src[pos+2] | 0x20) == 'i') &&
				((src[pos+3] | 0x20) == 'o') &&
				((src[pos+4] | 0x20) == 'n') &&
				(src[pos+5] == '-') &&
				((src[pos+6] | 0x20) == 'e') &&
				((src[pos+7] | 0x20) == 'x') &&
				((src[pos+8] | 0x20) == 'p') &&
				((src[pos+9] | 0x20) == 'i') &&
				((src[pos+10] | 0x20) == 'r') &&
				((src[pos+11] | 0x20) == 'e') &&
				((src[pos+12] | 0x20) == 's') {
				pos += 13
				if pos >= len1 {
					return ABNF_SIP_HDR_SESSION_EXPIRES
				}
			}
			return ABNF_SIP_HDR_UNKNOWN
		case 'u':
			pos++
			switch src[pos] | 0x20 {
			case 'b':
				pos++
				if (pos + 3) >= len1 {
					return ABNF_SIP_HDR_UNKNOWN
				}
				if ((src[pos] | 0x20) == 'j') &&
					((src[pos+1] | 0x20) == 'e') &&
					((src[pos+2] | 0x20) == 'c') &&
					((src[pos+3] | 0x20) == 't') {
					pos += 4
					if pos >= len1 {
						return ABNF_SIP_HDR_SUBJECT
					}
				}
				return ABNF_SIP_HDR_UNKNOWN
			case 'p':
				pos++
				if (pos + 5) >= len1 {
					return ABNF_SIP_HDR_UNKNOWN
				}
				if ((src[pos] | 0x20) == 'p') &&
					((src[pos+1] | 0x20) == 'o') &&
					((src[pos+2] | 0x20) == 'r') &&
					((src[pos+3] | 0x20) == 't') &&
					((src[pos+4] | 0x20) == 'e') &&
					((src[pos+5] | 0x20) == 'd') {
					pos += 6
					if pos >= len1 {
						return ABNF_SIP_HDR_SUPPORTED
					}
				}
				return ABNF_SIP_HDR_UNKNOWN
			}
			return ABNF_SIP_HDR_UNKNOWN
		}
		return ABNF_SIP_HDR_UNKNOWN
	case 't':
		pos++
		if pos >= len1 {
			return ABNF_SIP_HDR_TO
		}
		if (pos < len1) && ((src[pos] | 0x20) == 'o') {
			pos++
			if pos >= len1 {
				return ABNF_SIP_HDR_TO
			}
		}
		return ABNF_SIP_HDR_UNKNOWN
	case 'u':
		pos++
		if pos >= len1 {
			return ABNF_SIP_HDR_ALLOW_EVENTS
		}
		return ABNF_SIP_HDR_UNKNOWN
	case 'v':
		pos++
		if pos >= len1 {
			return ABNF_SIP_HDR_VIA
		}
		if (pos + 1) >= len1 {
			return ABNF_SIP_HDR_UNKNOWN
		}
		if ((src[pos] | 0x20) == 'i') &&
			((src[pos+1] | 0x20) == 'a') {
			pos += 2
			if pos >= len1 {
				return ABNF_SIP_HDR_VIA
			}
		}
		return ABNF_SIP_HDR_UNKNOWN
	case 'x':
		pos++
		if pos >= len1 {
			return ABNF_SIP_HDR_SESSION_EXPIRES
		}
		return ABNF_SIP_HDR_UNKNOWN
	}

	return ABNF_SIP_HDR_UNKNOWN
}
