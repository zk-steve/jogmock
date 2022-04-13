import F from"svelte-preprocess";import*as o from"path";import*as g from"fs/promises";import*as i from"path";import*as x from"fs/promises";var I=(s,...t)=>c.merging(...Object.getOwnPropertyNames(s).map(e=>c.alias({alias:e,directory:i.resolve(s[e])},...t)));var R=s=>{let t=I(s,c.sass,c.pkgJson);return function(e,r,n){Promise.resolve(t(e)).then(l=>{if(l!==void 0&&l.length>0){n({file:l[0]});return}n(null)})}};function P(s,t,e){if(!s.startsWith("json:")){e(null);return}s=s.slice(5),x.readFile(i.join(i.dirname(t),`${s}.json`)).then(r=>{let n=JSON.parse(r.toString()),l=p=>p.replace(/([A-Z]+)/g," $1").split(" ").filter(a=>a!=="").map(a=>a.toLowerCase()).join("-"),u="",b=(p,a)=>{for(let m in a){let d=`${p}${p&&"-"}${l(m)}`,f=a[m];typeof f=="string"?u+=`$${d}: ${a[m]};
`:f instanceof Array?u+=`$${d}: ${f.join(", ")};
`:b(d,f)}};b(l(i.parse(s).name),n),e({contents:u})})}function L(s){return[s]}async function y(s,t){return g.stat(s).then(e=>t&&t(e)||!t?[s]:void 0).catch(()=>!1)}async function v(s){return y(s,t=>t.isFile())}async function O(s){return y(s,t=>t.isDirectory())}function h(...s){return async function(t){let e=[];for(let r of s.map(n=>`${t}.${n}`))await v(r)&&e.push(r);return e.length>0?e:void 0}}var j=h("js","cjs","mjs"),$=h("sass","scss","css");async function E(s){if(/\.(css|scss|sass)$/.test(s))return[s];let t=[s,`${s}/index`,`${s}/_index`];o.basename(s).startsWith("_")||t.push(`${o.dirname(s)}/_${o.basename(s)}`);let e=[];for(let r of t){let n=await $(r);n&&e.push(...n)}return e}async function D(s){let t=o.join(s,"package.json");try{if(await v(t)){let e=JSON.parse((await g.readFile(t)).toString()),r=e.main?o.join(s,e.main):void 0;if(r)return[r]}}catch(e){console.error(`Error resolving module package.json for path ${s}: ${e instanceof Error?e.message:e}`)}}async function J(s){let t=await j(o.join(s,"index"));if(t)return t}function w(...s){return async t=>{let e=[];for(let r of s)await Promise.resolve(r(t)).then(n=>{n!==void 0&&e.push(...n)});return e}}function S({alias:s,directory:t},...e){let r=w(...e);return async n=>{if(n.startsWith(s)&&n.length>s.length)return r(o.join(t,n.slice(s.length)))}}var c={identity:L,existing:y,existingFile:v,existingDir:O,extension:h,jsExtension:j,sassExtension:$,merging:w,sass:E,pkgJson:D,jsIndex:J,alias:S};function A(s){return s===void 0?process.env.NODE_ENV==="production":s==="production"}function V(s){return!A(s)}var k=V(),T=A();var N={preprocess:F({sourceMap:k,scss:{importer:[R({"@":"./src","~":"./node_modules"}),P],prependData:'@use "@styles/variables" as *;'}})},z=N;export{z as default};
