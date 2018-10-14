/*---------------- generated by ps_abnf v0.1.0 2018-04-23 23:48:09.186 ----------------*/
package sipparser

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

func GetSipHeaderIndex(src []byte, pos AbnfPos) (SipHeaderIndexType, AbnfPos) {
	len1 := AbnfPos(len(src))

	if pos >= len1 {
		return ABNF_SIP_HDR_UNKNOWN, pos
	}

	switch src[pos] {
	case 'A':
		pos++
		switch src[pos] {
		case 'c':
			pos++
			if (pos + 11) >= len1 {
				return ABNF_SIP_HDR_UNKNOWN, pos
			}
			if (src[pos] == 'c') &&
				(src[pos+1] == 'e') &&
				(src[pos+2] == 'p') &&
				(src[pos+3] == 't') &&
				(src[pos+4] == '-') &&
				(src[pos+5] == 'C') &&
				(src[pos+6] == 'o') &&
				(src[pos+7] == 'n') &&
				(src[pos+8] == 't') &&
				(src[pos+9] == 'a') &&
				(src[pos+10] == 'c') &&
				(src[pos+11] == 't') {
				pos += 12
				if (pos >= len1) || !IsSipToken(src[pos]) {
					return ABNF_SIP_HDR_ACCEPT_CONTACT, pos
				}
			}
			return ABNF_SIP_HDR_UNKNOWN, pos
		case 'l':
			pos++
			if (pos < len1) && (src[pos] == 'l') {
				pos++
				if (pos < len1) && (src[pos] == 'o') {
					pos++
					if (pos < len1) && (src[pos] == 'w') {
						pos++
						if (pos >= len1) || !IsSipToken(src[pos]) {
							return ABNF_SIP_HDR_ALLOW, pos
						}
						if (pos + 6) >= len1 {
							return ABNF_SIP_HDR_UNKNOWN, pos
						}
						if (src[pos] == '-') &&
							(src[pos+1] == 'E') &&
							(src[pos+2] == 'v') &&
							(src[pos+3] == 'e') &&
							(src[pos+4] == 'n') &&
							(src[pos+5] == 't') &&
							(src[pos+6] == 's') {
							pos += 7
							if (pos >= len1) || !IsSipToken(src[pos]) {
								return ABNF_SIP_HDR_ALLOW_EVENTS, pos
							}
						}
					}
				}
			}
			return ABNF_SIP_HDR_UNKNOWN, pos
		}
		return ABNF_SIP_HDR_UNKNOWN, pos
	case 'C':
		pos++
		switch src[pos] {
		case 'S':
			pos++
			if (pos + 1) >= len1 {
				return ABNF_SIP_HDR_UNKNOWN, pos
			}
			if (src[pos] == 'e') &&
				(src[pos+1] == 'q') {
				pos += 2
				if (pos >= len1) || !IsSipToken(src[pos]) {
					return ABNF_SIP_HDR_CSEQ, pos
				}
			}
			return ABNF_SIP_HDR_UNKNOWN, pos
		case 'a':
			pos++
			if (pos + 4) >= len1 {
				return ABNF_SIP_HDR_UNKNOWN, pos
			}
			if (src[pos] == 'l') &&
				(src[pos+1] == 'l') &&
				(src[pos+2] == '-') &&
				(src[pos+3] == 'I') &&
				(src[pos+4] == 'D') {
				pos += 5
				if (pos >= len1) || !IsSipToken(src[pos]) {
					return ABNF_SIP_HDR_CALL_ID, pos
				}
			}
			return ABNF_SIP_HDR_UNKNOWN, pos
		case 'o':
			pos++
			if (pos < len1) && (src[pos] == 'n') {
				pos++
				if (pos < len1) && (src[pos] == 't') {
					pos++
					switch src[pos] {
					case 'a':
						pos++
						if (pos + 1) >= len1 {
							return ABNF_SIP_HDR_UNKNOWN, pos
						}
						if (src[pos] == 'c') &&
							(src[pos+1] == 't') {
							pos += 2
							if (pos >= len1) || !IsSipToken(src[pos]) {
								return ABNF_SIP_HDR_CONTACT, pos
							}
						}
						return ABNF_SIP_HDR_UNKNOWN, pos
					case 'e':
						pos++
						if (pos < len1) && (src[pos] == 'n') {
							pos++
							if (pos < len1) && (src[pos] == 't') {
								pos++
								if (pos < len1) && (src[pos] == '-') {
									pos++
									switch src[pos] {
									case 'D':
										pos++
										if (pos + 9) >= len1 {
											return ABNF_SIP_HDR_UNKNOWN, pos
										}
										if (src[pos] == 'i') &&
											(src[pos+1] == 's') &&
											(src[pos+2] == 'p') &&
											(src[pos+3] == 'o') &&
											(src[pos+4] == 's') &&
											(src[pos+5] == 'i') &&
											(src[pos+6] == 't') &&
											(src[pos+7] == 'i') &&
											(src[pos+8] == 'o') &&
											(src[pos+9] == 'n') {
											pos += 10
											if (pos >= len1) || !IsSipToken(src[pos]) {
												return ABNF_SIP_HDR_CONTENT_DISPOSITION, pos
											}
										}
										return ABNF_SIP_HDR_UNKNOWN, pos
									case 'E':
										pos++
										if (pos + 6) >= len1 {
											return ABNF_SIP_HDR_UNKNOWN, pos
										}
										if (src[pos] == 'n') &&
											(src[pos+1] == 'c') &&
											(src[pos+2] == 'o') &&
											(src[pos+3] == 'd') &&
											(src[pos+4] == 'i') &&
											(src[pos+5] == 'n') &&
											(src[pos+6] == 'g') {
											pos += 7
											if (pos >= len1) || !IsSipToken(src[pos]) {
												return ABNF_SIP_HDR_CONTENT_ENCODING, pos
											}
										}
										return ABNF_SIP_HDR_UNKNOWN, pos
									case 'L':
										pos++
										if (pos + 4) >= len1 {
											return ABNF_SIP_HDR_UNKNOWN, pos
										}
										if (src[pos] == 'e') &&
											(src[pos+1] == 'n') &&
											(src[pos+2] == 'g') &&
											(src[pos+3] == 't') &&
											(src[pos+4] == 'h') {
											pos += 5
											if (pos >= len1) || !IsSipToken(src[pos]) {
												return ABNF_SIP_HDR_CONTENT_LENGTH, pos
											}
										}
										return ABNF_SIP_HDR_UNKNOWN, pos
									case 'T':
										pos++
										if (pos + 2) >= len1 {
											return ABNF_SIP_HDR_UNKNOWN, pos
										}
										if (src[pos] == 'y') &&
											(src[pos+1] == 'p') &&
											(src[pos+2] == 'e') {
											pos += 3
											if (pos >= len1) || !IsSipToken(src[pos]) {
												return ABNF_SIP_HDR_CONTENT_TYPE, pos
											}
										}
										return ABNF_SIP_HDR_UNKNOWN, pos
									}
								}
							}
						}
						return ABNF_SIP_HDR_UNKNOWN, pos
					}
				}
			}
			return ABNF_SIP_HDR_UNKNOWN, pos
		}
		return ABNF_SIP_HDR_UNKNOWN, pos
	case 'D':
		pos++
		if (pos + 2) >= len1 {
			return ABNF_SIP_HDR_UNKNOWN, pos
		}
		if (src[pos] == 'a') &&
			(src[pos+1] == 't') &&
			(src[pos+2] == 'e') {
			pos += 3
			if (pos >= len1) || !IsSipToken(src[pos]) {
				return ABNF_SIP_HDR_DATE, pos
			}
		}
		return ABNF_SIP_HDR_UNKNOWN, pos
	case 'E':
		pos++
		if (pos + 3) >= len1 {
			return ABNF_SIP_HDR_UNKNOWN, pos
		}
		if (src[pos] == 'v') &&
			(src[pos+1] == 'e') &&
			(src[pos+2] == 'n') &&
			(src[pos+3] == 't') {
			pos += 4
			if (pos >= len1) || !IsSipToken(src[pos]) {
				return ABNF_SIP_HDR_EVENT, pos
			}
		}
		return ABNF_SIP_HDR_UNKNOWN, pos
	case 'F':
		pos++
		if (pos + 2) >= len1 {
			return ABNF_SIP_HDR_UNKNOWN, pos
		}
		if (src[pos] == 'r') &&
			(src[pos+1] == 'o') &&
			(src[pos+2] == 'm') {
			pos += 3
			if (pos >= len1) || !IsSipToken(src[pos]) {
				return ABNF_SIP_HDR_FROM, pos
			}
		}
		return ABNF_SIP_HDR_UNKNOWN, pos
	case 'M':
		pos++
		switch src[pos] {
		case 'I':
			pos++
			if (pos + 9) >= len1 {
				return ABNF_SIP_HDR_UNKNOWN, pos
			}
			if (src[pos] == 'M') &&
				(src[pos+1] == 'E') &&
				(src[pos+2] == '-') &&
				(src[pos+3] == 'V') &&
				(src[pos+4] == 'e') &&
				(src[pos+5] == 'r') &&
				(src[pos+6] == 's') &&
				(src[pos+7] == 'i') &&
				(src[pos+8] == 'o') &&
				(src[pos+9] == 'n') {
				pos += 10
				if (pos >= len1) || !IsSipToken(src[pos]) {
					return ABNF_SIP_HDR_MIME_VERSION, pos
				}
			}
			return ABNF_SIP_HDR_UNKNOWN, pos
		case 'a':
			pos++
			if (pos + 9) >= len1 {
				return ABNF_SIP_HDR_UNKNOWN, pos
			}
			if (src[pos] == 'x') &&
				(src[pos+1] == '-') &&
				(src[pos+2] == 'F') &&
				(src[pos+3] == 'o') &&
				(src[pos+4] == 'r') &&
				(src[pos+5] == 'w') &&
				(src[pos+6] == 'a') &&
				(src[pos+7] == 'r') &&
				(src[pos+8] == 'd') &&
				(src[pos+9] == 's') {
				pos += 10
				if (pos >= len1) || !IsSipToken(src[pos]) {
					return ABNF_SIP_HDR_MAX_FORWARDS, pos
				}
			}
			return ABNF_SIP_HDR_UNKNOWN, pos
		}
		return ABNF_SIP_HDR_UNKNOWN, pos
	case 'R':
		pos++
		switch src[pos] {
		case 'e':
			pos++
			switch src[pos] {
			case 'c':
				pos++
				if (pos + 8) >= len1 {
					return ABNF_SIP_HDR_UNKNOWN, pos
				}
				if (src[pos] == 'o') &&
					(src[pos+1] == 'r') &&
					(src[pos+2] == 'd') &&
					(src[pos+3] == '-') &&
					(src[pos+4] == 'R') &&
					(src[pos+5] == 'o') &&
					(src[pos+6] == 'u') &&
					(src[pos+7] == 't') &&
					(src[pos+8] == 'e') {
					pos += 9
					if (pos >= len1) || !IsSipToken(src[pos]) {
						return ABNF_SIP_HDR_RECORD_ROUTE, pos
					}
				}
				return ABNF_SIP_HDR_UNKNOWN, pos
			case 'f':
				pos++
				if (pos < len1) && (src[pos] == 'e') {
					pos++
					if (pos < len1) && (src[pos] == 'r') {
						pos++
						switch src[pos] {
						case '-':
							pos++
							if (pos + 1) >= len1 {
								return ABNF_SIP_HDR_UNKNOWN, pos
							}
							if (src[pos] == 'T') &&
								(src[pos+1] == 'o') {
								pos += 2
								if (pos >= len1) || !IsSipToken(src[pos]) {
									return ABNF_SIP_HDR_REFER_TO, pos
								}
							}
							return ABNF_SIP_HDR_UNKNOWN, pos
						case 'r':
							pos++
							if (pos + 4) >= len1 {
								return ABNF_SIP_HDR_UNKNOWN, pos
							}
							if (src[pos] == 'e') &&
								(src[pos+1] == 'd') &&
								(src[pos+2] == '-') &&
								(src[pos+3] == 'B') &&
								(src[pos+4] == 'y') {
								pos += 5
								if (pos >= len1) || !IsSipToken(src[pos]) {
									return ABNF_SIP_HDR_REFERRED_BY, pos
								}
							}
							return ABNF_SIP_HDR_UNKNOWN, pos
						}
					}
				}
				return ABNF_SIP_HDR_UNKNOWN, pos
			case 'j':
				pos++
				if (pos + 10) >= len1 {
					return ABNF_SIP_HDR_UNKNOWN, pos
				}
				if (src[pos] == 'e') &&
					(src[pos+1] == 'c') &&
					(src[pos+2] == 't') &&
					(src[pos+3] == '-') &&
					(src[pos+4] == 'C') &&
					(src[pos+5] == 'o') &&
					(src[pos+6] == 'n') &&
					(src[pos+7] == 't') &&
					(src[pos+8] == 'a') &&
					(src[pos+9] == 'c') &&
					(src[pos+10] == 't') {
					pos += 11
					if (pos >= len1) || !IsSipToken(src[pos]) {
						return ABNF_SIP_HDR_REJECT_CONTACT, pos
					}
				}
				return ABNF_SIP_HDR_UNKNOWN, pos
			case 'q':
				pos++
				if (pos + 15) >= len1 {
					return ABNF_SIP_HDR_UNKNOWN, pos
				}
				if (src[pos] == 'u') &&
					(src[pos+1] == 'e') &&
					(src[pos+2] == 's') &&
					(src[pos+3] == 't') &&
					(src[pos+4] == '-') &&
					(src[pos+5] == 'D') &&
					(src[pos+6] == 'i') &&
					(src[pos+7] == 's') &&
					(src[pos+8] == 'p') &&
					(src[pos+9] == 'o') &&
					(src[pos+10] == 's') &&
					(src[pos+11] == 'i') &&
					(src[pos+12] == 't') &&
					(src[pos+13] == 'i') &&
					(src[pos+14] == 'o') &&
					(src[pos+15] == 'n') {
					pos += 16
					if (pos >= len1) || !IsSipToken(src[pos]) {
						return ABNF_SIP_HDR_REQUEST_DISPOSITION, pos
					}
				}
				return ABNF_SIP_HDR_UNKNOWN, pos
			}
			return ABNF_SIP_HDR_UNKNOWN, pos
		case 'o':
			pos++
			if (pos + 2) >= len1 {
				return ABNF_SIP_HDR_UNKNOWN, pos
			}
			if (src[pos] == 'u') &&
				(src[pos+1] == 't') &&
				(src[pos+2] == 'e') {
				pos += 3
				if (pos >= len1) || !IsSipToken(src[pos]) {
					return ABNF_SIP_HDR_ROUTE, pos
				}
			}
			return ABNF_SIP_HDR_UNKNOWN, pos
		}
		return ABNF_SIP_HDR_UNKNOWN, pos
	case 'S':
		pos++
		switch src[pos] {
		case 'e':
			pos++
			if (pos + 12) >= len1 {
				return ABNF_SIP_HDR_UNKNOWN, pos
			}
			if (src[pos] == 's') &&
				(src[pos+1] == 's') &&
				(src[pos+2] == 'i') &&
				(src[pos+3] == 'o') &&
				(src[pos+4] == 'n') &&
				(src[pos+5] == '-') &&
				(src[pos+6] == 'E') &&
				(src[pos+7] == 'x') &&
				(src[pos+8] == 'p') &&
				(src[pos+9] == 'i') &&
				(src[pos+10] == 'r') &&
				(src[pos+11] == 'e') &&
				(src[pos+12] == 's') {
				pos += 13
				if (pos >= len1) || !IsSipToken(src[pos]) {
					return ABNF_SIP_HDR_SESSION_EXPIRES, pos
				}
			}
			return ABNF_SIP_HDR_UNKNOWN, pos
		case 'u':
			pos++
			switch src[pos] {
			case 'b':
				pos++
				if (pos + 3) >= len1 {
					return ABNF_SIP_HDR_UNKNOWN, pos
				}
				if (src[pos] == 'j') &&
					(src[pos+1] == 'e') &&
					(src[pos+2] == 'c') &&
					(src[pos+3] == 't') {
					pos += 4
					if (pos >= len1) || !IsSipToken(src[pos]) {
						return ABNF_SIP_HDR_SUBJECT, pos
					}
				}
				return ABNF_SIP_HDR_UNKNOWN, pos
			case 'p':
				pos++
				if (pos + 5) >= len1 {
					return ABNF_SIP_HDR_UNKNOWN, pos
				}
				if (src[pos] == 'p') &&
					(src[pos+1] == 'o') &&
					(src[pos+2] == 'r') &&
					(src[pos+3] == 't') &&
					(src[pos+4] == 'e') &&
					(src[pos+5] == 'd') {
					pos += 6
					if (pos >= len1) || !IsSipToken(src[pos]) {
						return ABNF_SIP_HDR_SUPPORTED, pos
					}
				}
				return ABNF_SIP_HDR_UNKNOWN, pos
			}
			return ABNF_SIP_HDR_UNKNOWN, pos
		}
		return ABNF_SIP_HDR_UNKNOWN, pos
	case 'T':
		pos++
		if (pos < len1) && (src[pos] == 'o') {
			pos++
			if (pos >= len1) || !IsSipToken(src[pos]) {
				return ABNF_SIP_HDR_TO, pos
			}
		}
		return ABNF_SIP_HDR_UNKNOWN, pos
	case 'V':
		pos++
		if (pos + 1) >= len1 {
			return ABNF_SIP_HDR_UNKNOWN, pos
		}
		if (src[pos] == 'i') &&
			(src[pos+1] == 'a') {
			pos += 2
			if (pos >= len1) || !IsSipToken(src[pos]) {
				return ABNF_SIP_HDR_VIA, pos
			}
		}
		return ABNF_SIP_HDR_UNKNOWN, pos
	case 'a':
		pos++
		if (pos >= len1) || !IsSipToken(src[pos]) {
			return ABNF_SIP_HDR_ACCEPT_CONTACT, pos
		}
		return ABNF_SIP_HDR_UNKNOWN, pos
	case 'b':
		pos++
		if (pos >= len1) || !IsSipToken(src[pos]) {
			return ABNF_SIP_HDR_REFERRED_BY, pos
		}
		return ABNF_SIP_HDR_UNKNOWN, pos
	case 'c':
		pos++
		if (pos >= len1) || !IsSipToken(src[pos]) {
			return ABNF_SIP_HDR_CONTENT_TYPE, pos
		}
		return ABNF_SIP_HDR_UNKNOWN, pos
	case 'e':
		pos++
		if (pos >= len1) || !IsSipToken(src[pos]) {
			return ABNF_SIP_HDR_CONTENT_ENCODING, pos
		}
		return ABNF_SIP_HDR_UNKNOWN, pos
	case 'f':
		pos++
		if (pos >= len1) || !IsSipToken(src[pos]) {
			return ABNF_SIP_HDR_FROM, pos
		}
		return ABNF_SIP_HDR_UNKNOWN, pos
	case 'i':
		pos++
		if (pos >= len1) || !IsSipToken(src[pos]) {
			return ABNF_SIP_HDR_CALL_ID, pos
		}
		return ABNF_SIP_HDR_UNKNOWN, pos
	case 'j':
		pos++
		if (pos >= len1) || !IsSipToken(src[pos]) {
			return ABNF_SIP_HDR_REJECT_CONTACT, pos
		}
		return ABNF_SIP_HDR_UNKNOWN, pos
	case 'k':
		pos++
		if (pos >= len1) || !IsSipToken(src[pos]) {
			return ABNF_SIP_HDR_SUPPORTED, pos
		}
		return ABNF_SIP_HDR_UNKNOWN, pos
	case 'l':
		pos++
		if (pos >= len1) || !IsSipToken(src[pos]) {
			return ABNF_SIP_HDR_CONTENT_LENGTH, pos
		}
		return ABNF_SIP_HDR_UNKNOWN, pos
	case 'm':
		pos++
		if (pos >= len1) || !IsSipToken(src[pos]) {
			return ABNF_SIP_HDR_CONTACT, pos
		}
		return ABNF_SIP_HDR_UNKNOWN, pos
	case 'o':
		pos++
		if (pos >= len1) || !IsSipToken(src[pos]) {
			return ABNF_SIP_HDR_EVENT, pos
		}
		return ABNF_SIP_HDR_UNKNOWN, pos
	case 'r':
		pos++
		if (pos >= len1) || !IsSipToken(src[pos]) {
			return ABNF_SIP_HDR_REFER_TO, pos
		}
		return ABNF_SIP_HDR_UNKNOWN, pos
	case 's':
		pos++
		if (pos >= len1) || !IsSipToken(src[pos]) {
			return ABNF_SIP_HDR_SUBJECT, pos
		}
		return ABNF_SIP_HDR_UNKNOWN, pos
	case 't':
		pos++
		if (pos >= len1) || !IsSipToken(src[pos]) {
			return ABNF_SIP_HDR_TO, pos
		}
		return ABNF_SIP_HDR_UNKNOWN, pos
	case 'u':
		pos++
		if (pos >= len1) || !IsSipToken(src[pos]) {
			return ABNF_SIP_HDR_ALLOW_EVENTS, pos
		}
		return ABNF_SIP_HDR_UNKNOWN, pos
	case 'v':
		pos++
		if (pos >= len1) || !IsSipToken(src[pos]) {
			return ABNF_SIP_HDR_VIA, pos
		}
		return ABNF_SIP_HDR_UNKNOWN, pos
	case 'x':
		pos++
		if (pos >= len1) || !IsSipToken(src[pos]) {
			return ABNF_SIP_HDR_SESSION_EXPIRES, pos
		}
		return ABNF_SIP_HDR_UNKNOWN, pos
	}

	return ABNF_SIP_HDR_UNKNOWN, pos
}