package middleware

import "net/http"

type Middleware func(http.HandlerFunc) http.HandlerFunc

func Handler(hd http.HandlerFunc, mid ...Middleware) http.HandlerFunc {
	
	for i := len(mid); i > 0; i-- {
		hd = mid[i-1](hd)
	}

	return hd
}