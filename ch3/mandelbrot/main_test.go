package main

import (
	"math"
	"math/big"
	"math/rand"
	"testing"
)

func TestComplexMultiplication(t *testing.T) {
	a := Complex{big.NewFloat(1), big.NewFloat(2)}
	b := Complex{big.NewFloat(3), big.NewFloat(4)}

	expected := Complex{big.NewFloat(-5), big.NewFloat(10)}
	actual := complexMulBig(a, b)

	if actual.Real.Cmp(expected.Real) != 0 {
		t.Errorf("expected %v, got %v", expected.Real, actual.Real)
	}

	if actual.Imag.Cmp(expected.Imag) != 0 {
		t.Errorf("expected %v, got %v", expected.Imag, actual.Imag)
	}

	regularComplex1 := complex(rand.Float64(), rand.Float64())
	regularComplex2 := complex(rand.Float64(), rand.Float64())

	a.Real = big.NewFloat(real(regularComplex1))
	a.Imag = big.NewFloat(imag(regularComplex1))

	b.Real = big.NewFloat(real(regularComplex2))
	b.Imag = big.NewFloat(imag(regularComplex2))

	expectedComplex := regularComplex1 * regularComplex2

	expected = Complex{big.NewFloat(real(expectedComplex)), big.NewFloat(imag(expectedComplex))}
	actual = complexMulBig(a, b)

	if actual.Real.Cmp(expected.Real) != 0 {
		t.Errorf("expected %v, got %v", expected.Real, actual.Real)
	}

	if actual.Imag.Cmp(expected.Imag) != 0 {
		t.Errorf("expected %v, got %v", expected.Imag, actual.Imag)
	}
}

func TestComplexAbs(t *testing.T) {
	a := Complex{big.NewFloat(1), big.NewFloat(2)}

	expected := big.NewFloat(math.Sqrt(5))
	actual := complexAbsBig(a)

	if actual.Cmp(expected) != 0 {
		t.Errorf("expected %v, got %v", expected, actual)
	}

	if actual.Cmp(expected) != 0 {
		t.Errorf("expected %v, got %v", expected, actual)
	}

	regularComplex1 := complex(rand.Float64(), rand.Float64())

	a.Real = big.NewFloat(real(regularComplex1))
	a.Imag = big.NewFloat(imag(regularComplex1))

	real := real(regularComplex1)
	imag := imag(regularComplex1)

	expected = big.NewFloat(math.Sqrt(real*real + imag*imag))
	actual = complexAbsBig(a)

	if actual.Cmp(expected) != 0 {
		t.Errorf("expected %v, got %v", expected, actual)
	}

	if actual.Cmp(expected) != 0 {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}
