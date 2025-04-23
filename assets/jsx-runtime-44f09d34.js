import{g as l}from"./_commonjsHelpers-39b5b250.js";import{a as u}from"./alpha_mf_2_uhmi__loadShare__react__loadShare__-56e9ba3c.js";function m(o,r){for(var _=0;_<r.length;_++){const e=r[_];if(typeof e!="string"&&!Array.isArray(e)){for(const t in e)if(t!=="default"&&!(t in o)){const n=Object.getOwnPropertyDescriptor(e,t);n&&Object.defineProperty(o,t,n.get?n:{enumerable:!0,get:()=>e[t]})}}}return Object.freeze(Object.defineProperty(o,Symbol.toStringTag,{value:"Module"}))}var a={exports:{}},s={};/**
 * @license React
 * react-jsx-runtime.production.min.js
 *
 * Copyright (c) Facebook, Inc. and its affiliates.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */var c=u,y=Symbol.for("react.element"),d=Symbol.for("react.fragment"),j=Object.prototype.hasOwnProperty,x=c.__SECRET_INTERNALS_DO_NOT_USE_OR_YOU_WILL_BE_FIRED.ReactCurrentOwner,O={key:!0,ref:!0,__self:!0,__source:!0};function i(o,r,_){var e,t={},n=null,f=null;_!==void 0&&(n=""+_),r.key!==void 0&&(n=""+r.key),r.ref!==void 0&&(f=r.ref);for(e in r)j.call(r,e)&&!O.hasOwnProperty(e)&&(t[e]=r[e]);if(o&&o.defaultProps)for(e in r=o.defaultProps,r)t[e]===void 0&&(t[e]=r[e]);return{$$typeof:y,type:o,key:n,ref:f,props:t,_owner:x.current}}s.Fragment=d;s.jsx=i;s.jsxs=i;a.exports=s;var p=a.exports;const g=l(p),R=m({__proto__:null,default:g},[p]);export{R as j};
