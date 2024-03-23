/*
Copyright © 2024 JOSEPH INNES <avianpneuma@gmail.com>
*/
package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/intob/logd/store"
	"golang.org/x/time/rate"
)

type App struct {
	// cfg
	ctx                      context.Context
	port                     int
	logStore                 *store.Store
	rateLimitEvery           time.Duration
	rateLimitBurst           int
	accessControlAllowOrigin string
	// state
	commit     string
	started    time.Time
	clientMu   sync.Mutex
	clients    map[string]*client
	statusJson []byte
}

type Cfg struct {
	Ctx                      context.Context
	Port                     int
	LogStore                 *store.Store
	RateLimitEvery           time.Duration
	RateLimitBurst           int
	AccessControlAllowOrigin string
}

type client struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

func NewApp(cfg *Cfg) *App {
	// read commit file
	commit, err := os.ReadFile("commit")
	if err != nil {
		panic(fmt.Sprintf("failed to read commit file: %s", err))
	}
	app := &App{
		// cfg
		ctx:                      cfg.Ctx,
		port:                     cfg.Port,
		logStore:                 cfg.LogStore,
		rateLimitEvery:           cfg.RateLimitEvery,
		rateLimitBurst:           cfg.RateLimitBurst,
		accessControlAllowOrigin: cfg.AccessControlAllowOrigin,
		// state
		started: time.Now(),
		commit:  string(commit),
		clients: make(map[string]*client),
	}
	go app.cleanupClients()
	go app.measureStatus()
	go app.serve()
	return app
}

func (app *App) serve() {
	mux := http.NewServeMux()
	mux.Handle("/", app.rateLimitMiddleware(
		app.corsMiddleware(
			http.HandlerFunc(app.handleRequest))))
	laddr := fmt.Sprintf(":%d", app.port)
	fmt.Println("app listening http on", laddr)
	server := &http.Server{Addr: laddr, Handler: mux}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(fmt.Sprintf("failed to listen http: %v\n", err))
		}
	}()
	<-app.ctx.Done()
	app.shutdown(server)
}

func (app *App) handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	app.handleStatus(w)
}

func (app *App) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", app.accessControlAllowOrigin)
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Fly-Client-IP")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		next.ServeHTTP(w, r)
	})
}

// rateLimitMiddleware is a middleware that limits the rate of requests.
func (app *App) rateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		limiter := app.getRateLimiter(r)
		if !limiter.Allow() &&
			// TODO: REMOVE. TEST ONLY! Disables rate limit for POST
			r.Method != "POST" {
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// getRateLimiter returns a rate limiter for the given request.
func (app *App) getRateLimiter(r *http.Request) *rate.Limiter {
	addr := r.Header.Get("Fly-Client-IP")
	if addr == "" {
		addr = r.RemoteAddr // fallback when local
	}
	app.clientMu.Lock()
	defer app.clientMu.Unlock()
	key := r.Method + addr
	v, exists := app.clients[key]
	if !exists {
		limiter := rate.NewLimiter(rate.Every(app.rateLimitEvery), app.rateLimitBurst)
		app.clients[key] = &client{limiter, time.Now()}
		return limiter
	}
	v.lastSeen = time.Now()
	return v.limiter
}

// cleanupClients removes clients that have not been seen for 10 seconds.
func (a *App) cleanupClients() {
	for {
		select {
		case <-a.ctx.Done():
			return
		case <-time.After(10 * time.Second):
			a.clientMu.Lock()
			for key, client := range a.clients {
				if time.Since(client.lastSeen) > 10*time.Second {
					delete(a.clients, key)
				}
			}
			a.clientMu.Unlock()
		}
	}
}

// shutdown attempts to gracefully shutdown the server.
func (a *App) shutdown(server *http.Server) {
	// Create a context with timeout for the server shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// Attempt to gracefully shutdown the server
	if err := server.Shutdown(ctx); err != nil {
		panic(fmt.Sprintf("server shutdown failed: %v", err))
	}
	fmt.Println("server shutdown gracefully")
}
