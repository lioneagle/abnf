package sip_charset

/*---------------- mask definition ----------------*/
const (
	PS_SIP_CHARSETS_MASK_DIGIT              PS_BYTE = 0x01
	PS_SIP_CHARSETS_MASK_ALPHA              PS_BYTE = 0x02
	PS_SIP_CHARSETS_MASK_LOWER              PS_BYTE = 0x04
	PS_SIP_CHARSETS_MASK_UPPER              PS_BYTE = 0x08
	PS_SIP_CHARSETS_MASK_ALPHANUM           PS_BYTE = 0x10
	PS_SIP_CHARSETS_MASK_HEX                PS_BYTE = 0x20
	PS_SIP_CHARSETS_MASK_LOWER_HEX_ALPHA    PS_BYTE = 0x40
	PS_SIP_CHARSETS_MASK_UPPER_HEX_ALPHA    PS_BYTE = 0x80
	PS_SIP_CHARSETS_MASK_WSP                PS_BYTE = 0x01
)

/*---------------- action definition ----------------*/
func PS_SIP_isDigit(ch byte)          { return (g_sipCharsets0[ch] & PS_SIP_CHARSETS_MASK_DIGIT) != 0 }
func PS_SIP_isAlpha(ch byte)          { return (g_sipCharsets0[ch] & PS_SIP_CHARSETS_MASK_ALPHA) != 0 }
func PS_SIP_isLower(ch byte)          { return (g_sipCharsets0[ch] & PS_SIP_CHARSETS_MASK_LOWER) != 0 }
func PS_SIP_isUpper(ch byte)          { return (g_sipCharsets0[ch] & PS_SIP_CHARSETS_MASK_UPPER) != 0 }
func PS_SIP_isAlphanum(ch byte)       { return (g_sipCharsets0[ch] & PS_SIP_CHARSETS_MASK_ALPHANUM) != 0 }
func PS_SIP_isHex(ch byte)            { return (g_sipCharsets0[ch] & PS_SIP_CHARSETS_MASK_HEX) != 0 }
func PS_SIP_isLowerHexAlpha(ch byte)  { return (g_sipCharsets0[ch] & PS_SIP_CHARSETS_MASK_LOWER_HEX_ALPHA) != 0 }
func PS_SIP_isUpperHexAlpha(ch byte)  { return (g_sipCharsets0[ch] & PS_SIP_CHARSETS_MASK_UPPER_HEX_ALPHA) != 0 }
func PS_SIP_isWsp(ch byte)            { return (g_sipCharsets1[ch] & PS_SIP_CHARSETS_MASK_WSP) != 0 }

/*---------------- var definition ----------------*/
var g_sipCharsets0 = [256]PS_BYTE{
	0x00,  /* position 000 */
	0x00,  /* position 001 */
	0x00,  /* position 002 */
	0x00,  /* position 003 */
	0x00,  /* position 004 */
	0x00,  /* position 005 */
	0x00,  /* position 006 */
	0x00,  /* position 007 */
	0x00,  /* position 008 */
	0x00,  /* position 009 */
	0x00,  /* position 010 */
	0x00,  /* position 011 */
	0x00,  /* position 012 */
	0x00,  /* position 013 */
	0x00,  /* position 014 */
	0x00,  /* position 015 */
	0x00,  /* position 016 */
	0x00,  /* position 017 */
	0x00,  /* position 018 */
	0x00,  /* position 019 */
	0x00,  /* position 020 */
	0x00,  /* position 021 */
	0x00,  /* position 022 */
	0x00,  /* position 023 */
	0x00,  /* position 024 */
	0x00,  /* position 025 */
	0x00,  /* position 026 */
	0x00,  /* position 027 */
	0x00,  /* position 028 */
	0x00,  /* position 029 */
	0x00,  /* position 030 */
	0x00,  /* position 031 */
	0x00,  /* position 032  ' ' */
	0x00,  /* position 033  '!' */
	0x00,  /* position 034  '"' */
	0x00,  /* position 035  '#' */
	0x00,  /* position 036  '$' */
	0x00,  /* position 037  '%' */
	0x00,  /* position 038  '&' */
	0x00,  /* position 039  ''' */
	0x00,  /* position 040  '(' */
	0x00,  /* position 041  ')' */
	0x00,  /* position 042  '*' */
	0x00,  /* position 043  '+' */
	0x00,  /* position 044  ',' */
	0x00,  /* position 045  '-' */
	0x00,  /* position 046  '.' */
	0x00,  /* position 047  '/' */
	0x31,  /* position 048  '0' */
	0x31,  /* position 049  '1' */
	0x31,  /* position 050  '2' */
	0x31,  /* position 051  '3' */
	0x31,  /* position 052  '4' */
	0x31,  /* position 053  '5' */
	0x31,  /* position 054  '6' */
	0x31,  /* position 055  '7' */
	0x31,  /* position 056  '8' */
	0x31,  /* position 057  '9' */
	0x00,  /* position 058  ':' */
	0x00,  /* position 059  ';' */
	0x00,  /* position 060  '<' */
	0x00,  /* position 061  '=' */
	0x00,  /* position 062  '>' */
	0x00,  /* position 063  '?' */
	0x00,  /* position 064  '@' */
	0xba,  /* position 065  'A' */
	0xba,  /* position 066  'B' */
	0xba,  /* position 067  'C' */
	0xba,  /* position 068  'D' */
	0xba,  /* position 069  'E' */
	0xba,  /* position 070  'F' */
	0x1a,  /* position 071  'G' */
	0x1a,  /* position 072  'H' */
	0x1a,  /* position 073  'I' */
	0x1a,  /* position 074  'J' */
	0x1a,  /* position 075  'K' */
	0x1a,  /* position 076  'L' */
	0x1a,  /* position 077  'M' */
	0x1a,  /* position 078  'N' */
	0x1a,  /* position 079  'O' */
	0x1a,  /* position 080  'P' */
	0x1a,  /* position 081  'Q' */
	0x1a,  /* position 082  'R' */
	0x1a,  /* position 083  'S' */
	0x1a,  /* position 084  'T' */
	0x1a,  /* position 085  'U' */
	0x1a,  /* position 086  'V' */
	0x1a,  /* position 087  'W' */
	0x1a,  /* position 088  'X' */
	0x1a,  /* position 089  'Y' */
	0x1a,  /* position 090  'Z' */
	0x00,  /* position 091  '[' */
	0x00,  /* position 092  '\' */
	0x00,  /* position 093  ']' */
	0x00,  /* position 094  '^' */
	0x00,  /* position 095  '_' */
	0x00,  /* position 096  '`' */
	0x76,  /* position 097  'a' */
	0x76,  /* position 098  'b' */
	0x76,  /* position 099  'c' */
	0x76,  /* position 100  'd' */
	0x76,  /* position 101  'e' */
	0x76,  /* position 102  'f' */
	0x16,  /* position 103  'g' */
	0x16,  /* position 104  'h' */
	0x16,  /* position 105  'i' */
	0x16,  /* position 106  'j' */
	0x16,  /* position 107  'k' */
	0x16,  /* position 108  'l' */
	0x16,  /* position 109  'm' */
	0x16,  /* position 110  'n' */
	0x16,  /* position 111  'o' */
	0x16,  /* position 112  'p' */
	0x16,  /* position 113  'q' */
	0x16,  /* position 114  'r' */
	0x16,  /* position 115  's' */
	0x16,  /* position 116  't' */
	0x16,  /* position 117  'u' */
	0x16,  /* position 118  'v' */
	0x16,  /* position 119  'w' */
	0x16,  /* position 120  'x' */
	0x16,  /* position 121  'y' */
	0x16,  /* position 122  'z' */
	0x00,  /* position 123  '{' */
	0x00,  /* position 124  '|' */
	0x00,  /* position 125  '}' */
	0x00,  /* position 126  '~' */
	0x00,  /* position 127 */
	0x00,  /* position 128 */
	0x00,  /* position 129 */
	0x00,  /* position 130 */
	0x00,  /* position 131 */
	0x00,  /* position 132 */
	0x00,  /* position 133 */
	0x00,  /* position 134 */
	0x00,  /* position 135 */
	0x00,  /* position 136 */
	0x00,  /* position 137 */
	0x00,  /* position 138 */
	0x00,  /* position 139 */
	0x00,  /* position 140 */
	0x00,  /* position 141 */
	0x00,  /* position 142 */
	0x00,  /* position 143 */
	0x00,  /* position 144 */
	0x00,  /* position 145 */
	0x00,  /* position 146 */
	0x00,  /* position 147 */
	0x00,  /* position 148 */
	0x00,  /* position 149 */
	0x00,  /* position 150 */
	0x00,  /* position 151 */
	0x00,  /* position 152 */
	0x00,  /* position 153 */
	0x00,  /* position 154 */
	0x00,  /* position 155 */
	0x00,  /* position 156 */
	0x00,  /* position 157 */
	0x00,  /* position 158 */
	0x00,  /* position 159 */
	0x00,  /* position 160 */
	0x00,  /* position 161 */
	0x00,  /* position 162 */
	0x00,  /* position 163 */
	0x00,  /* position 164 */
	0x00,  /* position 165 */
	0x00,  /* position 166 */
	0x00,  /* position 167 */
	0x00,  /* position 168 */
	0x00,  /* position 169 */
	0x00,  /* position 170 */
	0x00,  /* position 171 */
	0x00,  /* position 172 */
	0x00,  /* position 173 */
	0x00,  /* position 174 */
	0x00,  /* position 175 */
	0x00,  /* position 176 */
	0x00,  /* position 177 */
	0x00,  /* position 178 */
	0x00,  /* position 179 */
	0x00,  /* position 180 */
	0x00,  /* position 181 */
	0x00,  /* position 182 */
	0x00,  /* position 183 */
	0x00,  /* position 184 */
	0x00,  /* position 185 */
	0x00,  /* position 186 */
	0x00,  /* position 187 */
	0x00,  /* position 188 */
	0x00,  /* position 189 */
	0x00,  /* position 190 */
	0x00,  /* position 191 */
	0x00,  /* position 192 */
	0x00,  /* position 193 */
	0x00,  /* position 194 */
	0x00,  /* position 195 */
	0x00,  /* position 196 */
	0x00,  /* position 197 */
	0x00,  /* position 198 */
	0x00,  /* position 199 */
	0x00,  /* position 200 */
	0x00,  /* position 201 */
	0x00,  /* position 202 */
	0x00,  /* position 203 */
	0x00,  /* position 204 */
	0x00,  /* position 205 */
	0x00,  /* position 206 */
	0x00,  /* position 207 */
	0x00,  /* position 208 */
	0x00,  /* position 209 */
	0x00,  /* position 210 */
	0x00,  /* position 211 */
	0x00,  /* position 212 */
	0x00,  /* position 213 */
	0x00,  /* position 214 */
	0x00,  /* position 215 */
	0x00,  /* position 216 */
	0x00,  /* position 217 */
	0x00,  /* position 218 */
	0x00,  /* position 219 */
	0x00,  /* position 220 */
	0x00,  /* position 221 */
	0x00,  /* position 222 */
	0x00,  /* position 223 */
	0x00,  /* position 224 */
	0x00,  /* position 225 */
	0x00,  /* position 226 */
	0x00,  /* position 227 */
	0x00,  /* position 228 */
	0x00,  /* position 229 */
	0x00,  /* position 230 */
	0x00,  /* position 231 */
	0x00,  /* position 232 */
	0x00,  /* position 233 */
	0x00,  /* position 234 */
	0x00,  /* position 235 */
	0x00,  /* position 236 */
	0x00,  /* position 237 */
	0x00,  /* position 238 */
	0x00,  /* position 239 */
	0x00,  /* position 240 */
	0x00,  /* position 241 */
	0x00,  /* position 242 */
	0x00,  /* position 243 */
	0x00,  /* position 244 */
	0x00,  /* position 245 */
	0x00,  /* position 246 */
	0x00,  /* position 247 */
	0x00,  /* position 248 */
	0x00,  /* position 249 */
	0x00,  /* position 250 */
	0x00,  /* position 251 */
	0x00,  /* position 252 */
	0x00,  /* position 253 */
	0x00,  /* position 254 */
	0x00,  /* position 255 */
}

var g_sipCharsets1 = [256]PS_BYTE{
	0x00,  /* position 000 */
	0x00,  /* position 001 */
	0x00,  /* position 002 */
	0x00,  /* position 003 */
	0x00,  /* position 004 */
	0x00,  /* position 005 */
	0x00,  /* position 006 */
	0x00,  /* position 007 */
	0x00,  /* position 008 */
	0x01,  /* position 009 */
	0x00,  /* position 010 */
	0x00,  /* position 011 */
	0x00,  /* position 012 */
	0x00,  /* position 013 */
	0x00,  /* position 014 */
	0x00,  /* position 015 */
	0x00,  /* position 016 */
	0x00,  /* position 017 */
	0x00,  /* position 018 */
	0x00,  /* position 019 */
	0x00,  /* position 020 */
	0x00,  /* position 021 */
	0x00,  /* position 022 */
	0x00,  /* position 023 */
	0x00,  /* position 024 */
	0x00,  /* position 025 */
	0x00,  /* position 026 */
	0x00,  /* position 027 */
	0x00,  /* position 028 */
	0x00,  /* position 029 */
	0x00,  /* position 030 */
	0x00,  /* position 031 */
	0x01,  /* position 032  ' ' */
	0x00,  /* position 033  '!' */
	0x00,  /* position 034  '"' */
	0x00,  /* position 035  '#' */
	0x00,  /* position 036  '$' */
	0x00,  /* position 037  '%' */
	0x00,  /* position 038  '&' */
	0x00,  /* position 039  ''' */
	0x00,  /* position 040  '(' */
	0x00,  /* position 041  ')' */
	0x00,  /* position 042  '*' */
	0x00,  /* position 043  '+' */
	0x00,  /* position 044  ',' */
	0x00,  /* position 045  '-' */
	0x00,  /* position 046  '.' */
	0x00,  /* position 047  '/' */
	0x00,  /* position 048  '0' */
	0x00,  /* position 049  '1' */
	0x00,  /* position 050  '2' */
	0x00,  /* position 051  '3' */
	0x00,  /* position 052  '4' */
	0x00,  /* position 053  '5' */
	0x00,  /* position 054  '6' */
	0x00,  /* position 055  '7' */
	0x00,  /* position 056  '8' */
	0x00,  /* position 057  '9' */
	0x00,  /* position 058  ':' */
	0x00,  /* position 059  ';' */
	0x00,  /* position 060  '<' */
	0x00,  /* position 061  '=' */
	0x00,  /* position 062  '>' */
	0x00,  /* position 063  '?' */
	0x00,  /* position 064  '@' */
	0x00,  /* position 065  'A' */
	0x00,  /* position 066  'B' */
	0x00,  /* position 067  'C' */
	0x00,  /* position 068  'D' */
	0x00,  /* position 069  'E' */
	0x00,  /* position 070  'F' */
	0x00,  /* position 071  'G' */
	0x00,  /* position 072  'H' */
	0x00,  /* position 073  'I' */
	0x00,  /* position 074  'J' */
	0x00,  /* position 075  'K' */
	0x00,  /* position 076  'L' */
	0x00,  /* position 077  'M' */
	0x00,  /* position 078  'N' */
	0x00,  /* position 079  'O' */
	0x00,  /* position 080  'P' */
	0x00,  /* position 081  'Q' */
	0x00,  /* position 082  'R' */
	0x00,  /* position 083  'S' */
	0x00,  /* position 084  'T' */
	0x00,  /* position 085  'U' */
	0x00,  /* position 086  'V' */
	0x00,  /* position 087  'W' */
	0x00,  /* position 088  'X' */
	0x00,  /* position 089  'Y' */
	0x00,  /* position 090  'Z' */
	0x00,  /* position 091  '[' */
	0x00,  /* position 092  '\' */
	0x00,  /* position 093  ']' */
	0x00,  /* position 094  '^' */
	0x00,  /* position 095  '_' */
	0x00,  /* position 096  '`' */
	0x00,  /* position 097  'a' */
	0x00,  /* position 098  'b' */
	0x00,  /* position 099  'c' */
	0x00,  /* position 100  'd' */
	0x00,  /* position 101  'e' */
	0x00,  /* position 102  'f' */
	0x00,  /* position 103  'g' */
	0x00,  /* position 104  'h' */
	0x00,  /* position 105  'i' */
	0x00,  /* position 106  'j' */
	0x00,  /* position 107  'k' */
	0x00,  /* position 108  'l' */
	0x00,  /* position 109  'm' */
	0x00,  /* position 110  'n' */
	0x00,  /* position 111  'o' */
	0x00,  /* position 112  'p' */
	0x00,  /* position 113  'q' */
	0x00,  /* position 114  'r' */
	0x00,  /* position 115  's' */
	0x00,  /* position 116  't' */
	0x00,  /* position 117  'u' */
	0x00,  /* position 118  'v' */
	0x00,  /* position 119  'w' */
	0x00,  /* position 120  'x' */
	0x00,  /* position 121  'y' */
	0x00,  /* position 122  'z' */
	0x00,  /* position 123  '{' */
	0x00,  /* position 124  '|' */
	0x00,  /* position 125  '}' */
	0x00,  /* position 126  '~' */
	0x00,  /* position 127 */
	0x00,  /* position 128 */
	0x00,  /* position 129 */
	0x00,  /* position 130 */
	0x00,  /* position 131 */
	0x00,  /* position 132 */
	0x00,  /* position 133 */
	0x00,  /* position 134 */
	0x00,  /* position 135 */
	0x00,  /* position 136 */
	0x00,  /* position 137 */
	0x00,  /* position 138 */
	0x00,  /* position 139 */
	0x00,  /* position 140 */
	0x00,  /* position 141 */
	0x00,  /* position 142 */
	0x00,  /* position 143 */
	0x00,  /* position 144 */
	0x00,  /* position 145 */
	0x00,  /* position 146 */
	0x00,  /* position 147 */
	0x00,  /* position 148 */
	0x00,  /* position 149 */
	0x00,  /* position 150 */
	0x00,  /* position 151 */
	0x00,  /* position 152 */
	0x00,  /* position 153 */
	0x00,  /* position 154 */
	0x00,  /* position 155 */
	0x00,  /* position 156 */
	0x00,  /* position 157 */
	0x00,  /* position 158 */
	0x00,  /* position 159 */
	0x00,  /* position 160 */
	0x00,  /* position 161 */
	0x00,  /* position 162 */
	0x00,  /* position 163 */
	0x00,  /* position 164 */
	0x00,  /* position 165 */
	0x00,  /* position 166 */
	0x00,  /* position 167 */
	0x00,  /* position 168 */
	0x00,  /* position 169 */
	0x00,  /* position 170 */
	0x00,  /* position 171 */
	0x00,  /* position 172 */
	0x00,  /* position 173 */
	0x00,  /* position 174 */
	0x00,  /* position 175 */
	0x00,  /* position 176 */
	0x00,  /* position 177 */
	0x00,  /* position 178 */
	0x00,  /* position 179 */
	0x00,  /* position 180 */
	0x00,  /* position 181 */
	0x00,  /* position 182 */
	0x00,  /* position 183 */
	0x00,  /* position 184 */
	0x00,  /* position 185 */
	0x00,  /* position 186 */
	0x00,  /* position 187 */
	0x00,  /* position 188 */
	0x00,  /* position 189 */
	0x00,  /* position 190 */
	0x00,  /* position 191 */
	0x00,  /* position 192 */
	0x00,  /* position 193 */
	0x00,  /* position 194 */
	0x00,  /* position 195 */
	0x00,  /* position 196 */
	0x00,  /* position 197 */
	0x00,  /* position 198 */
	0x00,  /* position 199 */
	0x00,  /* position 200 */
	0x00,  /* position 201 */
	0x00,  /* position 202 */
	0x00,  /* position 203 */
	0x00,  /* position 204 */
	0x00,  /* position 205 */
	0x00,  /* position 206 */
	0x00,  /* position 207 */
	0x00,  /* position 208 */
	0x00,  /* position 209 */
	0x00,  /* position 210 */
	0x00,  /* position 211 */
	0x00,  /* position 212 */
	0x00,  /* position 213 */
	0x00,  /* position 214 */
	0x00,  /* position 215 */
	0x00,  /* position 216 */
	0x00,  /* position 217 */
	0x00,  /* position 218 */
	0x00,  /* position 219 */
	0x00,  /* position 220 */
	0x00,  /* position 221 */
	0x00,  /* position 222 */
	0x00,  /* position 223 */
	0x00,  /* position 224 */
	0x00,  /* position 225 */
	0x00,  /* position 226 */
	0x00,  /* position 227 */
	0x00,  /* position 228 */
	0x00,  /* position 229 */
	0x00,  /* position 230 */
	0x00,  /* position 231 */
	0x00,  /* position 232 */
	0x00,  /* position 233 */
	0x00,  /* position 234 */
	0x00,  /* position 235 */
	0x00,  /* position 236 */
	0x00,  /* position 237 */
	0x00,  /* position 238 */
	0x00,  /* position 239 */
	0x00,  /* position 240 */
	0x00,  /* position 241 */
	0x00,  /* position 242 */
	0x00,  /* position 243 */
	0x00,  /* position 244 */
	0x00,  /* position 245 */
	0x00,  /* position 246 */
	0x00,  /* position 247 */
	0x00,  /* position 248 */
	0x00,  /* position 249 */
	0x00,  /* position 250 */
	0x00,  /* position 251 */
	0x00,  /* position 252 */
	0x00,  /* position 253 */
	0x00,  /* position 254 */
	0x00,  /* position 255 */
}

