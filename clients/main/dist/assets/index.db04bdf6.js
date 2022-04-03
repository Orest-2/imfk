import{C as e,c as t,a,u as r,b as l,r as n,d as s,o,w as i,e as m,f as d,g as u,t as p,F as f,h as c,i as g,v,j as y,n as b,k as h,l as x,m as _}from"./vendor.cd99b205.js";!function(e=".",t="__import__"){try{self[t]=new Function("u","return import(u)")}catch(a){const r=new URL(e,location),l=e=>{URL.revokeObjectURL(e.src),e.remove()};self[t]=e=>new Promise(((a,n)=>{const s=new URL(e,r);if(self[t].moduleMap[s])return a(self[t].moduleMap[s]);const o=new Blob([`import * as m from '${s}';`,`${t}.moduleMap['${s}']=m;`],{type:"text/javascript"}),i=Object.assign(document.createElement("script"),{type:"module",src:URL.createObjectURL(o),onerror(){n(new Error(`Failed to import: ${e}`)),l(i)},onload(){a(self[t].moduleMap[s]),l(i)}});document.head.appendChild(i)})),self[t].moduleMap={}}}("/assets/");const M={trimf:[.2,.5,.8],trapmf:[.2,.4,.6,.75],szmf:[.25,.75],hzmf:[.25,.75],zsigmf:[-20,.5],ssigmf:[20,.5],zlinemf:[.25,.75],slinemf:[.25,.75],hsmf:[.25,.75],ssmf:[.25,.75],gbellmf:[.2,4,.5],gaussmf:[.1,.5],conemf:[[.5,.5],[.4,.4]],pyrammf:[[.5,.5],[.4,.4]],trappyrammf:[[.5,.5],[.2,1],[.4,.4]],gsigmf:[[.5,.5],[0,10]],gbell3dmf:[[.5,.5],[.2,.2],[2.5,2.5]],gauss3dmf:[[.5,.5],[.25,.25]],hyperbolmf:[[.5,.5],[.25,.25]],ellipsmf:[[.5,.5],[.4,.4]]},k={conemf:["x0","h"],pyrammf:["x0","h"],trappyrammf:["x0","h","a"],gsigmf:["x0","a"],gbell3dmf:["x0","a","b"],gauss3dmf:["x0","a"],hyperbolmf:["x0","a"],ellipsmf:["x0","a"]},w=location.origin,T={settings:`${w}/api/v1/settings`,eval:`${w}/api/v1/mf/:type/eval`,plot:`${w}/api/v1/mf/:type/plot`,operationPlot:`${w}/api/v1/mf/operation/:operation/:type/plot`,operationEval:`${w}/api/v1/mf/operation/:operation/:type/eval`},D=()=>Math.random().toString(36).substr(2,5),S=e=>JSON.parse(JSON.stringify(e)),P={mf:null,plotParams:[1,.01],funcParams:[],plotTraces:[],evalData:[],evalDataResult:[]},B={operands:[],operation:"intersection",plotParams:[1,.01],plotTraces:[],evalData:[],evalDataResult:[]},E=t({modules:{general:{namespaced:!0,state:()=>({type:"2d",selectedMfs:{"2d":{[D()]:S(P)},"3d":{[D()]:S(P)}}}),getters:{getSelectedMfsByType:e=>(t=e.type)=>Object.entries(e.selectedMfs[t]),getSelectedMfDataByKeyAndType:e=>(t,a=e.type)=>e.selectedMfs[a][t],getSelectedMfByKeyAndType:e=>(t,a=e.type)=>{var r;return(null==(r=e.selectedMfs[a][t])?void 0:r.mf)||null},getOperandByKeyAndOperandIndxAndType:e=>(t,a,r=e.type)=>{var l;return(null==(l=e.selectedMfs[r][t])?void 0:l.operands[a])||null}},mutations:{setType(e,t){e.type=t},addMfByType(e,{k:t,data:a}){e.selectedMfs[e.type][t]=a},setMfByType(e,{k:t,v:a}){e.selectedMfs[e.type][t].mf=a,M[null==a?void 0:a.code]?e.selectedMfs[e.type][t].funcParams=S(M[a.code]):e.selectedMfs[e.type][t].funcParams=Array(null==a?void 0:a.params_count).fill(0)},setOperandByType(e,{k:t,i:a,v:r}){e.selectedMfs[e.type][t].operands[a]={mf:r,funcParams:[],evalData:[]},M[null==r?void 0:r.code]?e.selectedMfs[e.type][t].operands[a].funcParams=S(M[r.code]):e.selectedMfs[e.type][t].operands[a].funcParams=Array(null==r?void 0:r.params_count).fill(0)},setMfOpretionByType(e,{k:t,v:a}){e.selectedMfs[e.type][t].operation=a},setMfPlotTracesByType(e,{k:t,v:a}){e.selectedMfs[e.type][t].plotTraces=a},setMfEvalDataByType(e,{k:t,v:a}){e.selectedMfs[e.type][t].evalData=a},setMfEvalDataResultByType(e,{k:t,v:a}){e.selectedMfs[e.type][t]&&(e.selectedMfs[e.type][t].evalDataResult=a)},removeMfByTypeAndKey(e,{k:t}){delete e.selectedMfs[e.type][t]}},actions:{addMf({commit:e}){e("addMfByType",{k:D(),data:S(P)})},addMfOperation({commit:e}){e("addMfByType",{k:D(),data:S(B)})},removeMf({commit:e},t){e("removeMfByTypeAndKey",{k:t})},makePlot({state:t,commit:a,dispatch:r},{k:l,payload:n}){e.post(T.plot.replace(":type",t.type),n).then((({data:e})=>{r("error/clearErrors",null,{root:!0}),a("setMfPlotTracesByType",{k:l,v:[e.data]})})).catch((e=>{r("error/setErrors",e,{root:!0})}))},operationMakePlot({state:t,getters:a,commit:r,dispatch:l},{k:n,payload:s}){const{operation:o}=a.getSelectedMfDataByKeyAndType(n);e.post(T.operationPlot.replace(":type",t.type).replace(":operation",o),s).then((({data:e})=>{l("error/clearErrors",null,{root:!0});const t=e.data.map(((e,t)=>({x:e.x,y:e.y,name:e.membership_func_id?"function":"result",color:e.membership_func_id?"blue":"red",dash:e.membership_func_id?"dot":"solid"})));r("setMfPlotTracesByType",{k:n,v:t})})).catch((e=>{l("error/setErrors",e,{root:!0})}))},evalData({state:t,commit:a,dispatch:r},{k:l,payload:n}){e.post(T.eval.replace(":type",t.type),n).then((({data:e})=>{r("error/clearErrors",null,{root:!0}),a("setMfEvalDataResultByType",{k:l,v:[...e.data.result]})})).catch((e=>{r("error/setErrors",e,{root:!0})}))},operationEvalData({state:t,getters:a,commit:r,dispatch:l},{k:n,payload:s}){const{operation:o}=a.getSelectedMfDataByKeyAndType(n);e.post(T.operationEval.replace(":type",t.type).replace(":operation",o),s).then((({data:e})=>{l("error/clearErrors",null,{root:!0}),r("setMfEvalDataResultByType",{k:n,v:[...e.data[2].result]})})).catch((e=>{l("error/setErrors",e,{root:!0})}))}}},settings:{namespaced:!0,state:()=>({membership_funcs:null}),getters:{getMembershipFuncsByType:e=>t=>{var a;return(null==(a=e.membership_funcs)?void 0:a.filter((e=>e.type===t)))||[]}},mutations:{setMembershipFuncs(e,t){e.membership_funcs=t}},actions:{fetchSettings({commit:t}){e.get(T.settings).then((({data:e})=>{t("setMembershipFuncs",e.data.membership_funcs)}))}}},error:{namespaced:!0,state:()=>({list:[]}),mutations:{setErrorList(e,t){e.list=t}},actions:{setErrors({commit:e},t){var a;if(null==t?void 0:t.response){const{error:r}=(null==(a=null==t?void 0:t.response)?void 0:a.data)||{error:"Internal server error"};e("setErrorList",[{detail:r}])}else e("setErrorList",[{detail:t}])},clearErrors({commit:e}){e("setErrorList",[])}}}}}),z=a({locale:"en",fallbackLocale:"en",messages:{en:{general:{membership_functions:"Membership functions",one_variable:"One variable",many_variables:"Many variables",select_the_membership_function:"Select the membership function",add_membership_function:"+ Add membership function",add_operation:"+ Add operation",graph_parameters:"Graph parameters",enter_the_information_you_want_to_formalize:"Enter the information you want to formalize",evaluate:"Evaluate",select_an_operation:"Select an operation",result:"Result",function:"Function",empty:"Empty",number_of_data_sets:"Number of data sets",number_of_arguments:"Number of arguments",function_parameters:"Function parameters"},operations:{intersection:"Intersection",association:"Association",difference:"Difference",symmetrical_difference:"Symmetrical difference",disjunctive_sum:"Disjunctive sum"},membership_functions:{trimf:"Triangular membership function",trapmf:"Trapezoidal membership function",szmf:"Square Z-spline",hzmf:"Harmonic Z-spline",zsigmf:"Z-sigmoidal membership function",zlinemf:"Z-linear membership function",ssmf:"Square S-spline",hsmf:"Harmonic S-spline",ssigmf:"S-sigmoidal membership function",slinemf:"S-linear membership function",gbellmf:"Bell-like membership function",gaussmf:"Gaussian membership function",conemf:"Cone-like membership function",pyrammf:"Pyramid membership function",trappyrammf:"Trapezoidal-pyramidal membership function",gsigmf:"Generalized sigmoidal membership function",gbell3dmf:"Bell-like membership function",gauss3dmf:"Gaussian membership function",hyperbolmf:"Hyperboloid membership function",ellipsmf:"Ellipsoid membership function"}},ua:{general:{membership_functions:"Функції належності",one_variable:"Однієї змінної",many_variables:"Багатьох змінних",select_the_membership_function:"Виберіть функцію належності",add_membership_function:"+ Добавити функцію належності",add_operation:"+ Добавити операцію",graph_parameters:"Параметри графіка",enter_the_information_you_want_to_formalize:"Введіть інформацію, що потрібно формалізувати",evaluate:"Виконати",select_an_operation:"Виберіть операцію",result:"Результат",function:"Функція",empty:"Нічого не вибрано",number_of_data_sets:"Кількість наборів даних",number_of_arguments:"Кількість аргументів",function_parameters:"Параметри функції"},operations:{intersection:"Перетин",association:"Об'єднання",difference:"Різниця",symmetrical_difference:"Симетрична різниця",disjunctive_sum:"Диз'юнктивна сума"},membership_functions:{trimf:"Трикутна функції належності",trapmf:"Трапециідальна функція належності",szmf:"Квадратичний Z-сплайн",hzmf:"Гармонійний Z-сплайн",zsigmf:"Z-сигмоїдальна функція належності",zlinemf:"Z-лінійна функція належності",ssmf:"Квадратичний S-сплайн",hsmf:"Гармонійний S-сплайн",ssigmf:"S-сигмоїдальна функція належності",slinemf:"S-лінійна функція належності",gbellmf:"Дзвоноподібна функція належності",gaussmf:"Гаусова функція належності",conemf:"Конусоподібна функція належності",pyrammf:"Пірамідальна функція належності",trappyrammf:"Трапецієподібно-пірамідальна функція належності",gsigmf:"Узагальнена сигмоїдальна функція належності",gbell3dmf:"Дзвоноподібна функція належності",gauss3dmf:"Гаусова функція належності",hyperbolmf:"Гіперболоїдна функція належності",ellipsmf:"Еліпсоїдна функція належності"}},hu:{general:{membership_functions:"Tagsági függvények",one_variable:"Egyváltozós",many_variables:"Többváltozós",select_the_membership_function:"Válasszon tagsági függvényt",add_membership_function:"+ Tagsági függvény hozzáadása",add_operation:"+ Művelet hozzáadása",graph_parameters:"Grafikon paraméterek",enter_the_information_you_want_to_formalize:"Adja meg a formalizálni kívánt információkat",evaluate:"Elvégzés",select_an_operation:"Válasszon műveletet",result:"Eredmény",function:"Függvény",empty:"Üres kijelölés",number_of_data_sets:"Adatkészletek száma",number_of_arguments:"Paraméterek száma",function_parameters:"Függvény paraméterei"},operations:{intersection:"Metszés",association:"Unió",difference:"Különbség",symmetrical_difference:"Szimmetrikus különbség",disjunctive_sum:"Diszjunktív összeg"},membership_functions:{trimf:"Háromszög tagsági függvény",trapmf:"Trapéz tagsági függvény",szmf:"Másodfokú Z-spline",hzmf:"Harmonikus Z-spline",zsigmf:"Z-szigmoid tagsági függvény",zlinemf:"Z-lineáris tagsági függvény",ssmf:"Másodfokú S-spline",hsmf:"Harmonikus S-spline",ssigmf:"S-szigmoid tagsági függvény",slinemf:"S-lineáris tagsági függvény",gbellmf:"Harangtípusú tagsági függvény",gaussmf:"Gauss tagsági függvény",conemf:"Kúpos tagsági függvény",pyrammf:"Piramis tagsági függvény",trappyrammf:"Trapéz alakú piramis tagsági függvény",gsigmf:"Általánosított szigmoid tagsági függvény",gbell3dmf:"Harangtípusú tagsági függvény",gauss3dmf:"Gauss tagsági függvény",hyperbolmf:"Hiperboloid tagsági függvény",ellipsmf:"Ellipszoid tagsági függvény"}}}}),$=()=>{let e=null;return(t,a)=>{clearTimeout(e),e=setTimeout((()=>{t()}),a||500)}},A={props:{mfid:{type:String,default:""}},setup(e){const{t:t}=r({useScope:"global"}),a=l(),m=n(null),d=s((()=>a.state.general.type)),u=s((()=>a.getters["general/getSelectedMfDataByKeyAndType"](e.mfid))),p=s((()=>{var e;return(null==(e=u.value)?void 0:e.plotTraces)||[]})),f=s((()=>p.value.map((e=>{const a={x:e.x,y:e.y};return e.name&&(a.name=t(`general.${e.name}`)),"3d"===d.value?(a.z=e.z||[],a.type="surface"):(a.line={color:e.color||"red",width:2},e.dash&&(a.line.dash=e.dash)),a})))),c={dragmode:"2d"===d.value?"pan":"turntable"},g={scrollZoom:!0,responsive:!0};return o((()=>{Plotly.newPlot(m.value,f.value,c,g)})),i(f,(()=>{Plotly.newPlot(m.value,f.value,c,g)})),{plotEl:m}}},R={ref:"plotEl",class:"lg:w-1/2 md:h-150 h-100 w-full"};A.render=function(e,t,a,r,l,n){return m(),d("div",R,null,512)};const U={props:{mfid:{type:String,default:""},operand:{type:Number,default:-1}},setup(e){const t=l(),a=s((()=>t.getters["general/getSelectedMfDataByKeyAndType"](e.mfid))),r=s((()=>{var e;return(null==(e=a.value)?void 0:e.operands)||[]}));return{selectedMf:s((()=>{var t;return null==(t=r.value.length&&r.value[e.operand]||a.value)?void 0:t.mf})),params:s((()=>{var t;return null==(t=r.value.length&&r.value[e.operand]||a.value)?void 0:t.funcParams}))}}},O={class:"mb-2"},C={class:"font-mono mb-1"},K={class:"list-lowerlatin list-inside"},j={class:"flex flex-wrap space md:-mx-1"},V=u("li",{class:"mr-1"},null,-1);U.render=function(e,t,a,r,l,n){return m(),d("fieldset",O,[u("legend",C,p(e.$t("general.function_parameters"))+" "+p(a.operand>=0&&a.operand+1||""),1),u("ul",K,[u("div",j,[(m(!0),d(f,null,c(r.selectedMf.params_count,((e,t)=>(m(),d("div",{key:e+t,class:"flex items-center w-full md:w-1/8 md:mx-1 md:mb-0 mb-1"},[V,g(u("input",{"onUpdate:modelValue":e=>r.params[t]=e,type:"number",step:"0.2",class:"border-3 border-gray-500 rounded w-full p-1"},null,8,["onUpdate:modelValue"]),[[v,r.params[t],void 0,{number:!0}]])])))),128))])])])};const L={props:{mfid:{type:String,default:""},operand:{type:Number,default:-1}},setup(e){const t=l(),a=s((()=>t.getters["settings/getMembershipFuncsByType"](t.state.general.type)));return{selectedMf:s({get(){var a;return e.operand>=0?(null==(a=t.getters["general/getOperandByKeyAndOperandIndxAndType"](e.mfid,e.operand))?void 0:a.mf)||null:t.getters["general/getSelectedMfByKeyAndType"](e.mfid)},set(a){e.operand>=0?t.commit("general/setOperandByType",{k:e.mfid,i:e.operand,v:a}):t.commit("general/setMfByType",{k:e.mfid,v:a})}}),membershipFuncs:a}}},F={class:"w-full"},Z={class:"font-mono mb-1"},H={value:null,selected:""};L.render=function(e,t,a,r,l,n){return m(),d("fieldset",F,[u("legend",Z,p(e.$t("general.select_the_membership_function")),1),g(u("select",{"onUpdate:modelValue":t[1]||(t[1]=e=>r.selectedMf=e),class:"border-3 border-gray-500 rounded w-full p-1.2"},[u("option",H,p(e.$t("general.empty")),1),(m(!0),d(f,null,c(r.membershipFuncs,(t=>(m(),d("option",{key:t.id,value:t},p(e.$t(`membership_functions.${t.code}`)),9,["value"])))),128))],512),[[y,r.selectedMf]])])};const G={setup(){const e=l(),t=s((()=>e.state.error.list));return{store:e,errorList:t}}};G.render=function(e,t,a,r,l,n){return m(),d("div",null,[(m(!0),d(f,null,c(r.errorList,((e,t)=>(m(),d("div",{key:t,class:"border-3 p-4 rounded border-red-500 bg-red-200"},p(e.detail),1)))),128))])};const N={props:{mfid:{type:String,default:""}},setup(e){const t=l(),a=s((()=>t.getters["general/getSelectedMfDataByKeyAndType"](e.mfid))),r=s((()=>{var e;return null==(e=a.value)?void 0:e.mf})),n=s((()=>{var e;return null==(e=a.value)?void 0:e.funcParams}));return{selectedMf:r,paramsLetters:k,params:n}}},I={class:"mb-2"},q={class:"font-mono mb-1"},J={class:"list-lowerlatin list-inside"},X={class:"flex flex-wrap space md:-mx-1"},Q={class:"min-w-4 mr-1"};N.render=function(e,t,a,r,l,n){return m(),d("fieldset",I,[u("legend",q,p(e.$t("general.function_parameters")),1),u("ul",J,[u("div",X,[(m(!0),d(f,null,c(r.params.length,((e,t)=>(m(),d("div",{key:e+t,class:"flex items-center w-full md:mx-1 mb-1"},[u("span",Q,p(r.paramsLetters[r.selectedMf.code][t]),1),(m(!0),d(f,null,c(r.params[t],((e,a)=>g((m(),d("input",{key:`${t}+${a}`,"onUpdate:modelValue":e=>r.params[t][a]=e,type:"number",step:"1",class:"border-3 border-gray-500 rounded md:w-1/4 p-1 mx-1"},null,8,["onUpdate:modelValue"])),[[v,r.params[t][a],void 0,{number:!0}]]))),128))])))),128))])])])};const W={get params(){return new URLSearchParams(location.search)}},Y={props:{mfid:{type:String,default:""}},emits:["eval"],setup(e){const t=l(),a=s((()=>t.getters["general/getSelectedMfDataByKeyAndType"](e.mfid))),r=s((()=>{var e;return null==(e=a.value)?void 0:e.evalData})),n=s({get:()=>r.value.length,set(e){const t=r.value.length;t<e&&r.value.push(...Array(e-t).fill(0)),t>e&&(r.value=r.value.slice(0,e))}});return o((()=>{"1"===W.params.get("test")?t.commit("general/setMfEvalDataByType",{k:e.mfid,v:[.415,.35,.613,.283,.927]}):t.commit("general/setMfEvalDataByType",{k:e.mfid,v:[0]})})),{data:r,count:n}}},ee={class:"mb-2"},te={class:"font-mono mb-1"},ae={class:"flex flex-wrap space md:-mx-1"},re={class:"flex items-center w-full my-1"},le={class:"flex flex-wrap space mx-1 md:w-1/1 text-center"},ne={class:"flex items-center w-full"},se={class:"flex py-2"};Y.render=function(e,t,a,r,l,n){return m(),d("fieldset",ee,[u("legend",te,p(e.$t("general.enter_the_information_you_want_to_formalize")),1),u("div",ae,[(m(!0),d(f,null,c(r.data,((e,t)=>(m(),d("div",{key:t,class:"flex items-center w-full md:w-1/8 md:mx-1 mb-1"},[g(u("input",{"onUpdate:modelValue":e=>r.data[t]=e,type:"number",class:"border-3 border-gray-500 rounded w-full p-1"},null,8,["onUpdate:modelValue"]),[[v,r.data[t],void 0,{number:!0}]])])))),128))]),u("div",re,[u("button",{class:"mx-auto px-15px py-5px border-3 border-gray-500 rounded w-full font-mono hover:bg-true-gray-100 bg-true-gray-200",onClick:t[1]||(t[1]=e=>r.data.pop())}," - "),u("div",le,[u("div",ne,[g(u("input",{"onUpdate:modelValue":t[2]||(t[2]=e=>r.count=e),type:"number",class:"border-3 border-gray-500 rounded w-full p-1"},null,512),[[v,r.count,void 0,{number:!0}]])])]),u("button",{class:"mx-auto px-15px py-5px border-3 border-gray-500 rounded w-full font-mono hover:bg-true-gray-100 bg-true-gray-200",onClick:t[3]||(t[3]=e=>r.data.push(0))}," + ")]),u("div",se,[u("button",{class:"mx-auto px-15px py-5px border-3 border-gray-500 rounded w-full font-mono hover:bg-true-gray-100 bg-true-gray-200",onClick:t[4]||(t[4]=t=>e.$emit("eval"))},p(e.$t("general.evaluate")),1)])])};const oe={props:{mfid:{type:String,default:""}},emits:["eval"],setup(e){const t=l(),a=s((()=>t.state.general.type)),r=s((()=>t.getters["general/getSelectedMfDataByKeyAndType"](e.mfid))),n=s((()=>{var e;return null==(e=r.value)?void 0:e.mf})),m=s((()=>{var e;return null==(e=r.value)?void 0:e.evalData})),d=s({get(){var e;return(null==(e=m.value[0])?void 0:e.length)||2},set(e){m.value.forEach(((t,a)=>{const r=t.length||2;r<e&&t.push(...Array(e-r).fill(0)),r>e&&(m.value[a]=t.slice(0,e))}))}}),u=s({get:()=>m.value.length||1,set(a){const r=m.value.length||1;r<a&&m.value.push(Array(d.value).fill(0)),r>a&&t.commit("general/setMfEvalDataByType",{k:e.mfid,v:m.value.slice(0,a)})}}),p=()=>{var e;if("3d"===a.value&&m.value){const t=null==(e=r.value)?void 0:e.funcParams;b((()=>{const e=m.value.length||1;t.forEach(((a,r)=>{const l=a.length||2;l<e&&(t[r]=[...t[r],...Array(e-l).fill(0)]),l>e&&(t[r]=a.slice(0,e))}))}))}};return i(n,p,{deep:!0}),i(m,p,{deep:!0}),o((()=>{"1"===W.params.get("test")?t.commit("general/setMfEvalDataByType",{k:e.mfid,v:[[.87,.82,.6,.77,.69],[.66,.83,.71,.98,.91],[.78,.4,.54,.85,.82]]}):t.commit("general/setMfEvalDataByType",{k:e.mfid,v:[[0,0],[0,0]]})})),{data:m,countCol:d,countRow:u}}},ie={class:"mb-2"},me={class:"font-mono mb-1"},de={class:"flex flex-wrap space md:-mx-1"},ue={class:"mb-1"},pe={class:"font-mono mb-1"},fe={class:"flex items-center w-full mb-1"},ce={class:"flex flex-wrap space mx-1 md:w-1/1 text-center"},ge={class:"flex items-center w-full"},ve={class:"mb-1"},ye={class:"font-mono mb-1"},be={class:"flex items-center w-full mb-1"},he={class:"flex flex-wrap space mx-1 md:w-1/1 text-center"},xe={class:"flex items-center w-full"},_e={class:"flex py-2"};oe.render=function(e,t,a,r,l,n){return m(),d("fieldset",ie,[u("legend",me,p(e.$t("general.enter_the_information_you_want_to_formalize")),1),u("div",de,[(m(!0),d(f,null,c(r.data,((e,t)=>(m(),d("div",{key:t,class:"flex flex-wrap items-center w-full md:mx-1 mb-1"},[(m(!0),d(f,null,c(r.data[t],((e,a)=>(m(),d("div",{key:`${t}+${a}`,class:"flex items-center w-full md:w-1/8 md:mx-1 mb-1"},[g(u("input",{"onUpdate:modelValue":e=>r.data[t][a]=e,type:"number",class:"border-3 border-gray-500 rounded w-full p-1"},null,8,["onUpdate:modelValue"]),[[v,r.data[t][a],void 0,{number:!0}]])])))),128))])))),128))]),u("fieldset",ue,[u("legend",pe,p(e.$t("general.number_of_data_sets")),1),u("div",fe,[u("button",{class:"mx-auto px-15px py-5px border-3 border-gray-500 rounded w-full font-mono hover:bg-true-gray-100 bg-true-gray-200",onClick:t[1]||(t[1]=e=>r.countCol--)}," - "),u("div",ce,[u("div",ge,[g(u("input",{"onUpdate:modelValue":t[2]||(t[2]=e=>r.countCol=e),type:"number",class:"border-3 border-gray-500 rounded w-full p-1"},null,512),[[v,r.countCol,void 0,{number:!0}]])])]),u("button",{class:"mx-auto px-15px py-5px border-3 border-gray-500 rounded w-full font-mono hover:bg-true-gray-100 bg-true-gray-200",onClick:t[3]||(t[3]=e=>r.countCol++)}," + ")])]),u("fieldset",ve,[u("legend",ye,p(e.$t("general.number_of_arguments")),1),u("div",be,[u("button",{class:"mx-auto px-15px py-5px border-3 border-gray-500 rounded w-full font-mono hover:bg-true-gray-100 bg-true-gray-200",onClick:t[4]||(t[4]=e=>r.countRow--)}," - "),u("div",he,[u("div",xe,[g(u("input",{"onUpdate:modelValue":t[5]||(t[5]=e=>r.countRow=e),type:"number",class:"border-3 border-gray-500 rounded w-full p-1"},null,512),[[v,r.countRow,void 0,{number:!0}]])])]),u("button",{class:"mx-auto px-15px py-5px border-3 border-gray-500 rounded w-full font-mono hover:bg-true-gray-100 bg-true-gray-200",onClick:t[6]||(t[6]=e=>r.countRow++)}," + ")])]),u("div",_e,[u("button",{class:"mx-auto px-15px py-5px border-3 border-gray-500 rounded w-full font-mono hover:bg-true-gray-100 bg-true-gray-200",onClick:t[7]||(t[7]=t=>e.$emit("eval"))},p(e.$t("general.evaluate")),1)])])};const Me={props:{mfid:{type:String,default:""}},setup(e){const t=l(),a=s((()=>t.getters["general/getSelectedMfDataByKeyAndType"](e.mfid))),r=s((()=>{var e;return null==(e=a.value)?void 0:e.evalDataResult.map((e=>e.toFixed(4)))}));return i([()=>{var e;return null==(e=a.value)?void 0:e.funcParams},()=>{var e;return null==(e=a.value)?void 0:e.plotParams},()=>{var e;return null==(e=a.value)?void 0:e.mf},()=>{var e;return null==(e=a.value)?void 0:e.evalData}],(()=>{t.commit("general/setMfEvalDataResultByType",{k:e.mfid,v:[]})}),{deep:!0}),{resf:r}}},ke={key:0,class:"w-full text-center"},we={class:"text-3xl mb-3 font-mono"},Te={class:"text-xl"};Me.render=function(e,t,a,r,l,n){return r.resf&&r.resf.length?(m(),d("div",ke,[u("div",we,p(e.$t("general.result")),1),u("div",Te,p(`${r.resf.join(" | ")}`),1)])):h("",!0)};const De={props:{mfid:{type:String,default:""}},setup(e){const t=l(),a=s((()=>t.getters["general/getSelectedMfDataByKeyAndType"](e.mfid)));return{plotParams:s((()=>{var e;return null==(e=a.value)?void 0:e.plotParams}))}}},Se={key:0,class:"mb-2"},Pe={class:"font-mono mb-1"},Be={class:"flex flex-wrap space md:-mx-1"},Ee={class:"flex items-center w-full md:w-1/8 md:mx-1 md:mb-0 mb-1"},ze={class:"flex items-center w-full md:w-1/8 md:mx-1 md:mb-0 mb-1"};De.render=function(e,t,a,r,l,n){return e.hidePlotParams?h("",!0):(m(),d("fieldset",Se,[u("legend",Pe,p(e.$t("general.graph_parameters")),1),u("div",Be,[u("div",Ee,[g(u("input",{"onUpdate:modelValue":t[1]||(t[1]=e=>r.plotParams[0]=e),class:"border-3 border-gray-500 rounded w-full p-1",type:"number",step:"1"},null,512),[[v,r.plotParams[0],void 0,{number:!0}]])]),u("div",ze,[g(u("input",{"onUpdate:modelValue":t[2]||(t[2]=e=>r.plotParams[1]=e),class:"border-3 border-gray-500 rounded w-full p-1",type:"number",step:"0.1"},null,512),[[v,r.plotParams[1],void 0,{number:!0}]])])])]))};const $e={components:{Polot:A,MfParams:U,MfSelector:L,DangerAlert:G,Mf3DParams:N,MfEvalData:Y,Mf3DEvalData:oe,MfResult:Me,PolotParams:De},props:{showDelBtn:Boolean,mfid:{type:String,default:""}},emits:["remove"],setup(e){const t=l(),a=s((()=>t.state.general.type)),r=s((()=>t.getters["general/getSelectedMfDataByKeyAndType"](e.mfid))),n=s((()=>{var e;return null==(e=r.value)?void 0:e.mf})),o=s((()=>{var e;return null==(e=r.value)?void 0:e.evalData})),m=$();return i([()=>{var e;return null==(e=r.value)?void 0:e.funcParams},()=>{var e;return null==(e=r.value)?void 0:e.plotParams}],(([a,r])=>{if(n.value&&a&&r){const l={membership_func_id:n.value.id,func_params:a,plot_params:r};m((()=>t.dispatch("general/makePlot",{k:e.mfid,payload:l})),1e3)}}),{deep:!0}),{type:a,selectedMf:n,evalData:()=>{var l;const s={membership_func_id:n.value.id,func_params:null==(l=r.value)?void 0:l.funcParams,in_data:"3d"===a.value?(i=o.value,i.reduce(((e,t)=>t.map(((a,r)=>[...e[r]||[],t[r]]))),[])):o.value};var i;t.dispatch("general/evalData",{k:e.mfid,payload:s})}}}},Ae={class:"py-1rem"},Re={key:0,class:"md:w-1/10 self-end"},Ue={key:0,class:"flex flex-wrap lg:flex-row flex-col-reverse"},Oe={class:"lg:w-1/2 pl-1rem -lg:px-1rem sm:w-full"};$e.render=function(e,t,a,r,l,n){const s=x("mf-selector"),o=x("polot"),i=x("danger-alert"),p=x("polot-params"),f=x("mf-params"),c=x("mf-3-d-params"),g=x("mf-eval-data"),v=x("mf-3-d-eval-data"),y=x("mf-result");return m(),d("div",Ae,[u("div",{class:["-lg:px-1rem flex space-x-2",{"mb-1rem":r.selectedMf}]},[u(s,{mfid:a.mfid},null,8,["mfid"]),a.showDelBtn?(m(),d("div",Re,[u("button",{class:"mx-auto px-15px py-5px border-3 border-red-500 rounded w-full font-mono hover:bg-red-300 bg-red-200",onClick:t[1]||(t[1]=t=>e.$emit("remove"))}," X ")])):h("",!0)],2),r.selectedMf?(m(),d("div",Ue,[u(o,{mfid:a.mfid},null,8,["mfid"]),u("div",Oe,[u(i,{class:"mb-2"}),u(p,{mfid:a.mfid},null,8,["mfid"]),"2d"===r.type?(m(),d(f,{key:0,mfid:a.mfid},null,8,["mfid"])):h("",!0),"3d"===r.type?(m(),d(c,{key:1,mfid:a.mfid},null,8,["mfid"])):h("",!0),"2d"===r.type?(m(),d(g,{key:2,mfid:a.mfid,onEval:r.evalData},null,8,["mfid","onEval"])):h("",!0),"3d"===r.type?(m(),d(v,{key:3,mfid:a.mfid,onEval:r.evalData},null,8,["mfid","onEval"])):h("",!0)]),u(y,{mfid:a.mfid},null,8,["mfid"])])):h("",!0)])};const Ce={props:{mfid:{type:String,default:""}},setup(e){const{t:t}=r({useScope:"global"}),a=l(),n=s((()=>[{value:"intersection",text:t("operations.intersection")},{value:"association",text:t("operations.association")},{value:"difference",text:t("operations.difference")},{value:"symmetrical_difference",text:t("operations.symmetrical_difference")},{value:"disjunctive_sum",text:t("operations.disjunctive_sum")}])),o=s((()=>a.getters["general/getSelectedMfDataByKeyAndType"](e.mfid)));return{selectedMfOpretion:s({get(){var e;return(null==(e=o.value)?void 0:e.opretion)||"intersection"},set(t){a.commit("general/setMfOpretionByType",{k:e.mfid,v:t})}}),operations:n}}},Ke={class:"w-full"},je={class:"font-mono mb-1"};Ce.render=function(e,t,a,r,l,n){return m(),d("fieldset",Ke,[u("legend",je,p(e.$t("general.select_an_operation")),1),g(u("select",{"onUpdate:modelValue":t[1]||(t[1]=e=>r.selectedMfOpretion=e),class:"border-3 border-gray-500 rounded w-full p-1.2"},[(m(!0),d(f,null,c(r.operations,(e=>(m(),d("option",{key:e.value,value:e.value},p(e.text),9,["value"])))),128))],512),[[y,r.selectedMfOpretion]])])};const Ve={components:{MfSelector:L,MfOperationSelector:Ce,Polot:A,DangerAlert:G,MfParams:U,MfEvalData:Y,MfResult:Me,PolotParams:De},props:{showDelBtn:Boolean,mfid:{type:String,default:""}},emits:["remove"],setup(e){const t=l(),a=s((()=>t.state.general.type)),r=s((()=>t.getters["general/getSelectedMfDataByKeyAndType"](e.mfid))),n=s((()=>{var e;return(null==(e=r.value)?void 0:e.operands)||[]})),o=s((()=>{var e;return null==(e=r.value)?void 0:e.evalData})),m=$();return i([()=>n.value,()=>{var e;return null==(e=r.value)?void 0:e.operation},()=>{var e;return null==(e=r.value)?void 0:e.plotParams}],(([a,r,l])=>{if(a.length>1&&r&&l){const r={data:a.map((e=>({membership_func_id:e.mf.id,func_params:e.funcParams,plot_params:l})))};m((()=>t.dispatch("general/operationMakePlot",{k:e.mfid,payload:r})),1e3)}}),{deep:!0}),{type:a,operands:n,evalData:()=>{const r={data:n.value.map((e=>{return{membership_func_id:e.mf.id,func_params:e.funcParams,in_data:"3d"===a.value?(t=o.value,t.reduce(((e,t)=>t.map(((a,r)=>[...e[r]||[],t[r]]))),[])):o.value};var t}))};t.dispatch("general/operationEvalData",{k:e.mfid,payload:r})}}}},Le={class:"py-1rem"},Fe={key:0,class:"md:w-1/10 self-end"},Ze={key:0,class:"flex flex-wrap lg:flex-row flex-col-reverse"},He={class:"lg:w-1/2 pl-1rem -lg:px-1rem sm:w-full"};Ve.render=function(e,t,a,r,l,n){const s=x("mf-selector"),o=x("mf-operation-selector"),i=x("polot"),p=x("danger-alert"),f=x("PolotParams"),c=x("mf-params"),g=x("mf-eval-data"),v=x("mf-result");return m(),d("div",Le,[u("div",{class:["-lg:px-1rem flex space-x-2",{"mb-1rem":r.operands.length>1}]},[u(s,{mfid:a.mfid,operand:0},null,8,["mfid"]),u(o,{mfid:a.mfid},null,8,["mfid"]),u(s,{mfid:a.mfid,operand:1},null,8,["mfid"]),a.showDelBtn?(m(),d("div",Fe,[u("button",{class:"mx-auto px-15px py-5px border-3 border-red-500 rounded w-full font-mono hover:bg-red-300 bg-red-200",onClick:t[1]||(t[1]=t=>e.$emit("remove"))}," X ")])):h("",!0)],2),r.operands.length>1?(m(),d("div",Ze,[u(i,{mfid:a.mfid},null,8,["mfid"]),u("div",He,[u(p,{class:"mb-2"}),u(f,{mfid:a.mfid},null,8,["mfid"]),"2d"===r.type?(m(),d(c,{key:0,mfid:a.mfid,operand:0},null,8,["mfid"])):h("",!0),"2d"===r.type?(m(),d(c,{key:1,mfid:a.mfid,operand:1},null,8,["mfid"])):h("",!0),"2d"===r.type?(m(),d(g,{key:2,mfid:a.mfid,onEval:r.evalData},null,8,["mfid","onEval"])):h("",!0)]),u(v,{mfid:a.mfid},null,8,["mfid"])])):h("",!0)])};const Ge={components:{MfCard:$e,MfOperationCard:Ve},setup(){const e=l();return{type:s((()=>e.state.general.type)),mfs:s((()=>e.getters["general/getSelectedMfsByType"]())),addMf:()=>{e.dispatch("general/addMf")},addMfOperation:()=>{e.dispatch("general/addMfOperation")},removeMf:t=>{e.dispatch("general/removeMf",t)}}}},Ne={class:"divide-y-2 divide-true-gray-500 mb-2rem"},Ie={class:"flex py-15px"};Ge.render=function(e,t,a,r,l,n){const s=x("mf-card"),o=x("mf-operation-card");return m(),d("div",Ne,[(m(!0),d(f,null,c(r.mfs,(([e,t])=>(m(),d("div",{key:e},[t.operation?h("",!0):(m(),d(s,{key:0,mfid:e,"show-del-btn":r.mfs.length>1,onRemove:t=>r.removeMf(e)},null,8,["mfid","show-del-btn","onRemove"])),t.operation?(m(),d(o,{key:1,mfid:e,"show-del-btn":r.mfs.length>1,onRemove:t=>r.removeMf(e)},null,8,["mfid","show-del-btn","onRemove"])):h("",!0)])))),128)),u("div",Ie,[u("button",{class:"mx-15px px-15px py-5px border-3 border-gray-500 rounded w-full font-mono hover:bg-true-gray-100 bg-true-gray-200",onClick:t[1]||(t[1]=(...e)=>r.addMf&&r.addMf(...e))},p(e.$t("general.add_membership_function")),1),"2d"===r.type?(m(),d("button",{key:0,class:"mx-15px px-15px py-5px border-3 border-gray-500 rounded w-full font-mono hover:bg-true-gray-100 bg-true-gray-200",onClick:t[2]||(t[2]=(...e)=>r.addMfOperation&&r.addMfOperation(...e))},p(e.$t("general.add_operation")),1)):h("",!0)])])};const qe={components:{MfList:Ge},setup(){const e=l(),t=s({get:()=>e.state.general.type,set(t){e.commit("general/setType",t)}});return e.dispatch("settings/fetchSettings"),{type:t}}},Je={class:"lg:container mx-auto min-h-screen bg-white"},Xe={class:"pt-10px"},Qe={class:" text-center"},We={class:"text-3xl font-mono"},Ye=u("option",{value:"en"}," EN ",-1),et=u("option",{value:"ua"}," UA ",-1),tt=u("option",{value:"hu"}," HU ",-1),at={class:"font-mono flex justify-center divide-x-2 divide-black border-b-2 p-15px border-true-gray-500"};qe.render=function(e,t,a,r,l,n){const s=x("mf-list");return m(),d("div",null,[u("div",Je,[u("div",Xe,[u("div",Qe,[u("span",We,p(e.$t("general.membership_functions")),1),g(u("select",{"onUpdate:modelValue":t[1]||(t[1]=t=>e.$i18n.locale=t),class:"float-right"},[Ye,et,tt],512),[[y,e.$i18n.locale]])])]),u("div",at,[u("div",{class:["px-15px",{"text-true-gray-400 hover:text-black":"2d"!==r.type}],role:"button",onClick:t[2]||(t[2]=e=>r.type="2d")},p(e.$t("general.one_variable")),3),u("div",{role:"button",class:["px-15px",{"text-true-gray-400 hover:text-black":"3d"!==r.type}],onClick:t[3]||(t[3]=e=>r.type="3d")},p(e.$t("general.many_variables")),3)]),u(s)])])};const rt=_(qe);rt.use(E),rt.use(z),rt.mount("#app");
