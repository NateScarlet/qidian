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

  const documentPossibleProps = [
    "location",
    "getElementsByTagName",
    "getElementsByTagNameNS",
    "getElementsByClassName",
    "getElementById",
    "createElement",
    "createElementNS",
    "createDocumentFragment",
    "createTextNode",
    "createComment",
    "createProcessingInstruction",
    "importNode",
    "adoptNode",
    "createEvent",
    "createRange",
    "createNodeIterator",
    "createTreeWalker",
    "createCDATASection",
    "createAttribute",
    "createAttributeNS",
    "getElementsByName",
    "open",
    "close",
    "write",
    "writeln",
    "hasFocus",
    "execCommand",
    "queryCommandEnabled",
    "queryCommandIndeterm",
    "queryCommandState",
    "queryCommandSupported",
    "queryCommandValue",
    "releaseCapture",
    "mozSetImageElement",
    "clear",
    "captureEvents",
    "releaseEvents",
    "exitFullscreen",
    "mozCancelFullScreen",
    "exitPointerLock",
    "enableStyleSheetsForSet",
    "caretPositionFromPoint",
    "querySelector",
    "querySelectorAll",
    "getSelection",
    "hasStorageAccess",
    "requestStorageAccess",
    "elementFromPoint",
    "elementsFromPoint",
    "getAnimations",
    "prepend",
    "append",
    "replaceChildren",
    "createExpression",
    "createNSResolver",
    "evaluate",
    "implementation",
    "URL",
    "documentURI",
    "compatMode",
    "characterSet",
    "charset",
    "inputEncoding",
    "contentType",
    "doctype",
    "documentElement",
    "domain",
    "referrer",
    "cookie",
    "lastModified",
    "readyState",
    "title",
    "dir",
    "body",
    "head",
    "images",
    "embeds",
    "plugins",
    "links",
    "forms",
    "scripts",
    "defaultView",
    "designMode",
    "onreadystatechange",
    "onbeforescriptexecute",
    "onafterscriptexecute",
    "currentScript",
    "fgColor",
    "linkColor",
    "vlinkColor",
    "alinkColor",
    "bgColor",
    "anchors",
    "applets",
    "all",
    "fullscreen",
    "mozFullScreen",
    "fullscreenEnabled",
    "mozFullScreenEnabled",
    "onfullscreenchange",
    "onfullscreenerror",
    "onpointerlockchange",
    "onpointerlockerror",
    "hidden",
    "visibilityState",
    "onvisibilitychange",
    "selectedStyleSheetSet",
    "lastStyleSheetSet",
    "preferredStyleSheetSet",
    "styleSheetSets",
    "scrollingElement",
    "timeline",
    "rootElement",
    "oncopy",
    "oncut",
    "onpaste",
    "activeElement",
    "styleSheets",
    "pointerLockElement",
    "fullscreenElement",
    "mozFullScreenElement",
    "adoptedStyleSheets",
    "fonts",
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
    "onerror",
    "children",
    "firstElementChild",
    "lastElementChild",
    "childElementCount",
    "getRootNode",
    "hasChildNodes",
    "insertBefore",
    "appendChild",
    "replaceChild",
    "removeChild",
    "normalize",
    "cloneNode",
    "isSameNode",
    "isEqualNode",
    "compareDocumentPosition",
    "contains",
    "lookupPrefix",
    "lookupNamespaceURI",
    "isDefaultNamespace",
    "nodeType",
    "nodeName",
    "baseURI",
    "isConnected",
    "ownerDocument",
    "parentNode",
    "parentElement",
    "childNodes",
    "firstChild",
    "lastChild",
    "previousSibling",
    "nextSibling",
    "nodeValue",
    "textContent",
    "ELEMENT_NODE",
    "ATTRIBUTE_NODE",
    "TEXT_NODE",
    "CDATA_SECTION_NODE",
    "ENTITY_REFERENCE_NODE",
    "ENTITY_NODE",
    "PROCESSING_INSTRUCTION_NODE",
    "COMMENT_NODE",
    "DOCUMENT_NODE",
    "DOCUMENT_TYPE_NODE",
    "DOCUMENT_FRAGMENT_NODE",
    "NOTATION_NODE",
    "DOCUMENT_POSITION_DISCONNECTED",
    "DOCUMENT_POSITION_PRECEDING",
    "DOCUMENT_POSITION_FOLLOWING",
    "DOCUMENT_POSITION_CONTAINS",
    "DOCUMENT_POSITION_CONTAINED_BY",
    "DOCUMENT_POSITION_IMPLEMENTATION_SPECIFIC",
    "addEventListener",
    "removeEventListener",
    "dispatchEvent",
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
  const document = proxyGet(
    {
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
        removeChild(el) {
          if (el.id === "_rspj") {
            return;
          }
          throw ["head.removeChild", ...arguments, this.children];
        },
      }),
      createElement(tag) {
        if (tag === "script") {
          return proxyGet({
            tagName: "SCRIPT",
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
        if (name === "base") {
          return [];
        }
        if (name === "script") {
          return proxyGet([
            proxyGet({
              id: "_rspj",
              getAttribute(name) {
                if (name === "r") {
                  return "m";
                }
                throw ["script.0.getAttribute", ...arguments];
              },
              get parentElement() {
                return document.head;
              },
            }),
          ]);
        }
        throw new Error("document.getElementsByTagName(" + name + ")");
      },
      characterSet: "UTF-8",
      getElementById(id) {
        if (id === "__anchor__") {
          return proxyGet({
            id,
            addEventListener(name) {
              if (name === "dblclick") {
                return;
              }
              throw arguments;
            },
          });
        }
        if (id === "__onload__") {
          return proxyGet({
            id,
            name: "ehYvWYh7dIVSRWtku20i.7x35mAh2xs8SGcig01aqZs77n5XB8e2L4paz7CyFe1DAu4NdWvxkftV1LDB5jVm4pxGBxW316q3EyPkoTy8Pck8D.Dy1xQHiKtULpFosDPx",
            value: "JopNCuCwfWnOpKIJVKbFPa",
          });
        }
        if (id == "root-hammerhead-shadow-ui") {
          return null;
        }
        throw ["getElementById", ...arguments, this];
      },
      addEventListener(name, cb) {
        if (name.startsWith("mouse")) {
          return;
        }
        if (name.startsWith("key")) {
          return;
        }
        if (name.startsWith("touch")) {
          return;
        }
        if (name === "input") {
          return;
        }
        if (name === "click") {
          return;
        }
        if (name === "scroll") {
          return;
        }
        if (name === "mousemove") {
          // cb();
          return;
        }
        if (name === "driver-evaluate") {
          return;
        }
        if (name === "webdriver-evaluate") {
          return;
        }
        if (name === "selenium-evaluate") {
          return;
        }
        if (name === "error") {
          return;
        }
        throw ["document.addEventListener", ...arguments];
      },
    },
    documentPossibleProps
  );
  const onload = [() => {}];
  const preventAddEventListener = new Set("");
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
            hostname: "{{.URL.Host}}",
            href: "{{.URL.String}}",
            pathname: "{{.URL.Path}}",
            port: "",
            search: "",
            replace(url) {
              if (this.pathname === url) {
                return;
              }
              throw ["location.replace", ...arguments, this];
            },
          }),
          setTimeout(cb, d) {
            if (d === 0) {
              cb();
              return;
            }
            throw ["setTimeout", ...arguments];
          },
          setInterval(cb, d) {
            if (d === 2047 || d == 50e3 || d === 2000 || d == 1500) {
              return;
            }

            throw ["setInterval", ...arguments];
          },
          XMLHttpRequest() {
            throw arguments;
          },
          onload: () => {},
          onunload: () => {},
          onbeforeunload: () => {},
          addEventListener(name, cb) {
            if (preventAddEventListener.has(name)) {
              throw arguments;
            }
            if (name === "load") {
              onload.push(cb);
              // preventAddEventListener.add(name);
              return;
            }
            if (name === "unload") {
              onunload.push(cb);
              preventAddEventListener.add(name);
              return;
            }
            if (name === "error") {
              return;
            }
            throw ["window.addEventListener", ...arguments];
          },
          /**
           *
           * @param { {type: string} } e
           */
          dispatchEvent(e) {
            if (e.type === "load") {
              this.onload();
              this.onload = () => {};
              onload.forEach((i) => i());
              onload.length = 0;
              return;
            }
            if (e.type === "unload") {
              this.onunload();
              this.onunload = () => {};
              onunload.forEach((i) => i());
              onunload.length = 0;
              return;
            }
            throw arguments;
          },
          $_ts: undefined,
          "2eab37c0ead4b0b0": undefined,
          $b_onBridgeReady: undefined,
          $b_setup: undefined,
          navigator: {
            userAgent:
              "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:102.0) Gecko/20100101 Firefox/102.0",
            language: "zh-CN",
          },
        },

        windowPossibleProps
      ),
      {
        set(obj, prop, value) {
          obj[prop] = value;
          globalThis[prop] = value;
          return true;
        },
        get(obj, prop) {
          if (prop in obj) {
            return obj[prop];
          }
          if (prop in globalThis) {
            return globalThis[prop];
          }
        },
      }
    )
  );

  return {
    window,
    document,
  };
})();
