(window.webpackJsonp=window.webpackJsonp||[]).push([[4],{"Sx/Y":function(A,d,e){"use strict";e.r(d);var m=e("9og8"),g=e("WmNS"),u=e.n(g),o=e("q1tI"),h=e.n(o),f=e("fe1z"),v=h.a.Fragment,p=function(_){var y=_.children;return Object(o.useEffect)(function(){var i=function(){var s=Object(m.a)(u.a.mark(function n(){var a,r,t;return u.a.wrap(function(l){for(;;)switch(l.prev=l.next){case 0:return l.next=2,Object(f.a)();case 2:a=l.sent,r=a.code,t=a.data,r==200&&localStorage.setItem("address",t.address);case 6:case"end":return l.stop()}},n)}));return function(){return s.apply(this,arguments)}}();localStorage.getItem("address")||i()}),h.a.createElement(v,null,y)};d.default=p},fe1z:function(A,d,e){"use strict";e.d(d,"a",function(){return h}),e.d(d,"b",function(){return v}),e.d(d,"c",function(){return E});var m=e("9og8"),g=e("WmNS"),u=e.n(g),o=e("t3Un");function h(){return f.apply(this,arguments)}function f(){return f=Object(m.a)(u.a.mark(function s(){var n,a=arguments;return u.a.wrap(function(t){for(;;)switch(t.prev=t.next){case 0:return n=a.length>0&&a[0]!==void 0?a[0]:{},t.abrupt("return",Object(o.a)("host",{method:"GET",params:{},body:JSON.stringify(n)},Object({NODE_ENV:"production"}).getAPI));case 2:case"end":return t.stop()}},s)})),f.apply(this,arguments)}function v(){return p.apply(this,arguments)}function p(){return p=Object(m.a)(u.a.mark(function s(){var n,a=arguments;return u.a.wrap(function(t){for(;;)switch(t.prev=t.next){case 0:return n=a.length>0&&a[0]!==void 0?a[0]:{},t.abrupt("return",Object(o.a)("user/login",{method:"POST",params:{},body:JSON.stringify(n)},Object({NODE_ENV:"production"}).API));case 2:case"end":return t.stop()}},s)})),p.apply(this,arguments)}function E(){return _.apply(this,arguments)}function _(){return _=Object(m.a)(u.a.mark(function s(){var n,a=arguments;return u.a.wrap(function(t){for(;;)switch(t.prev=t.next){case 0:return n=a.length>0&&a[0]!==void 0?a[0]:{},t.abrupt("return",Object(o.a)("user/exit",{method:"GET",params:{},body:JSON.stringify(n)},Object({NODE_ENV:"production"}).API));case 2:case"end":return t.stop()}},s)})),_.apply(this,arguments)}function y(){return i.apply(this,arguments)}function i(){return i=Object(m.a)(u.a.mark(function s(){var n,a=arguments;return u.a.wrap(function(t){for(;;)switch(t.prev=t.next){case 0:return n=a.length>0&&a[0]!==void 0?a[0]:{},t.abrupt("return",Object(o.a)("miner/summary",{method:"GET",params:{},body:JSON.stringify(n)},Object({NODE_ENV:"production"}).API));case 2:case"end":return t.stop()}},s)})),i.apply(this,arguments)}},t3Un:function(A,d,e){"use strict";e.d(d,"a",function(){return y});var m=e("9og8"),g=e("k1fw"),u=e("WmNS"),o=e.n(u),h=e("Qyje"),f=e.n(h),v=function(s,n){var a=s,r=Object.keys(n);return r.length!=0&&(a+="?"+f.a.stringify(n)),a},p=e("p46w"),E=e.n(p),_=e("Hg0r").fetch;function y(i){var s=arguments.length>1&&arguments[1]!==void 0?arguments[1]:{},n=arguments.length>2&&arguments[2]!==void 0?arguments[2]:"",a={method:"GET"},r=Object(g.a)(Object(g.a)({},a),s);r.headers=Object(g.a)({Accept:"application/json"},r.headers),E.a.get("token")&&(r.headers.token=E.a.get("token"));var t="";return i=="host"?t=n:t=localStorage.getItem("address")?"http://"+localStorage.getItem("address")+"/":n,i=v(i,s.params),r.method=="GET"&&delete r.body,_(t+i,r).then(function(){var P=Object(m.a)(o.a.mark(function l(O){var j,D,b;return o.a.wrap(function(c){for(;;)switch(c.prev=c.next){case 0:return j=O.status,c.next=3,O.json();case 3:if(D=c.sent,!(O.status>=200&&O.status<300)){c.next=6;break}return c.abrupt("return",D);case 6:if(!(j>=400||j>=500)){c.next=11;break}throw b=new Error,b.response=O,b.data=D,b;case 11:case"end":return c.stop()}},l)}));return function(l){return P.apply(this,arguments)}}()).catch(function(P){var l=503,O="\u94FE\u63A5\u670D\u52A1\u5668\u8D85\u65F6";throw O})}}}]);
