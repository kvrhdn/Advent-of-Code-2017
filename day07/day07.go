package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Advent of Code 2017 - day 07")

	programTree := ParseInput(input)

	fmt.Printf("The bottom program is %q\n", programTree.Name)

	unbalancedProgram := FindTopMostUnbalancedProgram(programTree)

	fmt.Printf("Unbalanced program is %q, which holds:\n", unbalancedProgram.Name)
	for _, p := range unbalancedProgram.Holds {
		fmt.Printf("- %q (%v), total weight = %v\n", p.Name, p.Weight, p.RecursiveWeight())
	}
}

type Program struct {
	Name       string
	Weight     int
	HoldsNames []string
	Holds      []Program
}

func (p Program) String() string {
	return fmt.Sprintf("%q (%v) holds %v", p.Name, p.Weight, p.HoldsNames)
}

func (p Program) RecursiveWeight() int {
	weight := p.Weight

	for _, programHeld := range p.Holds {
		weight += programHeld.RecursiveWeight()
	}

	return weight
}

func (p Program) IsBalanced() bool {
	if len(p.Holds) == 0 {
		return true
	}

	weight1 := p.Holds[0].RecursiveWeight()

	for _, p := range p.Holds[1:] {
		if weight1 != p.RecursiveWeight() {
			return false
		}
	}
	return true
}

func FindTopMostUnbalancedProgram(programTree Program) Program {
	for _, p := range programTree.Holds {
		if !p.IsBalanced() {
			return FindTopMostUnbalancedProgram(p)
		}
	}
	return programTree
}

func ParseProgram(line string) Program {
	var p Program

	_, err := fmt.Sscanf(line, "%s (%d)", &p.Name, &p.Weight)
	if err != nil {
		panic(fmt.Sprintf("could not parse input %q", line))
	}

	indexProgramsAbove := strings.Index(line, "->")
	if indexProgramsAbove > 0 {
		p.HoldsNames = strings.Split(line[indexProgramsAbove+3:], ", ")
	}

	return p
}

func ParseInput(input string) Program {
	lines := strings.Split(input, "\n")

	programs := make(map[string]Program, len(lines))

	for _, line := range lines {
		program := ParseProgram(line)
		programs[program.Name] = program
	}

	bottomProgram := findBottomProgram(programs)

	linkPrograms(&bottomProgram, programs)

	return bottomProgram
}

func findBottomProgram(programs map[string]Program) Program {
	for _, candidate := range programs {
		if !isProgramHeld(programs, candidate.Name) {
			return candidate
		}
	}
	panic("no bottom program could be found")
}

func isProgramHeld(programs map[string]Program, programName string) bool {
	for _, p := range programs {
		for _, programHeld := range p.HoldsNames {
			if programHeld == programName {
				return true
			}
		}
	}
	return false
}

func linkPrograms(p *Program, programs map[string]Program) {
	for _, heldProgramName := range p.HoldsNames {
		heldProgram, ok := programs[heldProgramName]
		if !ok {
			panic(fmt.Sprintf("program %q could not be found", heldProgramName))
		}
		p.Holds = append(p.Holds, heldProgram)
	}

	for i := 0; i < len(p.Holds); i++ {
		linkPrograms(&p.Holds[i], programs)
	}
}

var input = `mqdjo (83)
jzgxy (15) -> usdayz, zvbru
altep (75)
gaieus (117) -> unkring, wjgvjlg
ejlbps (59)
uralh (17)
tzdmi (891) -> mhzwwo, mgybhs, pptdd
whntvd (2133) -> vuzxyz, hcfgnt, kuvek
aoqjxqk (99) -> qjfdh, vmusp, yqmclp
wwmjf (404) -> vvoadnv, sgujp
btgbr (54)
ftkfi (27)
vamssbi (194) -> picmfs, eutmjw
rrmrv (23)
vdavxhx (31)
ensgy (13)
ayngc (51)
qcblt (87)
wgqalow (6)
fvypts (47)
rioda (97)
fbmhsy (58) -> lqggzbu, xwete, vdarulb, rkobinu, ztoyd, vozjzra
ctmlo (29)
uezrt (28)
xvuxc (39164) -> nieyygi, hmcjz, ceizm
lvejxwu (71)
enogs (87)
lhdwm (58)
yeoeyme (55)
iwqzx (74)
xmnts (1113) -> rcmcf, dgihmb, xyhyumy
xhlgx (105) -> pamyclw, rrmrv, ulseuu, ihkrezl
mivzz (10)
kqwlhff (72)
lrlfs (152) -> jlgvr, sttwla
trzae (35)
jqgfz (63)
ymmpp (66)
vwuif (26) -> lrdpo, hxejeo, jxlzfjd
osxdc (172) -> qlxtmmw, gzvxqn, krboup
elfkuam (92)
sjsfqhv (143) -> rwxwe, dyjyjd, nwbcxj
kuvek (92) -> yqnfpeq, fcwev
gcsftd (765) -> lmpvxf, epqbl, rlsom
jfpbi (28) -> altep, ajsgpnx, vygscor
woqtzz (51)
mrqbmg (15)
iwdkssd (78)
gwqzfk (28)
zfvweqs (34)
anqvh (2135) -> gprcxh, oyahwdc, hxtrzk
eeikp (24)
nrlmfax (148) -> ammwyf, uaypo
eutmjw (12)
gsjah (31)
lgofh (90) -> alhvq, fgreoy, lmgwjss
untdlz (88)
rtdpm (24384) -> xhvyt, mrqug, kbkcw, gmlxpn
yfejqb (493) -> osxdc, jmyovo, nfmkvob
tzccgz (74)
cqksxxm (56)
vclagjf (40) -> tcziis, oxkbg
mqhuedy (7)
vvfukul (88)
acmkv (65)
gvfmsy (259) -> tokzfq, mnenfrc, feyacx, vpjiwwm
zcxiup (35)
dtivzu (36)
rwpaht (55)
yzuuvbm (63)
xuavduh (125) -> ftcxb, vynaaf
xmxuqtu (13) -> edwlp, xphhbib
grqiik (165) -> ookot, hazds
lnvdcr (31)
qdxkysa (155) -> nqcdwx, boliuot, dsjxy, mfjan
pdbruf (76) -> wwpagi, svmgylv
xilwqcz (106) -> gsjah, wntacd, ifvwm
abdnunu (35)
kihew (46)
jteda (97)
hzdox (43) -> oclbecq, iwdkssd
zdtzt (77)
arlra (258)
omynf (318) -> dwdjp, kiqfls, sfepx
xycww (64)
vlsmybk (83)
umhfxkg (40) -> cpfswhf, jpgbqgr, guzav, xumyvt, gmgwaa
ngdhmz (92) -> qgdeiw, dhcgig
mxgpbvf (52) -> jkaah, tfovn, foaqmc, gaieus, zxgptth
diizvrq (69)
dnuov (72)
cjofkgw (10)
dtacyn (42) -> xvuxc, eyyrn, udtsgjk, zprkrrn, rhhdc, rtdpm
grdpxxr (45)
bruwizl (11)
eetpan (48) -> mbffn, hhmwxe
ceupnk (31)
tuijc (27)
dhoqdtf (81) -> fzvnyc, ofazxj, vudnna
svmgylv (88)
bxoqjhj (70) -> woqtzz, ayngc
mpthchj (1933) -> acowb, ojaql, tqlyzrj
uovufgc (91) -> sumnuv, iklpkqe
reslc (26) -> epnsvw, zgpuqmv
sovwn (63)
mhgzyqt (30)
rfhlu (186) -> ifpae, bivky
yimjq (285) -> vvvgmqe, rjqdo, uavna
mgybhs (300) -> sgwxlb, beopgpz, ftkfi, tqndz
yareka (95)
qyiwwwy (39)
vudnna (93)
rbeigk (47)
qzhrv (360)
imgrm (12512) -> gcsftd, odxjz, umhfxkg, oqisjkn
nieyygi (7308) -> ptshtrn, mxgpbvf, cfqdsb, yfejqb
pytprd (75)
ykaodk (24)
upbwakj (66)
byxzbw (86)
xlzryhc (1814) -> ujxume, zrqzw, aqwfxoj, mxtgoh
ovonpn (93)
pfzsu (84)
twucd (7)
vtrgext (50)
vpjsxth (70) -> kobwl, eeupi
nwdzcl (54)
xwyxvr (29) -> pffhee, ycmdmvr, cupvdjn
mktppi (86)
oouozuv (46)
zhuue (5)
kpoou (34)
blskny (29)
grhibzl (43)
okwiabh (304) -> hyufw, hmhmfd, mwhef, lnkeeie
nlpdso (8)
vyqni (76)
mbffn (85)
tsuti (34)
plgldf (90)
obelp (105) -> icdtprw, mrqbmg, wwmzjie, qxhuuf
uueqz (52)
mumpgs (94)
ioxkn (67) -> mxqme, tcugzqe, wbopu, ipjaf
fyecx (50) -> gcyzqup, zkljjip
gmfzmby (90)
sxbko (84)
unkring (48)
kbdrjwl (2012) -> skjdw, ukqofsd, gonru
guzav (73) -> tuzwlaj, hsdwgl
qicxip (18)
arbiru (168) -> znbxm, zsetpb
rjqdo (36)
slvwki (11)
uanii (10)
zxvtq (1702) -> brenhzl, aekrgip, eojaq
izmnzef (330) -> idwncn, tcsoa, wagjjy
gcyzqup (63)
pjqtv (29) -> qvmdj, ovzrat, hjtmfck, yqkzm
awzbq (25) -> qjhugv, grqiik, ojcbjkx, blwgxoz, yspyrs, qiwio
badle (30) -> anuyfh, fcesi
ltxiy (341) -> smusd, kpoou
mwzana (12)
elkqdsi (52)
ggymz (83)
hxtrzk (172)
lufwlhx (34) -> mhrno, mtyuix, crzqw, mktppi
pvkps (92)
lrgsjqn (66)
fgtzv (63)
mtrmn (79)
yshpwl (19)
xnmmdhk (49)
aozsepf (113) -> rnaud, bclda
xhnlro (46)
gvrwy (71) -> idhdqo, bvjds
znbxm (12)
uaypo (61)
skoloe (28)
ypfxnr (233) -> avxapnw, maalkoh
fywvj (71)
ltpjwb (82)
atbxzb (221) -> jepepf, torpy
oxhnlv (13) -> nrozgxx, kjjzdvp
eyyrn (44234) -> rchso, dqwet, qoqdhpx
jxvod (9)
jrvumy (92)
ukejvdd (34)
ulseuu (23)
lmgwjss (34)
vpgsx (65)
ykvww (51)
zwmbuyl (71)
vcsfjyd (24)
vdxfff (38)
vozjzra (70) -> ginyxrk, sgbejy
cizxbj (193) -> abndmkw, qyiwwwy
yroyqeo (13)
hkanwp (13)
plctc (345) -> rngety, zebcbo
ayode (38)
rwwww (22)
rundk (116) -> cwwsi, enogs
ixqqqjj (269) -> aozcb, ahvhqoe, qycnzr, ojwezo
jdjvp (54)
tcsoa (165) -> rwpaht, idcowch
owlwzwx (36)
yhljr (26)
rciod (54)
tnhnp (154)
shmvzp (43)
gkoze (158) -> lrpzf, zkbzbn
picmfs (12)
ghjcbf (3487) -> ccpeyc, okwiabh, qxsie
tcziis (77)
raugvuk (78)
jepix (56)
wlbuyk (73)
fgreoy (34)
idqsz (27)
akcgcxy (90)
hkcobh (20)
hgmzkam (63)
teauca (399) -> cwlvxo, jsohd
dkjcz (34)
xulggnd (253) -> spfyg, qqpvjbu
gyuiagl (126) -> xwkudo, irlntnq
wnkaz (9) -> ovonpn, idnken
mlqfisf (150) -> lrgsjqn, rfewc
ikknx (23)
aekrgip (155) -> xbtcnx, odfgjh
qfvbg (420) -> jdyblpx, llqfm, vydlsvt, gkoze, ofvidy
xddngkv (75)
hazds (17)
awvyj (43)
swsbeo (1351) -> crhrx, ylbuzsq, wcuqhwq
irlntnq (23)
asdevc (66)
dniwy (7)
qedlf (60)
yrsfcq (72) -> emdaoqf, qezss, frozh, aakyzvs
ptocen (64)
vqxyr (66) -> jdjvp, rciod, zwofqok, igkxvqh
kjgfm (73)
hmhmfd (100) -> hbuwd, rrohrwo
lezpa (84)
cfqdsb (556) -> feizivp, vjhsczt, gmxbjg
oclbecq (78)
qqypffg (67)
lrdpo (91)
kvolhll (77)
nhiujz (271)
mxqme (32)
wjmfbz (75)
sgltbu (80) -> guphfh, uzxjy
wvjumh (90)
tlwfzzu (87)
vmusp (83)
jlslmo (13)
bhysq (10)
daxfuxl (13)
arqxnjl (1242) -> xvtdoyr, rkupoek, izapo
mtyuix (86)
qjuvyr (56)
pfctdp (127) -> hjibwmq, bseajmt
erjsdm (65)
jkaah (97) -> blskny, ctmlo, sxlecxm, fjixq
lldszk (2467) -> xxslxbt, elfkuam
wzejrx (43)
yqnfpeq (20)
ztoyd (227) -> cuwjy, rhamh, idwgpl
ocvjwr (63)
bxgyjw (60)
bnggqtn (58)
cupvdjn (50)
elebwva (173) -> skmum, hibcfm
tmhgfe (223) -> tqfel, daxfuxl
slduca (71)
ybcta (17)
cyniz (855) -> cgzpd, kimrt, sxaib, syppxhj, pfctdp, qdxkysa
ledbvvq (240) -> cchmz, grdyhff
alhvq (34)
eoxmesy (24)
hzzwzx (7)
odtzcu (37)
rnviq (89) -> rkcdg, fyalh
qvmdj (30)
phqjxv (161) -> kihew, qiyidu
ifpae (87)
ooqsaz (97)
jnsph (93) -> aeiee, vcpkha, zvdmblb, sbvcz, yriybrc, aoqjxqk, nhlbui
vvoadnv (23)
btabyp (203) -> gegad, rmyjxl
aosrim (40)
qgkhdex (41)
qvhpl (84)
rkobinu (164) -> ymccua, xtsnf
fcesi (73)
pxdezu (43)
phjllc (132) -> slvwki, fhqqbsx
ezrpvp (94)
jhlxf (74)
ypfss (172) -> kltwy, grhibzl
jrqmxt (23)
nfchsda (43)
ctnfhim (71)
dksbpu (50)
sxlecxm (29)
exnlyy (26) -> nwooizm, jhlxf
hyurmnj (70)
rarvsbv (91)
pmpmgak (19)
pnkce (63)
isozjve (84)
axondt (154) -> nmqmbc, kmhzija
yuzamn (79) -> rxwbam, ptocen
orlwcra (71)
sxghqf (17) -> hswecad, pbdww
saksuh (72)
hromhg (50)
lqplgho (198) -> fywvj, jbxkrhf, slduca
pptdd (76) -> cthyoqb, mvjppl, mhtfkag, dsxamd
baosr (80) -> byxzbw, meajs
qzslx (83) -> vtrgext, fnwzvf
rhqmcl (15)
nxvfovk (97)
nwbcxj (68)
yxstzne (21) -> dnuov, mqwmam
wlheoem (6)
luqtj (171) -> kjgfm, wlbuyk
sgujp (23)
odxjz (37) -> luqtj, igpyz, qzltox, ufwsbx
gprcxh (84) -> xhlqyt, rwwww, xwsgv, gnqasz
iklpkqe (49)
vmianfx (51)
zkljjip (63)
oaucz (74)
zaahfws (72) -> cpncdpn, kqwlhff, saksuh, rwnnzj
omnflj (26)
yzzkb (98) -> nxywzy, zukmsws
upttsy (46) -> lhusw, cabbpbp
awdam (122) -> zboinwm, erjsdm
idnken (93)
xpxwq (56) -> cizxbj, ynxwu, cpbfkf, izfny, dontia, viohwud, nhiujz
adklqb (13290) -> pzsdbxi, jahpohn, yurlqf
yurlqf (1176) -> wiyuhex, cujasg, hzdox
qjagx (265)
ccpeyc (13) -> eowou, kbbck, oxhnlv, obelp, rxmnfgj, ahfsddo, yxstzne
amdyzs (28)
nxywzy (96)
vjunjns (170) -> dgrfnj, ceathvu, lspkp, nqmfyxs
juzzs (30)
hrfdy (79)
klbkp (1736) -> qzslx, uycbfkf, gvrwy, eadxb, rysbl
ggijqt (21)
ebeltx (60)
lnkeeie (58) -> hrfdy, mtrmn
pwsfhgi (52)
zdymtqx (49)
zenuyzk (84) -> teqabnc, kmqopjq
eowou (119) -> jrqmxt, ikknx
mwdjsxv (51)
oritewf (164) -> jxvjp, htvenp, cgzcdzi, ukysk, ocakmk
vygscor (75)
zkvql (20)
sttwla (65)
mmlkoh (2271) -> jmrbakg, bpatyq, yucyi
yiupwfc (88) -> izrwts, yimjq, vfqbff
ynxwu (127) -> ghicha, xohoejm
fqffx (94)
xhlqyt (22)
tqndz (27)
crzqw (86)
epqbl (140) -> zdxksj, zkvql
ujckr (1913) -> phjllc, ibuys, tnhnp, sgltbu
lezek (775) -> rkgwky, menclt, ekxdgkq
qzoqo (79)
rnaud (68)
zwfvctv (5)
lyebndp (49)
ircvo (11)
dpqnim (39) -> pxdezu, awvyj, yiigtz, wzejrx
pmtssy (18)
zjyhf (56)
kdzuk (78)
mnojup (73) -> sxbko, umnvlw, zvodhi, pfzsu
befeb (96)
tpcir (16) -> jcwdnlr, goeprx, wwskbob, tneetox
xohoejm (72)
qdufkz (94)
ghyii (60) -> esmmub, qjuvyr
zprkrrn (46528) -> ghjcbf, jsrdr, ywyvxv, ujsqcr
yxpavwg (38)
dqwet (6621) -> uibut, bgugou, izmnzef
zqhmft (98)
rttcn (108) -> ylzoeee, ooqsaz
qzjtxi (13)
yssxm (1494) -> yuzamn, hybkngi, jzgxy, umryvx, mykoaba
xbtvul (30)
vgmnvqr (79)
rchso (42) -> viwwie, kppigwv, zesja, azaxxvz
njxkbbs (37)
sambp (84)
ctmyow (14952) -> lwzfwsb, awzbq, dvdbswf
bjenf (94) -> bjxuox, offqpoc
ekhsurt (223) -> skoloe, ijzzl
mhnjio (29) -> whntvd, ujckr, dysfcb, yssxm, cyniz, jnsph, mmlkoh
uycbfkf (9) -> tlwfzzu, mnieqt
rcmcf (86) -> jteda, iydwka
zboinwm (65)
boliuot (31)
udcgj (455) -> ynydd, uxbgd, xogbbc
uozayo (39)
fltthw (117) -> hdpizox, dfqahip
dgrfnj (22)
wyusin (38)
zvbru (96)
owhgyez (37)
dtwdad (81) -> zwxvo, nsldzh
yqmclp (83)
semtwrh (82)
ihteano (97)
mhrno (86)
wwpagi (88)
ngcyru (56)
ngwldiu (55)
virmpl (175)
aakyzvs (367)
vubuie (79)
kpczbnm (26)
jahpohn (51) -> uawat, zenuyzk, pwonvpm, bborg, zjnaor, lniuonj, yxmkycd
jxlzfjd (91)
mgvgcf (30)
xphhbib (83)
vzubyy (84)
spfyg (13)
emaqhkl (49)
bvjds (56)
aqwfxoj (50) -> byytq, qybcrhn
ukysk (234) -> ircvo, oeinoe
lpwze (175) -> ltaujs, hosjc
wikrvq (15)
anuyfh (73)
fgfkrxp (69)
rrtqw (8)
lehgocg (42) -> vpgsx, pydanp, nppyzr, kdgxoi
nzuvke (16)
ehcjup (18) -> kqduo, sambp, rxehio
jjqwgq (34)
jepepf (63)
oikfudn (135) -> ynubz, xqqdtu
maicx (48)
ammwyf (61)
xqkrdp (30)
ylbuzsq (89)
ulrqcxx (17)
bfhjnn (13)
hyhcihf (28)
jqtbsy (32) -> yqzkn, dhiavwc
shelxb (66)
jmfybp (100) -> qzhrv, ckyvnnj, dhoqdtf, hhrqizc
fgscpqx (96)
viohwud (256) -> zwfvctv, rjfktu, gfygwj
erqjppo (97)
gldbbde (24) -> hqxsfn, vwgmnj, rphmi
hvrve (36)
ajsgpnx (75)
oklhlne (12)
radodhv (5)
kycxdgx (29) -> wjfar, pmjvz
mvjppl (83)
brilycx (97)
duklm (47)
hibcfm (87)
yjeavx (53)
vwgmnj (97)
uavna (36)
cujasg (65) -> qqypffg, pgpmlgz
rqura (47)
rfewc (66)
ayklc (595) -> lrwlhbc, ypfxnr, gldbbde
nvldfqt (57)
bzvwi (20)
rcstlvu (10)
fnwiq (63)
twikxu (59)
byytq (60)
mmlcvqo (37)
sqowhp (34)
nwkqcsc (24)
czlcjl (626) -> qmtofe, kefdbed, rnviq
pqavmdg (17)
wfjme (21)
latbj (240) -> vdavxhx, lnvdcr
pthrzgd (38)
ghicha (72)
vlrtfe (114) -> aosrim, pnlmj
izrwts (337) -> amdyzs, epilojy
cjksouo (154) -> oaucz, iwqzx, dlszm, tzccgz
rysbl (21) -> dzmqhbp, rscwd
qilfd (86)
zebcbo (42)
atdfonp (52)
oukcyj (7)
pydanp (65)
bpatyq (32) -> iszhqi, gwnxq
bcuxejj (63)
hdsvjwr (1480) -> stryo, jmtmw, tbdfmz
ywyvxv (3190) -> gydenqe, ydhxt, yiupwfc
xqqdtu (20)
xkqiqns (69)
wcuqhwq (89)
pczqv (88)
yhmbooy (35)
bjxuox (39)
kqduo (84)
qinuxsr (238) -> ajcpmg, mdtuwks, huvag, bxtzi
wjfar (85)
nfmkvob (208)
wlrrckg (697) -> pcxcr, lmyinxy, ekkgfax
plifn (26)
crhrx (89)
zwofqok (54)
afurlcd (63)
tohwd (1394) -> eeikp, wsqspc
ojcbjkx (185) -> mqhuedy, twucd
wbopu (32)
fjixq (29)
mwhef (136) -> zcvqw, tnxhf
jdyblpx (130) -> mmhim, phomdt
nhdfnv (51) -> fltthw, kycxdgx, xilwqcz, prbrv, dzibhav, orxuz, zozpw
ckyvnnj (234) -> pnkce, yzuuvbm
vllgqt (5)
kbrbq (42)
stryo (219) -> idqsz, czala
ksrbxbh (255)
vqrcq (127) -> zcxiup, vvmfk
orxuz (172) -> bfltqak, skogcpe, jxvod
jbxkrhf (71)
egxng (20)
fswhfws (90) -> ocikhq, ayyfw, gwqzfk
gegad (26)
ifvwm (31)
kmqopjq (81)
rdxfau (8)
vvvgmqe (36)
rssmrnc (53)
eofrs (35)
vdarulb (246) -> hzzwzx, oukcyj
rtbjlii (157) -> pcxny, zjzyo
zjnaor (118) -> uuozpc, xycww
zmkyj (65) -> fqffx, ikzsd, qdufkz
bivky (87)
roirllq (75)
osehy (12)
odqfu (37)
rhhdc (3564) -> mhnjio, mzrntf, imgrm, sgamgp
xghyc (35)
umeup (55)
pihhh (859) -> jfpbi, xvsrff, phqjxv
yucyi (35) -> pqavmdg, ulrqcxx, eaemt
kobwl (74)
zudwpn (2104) -> gvfcl, dicxtfr, reslc
acheq (179)
hxxvca (41)
sbvcz (198) -> mveyh, xddngkv
rphmi (97)
kotdcly (26)
rwrmuv (78)
skmum (87)
rngety (42)
lhusw (74)
srrbgff (42)
fvpbw (23) -> vmianfx, ykvww, mwdjsxv
huvag (23)
tokzfq (22)
jylrfoj (72)
cpbfkf (77) -> brilycx, ihteano
fzofpn (75)
xogbbc (32) -> asdevc, shelxb, rurnike, ymmpp
ajihys (56) -> bnggqtn, lhdwm
rhamh (11)
xvsrff (205) -> eoxmesy, nwkqcsc
hbuwd (58)
xjjhicj (28)
iszhqi (27)
qqpvjbu (13)
lmyinxy (287) -> bhysq, rcstlvu
ubdju (438) -> wgqalow, xoknnac
kmhzija (88)
lwzcyq (42)
omwhgx (79)
fvqgb (71)
ufsus (26)
vynaaf (21)
wwdhply (52) -> lldszk, klbkp, vysbz, xgtged, kbdrjwl, hcxrwwv, anqvh
sgamgp (14279) -> czlcjl, wlruzq, koklf
nfiyxwf (46)
smusd (34)
cuuvm (195)
krboup (12)
vysbz (1571) -> jzxcnkv, rfhlu, zaahfws
rojomwr (32)
scrykpr (47) -> dlochnp, kuxdz
zesja (57) -> ltxiy, ngskki, emjhgx, mnojup, ixqqqjj, wnpyrz
skogcpe (9)
brenhzl (119) -> bzvwi, lcvtfz, egxng, hkcobh
rkgwky (75) -> grdpxxr, xflynr, yrixhb, lalemvr
odfgjh (22)
wlruzq (1135) -> rrtqw, jmcafz
torpy (63)
yqzkn (93)
hcfgnt (132)
tvobw (69)
yxmkycd (120) -> afurlcd, qyerfcc
ikzsd (94)
ludhiab (96)
xiayty (80)
jsohd (15)
tqfel (13)
huvrn (86)
wzlrws (26)
ywxjbc (39)
dwdjp (31)
hpfhhg (63)
hybkngi (195) -> wlheoem, oxbfji
wsqspc (24)
eaemt (17)
lkwhxen (85)
wmzmy (59)
yspyrs (147) -> kpczbnm, yhljr
fuwat (65)
alsay (7) -> snloung, djdmsc, sdqll
zxgptth (189) -> fknto, dpfuu
kaihsu (228) -> ggijqt, wfjme
kbkcw (6055) -> swsbeo, pihhh, wlrrckg, fbmhsy
rnatyw (59)
xhvyt (57) -> xlzryhc, zudwpn, swsva, mpthchj, wxtlbn
ywrnqrp (179)
fzvnyc (93)
tantkl (77)
xvtdoyr (45) -> tatvgv, nrvpgth, luugrb, maicx
cpzdjum (55)
blwgxoz (87) -> mzeqa, abghrf
uxbgd (184) -> zbnynh, twrim
dnmrtok (221) -> epzukel, fgtzv
tnxhf (40)
dgihmb (175) -> liglllj, xghyc, eofrs
jmrbakg (86)
rurnike (66)
ceizm (58) -> ebqjlp, xpxwq, vfigul, arqxnjl, xmnts, sdbkh
cmmnbug (77)
sgwxlb (27)
lmbucr (55)
dsxamd (83)
qcgykv (88)
wxtlbn (604) -> lmipr, pxwsb, lufwlhx, mixmiv, ledbvvq
qbwfy (19)
omznvh (320) -> zmemxd, ihjpab
xoldlj (139) -> umeup, ewvfvx
zdxksj (20)
foaqmc (71) -> ctnfhim, zwmbuyl
azaxxvz (1383) -> lrlfs, gqvnw, vqxyr, mlqfisf
hcdyhi (63)
hosjc (52)
sfepx (31)
gzvxqn (12)
maalkoh (41)
qwnexz (8)
guphfh (37)
uuozpc (64)
nrvpgth (48)
aljvja (98) -> arbiru, iumni, lgofh, nkhdzon, yxpkynj, qqqsv, qiduej
pctyehs (304) -> ensgy, bfhjnn
xkzwk (17413) -> tdbcau, rtbjlii, vwuif, hfyaud
izfny (148) -> qgkhdex, qhrmshn, sewsk
dhiavwc (93)
wnpyrz (323) -> shmvzp, nfchsda
cdaue (59) -> odqfu, knqruz
syppxhj (264) -> vllgqt, radodhv, zhuue
yfozx (55)
rrohrwo (58)
cwlvxo (15)
hhukemh (155) -> osehy, lfsqjx
luugrb (48)
ojoes (39)
zjzyo (71)
ngskki (409)
rjfktu (5)
hcxrwwv (2126) -> sxghqf, oikfudn, virmpl
umryvx (96) -> owhgyez, jkorfu, odtzcu
zvdmblb (222) -> hpfhhg, bcuxejj
swspgk (55)
dontia (127) -> ofcfw, zjstwhc
qlxtmmw (12)
ceathvu (22)
llqfm (64) -> omvsg, xiayty
ekxdgkq (83) -> qilfd, huvrn
gwnxq (27)
mveyh (75)
eadxb (167) -> mnqis, fglxnn
bexvpzn (55)
ywquvbe (17)
zvodhi (84)
aleenbl (55)
quqts (63)
kimrt (30) -> qdtel, zznqt, mqdjo
eojaq (73) -> hcdyhi, jqgfz
blhxnrj (55)
rvycq (47)
ycmdmvr (50)
bfltqak (9)
ydhxt (275) -> hqekpst, jpwof, dbuno, jfavd
aeiee (348)
jmcafz (8)
lfsqjx (12)
epzukel (63)
jsqss (19)
qgdeiw (83)
hfyaud (299)
ybtmtno (26)
ridid (16)
uslizt (59) -> xbtvul, mhgzyqt, mgvgcf
vsnpvm (87)
ojaql (157) -> ppudoq, ftsejjt, cjofkgw
wyjbjf (8)
wntacd (31)
ymccua (48)
cgzpd (201) -> vndxvi, uozayo
iumni (192)
ofazxj (93)
uoblmgo (46)
cftahaw (70)
rlsom (66) -> nvldfqt, mdakca
jlgvr (65)
ozbig (98)
qpbebnc (76)
ayyfw (28)
tkqlt (107) -> pmecds, dtivzu
dnlftxp (81)
mqwmam (72)
oxkbg (77)
nbycb (55)
edwlp (83)
nkhdzon (34) -> hjrhs, qzoqo
ekkgfax (70) -> ugxovve, vgmnvqr, wafanbg
hjrhs (79)
zwxvo (58)
oyahwdc (16) -> rwrmuv, lyuznk
uaatgf (11) -> uueqz, atdfonp, elkqdsi
dyjyjd (68)
rxehio (84)
jzxcnkv (202) -> vubuie, omwhgx
hxejeo (91)
zozpw (199)
ptshtrn (526) -> uslizt, pjqtv, bzbdorp, atvkttc
qzltox (97) -> lmbucr, blhxnrj, cpzdjum, swspgk
vywneg (53)
cigzzld (90)
wxqadgl (88)
goeprx (59)
nhcio (88)
jsrdr (2962) -> rezncs, udcgj, cjtrpm
lswgx (75)
czala (27)
jwftyhn (76)
pgpmlgz (67)
abqar (319) -> ooque, nyimka
cuwjy (11)
kltwy (43)
fknto (12)
zvgxbve (37) -> zqhmft, mwfno, mnyblkw, ozbig
xwete (92) -> lezpa, qvhpl
qdhte (168) -> ypysb, nwdzcl, btgbr
nkrsre (18)
hqekpst (68) -> eqakw, gmfzmby
nppyzr (65)
ojmudkf (68) -> rhqmcl, wmlvvzy
ycpgrap (52)
fnwzvf (50)
tfovn (213)
kefdbed (121) -> tuijc, tngvv
idwgpl (11)
qybcrhn (60)
zvfdl (72)
rwxwe (68)
tcugzqe (32)
liglllj (35)
zkbzbn (33)
mbsuxms (28)
ugxovve (79)
udtsgjk (56) -> adklqb, xkzwk, wwdhply, ctmyow
jlvflv (47)
ajcpmg (23)
hmcjz (11191) -> ioxkn, wnkaz, cuuvm
zjstwhc (72)
qxhuuf (15)
unoqw (34)
xmzvsd (55)
rwnnzj (72)
xyhyumy (104) -> vhdyc, vvfukul
ginyxrk (95)
cgzcdzi (70) -> uoagriu, jnfcc
oxbfji (6)
yybuekt (96)
zznqt (83)
yiflicd (97)
hhmwxe (85)
avxapnw (41)
nrozgxx (76)
prbrv (61) -> nfiyxwf, vlhtnof, xhnlro
xwsgv (22)
jmtmw (205) -> sqowhp, qjkqxp
epilojy (28)
qjkqxp (34)
cchmz (69)
kkfcate (114) -> juzzs, xqkrdp
ihkrezl (23)
jkdrr (7)
mqhelgf (31)
fyalh (43)
icdtprw (15)
qfuxpos (94) -> omnflj, ybtmtno, cowhe
vndxvi (39)
sbaswic (34)
rbbkcz (81)
aozcb (35)
tdbcau (109) -> zhfenxn, yareka
jpgbqgr (61) -> fgscpqx, ludhiab
ftcxb (21)
egrtqo (60) -> tigky, mbqfitx, lqplgho, ivkymv, omynf
mnyblkw (98)
ipjaf (32)
mfweuw (96)
kjrup (11)
usdayz (96)
vbfwgo (85)
lwzfwsb (820) -> alsay, sgzxb, cdaue
dzibhav (199)
bxtzi (23)
ihjpab (65)
xtsnf (48)
ofvidy (8) -> zvfdl, zhezjdt, jylrfoj
oehapu (54)
looasd (161) -> ycpgrap, pwsfhgi
mrqug (6751) -> oritewf, nhdfnv, pblscqj, rmknzz
rmknzz (649) -> looasd, qjagx, hkati
ykidq (61)
dlochnp (60)
fglxnn (8)
ooque (55)
mnqis (8)
lspkp (22)
dgpydnw (15)
rezncs (776) -> gjzwa, uovufgc, hdhlyxq
sumnuv (49)
zgpuqmv (52)
wagjjy (91) -> pvkps, jrvumy
wmlvvzy (15)
pnwrdla (91)
cthyoqb (83)
pwonvpm (162) -> srrbgff, kbrbq
smvkv (59)
atvkttc (41) -> msxbv, oehapu
ioaogd (46) -> wzlrws, ufsus
lmpvxf (78) -> ukejvdd, sbaswic, jcbdf
vfigul (1371) -> vlrtfe, upttsy, vclagjf
oeinoe (11)
uoagriu (93)
rkcdg (43)
qycnzr (35)
vhupxjs (157) -> pmtssy, nkrsre, qicxip
msxbv (54)
oifkd (69)
qjhugv (199)
nsfrscb (278) -> oklhlne, mwzana
pmjvz (85)
qmtofe (19) -> kdzuk, raugvuk
fhqqbsx (11)
zukmsws (96)
kuxdz (60)
ofcfw (72)
teewbxk (39)
hsdwgl (90)
offqpoc (39)
akvdjmc (75)
pcxcr (214) -> eqdud, mqhelgf, ceupnk
zhfenxn (95)
eudjgoy (351) -> bnomgct, ojoes
wjgvjlg (48)
bborg (10) -> tarxkki, ejlbps, smvkv, wmzmy
xoknnac (6)
bzbdorp (73) -> pthrzgd, vulklzo
qczjylz (19)
cpncdpn (72)
kjjzdvp (76)
dzmqhbp (81)
xbtcnx (22)
fcwev (20)
vlhtnof (46)
eqdud (31)
ukqofsd (139) -> njxkbbs, mmlcvqo
xxslxbt (92)
jkorfu (37)
nhlbui (348)
pkgri (41)
fsowc (88)
idhdqo (56)
uqmgy (1245) -> ysnvnu, rundk, yzzkb
hqgsj (69)
yrixhb (45)
yqkzm (30)
wdeclb (13)
rxmnfgj (27) -> hqgsj, diizvrq
uawat (230) -> rdxfau, qwnexz
rnzglhk (69)
psldjp (66)
yedhncv (16)
sxaib (279)
gfquwm (13)
jmzwqs (49)
dpfuu (12)
edgsqu (32)
acowb (33) -> cmmnbug, tantkl
fbtqea (49) -> ubdju, omznvh, cjksouo, wwmjf, fdbehfx
xllppou (61)
zqrfz (29)
mdakca (57)
abndmkw (39)
ejtuc (42)
mmhim (47)
dicxtfr (110) -> uanii, xiiamxd
gmlxpn (6182) -> tzdmi, egrtqo, uqmgy
sgzxb (119) -> dniwy, jkdrr
tuzwlaj (90)
gmxbjg (187)
nwooizm (74)
ovzrat (30)
mwfno (98)
mbqfitx (249) -> rbbkcz, dnlftxp
ftsejjt (10)
beopgpz (27)
oqisjkn (672) -> kktfosp, vhupxjs, dpqnim
skjdw (165) -> ridid, yedhncv, nzuvke
jfavd (72) -> nhcio, qcgykv
gydenqe (739) -> fvpbw, badle, fyecx
mhtfkag (83)
tneetox (59)
hkati (58) -> tvobw, xkqiqns, oifkd
qxsie (136) -> bjenf, bxoqjhj, ajihys, ghyii, qfuxpos, gyuiagl
ynydd (220) -> hypma, qbwfy, jsqss, yshpwl
esmmub (56)
ypysb (54)
feizivp (40) -> jmzwqs, zdymtqx, emaqhkl
lalemvr (45)
mxtgoh (132) -> pmpmgak, qczjylz
ujxume (88) -> pkgri, hxxvca
ylzoeee (97)
zhezjdt (72)
ojwezo (35)
cjtrpm (90) -> xmxuqtu, acheq, tkqlt, xwyxvr, hhukemh, eqyac, ywrnqrp
sewsk (41)
lqggzbu (164) -> rojomwr, edgsqu, ifrxho
ppudoq (10)
vydlsvt (176) -> vcsfjyd, ykaodk
uswkbi (13)
cpfswhf (223) -> dgpydnw, wxexs
feyacx (22)
menclt (255)
igkxvqh (54)
cabbpbp (74)
idwncn (157) -> rnatyw, twikxu
iuwokd (89) -> vlsmybk, ggymz
qiduej (90) -> jjqwgq, tsuti, unoqw
imuik (434) -> awdam, baosr, tpcir, pdbruf
jijdq (202) -> dkjcz, zfvweqs
sgbejy (95)
rxwbam (64)
mzeqa (56)
mpbjcj (202) -> dksbpu, hromhg
kdgxoi (65)
tatvgv (48)
lyuznk (78)
vhgimqo (75)
wiyuhex (160) -> gfquwm, qzjtxi, uswkbi
hswecad (79)
qwnxx (353) -> yxpavwg, ayode
qjexd (10)
htvenp (130) -> quqts, ocvjwr
vcpkha (309) -> jlslmo, wdeclb, hkanwp
gyyld (91)
tigky (335) -> wyusin, vdxfff
dysfcb (1833) -> vtnsti, fswhfws, exnlyy, kkfcate
mdtuwks (23)
pbdww (79)
ewvfvx (55)
pxwsb (210) -> isozjve, vzubyy
rscwd (81)
abghrf (56)
ujsqcr (94) -> hdsvjwr, fbtqea, zxvtq
qiyidu (46)
xumyvt (197) -> mbsuxms, uezrt
fdbehfx (86) -> pnwrdla, gyyld, twxghfp, rarvsbv
avwkq (63)
ocikhq (28)
xzmtk (61)
qezss (83) -> vnmhbwf, fvqgb, lvejxwu, orlwcra
jcbdf (34)
rkupoek (83) -> kvolhll, zdtzt
bgugou (39) -> ekhsurt, lpwze, xulggnd, mjckr
ocakmk (115) -> fvypts, rbeigk, jlvflv
yiigtz (43)
rovmviv (164) -> rnzglhk, fgfkrxp
zmemxd (65)
omvsg (80)
ysnvnu (110) -> cigzzld, wvjumh
zrqzw (38) -> upbwakj, psldjp
tbdfmz (79) -> nxvfovk, yiflicd
ydishxc (75)
tqlyzrj (135) -> kotdcly, plifn
lrwlhbc (257) -> wvibg, zqrfz
xwkudo (23)
vjhsczt (77) -> xmzvsd, nbycb
qjfdh (83)
izapo (237)
fipmpb (13)
dsjxy (31)
bclda (68)
zbnynh (56)
hdhlyxq (63) -> hgmzkam, avwkq
ijzzl (28)
gmgwaa (73) -> akcgcxy, plgldf
hyufw (216)
jnfcc (93)
swsva (65) -> atbxzb, zmkyj, elebwva, dnmrtok, gvfmsy, sjsfqhv, rehfw
yriybrc (196) -> qpbebnc, fwmpx
hhrqizc (250) -> aleenbl, yeoeyme
idcowch (55)
nsldzh (58)
ufewl (6) -> uoblmgo, oouozuv
ebqjlp (933) -> btabyp, byoksw, iuwokd, ksrbxbh
hjtmfck (30)
gjzwa (163) -> fipmpb, yroyqeo
usnfnjb (53)
eviqdbq (15)
eqakw (90)
ltaujs (52)
ifrxho (32)
gnqasz (22)
gonru (25) -> mumpgs, ezrpvp
lmipr (195) -> ykidq, xzmtk, xllppou
vvmfk (35)
sdqll (42)
jpwof (178) -> abdnunu, yhmbooy
vfqbff (93) -> ydishxc, vhgimqo, wbjbpjf, roirllq
fwmpx (76)
knqruz (37)
epnsvw (52)
akznnid (69) -> bxgyjw, ebeltx, qedlf
vhdyc (88)
hqxsfn (97)
pgqenlk (143) -> rssmrnc, usnfnjb
jmyovo (68) -> hyurmnj, cftahaw
wvibg (29)
pnlmj (40)
lniuonj (216) -> eviqdbq, wikrvq
gmuco (35)
qqqsv (18) -> vsnpvm, qcblt
igpyz (245) -> hvrve, owlwzwx
kppigwv (1920) -> dtwdad, xhlgx, vqrcq
rvwvokn (218)
ynubz (20)
viwwie (699) -> nsfrscb, mpbjcj, rttcn, latbj, lehgocg, rovmviv
rmyjxl (26)
qyerfcc (63)
dbuno (226) -> kjrup, bruwizl
iydwka (97)
yxpkynj (192)
ibdhatw (75)
umnvlw (84)
dfqahip (41)
pcxny (71)
bseajmt (76)
mqeep (460) -> ehcjup, kaihsu, jijdq, nrlmfax
vuzxyz (6) -> fnwiq, sovwn
qdtel (83)
eqyac (81) -> lyebndp, xnmmdhk
mzrntf (11964) -> pxhqf, tohwd, imuik, aljvja
nqcdwx (31)
sdbkh (303) -> qdhte, qinuxsr, ymjxszk, pctyehs, axondt
gqvnw (282)
eeupi (74)
qhrmshn (41)
hjibwmq (76)
nqmfyxs (22)
grdyhff (69)
twrim (56)
djdmsc (42)
mnieqt (87)
cwwsi (87)
hrbzakg (42)
jxvjp (236) -> mivzz, qjexd
xflynr (45)
hdpizox (41)
vpjiwwm (22)
teqabnc (81)
mjckr (85) -> rioda, erqjppo
wwskbob (59)
wxexs (15)
snloung (42)
mhzwwo (392) -> wyjbjf, nlpdso
phomdt (47)
mfjan (31)
frozh (367)
lcvtfz (20)
emjhgx (189) -> kteflr, ngwldiu, yfozx, bexvpzn
vtnsti (104) -> gmuco, trzae
nmqmbc (88)
dlszm (74)
lrpzf (33)
ibuys (98) -> xjjhicj, hyhcihf
bnomgct (39)
uibut (861) -> ojmudkf, ufewl, ioaogd
byoksw (91) -> semtwrh, ltpjwb
pblscqj (943) -> scrykpr, xuavduh, uaatgf
zsetpb (12)
cowhe (26)
dvdbswf (187) -> ypfss, vjunjns, arlra, ngdhmz
xgtged (77) -> eudjgoy, zvgxbve, plctc, teauca, qwnxx, abqar
pffhee (50)
ahvhqoe (35)
pamyclw (23)
ufwsbx (93) -> ngcyru, zjyhf, cqksxxm, jepix
pmecds (36)
qoqdhpx (846) -> yrsfcq, jmfybp, qfvbg, lezek, mqeep, ayklc
pxhqf (197) -> akznnid, xoldlj, pgqenlk, aozsepf, tmhgfe
wafanbg (79)
nyimka (55)
ynypbb (47)
wwmzjie (15)
tarxkki (59)
rehfw (221) -> lwzcyq, hrbzakg, ejtuc
mykoaba (37) -> vbfwgo, lkwhxen
pzsdbxi (1621) -> jwftyhn, vyqni
kktfosp (160) -> ywquvbe, uralh, ybcta
ahfsddo (35) -> fuwat, acmkv
kteflr (55)
hypma (19)
wbjbpjf (75)
emdaoqf (217) -> ibdhatw, akvdjmc
qiwio (11) -> rvycq, duklm, rqura, ynypbb
vnmhbwf (71)
xiiamxd (10)
uzxjy (37)
ymjxszk (30) -> fzofpn, wjmfbz, pytprd, lswgx
jcwdnlr (59)
vulklzo (38)
ookot (17)
mnenfrc (22)
mixmiv (90) -> mfweuw, befeb, yybuekt
zcvqw (40)
gfygwj (5)
kiqfls (31)
ivkymv (59) -> wxqadgl, fsowc, pczqv, untdlz
meajs (86)
dhcgig (83)
kbbck (87) -> teewbxk, ywxjbc
gvfcl (24) -> vywneg, yjeavx
twxghfp (91)
tngvv (27)
koklf (61) -> jqtbsy, eetpan, vamssbi, rvwvokn, vpjsxth`
