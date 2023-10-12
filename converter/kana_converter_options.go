package converter

import (
	"fmt"
)

type KanaConverterOptions struct {
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

func (r *KanaConverterOptions) EnableOptr() error {
	if r.optR {
		return fmt.Errorf("must not combine 'r' and 'R' flags")
	}
	if r.optA {
		return fmt.Errorf("must not combine 'r' and 'A' flags")
	}
	r.optr = true
	return nil
}

func (r *KanaConverterOptions) EnableOptR() error {
	if r.optr {
		return fmt.Errorf("must not combine 'r' and 'R' flags")
	}
	if r.opta {
		return fmt.Errorf("must not combine 'R' and 'a' flags")
	}
	r.optR = true
	return nil
}

func (r *KanaConverterOptions) EnableOptn() error {
	if r.optN {
		return fmt.Errorf("must not combine 'n' and 'N' flags")
	}
	if r.optA {
		return fmt.Errorf("must not combine 'n' and 'A' flags")
	}
	r.optn = true
	return nil
}

func (r *KanaConverterOptions) EnableOptN() error {
	if r.optn {
		return fmt.Errorf("must not combine 'n' and 'N' flags")
	}
	if r.opta {
		return fmt.Errorf("must not combine 'N' and 'a' flags")
	}
	r.optN = true
	return nil
}

func (r *KanaConverterOptions) EnableOpta() error {
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

func (r *KanaConverterOptions) EnableOptA() error {
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

func (r *KanaConverterOptions) EnableOpts() error {
	if r.optS {
		return fmt.Errorf("must not combine 's' and 'S' flags")
	}
	r.opts = true
	return nil
}

func (r *KanaConverterOptions) EnableOptS() error {
	if r.opts {
		return fmt.Errorf("must not combine 's' and 'S' flags")
	}
	r.optS = true
	return nil
}

func (r *KanaConverterOptions) EnableOptk() error {
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

func (r *KanaConverterOptions) EnableOptK() error {
	if r.optk {
		return fmt.Errorf("must not combine 'k' and 'K' flags")
	}
	if r.optH {
		return fmt.Errorf("must not combine 'K' and 'H' flags")
	}

	r.optK = true
	return nil
}

func (r *KanaConverterOptions) EnableOpth() error {
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

func (r *KanaConverterOptions) EnableOptH() error {
	if r.opth {
		return fmt.Errorf("must not combine 'h' and 'H' flags")
	}
	if r.optK {
		return fmt.Errorf("must not combine 'K' and 'H' flags")
	}

	r.optH = true
	return nil
}

func (r *KanaConverterOptions) EnableOptc() error {
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

func (r *KanaConverterOptions) EnableOptC() error {
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

func (r *KanaConverterOptions) EnableOptV() error {
	r.optV = true
	return nil
}

func NewKanaConverterOptions(mode string) (*KanaConverterOptions, error) {
	o := &KanaConverterOptions{}
	for _, char := range mode {
		var err error
		switch char {
		case rune('r'):
			err = o.EnableOptr()
		case rune('R'):
			err = o.EnableOptR()
		case rune('n'):
			err = o.EnableOptn()
		case rune('N'):
			err = o.EnableOptN()
		case rune('a'):
			err = o.EnableOpta()
		case rune('A'):
			err = o.EnableOptA()
		case rune('s'):
			err = o.EnableOpts()
		case rune('S'):
			err = o.EnableOptS()
		case rune('k'):
			err = o.EnableOptk()
		case rune('K'):
			err = o.EnableOptK()
		case rune('h'):
			err = o.EnableOpth()
		case rune('H'):
			err = o.EnableOptH()
		case rune('c'):
			err = o.EnableOptc()
		case rune('C'):
			err = o.EnableOptC()
		case rune('V'):
			err = o.EnableOptV()
		}
		if err != nil {
			return nil, err
		}
	}

	return o, nil
}
