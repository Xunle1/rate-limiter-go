package token_bucket

import (
	"fmt"
	"log"
	"time"
)

type Token struct {
	Id uint
}

type TokenBucket struct {
	Tokens   []*Token
	Remain   uint
	Total    uint
	Duration time.Duration
}

func NewTokenBucket(total uint, duration time.Duration) *TokenBucket {
	tokens := make([]*Token, 0)
	for i := 0; i < int(total); i++ {
		tokens = append(tokens, &Token{Id: uint(i)})
	}

	tb := &TokenBucket{
		Tokens:   tokens,
		Remain:   uint(total),
		Total:    uint(total),
		Duration: duration,
	}

	tb.startRefiller()

	return tb
}

func DefaultTokenBucket() *TokenBucket {
	return NewTokenBucket(4, 10*time.Second)
}

func (t *TokenBucket) Get() error {
	if t.Remain <= 0 {
		return fmt.Errorf("no token remain")
	}

	token := t.Tokens[t.Remain-1]
	log.Printf("Get token: %d.", token.Id)

	t.Remain--
	log.Printf("%d token(s) remain.", t.Remain)
	return nil
}

func (t *TokenBucket) startRefiller() {
	go func() {
		ticker := time.NewTicker(t.Duration)
		for range ticker.C {
			t.refiller()
		}
	}()
}

func (t *TokenBucket) refiller() {
	if t.Remain >= t.Total {
		log.Println("Token overflowed.")
		return
	}
	if t.Remain < t.Total {
		t.Tokens = append(t.Tokens, &Token{Id: t.Remain})
		t.Remain++
		log.Printf("Refill 1 token, %d remaining", t.Remain)
	}
}
