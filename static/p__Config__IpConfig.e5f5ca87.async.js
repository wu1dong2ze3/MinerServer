(window.webpackJsonp=window.webpackJsonp||[]).push([[8],{"5dcc":function(A,i,e){"use strict";e.d(i,"c",function(){return h}),e.d(i,"d",function(){return N}),e.d(i,"g",function(){return M}),e.d(i,"e",function(){return c}),e.d(i,"a",function(){return r}),e.d(i,"b",function(){return a}),e.d(i,"f",function(){return m});var f=e("jrin"),b=e("9og8"),s=e("WmNS"),d=e.n(s),j=e("MZJn"),E=e("fe1z"),T=e("p46w"),y=e.n(T),v=e("9kvl");function h(t){return t===""||t===void 0||t===null}function N(t){return t===0?j.a:t}function M(t,l){return h(t)?l:t}function I(t,l){var n=(t-l)/1e3/3600/24,D=Math.floor(n),p=(n-D)*24,C=Math.floor(p),B=(p-C)*60,R=Math.floor(B),F=(B-R)*60,U=Math.floor(F),O=D+"\u5929"+C+"\u5C0F\u65F6"+R+"\u5206";return O}function c(){return o.apply(this,arguments)}function o(){return o=Object(b.a)(d.a.mark(function t(){var l,n;return d.a.wrap(function(p){for(;;)switch(p.prev=p.next){case 0:return p.next=2,Object(E.c)();case 2:l=p.sent,n=l.code,n==200&&(v.e.push("/user/login"),localStorage.getItem("address")&&localStorage.removeItem("address"),y.a.remove("token"));case 5:case"end":return p.stop()}},t)})),o.apply(this,arguments)}function r(t){var l=_.trim(t);return l=l.replace(/[^\d^\.]+/g,"").replace(".","$#$").replace(/\./g,"").replace("$#$","."),l}function a(t){var l=_.trim(t);return l=l.replace(/[^\x00-\xff]/g,""),l}function m(t,l){for(var n in t)l.setFieldsValue(Object(f.a)({},n,t[n]))}},"6Ynu":function(A,i,e){"use strict";var f=e("k1fw"),b=e("q1tI"),s=e.n(b),d=function(E){return s.a.createElement("div",{style:Object(f.a)({height:"100%",background:"#fff",borderRadius:"7px"},E.style)},s.a.createElement("div",{style:{height:"54px",lineHeight:"54px",borderBottom:"1px solid #F5F7FD",fontWeight:"600",paddingLeft:"30px",color:"#333",fontSize:"18px"}},E.title," "),s.a.createElement("div",{style:{padding:"0px 30px",marginTop:"24px",color:"#333"}},E.children))};i.a=d},BdNu:function(A,i,e){A.exports={submitButton:"submitButton___17kn1"}},MZJn:function(A,i,e){"use strict";e.d(i,"a",function(){return f});var f="--"},fe1z:function(A,i,e){"use strict";e.d(i,"a",function(){return j}),e.d(i,"b",function(){return T}),e.d(i,"c",function(){return v});var f=e("9og8"),b=e("WmNS"),s=e.n(b),d=e("t3Un");function j(){return E.apply(this,arguments)}function E(){return E=Object(f.a)(s.a.mark(function I(){var c,o=arguments;return s.a.wrap(function(a){for(;;)switch(a.prev=a.next){case 0:return c=o.length>0&&o[0]!==void 0?o[0]:{},a.abrupt("return",Object(d.a)("host",{method:"GET",params:{},body:JSON.stringify(c)},Object({NODE_ENV:"production"}).getAPI));case 2:case"end":return a.stop()}},I)})),E.apply(this,arguments)}function T(){return y.apply(this,arguments)}function y(){return y=Object(f.a)(s.a.mark(function I(){var c,o=arguments;return s.a.wrap(function(a){for(;;)switch(a.prev=a.next){case 0:return c=o.length>0&&o[0]!==void 0?o[0]:{},a.abrupt("return",Object(d.a)("user/login",{method:"POST",params:{},body:JSON.stringify(c)},Object({NODE_ENV:"production"}).API));case 2:case"end":return a.stop()}},I)})),y.apply(this,arguments)}function v(){return h.apply(this,arguments)}function h(){return h=Object(f.a)(s.a.mark(function I(){var c,o=arguments;return s.a.wrap(function(a){for(;;)switch(a.prev=a.next){case 0:return c=o.length>0&&o[0]!==void 0?o[0]:{},a.abrupt("return",Object(d.a)("user/exit",{method:"GET",params:{},body:JSON.stringify(c)},Object({NODE_ENV:"production"}).API));case 2:case"end":return a.stop()}},I)})),h.apply(this,arguments)}function N(){return M.apply(this,arguments)}function M(){return M=Object(f.a)(s.a.mark(function I(){var c,o=arguments;return s.a.wrap(function(a){for(;;)switch(a.prev=a.next){case 0:return c=o.length>0&&o[0]!==void 0?o[0]:{},a.abrupt("return",Object(d.a)("miner/summary",{method:"GET",params:{},body:JSON.stringify(c)},Object({NODE_ENV:"production"}).API));case 2:case"end":return a.stop()}},I)})),M.apply(this,arguments)}},hy7T:function(A,i,e){"use strict";e.d(i,"e",function(){return j}),e.d(i,"f",function(){return T}),e.d(i,"c",function(){return v}),e.d(i,"d",function(){return N}),e.d(i,"a",function(){return I}),e.d(i,"b",function(){return o});var f=e("9og8"),b=e("WmNS"),s=e.n(b),d=e("t3Un");function j(){return E.apply(this,arguments)}function E(){return E=Object(f.a)(s.a.mark(function a(){var m,t=arguments;return s.a.wrap(function(n){for(;;)switch(n.prev=n.next){case 0:return m=t.length>0&&t[0]!==void 0?t[0]:{},n.abrupt("return",Object(d.a)("miner/user/info",{method:"GET",params:{},body:JSON.stringify(m)},Object({NODE_ENV:"production"}).API));case 2:case"end":return n.stop()}},a)})),E.apply(this,arguments)}function T(){return y.apply(this,arguments)}function y(){return y=Object(f.a)(s.a.mark(function a(){var m,t=arguments;return s.a.wrap(function(n){for(;;)switch(n.prev=n.next){case 0:return m=t.length>0&&t[0]!==void 0?t[0]:{},n.abrupt("return",Object(d.a)("miner/user/update",{method:"POST",params:{},body:JSON.stringify(m)},Object({NODE_ENV:"production"}).API));case 2:case"end":return n.stop()}},a)})),y.apply(this,arguments)}function v(){return h.apply(this,arguments)}function h(){return h=Object(f.a)(s.a.mark(function a(){var m,t=arguments;return s.a.wrap(function(n){for(;;)switch(n.prev=n.next){case 0:return m=t.length>0&&t[0]!==void 0?t[0]:{},n.abrupt("return",Object(d.a)("system/net/info",{method:"GET",params:{},body:JSON.stringify(m)},Object({NODE_ENV:"production"}).API));case 2:case"end":return n.stop()}},a)})),h.apply(this,arguments)}function N(){return M.apply(this,arguments)}function M(){return M=Object(f.a)(s.a.mark(function a(){var m,t=arguments;return s.a.wrap(function(n){for(;;)switch(n.prev=n.next){case 0:return m=t.length>0&&t[0]!==void 0?t[0]:{},n.abrupt("return",Object(d.a)("system/net/update",{method:"POST",params:{},body:JSON.stringify(m)},Object({NODE_ENV:"production"}).API));case 2:case"end":return n.stop()}},a)})),M.apply(this,arguments)}function I(){return c.apply(this,arguments)}function c(){return c=Object(f.a)(s.a.mark(function a(){var m,t=arguments;return s.a.wrap(function(n){for(;;)switch(n.prev=n.next){case 0:return m=t.length>0&&t[0]!==void 0?t[0]:{},n.abrupt("return",Object(d.a)("miner/mode/info",{method:"GET",params:{},body:JSON.stringify(m)},Object({NODE_ENV:"production"}).API));case 2:case"end":return n.stop()}},a)})),c.apply(this,arguments)}function o(){return r.apply(this,arguments)}function r(){return r=Object(f.a)(s.a.mark(function a(){var m,t=arguments;return s.a.wrap(function(n){for(;;)switch(n.prev=n.next){case 0:return m=t.length>0&&t[0]!==void 0?t[0]:{},n.abrupt("return",Object(d.a)("miner/mode/update",{method:"POST",params:{},body:JSON.stringify(m)},Object({NODE_ENV:"production"}).API));case 2:case"end":return n.stop()}},a)})),r.apply(this,arguments)}},kHRW:function(A,i,e){"use strict";e.r(i);var f=e("OaEy"),b=e("2fM7"),s=e("5NDa"),d=e("5rEg"),j=e("miYZ"),E=e("tsqr"),T=e("9og8"),y=e("jrin"),v=e("k1fw"),h=e("tJVT"),N=e("y8nQ"),M=e("Vl3Y"),I=e("WmNS"),c=e.n(I),o=e("q1tI"),r=e.n(o),a=e("6Ynu"),m=e("TSYQ"),t=e.n(m),l=e("ww3E"),n=e("9kvl"),D=e("5dcc"),p=e("hy7T"),C=M.a.Item,B=function(F){var U=Object(n.i)(),O=U.formatMessage,H=Object(o.useState)({mac:"",routerType:0,ip:"",subnetMask:"",gateway:"",dns1:"",dns2:""}),J=Object(h.a)(H,2),W=J[0],G=J[1],Y=M.a.useForm(),z=Object(h.a)(Y,1),$=z[0],Z=W.mac,Q=W.routerType,X=W.ip,x=W.subnetMask,k=W.gateway,w=W.dns1,q=W.dns2,K=function(u,L){G(Object(v.a)(Object(v.a)({},W),{},Object(y.a)({},u,L)))},ee=function(){var P=Object(T.a)(c.a.mark(function u(){var L,V,S;return c.a.wrap(function(g){for(;;)switch(g.prev=g.next){case 0:return g.next=2,Object(p.c)();case 2:L=g.sent,V=L.code,S=L.data,V==200&&!_.isEmpty(S.net_info)&&(G(S.net_info),Object(D.f)(S.net_info,$));case 6:case"end":return g.stop()}},u)}));return function(){return P.apply(this,arguments)}}();Object(o.useEffect)(function(){ee()},[]);var te=function(){var P=Object(T.a)(c.a.mark(function u(){var L,V,S;return c.a.wrap(function(g){for(;;)switch(g.prev=g.next){case 0:return g.prev=0,g.next=3,$.validateFields();case 3:if(L=g.sent,!(_.isEmpty(L.dns1)&&_.isEmpty(L.dns2))){g.next=6;break}return g.abrupt("return",E.default.error(O({id:"config.IpConfig.lastOneDns"})));case 6:return g.next=8,Object(p.d)(W);case 8:V=g.sent,S=V.code,S==200&&Object(D.e)(),g.next=16;break;case 13:g.prev=13,g.t0=g.catch(0),console.log("Failed:",g.t0);case 16:case"end":return g.stop()}},u,null,[[0,13]])}));return function(){return P.apply(this,arguments)}}();return r.a.createElement("div",{style:{height:"100%"},className:t()(["input_height_44","select_height_44"])},r.a.createElement(a.a,{title:O({id:"config.IpConfig.title"})},r.a.createElement("div",{style:{width:"600px"}},r.a.createElement(M.a,{form:$},r.a.createElement(C,{label:O({id:"config.IpConfig.macAdress"}),name:"mac",initialValue:Z||void 0,rules:[{required:!0,message:O({id:"config.IpConfig.enterMacAdress"})}],getValueFromEvent:function(u){return Object(D.b)(u.target.value)}},r.a.createElement(d.a,{autoComplete:"off",placeholder:O({id:"config.IpConfig.macAdress"}),onChange:function(u){K("mac",u.target.value)}})),r.a.createElement(C,{label:O({id:"config.IpConfig.getIp"}),name:"routerType",initialValue:Q},r.a.createElement(b.a,{style:{width:"100%",height:"44px"},onChange:function(u){console.log(u),K("routerType",u)}},r.a.createElement(b.a.Option,{value:0},O({id:"config.IpConfig.staticIp"})),r.a.createElement(b.a.Option,{value:1},"DHCP"))),r.a.createElement(C,{label:O({id:"config.IpConfig.ipAdress"}),name:"ip",initialValue:X||void 0,rules:[{required:!0,message:O({id:"config.IpConfig.enterIpAdress"})}],getValueFromEvent:function(u){return Object(D.b)(u.target.value)}},r.a.createElement(d.a,{autoComplete:"off",placeholder:O({id:"config.IpConfig.ipAdress"}),onChange:function(u){K("ip",u.target.value)}})),r.a.createElement(C,{label:O({id:"config.IpConfig.subnetMask"}),name:"subnetMask",initialValue:x||void 0,rules:[{required:!0,message:O({id:"config.IpConfig.enterSubnetMask"})}],getValueFromEvent:function(u){return Object(D.b)(u.target.value)}},r.a.createElement(d.a,{autoComplete:"off",placeholder:O({id:"config.IpConfig.subnetMask"}),onChange:function(u){K("subnetMask",u.target.value)}})),r.a.createElement(C,{label:O({id:"config.IpConfig.subnetMask"}),name:"gateway",initialValue:k||void 0,rules:[{required:!0,message:O({id:"config.IpConfig.enterGateway"})}],getValueFromEvent:function(u){return Object(D.b)(u.target.value)}},r.a.createElement(d.a,{autoComplete:"off",placeholder:O({id:"config.IpConfig.subnetMask"}),onChange:function(u){K("gateway",u.target.value)}})),r.a.createElement(C,{label:"DNS1",name:"dns1",initialValue:w||void 0,getValueFromEvent:function(u){return Object(D.b)(u.target.value)}},r.a.createElement(d.a,{autoComplete:"off",placeholder:"DNS1",onChange:function(u){K("dns1",u.target.value)}})),r.a.createElement(C,{label:"DNS2",name:"dns2",initialValue:q||void 0,getValueFromEvent:function(u){return Object(D.b)(u.target.value)}},r.a.createElement(d.a,{autoComplete:"off",placeholder:"DNS2",onChange:function(u){K("dns2",u.target.value)}})),r.a.createElement(l.a,{submitButtononClick:function(){return te()},style:{marginBottom:"48px",width:"160px"},text:O({id:"config.poolConfig.save"})})))))};i.default=B},t3Un:function(A,i,e){"use strict";e.d(i,"a",function(){return c});var f=e("9og8"),b=e("jrin"),s=e("k1fw"),d=e("WmNS"),j=e.n(d),E=e("Qyje"),T=e.n(E),y=function(r,a){var m=r,t=Object.keys(a);return t.length!=0&&(m+="?"+T.a.stringify(a)),m},v=e("p46w"),h=e.n(v),N=e("9kvl"),M=e("5dcc"),I=e("Hg0r").fetch;function c(o){var r=arguments.length>1&&arguments[1]!==void 0?arguments[1]:{},a=arguments.length>2&&arguments[2]!==void 0?arguments[2]:"",m={method:"GET"},t=Object(s.a)(Object(s.a)({},m),r);t.headers=Object(s.a)(Object(b.a)({Accept:"application/json"},"accept-language",localStorage.getItem("umi_locale")||"zh-CN"),t.headers),h.a.get("token")&&(t.headers.token=h.a.get("token"));var l="";return o=="host"?l=a:l=localStorage.getItem("address")?"http://"+localStorage.getItem("address")+"/":a,o=y(o,r.params),t.method=="GET"&&delete t.body,I(l+o,t).then(function(){var n=Object(f.a)(j.a.mark(function D(p){var C,B,R;return j.a.wrap(function(U){for(;;)switch(U.prev=U.next){case 0:return C=p.status,U.next=3,p.json();case 3:if(B=U.sent,!(p.status>=200&&p.status<300)){U.next=7;break}return B.code==680&&(N.e.push("/user/login"),localStorage.getItem("address")&&localStorage.removeItem("address"),h.a.remove("token")),U.abrupt("return",B);case 7:if(!(C>=400||C>=500)){U.next=12;break}throw R=new Error,R.response=p,R.data=B,R;case 12:case"end":return U.stop()}},D)}));return function(D){return n.apply(this,arguments)}}()).catch(function(n){var D=503,p="\u94FE\u63A5\u670D\u52A1\u5668\u8D85\u65F6";throw p})}},ww3E:function(A,i,e){"use strict";var f=e("+L6B"),b=e("2/Rp"),s=e("q1tI"),d=e.n(s),j=e("BdNu"),E=e.n(j),T=function(v){return d.a.createElement(b.a,{type:"primary",loading:v.uploading?v.uploading:!1,disabled:v.disabled?v.disabled:!1,className:E.a.submitButton,onClick:function(){return v.submitButtononClick()},style:v.style},v.text)};i.a=T}}]);
