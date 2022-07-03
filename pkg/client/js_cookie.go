package client

import (
	"bytes"
	"context"
	"net/url"
	"text/template"

	"github.com/PuerkitoBio/goquery"
)

var jsCookieTemplates = template.New("").Funcs(template.FuncMap{
	"backtickQuote": func(s string) string {
		return "`" + s + "`"
	},
})

func init() {
	jsCookieTemplates = template.Must(jsCookieTemplates.New("env").Parse(`
{{- /* */ -}}
const { window, document } = function () {
	const windowPossibleProps = ["close","stop","focus","blur","open","alert","confirm","prompt","print","postMessage","captureEvents","releaseEvents","getSelection","getComputedStyle","matchMedia","moveTo","moveBy","resizeTo","resizeBy","scroll","scrollTo","scrollBy","getDefaultComputedStyle","scrollByLines","scrollByPages","sizeToContent","updateCommands","find","dump","setResizable","requestIdleCallback","cancelIdleCallback","requestAnimationFrame","cancelAnimationFrame","reportError","btoa","atob","setTimeout","clearTimeout","setInterval","clearInterval","queueMicrotask","createImageBitmap","structuredClone","fetch","self","name","history","customElements","locationbar","menubar","personalbar","scrollbars","statusbar","toolbar","status","closed","event","frames","length","opener","parent","frameElement","navigator","clientInformation","external","applicationCache","screen","innerWidth","innerHeight","scrollX","pageXOffset","scrollY","pageYOffset","screenLeft","screenTop","screenX","screenY","outerWidth","outerHeight","performance","mozInnerScreenX","mozInnerScreenY","devicePixelRatio","scrollMaxX","scrollMaxY","fullScreen","ondevicemotion","ondeviceorientation","onabsolutedeviceorientation","InstallTrigger","visualViewport","crypto","onabort","onblur","onfocus","onauxclick","onbeforeinput","oncanplay","oncanplaythrough","onchange","onclick","onclose","oncontextmenu","oncuechange","ondblclick","ondrag","ondragend","ondragenter","ondragexit","ondragleave","ondragover","ondragstart","ondrop","ondurationchange","onemptied","onended","onformdata","oninput","oninvalid","onkeydown","onkeypress","onkeyup","onload","onloadeddata","onloadedmetadata","onloadend","onloadstart","onmousedown","onmouseenter","onmouseleave","onmousemove","onmouseout","onmouseover","onmouseup","onwheel","onpause","onplay","onplaying","onprogress","onratechange","onreset","onresize","onscroll","onsecuritypolicyviolation","onseeked","onseeking","onselect","onslotchange","onstalled","onsubmit","onsuspend","ontimeupdate","onvolumechange","onwaiting","onselectstart","onselectionchange","ontoggle","onpointercancel","onpointerdown","onpointerup","onpointermove","onpointerout","onpointerover","onpointerenter","onpointerleave","ongotpointercapture","onlostpointercapture","onmozfullscreenchange","onmozfullscreenerror","onanimationcancel","onanimationend","onanimationiteration","onanimationstart","ontransitioncancel","ontransitionend","ontransitionrun","ontransitionstart","onwebkitanimationend","onwebkitanimationiteration","onwebkitanimationstart","onwebkittransitionend","u2f","onerror","speechSynthesis","onafterprint","onbeforeprint","onbeforeunload","onhashchange","onlanguagechange","onmessage","onmessageerror","onoffline","ononline","onpagehide","onpageshow","onpopstate","onrejectionhandled","onstorage","onunhandledrejection","onunload","ongamepadconnected","ongamepaddisconnected","localStorage","origin","crossOriginIsolated","isSecureContext","indexedDB","caches","sessionStorage","window","document","location","top"];
	const proxy = (obj, possibleProps) => new Proxy(obj, {
		get(obj, prop) {
			if (prop in obj){
				return obj[prop];
			}
			if (typeof prop === "symbol") {
				return
			}
			if (possibleProps != null && !possibleProps.includes(prop)) {
				return
			}
			throw new Error({{ "get obj(${Object.keys(obj)}).${prop}" | backtickQuote }});
		},
	});

	const div = proxy({
		getElementsByTagName(name) {
			if (name === "i") {
				return proxy([undefined]);
			}
			throw ["div.getElementsByTagName", ...arguments, this];
		},
	});
	const document = proxy({
		head: proxy({
			children: [],
			appendChild(el) {
				this.children.push(el);
			},
		}),
		createElement(tag) {
			if (tag === "script") {
				return proxy({
					readyState: undefined,
				});
			};
			if (tag === "div") {
				return div;
			};
			throw ["document.createElement", ...arguments];
		},
		getElementsByTagName(name) {
			if (name === 'head') {
				return proxy([this.head]);
			}
			throw new Error('document.getElementsByTagName(' + name + ')');
		},
		characterSet: "UTF-8",
		getElementById(id) {
			if (id === "__anchor__") {
				return proxy({});
			}
			throw ["getElementById", ...arguments, this];
		},
	});
	const window = proxy({
		eval,
		escape,
		Number,
		decodeURIComponent,
		isFinite: Number.isFinite,
		JSON,
		document,
		DOMParser: proxy({}),
		RegExp,
		location: proxy({
			protocol: "{{.URL.Scheme}}:",
			host: "{{.URL.Host}}",
			href: "{{.URL.String}}",
			port: "",
		}),
		setTimeout(cb, d) {
			if (d === 0) {
				cb()
				return;
			}
			throw ["setTimeout", ...arguments];
		},
		setInterval() {
			throw ["setInterval", ...arguments];
		},
		XMLHttpRequest() {
			throw arguments;
		},
		top: undefined,
		addEventListener(name, cb) {
			if (name === "load") {
				this.onload = cb;
				return
			}
			if (name === "unload") {
				this.onunload = cb;
				return
			}
			throw arguments;
		},
	}, windowPossibleProps);

	window.top = window;
	return {
		window,
		document,
	};
}();
`))

	jsCookieTemplates = template.Must(jsCookieTemplates.New("src").Parse(`
{{- /* */ -}}
{{ template "env" . }}

{{ .Code }};

document.getElementsByTagName('head')[0].children[0].src;
`))

	jsCookieTemplates = template.Must(jsCookieTemplates.New("TODO").Parse(`
{{- /* */ -}}
{{ template "env" . }}

{{ .Script1 }};
{{ .Script2 }};

window.onload();
[document.head, document, window];
`))

	jsCookieTemplates = template.Must(jsCookieTemplates.New("TODO2").Parse(`
{{- /* */ -}}
{{ template "env" . }}



[document, window];
`))
}

func jsCookieSrc(ctx context.Context, rawURL string, doc goquery.Document) (script, src string, err error) {
	var scriptEl = doc.Find("script#_rspj")
	if scriptEl.Length() == 0 {
		return
	}
	script = scriptEl.Text()
	var jsEngine = ContextJavaScriptEngine(ctx)
	var b bytes.Buffer
	u, err := url.Parse(rawURL)
	if err != nil {
		return
	}
	err = jsCookieTemplates.Lookup("src").Execute(&b, struct {
		URL  *url.URL
		Code string
	}{u, script})
	if err != nil {
		return
	}
	src, err = jsEngine.Run(ctx, b.String())
	return
}

func jsCookieTODO(ctx context.Context, rawURL string, script1, script2 string) (script3 string, err error) {
	var jsEngine = ContextJavaScriptEngine(ctx)
	var b bytes.Buffer
	u, err := url.Parse(rawURL)
	if err != nil {
		return
	}
	err = jsCookieTemplates.Lookup("TODO").Execute(&b, struct {
		URL     *url.URL
		Script1 string
		Script2 string
	}{u, script1, script2})
	if err != nil {
		return
	}
	return jsEngine.Run(ctx, b.String())
}

func jsCookieTODO2(ctx context.Context, rawURL string, script1, script2, script3 string) (_ string, err error) {
	var jsEngine = ContextJavaScriptEngine(ctx)
	var b bytes.Buffer
	u, err := url.Parse(rawURL)
	if err != nil {
		return
	}
	err = jsCookieTemplates.Lookup("TODO2").Execute(&b, struct {
		URL     *url.URL
		Script1 string
		Script2 string
		Script3 string
	}{u, script1, script2, script3})
	if err != nil {
		return
	}
	return jsEngine.Run(ctx, b.String())
}
