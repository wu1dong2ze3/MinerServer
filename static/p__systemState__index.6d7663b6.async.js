(window.webpackJsonp=window.webpackJsonp||[]).push([[16],{"5dcc":function(S,o,e){"use strict";e.d(o,"c",function(){return j}),e.d(o,"d",function(){return I}),e.d(o,"g",function(){return E}),e.d(o,"e",function(){return d}),e.d(o,"a",function(){return v}),e.d(o,"b",function(){return r}),e.d(o,"f",function(){return f});var g=e("jrin"),O=e("9og8"),u=e("WmNS"),c=e.n(u),b=e("MZJn"),t=e("fe1z"),N=e("p46w"),P=e.n(N),i=e("9kvl");function j(a){return a===""||a===void 0||a===null}function I(a){return a===0?b.a:a}function E(a,l){return j(a)?l:a}function h(a,l){var n=(a-l)/1e3/3600/24,y=Math.floor(n),s=(n-y)*24,M=Math.floor(s),p=(s-M)*60,T=Math.floor(p),B=(p-T)*60,H=Math.floor(B),R=y+"\u5929"+M+"\u5C0F\u65F6"+T+"\u5206";return R}function d(){return m.apply(this,arguments)}function m(){return m=Object(O.a)(c.a.mark(function a(){var l,n;return c.a.wrap(function(s){for(;;)switch(s.prev=s.next){case 0:return s.next=2,Object(t.c)();case 2:l=s.sent,n=l.code,n==200&&(i.e.push("/user/login"),localStorage.getItem("address")&&localStorage.removeItem("address"),P.a.remove("token"));case 5:case"end":return s.stop()}},a)})),m.apply(this,arguments)}function v(a){var l=_.trim(a);return l=l.replace(/[^\d^\.]+/g,"").replace(".","$#$").replace(/\./g,"").replace("$#$","."),l}function r(a){var l=_.trim(a);return l=l.replace(/[^\x00-\xff]/g,""),l}function f(a,l){for(var n in a)l.setFieldsValue(Object(g.a)({},n,a[n]))}},"6Ynu":function(S,o,e){"use strict";var g=e("k1fw"),O=e("q1tI"),u=e.n(O),c=function(t){return u.a.createElement("div",{style:Object(g.a)({height:"100%",background:"#fff",borderRadius:"7px"},t.style)},u.a.createElement("div",{style:{height:"54px",lineHeight:"54px",borderBottom:"1px solid #F5F7FD",fontWeight:"600",paddingLeft:"30px",color:"#333",fontSize:"18px"}},t.title," "),u.a.createElement("div",{style:{padding:"0px 30px",marginTop:"24px",color:"#333"}},t.children))};o.a=c},MZJn:function(S,o,e){"use strict";e.d(o,"a",function(){return g});var g="--"},f7Z1:function(S,o,e){S.exports={itemcol:"itemcol___1J_wh",title:"title___2Szav",line:"line___21GXQ",progress:"progress___ZsQHr",rpds:"rpds___1cm0v"}},fe1z:function(S,o,e){"use strict";e.d(o,"a",function(){return b}),e.d(o,"b",function(){return N}),e.d(o,"c",function(){return i});var g=e("9og8"),O=e("WmNS"),u=e.n(O),c=e("t3Un");function b(){return t.apply(this,arguments)}function t(){return t=Object(g.a)(u.a.mark(function h(){var d,m=arguments;return u.a.wrap(function(r){for(;;)switch(r.prev=r.next){case 0:return d=m.length>0&&m[0]!==void 0?m[0]:{},r.abrupt("return",Object(c.a)("host",{method:"GET",params:{},body:JSON.stringify(d)},Object({NODE_ENV:"production"}).getAPI));case 2:case"end":return r.stop()}},h)})),t.apply(this,arguments)}function N(){return P.apply(this,arguments)}function P(){return P=Object(g.a)(u.a.mark(function h(){var d,m=arguments;return u.a.wrap(function(r){for(;;)switch(r.prev=r.next){case 0:return d=m.length>0&&m[0]!==void 0?m[0]:{},r.abrupt("return",Object(c.a)("user/login",{method:"POST",params:{},body:JSON.stringify(d)},Object({NODE_ENV:"production"}).API));case 2:case"end":return r.stop()}},h)})),P.apply(this,arguments)}function i(){return j.apply(this,arguments)}function j(){return j=Object(g.a)(u.a.mark(function h(){var d,m=arguments;return u.a.wrap(function(r){for(;;)switch(r.prev=r.next){case 0:return d=m.length>0&&m[0]!==void 0?m[0]:{},r.abrupt("return",Object(c.a)("user/exit",{method:"GET",params:{},body:JSON.stringify(d)},Object({NODE_ENV:"production"}).API));case 2:case"end":return r.stop()}},h)})),j.apply(this,arguments)}function I(){return E.apply(this,arguments)}function E(){return E=Object(g.a)(u.a.mark(function h(){var d,m=arguments;return u.a.wrap(function(r){for(;;)switch(r.prev=r.next){case 0:return d=m.length>0&&m[0]!==void 0?m[0]:{},r.abrupt("return",Object(c.a)("miner/summary",{method:"GET",params:{},body:JSON.stringify(d)},Object({NODE_ENV:"production"}).API));case 2:case"end":return r.stop()}},h)})),E.apply(this,arguments)}},hGqE:function(S,o,e){"use strict";e.r(o);var g=e("9og8"),O=e("tJVT"),u=e("WmNS"),c=e.n(u),b=e("q1tI"),t=e.n(b),N=e("6Ynu"),P=e("f7Z1"),i=e.n(P),j=e("9kvl"),I=e("hy7T"),E=e("t3Un");function h(){return d.apply(this,arguments)}function d(){return d=Object(g.a)(c.a.mark(function n(){var y,s=arguments;return c.a.wrap(function(p){for(;;)switch(p.prev=p.next){case 0:return y=s.length>0&&s[0]!==void 0?s[0]:{},p.abrupt("return",Object(E.a)("system/hardware/version",{method:"GET",params:{},body:JSON.stringify(y)},Object({NODE_ENV:"production"}).API));case 2:case"end":return p.stop()}},n)})),d.apply(this,arguments)}function m(){return v.apply(this,arguments)}function v(){return v=Object(g.a)(c.a.mark(function n(){var y,s=arguments;return c.a.wrap(function(p){for(;;)switch(p.prev=p.next){case 0:return y=s.length>0&&s[0]!==void 0?s[0]:{},p.abrupt("return",Object(E.a)("system/hardware/status",{method:"GET",params:{},body:JSON.stringify(y)},Object({NODE_ENV:"production"}).API));case 2:case"end":return p.stop()}},n)})),v.apply(this,arguments)}var r=e("MZJn"),f=e("5dcc"),a=function(){var y=Object(j.i)(),s=y.formatMessage,M=Object(b.useState)({model:"--",cpu:"--",firmwareVersion:"--",firmwareDate:""}),p=Object(O.a)(M,2),T=p[0],B=p[1],H=Object(b.useState)({cpu:"",totalRam:"",availableRam:"",availableRom:"",totalRom:"",time:z}),R=Object(O.a)(H,2),U=R[0],x=R[1],q=Object(b.useState)({mac:"",routerType:0,ip:"",subnetMask:"",gateway:"",dns1:"",dns2:""}),Z=Object(O.a)(q,2),A=Z[0],ee=Z[1],te=function(){var ce=Object(g.a)(c.a.mark(function Q(){var Y,X,k;return c.a.wrap(function(D){for(;;)switch(D.prev=D.next){case 0:return D.next=2,h();case 2:return Y=D.sent,D.next=5,m();case 5:return X=D.sent,D.next=8,Object(I.c)();case 8:k=D.sent,Promise.all([Y,X,k]).then(function(J){var G=J[0],w=J[1],F=J[2];G.code==200&&!_.isEmpty(G.data.version)&&B(G.data.version),w.code==200&&!_.isEmpty(w.data.status)&&x(w.data.status),F.code==200&&!_.isEmpty(F.data.net_info)&&ee(F.data.net_info)});case 10:case"end":return D.stop()}},Q)}));return function(){return ce.apply(this,arguments)}}();Object(b.useEffect)(function(){te()},[]);var ae=T.model,ne=T.cpu,re=T.firmwareVersion,se=T.firmwareDate,C=U.totalRam,W=U.availableRam,z=U.time,L=U.totalRom,V=U.availableRom,ie=A.mac,ue=A.routerType,le=A.ip,oe=A.subnetMask,me=A.dns1,de=A.dns2,K=0;W&&C&&(W<C?K=W/C*100:K=100);var $=0;return V<L?$=V/L*100:$=100,t.a.createElement("div",{style:{height:"100%"}},t.a.createElement(N.a,{title:s({id:"systemState.title"}),style:{height:"auto",marginBottom:"14px"}},t.a.createElement("div",{className:i.a.itemcol},t.a.createElement("div",null,t.a.createElement("span",{className:i.a.title},s({id:"systemState.model"})),ae),t.a.createElement("div",null,t.a.createElement("span",{className:i.a.title},s({id:"systemState.cpu"})),ne),t.a.createElement("div",null,t.a.createElement("span",{className:i.a.title},s({id:"header.version"}),":"),re),t.a.createElement("div",null,t.a.createElement("span",{className:i.a.title},s({id:"maintain.update.time"}),":"),se))),t.a.createElement(N.a,{title:s({id:"systemState.netWorkState"}),style:{height:"auto",marginBottom:"14px"}},t.a.createElement("div",{className:i.a.itemcol},t.a.createElement("div",null,t.a.createElement("span",{className:i.a.title},s({id:"config.IpConfig.macAdress"}),":"),ie),t.a.createElement("div",null,t.a.createElement("span",{className:i.a.title},s({id:"systemState.IpStyle"}),":"),ue===0?s({id:"config.IpConfig.staticIp"}):"DHCP"),t.a.createElement("div",null,t.a.createElement("span",{className:i.a.title},s({id:"config.IpConfig.ipAdress"}),":"),le),t.a.createElement("div",null,t.a.createElement("span",{className:i.a.title},s({id:"config.IpConfig.subnetMask"}),":"),oe),t.a.createElement("div",null,t.a.createElement("span",{className:i.a.title},"DNS1:"),Object(f.g)(me,r.a)),t.a.createElement("div",null,t.a.createElement("span",{className:i.a.title},"DNS2:"),Object(f.g)(de,r.a)))),t.a.createElement(N.a,{title:s({id:"systemState.hardwareStatus"}),style:{height:"auto",marginBottom:"14px"}},t.a.createElement("div",{className:i.a.itemcol},t.a.createElement("div",null,t.a.createElement("span",{className:i.a.title},s({id:"systemState.runTime"}),":"),z),t.a.createElement("div",null,t.a.createElement("span",{className:i.a.title},s({id:"systemState.cpuUse"}),":"),t.a.createElement("div",{className:i.a.line},t.a.createElement("i",{className:i.a.progress,style:{width:U.cpu?U.cpu+"%":""}}))),t.a.createElement("div",null,t.a.createElement("span",{className:i.a.title},s({id:"systemState.availableRam"}),":"),t.a.createElement("div",{className:i.a.line},t.a.createElement("i",{className:i.a.progress,style:{width:K+"%"}})),t.a.createElement("span",{className:i.a.rpds},W," kb /",C)),t.a.createElement("div",null,t.a.createElement("span",{className:i.a.title},s({id:"systemState.availableRom"}),":"),t.a.createElement("div",{className:i.a.line},t.a.createElement("i",{className:i.a.progress,style:{width:$+"%"}})),t.a.createElement("span",{className:i.a.rpds},V,"kb /",L)))))},l=o.default=t.a.memo(a)},hy7T:function(S,o,e){"use strict";e.d(o,"e",function(){return b}),e.d(o,"f",function(){return N}),e.d(o,"c",function(){return i}),e.d(o,"d",function(){return I}),e.d(o,"a",function(){return h}),e.d(o,"b",function(){return m});var g=e("9og8"),O=e("WmNS"),u=e.n(O),c=e("t3Un");function b(){return t.apply(this,arguments)}function t(){return t=Object(g.a)(u.a.mark(function r(){var f,a=arguments;return u.a.wrap(function(n){for(;;)switch(n.prev=n.next){case 0:return f=a.length>0&&a[0]!==void 0?a[0]:{},n.abrupt("return",Object(c.a)("miner/user/info",{method:"GET",params:{},body:JSON.stringify(f)},Object({NODE_ENV:"production"}).API));case 2:case"end":return n.stop()}},r)})),t.apply(this,arguments)}function N(){return P.apply(this,arguments)}function P(){return P=Object(g.a)(u.a.mark(function r(){var f,a=arguments;return u.a.wrap(function(n){for(;;)switch(n.prev=n.next){case 0:return f=a.length>0&&a[0]!==void 0?a[0]:{},n.abrupt("return",Object(c.a)("miner/user/update",{method:"POST",params:{},body:JSON.stringify(f)},Object({NODE_ENV:"production"}).API));case 2:case"end":return n.stop()}},r)})),P.apply(this,arguments)}function i(){return j.apply(this,arguments)}function j(){return j=Object(g.a)(u.a.mark(function r(){var f,a=arguments;return u.a.wrap(function(n){for(;;)switch(n.prev=n.next){case 0:return f=a.length>0&&a[0]!==void 0?a[0]:{},n.abrupt("return",Object(c.a)("system/net/info",{method:"GET",params:{},body:JSON.stringify(f)},Object({NODE_ENV:"production"}).API));case 2:case"end":return n.stop()}},r)})),j.apply(this,arguments)}function I(){return E.apply(this,arguments)}function E(){return E=Object(g.a)(u.a.mark(function r(){var f,a=arguments;return u.a.wrap(function(n){for(;;)switch(n.prev=n.next){case 0:return f=a.length>0&&a[0]!==void 0?a[0]:{},n.abrupt("return",Object(c.a)("system/net/update",{method:"POST",params:{},body:JSON.stringify(f)},Object({NODE_ENV:"production"}).API));case 2:case"end":return n.stop()}},r)})),E.apply(this,arguments)}function h(){return d.apply(this,arguments)}function d(){return d=Object(g.a)(u.a.mark(function r(){var f,a=arguments;return u.a.wrap(function(n){for(;;)switch(n.prev=n.next){case 0:return f=a.length>0&&a[0]!==void 0?a[0]:{},n.abrupt("return",Object(c.a)("miner/mode/info",{method:"GET",params:{},body:JSON.stringify(f)},Object({NODE_ENV:"production"}).API));case 2:case"end":return n.stop()}},r)})),d.apply(this,arguments)}function m(){return v.apply(this,arguments)}function v(){return v=Object(g.a)(u.a.mark(function r(){var f,a=arguments;return u.a.wrap(function(n){for(;;)switch(n.prev=n.next){case 0:return f=a.length>0&&a[0]!==void 0?a[0]:{},n.abrupt("return",Object(c.a)("miner/mode/update",{method:"POST",params:{},body:JSON.stringify(f)},Object({NODE_ENV:"production"}).API));case 2:case"end":return n.stop()}},r)})),v.apply(this,arguments)}},t3Un:function(S,o,e){"use strict";e.d(o,"a",function(){return I});var g=e("9og8"),O=e("k1fw"),u=e("WmNS"),c=e.n(u),b=e("Qyje"),t=e.n(b),N=function(h,d){var m=h,v=Object.keys(d);return v.length!=0&&(m+="?"+t.a.stringify(d)),m},P=e("p46w"),i=e.n(P),j=e("Hg0r").fetch;function I(E){var h=arguments.length>1&&arguments[1]!==void 0?arguments[1]:{},d=arguments.length>2&&arguments[2]!==void 0?arguments[2]:"",m={method:"GET"},v=Object(O.a)(Object(O.a)({},m),h);v.headers=Object(O.a)({Accept:"application/json"},v.headers),i.a.get("token")&&(v.headers.token=i.a.get("token"));var r="";return E=="host"?r=d:r=localStorage.getItem("address")?"http://"+localStorage.getItem("address")+"/":d,E=N(E,h.params),v.method=="GET"&&delete v.body,j(r+E,v).then(function(){var f=Object(g.a)(c.a.mark(function a(l){var n,y,s;return c.a.wrap(function(p){for(;;)switch(p.prev=p.next){case 0:return n=l.status,p.next=3,l.json();case 3:if(y=p.sent,!(l.status>=200&&l.status<300)){p.next=6;break}return p.abrupt("return",y);case 6:if(!(n>=400||n>=500)){p.next=11;break}throw s=new Error,s.response=l,s.data=y,s;case 11:case"end":return p.stop()}},a)}));return function(a){return f.apply(this,arguments)}}()).catch(function(f){var a=503,l="\u94FE\u63A5\u670D\u52A1\u5668\u8D85\u65F6";throw l})}}}]);