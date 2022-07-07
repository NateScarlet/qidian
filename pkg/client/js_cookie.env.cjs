// @ts-check
/// <reference no-default-lib="true" />
/// <reference types="./js_cookie.env"/>

const { window, document } = (function () {
  const windowPossibleProps = [
    "close",
    "stop",
    "focus",
    "blur",
    "open",
    "alert",
    "confirm",
    "prompt",
    "print",
    "postMessage",
    "captureEvents",
    "releaseEvents",
    "getSelection",
    "getComputedStyle",
    "matchMedia",
    "moveTo",
    "moveBy",
    "resizeTo",
    "resizeBy",
    "scroll",
    "scrollTo",
    "scrollBy",
    "getDefaultComputedStyle",
    "scrollByLines",
    "scrollByPages",
    "sizeToContent",
    "updateCommands",
    "find",
    "dump",
    "setResizable",
    "requestIdleCallback",
    "cancelIdleCallback",
    "requestAnimationFrame",
    "cancelAnimationFrame",
    "reportError",
    "btoa",
    "atob",
    "setTimeout",
    "clearTimeout",
    "setInterval",
    "clearInterval",
    "queueMicrotask",
    "createImageBitmap",
    "structuredClone",
    "fetch",
    "self",
    "name",
    "history",
    "customElements",
    "locationbar",
    "menubar",
    "personalbar",
    "scrollbars",
    "statusbar",
    "toolbar",
    "status",
    "closed",
    "event",
    "frames",
    "length",
    "opener",
    "parent",
    "frameElement",
    "navigator",
    "clientInformation",
    "external",
    "applicationCache",
    "screen",
    "innerWidth",
    "innerHeight",
    "scrollX",
    "pageXOffset",
    "scrollY",
    "pageYOffset",
    "screenLeft",
    "screenTop",
    "screenX",
    "screenY",
    "outerWidth",
    "outerHeight",
    "performance",
    "mozInnerScreenX",
    "mozInnerScreenY",
    "devicePixelRatio",
    "scrollMaxX",
    "scrollMaxY",
    "fullScreen",
    "ondevicemotion",
    "ondeviceorientation",
    "onabsolutedeviceorientation",
    "InstallTrigger",
    "visualViewport",
    "crypto",
    "onabort",
    "onblur",
    "onfocus",
    "onauxclick",
    "onbeforeinput",
    "oncanplay",
    "oncanplaythrough",
    "onchange",
    "onclick",
    "onclose",
    "oncontextmenu",
    "oncuechange",
    "ondblclick",
    "ondrag",
    "ondragend",
    "ondragenter",
    "ondragexit",
    "ondragleave",
    "ondragover",
    "ondragstart",
    "ondrop",
    "ondurationchange",
    "onemptied",
    "onended",
    "onformdata",
    "oninput",
    "oninvalid",
    "onkeydown",
    "onkeypress",
    "onkeyup",
    "onload",
    "onloadeddata",
    "onloadedmetadata",
    "onloadend",
    "onloadstart",
    "onmousedown",
    "onmouseenter",
    "onmouseleave",
    "onmousemove",
    "onmouseout",
    "onmouseover",
    "onmouseup",
    "onwheel",
    "onpause",
    "onplay",
    "onplaying",
    "onprogress",
    "onratechange",
    "onreset",
    "onresize",
    "onscroll",
    "onsecuritypolicyviolation",
    "onseeked",
    "onseeking",
    "onselect",
    "onslotchange",
    "onstalled",
    "onsubmit",
    "onsuspend",
    "ontimeupdate",
    "onvolumechange",
    "onwaiting",
    "onselectstart",
    "onselectionchange",
    "ontoggle",
    "onpointercancel",
    "onpointerdown",
    "onpointerup",
    "onpointermove",
    "onpointerout",
    "onpointerover",
    "onpointerenter",
    "onpointerleave",
    "ongotpointercapture",
    "onlostpointercapture",
    "onmozfullscreenchange",
    "onmozfullscreenerror",
    "onanimationcancel",
    "onanimationend",
    "onanimationiteration",
    "onanimationstart",
    "ontransitioncancel",
    "ontransitionend",
    "ontransitionrun",
    "ontransitionstart",
    "onwebkitanimationend",
    "onwebkitanimationiteration",
    "onwebkitanimationstart",
    "onwebkittransitionend",
    "u2f",
    "onerror",
    "speechSynthesis",
    "onafterprint",
    "onbeforeprint",
    "onbeforeunload",
    "onhashchange",
    "onlanguagechange",
    "onmessage",
    "onmessageerror",
    "onoffline",
    "ononline",
    "onpagehide",
    "onpageshow",
    "onpopstate",
    "onrejectionhandled",
    "onstorage",
    "onunhandledrejection",
    "onunload",
    "ongamepadconnected",
    "ongamepaddisconnected",
    "localStorage",
    "origin",
    "crossOriginIsolated",
    "isSecureContext",
    "indexedDB",
    "caches",
    "sessionStorage",
    "window",
    "document",
    "location",
    "top",
  ];

  /**
   *
   * @template { {} } T
   * @param {T} obj
   * @param {string[]=} possibleProps
   * @returns {T}
   */
  const proxyGet = (obj, possibleProps) =>
    new Proxy(obj, {
      get(obj, prop) {
        if (prop in obj) {
          return obj[prop];
        }
        if (
          typeof prop === "string" &&
          possibleProps &&
          !possibleProps.includes(prop)
        ) {
          return;
        }
        throw new Error(
          `unexpected get obj(${Object.keys(obj)}).${String(prop)}`
        );
      },
    });

  /**
   *
   * @template { {} } T
   * @param {T} obj
   * @param {string[]=} possibleProps
   * @returns {T}
   */
  const proxySet = (obj, possibleProps) =>
    new Proxy(obj, {
      set(obj, prop, value) {
        if (prop in obj) {
          obj[prop] = value;
          return true;
        }
        if (
          typeof prop === "string" &&
          possibleProps != null &&
          possibleProps.includes(prop)
        ) {
          obj[prop] = value;
          return true;
        }
        throw new Error(
          `unexpected set obj(${Object.keys(obj)}).${String(prop)} = ${value}`
        );
      },
    });

  const div = proxyGet({
    getElementsByTagName(name) {
      if (name === "i") {
        return proxyGet([undefined]);
      }
      throw ["div.getElementsByTagName", ...arguments, this];
    },
  });
  const document = proxyGet({
    head: proxyGet({
      /**
       * @type {unknown[]}
       */
      children: [],
      /**
       *
       * @param {unknown} el
       */
      appendChild(el) {
        this.children.push(el);
      },
    }),
    createElement(tag) {
      if (tag === "script") {
        return proxyGet({
          readyState: undefined,
        });
      }
      if (tag === "div") {
        return div;
      }
      throw ["document.createElement", ...arguments];
    },
    getElementsByTagName(name) {
      if (name === "head") {
        return proxyGet([this.head]);
      }
      throw new Error("document.getElementsByTagName(" + name + ")");
    },
    characterSet: "UTF-8",
    getElementById(id) {
      if (id === "__anchor__") {
        return proxyGet({});
      }
      throw ["getElementById", ...arguments, this];
    },
  });
  const onload = [() => {}];
  const onunload = [() => {}];
  const window = proxySet(
    new Proxy(
      proxyGet(
        {
          eval,
          escape,
          Number,
          /** @type {unknown} */
          get top() {
            return window;
          },
          decodeURIComponent,
          isFinite: Number.isFinite,
          JSON,
          document,
          DOMParser: proxyGet({}),
          RegExp,
          location: proxyGet({
            protocol: "{{.URL.Scheme}}:",
            host: "{{.URL.Host}}",
            href: "{{.URL.String}}",
            port: "",
          }),
          setTimeout(cb, d) {
            if (d === 0) {
              cb();
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
          onload: () => {},
          onunload: () => {},
          addEventListener(name, cb) {
            if (name === "load") {
              onload.push(cb);
              return;
            }
            if (name === "unload") {
              onunload.push(cb);
              return;
            }
            throw arguments;
          },
          /**
           *
           * @param { {type: string} } e
           */
          dispatchEvent(e) {
            if (e.type === "load") {
              this.onload();
              onload.forEach((i) => i());
              return;
            }
            if (e.type === "unload") {
              this.onunload();
              onunload.forEach((i) => i());
              return;
            }
            throw arguments;
          },
          $_ts: undefined,
          "2eab37c0ead4b0b0": undefined,
          $b_onBridgeReady: undefined,
          $b_setup: undefined,
        },

        windowPossibleProps
      ),
      {
        set(obj, prop, value) {
          obj[prop] = value;
          globalThis[prop] = value;
          return true;
        },
      }
    )
  );

  return {
    window,
    document,
  };
})();
