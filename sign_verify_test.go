package bench_ed25519

import (
	"crypto/ed25519"
	"crypto/rand"
	"testing"
)

func Benchmark_Sign(b *testing.B) {
	_, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		b.Fatalf("Generation error : %s", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = priv.Sign(nil, body, &ed25519.Options{})
	}
}

func Benchmark_Verify(b *testing.B) {
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		b.Fatalf("Generation error : %s", err)
	}

	signed, err := priv.Sign(nil, body, &ed25519.Options{})
	if err != nil {
		b.Fatalf("Sign error: %s", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ed25519.VerifyWithOptions(pub, body, signed, &ed25519.Options{})
	}
}

var body = []byte(`
{
    "iss": "Some big string for testing purpose only",
    "iat": 1504699136,
    "sub": "github|353454354354353453",
    "exp": 1504699256,
    "iss1": "Some big string for testing purpose only",
    "iat1": 1504699136,
    "sub1": "github|353454354354353453",
    "exp1": 1504699256,
    "iss2": "Some big string for testing purpose only",
    "iat2": 1504699136,
    "sub2": "github|353454354354353453",
    "exp2": 1504699256,
    "iss3": "Some big string for testing purpose only",
    "iat3": 1504699136,
    "sub3": "github|353454354354353453",
    "exp3": 1504699256,
    "iss4": "Some big string for testing purpose only",
    "iat4": 1504699136,
    "sub4": "github|353454354354353453",
    "exp4": 1504699256,
    "iss5": "Some big string for testing purpose only",
    "iat5": 1504699136,
    "sub5": "github|353454354354353453",
    "exp5": 1504699256,
    "iss6: "Some big string for testing purpose only",
    "iat6": 1504699136,
    "sub6": "github|353454354354353453",
    "exp6": 1504699256,
    "iss7": "Some big string for testing purpose only",
    "iat7": 1504699136,
    "sub7": "github|353454354354353453",
    "exp7": 1504699256
}`)
