package main

import "testing"

// func TestRemoveDeadInstructions(t *testing.T) {
// 	for _, tc := range []struct {
// 		input string
// 		want  string
// 	}{
// 		{
// 			input: "",
// 			want:  "",
// 		},
// 		{
// 			input: "abcdon't()efgdo()abc",
// 			want:  "abcabc",
// 		},
// 		{
// 			input: "abcdon't()edon't()fgdo()abc",
// 			want:  "abcabc",
// 		},
// 		{
// 			input: "ado()bcdon't()edon't()fgdo()abc",
// 			want:  "ado()bcabc",
// 		},
// 		{
// 			input: "ado()bcdon't()edon't()fgdo()abdo()c",
// 			want:  "ado()bcabdo()c",
// 		},
// 		{
// 			input: "ado()bcdon't()edon't()fgdo()abcdo()",
// 			want:  "ado()bcabcdo()",
// 		},
// 		{
// 			input: "don't()do()",
// 			want:  "",
// 		},
// 		{
// 			input: "don't()ado()bdon't()",
// 			want:  "b",
// 		},
// 		{
// 			input: "abcdo()",
// 			want:  "abcdo()",
// 		},
// 		{
// 			input: "don't()abc",
// 			want:  "",
// 		},
// 		{
// 			input: "do()abc",
// 			want:  "do()abc",
// 		},
// 		{
// 			input: "select()} <*mul(843,597)!~mul(717,524)&?}'mul(928,721)>mul(194,52)'why()]-*select()what(898,458):#*mul(31,582)mul(209,470)'-mul(834,167)>}mul(188,914)where(344,689)select(90,321)where()-when()[{]mul(133,940)#-mul(732,657)why()$when()-how()?!>who(208,16)mul(332,604)?do()^:how():: mul(613,614)@);mul(458,93)@# @$~+#select(234,120)mul(184,924)~ why()^$' #{mul(882,83)~<&mul(878,13)select()who()?mul(947,828)select()select())?>^where()how()'*mul(837,672)-select():from())[^$^mul(325,979)}what()[select()mul(226,766)?where()'what():mul(964,728)/)~$~how()&]%]mul(420,781)$how()/]! ~]?mul(624,205)-how()<(*where()where()mul(33,801)mul(614,925)/mul(660,91)}/mul(205,753)$- }mul(136,460)where(671,425)from()@mul(316,262)*?'%(mul(290,988)why()<+[%from()/@>mul(694,616)what()from()what()#$where()mul(666,785)mul(668,125)!%/&*(~*mul(728,226)why() %/{'>how(615,392):mul(478,823)>do()mul(166,999)%~!!mul(368,199)how()!who()/<-[mul(43,454)how(673,426))do()when()$from()why()^mul(228,279)how():select()mul(413,456)where(){$-[where()who()mul(769,535)mul(290,859)!#do();who()&+%mul(734,882)mul(346,904)/mul(319,21)<mul(548,912)mul(776,306)/'when() -^-!/mul(805,224)who(),~{when()[mul(641,941)+*why()mul(880,39)!mul(367,322)+{;why())<where():+ do():who()who(){mul(297,155){>/@!}[(: mul(582,935);mul(924,453)where():]*]mul(962,327):?<where(583,199)/[{)how()don't()where()^mul(493,697)'%+#[!mul(64,397)*^;)how()}!mul(133,628):-&@%?&;when()(mul(245,450){who()*mul(8,731)!([&mul(862,150)who()why()who()'mul(129,583)@:why()[ mul(557,879)why()%,@mul(969,31)^*mul(243,857)#]'?^~mul(418,611)@who()[-mul(381,404)#? who()}}how(),$mul(992,433)#[~from()%mul(823,119)who();when()from()mul(436,64)when()<{when()how(141,705)$<{select(673,665)@mul(776,698)mul(859,821)where(733,36)}>$who()who()<~don't();mul(531,594) ^>&mul(246,430)select()how() %]/mul(80)')][,+-select()-mul(18,784){mul(360,42)) ~who()@+mul(113,212)/*(/mul(765,643)!mul(853,147)[/ mul(396,209)+[?what()mul(479,669))why();;mul(542,614);#}$select()mul(398,910)mul(687,370)mul(59,590)'&what(656,317)(what()*/+()mul(382,325)#:]when()%<,,where()<mul(124,796)what()$/,<@)mul(227,847);mul(588,764)<@select(),who()where()mul(496,225)where()how()<!mul(998,360)+-^){?where()<mul(584,368)why()why()why()>mul(607,933)[what()*what()mul(301,569)where()%{when();)[*%mul(223,95)[select()-why()where()from()?:[<mul(927,19)] *:^'mul(846,824)'~(&what(206,282)[,mul(791]? -&!mul(482,335)mul(835,793);/@mul(115,602)what()!:what()!&]where()>mul(718,286):}why(),mul(2,974) :/where()?! }what()when()mul(379,171)+]select()][{mul(530,485)>*-why()},-how()!who()mul(643,906)}who()mul(906,628)<;{mul(875,497)%;!#^{!how()(mul(29,450)*$how()+from()what():mul(298,289)-how()*mul(771,685)who(44,541)!?when()who()#{mul(87,962)]#mul(479,616)",
// 			want:  "select()} <*mul(843,597)!~mul(717,524)&?}'mul(928,721)>mul(194,52)'why()]-*select()what(898,458):#*mul(31,582)mul(209,470)'-mul(834,167)>}mul(188,914)where(344,689)select(90,321)where()-when()[{]mul(133,940)#-mul(732,657)why()$when()-how()?!>who(208,16)mul(332,604)?do()^:how():: mul(613,614)@);mul(458,93)@# @$~+#select(234,120)mul(184,924)~ why()^$' #{mul(882,83)~<&mul(878,13)select()who()?mul(947,828)select()select())?>^where()how()'*mul(837,672)-select():from())[^$^mul(325,979)}what()[select()mul(226,766)?where()'what():mul(964,728)/)~$~how()&]%]mul(420,781)$how()/]! ~]?mul(624,205)-how()<(*where()where()mul(33,801)mul(614,925)/mul(660,91)}/mul(205,753)$- }mul(136,460)where(671,425)from()@mul(316,262)*?'%(mul(290,988)why()<+[%from()/@>mul(694,616)what()from()what()#$where()mul(666,785)mul(668,125)!%/&*(~*mul(728,226)why() %/{'>how(615,392):mul(478,823)>do()mul(166,999)%~!!mul(368,199)how()!who()/<-[mul(43,454)how(673,426))do()when()$from()why()^mul(228,279)how():select()mul(413,456)where(){$-[where()who()mul(769,535)mul(290,859)!#do();who()&+%mul(734,882)mul(346,904)/mul(319,21)<mul(548,912)mul(776,306)/'when() -^-!/mul(805,224)who(),~{when()[mul(641,941)+*why()mul(880,39)!mul(367,322)+{;why())<where():+ do():who()who(){mul(297,155){>/@!}[(: mul(582,935);mul(924,453)where():]*]mul(962,327):?<where(583,199)/[{)how()",
// 		},
// 		{
// 			input: "mul(604,576)mul%}}how(153,807);@ what()(mul(273,600)!]mul(106,99) mul(461,886)(mul(121what()]'*@+;!mul(513,885)'why()(how()mul(830,191)(>(where()%how()when()what()mul(562,733)^*';:mul(21,307)@what()select()where()~ select()mul(789,818)]mul(11,673)mul(194,572)$#%[/'</ #mul(10who()^!>&}(mul(162,864);{{mul(548,916)(}><+;(}mul(325,72);mul(722,66)}what()mul(703,168),;where()^,mul(530,109)why()where()~from()mul(471,436)mulwhat()(from()~*why()mul(177,943)[select()when()'<?mul(229,627)&%what()[/{~'how(582,475)mul(56,986)mul(999,466),$where()how()select() ~$]select()mul(576,749)*$who()mul(847,95)mul(702,555)*]@when(),how()>!mul(734,260)( who()+-select()>:*mul(63,684)where()#{,)&mul(531,571)[~$where()^({<mul(502,674)mul(490,264)~why()from()select() mul(540,855):+@ select();,,do()}why()how()select(73,137)mul(268,58)-(where()?/'mul(741,485)when()?select()-)where()*<%don't()#mul? ~mul(990,88)$/mul(986,722)?;from()where(394,512)how(593,980)&+mul(667,464);mul(390,181)]who()what()*}+#;mul(429,936)'<:-'+^}mul(346,607-/where()select()%><who()mul(914,888)mul(781,920):*mul(954,791))({who()?~$)mul(769,183)#['{^)*-mul(330,184)select()%what()'who()where()when()(mul(988,148)&who()*&/[/'mul(327,74)/+;?/select()mul(315,381)don't()>}/-#~why()!*mul(721,722),>why() mul(583,596)when()%}$mul(482,164)$mul(230,264)mul(752,60)@'mul(47,57)(mul(17,292)select()where()%)>@why()mul(302,101)<<how()!-[!from()]mul(649,528)^]@^$mul(577,114)/<mul(579,480)who(808,216), #when()why()?mul(979,878)$+why(712,413)$mul:-from()%&}select()mul(539,991)$?>?when()how(761,642)from()mul(501,428)}#how()'what()/+{<+mul(793,630){[$;:[!:what()%don't()what()>from(825,940)' </^&mul(212,563)~when()mul(943,607)where(802,717)'[when()<who())&how()mul(696,659)mul(156//,(mul(477,156)*when(952,865)#*?!>%>!do()when()$<(mul(351,5);-<;,@!why()mul(234,498)who()when()]^/what()^how()do();$]&mul(120,466)*? -#{mul(815,705)why()mul+''*:<select()mul(363,529)# :(,!do()how()[$select()mul(959,174)]$ ^^&(} what()mul(57,223)]!][mul(887,199)who()! &<mul(344,455)% ^&^&mul#!where(431,184) select() do() !^(:+;!mul(342,49)&who(),mul(265,288)#(^mul(360,144)>mul(82,491)where()(don't()^ when()mul(609,399){:#who()$when()from()+mul(506,128)mul(930,401)*%$?mul(140,278)>who()mul(414,2)(-where()!mul(536,354)when()]select()@**mul(678,285){-+select()who(786,299)>$}mul(271,42)][-,[mul(94,592when(870,228)^'~{what()mul(760,390)#>*from()'mul(944,912)why()@#}&/'why(868,249)mul(538,835']),(<~'mul(50,117)}<mul(877,572)when()mul(228,855)}(when()*mul(633,680)/ mul(570,855)[what()mul(807,384)-mul(637,824)}!{who()?$/%mul,+;/how()when()how()mul(907,786)*[&!]~how()mul(580,380){*what()~~(:mul[{:! mul(801,883)}why():'!,}:'mul(419,843)&do()+'}select(),}{,,{mul(294,229)}when(),mul(967,538)/%;who()>*[what(752,345)how()mul(443,685)!mul(278,194)^mul(490,365)mul(197,282)why() what(444,573)from()!+from()how()mul(303,883))how()>mul(999,385)mul(226,867)$mul(726,25)!mul(630,295)when()mul(154,746)<+>when()how(439,850)$,/&mul(990,848)",
// 			want:  "mul(604,576)mul%}}how(153,807);@ what()(mul(273,600)!]mul(106,99) mul(461,886)(mul(121what()]'*@+;!mul(513,885)'why()(how()mul(830,191)(>(where()%how()when()what()mul(562,733)^*';:mul(21,307)@what()select()where()~ select()mul(789,818)]mul(11,673)mul(194,572)$#%[/'</ #mul(10who()^!>&}(mul(162,864);{{mul(548,916)(}><+;(}mul(325,72);mul(722,66)}what()mul(703,168),;where()^,mul(530,109)why()where()~from()mul(471,436)mulwhat()(from()~*why()mul(177,943)[select()when()'<?mul(229,627)&%what()[/{~'how(582,475)mul(56,986)mul(999,466),$where()how()select() ~$]select()mul(576,749)*$who()mul(847,95)mul(702,555)*]@when(),how()>!mul(734,260)( who()+-select()>:*mul(63,684)where()#{,)&mul(531,571)[~$where()^({<mul(502,674)mul(490,264)~why()from()select() mul(540,855):+@ select();,,do()}why()how()select(73,137)mul(268,58)-(where()?/'mul(741,485)when()?select()-)where()*<%when()$<(mul(351,5);-<;,@!why()mul(234,498)who()when()]^/what()^how()do();$]&mul(120,466)*? -#{mul(815,705)why()mul+''*:<select()mul(363,529)# :(,!do()how()[$select()mul(959,174)]$ ^^&(} what()mul(57,223)]!][mul(887,199)who()! &<mul(344,455)% ^&^&mul#!where(431,184) select() do() !^(:+;!mul(342,49)&who(),mul(265,288)#(^mul(360,144)>mul(82,491)where()(+'}select(),}{,,{mul(294,229)}when(),mul(967,538)/%;who()>*[what(752,345)how()mul(443,685)!mul(278,194)^mul(490,365)mul(197,282)why() what(444,573)from()!+from()how()mul(303,883))how()>mul(999,385)mul(226,867)$mul(726,25)!mul(630,295)when()mul(154,746)<+>when()how(439,850)$,/&mul(990,848)",
// 		},
// 	} {
// 		got := removeDeadInstructions(tc.input)

// 		if got != tc.want {
// 			t.Errorf("got\n%s, want\n%s", got, tc.want)
// 		}
// 	}
// }

func TestPart2(t *testing.T) {
	for _, tc := range []struct {
		lines   []string
		wantSum int
	}{
		{
			lines:   []string{"mul(2,5)don't()", "mul(1,2)do()mul(5,1)"},
			wantSum: 15,
		},
	} {
		gotSum := part2(tc.lines)
		if gotSum != tc.wantSum {
			t.Errorf("got %d, want %d", gotSum, tc.wantSum)
		}
	}
}
