package crawler

import "net/http"

//RequestModifierFunc is a function that takes a request and probably modifies it.
type RequestModifierFunc func(req *http.Request)

//RequestModifier A collection of functions to be applied on a request.
type RequestModifier struct {
	funcs []RequestModifierFunc
}

//With registers a request modifier function to be applied later with Do()
func (r *RequestModifier) With(modifierFunc RequestModifierFunc) {
	r.funcs = append(r.funcs, modifierFunc)
}

//Do executes all registered request modifier functions on a request.
func (r *RequestModifier) Do(req *http.Request) {
	for _, f := range r.funcs {
		f(req)
	}
}

//CreateRequestsFromUrls creates HTTP request objects from a list of given urls.
func CreateRequestsFromUrls(urls []string, modifiers RequestModifier) ([]*http.Request, error) {
	var requests []*http.Request

	for _, url := range urls {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}

		modifiers.Do(req)
		requests = append(requests, req)
	}

	return requests, nil
}
