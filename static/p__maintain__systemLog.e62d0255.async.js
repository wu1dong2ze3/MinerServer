(window.webpackJsonp=window.webpackJsonp||[]).push([[14],{"4fRq":function(u,h){var e=typeof crypto!="undefined"&&crypto.getRandomValues&&crypto.getRandomValues.bind(crypto)||typeof msCrypto!="undefined"&&typeof window.msCrypto.getRandomValues=="function"&&msCrypto.getRandomValues.bind(msCrypto);if(e){var c=new Uint8Array(16);u.exports=function(){return e(c),c}}else{var f=new Array(16);u.exports=function(){for(var o=0,n;o<16;o++)(o&3)==0&&(n=Math.random()*4294967296),f[o]=n>>>((o&3)<<3)&255;return f}}},"6Ynu":function(u,h,e){"use strict";var c=e("k1fw"),f=e("q1tI"),t=e.n(f),o=function(a){return t.a.createElement("div",{style:Object(c.a)({height:"100%",background:"#fff",borderRadius:"7px"},a.style)},t.a.createElement("div",{style:{height:"54px",lineHeight:"54px",borderBottom:"1px solid #F5F7FD",fontWeight:"600",paddingLeft:"30px",color:"#333",fontSize:"18px"}},a.title," "),t.a.createElement("div",{style:{padding:"0px 30px",marginTop:"24px",color:"#333"}},a.children))};h.a=o},I2ZF:function(u,h){for(var e=[],c=0;c<256;++c)e[c]=(c+256).toString(16).substr(1);function f(t,o){var n=o||0,a=e;return[a[t[n++]],a[t[n++]],a[t[n++]],a[t[n++]],"-",a[t[n++]],a[t[n++]],"-",a[t[n++]],a[t[n++]],"-",a[t[n++]],a[t[n++]],"-",a[t[n++]],a[t[n++]],a[t[n++]],a[t[n++]],a[t[n++]],a[t[n++]]].join("")}u.exports=f},IK7K:function(u,h,e){u.exports={box:"box___2g2_8"}},UbLO:function(u,h,e){"use strict";e.r(h);var c=e("q1tI"),f=e.n(c),t=e("6Ynu"),o=e("9kvl"),n=e("/MKj"),a=e("IK7K"),C=e.n(a),i=e("9og8"),I=e("oBTY"),M=e("WmNS"),s=e.n(M),l=e("p46w"),L=e.n(l),g=e("xDdU"),x=e.n(g),r,y,T=localStorage.getItem("address")?"ws://"+localStorage.getItem("address")+"/system/log":"/";function S(){var D=arguments.length>0&&arguments[0]!==void 0?arguments[0]:{};r&&r.close(),y&&clearInterval(y);var w=L.a.get("token")||"",R=x()();r=new WebSocket(T),r.onopen=function(){var m={op:"login",token:new String(w),sign:new String(R)};r.send(JSON.stringify(m));var m={op:"register",sign:new String(R)};r.send(JSON.stringify(m)),y=setInterval(function(){var d={op:"heartBeat"};r.send(JSON.stringify(d))},15e3)},r.onmessage=function(m){var d=JSON.parse(m.data),p=Object(o.c)()._store,j=p.getState(),v=j.systemLog;if(d.type=="normal"){var P=Object(I.a)(v.log);P.push(d.data),p.dispatch({type:"systemLog/setState",payload:{log:P}})}},r.onerror=function(){var m=Object(i.a)(s.a.mark(function d(p){return s.a.wrap(function(v){for(;;)switch(v.prev=v.next){case 0:p.code==1006&&E(S);case 1:case"end":return v.stop()}},d)}));return function(d){return m.apply(this,arguments)}}(),r.onclose=function(){var m=Object(i.a)(s.a.mark(function d(p){return s.a.wrap(function(v){for(;;)switch(v.prev=v.next){case 0:p.code==1006&&E(S);case 1:case"end":return v.stop()}},d)}));return function(d){return m.apply(this,arguments)}}()}var O;function E(){var D=arguments.length>0&&arguments[0]!==void 0?arguments[0]:function(){};clearTimeout(O),O=setTimeout(function(){D(),clearTimeout(O)},2e3)}function U(){r&&r.close()}var k=function(){var w=Object(n.f)(),R=Object(o.i)(),m=R.formatMessage,d=Object(n.g)(function(j){return j.systemLog}),p=d.log;return Object(c.useEffect)(function(){return S(),function(){U(),w({type:"systemLog/setState",payload:{log:[]}})}},[]),f.a.createElement("div",{style:{height:"100%"}},f.a.createElement(t.a,{title:m({id:"maintain.systemLog.title"})},f.a.createElement("div",{className:C.a.box},p&&p.map(function(j,v){return f.a.createElement("p",{key:v},j)}))))},N=h.default=f.a.memo(k)},xDdU:function(u,h,e){var c=e("4fRq"),f=e("I2ZF"),t,o,n=0,a=0;function C(i,I,M){var s=I&&M||0,l=I||[];i=i||{};var L=i.node||t,g=i.clockseq!==void 0?i.clockseq:o;if(L==null||g==null){var x=c();L==null&&(L=t=[x[0]|1,x[1],x[2],x[3],x[4],x[5]]),g==null&&(g=o=(x[6]<<8|x[7])&16383)}var r=i.msecs!==void 0?i.msecs:new Date().getTime(),y=i.nsecs!==void 0?i.nsecs:a+1,T=r-n+(y-a)/1e4;if(T<0&&i.clockseq===void 0&&(g=g+1&16383),(T<0||r>n)&&i.nsecs===void 0&&(y=0),y>=1e4)throw new Error("uuid.v1(): Can't create more than 10M uuids/sec");n=r,a=y,o=g,r+=122192928e5;var S=((r&268435455)*1e4+y)%4294967296;l[s++]=S>>>24&255,l[s++]=S>>>16&255,l[s++]=S>>>8&255,l[s++]=S&255;var O=r/4294967296*1e4&268435455;l[s++]=O>>>8&255,l[s++]=O&255,l[s++]=O>>>24&15|16,l[s++]=O>>>16&255,l[s++]=g>>>8|128,l[s++]=g&255;for(var E=0;E<6;++E)l[s+E]=L[E];return I||f(l)}u.exports=C}}]);