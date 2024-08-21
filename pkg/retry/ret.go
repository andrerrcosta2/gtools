package retry

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

func Of[K any](opt Config[any], fn func() (K, error)) (K, error) {
	var err error
	var res K
	var dl = opt.Dl
	var attErr []string
	if opt.Att == 0 {
		opt.Att = 1
	}

	for i := 0; i < int(opt.Att); i++ {
		select {
		case <-time.After(dl):
			res, err = fn()
			if err == nil {
				return res, nil
			}
			attErr = append(attErr, fmt.Sprintf("attempt %d failed: %s", i, err))

			if opt.Bf > 0.0 {
				dl = time.Duration(float64(dl) * opt.Bf)
			}

		case <-opt.Ctx.Done():
			return res, opt.Ctx.Err()
		}
	}
	return res, fmt.Errorf("sm failed: all %d attempts failed\n\n%s", opt.Att, strings.Join(attErr, "\n"))
}

func Each[T any, K any](opt Config[T], fn func(T) (K, error)) (K, error) {
	var err error
	var res K
	var dl = opt.Dl
	var prms = bprm(&opt)

	for i := 0; i < int(opt.Att); i++ {
		select {
		case <-time.After(dl):
			res, err = fn(prms[i])
			if err == nil {
				return res, nil
			}
			fmt.Printf("attempt %d failed: %s\n", i, err)

			if opt.Bf > 0.0 {
				dl = time.Duration(float64(dl) * opt.Bf)
			}

		case <-opt.Ctx.Done():
			return res, opt.Ctx.Err()
		}
	}
	return res, fmt.Errorf("all %d attempts failed\n", opt.Att)
}

func bprm[T any](opt *Config[T]) []T {
	if opt.Srt != nil {
		if opt.Prm == nil {
			return nil
		}
		sort.Slice(opt.Prm, func(x, y int) bool {
			return opt.Srt(opt.Prm[x], opt.Prm[y])
		})
	}
	return opt.Prm
}
