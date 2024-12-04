package web

// import (
// 	"context"
// 	"net/http"
// 	"sync"
// 	"time"

// 	"github.com/cufee/aftermath/internal/json"

// 	"github.com/cufee/aftermath/cmd/core/server"
// 	"github.com/cufee/aftermath/cmd/core/widget"
// 	"github.com/cufee/aftermath/internal/constants"
// 	"github.com/go-rod/rod"
// 	"github.com/go-rod/rod/lib/launcher"
// 	"github.com/go-rod/rod/lib/proto"
// 	"github.com/go-rod/rod/lib/utils"
// 	"github.com/lucsky/cuid"
// 	"github.com/pkg/errors"
// )

// func NewServer(port string) func() {
// 	renderPage := BuildRenderPage(constants.FrontendURL, constants.FrontendURL+"/assets/js/widget/amth-widget.js", 5)
// 	return server.NewServer(port, []server.Handler{
// 		{
// 			Path: "GET /render",
// 			Func: func(w http.ResponseWriter, r *http.Request) {
// 				w.Write([]byte(renderPage))
// 				w.Header().Add("Content-Type", "text/html; charset=utf-8")
// 			},
// 		},
// 	})
// }

// type rendererOptions struct {
// 	pagePoolSize int
// 	headless     bool
// 	devTools     bool
// }

// type rendererOption func(*rendererOptions)

// func WithPoolSize(size int) func(ro *rendererOptions) {
// 	return func(ro *rendererOptions) {
// 		if size > 0 {
// 			ro.pagePoolSize = size
// 		}
// 	}
// }
// func WithWindow() func(ro *rendererOptions) {
// 	return func(ro *rendererOptions) {
// 		ro.headless = false
// 	}
// }
// func WithDevTools() func(ro *rendererOptions) {
// 	return func(ro *rendererOptions) {
// 		ro.devTools = true
// 	}
// }

// func NewWebRenderer(renderURL string, o ...rendererOption) (*WebRenderer, error) {
// 	opts := rendererOptions{pagePoolSize: 1, headless: true, devTools: false}
// 	for _, apply := range o {
// 		apply(&opts)
// 	}

// 	renderer := &WebRenderer{
// 		renderURL: renderURL,
// 		launcher:  launcher.New().Headless(opts.headless).Devtools(opts.devTools),
// 		pool:      newRendererPool(),
// 	}
// 	url, err := renderer.launcher.Launch()
// 	if err != nil {
// 		return nil, errors.Wrap(err, "failed to launch a new browser")
// 	}

// 	renderer.browser = rod.New().ControlURL(url)
// 	if err := renderer.browser.Connect(); err != nil {
// 		return nil, errors.Wrap(err, "failed to connect controller to a browser")
// 	}

// 	for range opts.pagePoolSize {
// 		page, err := renderer.browser.Page(proto.TargetCreateTarget{URL: renderer.renderURL})
// 		if err != nil {
// 			return nil, errors.Wrap(err, "failed to create a new page")
// 		}

// 		err = renderer.pool.RegisterPage(page)
// 		if err != nil {
// 			return nil, err
// 		}
// 	}

// 	return renderer, nil
// }

// type WebRenderer struct {
// 	launcher  *launcher.Launcher
// 	browser   *rod.Browser
// 	renderURL string

// 	pool *rendererPool
// }

// func (r *WebRenderer) Cleanup() {
// 	r.launcher.Cleanup()
// }

// func (r *WebRenderer) Render(ctx context.Context, data widget.WidgetData, background string) ([]byte, error) {
// 	input, err := json.Marshal(data)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "failed to marshal data")
// 	}
// 	return r.RenderEncoded(ctx, string(input), background)
// }

// func (r *WebRenderer) RenderEncoded(ctx context.Context, data, background string) ([]byte, error) {
// 	node, err := r.pool.GetAndLockNode(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer node.lock.Unlock()
// 	return node.captureNodeImage(ctx, data, background)
// }

// type rendererPool struct {
// 	lock *sync.Mutex

// 	id    string
// 	nodes []*poolNode
// }

// type poolNode struct {
// 	lock     *sync.Mutex
// 	node     *rod.Element
// 	page     *rod.Page
// 	pageLock *sync.Mutex
// }

// func (p *rendererPool) GetAndLockNode(ctx context.Context) (*poolNode, error) {
// 	p.lock.Lock()
// 	defer p.lock.Unlock()

// 	ticker := time.NewTicker(time.Millisecond * 25)
// 	defer ticker.Stop()
// 	for {
// 		select {
// 		case <-ticker.C:
// 			for _, node := range p.nodes {
// 				if !node.lock.TryLock() {
// 					continue
// 				}
// 				return node, nil
// 			}
// 		case <-ctx.Done():
// 			return nil, ctx.Err()
// 		}
// 	}
// }

// func (p *poolNode) captureNodeImage(ctx context.Context, input, background string) ([]byte, error) {
// 	_, err := p.node.Eval("this.setData", input, background)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "failed to set data")
// 	}

// 	ticker := time.NewTicker(time.Millisecond * 10)
// 	defer ticker.Stop()

// outerLoop:
// 	for {
// 		select {
// 		case <-ticker.C:
// 			errValue, _ := p.node.Eval("() => this.shadowRoot.querySelector('.amth-error').dataset.error")
// 			if errValue != nil && errValue.Value.String() != "" {
// 				return nil, errors.New("failed to render cards: " + errValue.Value.String())
// 			}
// 			rendered, _ := p.node.Eval("() => !!this.parentElement.querySelector('input').checked && [...this.shadowRoot.querySelectorAll('img')].filter((img) => img.dataset?.loaded !== '1').length === 0")
// 			if rendered != nil && rendered.Value.Bool() {
// 				break outerLoop
// 			}
// 		case <-ctx.Done():
// 			return nil, ctx.Err()
// 		}
// 	}

// 	// There are some rare cases where the shape returns an old value, so we wait a little bit extra
// 	p.node.WaitStable(time.Millisecond * 10)
// 	shape, err := p.node.Shape()
// 	if err != nil {
// 		return nil, errors.Wrap(err, "failed to get node shape")
// 	}
// 	p.pageLock.Lock()
// 	image, err := p.page.Screenshot(true, &proto.PageCaptureScreenshot{
// 		Clip: &proto.PageViewport{
// 			X:      shape.Box().X,
// 			Y:      shape.Box().Y,
// 			Width:  shape.Box().Width,
// 			Height: shape.Box().Height,
// 			Scale:  1.5,
// 		},
// 		FromSurface:           false,
// 		CaptureBeyondViewport: true,
// 		OptimizeForSpeed:      true,
// 		Format:                proto.PageCaptureScreenshotFormatPng,
// 	})
// 	p.pageLock.Unlock()
// 	if err != nil {
// 		return nil, errors.Wrap(err, "failed to screenshot a page")
// 	}

// 	_, err = p.node.Eval("this.clear")
// 	if err != nil {
// 		return nil, errors.Wrap(err, "failed clear data after use")
// 	}

// 	return image, nil
// }

// func (p *rendererPool) RegisterPage(page *rod.Page) error {
// 	p.lock.Lock()
// 	defer p.lock.Unlock()

// 	page = page.Sleeper(func() utils.Sleeper {
// 		return utils.BackoffSleeper(time.Second, time.Second*5, func(d time.Duration) time.Duration { return d })
// 	})

// 	result, err := page.Eval("() => !!document.querySelector('#pool-id')")
// 	if err != nil {
// 		return errors.Wrap(err, "failed to find a pool id anchor element")
// 	}
// 	if result.Value.Bool() {
// 		return errors.New("page already registered as part of a pool")
// 	}

// 	_, err = page.Eval(`(id) => {
// 		const input = document.createElement("input")
// 		input.id = "pool-id"
// 		input.hidden = true
// 		input.value = id
// 		document.body.prepend(input)
// 	}`, p.id)
// 	if err != nil {
// 		return errors.Wrap(err, "failed to create a pool id anchor element on the page")
// 	}

// 	widgets, err := page.Elements("amth-render")
// 	if err != nil {
// 		return errors.Wrap(err, "failed to get render element on a page")
// 	}
// 	if len(widgets) < 1 {
// 		return errors.New("page does not have any render elements")

// 	}

// 	pageLock := &sync.Mutex{}
// 	for _, node := range widgets {
// 		p.nodes = append(p.nodes, &poolNode{
// 			lock:     &sync.Mutex{},
// 			pageLock: pageLock,
// 			page:     page,
// 			node:     node,
// 		})
// 	}
// 	return nil
// }

// func newRendererPool() *rendererPool {
// 	return &rendererPool{id: cuid.New(), lock: &sync.Mutex{}}
// }
