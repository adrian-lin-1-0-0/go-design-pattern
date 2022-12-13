package proxy

type Application struct {
}

func (a *Application) HandleRequest(url, method string) (int, string) {
	if url == "/app/status" && method == "GET" {
		return 200, "OK"
	}

	if url == "/create/user" && method == "POST" {
		return 201, "USER_CREATED"
	}

	return 404, "NOT_FOUND"
}

type Proxy struct {
	application       *Application
	maxAllowedRequest int
	rateLimiter       map[string]int
}

func NewProxyServer() *Proxy {
	return &Proxy{
		application:       &Application{},
		maxAllowedRequest: 1,
		rateLimiter:       make(map[string]int),
	}
}

func (p *Proxy) checkRateLimiting(url string) bool {
	if p.rateLimiter[url] == 0 {
		p.rateLimiter[url] = 1
	}
	if p.rateLimiter[url] > p.maxAllowedRequest {
		return false
	}
	p.rateLimiter[url] = p.rateLimiter[url] + 1
	return true
}

func (p *Proxy) HandleRequest(url, method string) (int, string) {
	allowed := p.checkRateLimiting(url)
	if !allowed {
		return 403, "Not Allowed"
	}
	return p.application.HandleRequest(url, method)
}
