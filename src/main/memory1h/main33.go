package main

import (
	"fmt"
	"math/rand"
	"net/http"
	_ "net/http/pprof" // 导入 pprof 包以启用 pprof 路由
	"reflect"
	"sync"
	"time"
)

const (
	numTurbines = 10
	numRecords  = 3600
)

// WindTurbineData33 结构体定义
type WindTurbineData33 struct {
	Timestamp time.Time
	MA001     bool
	MA002     bool
	MA003     bool
	MA004     bool
	MA005     bool
	MA006     bool
	MA007     bool
	MA008     bool
	MA009     bool
	MA010     bool
	MA011     bool
	MA012     bool
	MA013     bool
	MA014     bool
	MA015     bool
	MA016     bool
	MA017     bool
	MA018     bool
	MA019     bool
	MA020     bool
	MA021     bool
	MA022     bool
	MA023     bool
	MA024     bool
	MA025     bool
	MA026     bool
	MA027     bool
	MA028     bool
	MA029     bool
	MA030     bool
	MA031     bool
	MA032     bool
	MA033     bool
	MA034     bool
	MA035     bool
	MA036     bool
	MA037     bool
	MA038     bool
	MA039     bool
	MA040     bool
	MA041     bool
	MA042     bool
	MA043     bool
	MA044     bool
	MA045     bool
	MA046     bool
	MA047     bool
	MA048     bool
	MA049     bool
	MA050     bool
	MA051     bool
	MA052     bool
	MA053     bool
	MA054     bool
	MA055     bool
	MA056     bool
	MA057     bool
	MA058     bool
	MA059     bool
	MA060     bool
	MA061     bool
	MA062     bool
	MA063     bool
	MA064     bool
	MA065     bool
	MA066     bool
	MA067     bool
	MA068     bool
	MA069     bool
	MA070     bool
	MA071     bool
	MA072     bool
	MA073     bool
	MA074     bool
	MA075     bool
	MA076     bool
	MA077     bool
	MA078     bool
	MA079     bool
	MA080     bool
	MA081     bool
	MA082     bool
	MA083     bool
	MA084     bool
	MA085     bool
	MA086     bool
	MA087     bool
	MA088     bool
	MA089     bool
	MA090     bool
	MA091     bool
	MA092     bool
	MA093     bool
	MA094     bool
	MA095     bool
	MA096     bool
	MA097     bool
	MA098     bool
	MA099     bool
	MA100     bool
	MA101     bool
	MA102     bool
	MA103     bool
	MA104     bool
	MA105     bool
	MA106     bool
	MA107     bool
	MA108     bool
	MA109     bool
	MA110     bool
	MA111     bool
	MA112     bool
	MA113     bool
	MA114     bool
	MA115     bool
	MA116     bool
	MA117     bool
	MA118     bool
	MA119     bool
	MA120     bool
	MA121     bool
	MA122     bool
	MA123     bool
	MA124     bool
	MA125     bool
	MA126     bool
	MA127     bool
	MA128     bool
	MA129     bool
	MA130     bool
	MA131     bool
	MA132     bool
	MA133     bool
	MA134     bool
	MA135     bool
	MA136     bool
	MA137     bool
	MA138     bool
	MA139     bool
	MA140     bool
	MA141     bool
	MA142     bool
	MA143     bool
	MA144     bool
	MA145     bool
	MA146     bool
	MA147     bool
	MA148     bool
	MA149     bool
	MA150     bool
	MA151     bool
	MA152     bool
	MA153     bool
	MA154     bool
	MA155     bool
	MA156     bool
	MA157     bool
	MA158     bool
	MA159     bool
	MA160     bool
	MA161     bool
	MA162     bool
	MA163     bool
	MA164     bool
	MA165     bool
	MA166     bool
	MA167     bool
	MA168     bool
	MA169     bool
	MA170     bool
	MA171     bool
	MA172     bool
	MA173     bool
	MA174     bool
	MA175     bool
	MA176     bool
	MA177     bool
	MA178     bool
	MA179     bool
	MA180     bool
	MA181     bool
	MA182     bool
	MA183     bool
	MA184     bool
	MA185     bool
	MA186     bool
	MA187     bool
	MA188     bool
	MA189     bool
	MA190     bool
	MA191     bool
	MA192     bool
	MA193     bool
	MA194     bool
	MA195     bool
	MA196     bool
	MA197     bool
	MA198     bool
	MA199     bool
	MA200     bool
	MA201     bool
	MA202     bool
	MA203     bool
	MA204     bool
	MA205     bool
	MA206     bool
	MA207     bool
	MA208     bool
	MA209     bool
	MA210     bool
	MA211     bool
	MA212     bool
	MA213     bool
	MA214     bool
	MA215     bool
	MA216     bool
	MA217     bool
	MA218     bool
	MA219     bool
	MA220     bool
	MA221     bool
	MA222     bool
	MA223     bool
	MA224     bool
	MA225     bool
	MA226     bool
	MA227     bool
	MA228     bool
	MA229     bool
	MA230     bool
	MA231     bool
	MA232     bool
	MA233     bool
	MA234     bool
	MA235     bool
	MA236     bool
	MA237     bool
	MA238     bool
	MA239     bool
	MA240     bool
	MA241     bool
	MA242     bool
	MA243     bool
	MA244     bool
	MA245     bool
	MA246     bool
	MA247     bool
	MA248     bool
	MA249     bool
	MA250     bool
	MA251     bool
	MA252     bool
	MA253     bool
	MA254     bool
	MA255     bool
	MA256     bool
	MA257     bool
	MA258     bool
	MA259     bool
	MA260     bool
	MA261     bool
	MA262     bool
	MA263     bool
	MA264     bool
	MA265     bool
	MA266     bool
	MA267     bool
	MA268     bool
	MA269     bool
	MA270     bool
	MA271     bool
	MA272     bool
	MA273     bool
	MA274     bool
	MA275     bool
	MA276     bool
	MA277     bool
	MA278     bool
	MA279     bool
	MA280     bool
	MA281     bool
	MA282     bool
	MA283     bool
	MA284     bool
	MA285     bool
	MA286     bool
	MA287     bool
	MA288     bool
	MA289     bool
	MA290     bool
	MA291     bool
	MA292     bool
	MA293     bool
	MA294     bool
	MA295     bool
	MA296     bool
	MA297     bool
	MA298     bool
	MA299     bool
	MA300     bool
	MA301     bool
	MA302     bool
	MA303     bool
	MA304     bool
	MA305     bool
	MA306     bool
	MA307     bool
	MA308     bool
	MA309     bool
	MA310     bool
	MA311     bool
	MA312     bool
	MA313     bool
	MA314     bool
	MA315     bool
	MA316     bool
	MA317     bool
	MA318     bool
	MA319     bool
	MA320     bool
	MA321     bool
	MA322     bool
	MA323     bool
	MA324     bool
	MA325     bool
	MA326     bool
	MA327     bool
	MA328     bool
	MA329     bool
	MA330     bool
	MA331     bool
	MA332     bool
	MA333     bool
	MA334     bool
	MA335     bool
	MA336     bool
	MA337     bool
	MA338     bool
	MA339     bool
	MA340     bool
	MA341     bool
	MA342     bool
	MA343     bool
	MA344     bool
	MA345     bool
	MA346     bool
	MA347     bool
	MA348     bool
	MA349     bool
	MA350     bool
	MA351     bool
	MA352     bool
	MA353     bool
	MA354     bool
	MA355     bool
	MA356     bool
	MA357     bool
	MA358     bool
	MA359     bool
	MA360     bool
	MA361     bool
	MA362     bool
	MA363     bool
	MA364     bool
	MA365     bool
	MA366     bool
	MA367     bool
	MA368     bool
	MA369     bool
	MA370     bool
	MA371     bool
	MA372     bool
	MA373     bool
	MA374     bool
	MA375     bool
	MA376     bool
	MA377     bool
	MA378     bool
	MA379     bool
	MA380     bool
	MA381     bool
	MA382     bool
	MA383     bool
	MA384     bool
	MA385     bool
	MA386     bool
	MA387     bool
	MA388     bool
	MA389     bool
	MA390     bool
	MA391     bool
	MA392     bool
	MA393     bool
	MA394     bool
	MA395     bool
	MA396     bool
	MA397     bool
	MA398     bool
	MA399     bool
	MA400     bool
	MA401     bool
	MA402     bool
	MA403     bool
	MA404     bool
	MA405     bool
	MA406     bool
	MA407     bool
	MA408     bool
	MA409     bool
	MA410     bool
	MA411     bool
	MA412     bool
	MA413     bool
	MA414     bool
	MA415     bool
	MA416     bool
	MA417     bool
	MA418     bool
	MA419     bool
	MA420     bool
	MA421     bool
	MA422     bool
	MA423     bool
	MA424     bool
	MA425     bool
	MA426     bool
	MA427     bool
	MA428     bool
	MA429     bool
	MA430     bool
	MA431     bool
	MA432     bool
	MA433     bool
	MA434     bool
	MA435     bool
	MA436     bool
	MA437     bool
	MA438     bool
	MA439     bool
	MA440     bool
	MA441     bool
	MA442     bool
	MA443     bool
	MA444     bool
	MA445     bool
	MA446     bool
	MA447     bool
	MA448     bool
	MA449     bool
	MA450     bool
	MA451     bool
	MA452     bool
	MA453     bool
	MA454     bool
	MA455     bool
	MA456     bool
	MA457     bool
	MA458     bool
	MA459     bool
	MA460     bool
	MA461     bool
	MA462     bool
	MA463     bool
	MA464     bool
	MA465     bool
	MA466     bool
	MA467     bool
	MA468     bool
	MA469     bool
	MA470     bool
	MA471     bool
	MA472     bool
	MA473     bool
	MA474     bool
	MA475     bool
	MA476     bool
	MA477     bool
	MA478     bool
	MA479     bool
	MA480     bool
	MA481     bool
	MA482     bool
	MA483     bool
	MA484     bool
	MA485     bool
	MA486     bool
	MA487     bool
	MA488     bool
	MA489     bool
	MA490     bool
	MA491     bool
	MA492     bool
	MA493     bool
	MA494     bool
	MA495     bool
	MA496     bool
	MA497     bool
	MA498     bool
	MA499     bool
	MA500     bool
	MA501     bool
	MA502     bool
	MA503     bool
	MA504     bool
	MA505     bool
	MA506     bool
	MA507     bool
	MA508     bool
	MA509     bool
	MA510     bool
	MA511     bool
	MA512     bool
	MA513     bool
	MA514     bool
	MA515     bool
	MA516     bool
	MA517     bool
	MA518     bool
	MA519     bool
	MA520     bool
	MA521     bool
	MA522     bool
	MA523     bool
	MA524     bool
	MA525     bool
	MA526     bool
	MA527     bool
	MA528     bool
	MA529     bool
	MA530     bool
	MA531     bool
	MA532     bool
	MA533     bool
	MA534     bool
	MA535     bool
	MA536     bool
	MA537     bool
	MA538     bool
	MA539     bool
	MA540     bool
	MA541     bool
	MA542     bool
	MA543     bool
	MA544     bool
	MA545     bool
	MA546     bool
	MA547     bool
	MA548     bool
	MA549     bool
	MA550     bool
	MA551     bool
	MA552     bool
	MA553     bool
	MA554     bool
	MA555     bool
	MA556     bool
	MA557     bool
	MA558     bool
	MA559     bool
	MA560     bool
	MA561     bool
	MA562     bool
	MA563     bool
	MA564     bool
	MA565     bool
	MA566     bool
	MA567     bool
	MA568     bool
	MA569     bool
	MA570     bool
	MA571     bool
	MA572     bool
	MA573     bool
	MA574     bool
	MA575     bool
	MA576     bool
	MA577     bool
	MA578     bool
	MA579     bool
	MA580     bool
	MA581     bool
	MA582     bool
	MA583     bool
	MA584     bool
	MA585     bool
	MA586     bool
	MA587     bool
	MA588     bool
	MA589     bool
	MA590     bool
	MA591     bool
	MA592     bool
	MA593     bool
	MA594     bool
	MA595     bool
	MA596     bool
	MA597     bool
	MA598     bool
	MA599     bool
	MA600     bool
	MA601     bool
	MA602     bool
	MA603     bool
	MA604     bool
	MA605     bool
	MA606     bool
	MA607     bool
	MA608     bool
	MA609     bool
	MA610     bool
	MA611     bool
	MA612     bool
	MA613     bool
	MA614     bool
	MA615     bool
	MA616     bool
	MA617     bool
	MA618     bool
	MA619     bool
	MA620     bool
	MA621     bool
	MA622     bool
	MA623     bool
	MA624     bool
	MA625     bool
	MA626     bool
	MA627     bool
	MA628     bool
	MA629     bool
	MA630     bool
	MA631     bool
	MA632     bool
	MA633     bool
	MA634     bool
	MA635     bool
	MA636     bool
	MA637     bool
	MA638     bool
	MA639     bool
	MA640     bool
	MA641     bool
	MA642     bool
	MA643     bool
	MA644     bool
	MA645     bool
	MA646     bool
	MA647     bool
	MA648     bool
	MA649     bool
	MA650     bool
	MA651     bool
	MA652     bool
	MA653     bool
	MA654     bool
	MA655     bool
	MA656     bool
	MA657     bool
	MA658     bool
	MA659     bool
	MA660     bool
	MA661     bool
	MA662     bool
	MA663     bool
	MA664     bool
	MA665     bool
	MA666     bool
	MA667     bool
	MA668     bool
	MA669     bool
	MA670     bool
	MA671     bool
	MA672     bool
	MA673     bool
	MA674     bool
	MA675     bool
	MA676     bool
	MA677     bool
	MA678     bool
	MA679     bool
	MA680     bool
	MA681     bool
	MA682     bool
	MA683     bool
	MA684     bool
	MA685     bool
	MA686     bool
	MA687     bool
	MA688     bool
	MA689     bool
	MA690     bool
	MA691     bool
	MA692     bool
	MA693     bool
	MA694     bool
	MA695     bool
	MA696     bool
	MA697     bool
	MA698     bool
	MA699     bool
	MA700     bool
	MA701     bool
	MA702     bool
	MA703     bool
	MA704     bool
	MA705     bool
	MA706     bool
	MA707     bool
	MA708     bool
	MA709     bool
	MA710     bool
	MA711     bool
	MA712     bool
	MA713     bool
	MA714     bool
	MA715     bool
	MA716     bool
	MA717     bool
	MA718     bool
	MA719     bool
	MA720     bool
	MA721     bool
	MA722     bool
	MA723     bool
	MA724     bool
	MA725     bool
	MA726     bool
	MA727     bool
	MA728     bool
	MA729     bool
	MA730     bool
	MA731     bool
	MA732     bool
	MA733     bool
	MA734     bool
	MA735     bool
	MA736     bool
	MA737     bool
	MA738     bool
	MA739     bool
	MA740     bool
	MA741     bool
	MA742     bool
	MA743     bool
	MA744     bool
	MA745     bool
	MA746     bool
	MA747     bool
	MA748     bool
	MA749     bool
	MA750     bool
	MA751     bool
	MA752     bool
	MA753     bool
	MA754     bool
	MA755     bool
	MA756     bool
	MA757     bool
	MA758     bool
	MA759     bool
	MA760     bool
	MA761     bool
	MA762     bool
	MA763     bool
	MA764     bool
	MA765     bool
	MA766     bool
	MA767     bool
	MA768     bool
	MA769     bool
	MA770     bool
	MA771     bool
	MA772     bool
	MA773     bool
	MA774     bool
	MA775     bool
	MA776     bool
	MA777     bool
	MA778     bool
	MA779     bool
	MA780     bool
	MA781     bool
	MA782     bool
	MA783     bool
	MA784     bool
	MA785     bool
	MA786     bool
	MA787     bool
	MA788     bool
	MA789     bool
	MA790     bool
	MA791     bool
	MA792     bool
	MA793     bool
	MA794     bool
	MA795     bool
	MA796     bool
	MA797     bool
	MA798     bool
	MA799     bool
	MA800     bool
	MA801     bool
	MA802     bool
	MA803     bool
	MA804     bool
	MA805     bool
	MA806     bool
	MA807     bool
	MA808     bool
	MA809     bool
	MA810     bool
	MA811     bool
	MA812     bool
	MA813     bool
	MA814     bool
	MA815     bool
	MA816     bool
	MA817     bool
	MA818     bool
	MA819     bool
	MA820     bool
	MA821     bool
	MA822     bool
	MA823     bool
	MA824     bool
	MA825     bool
	MA826     bool
	MA827     bool
	MA828     bool
	MA829     bool
	MA830     bool
	MA831     bool
	MA832     bool
	MA833     bool
	MA834     bool
	MA835     bool
	MA836     bool
	MA837     bool
	MA838     bool
	MA839     bool
	MA840     bool
	MA841     bool
	MA842     bool
	MA843     bool
	MA844     bool
	MA845     bool
	MA846     bool
	MA847     bool
	MA848     bool
	MA849     bool
	MA850     bool
	MA851     bool
	MA852     bool
	MA853     bool
	MA854     bool
	MA855     bool
	MA856     bool
	MA857     bool
	MA858     bool
	MA859     bool
	MA860     bool
	MA861     bool
	MA862     bool
	MA863     bool
	MA864     bool
	MA865     bool
	MA866     bool
	MA867     bool
	MA868     bool
	MA869     bool
	MA870     bool
	MA871     bool
	MA872     bool
	MA873     bool
	MA874     bool
	MA875     bool
	MA876     bool
	MA877     bool
	MA878     bool
	MA879     bool
	MA880     bool
	MA881     bool
	MA882     bool
	MA883     bool
	MA884     bool
	MA885     bool
	MA886     bool
	MA887     bool
	MA888     bool
	MA889     bool
	MA890     bool
	MA891     bool
	MA892     bool
	MA893     bool
	MA894     bool
	MA895     bool
	MA896     bool
	MA897     bool
	MA898     bool
	MA899     bool
	MA900     bool
	MA901     bool
	MA902     bool
	MA903     bool
	MA904     bool
	MA905     bool
	MA906     bool
	MA907     bool
	MA908     bool
	MA909     bool
	MA910     bool
	MA911     bool
	MA912     bool
	MA913     bool
	MA914     bool
	MA915     bool
	MA916     bool
	MA917     bool
	MA918     bool
	MA919     bool
	MA920     bool
	MA921     bool
	MA922     bool
	MA923     bool
	MA924     bool
	MA925     bool
	MA926     bool
	MA927     bool
	MA928     bool
	MA929     bool
	MA930     bool
	MA931     bool
	MA932     bool
	MA933     bool
	MA934     bool
	MA935     bool
	MA936     bool
	MA937     bool
	MA938     bool
	MA939     bool
	MA940     bool
	MA941     bool
	MA942     bool
	MA943     bool
	MA944     bool
	MA945     bool
	MA946     bool
	MA947     bool
	MA948     bool
	MA949     bool
	MA950     bool
	MA951     bool
	MA952     bool
	MA953     bool
	MA954     bool
	MA955     bool
	MA956     bool
	MA957     bool
	MA958     bool
	MA959     bool
	MA960     bool
	MA961     bool
	MA962     bool
	MA963     bool
	MA964     bool
	MA965     bool
	MA966     bool
	MA967     bool
	MA968     bool
	MA969     bool
	MA970     bool
	MA971     bool
	MA972     bool
	MA973     bool
	MA974     bool
	MA975     bool
	MA976     bool
	MA977     bool
	MA978     bool
	MA979     bool
	MA980     bool
	MA981     bool
	MA982     bool
	MA983     bool
	MA984     bool
	MA985     bool
	MA986     bool
	MA987     bool
	MA988     bool
	MA989     bool
	MA990     bool
	MA991     bool
	MA992     bool
	MA993     bool
	MA994     bool
	MA995     bool
	MA996     bool
	MA997     bool
	MA998     bool
	MA999     bool
	MA1000    bool

	// ... 省略中间的 MA004 - MA998 ...
	// 补全省略的部分

	MC001  float32
	MC002  float32
	MC003  float32
	MC004  float32
	MC005  float32
	MC006  float32
	MC007  float32
	MC008  float32
	MC009  float32
	MC010  float32
	MC011  float32
	MC012  float32
	MC013  float32
	MC014  float32
	MC015  float32
	MC016  float32
	MC017  float32
	MC018  float32
	MC019  float32
	MC020  float32
	MC021  float32
	MC022  float32
	MC023  float32
	MC024  float32
	MC025  float32
	MC026  float32
	MC027  float32
	MC028  float32
	MC029  float32
	MC030  float32
	MC031  float32
	MC032  float32
	MC033  float32
	MC034  float32
	MC035  float32
	MC036  float32
	MC037  float32
	MC038  float32
	MC039  float32
	MC040  float32
	MC041  float32
	MC042  float32
	MC043  float32
	MC044  float32
	MC045  float32
	MC046  float32
	MC047  float32
	MC048  float32
	MC049  float32
	MC050  float32
	MC051  float32
	MC052  float32
	MC053  float32
	MC054  float32
	MC055  float32
	MC056  float32
	MC057  float32
	MC058  float32
	MC059  float32
	MC060  float32
	MC061  float32
	MC062  float32
	MC063  float32
	MC064  float32
	MC065  float32
	MC066  float32
	MC067  float32
	MC068  float32
	MC069  float32
	MC070  float32
	MC071  float32
	MC072  float32
	MC073  float32
	MC074  float32
	MC075  float32
	MC076  float32
	MC077  float32
	MC078  float32
	MC079  float32
	MC080  float32
	MC081  float32
	MC082  float32
	MC083  float32
	MC084  float32
	MC085  float32
	MC086  float32
	MC087  float32
	MC088  float32
	MC089  float32
	MC090  float32
	MC091  float32
	MC092  float32
	MC093  float32
	MC094  float32
	MC095  float32
	MC096  float32
	MC097  float32
	MC098  float32
	MC099  float32
	MC100  float32
	MC101  float32
	MC102  float32
	MC103  float32
	MC104  float32
	MC105  float32
	MC106  float32
	MC107  float32
	MC108  float32
	MC109  float32
	MC110  float32
	MC111  float32
	MC112  float32
	MC113  float32
	MC114  float32
	MC115  float32
	MC116  float32
	MC117  float32
	MC118  float32
	MC119  float32
	MC120  float32
	MC121  float32
	MC122  float32
	MC123  float32
	MC124  float32
	MC125  float32
	MC126  float32
	MC127  float32
	MC128  float32
	MC129  float32
	MC130  float32
	MC131  float32
	MC132  float32
	MC133  float32
	MC134  float32
	MC135  float32
	MC136  float32
	MC137  float32
	MC138  float32
	MC139  float32
	MC140  float32
	MC141  float32
	MC142  float32
	MC143  float32
	MC144  float32
	MC145  float32
	MC146  float32
	MC147  float32
	MC148  float32
	MC149  float32
	MC150  float32
	MC151  float32
	MC152  float32
	MC153  float32
	MC154  float32
	MC155  float32
	MC156  float32
	MC157  float32
	MC158  float32
	MC159  float32
	MC160  float32
	MC161  float32
	MC162  float32
	MC163  float32
	MC164  float32
	MC165  float32
	MC166  float32
	MC167  float32
	MC168  float32
	MC169  float32
	MC170  float32
	MC171  float32
	MC172  float32
	MC173  float32
	MC174  float32
	MC175  float32
	MC176  float32
	MC177  float32
	MC178  float32
	MC179  float32
	MC180  float32
	MC181  float32
	MC182  float32
	MC183  float32
	MC184  float32
	MC185  float32
	MC186  float32
	MC187  float32
	MC188  float32
	MC189  float32
	MC190  float32
	MC191  float32
	MC192  float32
	MC193  float32
	MC194  float32
	MC195  float32
	MC196  float32
	MC197  float32
	MC198  float32
	MC199  float32
	MC200  float32
	MC201  float32
	MC202  float32
	MC203  float32
	MC204  float32
	MC205  float32
	MC206  float32
	MC207  float32
	MC208  float32
	MC209  float32
	MC210  float32
	MC211  float32
	MC212  float32
	MC213  float32
	MC214  float32
	MC215  float32
	MC216  float32
	MC217  float32
	MC218  float32
	MC219  float32
	MC220  float32
	MC221  float32
	MC222  float32
	MC223  float32
	MC224  float32
	MC225  float32
	MC226  float32
	MC227  float32
	MC228  float32
	MC229  float32
	MC230  float32
	MC231  float32
	MC232  float32
	MC233  float32
	MC234  float32
	MC235  float32
	MC236  float32
	MC237  float32
	MC238  float32
	MC239  float32
	MC240  float32
	MC241  float32
	MC242  float32
	MC243  float32
	MC244  float32
	MC245  float32
	MC246  float32
	MC247  float32
	MC248  float32
	MC249  float32
	MC250  float32
	MC251  float32
	MC252  float32
	MC253  float32
	MC254  float32
	MC255  float32
	MC256  float32
	MC257  float32
	MC258  float32
	MC259  float32
	MC260  float32
	MC261  float32
	MC262  float32
	MC263  float32
	MC264  float32
	MC265  float32
	MC266  float32
	MC267  float32
	MC268  float32
	MC269  float32
	MC270  float32
	MC271  float32
	MC272  float32
	MC273  float32
	MC274  float32
	MC275  float32
	MC276  float32
	MC277  float32
	MC278  float32
	MC279  float32
	MC280  float32
	MC281  float32
	MC282  float32
	MC283  float32
	MC284  float32
	MC285  float32
	MC286  float32
	MC287  float32
	MC288  float32
	MC289  float32
	MC290  float32
	MC291  float32
	MC292  float32
	MC293  float32
	MC294  float32
	MC295  float32
	MC296  float32
	MC297  float32
	MC298  float32
	MC299  float32
	MC300  float32
	MC301  float32
	MC302  float32
	MC303  float32
	MC304  float32
	MC305  float32
	MC306  float32
	MC307  float32
	MC308  float32
	MC309  float32
	MC310  float32
	MC311  float32
	MC312  float32
	MC313  float32
	MC314  float32
	MC315  float32
	MC316  float32
	MC317  float32
	MC318  float32
	MC319  float32
	MC320  float32
	MC321  float32
	MC322  float32
	MC323  float32
	MC324  float32
	MC325  float32
	MC326  float32
	MC327  float32
	MC328  float32
	MC329  float32
	MC330  float32
	MC331  float32
	MC332  float32
	MC333  float32
	MC334  float32
	MC335  float32
	MC336  float32
	MC337  float32
	MC338  float32
	MC339  float32
	MC340  float32
	MC341  float32
	MC342  float32
	MC343  float32
	MC344  float32
	MC345  float32
	MC346  float32
	MC347  float32
	MC348  float32
	MC349  float32
	MC350  float32
	MC351  float32
	MC352  float32
	MC353  float32
	MC354  float32
	MC355  float32
	MC356  float32
	MC357  float32
	MC358  float32
	MC359  float32
	MC360  float32
	MC361  float32
	MC362  float32
	MC363  float32
	MC364  float32
	MC365  float32
	MC366  float32
	MC367  float32
	MC368  float32
	MC369  float32
	MC370  float32
	MC371  float32
	MC372  float32
	MC373  float32
	MC374  float32
	MC375  float32
	MC376  float32
	MC377  float32
	MC378  float32
	MC379  float32
	MC380  float32
	MC381  float32
	MC382  float32
	MC383  float32
	MC384  float32
	MC385  float32
	MC386  float32
	MC387  float32
	MC388  float32
	MC389  float32
	MC390  float32
	MC391  float32
	MC392  float32
	MC393  float32
	MC394  float32
	MC395  float32
	MC396  float32
	MC397  float32
	MC398  float32
	MC399  float32
	MC400  float32
	MC401  float32
	MC402  float32
	MC403  float32
	MC404  float32
	MC405  float32
	MC406  float32
	MC407  float32
	MC408  float32
	MC409  float32
	MC410  float32
	MC411  float32
	MC412  float32
	MC413  float32
	MC414  float32
	MC415  float32
	MC416  float32
	MC417  float32
	MC418  float32
	MC419  float32
	MC420  float32
	MC421  float32
	MC422  float32
	MC423  float32
	MC424  float32
	MC425  float32
	MC426  float32
	MC427  float32
	MC428  float32
	MC429  float32
	MC430  float32
	MC431  float32
	MC432  float32
	MC433  float32
	MC434  float32
	MC435  float32
	MC436  float32
	MC437  float32
	MC438  float32
	MC439  float32
	MC440  float32
	MC441  float32
	MC442  float32
	MC443  float32
	MC444  float32
	MC445  float32
	MC446  float32
	MC447  float32
	MC448  float32
	MC449  float32
	MC450  float32
	MC451  float32
	MC452  float32
	MC453  float32
	MC454  float32
	MC455  float32
	MC456  float32
	MC457  float32
	MC458  float32
	MC459  float32
	MC460  float32
	MC461  float32
	MC462  float32
	MC463  float32
	MC464  float32
	MC465  float32
	MC466  float32
	MC467  float32
	MC468  float32
	MC469  float32
	MC470  float32
	MC471  float32
	MC472  float32
	MC473  float32
	MC474  float32
	MC475  float32
	MC476  float32
	MC477  float32
	MC478  float32
	MC479  float32
	MC480  float32
	MC481  float32
	MC482  float32
	MC483  float32
	MC484  float32
	MC485  float32
	MC486  float32
	MC487  float32
	MC488  float32
	MC489  float32
	MC490  float32
	MC491  float32
	MC492  float32
	MC493  float32
	MC494  float32
	MC495  float32
	MC496  float32
	MC497  float32
	MC498  float32
	MC499  float32
	MC500  float32
	MC501  float32
	MC502  float32
	MC503  float32
	MC504  float32
	MC505  float32
	MC506  float32
	MC507  float32
	MC508  float32
	MC509  float32
	MC510  float32
	MC511  float32
	MC512  float32
	MC513  float32
	MC514  float32
	MC515  float32
	MC516  float32
	MC517  float32
	MC518  float32
	MC519  float32
	MC520  float32
	MC521  float32
	MC522  float32
	MC523  float32
	MC524  float32
	MC525  float32
	MC526  float32
	MC527  float32
	MC528  float32
	MC529  float32
	MC530  float32
	MC531  float32
	MC532  float32
	MC533  float32
	MC534  float32
	MC535  float32
	MC536  float32
	MC537  float32
	MC538  float32
	MC539  float32
	MC540  float32
	MC541  float32
	MC542  float32
	MC543  float32
	MC544  float32
	MC545  float32
	MC546  float32
	MC547  float32
	MC548  float32
	MC549  float32
	MC550  float32
	MC551  float32
	MC552  float32
	MC553  float32
	MC554  float32
	MC555  float32
	MC556  float32
	MC557  float32
	MC558  float32
	MC559  float32
	MC560  float32
	MC561  float32
	MC562  float32
	MC563  float32
	MC564  float32
	MC565  float32
	MC566  float32
	MC567  float32
	MC568  float32
	MC569  float32
	MC570  float32
	MC571  float32
	MC572  float32
	MC573  float32
	MC574  float32
	MC575  float32
	MC576  float32
	MC577  float32
	MC578  float32
	MC579  float32
	MC580  float32
	MC581  float32
	MC582  float32
	MC583  float32
	MC584  float32
	MC585  float32
	MC586  float32
	MC587  float32
	MC588  float32
	MC589  float32
	MC590  float32
	MC591  float32
	MC592  float32
	MC593  float32
	MC594  float32
	MC595  float32
	MC596  float32
	MC597  float32
	MC598  float32
	MC599  float32
	MC600  float32
	MC601  float32
	MC602  float32
	MC603  float32
	MC604  float32
	MC605  float32
	MC606  float32
	MC607  float32
	MC608  float32
	MC609  float32
	MC610  float32
	MC611  float32
	MC612  float32
	MC613  float32
	MC614  float32
	MC615  float32
	MC616  float32
	MC617  float32
	MC618  float32
	MC619  float32
	MC620  float32
	MC621  float32
	MC622  float32
	MC623  float32
	MC624  float32
	MC625  float32
	MC626  float32
	MC627  float32
	MC628  float32
	MC629  float32
	MC630  float32
	MC631  float32
	MC632  float32
	MC633  float32
	MC634  float32
	MC635  float32
	MC636  float32
	MC637  float32
	MC638  float32
	MC639  float32
	MC640  float32
	MC641  float32
	MC642  float32
	MC643  float32
	MC644  float32
	MC645  float32
	MC646  float32
	MC647  float32
	MC648  float32
	MC649  float32
	MC650  float32
	MC651  float32
	MC652  float32
	MC653  float32
	MC654  float32
	MC655  float32
	MC656  float32
	MC657  float32
	MC658  float32
	MC659  float32
	MC660  float32
	MC661  float32
	MC662  float32
	MC663  float32
	MC664  float32
	MC665  float32
	MC666  float32
	MC667  float32
	MC668  float32
	MC669  float32
	MC670  float32
	MC671  float32
	MC672  float32
	MC673  float32
	MC674  float32
	MC675  float32
	MC676  float32
	MC677  float32
	MC678  float32
	MC679  float32
	MC680  float32
	MC681  float32
	MC682  float32
	MC683  float32
	MC684  float32
	MC685  float32
	MC686  float32
	MC687  float32
	MC688  float32
	MC689  float32
	MC690  float32
	MC691  float32
	MC692  float32
	MC693  float32
	MC694  float32
	MC695  float32
	MC696  float32
	MC697  float32
	MC698  float32
	MC699  float32
	MC700  float32
	MC701  float32
	MC702  float32
	MC703  float32
	MC704  float32
	MC705  float32
	MC706  float32
	MC707  float32
	MC708  float32
	MC709  float32
	MC710  float32
	MC711  float32
	MC712  float32
	MC713  float32
	MC714  float32
	MC715  float32
	MC716  float32
	MC717  float32
	MC718  float32
	MC719  float32
	MC720  float32
	MC721  float32
	MC722  float32
	MC723  float32
	MC724  float32
	MC725  float32
	MC726  float32
	MC727  float32
	MC728  float32
	MC729  float32
	MC730  float32
	MC731  float32
	MC732  float32
	MC733  float32
	MC734  float32
	MC735  float32
	MC736  float32
	MC737  float32
	MC738  float32
	MC739  float32
	MC740  float32
	MC741  float32
	MC742  float32
	MC743  float32
	MC744  float32
	MC745  float32
	MC746  float32
	MC747  float32
	MC748  float32
	MC749  float32
	MC750  float32
	MC751  float32
	MC752  float32
	MC753  float32
	MC754  float32
	MC755  float32
	MC756  float32
	MC757  float32
	MC758  float32
	MC759  float32
	MC760  float32
	MC761  float32
	MC762  float32
	MC763  float32
	MC764  float32
	MC765  float32
	MC766  float32
	MC767  float32
	MC768  float32
	MC769  float32
	MC770  float32
	MC771  float32
	MC772  float32
	MC773  float32
	MC774  float32
	MC775  float32
	MC776  float32
	MC777  float32
	MC778  float32
	MC779  float32
	MC780  float32
	MC781  float32
	MC782  float32
	MC783  float32
	MC784  float32
	MC785  float32
	MC786  float32
	MC787  float32
	MC788  float32
	MC789  float32
	MC790  float32
	MC791  float32
	MC792  float32
	MC793  float32
	MC794  float32
	MC795  float32
	MC796  float32
	MC797  float32
	MC798  float32
	MC799  float32
	MC800  float32
	MC801  float32
	MC802  float32
	MC803  float32
	MC804  float32
	MC805  float32
	MC806  float32
	MC807  float32
	MC808  float32
	MC809  float32
	MC810  float32
	MC811  float32
	MC812  float32
	MC813  float32
	MC814  float32
	MC815  float32
	MC816  float32
	MC817  float32
	MC818  float32
	MC819  float32
	MC820  float32
	MC821  float32
	MC822  float32
	MC823  float32
	MC824  float32
	MC825  float32
	MC826  float32
	MC827  float32
	MC828  float32
	MC829  float32
	MC830  float32
	MC831  float32
	MC832  float32
	MC833  float32
	MC834  float32
	MC835  float32
	MC836  float32
	MC837  float32
	MC838  float32
	MC839  float32
	MC840  float32
	MC841  float32
	MC842  float32
	MC843  float32
	MC844  float32
	MC845  float32
	MC846  float32
	MC847  float32
	MC848  float32
	MC849  float32
	MC850  float32
	MC851  float32
	MC852  float32
	MC853  float32
	MC854  float32
	MC855  float32
	MC856  float32
	MC857  float32
	MC858  float32
	MC859  float32
	MC860  float32
	MC861  float32
	MC862  float32
	MC863  float32
	MC864  float32
	MC865  float32
	MC866  float32
	MC867  float32
	MC868  float32
	MC869  float32
	MC870  float32
	MC871  float32
	MC872  float32
	MC873  float32
	MC874  float32
	MC875  float32
	MC876  float32
	MC877  float32
	MC878  float32
	MC879  float32
	MC880  float32
	MC881  float32
	MC882  float32
	MC883  float32
	MC884  float32
	MC885  float32
	MC886  float32
	MC887  float32
	MC888  float32
	MC889  float32
	MC890  float32
	MC891  float32
	MC892  float32
	MC893  float32
	MC894  float32
	MC895  float32
	MC896  float32
	MC897  float32
	MC898  float32
	MC899  float32
	MC900  float32
	MC901  float32
	MC902  float32
	MC903  float32
	MC904  float32
	MC905  float32
	MC906  float32
	MC907  float32
	MC908  float32
	MC909  float32
	MC910  float32
	MC911  float32
	MC912  float32
	MC913  float32
	MC914  float32
	MC915  float32
	MC916  float32
	MC917  float32
	MC918  float32
	MC919  float32
	MC920  float32
	MC921  float32
	MC922  float32
	MC923  float32
	MC924  float32
	MC925  float32
	MC926  float32
	MC927  float32
	MC928  float32
	MC929  float32
	MC930  float32
	MC931  float32
	MC932  float32
	MC933  float32
	MC934  float32
	MC935  float32
	MC936  float32
	MC937  float32
	MC938  float32
	MC939  float32
	MC940  float32
	MC941  float32
	MC942  float32
	MC943  float32
	MC944  float32
	MC945  float32
	MC946  float32
	MC947  float32
	MC948  float32
	MC949  float32
	MC950  float32
	MC951  float32
	MC952  float32
	MC953  float32
	MC954  float32
	MC955  float32
	MC956  float32
	MC957  float32
	MC958  float32
	MC959  float32
	MC960  float32
	MC961  float32
	MC962  float32
	MC963  float32
	MC964  float32
	MC965  float32
	MC966  float32
	MC967  float32
	MC968  float32
	MC969  float32
	MC970  float32
	MC971  float32
	MC972  float32
	MC973  float32
	MC974  float32
	MC975  float32
	MC976  float32
	MC977  float32
	MC978  float32
	MC979  float32
	MC980  float32
	MC981  float32
	MC982  float32
	MC983  float32
	MC984  float32
	MC985  float32
	MC986  float32
	MC987  float32
	MC988  float32
	MC989  float32
	MC990  float32
	MC991  float32
	MC992  float32
	MC993  float32
	MC994  float32
	MC995  float32
	MC996  float32
	MC997  float32
	MC998  float32
	MC999  float32
	MC1000 float32
}

// Generate random WindTurbineData33
func generateRandomWindTurbineData33(timestamp time.Time) WindTurbineData33 {
	return WindTurbineData33{
		Timestamp: timestamp,

		MA001:  rand.Intn(2) == 1,
		MA002:  rand.Intn(2) == 2,
		MA003:  rand.Intn(2) == 3,
		MA004:  rand.Intn(2) == 4,
		MA005:  rand.Intn(2) == 5,
		MA006:  rand.Intn(2) == 6,
		MA007:  rand.Intn(2) == 7,
		MA008:  rand.Intn(2) == 8,
		MA009:  rand.Intn(2) == 9,
		MA010:  rand.Intn(2) == 10,
		MA011:  rand.Intn(2) == 11,
		MA012:  rand.Intn(2) == 12,
		MA013:  rand.Intn(2) == 13,
		MA014:  rand.Intn(2) == 14,
		MA015:  rand.Intn(2) == 15,
		MA016:  rand.Intn(2) == 16,
		MA017:  rand.Intn(2) == 17,
		MA018:  rand.Intn(2) == 18,
		MA019:  rand.Intn(2) == 19,
		MA020:  rand.Intn(2) == 20,
		MA021:  rand.Intn(2) == 21,
		MA022:  rand.Intn(2) == 22,
		MA023:  rand.Intn(2) == 23,
		MA024:  rand.Intn(2) == 24,
		MA025:  rand.Intn(2) == 25,
		MA026:  rand.Intn(2) == 26,
		MA027:  rand.Intn(2) == 27,
		MA028:  rand.Intn(2) == 28,
		MA029:  rand.Intn(2) == 29,
		MA030:  rand.Intn(2) == 30,
		MA031:  rand.Intn(2) == 31,
		MA032:  rand.Intn(2) == 32,
		MA033:  rand.Intn(2) == 33,
		MA034:  rand.Intn(2) == 34,
		MA035:  rand.Intn(2) == 35,
		MA036:  rand.Intn(2) == 36,
		MA037:  rand.Intn(2) == 37,
		MA038:  rand.Intn(2) == 38,
		MA039:  rand.Intn(2) == 39,
		MA040:  rand.Intn(2) == 40,
		MA041:  rand.Intn(2) == 41,
		MA042:  rand.Intn(2) == 42,
		MA043:  rand.Intn(2) == 43,
		MA044:  rand.Intn(2) == 44,
		MA045:  rand.Intn(2) == 45,
		MA046:  rand.Intn(2) == 46,
		MA047:  rand.Intn(2) == 47,
		MA048:  rand.Intn(2) == 48,
		MA049:  rand.Intn(2) == 49,
		MA050:  rand.Intn(2) == 50,
		MA051:  rand.Intn(2) == 51,
		MA052:  rand.Intn(2) == 52,
		MA053:  rand.Intn(2) == 53,
		MA054:  rand.Intn(2) == 54,
		MA055:  rand.Intn(2) == 55,
		MA056:  rand.Intn(2) == 56,
		MA057:  rand.Intn(2) == 57,
		MA058:  rand.Intn(2) == 58,
		MA059:  rand.Intn(2) == 59,
		MA060:  rand.Intn(2) == 60,
		MA061:  rand.Intn(2) == 61,
		MA062:  rand.Intn(2) == 62,
		MA063:  rand.Intn(2) == 63,
		MA064:  rand.Intn(2) == 64,
		MA065:  rand.Intn(2) == 65,
		MA066:  rand.Intn(2) == 66,
		MA067:  rand.Intn(2) == 67,
		MA068:  rand.Intn(2) == 68,
		MA069:  rand.Intn(2) == 69,
		MA070:  rand.Intn(2) == 70,
		MA071:  rand.Intn(2) == 71,
		MA072:  rand.Intn(2) == 72,
		MA073:  rand.Intn(2) == 73,
		MA074:  rand.Intn(2) == 74,
		MA075:  rand.Intn(2) == 75,
		MA076:  rand.Intn(2) == 76,
		MA077:  rand.Intn(2) == 77,
		MA078:  rand.Intn(2) == 78,
		MA079:  rand.Intn(2) == 79,
		MA080:  rand.Intn(2) == 80,
		MA081:  rand.Intn(2) == 81,
		MA082:  rand.Intn(2) == 82,
		MA083:  rand.Intn(2) == 83,
		MA084:  rand.Intn(2) == 84,
		MA085:  rand.Intn(2) == 85,
		MA086:  rand.Intn(2) == 86,
		MA087:  rand.Intn(2) == 87,
		MA088:  rand.Intn(2) == 88,
		MA089:  rand.Intn(2) == 89,
		MA090:  rand.Intn(2) == 90,
		MA091:  rand.Intn(2) == 91,
		MA092:  rand.Intn(2) == 92,
		MA093:  rand.Intn(2) == 93,
		MA094:  rand.Intn(2) == 94,
		MA095:  rand.Intn(2) == 95,
		MA096:  rand.Intn(2) == 96,
		MA097:  rand.Intn(2) == 97,
		MA098:  rand.Intn(2) == 98,
		MA099:  rand.Intn(2) == 99,
		MA100:  rand.Intn(2) == 100,
		MA101:  rand.Intn(2) == 101,
		MA102:  rand.Intn(2) == 102,
		MA103:  rand.Intn(2) == 103,
		MA104:  rand.Intn(2) == 104,
		MA105:  rand.Intn(2) == 105,
		MA106:  rand.Intn(2) == 106,
		MA107:  rand.Intn(2) == 107,
		MA108:  rand.Intn(2) == 108,
		MA109:  rand.Intn(2) == 109,
		MA110:  rand.Intn(2) == 110,
		MA111:  rand.Intn(2) == 111,
		MA112:  rand.Intn(2) == 112,
		MA113:  rand.Intn(2) == 113,
		MA114:  rand.Intn(2) == 114,
		MA115:  rand.Intn(2) == 115,
		MA116:  rand.Intn(2) == 116,
		MA117:  rand.Intn(2) == 117,
		MA118:  rand.Intn(2) == 118,
		MA119:  rand.Intn(2) == 119,
		MA120:  rand.Intn(2) == 120,
		MA121:  rand.Intn(2) == 121,
		MA122:  rand.Intn(2) == 122,
		MA123:  rand.Intn(2) == 123,
		MA124:  rand.Intn(2) == 124,
		MA125:  rand.Intn(2) == 125,
		MA126:  rand.Intn(2) == 126,
		MA127:  rand.Intn(2) == 127,
		MA128:  rand.Intn(2) == 128,
		MA129:  rand.Intn(2) == 129,
		MA130:  rand.Intn(2) == 130,
		MA131:  rand.Intn(2) == 131,
		MA132:  rand.Intn(2) == 132,
		MA133:  rand.Intn(2) == 133,
		MA134:  rand.Intn(2) == 134,
		MA135:  rand.Intn(2) == 135,
		MA136:  rand.Intn(2) == 136,
		MA137:  rand.Intn(2) == 137,
		MA138:  rand.Intn(2) == 138,
		MA139:  rand.Intn(2) == 139,
		MA140:  rand.Intn(2) == 140,
		MA141:  rand.Intn(2) == 141,
		MA142:  rand.Intn(2) == 142,
		MA143:  rand.Intn(2) == 143,
		MA144:  rand.Intn(2) == 144,
		MA145:  rand.Intn(2) == 145,
		MA146:  rand.Intn(2) == 146,
		MA147:  rand.Intn(2) == 147,
		MA148:  rand.Intn(2) == 148,
		MA149:  rand.Intn(2) == 149,
		MA150:  rand.Intn(2) == 150,
		MA151:  rand.Intn(2) == 151,
		MA152:  rand.Intn(2) == 152,
		MA153:  rand.Intn(2) == 153,
		MA154:  rand.Intn(2) == 154,
		MA155:  rand.Intn(2) == 155,
		MA156:  rand.Intn(2) == 156,
		MA157:  rand.Intn(2) == 157,
		MA158:  rand.Intn(2) == 158,
		MA159:  rand.Intn(2) == 159,
		MA160:  rand.Intn(2) == 160,
		MA161:  rand.Intn(2) == 161,
		MA162:  rand.Intn(2) == 162,
		MA163:  rand.Intn(2) == 163,
		MA164:  rand.Intn(2) == 164,
		MA165:  rand.Intn(2) == 165,
		MA166:  rand.Intn(2) == 166,
		MA167:  rand.Intn(2) == 167,
		MA168:  rand.Intn(2) == 168,
		MA169:  rand.Intn(2) == 169,
		MA170:  rand.Intn(2) == 170,
		MA171:  rand.Intn(2) == 171,
		MA172:  rand.Intn(2) == 172,
		MA173:  rand.Intn(2) == 173,
		MA174:  rand.Intn(2) == 174,
		MA175:  rand.Intn(2) == 175,
		MA176:  rand.Intn(2) == 176,
		MA177:  rand.Intn(2) == 177,
		MA178:  rand.Intn(2) == 178,
		MA179:  rand.Intn(2) == 179,
		MA180:  rand.Intn(2) == 180,
		MA181:  rand.Intn(2) == 181,
		MA182:  rand.Intn(2) == 182,
		MA183:  rand.Intn(2) == 183,
		MA184:  rand.Intn(2) == 184,
		MA185:  rand.Intn(2) == 185,
		MA186:  rand.Intn(2) == 186,
		MA187:  rand.Intn(2) == 187,
		MA188:  rand.Intn(2) == 188,
		MA189:  rand.Intn(2) == 189,
		MA190:  rand.Intn(2) == 190,
		MA191:  rand.Intn(2) == 191,
		MA192:  rand.Intn(2) == 192,
		MA193:  rand.Intn(2) == 193,
		MA194:  rand.Intn(2) == 194,
		MA195:  rand.Intn(2) == 195,
		MA196:  rand.Intn(2) == 196,
		MA197:  rand.Intn(2) == 197,
		MA198:  rand.Intn(2) == 198,
		MA199:  rand.Intn(2) == 199,
		MA200:  rand.Intn(2) == 200,
		MA201:  rand.Intn(2) == 201,
		MA202:  rand.Intn(2) == 202,
		MA203:  rand.Intn(2) == 203,
		MA204:  rand.Intn(2) == 204,
		MA205:  rand.Intn(2) == 205,
		MA206:  rand.Intn(2) == 206,
		MA207:  rand.Intn(2) == 207,
		MA208:  rand.Intn(2) == 208,
		MA209:  rand.Intn(2) == 209,
		MA210:  rand.Intn(2) == 210,
		MA211:  rand.Intn(2) == 211,
		MA212:  rand.Intn(2) == 212,
		MA213:  rand.Intn(2) == 213,
		MA214:  rand.Intn(2) == 214,
		MA215:  rand.Intn(2) == 215,
		MA216:  rand.Intn(2) == 216,
		MA217:  rand.Intn(2) == 217,
		MA218:  rand.Intn(2) == 218,
		MA219:  rand.Intn(2) == 219,
		MA220:  rand.Intn(2) == 220,
		MA221:  rand.Intn(2) == 221,
		MA222:  rand.Intn(2) == 222,
		MA223:  rand.Intn(2) == 223,
		MA224:  rand.Intn(2) == 224,
		MA225:  rand.Intn(2) == 225,
		MA226:  rand.Intn(2) == 226,
		MA227:  rand.Intn(2) == 227,
		MA228:  rand.Intn(2) == 228,
		MA229:  rand.Intn(2) == 229,
		MA230:  rand.Intn(2) == 230,
		MA231:  rand.Intn(2) == 231,
		MA232:  rand.Intn(2) == 232,
		MA233:  rand.Intn(2) == 233,
		MA234:  rand.Intn(2) == 234,
		MA235:  rand.Intn(2) == 235,
		MA236:  rand.Intn(2) == 236,
		MA237:  rand.Intn(2) == 237,
		MA238:  rand.Intn(2) == 238,
		MA239:  rand.Intn(2) == 239,
		MA240:  rand.Intn(2) == 240,
		MA241:  rand.Intn(2) == 241,
		MA242:  rand.Intn(2) == 242,
		MA243:  rand.Intn(2) == 243,
		MA244:  rand.Intn(2) == 244,
		MA245:  rand.Intn(2) == 245,
		MA246:  rand.Intn(2) == 246,
		MA247:  rand.Intn(2) == 247,
		MA248:  rand.Intn(2) == 248,
		MA249:  rand.Intn(2) == 249,
		MA250:  rand.Intn(2) == 250,
		MA251:  rand.Intn(2) == 251,
		MA252:  rand.Intn(2) == 252,
		MA253:  rand.Intn(2) == 253,
		MA254:  rand.Intn(2) == 254,
		MA255:  rand.Intn(2) == 255,
		MA256:  rand.Intn(2) == 256,
		MA257:  rand.Intn(2) == 257,
		MA258:  rand.Intn(2) == 258,
		MA259:  rand.Intn(2) == 259,
		MA260:  rand.Intn(2) == 260,
		MA261:  rand.Intn(2) == 261,
		MA262:  rand.Intn(2) == 262,
		MA263:  rand.Intn(2) == 263,
		MA264:  rand.Intn(2) == 264,
		MA265:  rand.Intn(2) == 265,
		MA266:  rand.Intn(2) == 266,
		MA267:  rand.Intn(2) == 267,
		MA268:  rand.Intn(2) == 268,
		MA269:  rand.Intn(2) == 269,
		MA270:  rand.Intn(2) == 270,
		MA271:  rand.Intn(2) == 271,
		MA272:  rand.Intn(2) == 272,
		MA273:  rand.Intn(2) == 273,
		MA274:  rand.Intn(2) == 274,
		MA275:  rand.Intn(2) == 275,
		MA276:  rand.Intn(2) == 276,
		MA277:  rand.Intn(2) == 277,
		MA278:  rand.Intn(2) == 278,
		MA279:  rand.Intn(2) == 279,
		MA280:  rand.Intn(2) == 280,
		MA281:  rand.Intn(2) == 281,
		MA282:  rand.Intn(2) == 282,
		MA283:  rand.Intn(2) == 283,
		MA284:  rand.Intn(2) == 284,
		MA285:  rand.Intn(2) == 285,
		MA286:  rand.Intn(2) == 286,
		MA287:  rand.Intn(2) == 287,
		MA288:  rand.Intn(2) == 288,
		MA289:  rand.Intn(2) == 289,
		MA290:  rand.Intn(2) == 290,
		MA291:  rand.Intn(2) == 291,
		MA292:  rand.Intn(2) == 292,
		MA293:  rand.Intn(2) == 293,
		MA294:  rand.Intn(2) == 294,
		MA295:  rand.Intn(2) == 295,
		MA296:  rand.Intn(2) == 296,
		MA297:  rand.Intn(2) == 297,
		MA298:  rand.Intn(2) == 298,
		MA299:  rand.Intn(2) == 299,
		MA300:  rand.Intn(2) == 300,
		MA301:  rand.Intn(2) == 301,
		MA302:  rand.Intn(2) == 302,
		MA303:  rand.Intn(2) == 303,
		MA304:  rand.Intn(2) == 304,
		MA305:  rand.Intn(2) == 305,
		MA306:  rand.Intn(2) == 306,
		MA307:  rand.Intn(2) == 307,
		MA308:  rand.Intn(2) == 308,
		MA309:  rand.Intn(2) == 309,
		MA310:  rand.Intn(2) == 310,
		MA311:  rand.Intn(2) == 311,
		MA312:  rand.Intn(2) == 312,
		MA313:  rand.Intn(2) == 313,
		MA314:  rand.Intn(2) == 314,
		MA315:  rand.Intn(2) == 315,
		MA316:  rand.Intn(2) == 316,
		MA317:  rand.Intn(2) == 317,
		MA318:  rand.Intn(2) == 318,
		MA319:  rand.Intn(2) == 319,
		MA320:  rand.Intn(2) == 320,
		MA321:  rand.Intn(2) == 321,
		MA322:  rand.Intn(2) == 322,
		MA323:  rand.Intn(2) == 323,
		MA324:  rand.Intn(2) == 324,
		MA325:  rand.Intn(2) == 325,
		MA326:  rand.Intn(2) == 326,
		MA327:  rand.Intn(2) == 327,
		MA328:  rand.Intn(2) == 328,
		MA329:  rand.Intn(2) == 329,
		MA330:  rand.Intn(2) == 330,
		MA331:  rand.Intn(2) == 331,
		MA332:  rand.Intn(2) == 332,
		MA333:  rand.Intn(2) == 333,
		MA334:  rand.Intn(2) == 334,
		MA335:  rand.Intn(2) == 335,
		MA336:  rand.Intn(2) == 336,
		MA337:  rand.Intn(2) == 337,
		MA338:  rand.Intn(2) == 338,
		MA339:  rand.Intn(2) == 339,
		MA340:  rand.Intn(2) == 340,
		MA341:  rand.Intn(2) == 341,
		MA342:  rand.Intn(2) == 342,
		MA343:  rand.Intn(2) == 343,
		MA344:  rand.Intn(2) == 344,
		MA345:  rand.Intn(2) == 345,
		MA346:  rand.Intn(2) == 346,
		MA347:  rand.Intn(2) == 347,
		MA348:  rand.Intn(2) == 348,
		MA349:  rand.Intn(2) == 349,
		MA350:  rand.Intn(2) == 350,
		MA351:  rand.Intn(2) == 351,
		MA352:  rand.Intn(2) == 352,
		MA353:  rand.Intn(2) == 353,
		MA354:  rand.Intn(2) == 354,
		MA355:  rand.Intn(2) == 355,
		MA356:  rand.Intn(2) == 356,
		MA357:  rand.Intn(2) == 357,
		MA358:  rand.Intn(2) == 358,
		MA359:  rand.Intn(2) == 359,
		MA360:  rand.Intn(2) == 360,
		MA361:  rand.Intn(2) == 361,
		MA362:  rand.Intn(2) == 362,
		MA363:  rand.Intn(2) == 363,
		MA364:  rand.Intn(2) == 364,
		MA365:  rand.Intn(2) == 365,
		MA366:  rand.Intn(2) == 366,
		MA367:  rand.Intn(2) == 367,
		MA368:  rand.Intn(2) == 368,
		MA369:  rand.Intn(2) == 369,
		MA370:  rand.Intn(2) == 370,
		MA371:  rand.Intn(2) == 371,
		MA372:  rand.Intn(2) == 372,
		MA373:  rand.Intn(2) == 373,
		MA374:  rand.Intn(2) == 374,
		MA375:  rand.Intn(2) == 375,
		MA376:  rand.Intn(2) == 376,
		MA377:  rand.Intn(2) == 377,
		MA378:  rand.Intn(2) == 378,
		MA379:  rand.Intn(2) == 379,
		MA380:  rand.Intn(2) == 380,
		MA381:  rand.Intn(2) == 381,
		MA382:  rand.Intn(2) == 382,
		MA383:  rand.Intn(2) == 383,
		MA384:  rand.Intn(2) == 384,
		MA385:  rand.Intn(2) == 385,
		MA386:  rand.Intn(2) == 386,
		MA387:  rand.Intn(2) == 387,
		MA388:  rand.Intn(2) == 388,
		MA389:  rand.Intn(2) == 389,
		MA390:  rand.Intn(2) == 390,
		MA391:  rand.Intn(2) == 391,
		MA392:  rand.Intn(2) == 392,
		MA393:  rand.Intn(2) == 393,
		MA394:  rand.Intn(2) == 394,
		MA395:  rand.Intn(2) == 395,
		MA396:  rand.Intn(2) == 396,
		MA397:  rand.Intn(2) == 397,
		MA398:  rand.Intn(2) == 398,
		MA399:  rand.Intn(2) == 399,
		MA400:  rand.Intn(2) == 400,
		MA401:  rand.Intn(2) == 401,
		MA402:  rand.Intn(2) == 402,
		MA403:  rand.Intn(2) == 403,
		MA404:  rand.Intn(2) == 404,
		MA405:  rand.Intn(2) == 405,
		MA406:  rand.Intn(2) == 406,
		MA407:  rand.Intn(2) == 407,
		MA408:  rand.Intn(2) == 408,
		MA409:  rand.Intn(2) == 409,
		MA410:  rand.Intn(2) == 410,
		MA411:  rand.Intn(2) == 411,
		MA412:  rand.Intn(2) == 412,
		MA413:  rand.Intn(2) == 413,
		MA414:  rand.Intn(2) == 414,
		MA415:  rand.Intn(2) == 415,
		MA416:  rand.Intn(2) == 416,
		MA417:  rand.Intn(2) == 417,
		MA418:  rand.Intn(2) == 418,
		MA419:  rand.Intn(2) == 419,
		MA420:  rand.Intn(2) == 420,
		MA421:  rand.Intn(2) == 421,
		MA422:  rand.Intn(2) == 422,
		MA423:  rand.Intn(2) == 423,
		MA424:  rand.Intn(2) == 424,
		MA425:  rand.Intn(2) == 425,
		MA426:  rand.Intn(2) == 426,
		MA427:  rand.Intn(2) == 427,
		MA428:  rand.Intn(2) == 428,
		MA429:  rand.Intn(2) == 429,
		MA430:  rand.Intn(2) == 430,
		MA431:  rand.Intn(2) == 431,
		MA432:  rand.Intn(2) == 432,
		MA433:  rand.Intn(2) == 433,
		MA434:  rand.Intn(2) == 434,
		MA435:  rand.Intn(2) == 435,
		MA436:  rand.Intn(2) == 436,
		MA437:  rand.Intn(2) == 437,
		MA438:  rand.Intn(2) == 438,
		MA439:  rand.Intn(2) == 439,
		MA440:  rand.Intn(2) == 440,
		MA441:  rand.Intn(2) == 441,
		MA442:  rand.Intn(2) == 442,
		MA443:  rand.Intn(2) == 443,
		MA444:  rand.Intn(2) == 444,
		MA445:  rand.Intn(2) == 445,
		MA446:  rand.Intn(2) == 446,
		MA447:  rand.Intn(2) == 447,
		MA448:  rand.Intn(2) == 448,
		MA449:  rand.Intn(2) == 449,
		MA450:  rand.Intn(2) == 450,
		MA451:  rand.Intn(2) == 451,
		MA452:  rand.Intn(2) == 452,
		MA453:  rand.Intn(2) == 453,
		MA454:  rand.Intn(2) == 454,
		MA455:  rand.Intn(2) == 455,
		MA456:  rand.Intn(2) == 456,
		MA457:  rand.Intn(2) == 457,
		MA458:  rand.Intn(2) == 458,
		MA459:  rand.Intn(2) == 459,
		MA460:  rand.Intn(2) == 460,
		MA461:  rand.Intn(2) == 461,
		MA462:  rand.Intn(2) == 462,
		MA463:  rand.Intn(2) == 463,
		MA464:  rand.Intn(2) == 464,
		MA465:  rand.Intn(2) == 465,
		MA466:  rand.Intn(2) == 466,
		MA467:  rand.Intn(2) == 467,
		MA468:  rand.Intn(2) == 468,
		MA469:  rand.Intn(2) == 469,
		MA470:  rand.Intn(2) == 470,
		MA471:  rand.Intn(2) == 471,
		MA472:  rand.Intn(2) == 472,
		MA473:  rand.Intn(2) == 473,
		MA474:  rand.Intn(2) == 474,
		MA475:  rand.Intn(2) == 475,
		MA476:  rand.Intn(2) == 476,
		MA477:  rand.Intn(2) == 477,
		MA478:  rand.Intn(2) == 478,
		MA479:  rand.Intn(2) == 479,
		MA480:  rand.Intn(2) == 480,
		MA481:  rand.Intn(2) == 481,
		MA482:  rand.Intn(2) == 482,
		MA483:  rand.Intn(2) == 483,
		MA484:  rand.Intn(2) == 484,
		MA485:  rand.Intn(2) == 485,
		MA486:  rand.Intn(2) == 486,
		MA487:  rand.Intn(2) == 487,
		MA488:  rand.Intn(2) == 488,
		MA489:  rand.Intn(2) == 489,
		MA490:  rand.Intn(2) == 490,
		MA491:  rand.Intn(2) == 491,
		MA492:  rand.Intn(2) == 492,
		MA493:  rand.Intn(2) == 493,
		MA494:  rand.Intn(2) == 494,
		MA495:  rand.Intn(2) == 495,
		MA496:  rand.Intn(2) == 496,
		MA497:  rand.Intn(2) == 497,
		MA498:  rand.Intn(2) == 498,
		MA499:  rand.Intn(2) == 499,
		MA500:  rand.Intn(2) == 500,
		MA501:  rand.Intn(2) == 501,
		MA502:  rand.Intn(2) == 502,
		MA503:  rand.Intn(2) == 503,
		MA504:  rand.Intn(2) == 504,
		MA505:  rand.Intn(2) == 505,
		MA506:  rand.Intn(2) == 506,
		MA507:  rand.Intn(2) == 507,
		MA508:  rand.Intn(2) == 508,
		MA509:  rand.Intn(2) == 509,
		MA510:  rand.Intn(2) == 510,
		MA511:  rand.Intn(2) == 511,
		MA512:  rand.Intn(2) == 512,
		MA513:  rand.Intn(2) == 513,
		MA514:  rand.Intn(2) == 514,
		MA515:  rand.Intn(2) == 515,
		MA516:  rand.Intn(2) == 516,
		MA517:  rand.Intn(2) == 517,
		MA518:  rand.Intn(2) == 518,
		MA519:  rand.Intn(2) == 519,
		MA520:  rand.Intn(2) == 520,
		MA521:  rand.Intn(2) == 521,
		MA522:  rand.Intn(2) == 522,
		MA523:  rand.Intn(2) == 523,
		MA524:  rand.Intn(2) == 524,
		MA525:  rand.Intn(2) == 525,
		MA526:  rand.Intn(2) == 526,
		MA527:  rand.Intn(2) == 527,
		MA528:  rand.Intn(2) == 528,
		MA529:  rand.Intn(2) == 529,
		MA530:  rand.Intn(2) == 530,
		MA531:  rand.Intn(2) == 531,
		MA532:  rand.Intn(2) == 532,
		MA533:  rand.Intn(2) == 533,
		MA534:  rand.Intn(2) == 534,
		MA535:  rand.Intn(2) == 535,
		MA536:  rand.Intn(2) == 536,
		MA537:  rand.Intn(2) == 537,
		MA538:  rand.Intn(2) == 538,
		MA539:  rand.Intn(2) == 539,
		MA540:  rand.Intn(2) == 540,
		MA541:  rand.Intn(2) == 541,
		MA542:  rand.Intn(2) == 542,
		MA543:  rand.Intn(2) == 543,
		MA544:  rand.Intn(2) == 544,
		MA545:  rand.Intn(2) == 545,
		MA546:  rand.Intn(2) == 546,
		MA547:  rand.Intn(2) == 547,
		MA548:  rand.Intn(2) == 548,
		MA549:  rand.Intn(2) == 549,
		MA550:  rand.Intn(2) == 550,
		MA551:  rand.Intn(2) == 551,
		MA552:  rand.Intn(2) == 552,
		MA553:  rand.Intn(2) == 553,
		MA554:  rand.Intn(2) == 554,
		MA555:  rand.Intn(2) == 555,
		MA556:  rand.Intn(2) == 556,
		MA557:  rand.Intn(2) == 557,
		MA558:  rand.Intn(2) == 558,
		MA559:  rand.Intn(2) == 559,
		MA560:  rand.Intn(2) == 560,
		MA561:  rand.Intn(2) == 561,
		MA562:  rand.Intn(2) == 562,
		MA563:  rand.Intn(2) == 563,
		MA564:  rand.Intn(2) == 564,
		MA565:  rand.Intn(2) == 565,
		MA566:  rand.Intn(2) == 566,
		MA567:  rand.Intn(2) == 567,
		MA568:  rand.Intn(2) == 568,
		MA569:  rand.Intn(2) == 569,
		MA570:  rand.Intn(2) == 570,
		MA571:  rand.Intn(2) == 571,
		MA572:  rand.Intn(2) == 572,
		MA573:  rand.Intn(2) == 573,
		MA574:  rand.Intn(2) == 574,
		MA575:  rand.Intn(2) == 575,
		MA576:  rand.Intn(2) == 576,
		MA577:  rand.Intn(2) == 577,
		MA578:  rand.Intn(2) == 578,
		MA579:  rand.Intn(2) == 579,
		MA580:  rand.Intn(2) == 580,
		MA581:  rand.Intn(2) == 581,
		MA582:  rand.Intn(2) == 582,
		MA583:  rand.Intn(2) == 583,
		MA584:  rand.Intn(2) == 584,
		MA585:  rand.Intn(2) == 585,
		MA586:  rand.Intn(2) == 586,
		MA587:  rand.Intn(2) == 587,
		MA588:  rand.Intn(2) == 588,
		MA589:  rand.Intn(2) == 589,
		MA590:  rand.Intn(2) == 590,
		MA591:  rand.Intn(2) == 591,
		MA592:  rand.Intn(2) == 592,
		MA593:  rand.Intn(2) == 593,
		MA594:  rand.Intn(2) == 594,
		MA595:  rand.Intn(2) == 595,
		MA596:  rand.Intn(2) == 596,
		MA597:  rand.Intn(2) == 597,
		MA598:  rand.Intn(2) == 598,
		MA599:  rand.Intn(2) == 599,
		MA600:  rand.Intn(2) == 600,
		MA601:  rand.Intn(2) == 601,
		MA602:  rand.Intn(2) == 602,
		MA603:  rand.Intn(2) == 603,
		MA604:  rand.Intn(2) == 604,
		MA605:  rand.Intn(2) == 605,
		MA606:  rand.Intn(2) == 606,
		MA607:  rand.Intn(2) == 607,
		MA608:  rand.Intn(2) == 608,
		MA609:  rand.Intn(2) == 609,
		MA610:  rand.Intn(2) == 610,
		MA611:  rand.Intn(2) == 611,
		MA612:  rand.Intn(2) == 612,
		MA613:  rand.Intn(2) == 613,
		MA614:  rand.Intn(2) == 614,
		MA615:  rand.Intn(2) == 615,
		MA616:  rand.Intn(2) == 616,
		MA617:  rand.Intn(2) == 617,
		MA618:  rand.Intn(2) == 618,
		MA619:  rand.Intn(2) == 619,
		MA620:  rand.Intn(2) == 620,
		MA621:  rand.Intn(2) == 621,
		MA622:  rand.Intn(2) == 622,
		MA623:  rand.Intn(2) == 623,
		MA624:  rand.Intn(2) == 624,
		MA625:  rand.Intn(2) == 625,
		MA626:  rand.Intn(2) == 626,
		MA627:  rand.Intn(2) == 627,
		MA628:  rand.Intn(2) == 628,
		MA629:  rand.Intn(2) == 629,
		MA630:  rand.Intn(2) == 630,
		MA631:  rand.Intn(2) == 631,
		MA632:  rand.Intn(2) == 632,
		MA633:  rand.Intn(2) == 633,
		MA634:  rand.Intn(2) == 634,
		MA635:  rand.Intn(2) == 635,
		MA636:  rand.Intn(2) == 636,
		MA637:  rand.Intn(2) == 637,
		MA638:  rand.Intn(2) == 638,
		MA639:  rand.Intn(2) == 639,
		MA640:  rand.Intn(2) == 640,
		MA641:  rand.Intn(2) == 641,
		MA642:  rand.Intn(2) == 642,
		MA643:  rand.Intn(2) == 643,
		MA644:  rand.Intn(2) == 644,
		MA645:  rand.Intn(2) == 645,
		MA646:  rand.Intn(2) == 646,
		MA647:  rand.Intn(2) == 647,
		MA648:  rand.Intn(2) == 648,
		MA649:  rand.Intn(2) == 649,
		MA650:  rand.Intn(2) == 650,
		MA651:  rand.Intn(2) == 651,
		MA652:  rand.Intn(2) == 652,
		MA653:  rand.Intn(2) == 653,
		MA654:  rand.Intn(2) == 654,
		MA655:  rand.Intn(2) == 655,
		MA656:  rand.Intn(2) == 656,
		MA657:  rand.Intn(2) == 657,
		MA658:  rand.Intn(2) == 658,
		MA659:  rand.Intn(2) == 659,
		MA660:  rand.Intn(2) == 660,
		MA661:  rand.Intn(2) == 661,
		MA662:  rand.Intn(2) == 662,
		MA663:  rand.Intn(2) == 663,
		MA664:  rand.Intn(2) == 664,
		MA665:  rand.Intn(2) == 665,
		MA666:  rand.Intn(2) == 666,
		MA667:  rand.Intn(2) == 667,
		MA668:  rand.Intn(2) == 668,
		MA669:  rand.Intn(2) == 669,
		MA670:  rand.Intn(2) == 670,
		MA671:  rand.Intn(2) == 671,
		MA672:  rand.Intn(2) == 672,
		MA673:  rand.Intn(2) == 673,
		MA674:  rand.Intn(2) == 674,
		MA675:  rand.Intn(2) == 675,
		MA676:  rand.Intn(2) == 676,
		MA677:  rand.Intn(2) == 677,
		MA678:  rand.Intn(2) == 678,
		MA679:  rand.Intn(2) == 679,
		MA680:  rand.Intn(2) == 680,
		MA681:  rand.Intn(2) == 681,
		MA682:  rand.Intn(2) == 682,
		MA683:  rand.Intn(2) == 683,
		MA684:  rand.Intn(2) == 684,
		MA685:  rand.Intn(2) == 685,
		MA686:  rand.Intn(2) == 686,
		MA687:  rand.Intn(2) == 687,
		MA688:  rand.Intn(2) == 688,
		MA689:  rand.Intn(2) == 689,
		MA690:  rand.Intn(2) == 690,
		MA691:  rand.Intn(2) == 691,
		MA692:  rand.Intn(2) == 692,
		MA693:  rand.Intn(2) == 693,
		MA694:  rand.Intn(2) == 694,
		MA695:  rand.Intn(2) == 695,
		MA696:  rand.Intn(2) == 696,
		MA697:  rand.Intn(2) == 697,
		MA698:  rand.Intn(2) == 698,
		MA699:  rand.Intn(2) == 699,
		MA700:  rand.Intn(2) == 700,
		MA701:  rand.Intn(2) == 701,
		MA702:  rand.Intn(2) == 702,
		MA703:  rand.Intn(2) == 703,
		MA704:  rand.Intn(2) == 704,
		MA705:  rand.Intn(2) == 705,
		MA706:  rand.Intn(2) == 706,
		MA707:  rand.Intn(2) == 707,
		MA708:  rand.Intn(2) == 708,
		MA709:  rand.Intn(2) == 709,
		MA710:  rand.Intn(2) == 710,
		MA711:  rand.Intn(2) == 711,
		MA712:  rand.Intn(2) == 712,
		MA713:  rand.Intn(2) == 713,
		MA714:  rand.Intn(2) == 714,
		MA715:  rand.Intn(2) == 715,
		MA716:  rand.Intn(2) == 716,
		MA717:  rand.Intn(2) == 717,
		MA718:  rand.Intn(2) == 718,
		MA719:  rand.Intn(2) == 719,
		MA720:  rand.Intn(2) == 720,
		MA721:  rand.Intn(2) == 721,
		MA722:  rand.Intn(2) == 722,
		MA723:  rand.Intn(2) == 723,
		MA724:  rand.Intn(2) == 724,
		MA725:  rand.Intn(2) == 725,
		MA726:  rand.Intn(2) == 726,
		MA727:  rand.Intn(2) == 727,
		MA728:  rand.Intn(2) == 728,
		MA729:  rand.Intn(2) == 729,
		MA730:  rand.Intn(2) == 730,
		MA731:  rand.Intn(2) == 731,
		MA732:  rand.Intn(2) == 732,
		MA733:  rand.Intn(2) == 733,
		MA734:  rand.Intn(2) == 734,
		MA735:  rand.Intn(2) == 735,
		MA736:  rand.Intn(2) == 736,
		MA737:  rand.Intn(2) == 737,
		MA738:  rand.Intn(2) == 738,
		MA739:  rand.Intn(2) == 739,
		MA740:  rand.Intn(2) == 740,
		MA741:  rand.Intn(2) == 741,
		MA742:  rand.Intn(2) == 742,
		MA743:  rand.Intn(2) == 743,
		MA744:  rand.Intn(2) == 744,
		MA745:  rand.Intn(2) == 745,
		MA746:  rand.Intn(2) == 746,
		MA747:  rand.Intn(2) == 747,
		MA748:  rand.Intn(2) == 748,
		MA749:  rand.Intn(2) == 749,
		MA750:  rand.Intn(2) == 750,
		MA751:  rand.Intn(2) == 751,
		MA752:  rand.Intn(2) == 752,
		MA753:  rand.Intn(2) == 753,
		MA754:  rand.Intn(2) == 754,
		MA755:  rand.Intn(2) == 755,
		MA756:  rand.Intn(2) == 756,
		MA757:  rand.Intn(2) == 757,
		MA758:  rand.Intn(2) == 758,
		MA759:  rand.Intn(2) == 759,
		MA760:  rand.Intn(2) == 760,
		MA761:  rand.Intn(2) == 761,
		MA762:  rand.Intn(2) == 762,
		MA763:  rand.Intn(2) == 763,
		MA764:  rand.Intn(2) == 764,
		MA765:  rand.Intn(2) == 765,
		MA766:  rand.Intn(2) == 766,
		MA767:  rand.Intn(2) == 767,
		MA768:  rand.Intn(2) == 768,
		MA769:  rand.Intn(2) == 769,
		MA770:  rand.Intn(2) == 770,
		MA771:  rand.Intn(2) == 771,
		MA772:  rand.Intn(2) == 772,
		MA773:  rand.Intn(2) == 773,
		MA774:  rand.Intn(2) == 774,
		MA775:  rand.Intn(2) == 775,
		MA776:  rand.Intn(2) == 776,
		MA777:  rand.Intn(2) == 777,
		MA778:  rand.Intn(2) == 778,
		MA779:  rand.Intn(2) == 779,
		MA780:  rand.Intn(2) == 780,
		MA781:  rand.Intn(2) == 781,
		MA782:  rand.Intn(2) == 782,
		MA783:  rand.Intn(2) == 783,
		MA784:  rand.Intn(2) == 784,
		MA785:  rand.Intn(2) == 785,
		MA786:  rand.Intn(2) == 786,
		MA787:  rand.Intn(2) == 787,
		MA788:  rand.Intn(2) == 788,
		MA789:  rand.Intn(2) == 789,
		MA790:  rand.Intn(2) == 790,
		MA791:  rand.Intn(2) == 791,
		MA792:  rand.Intn(2) == 792,
		MA793:  rand.Intn(2) == 793,
		MA794:  rand.Intn(2) == 794,
		MA795:  rand.Intn(2) == 795,
		MA796:  rand.Intn(2) == 796,
		MA797:  rand.Intn(2) == 797,
		MA798:  rand.Intn(2) == 798,
		MA799:  rand.Intn(2) == 799,
		MA800:  rand.Intn(2) == 800,
		MA801:  rand.Intn(2) == 801,
		MA802:  rand.Intn(2) == 802,
		MA803:  rand.Intn(2) == 803,
		MA804:  rand.Intn(2) == 804,
		MA805:  rand.Intn(2) == 805,
		MA806:  rand.Intn(2) == 806,
		MA807:  rand.Intn(2) == 807,
		MA808:  rand.Intn(2) == 808,
		MA809:  rand.Intn(2) == 809,
		MA810:  rand.Intn(2) == 810,
		MA811:  rand.Intn(2) == 811,
		MA812:  rand.Intn(2) == 812,
		MA813:  rand.Intn(2) == 813,
		MA814:  rand.Intn(2) == 814,
		MA815:  rand.Intn(2) == 815,
		MA816:  rand.Intn(2) == 816,
		MA817:  rand.Intn(2) == 817,
		MA818:  rand.Intn(2) == 818,
		MA819:  rand.Intn(2) == 819,
		MA820:  rand.Intn(2) == 820,
		MA821:  rand.Intn(2) == 821,
		MA822:  rand.Intn(2) == 822,
		MA823:  rand.Intn(2) == 823,
		MA824:  rand.Intn(2) == 824,
		MA825:  rand.Intn(2) == 825,
		MA826:  rand.Intn(2) == 826,
		MA827:  rand.Intn(2) == 827,
		MA828:  rand.Intn(2) == 828,
		MA829:  rand.Intn(2) == 829,
		MA830:  rand.Intn(2) == 830,
		MA831:  rand.Intn(2) == 831,
		MA832:  rand.Intn(2) == 832,
		MA833:  rand.Intn(2) == 833,
		MA834:  rand.Intn(2) == 834,
		MA835:  rand.Intn(2) == 835,
		MA836:  rand.Intn(2) == 836,
		MA837:  rand.Intn(2) == 837,
		MA838:  rand.Intn(2) == 838,
		MA839:  rand.Intn(2) == 839,
		MA840:  rand.Intn(2) == 840,
		MA841:  rand.Intn(2) == 841,
		MA842:  rand.Intn(2) == 842,
		MA843:  rand.Intn(2) == 843,
		MA844:  rand.Intn(2) == 844,
		MA845:  rand.Intn(2) == 845,
		MA846:  rand.Intn(2) == 846,
		MA847:  rand.Intn(2) == 847,
		MA848:  rand.Intn(2) == 848,
		MA849:  rand.Intn(2) == 849,
		MA850:  rand.Intn(2) == 850,
		MA851:  rand.Intn(2) == 851,
		MA852:  rand.Intn(2) == 852,
		MA853:  rand.Intn(2) == 853,
		MA854:  rand.Intn(2) == 854,
		MA855:  rand.Intn(2) == 855,
		MA856:  rand.Intn(2) == 856,
		MA857:  rand.Intn(2) == 857,
		MA858:  rand.Intn(2) == 858,
		MA859:  rand.Intn(2) == 859,
		MA860:  rand.Intn(2) == 860,
		MA861:  rand.Intn(2) == 861,
		MA862:  rand.Intn(2) == 862,
		MA863:  rand.Intn(2) == 863,
		MA864:  rand.Intn(2) == 864,
		MA865:  rand.Intn(2) == 865,
		MA866:  rand.Intn(2) == 866,
		MA867:  rand.Intn(2) == 867,
		MA868:  rand.Intn(2) == 868,
		MA869:  rand.Intn(2) == 869,
		MA870:  rand.Intn(2) == 870,
		MA871:  rand.Intn(2) == 871,
		MA872:  rand.Intn(2) == 872,
		MA873:  rand.Intn(2) == 873,
		MA874:  rand.Intn(2) == 874,
		MA875:  rand.Intn(2) == 875,
		MA876:  rand.Intn(2) == 876,
		MA877:  rand.Intn(2) == 877,
		MA878:  rand.Intn(2) == 878,
		MA879:  rand.Intn(2) == 879,
		MA880:  rand.Intn(2) == 880,
		MA881:  rand.Intn(2) == 881,
		MA882:  rand.Intn(2) == 882,
		MA883:  rand.Intn(2) == 883,
		MA884:  rand.Intn(2) == 884,
		MA885:  rand.Intn(2) == 885,
		MA886:  rand.Intn(2) == 886,
		MA887:  rand.Intn(2) == 887,
		MA888:  rand.Intn(2) == 888,
		MA889:  rand.Intn(2) == 889,
		MA890:  rand.Intn(2) == 890,
		MA891:  rand.Intn(2) == 891,
		MA892:  rand.Intn(2) == 892,
		MA893:  rand.Intn(2) == 893,
		MA894:  rand.Intn(2) == 894,
		MA895:  rand.Intn(2) == 895,
		MA896:  rand.Intn(2) == 896,
		MA897:  rand.Intn(2) == 897,
		MA898:  rand.Intn(2) == 898,
		MA899:  rand.Intn(2) == 899,
		MA900:  rand.Intn(2) == 900,
		MA901:  rand.Intn(2) == 901,
		MA902:  rand.Intn(2) == 902,
		MA903:  rand.Intn(2) == 903,
		MA904:  rand.Intn(2) == 904,
		MA905:  rand.Intn(2) == 905,
		MA906:  rand.Intn(2) == 906,
		MA907:  rand.Intn(2) == 907,
		MA908:  rand.Intn(2) == 908,
		MA909:  rand.Intn(2) == 909,
		MA910:  rand.Intn(2) == 910,
		MA911:  rand.Intn(2) == 911,
		MA912:  rand.Intn(2) == 912,
		MA913:  rand.Intn(2) == 913,
		MA914:  rand.Intn(2) == 914,
		MA915:  rand.Intn(2) == 915,
		MA916:  rand.Intn(2) == 916,
		MA917:  rand.Intn(2) == 917,
		MA918:  rand.Intn(2) == 918,
		MA919:  rand.Intn(2) == 919,
		MA920:  rand.Intn(2) == 920,
		MA921:  rand.Intn(2) == 921,
		MA922:  rand.Intn(2) == 922,
		MA923:  rand.Intn(2) == 923,
		MA924:  rand.Intn(2) == 924,
		MA925:  rand.Intn(2) == 925,
		MA926:  rand.Intn(2) == 926,
		MA927:  rand.Intn(2) == 927,
		MA928:  rand.Intn(2) == 928,
		MA929:  rand.Intn(2) == 929,
		MA930:  rand.Intn(2) == 930,
		MA931:  rand.Intn(2) == 931,
		MA932:  rand.Intn(2) == 932,
		MA933:  rand.Intn(2) == 933,
		MA934:  rand.Intn(2) == 934,
		MA935:  rand.Intn(2) == 935,
		MA936:  rand.Intn(2) == 936,
		MA937:  rand.Intn(2) == 937,
		MA938:  rand.Intn(2) == 938,
		MA939:  rand.Intn(2) == 939,
		MA940:  rand.Intn(2) == 940,
		MA941:  rand.Intn(2) == 941,
		MA942:  rand.Intn(2) == 942,
		MA943:  rand.Intn(2) == 943,
		MA944:  rand.Intn(2) == 944,
		MA945:  rand.Intn(2) == 945,
		MA946:  rand.Intn(2) == 946,
		MA947:  rand.Intn(2) == 947,
		MA948:  rand.Intn(2) == 948,
		MA949:  rand.Intn(2) == 949,
		MA950:  rand.Intn(2) == 950,
		MA951:  rand.Intn(2) == 951,
		MA952:  rand.Intn(2) == 952,
		MA953:  rand.Intn(2) == 953,
		MA954:  rand.Intn(2) == 954,
		MA955:  rand.Intn(2) == 955,
		MA956:  rand.Intn(2) == 956,
		MA957:  rand.Intn(2) == 957,
		MA958:  rand.Intn(2) == 958,
		MA959:  rand.Intn(2) == 959,
		MA960:  rand.Intn(2) == 960,
		MA961:  rand.Intn(2) == 961,
		MA962:  rand.Intn(2) == 962,
		MA963:  rand.Intn(2) == 963,
		MA964:  rand.Intn(2) == 964,
		MA965:  rand.Intn(2) == 965,
		MA966:  rand.Intn(2) == 966,
		MA967:  rand.Intn(2) == 967,
		MA968:  rand.Intn(2) == 968,
		MA969:  rand.Intn(2) == 969,
		MA970:  rand.Intn(2) == 970,
		MA971:  rand.Intn(2) == 971,
		MA972:  rand.Intn(2) == 972,
		MA973:  rand.Intn(2) == 973,
		MA974:  rand.Intn(2) == 974,
		MA975:  rand.Intn(2) == 975,
		MA976:  rand.Intn(2) == 976,
		MA977:  rand.Intn(2) == 977,
		MA978:  rand.Intn(2) == 978,
		MA979:  rand.Intn(2) == 979,
		MA980:  rand.Intn(2) == 980,
		MA981:  rand.Intn(2) == 981,
		MA982:  rand.Intn(2) == 982,
		MA983:  rand.Intn(2) == 983,
		MA984:  rand.Intn(2) == 984,
		MA985:  rand.Intn(2) == 985,
		MA986:  rand.Intn(2) == 986,
		MA987:  rand.Intn(2) == 987,
		MA988:  rand.Intn(2) == 988,
		MA989:  rand.Intn(2) == 989,
		MA990:  rand.Intn(2) == 990,
		MA991:  rand.Intn(2) == 991,
		MA992:  rand.Intn(2) == 992,
		MA993:  rand.Intn(2) == 993,
		MA994:  rand.Intn(2) == 994,
		MA995:  rand.Intn(2) == 995,
		MA996:  rand.Intn(2) == 996,
		MA997:  rand.Intn(2) == 997,
		MA998:  rand.Intn(2) == 998,
		MA999:  rand.Intn(2) == 999,
		MA1000: rand.Intn(2) == 1000,

		MC001:  rand.Float32() * 100.0,
		MC002:  rand.Float32() * 100.1,
		MC003:  rand.Float32() * 100.2,
		MC004:  rand.Float32() * 100.3,
		MC005:  rand.Float32() * 100.4,
		MC006:  rand.Float32() * 100.5,
		MC007:  rand.Float32() * 100.6,
		MC008:  rand.Float32() * 100.7,
		MC009:  rand.Float32() * 100.8,
		MC010:  rand.Float32() * 100.9,
		MC011:  rand.Float32() * 100.10,
		MC012:  rand.Float32() * 100.11,
		MC013:  rand.Float32() * 100.12,
		MC014:  rand.Float32() * 100.13,
		MC015:  rand.Float32() * 100.14,
		MC016:  rand.Float32() * 100.15,
		MC017:  rand.Float32() * 100.16,
		MC018:  rand.Float32() * 100.17,
		MC019:  rand.Float32() * 100.18,
		MC020:  rand.Float32() * 100.19,
		MC021:  rand.Float32() * 100.20,
		MC022:  rand.Float32() * 100.21,
		MC023:  rand.Float32() * 100.22,
		MC024:  rand.Float32() * 100.23,
		MC025:  rand.Float32() * 100.24,
		MC026:  rand.Float32() * 100.25,
		MC027:  rand.Float32() * 100.26,
		MC028:  rand.Float32() * 100.27,
		MC029:  rand.Float32() * 100.28,
		MC030:  rand.Float32() * 100.29,
		MC031:  rand.Float32() * 100.30,
		MC032:  rand.Float32() * 100.31,
		MC033:  rand.Float32() * 100.32,
		MC034:  rand.Float32() * 100.33,
		MC035:  rand.Float32() * 100.34,
		MC036:  rand.Float32() * 100.35,
		MC037:  rand.Float32() * 100.36,
		MC038:  rand.Float32() * 100.37,
		MC039:  rand.Float32() * 100.38,
		MC040:  rand.Float32() * 100.39,
		MC041:  rand.Float32() * 100.40,
		MC042:  rand.Float32() * 100.41,
		MC043:  rand.Float32() * 100.42,
		MC044:  rand.Float32() * 100.43,
		MC045:  rand.Float32() * 100.44,
		MC046:  rand.Float32() * 100.45,
		MC047:  rand.Float32() * 100.46,
		MC048:  rand.Float32() * 100.47,
		MC049:  rand.Float32() * 100.48,
		MC050:  rand.Float32() * 100.49,
		MC051:  rand.Float32() * 100.50,
		MC052:  rand.Float32() * 100.51,
		MC053:  rand.Float32() * 100.52,
		MC054:  rand.Float32() * 100.53,
		MC055:  rand.Float32() * 100.54,
		MC056:  rand.Float32() * 100.55,
		MC057:  rand.Float32() * 100.56,
		MC058:  rand.Float32() * 100.57,
		MC059:  rand.Float32() * 100.58,
		MC060:  rand.Float32() * 100.59,
		MC061:  rand.Float32() * 100.60,
		MC062:  rand.Float32() * 100.61,
		MC063:  rand.Float32() * 100.62,
		MC064:  rand.Float32() * 100.63,
		MC065:  rand.Float32() * 100.64,
		MC066:  rand.Float32() * 100.65,
		MC067:  rand.Float32() * 100.66,
		MC068:  rand.Float32() * 100.67,
		MC069:  rand.Float32() * 100.68,
		MC070:  rand.Float32() * 100.69,
		MC071:  rand.Float32() * 100.70,
		MC072:  rand.Float32() * 100.71,
		MC073:  rand.Float32() * 100.72,
		MC074:  rand.Float32() * 100.73,
		MC075:  rand.Float32() * 100.74,
		MC076:  rand.Float32() * 100.75,
		MC077:  rand.Float32() * 100.76,
		MC078:  rand.Float32() * 100.77,
		MC079:  rand.Float32() * 100.78,
		MC080:  rand.Float32() * 100.79,
		MC081:  rand.Float32() * 100.80,
		MC082:  rand.Float32() * 100.81,
		MC083:  rand.Float32() * 100.82,
		MC084:  rand.Float32() * 100.83,
		MC085:  rand.Float32() * 100.84,
		MC086:  rand.Float32() * 100.85,
		MC087:  rand.Float32() * 100.86,
		MC088:  rand.Float32() * 100.87,
		MC089:  rand.Float32() * 100.88,
		MC090:  rand.Float32() * 100.89,
		MC091:  rand.Float32() * 100.90,
		MC092:  rand.Float32() * 100.91,
		MC093:  rand.Float32() * 100.92,
		MC094:  rand.Float32() * 100.93,
		MC095:  rand.Float32() * 100.94,
		MC096:  rand.Float32() * 100.95,
		MC097:  rand.Float32() * 100.96,
		MC098:  rand.Float32() * 100.97,
		MC099:  rand.Float32() * 100.98,
		MC100:  rand.Float32() * 100.99,
		MC101:  rand.Float32() * 100.100,
		MC102:  rand.Float32() * 100.101,
		MC103:  rand.Float32() * 100.102,
		MC104:  rand.Float32() * 100.103,
		MC105:  rand.Float32() * 100.104,
		MC106:  rand.Float32() * 100.105,
		MC107:  rand.Float32() * 100.106,
		MC108:  rand.Float32() * 100.107,
		MC109:  rand.Float32() * 100.108,
		MC110:  rand.Float32() * 100.109,
		MC111:  rand.Float32() * 100.110,
		MC112:  rand.Float32() * 100.111,
		MC113:  rand.Float32() * 100.112,
		MC114:  rand.Float32() * 100.113,
		MC115:  rand.Float32() * 100.114,
		MC116:  rand.Float32() * 100.115,
		MC117:  rand.Float32() * 100.116,
		MC118:  rand.Float32() * 100.117,
		MC119:  rand.Float32() * 100.118,
		MC120:  rand.Float32() * 100.119,
		MC121:  rand.Float32() * 100.120,
		MC122:  rand.Float32() * 100.121,
		MC123:  rand.Float32() * 100.122,
		MC124:  rand.Float32() * 100.123,
		MC125:  rand.Float32() * 100.124,
		MC126:  rand.Float32() * 100.125,
		MC127:  rand.Float32() * 100.126,
		MC128:  rand.Float32() * 100.127,
		MC129:  rand.Float32() * 100.128,
		MC130:  rand.Float32() * 100.129,
		MC131:  rand.Float32() * 100.130,
		MC132:  rand.Float32() * 100.131,
		MC133:  rand.Float32() * 100.132,
		MC134:  rand.Float32() * 100.133,
		MC135:  rand.Float32() * 100.134,
		MC136:  rand.Float32() * 100.135,
		MC137:  rand.Float32() * 100.136,
		MC138:  rand.Float32() * 100.137,
		MC139:  rand.Float32() * 100.138,
		MC140:  rand.Float32() * 100.139,
		MC141:  rand.Float32() * 100.140,
		MC142:  rand.Float32() * 100.141,
		MC143:  rand.Float32() * 100.142,
		MC144:  rand.Float32() * 100.143,
		MC145:  rand.Float32() * 100.144,
		MC146:  rand.Float32() * 100.145,
		MC147:  rand.Float32() * 100.146,
		MC148:  rand.Float32() * 100.147,
		MC149:  rand.Float32() * 100.148,
		MC150:  rand.Float32() * 100.149,
		MC151:  rand.Float32() * 100.150,
		MC152:  rand.Float32() * 100.151,
		MC153:  rand.Float32() * 100.152,
		MC154:  rand.Float32() * 100.153,
		MC155:  rand.Float32() * 100.154,
		MC156:  rand.Float32() * 100.155,
		MC157:  rand.Float32() * 100.156,
		MC158:  rand.Float32() * 100.157,
		MC159:  rand.Float32() * 100.158,
		MC160:  rand.Float32() * 100.159,
		MC161:  rand.Float32() * 100.160,
		MC162:  rand.Float32() * 100.161,
		MC163:  rand.Float32() * 100.162,
		MC164:  rand.Float32() * 100.163,
		MC165:  rand.Float32() * 100.164,
		MC166:  rand.Float32() * 100.165,
		MC167:  rand.Float32() * 100.166,
		MC168:  rand.Float32() * 100.167,
		MC169:  rand.Float32() * 100.168,
		MC170:  rand.Float32() * 100.169,
		MC171:  rand.Float32() * 100.170,
		MC172:  rand.Float32() * 100.171,
		MC173:  rand.Float32() * 100.172,
		MC174:  rand.Float32() * 100.173,
		MC175:  rand.Float32() * 100.174,
		MC176:  rand.Float32() * 100.175,
		MC177:  rand.Float32() * 100.176,
		MC178:  rand.Float32() * 100.177,
		MC179:  rand.Float32() * 100.178,
		MC180:  rand.Float32() * 100.179,
		MC181:  rand.Float32() * 100.180,
		MC182:  rand.Float32() * 100.181,
		MC183:  rand.Float32() * 100.182,
		MC184:  rand.Float32() * 100.183,
		MC185:  rand.Float32() * 100.184,
		MC186:  rand.Float32() * 100.185,
		MC187:  rand.Float32() * 100.186,
		MC188:  rand.Float32() * 100.187,
		MC189:  rand.Float32() * 100.188,
		MC190:  rand.Float32() * 100.189,
		MC191:  rand.Float32() * 100.190,
		MC192:  rand.Float32() * 100.191,
		MC193:  rand.Float32() * 100.192,
		MC194:  rand.Float32() * 100.193,
		MC195:  rand.Float32() * 100.194,
		MC196:  rand.Float32() * 100.195,
		MC197:  rand.Float32() * 100.196,
		MC198:  rand.Float32() * 100.197,
		MC199:  rand.Float32() * 100.198,
		MC200:  rand.Float32() * 100.199,
		MC201:  rand.Float32() * 100.200,
		MC202:  rand.Float32() * 100.201,
		MC203:  rand.Float32() * 100.202,
		MC204:  rand.Float32() * 100.203,
		MC205:  rand.Float32() * 100.204,
		MC206:  rand.Float32() * 100.205,
		MC207:  rand.Float32() * 100.206,
		MC208:  rand.Float32() * 100.207,
		MC209:  rand.Float32() * 100.208,
		MC210:  rand.Float32() * 100.209,
		MC211:  rand.Float32() * 100.210,
		MC212:  rand.Float32() * 100.211,
		MC213:  rand.Float32() * 100.212,
		MC214:  rand.Float32() * 100.213,
		MC215:  rand.Float32() * 100.214,
		MC216:  rand.Float32() * 100.215,
		MC217:  rand.Float32() * 100.216,
		MC218:  rand.Float32() * 100.217,
		MC219:  rand.Float32() * 100.218,
		MC220:  rand.Float32() * 100.219,
		MC221:  rand.Float32() * 100.220,
		MC222:  rand.Float32() * 100.221,
		MC223:  rand.Float32() * 100.222,
		MC224:  rand.Float32() * 100.223,
		MC225:  rand.Float32() * 100.224,
		MC226:  rand.Float32() * 100.225,
		MC227:  rand.Float32() * 100.226,
		MC228:  rand.Float32() * 100.227,
		MC229:  rand.Float32() * 100.228,
		MC230:  rand.Float32() * 100.229,
		MC231:  rand.Float32() * 100.230,
		MC232:  rand.Float32() * 100.231,
		MC233:  rand.Float32() * 100.232,
		MC234:  rand.Float32() * 100.233,
		MC235:  rand.Float32() * 100.234,
		MC236:  rand.Float32() * 100.235,
		MC237:  rand.Float32() * 100.236,
		MC238:  rand.Float32() * 100.237,
		MC239:  rand.Float32() * 100.238,
		MC240:  rand.Float32() * 100.239,
		MC241:  rand.Float32() * 100.240,
		MC242:  rand.Float32() * 100.241,
		MC243:  rand.Float32() * 100.242,
		MC244:  rand.Float32() * 100.243,
		MC245:  rand.Float32() * 100.244,
		MC246:  rand.Float32() * 100.245,
		MC247:  rand.Float32() * 100.246,
		MC248:  rand.Float32() * 100.247,
		MC249:  rand.Float32() * 100.248,
		MC250:  rand.Float32() * 100.249,
		MC251:  rand.Float32() * 100.250,
		MC252:  rand.Float32() * 100.251,
		MC253:  rand.Float32() * 100.252,
		MC254:  rand.Float32() * 100.253,
		MC255:  rand.Float32() * 100.254,
		MC256:  rand.Float32() * 100.255,
		MC257:  rand.Float32() * 100.256,
		MC258:  rand.Float32() * 100.257,
		MC259:  rand.Float32() * 100.258,
		MC260:  rand.Float32() * 100.259,
		MC261:  rand.Float32() * 100.260,
		MC262:  rand.Float32() * 100.261,
		MC263:  rand.Float32() * 100.262,
		MC264:  rand.Float32() * 100.263,
		MC265:  rand.Float32() * 100.264,
		MC266:  rand.Float32() * 100.265,
		MC267:  rand.Float32() * 100.266,
		MC268:  rand.Float32() * 100.267,
		MC269:  rand.Float32() * 100.268,
		MC270:  rand.Float32() * 100.269,
		MC271:  rand.Float32() * 100.270,
		MC272:  rand.Float32() * 100.271,
		MC273:  rand.Float32() * 100.272,
		MC274:  rand.Float32() * 100.273,
		MC275:  rand.Float32() * 100.274,
		MC276:  rand.Float32() * 100.275,
		MC277:  rand.Float32() * 100.276,
		MC278:  rand.Float32() * 100.277,
		MC279:  rand.Float32() * 100.278,
		MC280:  rand.Float32() * 100.279,
		MC281:  rand.Float32() * 100.280,
		MC282:  rand.Float32() * 100.281,
		MC283:  rand.Float32() * 100.282,
		MC284:  rand.Float32() * 100.283,
		MC285:  rand.Float32() * 100.284,
		MC286:  rand.Float32() * 100.285,
		MC287:  rand.Float32() * 100.286,
		MC288:  rand.Float32() * 100.287,
		MC289:  rand.Float32() * 100.288,
		MC290:  rand.Float32() * 100.289,
		MC291:  rand.Float32() * 100.290,
		MC292:  rand.Float32() * 100.291,
		MC293:  rand.Float32() * 100.292,
		MC294:  rand.Float32() * 100.293,
		MC295:  rand.Float32() * 100.294,
		MC296:  rand.Float32() * 100.295,
		MC297:  rand.Float32() * 100.296,
		MC298:  rand.Float32() * 100.297,
		MC299:  rand.Float32() * 100.298,
		MC300:  rand.Float32() * 100.299,
		MC301:  rand.Float32() * 100.300,
		MC302:  rand.Float32() * 100.301,
		MC303:  rand.Float32() * 100.302,
		MC304:  rand.Float32() * 100.303,
		MC305:  rand.Float32() * 100.304,
		MC306:  rand.Float32() * 100.305,
		MC307:  rand.Float32() * 100.306,
		MC308:  rand.Float32() * 100.307,
		MC309:  rand.Float32() * 100.308,
		MC310:  rand.Float32() * 100.309,
		MC311:  rand.Float32() * 100.310,
		MC312:  rand.Float32() * 100.311,
		MC313:  rand.Float32() * 100.312,
		MC314:  rand.Float32() * 100.313,
		MC315:  rand.Float32() * 100.314,
		MC316:  rand.Float32() * 100.315,
		MC317:  rand.Float32() * 100.316,
		MC318:  rand.Float32() * 100.317,
		MC319:  rand.Float32() * 100.318,
		MC320:  rand.Float32() * 100.319,
		MC321:  rand.Float32() * 100.320,
		MC322:  rand.Float32() * 100.321,
		MC323:  rand.Float32() * 100.322,
		MC324:  rand.Float32() * 100.323,
		MC325:  rand.Float32() * 100.324,
		MC326:  rand.Float32() * 100.325,
		MC327:  rand.Float32() * 100.326,
		MC328:  rand.Float32() * 100.327,
		MC329:  rand.Float32() * 100.328,
		MC330:  rand.Float32() * 100.329,
		MC331:  rand.Float32() * 100.330,
		MC332:  rand.Float32() * 100.331,
		MC333:  rand.Float32() * 100.332,
		MC334:  rand.Float32() * 100.333,
		MC335:  rand.Float32() * 100.334,
		MC336:  rand.Float32() * 100.335,
		MC337:  rand.Float32() * 100.336,
		MC338:  rand.Float32() * 100.337,
		MC339:  rand.Float32() * 100.338,
		MC340:  rand.Float32() * 100.339,
		MC341:  rand.Float32() * 100.340,
		MC342:  rand.Float32() * 100.341,
		MC343:  rand.Float32() * 100.342,
		MC344:  rand.Float32() * 100.343,
		MC345:  rand.Float32() * 100.344,
		MC346:  rand.Float32() * 100.345,
		MC347:  rand.Float32() * 100.346,
		MC348:  rand.Float32() * 100.347,
		MC349:  rand.Float32() * 100.348,
		MC350:  rand.Float32() * 100.349,
		MC351:  rand.Float32() * 100.350,
		MC352:  rand.Float32() * 100.351,
		MC353:  rand.Float32() * 100.352,
		MC354:  rand.Float32() * 100.353,
		MC355:  rand.Float32() * 100.354,
		MC356:  rand.Float32() * 100.355,
		MC357:  rand.Float32() * 100.356,
		MC358:  rand.Float32() * 100.357,
		MC359:  rand.Float32() * 100.358,
		MC360:  rand.Float32() * 100.359,
		MC361:  rand.Float32() * 100.360,
		MC362:  rand.Float32() * 100.361,
		MC363:  rand.Float32() * 100.362,
		MC364:  rand.Float32() * 100.363,
		MC365:  rand.Float32() * 100.364,
		MC366:  rand.Float32() * 100.365,
		MC367:  rand.Float32() * 100.366,
		MC368:  rand.Float32() * 100.367,
		MC369:  rand.Float32() * 100.368,
		MC370:  rand.Float32() * 100.369,
		MC371:  rand.Float32() * 100.370,
		MC372:  rand.Float32() * 100.371,
		MC373:  rand.Float32() * 100.372,
		MC374:  rand.Float32() * 100.373,
		MC375:  rand.Float32() * 100.374,
		MC376:  rand.Float32() * 100.375,
		MC377:  rand.Float32() * 100.376,
		MC378:  rand.Float32() * 100.377,
		MC379:  rand.Float32() * 100.378,
		MC380:  rand.Float32() * 100.379,
		MC381:  rand.Float32() * 100.380,
		MC382:  rand.Float32() * 100.381,
		MC383:  rand.Float32() * 100.382,
		MC384:  rand.Float32() * 100.383,
		MC385:  rand.Float32() * 100.384,
		MC386:  rand.Float32() * 100.385,
		MC387:  rand.Float32() * 100.386,
		MC388:  rand.Float32() * 100.387,
		MC389:  rand.Float32() * 100.388,
		MC390:  rand.Float32() * 100.389,
		MC391:  rand.Float32() * 100.390,
		MC392:  rand.Float32() * 100.391,
		MC393:  rand.Float32() * 100.392,
		MC394:  rand.Float32() * 100.393,
		MC395:  rand.Float32() * 100.394,
		MC396:  rand.Float32() * 100.395,
		MC397:  rand.Float32() * 100.396,
		MC398:  rand.Float32() * 100.397,
		MC399:  rand.Float32() * 100.398,
		MC400:  rand.Float32() * 100.399,
		MC401:  rand.Float32() * 100.400,
		MC402:  rand.Float32() * 100.401,
		MC403:  rand.Float32() * 100.402,
		MC404:  rand.Float32() * 100.403,
		MC405:  rand.Float32() * 100.404,
		MC406:  rand.Float32() * 100.405,
		MC407:  rand.Float32() * 100.406,
		MC408:  rand.Float32() * 100.407,
		MC409:  rand.Float32() * 100.408,
		MC410:  rand.Float32() * 100.409,
		MC411:  rand.Float32() * 100.410,
		MC412:  rand.Float32() * 100.411,
		MC413:  rand.Float32() * 100.412,
		MC414:  rand.Float32() * 100.413,
		MC415:  rand.Float32() * 100.414,
		MC416:  rand.Float32() * 100.415,
		MC417:  rand.Float32() * 100.416,
		MC418:  rand.Float32() * 100.417,
		MC419:  rand.Float32() * 100.418,
		MC420:  rand.Float32() * 100.419,
		MC421:  rand.Float32() * 100.420,
		MC422:  rand.Float32() * 100.421,
		MC423:  rand.Float32() * 100.422,
		MC424:  rand.Float32() * 100.423,
		MC425:  rand.Float32() * 100.424,
		MC426:  rand.Float32() * 100.425,
		MC427:  rand.Float32() * 100.426,
		MC428:  rand.Float32() * 100.427,
		MC429:  rand.Float32() * 100.428,
		MC430:  rand.Float32() * 100.429,
		MC431:  rand.Float32() * 100.430,
		MC432:  rand.Float32() * 100.431,
		MC433:  rand.Float32() * 100.432,
		MC434:  rand.Float32() * 100.433,
		MC435:  rand.Float32() * 100.434,
		MC436:  rand.Float32() * 100.435,
		MC437:  rand.Float32() * 100.436,
		MC438:  rand.Float32() * 100.437,
		MC439:  rand.Float32() * 100.438,
		MC440:  rand.Float32() * 100.439,
		MC441:  rand.Float32() * 100.440,
		MC442:  rand.Float32() * 100.441,
		MC443:  rand.Float32() * 100.442,
		MC444:  rand.Float32() * 100.443,
		MC445:  rand.Float32() * 100.444,
		MC446:  rand.Float32() * 100.445,
		MC447:  rand.Float32() * 100.446,
		MC448:  rand.Float32() * 100.447,
		MC449:  rand.Float32() * 100.448,
		MC450:  rand.Float32() * 100.449,
		MC451:  rand.Float32() * 100.450,
		MC452:  rand.Float32() * 100.0,
		MC453:  rand.Float32() * 100.1,
		MC454:  rand.Float32() * 100.2,
		MC455:  rand.Float32() * 100.3,
		MC456:  rand.Float32() * 100.4,
		MC457:  rand.Float32() * 100.5,
		MC458:  rand.Float32() * 100.6,
		MC459:  rand.Float32() * 100.7,
		MC460:  rand.Float32() * 100.8,
		MC461:  rand.Float32() * 100.9,
		MC462:  rand.Float32() * 100.10,
		MC463:  rand.Float32() * 100.11,
		MC464:  rand.Float32() * 100.12,
		MC465:  rand.Float32() * 100.13,
		MC466:  rand.Float32() * 100.14,
		MC467:  rand.Float32() * 100.15,
		MC468:  rand.Float32() * 100.16,
		MC469:  rand.Float32() * 100.17,
		MC470:  rand.Float32() * 100.18,
		MC471:  rand.Float32() * 100.19,
		MC472:  rand.Float32() * 100.20,
		MC473:  rand.Float32() * 100.21,
		MC474:  rand.Float32() * 100.22,
		MC475:  rand.Float32() * 100.23,
		MC476:  rand.Float32() * 100.24,
		MC477:  rand.Float32() * 100.25,
		MC478:  rand.Float32() * 100.26,
		MC479:  rand.Float32() * 100.27,
		MC480:  rand.Float32() * 100.28,
		MC481:  rand.Float32() * 100.29,
		MC482:  rand.Float32() * 100.30,
		MC483:  rand.Float32() * 100.31,
		MC484:  rand.Float32() * 100.32,
		MC485:  rand.Float32() * 100.33,
		MC486:  rand.Float32() * 100.34,
		MC487:  rand.Float32() * 100.35,
		MC488:  rand.Float32() * 100.36,
		MC489:  rand.Float32() * 100.37,
		MC490:  rand.Float32() * 100.38,
		MC491:  rand.Float32() * 100.39,
		MC492:  rand.Float32() * 100.40,
		MC493:  rand.Float32() * 100.41,
		MC494:  rand.Float32() * 100.42,
		MC495:  rand.Float32() * 100.43,
		MC496:  rand.Float32() * 100.44,
		MC497:  rand.Float32() * 100.45,
		MC498:  rand.Float32() * 100.46,
		MC499:  rand.Float32() * 100.47,
		MC500:  rand.Float32() * 100.48,
		MC501:  rand.Float32() * 100.49,
		MC502:  rand.Float32() * 100.50,
		MC503:  rand.Float32() * 100.51,
		MC504:  rand.Float32() * 100.52,
		MC505:  rand.Float32() * 100.53,
		MC506:  rand.Float32() * 100.54,
		MC507:  rand.Float32() * 100.55,
		MC508:  rand.Float32() * 100.56,
		MC509:  rand.Float32() * 100.57,
		MC510:  rand.Float32() * 100.58,
		MC511:  rand.Float32() * 100.59,
		MC512:  rand.Float32() * 100.60,
		MC513:  rand.Float32() * 100.61,
		MC514:  rand.Float32() * 100.62,
		MC515:  rand.Float32() * 100.63,
		MC516:  rand.Float32() * 100.64,
		MC517:  rand.Float32() * 100.65,
		MC518:  rand.Float32() * 100.66,
		MC519:  rand.Float32() * 100.67,
		MC520:  rand.Float32() * 100.68,
		MC521:  rand.Float32() * 100.69,
		MC522:  rand.Float32() * 100.70,
		MC523:  rand.Float32() * 100.71,
		MC524:  rand.Float32() * 100.72,
		MC525:  rand.Float32() * 100.73,
		MC526:  rand.Float32() * 100.74,
		MC527:  rand.Float32() * 100.75,
		MC528:  rand.Float32() * 100.76,
		MC529:  rand.Float32() * 100.77,
		MC530:  rand.Float32() * 100.78,
		MC531:  rand.Float32() * 100.79,
		MC532:  rand.Float32() * 100.80,
		MC533:  rand.Float32() * 100.81,
		MC534:  rand.Float32() * 100.82,
		MC535:  rand.Float32() * 100.83,
		MC536:  rand.Float32() * 100.84,
		MC537:  rand.Float32() * 100.85,
		MC538:  rand.Float32() * 100.86,
		MC539:  rand.Float32() * 100.87,
		MC540:  rand.Float32() * 100.88,
		MC541:  rand.Float32() * 100.89,
		MC542:  rand.Float32() * 100.90,
		MC543:  rand.Float32() * 100.91,
		MC544:  rand.Float32() * 100.92,
		MC545:  rand.Float32() * 100.93,
		MC546:  rand.Float32() * 100.94,
		MC547:  rand.Float32() * 100.95,
		MC548:  rand.Float32() * 100.96,
		MC549:  rand.Float32() * 100.97,
		MC550:  rand.Float32() * 100.98,
		MC551:  rand.Float32() * 100.99,
		MC552:  rand.Float32() * 100.100,
		MC553:  rand.Float32() * 100.101,
		MC554:  rand.Float32() * 100.102,
		MC555:  rand.Float32() * 100.103,
		MC556:  rand.Float32() * 100.104,
		MC557:  rand.Float32() * 100.105,
		MC558:  rand.Float32() * 100.106,
		MC559:  rand.Float32() * 100.107,
		MC560:  rand.Float32() * 100.108,
		MC561:  rand.Float32() * 100.109,
		MC562:  rand.Float32() * 100.110,
		MC563:  rand.Float32() * 100.111,
		MC564:  rand.Float32() * 100.112,
		MC565:  rand.Float32() * 100.113,
		MC566:  rand.Float32() * 100.114,
		MC567:  rand.Float32() * 100.115,
		MC568:  rand.Float32() * 100.116,
		MC569:  rand.Float32() * 100.117,
		MC570:  rand.Float32() * 100.118,
		MC571:  rand.Float32() * 100.119,
		MC572:  rand.Float32() * 100.120,
		MC573:  rand.Float32() * 100.121,
		MC574:  rand.Float32() * 100.122,
		MC575:  rand.Float32() * 100.123,
		MC576:  rand.Float32() * 100.124,
		MC577:  rand.Float32() * 100.125,
		MC578:  rand.Float32() * 100.126,
		MC579:  rand.Float32() * 100.127,
		MC580:  rand.Float32() * 100.128,
		MC581:  rand.Float32() * 100.129,
		MC582:  rand.Float32() * 100.130,
		MC583:  rand.Float32() * 100.131,
		MC584:  rand.Float32() * 100.132,
		MC585:  rand.Float32() * 100.133,
		MC586:  rand.Float32() * 100.134,
		MC587:  rand.Float32() * 100.135,
		MC588:  rand.Float32() * 100.136,
		MC589:  rand.Float32() * 100.137,
		MC590:  rand.Float32() * 100.138,
		MC591:  rand.Float32() * 100.139,
		MC592:  rand.Float32() * 100.140,
		MC593:  rand.Float32() * 100.141,
		MC594:  rand.Float32() * 100.142,
		MC595:  rand.Float32() * 100.143,
		MC596:  rand.Float32() * 100.144,
		MC597:  rand.Float32() * 100.145,
		MC598:  rand.Float32() * 100.146,
		MC599:  rand.Float32() * 100.147,
		MC600:  rand.Float32() * 100.148,
		MC601:  rand.Float32() * 100.149,
		MC602:  rand.Float32() * 100.150,
		MC603:  rand.Float32() * 100.151,
		MC604:  rand.Float32() * 100.152,
		MC605:  rand.Float32() * 100.153,
		MC606:  rand.Float32() * 100.154,
		MC607:  rand.Float32() * 100.155,
		MC608:  rand.Float32() * 100.156,
		MC609:  rand.Float32() * 100.157,
		MC610:  rand.Float32() * 100.158,
		MC611:  rand.Float32() * 100.159,
		MC612:  rand.Float32() * 100.160,
		MC613:  rand.Float32() * 100.161,
		MC614:  rand.Float32() * 100.162,
		MC615:  rand.Float32() * 100.163,
		MC616:  rand.Float32() * 100.164,
		MC617:  rand.Float32() * 100.165,
		MC618:  rand.Float32() * 100.166,
		MC619:  rand.Float32() * 100.167,
		MC620:  rand.Float32() * 100.168,
		MC621:  rand.Float32() * 100.169,
		MC622:  rand.Float32() * 100.170,
		MC623:  rand.Float32() * 100.171,
		MC624:  rand.Float32() * 100.172,
		MC625:  rand.Float32() * 100.173,
		MC626:  rand.Float32() * 100.174,
		MC627:  rand.Float32() * 100.175,
		MC628:  rand.Float32() * 100.176,
		MC629:  rand.Float32() * 100.177,
		MC630:  rand.Float32() * 100.178,
		MC631:  rand.Float32() * 100.179,
		MC632:  rand.Float32() * 100.180,
		MC633:  rand.Float32() * 100.181,
		MC634:  rand.Float32() * 100.182,
		MC635:  rand.Float32() * 100.183,
		MC636:  rand.Float32() * 100.184,
		MC637:  rand.Float32() * 100.185,
		MC638:  rand.Float32() * 100.186,
		MC639:  rand.Float32() * 100.187,
		MC640:  rand.Float32() * 100.188,
		MC641:  rand.Float32() * 100.189,
		MC642:  rand.Float32() * 100.190,
		MC643:  rand.Float32() * 100.191,
		MC644:  rand.Float32() * 100.192,
		MC645:  rand.Float32() * 100.193,
		MC646:  rand.Float32() * 100.194,
		MC647:  rand.Float32() * 100.195,
		MC648:  rand.Float32() * 100.196,
		MC649:  rand.Float32() * 100.197,
		MC650:  rand.Float32() * 100.198,
		MC651:  rand.Float32() * 100.199,
		MC652:  rand.Float32() * 100.200,
		MC653:  rand.Float32() * 100.201,
		MC654:  rand.Float32() * 100.202,
		MC655:  rand.Float32() * 100.203,
		MC656:  rand.Float32() * 100.204,
		MC657:  rand.Float32() * 100.205,
		MC658:  rand.Float32() * 100.206,
		MC659:  rand.Float32() * 100.207,
		MC660:  rand.Float32() * 100.208,
		MC661:  rand.Float32() * 100.209,
		MC662:  rand.Float32() * 100.210,
		MC663:  rand.Float32() * 100.211,
		MC664:  rand.Float32() * 100.212,
		MC665:  rand.Float32() * 100.213,
		MC666:  rand.Float32() * 100.214,
		MC667:  rand.Float32() * 100.215,
		MC668:  rand.Float32() * 100.216,
		MC669:  rand.Float32() * 100.217,
		MC670:  rand.Float32() * 100.218,
		MC671:  rand.Float32() * 100.219,
		MC672:  rand.Float32() * 100.220,
		MC673:  rand.Float32() * 100.221,
		MC674:  rand.Float32() * 100.222,
		MC675:  rand.Float32() * 100.223,
		MC676:  rand.Float32() * 100.224,
		MC677:  rand.Float32() * 100.225,
		MC678:  rand.Float32() * 100.226,
		MC679:  rand.Float32() * 100.227,
		MC680:  rand.Float32() * 100.228,
		MC681:  rand.Float32() * 100.229,
		MC682:  rand.Float32() * 100.230,
		MC683:  rand.Float32() * 100.231,
		MC684:  rand.Float32() * 100.232,
		MC685:  rand.Float32() * 100.233,
		MC686:  rand.Float32() * 100.234,
		MC687:  rand.Float32() * 100.235,
		MC688:  rand.Float32() * 100.236,
		MC689:  rand.Float32() * 100.237,
		MC690:  rand.Float32() * 100.238,
		MC691:  rand.Float32() * 100.239,
		MC692:  rand.Float32() * 100.240,
		MC693:  rand.Float32() * 100.241,
		MC694:  rand.Float32() * 100.242,
		MC695:  rand.Float32() * 100.243,
		MC696:  rand.Float32() * 100.244,
		MC697:  rand.Float32() * 100.245,
		MC698:  rand.Float32() * 100.246,
		MC699:  rand.Float32() * 100.247,
		MC700:  rand.Float32() * 100.248,
		MC701:  rand.Float32() * 100.249,
		MC702:  rand.Float32() * 100.250,
		MC703:  rand.Float32() * 100.251,
		MC704:  rand.Float32() * 100.252,
		MC705:  rand.Float32() * 100.253,
		MC706:  rand.Float32() * 100.254,
		MC707:  rand.Float32() * 100.255,
		MC708:  rand.Float32() * 100.256,
		MC709:  rand.Float32() * 100.257,
		MC710:  rand.Float32() * 100.258,
		MC711:  rand.Float32() * 100.259,
		MC712:  rand.Float32() * 100.260,
		MC713:  rand.Float32() * 100.261,
		MC714:  rand.Float32() * 100.262,
		MC715:  rand.Float32() * 100.263,
		MC716:  rand.Float32() * 100.264,
		MC717:  rand.Float32() * 100.265,
		MC718:  rand.Float32() * 100.266,
		MC719:  rand.Float32() * 100.267,
		MC720:  rand.Float32() * 100.268,
		MC721:  rand.Float32() * 100.269,
		MC722:  rand.Float32() * 100.270,
		MC723:  rand.Float32() * 100.271,
		MC724:  rand.Float32() * 100.272,
		MC725:  rand.Float32() * 100.273,
		MC726:  rand.Float32() * 100.274,
		MC727:  rand.Float32() * 100.275,
		MC728:  rand.Float32() * 100.276,
		MC729:  rand.Float32() * 100.277,
		MC730:  rand.Float32() * 100.278,
		MC731:  rand.Float32() * 100.279,
		MC732:  rand.Float32() * 100.280,
		MC733:  rand.Float32() * 100.281,
		MC734:  rand.Float32() * 100.282,
		MC735:  rand.Float32() * 100.283,
		MC736:  rand.Float32() * 100.284,
		MC737:  rand.Float32() * 100.285,
		MC738:  rand.Float32() * 100.286,
		MC739:  rand.Float32() * 100.287,
		MC740:  rand.Float32() * 100.288,
		MC741:  rand.Float32() * 100.289,
		MC742:  rand.Float32() * 100.290,
		MC743:  rand.Float32() * 100.291,
		MC744:  rand.Float32() * 100.292,
		MC745:  rand.Float32() * 100.293,
		MC746:  rand.Float32() * 100.294,
		MC747:  rand.Float32() * 100.295,
		MC748:  rand.Float32() * 100.296,
		MC749:  rand.Float32() * 100.297,
		MC750:  rand.Float32() * 100.298,
		MC751:  rand.Float32() * 100.299,
		MC752:  rand.Float32() * 100.300,
		MC753:  rand.Float32() * 100.301,
		MC754:  rand.Float32() * 100.302,
		MC755:  rand.Float32() * 100.303,
		MC756:  rand.Float32() * 100.304,
		MC757:  rand.Float32() * 100.305,
		MC758:  rand.Float32() * 100.306,
		MC759:  rand.Float32() * 100.307,
		MC760:  rand.Float32() * 100.308,
		MC761:  rand.Float32() * 100.309,
		MC762:  rand.Float32() * 100.310,
		MC763:  rand.Float32() * 100.311,
		MC764:  rand.Float32() * 100.312,
		MC765:  rand.Float32() * 100.313,
		MC766:  rand.Float32() * 100.314,
		MC767:  rand.Float32() * 100.315,
		MC768:  rand.Float32() * 100.316,
		MC769:  rand.Float32() * 100.317,
		MC770:  rand.Float32() * 100.318,
		MC771:  rand.Float32() * 100.319,
		MC772:  rand.Float32() * 100.320,
		MC773:  rand.Float32() * 100.321,
		MC774:  rand.Float32() * 100.322,
		MC775:  rand.Float32() * 100.323,
		MC776:  rand.Float32() * 100.324,
		MC777:  rand.Float32() * 100.325,
		MC778:  rand.Float32() * 100.326,
		MC779:  rand.Float32() * 100.327,
		MC780:  rand.Float32() * 100.328,
		MC781:  rand.Float32() * 100.329,
		MC782:  rand.Float32() * 100.330,
		MC783:  rand.Float32() * 100.331,
		MC784:  rand.Float32() * 100.332,
		MC785:  rand.Float32() * 100.333,
		MC786:  rand.Float32() * 100.334,
		MC787:  rand.Float32() * 100.335,
		MC788:  rand.Float32() * 100.336,
		MC789:  rand.Float32() * 100.337,
		MC790:  rand.Float32() * 100.338,
		MC791:  rand.Float32() * 100.339,
		MC792:  rand.Float32() * 100.340,
		MC793:  rand.Float32() * 100.341,
		MC794:  rand.Float32() * 100.342,
		MC795:  rand.Float32() * 100.343,
		MC796:  rand.Float32() * 100.344,
		MC797:  rand.Float32() * 100.345,
		MC798:  rand.Float32() * 100.346,
		MC799:  rand.Float32() * 100.347,
		MC800:  rand.Float32() * 100.348,
		MC801:  rand.Float32() * 100.349,
		MC802:  rand.Float32() * 100.350,
		MC803:  rand.Float32() * 100.351,
		MC804:  rand.Float32() * 100.352,
		MC805:  rand.Float32() * 100.353,
		MC806:  rand.Float32() * 100.354,
		MC807:  rand.Float32() * 100.355,
		MC808:  rand.Float32() * 100.356,
		MC809:  rand.Float32() * 100.357,
		MC810:  rand.Float32() * 100.358,
		MC811:  rand.Float32() * 100.359,
		MC812:  rand.Float32() * 100.360,
		MC813:  rand.Float32() * 100.361,
		MC814:  rand.Float32() * 100.362,
		MC815:  rand.Float32() * 100.363,
		MC816:  rand.Float32() * 100.364,
		MC817:  rand.Float32() * 100.365,
		MC818:  rand.Float32() * 100.366,
		MC819:  rand.Float32() * 100.367,
		MC820:  rand.Float32() * 100.368,
		MC821:  rand.Float32() * 100.369,
		MC822:  rand.Float32() * 100.370,
		MC823:  rand.Float32() * 100.371,
		MC824:  rand.Float32() * 100.372,
		MC825:  rand.Float32() * 100.373,
		MC826:  rand.Float32() * 100.374,
		MC827:  rand.Float32() * 100.375,
		MC828:  rand.Float32() * 100.376,
		MC829:  rand.Float32() * 100.377,
		MC830:  rand.Float32() * 100.378,
		MC831:  rand.Float32() * 100.379,
		MC832:  rand.Float32() * 100.380,
		MC833:  rand.Float32() * 100.381,
		MC834:  rand.Float32() * 100.382,
		MC835:  rand.Float32() * 100.383,
		MC836:  rand.Float32() * 100.384,
		MC837:  rand.Float32() * 100.385,
		MC838:  rand.Float32() * 100.386,
		MC839:  rand.Float32() * 100.387,
		MC840:  rand.Float32() * 100.388,
		MC841:  rand.Float32() * 100.389,
		MC842:  rand.Float32() * 100.390,
		MC843:  rand.Float32() * 100.391,
		MC844:  rand.Float32() * 100.392,
		MC845:  rand.Float32() * 100.393,
		MC846:  rand.Float32() * 100.394,
		MC847:  rand.Float32() * 100.395,
		MC848:  rand.Float32() * 100.396,
		MC849:  rand.Float32() * 100.397,
		MC850:  rand.Float32() * 100.398,
		MC851:  rand.Float32() * 100.399,
		MC852:  rand.Float32() * 100.400,
		MC853:  rand.Float32() * 100.401,
		MC854:  rand.Float32() * 100.402,
		MC855:  rand.Float32() * 100.403,
		MC856:  rand.Float32() * 100.404,
		MC857:  rand.Float32() * 100.405,
		MC858:  rand.Float32() * 100.406,
		MC859:  rand.Float32() * 100.407,
		MC860:  rand.Float32() * 100.408,
		MC861:  rand.Float32() * 100.409,
		MC862:  rand.Float32() * 100.410,
		MC863:  rand.Float32() * 100.411,
		MC864:  rand.Float32() * 100.412,
		MC865:  rand.Float32() * 100.413,
		MC866:  rand.Float32() * 100.414,
		MC867:  rand.Float32() * 100.415,
		MC868:  rand.Float32() * 100.416,
		MC869:  rand.Float32() * 100.417,
		MC870:  rand.Float32() * 100.418,
		MC871:  rand.Float32() * 100.419,
		MC872:  rand.Float32() * 100.420,
		MC873:  rand.Float32() * 100.421,
		MC874:  rand.Float32() * 100.422,
		MC875:  rand.Float32() * 100.423,
		MC876:  rand.Float32() * 100.424,
		MC877:  rand.Float32() * 100.425,
		MC878:  rand.Float32() * 100.426,
		MC879:  rand.Float32() * 100.427,
		MC880:  rand.Float32() * 100.428,
		MC881:  rand.Float32() * 100.429,
		MC882:  rand.Float32() * 100.430,
		MC883:  rand.Float32() * 100.431,
		MC884:  rand.Float32() * 100.432,
		MC885:  rand.Float32() * 100.433,
		MC886:  rand.Float32() * 100.434,
		MC887:  rand.Float32() * 100.435,
		MC888:  rand.Float32() * 100.436,
		MC889:  rand.Float32() * 100.437,
		MC890:  rand.Float32() * 100.438,
		MC891:  rand.Float32() * 100.439,
		MC892:  rand.Float32() * 100.440,
		MC893:  rand.Float32() * 100.441,
		MC894:  rand.Float32() * 100.442,
		MC895:  rand.Float32() * 100.443,
		MC896:  rand.Float32() * 100.444,
		MC897:  rand.Float32() * 100.445,
		MC898:  rand.Float32() * 100.446,
		MC899:  rand.Float32() * 100.447,
		MC900:  rand.Float32() * 100.448,
		MC901:  rand.Float32() * 100.449,
		MC902:  rand.Float32() * 100.450,
		MC903:  rand.Float32() * 100.0,
		MC904:  rand.Float32() * 100.1,
		MC905:  rand.Float32() * 100.2,
		MC906:  rand.Float32() * 100.3,
		MC907:  rand.Float32() * 100.4,
		MC908:  rand.Float32() * 100.5,
		MC909:  rand.Float32() * 100.6,
		MC910:  rand.Float32() * 100.7,
		MC911:  rand.Float32() * 100.8,
		MC912:  rand.Float32() * 100.9,
		MC913:  rand.Float32() * 100.10,
		MC914:  rand.Float32() * 100.11,
		MC915:  rand.Float32() * 100.12,
		MC916:  rand.Float32() * 100.13,
		MC917:  rand.Float32() * 100.14,
		MC918:  rand.Float32() * 100.15,
		MC919:  rand.Float32() * 100.16,
		MC920:  rand.Float32() * 100.17,
		MC921:  rand.Float32() * 100.18,
		MC922:  rand.Float32() * 100.19,
		MC923:  rand.Float32() * 100.20,
		MC924:  rand.Float32() * 100.21,
		MC925:  rand.Float32() * 100.22,
		MC926:  rand.Float32() * 100.23,
		MC927:  rand.Float32() * 100.24,
		MC928:  rand.Float32() * 100.25,
		MC929:  rand.Float32() * 100.26,
		MC930:  rand.Float32() * 100.27,
		MC931:  rand.Float32() * 100.28,
		MC932:  rand.Float32() * 100.29,
		MC933:  rand.Float32() * 100.30,
		MC934:  rand.Float32() * 100.31,
		MC935:  rand.Float32() * 100.32,
		MC936:  rand.Float32() * 100.33,
		MC937:  rand.Float32() * 100.34,
		MC938:  rand.Float32() * 100.35,
		MC939:  rand.Float32() * 100.36,
		MC940:  rand.Float32() * 100.37,
		MC941:  rand.Float32() * 100.38,
		MC942:  rand.Float32() * 100.39,
		MC943:  rand.Float32() * 100.40,
		MC944:  rand.Float32() * 100.41,
		MC945:  rand.Float32() * 100.42,
		MC946:  rand.Float32() * 100.43,
		MC947:  rand.Float32() * 100.44,
		MC948:  rand.Float32() * 100.45,
		MC949:  rand.Float32() * 100.46,
		MC950:  rand.Float32() * 100.47,
		MC951:  rand.Float32() * 100.48,
		MC952:  rand.Float32() * 100.49,
		MC953:  rand.Float32() * 100.50,
		MC954:  rand.Float32() * 100.51,
		MC955:  rand.Float32() * 100.52,
		MC956:  rand.Float32() * 100.53,
		MC957:  rand.Float32() * 100.54,
		MC958:  rand.Float32() * 100.55,
		MC959:  rand.Float32() * 100.56,
		MC960:  rand.Float32() * 100.57,
		MC961:  rand.Float32() * 100.58,
		MC962:  rand.Float32() * 100.59,
		MC963:  rand.Float32() * 100.60,
		MC964:  rand.Float32() * 100.61,
		MC965:  rand.Float32() * 100.62,
		MC966:  rand.Float32() * 100.63,
		MC967:  rand.Float32() * 100.64,
		MC968:  rand.Float32() * 100.65,
		MC969:  rand.Float32() * 100.66,
		MC970:  rand.Float32() * 100.67,
		MC971:  rand.Float32() * 100.68,
		MC972:  rand.Float32() * 100.69,
		MC973:  rand.Float32() * 100.70,
		MC974:  rand.Float32() * 100.71,
		MC975:  rand.Float32() * 100.72,
		MC976:  rand.Float32() * 100.73,
		MC977:  rand.Float32() * 100.74,
		MC978:  rand.Float32() * 100.75,
		MC979:  rand.Float32() * 100.76,
		MC980:  rand.Float32() * 100.77,
		MC981:  rand.Float32() * 100.78,
		MC982:  rand.Float32() * 100.79,
		MC983:  rand.Float32() * 100.80,
		MC984:  rand.Float32() * 100.81,
		MC985:  rand.Float32() * 100.82,
		MC986:  rand.Float32() * 100.83,
		MC987:  rand.Float32() * 100.84,
		MC988:  rand.Float32() * 100.85,
		MC989:  rand.Float32() * 100.86,
		MC990:  rand.Float32() * 100.87,
		MC991:  rand.Float32() * 100.88,
		MC992:  rand.Float32() * 100.89,
		MC993:  rand.Float32() * 100.90,
		MC994:  rand.Float32() * 100.91,
		MC995:  rand.Float32() * 100.92,
		MC996:  rand.Float32() * 100.93,
		MC997:  rand.Float32() * 100.94,
		MC998:  rand.Float32() * 100.95,
		MC999:  rand.Float32() * 100.96,
		MC1000: rand.Float32() * 100.97,
	}
}

// Calculate average for specified fields
func calculateAverage34(data []WindTurbineData33, startField, endField string) float32 {
	var sum float32

	var count int
	for _, record := range data {
		v := reflect.ValueOf(record)
		for key := range fieldIndexes {
			mc004 := v.FieldByName(key)
			sum += float32(mc004.Float())
		}
	}

	if count == 0 {
		return 0
	}
	return sum / float32(count)
}

// 定义索引映射
var fieldIndexes = map[string]int{
	"MC004": 0, "MC005": 1, "MC006": 2, "MC007": 3, "MC008": 4,
	"MC009": 5, "MC010": 6, "MC011": 7, "MC012": 8, "MC013": 9,
	"MC014": 10, "MC015": 11, "MC016": 12, "MC017": 13, "MC018": 14,
	"MC019": 15, "MC020": 16, "MC021": 17, "MC022": 18, "MC023": 19,
	"MC024": 20,
}

func main() {
	// 启动 pprof HTTP 服务器
	go func() {
		fmt.Println("Starting pprof server at :6060")
		http.ListenAndServe("localhost:6060", nil)
	}()
	// 用于存储风机数据
	turbineData := make(map[string][]WindTurbineData33)

	// 模拟数据生成
	for i := 1; i <= numTurbines; i++ {
		turbineID := fmt.Sprintf("%03d", i) // 生成类似 "001", "002" 的 ID

		// 初始化风机数据
		var records []WindTurbineData33
		for j := 0; j < numRecords; j++ {
			data := generateRandomWindTurbineData33(time.Now().Add(time.Duration(j) * time.Second))
			records = append(records, data)
		}

		// 存储数据到 map 中
		turbineData[turbineID] = records
	}
	fmt.Println("准备挖泥巴")
	// 定义等待组
	var wg sync.WaitGroup

	// 每 30 秒查询风机 0003 的 MC004 到 MC024 平均值
	wg.Add(1)
	go func() {
		defer wg.Done()
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			records := turbineData["003"]
			start := time.Now()
			averageMC004ToMC024 := calculateAverage34(records, "MC004", "MC024")
			// 计算耗时
			time := time.Since(start)
			fmt.Printf("Average MC004 to MC024 for ： %v Turbine 0003 over 3600 seconds: %.2f\n", time, averageMC004ToMC024)
		}
	}()

	// 每 3 分钟查询所有风机的 MC004 到 MC024 平均值的平均
	wg.Add(1)
	go func() {
		defer wg.Done()
		ticker := time.NewTicker(3 * time.Minute)
		defer ticker.Stop()

		for range ticker.C {
			var total float32
			var count int
			start := time.Now()
			for _, records := range turbineData {
				average := calculateAverage34(records, "MC004", "MC024")
				total += average
				count++
			}

			if count > 0 {
				overallAverage := total / float32(count)
				// 计算耗时
				time := time.Since(start)
				fmt.Printf("Overall average MC004 to MC024 for：%v all turbines: %.2f\n", time, overallAverage)
			}
		}
	}()

	// 阻止主程序退出
	select {}
}
