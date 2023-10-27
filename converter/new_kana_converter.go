package converter

import (
	"errors"
	"fmt"
	"unicode/utf8"

	"golang.org/x/text/transform"
)

type KanaConverter struct {
	transform.NopResetter
	optr bool
	optR bool
	optn bool
	optN bool
	opta bool
	optA bool
	opts bool
	optS bool
	optk bool
	optK bool
	opth bool
	optH bool
	optc bool
	optC bool
	optV bool
}

func (r *KanaConverter) enableOptr() error {
	if r.optR {
		return fmt.Errorf("must not combine 'r' and 'R' flags")
	}
	if r.optA {
		return fmt.Errorf("must not combine 'r' and 'A' flags")
	}
	r.optr = true
	return nil
}

func (r *KanaConverter) enableOptR() error {
	if r.optr {
		return fmt.Errorf("must not combine 'r' and 'R' flags")
	}
	if r.opta {
		return fmt.Errorf("must not combine 'R' and 'a' flags")
	}
	r.optR = true
	return nil
}

func (r *KanaConverter) enableOptn() error {
	if r.optN {
		return fmt.Errorf("must not combine 'n' and 'N' flags")
	}
	if r.optA {
		return fmt.Errorf("must not combine 'n' and 'A' flags")
	}
	r.optn = true
	return nil
}

func (r *KanaConverter) enableOptN() error {
	if r.optn {
		return fmt.Errorf("must not combine 'n' and 'N' flags")
	}
	if r.opta {
		return fmt.Errorf("must not combine 'N' and 'a' flags")
	}
	r.optN = true
	return nil
}

func (r *KanaConverter) enableOpta() error {
	if r.optA {
		return fmt.Errorf("must not combine 'a' and 'A' flags")
	}
	if r.optR {
		return fmt.Errorf("must not combine 'R' and 'a' flags")
	}
	if r.optN {
		return fmt.Errorf("must not combine 'N' and 'a' flags")
	}
	r.opta = true
	return nil
}

func (r *KanaConverter) enableOptA() error {
	if r.opta {
		return fmt.Errorf("must not combine 'a' and 'A' flags")
	}
	if r.optr {
		return fmt.Errorf("must not combine 'r' and 'A' flags")
	}
	if r.optn {
		return fmt.Errorf("must not combine 'n' and 'A' flags")
	}
	r.optA = true
	return nil
}

func (r *KanaConverter) enableOpts() error {
	if r.optS {
		return fmt.Errorf("must not combine 's' and 'S' flags")
	}
	r.opts = true
	return nil
}

func (r *KanaConverter) enableOptS() error {
	if r.opts {
		return fmt.Errorf("must not combine 's' and 'S' flags")
	}
	r.optS = true
	return nil
}

func (r *KanaConverter) enableOptk() error {
	if r.optK {
		return fmt.Errorf("must not combine 'k' and 'K' flags")
	}
	if r.optc {
		return fmt.Errorf("must not combine 'k' and 'c' flags")
	}
	if r.optC {
		return fmt.Errorf("must not combine 'k' and 'C' flags")
	}

	r.optk = true
	return nil
}

func (r *KanaConverter) enableOptK() error {
	if r.optk {
		return fmt.Errorf("must not combine 'k' and 'K' flags")
	}
	if r.optH {
		return fmt.Errorf("must not combine 'K' and 'H' flags")
	}

	r.optK = true
	return nil
}

func (r *KanaConverter) enableOpth() error {
	if r.optH {
		return fmt.Errorf("must not combine 'h' and 'H' flags")
	}
	if r.optc {
		return fmt.Errorf("must not combine 'h' and 'c' flags")
	}
	if r.optC {
		return fmt.Errorf("must not combine 'h' and 'C' flags")
	}

	r.opth = true
	return nil
}

func (r *KanaConverter) enableOptH() error {
	if r.opth {
		return fmt.Errorf("must not combine 'h' and 'H' flags")
	}
	if r.optK {
		return fmt.Errorf("must not combine 'K' and 'H' flags")
	}

	r.optH = true
	return nil
}

func (r *KanaConverter) enableOptc() error {
	if r.optC {
		return fmt.Errorf("must not combine 'c' and 'C' flags")
	}
	if r.optk {
		return fmt.Errorf("must not combine 'k' and 'c' flags")
	}
	if r.opth {
		return fmt.Errorf("must not combine 'h' and 'c' flags")
	}

	r.optc = true
	return nil
}

func (r *KanaConverter) enableOptC() error {
	if r.optc {
		return fmt.Errorf("must not combine 'c' and 'C' flags")
	}
	if r.optk {
		return fmt.Errorf("must not combine 'k' and 'C' flags")
	}
	if r.opth {
		return fmt.Errorf("must not combine 'h' and 'C' flags")
	}

	r.optC = true
	return nil
}

func (r *KanaConverter) enableOptV() error {
	r.optV = true
	return nil
}

func NewKanaConverter(mode string) (*KanaConverter, error) {
	o := &KanaConverter{}
	for _, char := range mode {
		var err error
		switch char {
		case rune('r'):
			err = o.enableOptr()
		case rune('R'):
			err = o.enableOptR()
		case rune('n'):
			err = o.enableOptn()
		case rune('N'):
			err = o.enableOptN()
		case rune('a'):
			err = o.enableOpta()
		case rune('A'):
			err = o.enableOptA()
		case rune('s'):
			err = o.enableOpts()
		case rune('S'):
			err = o.enableOptS()
		case rune('k'):
			err = o.enableOptk()
		case rune('K'):
			err = o.enableOptK()
		case rune('h'):
			err = o.enableOpth()
		case rune('H'):
			err = o.enableOptH()
		case rune('c'):
			err = o.enableOptc()
		case rune('C'):
			err = o.enableOptC()
		case rune('V'):
			err = o.enableOptV()
		}
		if err != nil {
			return nil, err
		}
	}

	return o, nil
}

var ErrInvalidUTF8 = errors.New("invalid UTF-8 character")

func (t *KanaConverter) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	r, size := rune(0), 0
	before := rune(0)
	buf := make([]byte, 4)
loop:
	for ; nSrc < len(src); nSrc += size {
		r = rune(src[nSrc])

		// Decode a 1-byte rune.
		if r < utf8.RuneSelf {
			size = 1
		} else {
			// Decode a multi-byte rune.
			r, size = utf8.DecodeRune(src[nSrc:])
			if size == 1 {
				// All valid runes of size 1 (those below utf8.RuneSelf) were
				// handled above. We have invalid UTF-8 or we haven't seen the
				// full character yet.
				if !atEOF && !utf8.FullRune(src[nSrc:]) {
					err = transform.ErrShortSrc
					break loop
				}
				err = ErrInvalidUTF8
				break loop
			}
		}

		if t.optK && t.optV {
			if before == 0 {
				switch r {
				case 'ｳ', 'ｶ', 'ｷ', 'ｸ', 'ｹ', 'ｺ', 'ｻ', 'ｼ', 'ｽ', 'ｾ', 'ｿ', 'ﾀ', 'ﾁ', 'ﾂ', 'ﾃ', 'ﾄ', 'ﾊ', 'ﾋ', 'ﾌ', 'ﾍ', 'ﾎ':
					before = r
					continue
				}
			}
			switch r {
			case 'ﾞ':
				switch before {
				case 'ｳ':
					r = 'ヴ'
					goto writeZenkakuHiraganaOrZenkakuKatakana
				case 'ｶ':
					r = 'ガ'
					goto writeZenkakuHiraganaOrZenkakuKatakana
				case 'ｷ':
					r = 'ギ'
					goto writeZenkakuHiraganaOrZenkakuKatakana
				case 'ｸ':
					r = 'グ'
					goto writeZenkakuHiraganaOrZenkakuKatakana
				case 'ｹ':
					r = 'ゲ'
					goto writeZenkakuHiraganaOrZenkakuKatakana
				case 'ｺ':
					r = 'ゴ'
					goto writeZenkakuHiraganaOrZenkakuKatakana
				case 'ｻ':
					r = 'ザ'
					goto writeZenkakuHiraganaOrZenkakuKatakana
				case 'ｼ':
					r = 'ジ'
					goto writeZenkakuHiraganaOrZenkakuKatakana
				case 'ｽ':
					r = 'ズ'
					goto writeZenkakuHiraganaOrZenkakuKatakana
				case 'ｾ':
					r = 'ゼ'
					goto writeZenkakuHiraganaOrZenkakuKatakana
				case 'ｿ':
					r = 'ゾ'
					goto writeZenkakuHiraganaOrZenkakuKatakana
				case 'ﾀ':
					r = 'ダ'
					goto writeZenkakuHiraganaOrZenkakuKatakana
				case 'ﾁ':
					r = 'ヂ'
					goto writeZenkakuHiraganaOrZenkakuKatakana
				case 'ﾂ':
					r = 'ヅ'
					goto writeZenkakuHiraganaOrZenkakuKatakana
				case 'ﾃ':
					r = 'デ'
					goto writeZenkakuHiraganaOrZenkakuKatakana
				case 'ﾄ':
					r = 'ド'
					goto writeZenkakuHiraganaOrZenkakuKatakana
				case 'ﾊ':
					r = 'バ'
					goto writeZenkakuHiraganaOrZenkakuKatakana
				case 'ﾋ':
					r = 'ビ'
					goto writeZenkakuHiraganaOrZenkakuKatakana
				case 'ﾌ':
					r = 'ブ'
					goto writeZenkakuHiraganaOrZenkakuKatakana
				case 'ﾍ':
					r = 'ベ'
					goto writeZenkakuHiraganaOrZenkakuKatakana
				case 'ﾎ':
					r = 'ボ'
					goto writeZenkakuHiraganaOrZenkakuKatakana
				}
			case 'ﾟ':
				switch before {
				case 'ﾊ':
					r = 'パ'
					goto writeZenkakuHiraganaOrZenkakuKatakana
				case 'ﾋ':
					r = 'ピ'
					goto writeZenkakuHiraganaOrZenkakuKatakana
				case 'ﾌ':
					r = 'プ'
					goto writeZenkakuHiraganaOrZenkakuKatakana
				case 'ﾍ':
					r = 'ペ'
					goto writeZenkakuHiraganaOrZenkakuKatakana
				case 'ﾎ':
					r = 'ポ'
					goto writeZenkakuHiraganaOrZenkakuKatakana
				}
			}
			if before != 0 {
				before = HankakuKatakanaToZenkakuKatakana2(before)
				utf8.EncodeRune(buf, before)
				for i := 0; i < 3; i++ {
					dst[nDst+i] = buf[i]
				}
				nDst += 3
				before = 0
			}
		}

		if t.optH && t.optV {
			if before == 0 {
				switch r {
				case 'ｶ', 'ｷ', 'ｸ', 'ｹ', 'ｺ', 'ｻ', 'ｼ', 'ｽ', 'ｾ', 'ｿ', 'ﾀ', 'ﾁ', 'ﾂ', 'ﾃ', 'ﾄ', 'ﾊ', 'ﾋ', 'ﾌ', 'ﾍ', 'ﾎ':
					before = r
					continue
				}
			}
			switch r {
			case 'ﾞ':
				switch before {
				case 'ｶ':
					r = 'が'
					goto writeZenkakuHiraganaOrZenkakuKatakana
				case 'ｷ':
					r = 'ぎ'
					goto writeZenkakuHiraganaOrZenkakuKatakana
				case 'ｸ':
					r = 'ぐ'
					goto writeZenkakuHiraganaOrZenkakuKatakana
				case 'ｹ':
					r = 'げ'
					goto writeZenkakuHiraganaOrZenkakuKatakana
				case 'ｺ':
					r = 'ご'
					goto writeZenkakuHiraganaOrZenkakuKatakana
				case 'ｻ':
					r = 'ざ'
					goto writeZenkakuHiraganaOrZenkakuKatakana
				case 'ｼ':
					r = 'じ'
					goto writeZenkakuHiraganaOrZenkakuKatakana
				case 'ｽ':
					r = 'ず'
					goto writeZenkakuHiraganaOrZenkakuKatakana
				case 'ｾ':
					r = 'ぜ'
					goto writeZenkakuHiraganaOrZenkakuKatakana
				case 'ｿ':
					r = 'ぞ'
					goto writeZenkakuHiraganaOrZenkakuKatakana
				case 'ﾀ':
					r = 'だ'
					goto writeZenkakuHiraganaOrZenkakuKatakana
				case 'ﾁ':
					r = 'ぢ'
					goto writeZenkakuHiraganaOrZenkakuKatakana
				case 'ﾂ':
					r = 'づ'
					goto writeZenkakuHiraganaOrZenkakuKatakana
				case 'ﾃ':
					r = 'で'
					goto writeZenkakuHiraganaOrZenkakuKatakana
				case 'ﾄ':
					r = 'ど'
					goto writeZenkakuHiraganaOrZenkakuKatakana
				case 'ﾊ':
					r = 'ば'
					goto writeZenkakuHiraganaOrZenkakuKatakana
				case 'ﾋ':
					r = 'び'
					goto writeZenkakuHiraganaOrZenkakuKatakana
				case 'ﾌ':
					r = 'ぶ'
					goto writeZenkakuHiraganaOrZenkakuKatakana
				case 'ﾍ':
					r = 'べ'
					goto writeZenkakuHiraganaOrZenkakuKatakana
				case 'ﾎ':
					r = 'ぼ'
					goto writeZenkakuHiraganaOrZenkakuKatakana
				}
			case 'ﾟ':
				switch before {
				case 'ﾊ':
					r = 'ぱ'
					goto writeZenkakuHiraganaOrZenkakuKatakana
				case 'ﾋ':
					r = 'ぴ'
					goto writeZenkakuHiraganaOrZenkakuKatakana
				case 'ﾌ':
					r = 'ぷ'
					goto writeZenkakuHiraganaOrZenkakuKatakana
				case 'ﾍ':
					r = 'ぺ'
					goto writeZenkakuHiraganaOrZenkakuKatakana
				case 'ﾎ':
					r = 'ぽ'
					goto writeZenkakuHiraganaOrZenkakuKatakana
				}
			}
			if before != 0 {
				before = HankakuKatakanaToZenkakuHiragana2(before)
				utf8.EncodeRune(buf, before)
				for i := 0; i < 3; i++ {
					dst[nDst+i] = buf[i]
				}
				nDst += 3
				before = 0
			}
		}

		if t.optr && ((r >= '\uFF21' && r <= '\uFF3A') || (r >= '\uFF41' && r <= '\uFF5A')) {
			goto writeZenkakuEnglishToHankakuEnglish
		}
		if t.optR && ((r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z')) {
			goto writeHankakuEnglishToZenkakuEnglish
		}
		if t.optn && (r >= '\uFF10' && r <= '\uFF19') {
			goto writeZenkakuNumberToHankakuNumber
		}
		if t.optN && (r >= '0' && r <= '9') {
			goto writeHankakuNumberToZenkakuNumber
		}
		if t.opta && (r >= '\uFF01' && r <= '\uFF5E' && !(r == '\uFF02' || r == '\uFF07' || r == '\uFF3C' || r == '\uFF5E')) {
			goto writeZenkakuEnglishNumberToHankakuEnglishNumber
		}
		if t.optA && (r >= '\u0021' && r <= '\u007E' && !(r == '\u0022' || r == '\u0027' || r == '\u005C' || r == '\u007E')) {
			goto writeHankakuEnglishNumberToZenkakuEnglishNumber
		}
		if t.opts && r == '\u3000' {
			goto writeZenkakuSpaceToHankakuSpace
		}
		if t.optS && r == ' ' {
			goto writeHankakuSpaceToZenkakuSpace
		}
		if t.optk && (r == '\u3001' || r == '\u3002' || r == '\u300C' || r == '\u300D' || r == '\u309B' || r == '\u309C' || (r >= '\u30A1' && r <= '\u30F4') || r == '\u30FB' || r == '\u30FC') {
			goto writeZenkakuKatakanaToHankakuKatakana
		}
		if t.optK && r >= '\uFF61' && r <= '\uFF9F' {
			goto writeHankakuKatakanaToZenkakuKatakana
		}
		if t.opth && (r == '\u3001' || r == '\u3002' || r == '\u300C' || r == '\u300D' || (r >= '\u3041' && r <= '\u3093') || r == '\u309B' || r == '\u309C' || r == '\u30FB' || r == '\u30FC') {
			goto writeZenkakuHiraganaToHankakuKatakana
		}
		if t.optH && (r >= '\uFF61' && r <= '\uFF9F') {
			goto writeHankakuKatakanaToZenkakuHiragana
		}
		if t.optc && ((r >= '\u30A1' && r <= '\u30F4') || r == '\u30FD' || r == '\u30FE') {
			goto writeZenkakuKatakanaToZenkakuHiragana
		}
		if t.optC && ((r >= '\u3041' && r <= '\u3093') || r == '\u309D' || r == '\u309E') {
			goto writeZenkakuHiraganaToZenkakuKatakana
		}

		// write as is
		if nDst+size > len(dst) {
			err = transform.ErrShortDst
			break
		}
		for n := 0; n < size; n++ {
			dst[nDst+n] = src[nSrc+n]
		}
		nDst += size
		continue
	writeZenkakuHiraganaOrZenkakuKatakana:
		if nDst+3 > len(dst) {
			err = transform.ErrShortDst
			break
		}
		before = 0
		utf8.EncodeRune(buf, r)
		for i := 0; i < 3; i++ {
			dst[nDst+i] = buf[i]
		}
		nDst += 3
		continue
	writeZenkakuEnglishToHankakuEnglish:
		if nDst >= len(dst) {
			err = transform.ErrShortDst
			break
		}
		dst[nDst] = uint8('A' + r - '\uFF21')
		nDst++
		continue
	writeHankakuEnglishToZenkakuEnglish:
		if nDst+3 > len(dst) {
			err = transform.ErrShortDst
			break
		}
		r = '\uFF21' + r - 'A'
		utf8.EncodeRune(buf, r)
		for i := 0; i < 3; i++ {
			dst[nDst+i] = buf[i]
		}
		nDst += 3
		continue
	writeZenkakuNumberToHankakuNumber:
		if nDst >= len(dst) {
			err = transform.ErrShortDst
			break
		}
		dst[nDst] = uint8('0' + r - '\uFF10')
		nDst++
		continue
	writeHankakuNumberToZenkakuNumber:
		if nDst+3 > len(dst) {
			err = transform.ErrShortDst
			break
		}
		r = '\uFF10' + r - '0'
		utf8.EncodeRune(buf, r)
		for i := 0; i < 3; i++ {
			dst[nDst+i] = buf[i]
		}
		nDst += 3
		continue
	writeZenkakuEnglishNumberToHankakuEnglishNumber:
		if nDst >= len(dst) {
			err = transform.ErrShortDst
			break
		}
		dst[nDst] = uint8('!' + r - '\uFF01')
		nDst++
		continue
	writeHankakuEnglishNumberToZenkakuEnglishNumber:
		if nDst+3 > len(dst) {
			err = transform.ErrShortDst
			break
		}
		r = '\uFF01' + r - '!'
		utf8.EncodeRune(buf, r)
		for i := 0; i < 3; i++ {
			dst[nDst+i] = buf[i]
		}

		nDst += 3
		continue
	writeZenkakuSpaceToHankakuSpace:
		if nDst >= len(dst) {
			err = transform.ErrShortDst
			break
		}
		dst[nDst] = uint8(' ')
		nDst++
		continue
	writeHankakuSpaceToZenkakuSpace:
		if nDst+3 > len(dst) {
			err = transform.ErrShortDst
			break
		}
		dst[nDst+0] = 0xE3
		dst[nDst+1] = 0x80
		dst[nDst+2] = 0x80
		nDst += 3
		continue
	writeZenkakuKatakanaToHankakuKatakana:
		{
			r1, r2 := ZenkakuKatakanaToHankakuKatakana2(r)
			if r2 == 0 {
				if nDst+3 > len(dst) {
					err = transform.ErrShortDst
					break
				}
				utf8.EncodeRune(buf, r1)
				for i := 0; i < 3; i++ {
					dst[nDst+i] = buf[i]
				}
				nDst += 3
				continue
			}
			if nDst+3*2 > len(dst) {
				err = transform.ErrShortDst
				break
			}
			utf8.EncodeRune(buf, r1)
			for i := 0; i < 3; i++ {
				dst[nDst+i] = buf[i]
			}
			utf8.EncodeRune(buf, r2)
			for i := 0; i < 3; i++ {
				dst[nDst+3+i] = buf[i]
			}
			nDst += 6
		}
		continue
	writeHankakuKatakanaToZenkakuKatakana:
		if nDst+3 > len(dst) {
			err = transform.ErrShortDst
			break
		}
		r = HankakuKatakanaToZenkakuKatakana2(r)
		utf8.EncodeRune(buf, r)
		for i := 0; i < 3; i++ {
			dst[nDst+i] = buf[i]
		}
		nDst += 3
		continue
	writeZenkakuHiraganaToHankakuKatakana:
		{
			r1, r2 := ZenkakuHiraganaToHankakuKatakana2(r)
			if r2 == 0 {
				if nDst+3 > len(dst) {
					err = transform.ErrShortDst
					break
				}
				utf8.EncodeRune(buf, r1)
				for i := 0; i < 3; i++ {
					dst[nDst+i] = buf[i]
				}
				nDst += 3
				continue
			}
			if nDst+3*2 > len(dst) {
				err = transform.ErrShortDst
				break
			}
			utf8.EncodeRune(buf, r1)
			for i := 0; i < 3; i++ {
				dst[nDst+i] = buf[i]
			}
			utf8.EncodeRune(buf, r2)
			for i := 0; i < 3; i++ {
				dst[nDst+3+i] = buf[i]
			}
			nDst += 6
		}
		continue

	writeHankakuKatakanaToZenkakuHiragana:
		if nDst+3 > len(dst) {
			err = transform.ErrShortDst
			break
		}
		r = HankakuKatakanaToZenkakuHiragana2(r)
		utf8.EncodeRune(buf, r)
		for i := 0; i < 3; i++ {
			dst[nDst+i] = buf[i]
		}
		nDst += 3
		continue
	writeZenkakuKatakanaToZenkakuHiragana:
		if nDst+3 > len(dst) {
			err = transform.ErrShortDst
			break
		}
		r = '\u3041' + r - '\u30A1'
		utf8.EncodeRune(buf, r)
		for i := 0; i < 3; i++ {
			dst[nDst+i] = buf[i]
		}
		nDst += 3
		continue
	writeZenkakuHiraganaToZenkakuKatakana:
		if nDst+3 > len(dst) {
			err = transform.ErrShortDst
			break
		}
		r = '\u30A1' + r - '\u3041'
		utf8.EncodeRune(buf, r)
		for i := 0; i < 3; i++ {
			dst[nDst+i] = buf[i]
		}
		nDst += 3
		continue
	}

	if before != 0 {
		if t.optK {
			fmt.Println("optK")
			before = HankakuKatakanaToZenkakuKatakana2(before)
		}
		if t.optH {
			fmt.Println("optH")
			before = HankakuKatakanaToZenkakuHiragana2(before)
		}
		utf8.EncodeRune(buf, before)
		for i := 0; i < 3; i++ {
			dst[nDst+i] = buf[i]
		}
		nDst += 3
	}
	return nDst, nSrc, err
}

const hankaku_voiced rune = '\uFF9E'
const hankaku_semi_voiced rune = '\uFF9F'

func ZenkakuKatakanaToHankakuKatakana2(r rune) (rune, rune) {
	switch r {
	case '\u3001':
		return '\uFF64', 0
	case '\u3002':
		return '\uFF61', 0
	case '\u300C':
		return '\uFF62', 0
	case '\u300D':
		return '\uFF63', 0
	case '\u309B':
		return hankaku_voiced, 0
	case '\u309C':
		return hankaku_semi_voiced, 0
	case '\u30A1':
		return '\uFF67', 0
	case '\u30A2':
		return '\uFF71', 0
	case '\u30A3':
		return '\uFF68', 0
	case '\u30A4':
		return '\uFF72', 0
	case '\u30A5':
		return '\uFF69', 0
	case '\u30A6':
		return '\uFF73', 0
	case '\u30A7':
		return '\uFF6A', 0
	case '\u30A8':
		return '\uFF74', 0
	case '\u30A9':
		return '\uFF6B', 0
	case '\u30AA':
		return '\uFF75', 0
	case '\u30AB':
		return '\uFF76', 0
	case '\u30AC':
		return '\uFF76', hankaku_voiced
	case '\u30AD':
		return '\uFF77', 0
	case '\u30AE':
		return '\uFF77', hankaku_voiced
	case '\u30AF':
		return '\uFF78', 0
	case '\u30B0':
		return '\uFF78', hankaku_voiced
	case '\u30B1':
		return '\uFF79', 0
	case '\u30B2':
		return '\uFF79', hankaku_voiced
	case '\u30B3':
		return '\uFF7A', 0
	case '\u30B4':
		return '\uFF7A', hankaku_voiced
	case '\u30B5':
		return '\uFF7B', 0
	case '\u30B6':
		return '\uFF7B', hankaku_voiced
	case '\u30B7':
		return '\uFF7C', 0
	case '\u30B8':
		return '\uFF7C', hankaku_voiced
	case '\u30B9':
		return '\uFF7D', 0
	case '\u30BA':
		return '\uFF7D', hankaku_voiced
	case '\u30BB':
		return '\uFF7E', 0
	case '\u30BC':
		return '\uFF7E', hankaku_voiced
	case '\u30BD':
		return '\uFF7F', 0
	case '\u30BE':
		return '\uFF7F', hankaku_voiced
	case '\u30BF':
		return '\uFF80', 0
	case '\u30C0':
		return '\uFF80', hankaku_voiced
	case '\u30C1':
		return '\uFF81', 0
	case '\u30C2':
		return '\uFF81', hankaku_voiced
	case '\u30C3':
		return '\uFF6F', 0
	case '\u30C4':
		return '\uFF82', 0
	case '\u30C5':
		return '\uFF82', hankaku_voiced
	case '\u30C6':
		return '\uFF83', 0
	case '\u30C7':
		return '\uFF83', hankaku_voiced
	case '\u30C8':
		return '\uFF84', 0
	case '\u30C9':
		return '\uFF84', hankaku_voiced
	case '\u30CA':
		return '\uFF85', 0
	case '\u30CB':
		return '\uFF86', 0
	case '\u30CC':
		return '\uFF87', 0
	case '\u30CD':
		return '\uFF88', 0
	case '\u30CE':
		return '\uFF89', 0
	case '\u30CF':
		return '\uFF8A', 0
	case '\u30D0':
		return '\uFF8A', hankaku_voiced
	case '\u30D1':
		return '\uFF8A', hankaku_semi_voiced
	case '\u30D2':
		return '\uFF8B', 0
	case '\u30D3':
		return '\uFF8B', hankaku_voiced
	case '\u30D4':
		return '\uFF8B', hankaku_semi_voiced
	case '\u30D5':
		return '\uFF8C', 0
	case '\u30D6':
		return '\uFF8C', hankaku_voiced
	case '\u30D7':
		return '\uFF8C', hankaku_semi_voiced
	case '\u30D8':
		return '\uFF8D', 0
	case '\u30D9':
		return '\uFF8D', hankaku_voiced
	case '\u30DA':
		return '\uFF8D', hankaku_semi_voiced
	case '\u30DB':
		return '\uFF8E', 0
	case '\u30DC':
		return '\uFF8E', hankaku_voiced
	case '\u30DD':
		return '\uFF8E', hankaku_semi_voiced
	case '\u30DE':
		return '\uFF8F', 0
	case '\u30DF':
		return '\uFF90', 0
	case '\u30E0':
		return '\uFF91', 0
	case '\u30E1':
		return '\uFF92', 0
	case '\u30E2':
		return '\uFF93', 0
	case '\u30E3':
		return '\uFF6C', 0
	case '\u30E4':
		return '\uFF94', 0
	case '\u30E5':
		return '\uFF6D', 0
	case '\u30E6':
		return '\uFF95', 0
	case '\u30E7':
		return '\uFF6E', 0
	case '\u30E8':
		return '\uFF96', 0
	case '\u30E9':
		return '\uFF97', 0
	case '\u30EA':
		return '\uFF98', 0
	case '\u30EB':
		return '\uFF99', 0
	case '\u30EC':
		return '\uFF9A', 0
	case '\u30ED':
		return '\uFF9B', 0
	case '\u30EE':
		return '\uFF9C', 0
	case '\u30EF':
		return '\uFF9C', 0
	case '\u30F0':
		return '\uFF72', 0
	case '\u30F1':
		return '\uFF74', 0
	case '\u30F2':
		return '\uFF66', 0
	case '\u30F3':
		return '\uFF9D', 0
	case '\u30F4':
		return '\uFF73', hankaku_voiced
	case '\u30FB':
		return '\uFF65', 0
	case '\u30FC':
		return '\uFF70', 0
	}
	return r, 0
}

func HankakuKatakanaToZenkakuKatakana2(r rune) rune {
	switch r {
	case '\uFF61':
		return '\u3002'
	case '\uFF62':
		return '\u300C'
	case '\uFF63':
		return '\u300D'
	case '\uFF64':
		return '\u3001'
	case '\uFF65':
		return '\u30FB'
	case '\uFF66':
		return '\u30F2'
	case '\uFF67':
		return '\u30A1'
	case '\uFF68':
		return '\u30A3'
	case '\uFF69':
		return '\u30A5'
	case '\uFF6A':
		return '\u30A7'
	case '\uFF6B':
		return '\u30A9'
	case '\uFF6C':
		return '\u30E3'
	case '\uFF6D':
		return '\u30E5'
	case '\uFF6E':
		return '\u30E7'
	case '\uFF6F':
		return '\u30C3'
	case '\uFF70':
		return '\u30FC'
	case '\uFF71':
		return '\u30A2'
	case '\uFF72':
		return '\u30A4'
	case '\uFF73':
		return '\u30A6'
	case '\uFF74':
		return '\u30A8'
	case '\uFF75':
		return '\u30AA'
	case '\uFF76':
		return '\u30AB'
	case '\uFF77':
		return '\u30AD'
	case '\uFF78':
		return '\u30AF'
	case '\uFF79':
		return '\u30B1'
	case '\uFF7A':
		return '\u30B3'
	case '\uFF7B':
		return '\u30B5'
	case '\uFF7C':
		return '\u30B7'
	case '\uFF7D':
		return '\u30B9'
	case '\uFF7E':
		return '\u30BB'
	case '\uFF7F':
		return '\u30BD'
	case '\uFF80':
		return '\u30BF'
	case '\uFF81':
		return '\u30C1'
	case '\uFF82':
		return '\u30C4'
	case '\uFF83':
		return '\u30C6'
	case '\uFF84':
		return '\u30C8'
	case '\uFF85':
		return '\u30CA'
	case '\uFF86':
		return '\u30CB'
	case '\uFF87':
		return '\u30CC'
	case '\uFF88':
		return '\u30CD'
	case '\uFF89':
		return '\u30CE'
	case '\uFF8A':
		return '\u30CF'
	case '\uFF8B':
		return '\u30D2'
	case '\uFF8C':
		return '\u30D5'
	case '\uFF8D':
		return '\u30D8'
	case '\uFF8E':
		return '\u30DB'
	case '\uFF8F':
		return '\u30DE'
	case '\uFF90':
		return '\u30DF'
	case '\uFF91':
		return '\u30E0'
	case '\uFF92':
		return '\u30E1'
	case '\uFF93':
		return '\u30E2'
	case '\uFF94':
		return '\u30E4'
	case '\uFF95':
		return '\u30E6'
	case '\uFF96':
		return '\u30E8'
	case '\uFF97':
		return '\u30E9'
	case '\uFF98':
		return '\u30EA'
	case '\uFF99':
		return '\u30EB'
	case '\uFF9A':
		return '\u30EC'
	case '\uFF9B':
		return '\u30ED'
	case '\uFF9C':
		return '\u30EF'
	case '\uFF9D':
		return '\u30F3'
	case '\uFF9E':
		return '\u309B'
	case '\uFF9F':
		return '\u309C'
	}
	return r
}

func ZenkakuHiraganaToHankakuKatakana2(r rune) (rune, rune) {
	switch r {
	case '\u3001':
		return '\uFF64', 0
	case '\u3002':
		return '\uFF61', 0
	case '\u300C':
		return '\uFF62', 0
	case '\u300D':
		return '\uFF63', 0
	case '\u309B':
		return hankaku_voiced, 0
	case '\u309C':
		return hankaku_semi_voiced, 0
	case '\u3041':
		return '\uFF67', 0
	case '\u3042':
		return '\uFF71', 0
	case '\u3043':
		return '\uFF68', 0
	case '\u3044':
		return '\uFF72', 0
	case '\u3045':
		return '\uFF69', 0
	case '\u3046':
		return '\uFF73', 0
	case '\u3047':
		return '\uFF6A', 0
	case '\u3048':
		return '\uFF74', 0
	case '\u3049':
		return '\uFF6B', 0
	case '\u304A':
		return '\uFF75', 0
	case '\u304B':
		return '\uFF76', 0
	case '\u304C':
		return '\uFF76', hankaku_voiced
	case '\u304D':
		return '\uFF77', 0
	case '\u304E':
		return '\uFF77', hankaku_voiced
	case '\u304F':
		return '\uFF78', 0
	case '\u3050':
		return '\uFF78', hankaku_voiced
	case '\u3051':
		return '\uFF79', 0
	case '\u3052':
		return '\uFF79', hankaku_voiced
	case '\u3053':
		return '\uFF7A', 0
	case '\u3054':
		return '\uFF7A', hankaku_voiced
	case '\u3055':
		return '\uFF7B', 0
	case '\u3056':
		return '\uFF7B', hankaku_voiced
	case '\u3057':
		return '\uFF7C', 0
	case '\u3058':
		return '\uFF7C', hankaku_voiced
	case '\u3059':
		return '\uFF7D', 0
	case '\u305A':
		return '\uFF7D', hankaku_voiced
	case '\u305B':
		return '\uFF7E', 0
	case '\u305C':
		return '\uFF7E', hankaku_voiced
	case '\u305D':
		return '\uFF7F', 0
	case '\u305E':
		return '\uFF7F', hankaku_voiced
	case '\u305F':
		return '\uFF80', 0
	case '\u3060':
		return '\uFF80', hankaku_voiced
	case '\u3061':
		return '\uFF81', 0
	case '\u3062':
		return '\uFF81', hankaku_voiced
	case '\u3063':
		return '\uFF6F', 0
	case '\u3064':
		return '\uFF82', 0
	case '\u3065':
		return '\uFF82', hankaku_voiced
	case '\u3066':
		return '\uFF83', 0
	case '\u3067':
		return '\uFF83', hankaku_voiced
	case '\u3068':
		return '\uFF84', 0
	case '\u3069':
		return '\uFF84', hankaku_voiced
	case '\u306A':
		return '\uFF85', 0
	case '\u306B':
		return '\uFF86', 0
	case '\u306C':
		return '\uFF87', 0
	case '\u306D':
		return '\uFF88', 0
	case '\u306E':
		return '\uFF89', 0
	case '\u306F':
		return '\uFF8A', 0
	case '\u3070':
		return '\uFF8A', hankaku_voiced
	case '\u3071':
		return '\uFF8A', hankaku_semi_voiced
	case '\u3072':
		return '\uFF8B', 0
	case '\u3073':
		return '\uFF8B', hankaku_voiced
	case '\u3074':
		return '\uFF8B', hankaku_semi_voiced
	case '\u3075':
		return '\uFF8C', 0
	case '\u3076':
		return '\uFF8C', hankaku_voiced
	case '\u3077':
		return '\uFF8C', hankaku_semi_voiced
	case '\u3078':
		return '\uFF8D', 0
	case '\u3079':
		return '\uFF8D', hankaku_voiced
	case '\u307A':
		return '\uFF8D', hankaku_semi_voiced
	case '\u307B':
		return '\uFF8E', 0
	case '\u307C':
		return '\uFF8E', hankaku_voiced
	case '\u307D':
		return '\uFF8E', hankaku_semi_voiced
	case '\u307E':
		return '\uFF8F', 0
	case '\u307F':
		return '\uFF90', 0
	case '\u3080':
		return '\uFF91', 0
	case '\u3081':
		return '\uFF92', 0
	case '\u3082':
		return '\uFF93', 0
	case '\u3083':
		return '\uFF6C', 0
	case '\u3084':
		return '\uFF94', 0
	case '\u3085':
		return '\uFF6D', 0
	case '\u3086':
		return '\uFF95', 0
	case '\u3087':
		return '\uFF6E', 0
	case '\u3088':
		return '\uFF96', 0
	case '\u3089':
		return '\uFF97', 0
	case '\u308A':
		return '\uFF98', 0
	case '\u308B':
		return '\uFF99', 0
	case '\u308C':
		return '\uFF9A', 0
	case '\u308D':
		return '\uFF9B', 0
	case '\u308E':
		return '\uFF9C', 0
	case '\u308F':
		return '\uFF9C', 0
	case '\u3090':
		return '\uFF72', 0
	case '\u3091':
		return '\uFF74', 0
	case '\u3092':
		return '\uFF66', 0
	case '\u3093':
		return '\uFF9D', 0
	case '\u30FB':
		return '\uFF65', 0
	case '\u30FC':
		return '\uFF70', 0
	}
	return r, 0
}

func HankakuKatakanaToZenkakuHiragana2(r rune) rune {
	switch r {
	case '\uFF61':
		return '\u3002'
	case '\uFF62':
		return '\u300C'
	case '\uFF63':
		return '\u300D'
	case '\uFF64':
		return '\u3001'
	case '\uFF65':
		return '\u30FB'
	case '\uFF66':
		return '\u3092'
	case '\uFF67':
		return '\u3041'
	case '\uFF68':
		return '\u3043'
	case '\uFF69':
		return '\u3045'
	case '\uFF6A':
		return '\u3047'
	case '\uFF6B':
		return '\u3049'
	case '\uFF6C':
		return '\u3083'
	case '\uFF6D':
		return '\u3085'
	case '\uFF6E':
		return '\u3087'
	case '\uFF6F':
		return '\u3063'
	case '\uFF70':
		return '\u30FC'
	case '\uFF71':
		return '\u3042'
	case '\uFF72':
		return '\u3044'
	case '\uFF73':
		return '\u3046'
	case '\uFF74':
		return '\u3048'
	case '\uFF75':
		return '\u304A'
	case '\uFF76':
		return '\u304B'
	case '\uFF77':
		return '\u304D'
	case '\uFF78':
		return '\u304F'
	case '\uFF79':
		return '\u3051'
	case '\uFF7A':
		return '\u3053'
	case '\uFF7B':
		return '\u3055'
	case '\uFF7C':
		return '\u3057'
	case '\uFF7D':
		return '\u3059'
	case '\uFF7E':
		return '\u305B'
	case '\uFF7F':
		return '\u305D'
	case '\uFF80':
		return '\u305F'
	case '\uFF81':
		return '\u3061'
	case '\uFF82':
		return '\u3064'
	case '\uFF83':
		return '\u3066'
	case '\uFF84':
		return '\u3068'
	case '\uFF85':
		return '\u306A'
	case '\uFF86':
		return '\u306B'
	case '\uFF87':
		return '\u306C'
	case '\uFF88':
		return '\u306D'
	case '\uFF89':
		return '\u306E'
	case '\uFF8A':
		return '\u306F'
	case '\uFF8B':
		return '\u3072'
	case '\uFF8C':
		return '\u3075'
	case '\uFF8D':
		return '\u3078'
	case '\uFF8E':
		return '\u307B'
	case '\uFF8F':
		return '\u307E'
	case '\uFF90':
		return '\u307F'
	case '\uFF91':
		return '\u3080'
	case '\uFF92':
		return '\u3081'
	case '\uFF93':
		return '\u3082'
	case '\uFF94':
		return '\u3084'
	case '\uFF95':
		return '\u3086'
	case '\uFF96':
		return '\u3088'
	case '\uFF97':
		return '\u3089'
	case '\uFF98':
		return '\u308A'
	case '\uFF99':
		return '\u308B'
	case '\uFF9A':
		return '\u308C'
	case '\uFF9B':
		return '\u308D'
	case '\uFF9C':
		return '\u308F'
	case '\uFF9D':
		return '\u3093'
	case '\uFF9E':
		return '\u309B'
	case '\uFF9F':
		return '\u309C'
	}
	return r
}
