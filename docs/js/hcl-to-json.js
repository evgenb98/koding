(function e(t,n,r){function s(o,u){if(!n[o]){if(!t[o]){var a=typeof require=="function"&&require;if(!u&&a)return a(o,!0);if(i)return i(o,!0);var f=new Error("Cannot find module '"+o+"'");throw f.code="MODULE_NOT_FOUND",f}var l=n[o]={exports:{}};t[o][0].call(l.exports,function(e){var n=t[o][1][e];return s(n?n:e)},l,l.exports,e,t,n,r)}return n[o].exports}var i=typeof require=="function"&&require;for(var o=0;o<r.length;o++)s(r[o]);return s})({1:[function(require,module,exports){
var hcltojson;hcltojson=require("./index"),function(){return"undefined"!=typeof window&&null!==window?null!=window.hcltojson?window.hcltojson:window.hcltojson=hcltojson:void 0}();
},{"./index":2}],2:[function(require,module,exports){
var TOKENS,debug,get,hcltojson,set,tokenize,utils;debug=require("debug")("hcltojson:main"),get=require("lodash.get"),set=require("lodash.set"),TOKENS=require("./tokens"),tokenize=require("./tokenize"),utils=require("./utils"),module.exports=hcltojson=function(e){var t,s,i,r,u,n,o,a,l,g,E,c,p,d,h,b,f;if(debug("input",e),!(null!=e&&"function"==typeof e.trim?e.trim():void 0))return{};for(l=""+e,l=tokenize(l),E={},f=[],o=0,a=0,g=0,t=!1,i=!1,r="",u=null,b=l.split(TOKENS.SPACE),n=0,c=b.length;n<c;n++)if(d=b[n],d&&d!==TOKENS.COMMA)if(debug("json",JSON.stringify(E)),h=f.join("."),debug("path",h),t||(d=d.replace(/\"/g,"")),i)u=d,i=!1,r=utils.getHereDoc(e,u),debug("heredoc",d,r);else{if(u){if(d!==u)continue;d=r,u=null,r="",t=!0}if(t){switch(debug("assignment",d),d){case TOKENS.LBRACKET:set(E,h,[]),a++;break;case TOKENS.LBRACE:o++;break;case TOKENS.RBRACE:o--,o<=0&&(f=[]);break;case TOKENS.HEREDOC:i=!0;break;default:d=utils.doConversions(d),set(E,h,d),f.pop(),o<=0&&(f=[])}t=!1}else switch(d){case TOKENS.LBRACE:o++;break;case TOKENS.RBRACE:f.pop(),o--,g>0&&(f.pop(),g--),o<=0&&(f=[]);break;case TOKENS.RBRACKET:f.pop(),a--;break;case TOKENS.ASSIGN:t=!0;break;default:if(a>0)p=get(E,h),p.push(d),set(E,h,p);else if(f.push(d),h=f.join("."),debug("path end",h),f.length>2&&(s=get(E,h)))g++,Array.isArray(s)?(s.push({}),set(E,h,s),f.push(""+(s.length-1))):(s=[s,{}],set(E,h,s),f.push(""+g));else{if(get(E,h))continue;set(E,h,{})}}}return E};
},{"./tokenize":3,"./tokens":4,"./utils":5,"debug":6,"lodash.get":9,"lodash.set":10}],3:[function(require,module,exports){
var CHARS,TOKENS,debug,tokenize;debug=require("debug")("hcltojson:tokenize"),TOKENS=require("./tokens"),CHARS=[[",",TOKENS.COMMA],[":",TOKENS.COLON],["=",TOKENS.ASSIGN],["{",TOKENS.LBRACE],["}",TOKENS.RBRACE],["[",TOKENS.LBRACKET],["]",TOKENS.RBRACKET],["<<",TOKENS.HEREDOC]],module.exports=tokenize=function(e){var r,n,E,S,u,o;for(debug("input",e),e=e.replace(/#.+/gim,""),e=e.replace(/".+?"|\/\*[\s\S]*?\*\/|\/\/.*/gm,function(e){return'"'===(null!=e?e[0]:void 0)?e:""}),debug("comments.removed",e),n=0,E=CHARS.length;n<E;n++)S=CHARS[n],r=S[0],o=S[1],r="\\"+r,u=RegExp('".+?"|('+r+")","g"),e=e.replace(u,function(e,r){return r?" "+o+" ":e});return debug("tokenized",e),e=e.replace(/\s+(?=([^"]*"[^"]*")*[^"]*$)/gm,TOKENS.SPACE),debug("spaces.marked",e),e};

},{"./tokens":4,"debug":6}],4:[function(require,module,exports){
var TOKENS;module.exports=TOKENS={COMMA:"&#x2C",COLON:"&#x3A",ASSIGN:"&#x3D",LBRACE:"&#x7B",RBRACE:"&#x7D",LBRACKET:"&#x5B",RBRACKET:"&#x5D",HEREDOC:"&#xAB",SPACE:"&#x20"};
},{}],5:[function(require,module,exports){
var debug,utils;debug=require("debug")("hcltojson:utils"),module.exports=utils={getHereDoc:function(e,o){var s,r,n;return debug("getHereDoc",e,o),r=RegExp("<<\\s?(.+?)\\s((?:.|\\n)*)"+o,"gi"),n=r.exec(e),debug("getHereDoc.res",n),s=n[n.length-1],s.trimRight("\n")},doConversions:function(e){return debug("doConversions",e),e=e.replace(/^\"/,"").replace(/\"$/,"").replace(/\"/g,'"'),debug("doConversions.quotes.processed",e),"false"===e&&(e=!1),"true"===e&&(e=!0),debug("doConversions.boolean.processed",e),isNaN(""+e)||(e=+e),debug("doConversions.numerics.processed",e),e}};

},{"debug":6}],6:[function(require,module,exports){
function useColors(){return"WebkitAppearance"in document.documentElement.style||window.console&&(console.firebug||console.exception&&console.table)||navigator.userAgent.toLowerCase().match(/firefox\/(\d+)/)&&parseInt(RegExp.$1,10)>=31}function formatArgs(){var o=arguments,e=this.useColors;if(o[0]=(e?"%c":"")+this.namespace+(e?" %c":" ")+o[0]+(e?"%c ":" ")+"+"+exports.humanize(this.diff),!e)return o;var r="color: "+this.color;o=[o[0],r,"color: inherit"].concat(Array.prototype.slice.call(o,1));var t=0,s=0;return o[0].replace(/%[a-z%]/g,function(o){"%%"!==o&&(t++,"%c"===o&&(s=t))}),o.splice(s,0,r),o}function log(){return"object"==typeof console&&console.log&&Function.prototype.apply.call(console.log,console,arguments)}function save(o){try{null==o?exports.storage.removeItem("debug"):exports.storage.debug=o}catch(o){}}function load(){var o;try{o=exports.storage.debug}catch(o){}return o}function localstorage(){try{return window.localStorage}catch(o){}}exports=module.exports=require("./debug"),exports.log=log,exports.formatArgs=formatArgs,exports.save=save,exports.load=load,exports.useColors=useColors,exports.storage="undefined"!=typeof chrome&&"undefined"!=typeof chrome.storage?chrome.storage.local:localstorage(),exports.colors=["lightseagreen","forestgreen","goldenrod","dodgerblue","darkorchid","crimson"],exports.formatters.j=function(o){return JSON.stringify(o)},exports.enable(load());
},{"./debug":7}],7:[function(require,module,exports){
function selectColor(){return exports.colors[prevColor++%exports.colors.length]}function debug(e){function r(){}function o(){var e=o,r=+new Date,s=r-(prevTime||r);e.diff=s,e.prev=prevTime,e.curr=r,prevTime=r,null==e.useColors&&(e.useColors=exports.useColors()),null==e.color&&e.useColors&&(e.color=selectColor());var t=Array.prototype.slice.call(arguments);t[0]=exports.coerce(t[0]),"string"!=typeof t[0]&&(t=["%o"].concat(t));var n=0;t[0]=t[0].replace(/%([a-z%])/g,function(r,o){if("%%"===r)return r;n++;var s=exports.formatters[o];if("function"==typeof s){var p=t[n];r=s.call(e,p),t.splice(n,1),n--}return r}),"function"==typeof exports.formatArgs&&(t=exports.formatArgs.apply(e,t));var p=o.log||exports.log||console.log.bind(console);p.apply(e,t)}r.enabled=!1,o.enabled=!0;var s=exports.enabled(e)?o:r;return s.namespace=e,s}function enable(e){exports.save(e);for(var r=(e||"").split(/[\s,]+/),o=r.length,s=0;s<o;s++)r[s]&&(e=r[s].replace(/\*/g,".*?"),"-"===e[0]?exports.skips.push(new RegExp("^"+e.substr(1)+"$")):exports.names.push(new RegExp("^"+e+"$")))}function disable(){exports.enable("")}function enabled(e){var r,o;for(r=0,o=exports.skips.length;r<o;r++)if(exports.skips[r].test(e))return!1;for(r=0,o=exports.names.length;r<o;r++)if(exports.names[r].test(e))return!0;return!1}function coerce(e){return e instanceof Error?e.stack||e.message:e}exports=module.exports=debug,exports.coerce=coerce,exports.disable=disable,exports.enable=enable,exports.enabled=enabled,exports.humanize=require("ms"),exports.names=[],exports.skips=[],exports.formatters={};var prevColor=0,prevTime;
},{"ms":8}],8:[function(require,module,exports){
function parse(e){if(e=""+e,!(e.length>1e4)){var a=/^((?:\d+)?\.?\d+) *(milliseconds?|msecs?|ms|seconds?|secs?|s|minutes?|mins?|m|hours?|hrs?|h|days?|d|years?|yrs?|y)?$/i.exec(e);if(a){var r=parseFloat(a[1]),c=(a[2]||"ms").toLowerCase();switch(c){case"years":case"year":case"yrs":case"yr":case"y":return r*y;case"days":case"day":case"d":return r*d;case"hours":case"hour":case"hrs":case"hr":case"h":return r*h;case"minutes":case"minute":case"mins":case"min":case"m":return r*m;case"seconds":case"second":case"secs":case"sec":case"s":return r*s;case"milliseconds":case"millisecond":case"msecs":case"msec":case"ms":return r}}}}function short(e){return e>=d?Math.round(e/d)+"d":e>=h?Math.round(e/h)+"h":e>=m?Math.round(e/m)+"m":e>=s?Math.round(e/s)+"s":e+"ms"}function long(e){return plural(e,d,"day")||plural(e,h,"hour")||plural(e,m,"minute")||plural(e,s,"second")||e+" ms"}function plural(s,e,a){if(!(s<e))return s<1.5*e?Math.floor(s/e)+" "+a:Math.ceil(s/e)+" "+a+"s"}var s=1e3,m=60*s,h=60*m,d=24*h,y=365.25*d;module.exports=function(s,e){return e=e||{},"string"==typeof s?parse(s):e.long?long(s):short(s)};
},{}],9:[function(require,module,exports){
(function (global){
function getValue(t,e){return null==t?void 0:t[e]}function isHostObject(t){var e=!1;if(null!=t&&"function"!=typeof t.toString)try{e=!!(t+"")}catch(t){}return e}function Hash(t){var e=-1,r=t?t.length:0;for(this.clear();++e<r;){var a=t[e];this.set(a[0],a[1])}}function hashClear(){this.__data__=nativeCreate?nativeCreate(null):{}}function hashDelete(t){return this.has(t)&&delete this.__data__[t]}function hashGet(t){var e=this.__data__;if(nativeCreate){var r=e[t];return r===HASH_UNDEFINED?void 0:r}return hasOwnProperty.call(e,t)?e[t]:void 0}function hashHas(t){var e=this.__data__;return nativeCreate?void 0!==e[t]:hasOwnProperty.call(e,t)}function hashSet(t,e){var r=this.__data__;return r[t]=nativeCreate&&void 0===e?HASH_UNDEFINED:e,this}function ListCache(t){var e=-1,r=t?t.length:0;for(this.clear();++e<r;){var a=t[e];this.set(a[0],a[1])}}function listCacheClear(){this.__data__=[]}function listCacheDelete(t){var e=this.__data__,r=assocIndexOf(e,t);if(r<0)return!1;var a=e.length-1;return r==a?e.pop():splice.call(e,r,1),!0}function listCacheGet(t){var e=this.__data__,r=assocIndexOf(e,t);return r<0?void 0:e[r][1]}function listCacheHas(t){return assocIndexOf(this.__data__,t)>-1}function listCacheSet(t,e){var r=this.__data__,a=assocIndexOf(r,t);return a<0?r.push([t,e]):r[a][1]=e,this}function MapCache(t){var e=-1,r=t?t.length:0;for(this.clear();++e<r;){var a=t[e];this.set(a[0],a[1])}}function mapCacheClear(){this.__data__={hash:new Hash,map:new(Map||ListCache),string:new Hash}}function mapCacheDelete(t){return getMapData(this,t).delete(t)}function mapCacheGet(t){return getMapData(this,t).get(t)}function mapCacheHas(t){return getMapData(this,t).has(t)}function mapCacheSet(t,e){return getMapData(this,t).set(t,e),this}function assocIndexOf(t,e){for(var r=t.length;r--;)if(eq(t[r][0],e))return r;return-1}function baseGet(t,e){e=isKey(e,t)?[e]:castPath(e);for(var r=0,a=e.length;null!=t&&r<a;)t=t[toKey(e[r++])];return r&&r==a?t:void 0}function baseIsNative(t){if(!isObject(t)||isMasked(t))return!1;var e=isFunction(t)||isHostObject(t)?reIsNative:reIsHostCtor;return e.test(toSource(t))}function baseToString(t){if("string"==typeof t)return t;if(isSymbol(t))return symbolToString?symbolToString.call(t):"";var e=t+"";return"0"==e&&1/t==-INFINITY?"-0":e}function castPath(t){return isArray(t)?t:stringToPath(t)}function getMapData(t,e){var r=t.__data__;return isKeyable(e)?r["string"==typeof e?"string":"hash"]:r.map}function getNative(t,e){var r=getValue(t,e);return baseIsNative(r)?r:void 0}function isKey(t,e){if(isArray(t))return!1;var r=typeof t;return!("number"!=r&&"symbol"!=r&&"boolean"!=r&&null!=t&&!isSymbol(t))||(reIsPlainProp.test(t)||!reIsDeepProp.test(t)||null!=e&&t in Object(e))}function isKeyable(t){var e=typeof t;return"string"==e||"number"==e||"symbol"==e||"boolean"==e?"__proto__"!==t:null===t}function isMasked(t){return!!maskSrcKey&&maskSrcKey in t}function toKey(t){if("string"==typeof t||isSymbol(t))return t;var e=t+"";return"0"==e&&1/t==-INFINITY?"-0":e}function toSource(t){if(null!=t){try{return funcToString.call(t)}catch(t){}try{return t+""}catch(t){}}return""}function memoize(t,e){if("function"!=typeof t||e&&"function"!=typeof e)throw new TypeError(FUNC_ERROR_TEXT);var r=function(){var a=arguments,o=e?e.apply(this,a):a[0],n=r.cache;if(n.has(o))return n.get(o);var i=t.apply(this,a);return r.cache=n.set(o,i),i};return r.cache=new(memoize.Cache||MapCache),r}function eq(t,e){return t===e||t!==t&&e!==e}function isFunction(t){var e=isObject(t)?objectToString.call(t):"";return e==funcTag||e==genTag}function isObject(t){var e=typeof t;return!!t&&("object"==e||"function"==e)}function isObjectLike(t){return!!t&&"object"==typeof t}function isSymbol(t){return"symbol"==typeof t||isObjectLike(t)&&objectToString.call(t)==symbolTag}function toString(t){return null==t?"":baseToString(t)}function get(t,e,r){var a=null==t?void 0:baseGet(t,e);return void 0===a?r:a}var FUNC_ERROR_TEXT="Expected a function",HASH_UNDEFINED="__lodash_hash_undefined__",INFINITY=1/0,funcTag="[object Function]",genTag="[object GeneratorFunction]",symbolTag="[object Symbol]",reIsDeepProp=/\.|\[(?:[^[\]]*|(["'])(?:(?!\1)[^\\]|\\.)*?\1)\]/,reIsPlainProp=/^\w*$/,reLeadingDot=/^\./,rePropName=/[^.[\]]+|\[(?:(-?\d+(?:\.\d+)?)|(["'])((?:(?!\2)[^\\]|\\.)*?)\2)\]|(?=(?:\.|\[\])(?:\.|\[\]|$))/g,reRegExpChar=/[\\^$.*+?()[\]{}|]/g,reEscapeChar=/\\(\\)?/g,reIsHostCtor=/^\[object .+?Constructor\]$/,freeGlobal="object"==typeof global&&global&&global.Object===Object&&global,freeSelf="object"==typeof self&&self&&self.Object===Object&&self,root=freeGlobal||freeSelf||Function("return this")(),arrayProto=Array.prototype,funcProto=Function.prototype,objectProto=Object.prototype,coreJsData=root["__core-js_shared__"],maskSrcKey=function(){var t=/[^.]+$/.exec(coreJsData&&coreJsData.keys&&coreJsData.keys.IE_PROTO||"");return t?"Symbol(src)_1."+t:""}(),funcToString=funcProto.toString,hasOwnProperty=objectProto.hasOwnProperty,objectToString=objectProto.toString,reIsNative=RegExp("^"+funcToString.call(hasOwnProperty).replace(reRegExpChar,"\\$&").replace(/hasOwnProperty|(function).*?(?=\\\()| for .+?(?=\\\])/g,"$1.*?")+"$"),Symbol=root.Symbol,splice=arrayProto.splice,Map=getNative(root,"Map"),nativeCreate=getNative(Object,"create"),symbolProto=Symbol?Symbol.prototype:void 0,symbolToString=symbolProto?symbolProto.toString:void 0;Hash.prototype.clear=hashClear,Hash.prototype.delete=hashDelete,Hash.prototype.get=hashGet,Hash.prototype.has=hashHas,Hash.prototype.set=hashSet,ListCache.prototype.clear=listCacheClear,ListCache.prototype.delete=listCacheDelete,ListCache.prototype.get=listCacheGet,ListCache.prototype.has=listCacheHas,ListCache.prototype.set=listCacheSet,MapCache.prototype.clear=mapCacheClear,MapCache.prototype.delete=mapCacheDelete,MapCache.prototype.get=mapCacheGet,MapCache.prototype.has=mapCacheHas,MapCache.prototype.set=mapCacheSet;var stringToPath=memoize(function(t){t=toString(t);var e=[];return reLeadingDot.test(t)&&e.push(""),t.replace(rePropName,function(t,r,a,o){e.push(a?o.replace(reEscapeChar,"$1"):r||t)}),e});memoize.Cache=MapCache;var isArray=Array.isArray;module.exports=get;

}).call(this,typeof global !== "undefined" ? global : typeof self !== "undefined" ? self : typeof window !== "undefined" ? window : {})
},{}],10:[function(require,module,exports){
(function (global){
function getValue(t,e){return null==t?void 0:t[e]}function isHostObject(t){var e=!1;if(null!=t&&"function"!=typeof t.toString)try{e=!!(t+"")}catch(t){}return e}function Hash(t){var e=-1,r=t?t.length:0;for(this.clear();++e<r;){var a=t[e];this.set(a[0],a[1])}}function hashClear(){this.__data__=nativeCreate?nativeCreate(null):{}}function hashDelete(t){return this.has(t)&&delete this.__data__[t]}function hashGet(t){var e=this.__data__;if(nativeCreate){var r=e[t];return r===HASH_UNDEFINED?void 0:r}return hasOwnProperty.call(e,t)?e[t]:void 0}function hashHas(t){var e=this.__data__;return nativeCreate?void 0!==e[t]:hasOwnProperty.call(e,t)}function hashSet(t,e){var r=this.__data__;return r[t]=nativeCreate&&void 0===e?HASH_UNDEFINED:e,this}function ListCache(t){var e=-1,r=t?t.length:0;for(this.clear();++e<r;){var a=t[e];this.set(a[0],a[1])}}function listCacheClear(){this.__data__=[]}function listCacheDelete(t){var e=this.__data__,r=assocIndexOf(e,t);if(r<0)return!1;var a=e.length-1;return r==a?e.pop():splice.call(e,r,1),!0}function listCacheGet(t){var e=this.__data__,r=assocIndexOf(e,t);return r<0?void 0:e[r][1]}function listCacheHas(t){return assocIndexOf(this.__data__,t)>-1}function listCacheSet(t,e){var r=this.__data__,a=assocIndexOf(r,t);return a<0?r.push([t,e]):r[a][1]=e,this}function MapCache(t){var e=-1,r=t?t.length:0;for(this.clear();++e<r;){var a=t[e];this.set(a[0],a[1])}}function mapCacheClear(){this.__data__={hash:new Hash,map:new(Map||ListCache),string:new Hash}}function mapCacheDelete(t){return getMapData(this,t).delete(t)}function mapCacheGet(t){return getMapData(this,t).get(t)}function mapCacheHas(t){return getMapData(this,t).has(t)}function mapCacheSet(t,e){return getMapData(this,t).set(t,e),this}function assignValue(t,e,r){var a=t[e];hasOwnProperty.call(t,e)&&eq(a,r)&&(void 0!==r||e in t)||(t[e]=r)}function assocIndexOf(t,e){for(var r=t.length;r--;)if(eq(t[r][0],e))return r;return-1}function baseIsNative(t){if(!isObject(t)||isMasked(t))return!1;var e=isFunction(t)||isHostObject(t)?reIsNative:reIsHostCtor;return e.test(toSource(t))}function baseSet(t,e,r,a){if(!isObject(t))return t;e=isKey(e,t)?[e]:castPath(e);for(var o=-1,n=e.length,i=n-1,s=t;null!=s&&++o<n;){var c=toKey(e[o]),u=r;if(o!=i){var l=s[c];u=a?a(l,c,s):void 0,void 0===u&&(u=isObject(l)?l:isIndex(e[o+1])?[]:{})}assignValue(s,c,u),s=s[c]}return t}function baseToString(t){if("string"==typeof t)return t;if(isSymbol(t))return symbolToString?symbolToString.call(t):"";var e=t+"";return"0"==e&&1/t==-INFINITY?"-0":e}function castPath(t){return isArray(t)?t:stringToPath(t)}function getMapData(t,e){var r=t.__data__;return isKeyable(e)?r["string"==typeof e?"string":"hash"]:r.map}function getNative(t,e){var r=getValue(t,e);return baseIsNative(r)?r:void 0}function isIndex(t,e){return e=null==e?MAX_SAFE_INTEGER:e,!!e&&("number"==typeof t||reIsUint.test(t))&&t>-1&&t%1==0&&t<e}function isKey(t,e){if(isArray(t))return!1;var r=typeof t;return!("number"!=r&&"symbol"!=r&&"boolean"!=r&&null!=t&&!isSymbol(t))||(reIsPlainProp.test(t)||!reIsDeepProp.test(t)||null!=e&&t in Object(e))}function isKeyable(t){var e=typeof t;return"string"==e||"number"==e||"symbol"==e||"boolean"==e?"__proto__"!==t:null===t}function isMasked(t){return!!maskSrcKey&&maskSrcKey in t}function toKey(t){if("string"==typeof t||isSymbol(t))return t;var e=t+"";return"0"==e&&1/t==-INFINITY?"-0":e}function toSource(t){if(null!=t){try{return funcToString.call(t)}catch(t){}try{return t+""}catch(t){}}return""}function memoize(t,e){if("function"!=typeof t||e&&"function"!=typeof e)throw new TypeError(FUNC_ERROR_TEXT);var r=function(){var a=arguments,o=e?e.apply(this,a):a[0],n=r.cache;if(n.has(o))return n.get(o);var i=t.apply(this,a);return r.cache=n.set(o,i),i};return r.cache=new(memoize.Cache||MapCache),r}function eq(t,e){return t===e||t!==t&&e!==e}function isFunction(t){var e=isObject(t)?objectToString.call(t):"";return e==funcTag||e==genTag}function isObject(t){var e=typeof t;return!!t&&("object"==e||"function"==e)}function isObjectLike(t){return!!t&&"object"==typeof t}function isSymbol(t){return"symbol"==typeof t||isObjectLike(t)&&objectToString.call(t)==symbolTag}function toString(t){return null==t?"":baseToString(t)}function set(t,e,r){return null==t?t:baseSet(t,e,r)}var FUNC_ERROR_TEXT="Expected a function",HASH_UNDEFINED="__lodash_hash_undefined__",INFINITY=1/0,MAX_SAFE_INTEGER=9007199254740991,funcTag="[object Function]",genTag="[object GeneratorFunction]",symbolTag="[object Symbol]",reIsDeepProp=/\.|\[(?:[^[\]]*|(["'])(?:(?!\1)[^\\]|\\.)*?\1)\]/,reIsPlainProp=/^\w*$/,reLeadingDot=/^\./,rePropName=/[^.[\]]+|\[(?:(-?\d+(?:\.\d+)?)|(["'])((?:(?!\2)[^\\]|\\.)*?)\2)\]|(?=(?:\.|\[\])(?:\.|\[\]|$))/g,reRegExpChar=/[\\^$.*+?()[\]{}|]/g,reEscapeChar=/\\(\\)?/g,reIsHostCtor=/^\[object .+?Constructor\]$/,reIsUint=/^(?:0|[1-9]\d*)$/,freeGlobal="object"==typeof global&&global&&global.Object===Object&&global,freeSelf="object"==typeof self&&self&&self.Object===Object&&self,root=freeGlobal||freeSelf||Function("return this")(),arrayProto=Array.prototype,funcProto=Function.prototype,objectProto=Object.prototype,coreJsData=root["__core-js_shared__"],maskSrcKey=function(){var t=/[^.]+$/.exec(coreJsData&&coreJsData.keys&&coreJsData.keys.IE_PROTO||"");return t?"Symbol(src)_1."+t:""}(),funcToString=funcProto.toString,hasOwnProperty=objectProto.hasOwnProperty,objectToString=objectProto.toString,reIsNative=RegExp("^"+funcToString.call(hasOwnProperty).replace(reRegExpChar,"\\$&").replace(/hasOwnProperty|(function).*?(?=\\\()| for .+?(?=\\\])/g,"$1.*?")+"$"),Symbol=root.Symbol,splice=arrayProto.splice,Map=getNative(root,"Map"),nativeCreate=getNative(Object,"create"),symbolProto=Symbol?Symbol.prototype:void 0,symbolToString=symbolProto?symbolProto.toString:void 0;Hash.prototype.clear=hashClear,Hash.prototype.delete=hashDelete,Hash.prototype.get=hashGet,Hash.prototype.has=hashHas,Hash.prototype.set=hashSet,ListCache.prototype.clear=listCacheClear,ListCache.prototype.delete=listCacheDelete,ListCache.prototype.get=listCacheGet,ListCache.prototype.has=listCacheHas,ListCache.prototype.set=listCacheSet,MapCache.prototype.clear=mapCacheClear,MapCache.prototype.delete=mapCacheDelete,MapCache.prototype.get=mapCacheGet,MapCache.prototype.has=mapCacheHas,MapCache.prototype.set=mapCacheSet;var stringToPath=memoize(function(t){t=toString(t);var e=[];return reLeadingDot.test(t)&&e.push(""),t.replace(rePropName,function(t,r,a,o){e.push(a?o.replace(reEscapeChar,"$1"):r||t)}),e});memoize.Cache=MapCache;var isArray=Array.isArray;module.exports=set;

}).call(this,typeof global !== "undefined" ? global : typeof self !== "undefined" ? self : typeof window !== "undefined" ? window : {})
},{}]},{},[1]);
