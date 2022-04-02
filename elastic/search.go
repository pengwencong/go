package elastic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"go/help"
	"go/server"
	"log"
)

/*
doc1-title测试精确查询
doc2-title测精确查询
result: doc1
*/
func Term() {
	termOne := elastic.NewTermQuery("title","测试")

	search(termOne)
}

func TestSearch(c *gin.Context) {
	InitData()
	//MoreLikeThis()
	//SnapShot()
	//Span()
	//Suggester()
	//FuncScore()
	//Reindex()
	//GeoDis()
}

func GeoDis(){
	matchAll := elastic.NewMatchAllQuery()

	geoDis := elastic.NewGeoDistanceQuery("xing dong")
	geoDis.Distance("2km").
		Point(float64(22.58),float64(113.91))

	bool := elastic.NewBoolQuery()
	bool.Must(matchAll).
		Filter(geoDis)

	search(bool)
}

/*
区间查询只有first,near,都只能用spanterm的精确查询
first的end只在文本的查询结束位置，以中文分词为准
near跨区间查询，现在两个查询间的距离以及是否按顺序
*/
func Span(){
	match := elastic.NewSpanTermQuery("title","精确")
	spanF := elastic.NewSpanFirstQuery(match, 2)

	search(spanF)

	nearTerm1 := elastic.NewSpanTermQuery("title","测")
	nearTerm2 := elastic.NewSpanTermQuery("title", "查询")
	spanNear := elastic.NewSpanNearQuery(nearTerm2).
		Add(nearTerm1).
		Slop(2).
		InOrder(true)

	search(spanNear)
}

/*
需要在mapping那里的type类型为completion仅限于completeSuggester，term不需要
mode:missing,仅当查询的词语错误才能查出如：内容精确，无论查什么都查不出，因为容字错了，会被认为是内和错的容，所以要三个字的词组。
如：三生石，即便是此也是仅生字错了，或者只输入三生才会查出，输入声石查不出
*/
func Suggester(){
	match := elastic.NewMatchQuery("title", "测")
	suggester := elastic.NewTermSuggester("suggest").
		Field("title").
		Text("测试").
		SuggestMode("popular")

	searchSuggest(match, suggester)



}

/*
这个包的API没有切合官方api，具体看官方文档
*/
func FuncScore(){


}


/*
author 找不到
title:doc1,doc2,doc3
中文分词会把“精确查询”分成“精确”“查询”
同时文档也会根据中文分词的规则进行查询
如查找title:“测 试”则只能找到doc2
slop只作用于matchphrase,matchphraseprefix,spannear,
而且作用为：短语与短语之间的间隔数：这个数指的是分词后间隔的位数如下：能搜doc1,doc4
CutoffFrequency为词频，关键在于分词后的词语数量占比，可只判断以为小数点。只在CommonTermsQuery，match,multimatch
*/
func Match() {
	matchAuthor := elastic.NewMatchQuery("author","彭")
	matchTitle := elastic.NewMatchQuery("title","精确查询")
	matchSlopTitle := elastic.NewMatchPhraseQuery("title","测试查询")
	matchSlopTitle.Slop(1)
	matchFuzzyTitle := elastic.NewMatchQuery("title","精确查询")
	matchFuzzyTitle.CutoffFrequency(0.3)
	request0 := elastic.NewSearchRequest().Query(matchAuthor)
	request1 := elastic.NewSearchRequest().Query(matchTitle)
	request2 := elastic.NewSearchRequest().Query(matchSlopTitle)
	request3 := elastic.NewSearchRequest().Query(matchFuzzyTitle)

	multiSearch(request0, request1, request2, request3)
	//search(matchAuthor)
	//search(matchTitle)
	//search(matchSlopTitle)
	//search(matchFuzzyTitle)
}
/*
multimatch是针对一个字段搜索多个可能的值
每个field可以设置boost,其本身也可设置boost
同时还有MinimumShouldMatch，表示最少满足多少有多少字段的文档
*/

/*
MultiSearchService
*/

/*
fuzzy对于不分词类型也有作用
相对于max_expansions fuzzy也可作用于中间
*/
func Fuzzy() {
	Match := elastic.NewFuzzyQuery("author", "彭聪")
	Match.Fuzziness("1").
		PrefixLength(0)

	search(Match)
}

/*
满足positive才显示
满足negative的score跟下面的数组相乘如果小于1
相当于满足negative的权重更低
*/
func Boosting(){
	positiveQuery := elastic.NewMatchQuery("title", "精确")
	negativeQuery := elastic.NewMatchQuery("title", "测试")
	boostingQuery := elastic.NewBoostingQuery().
		Positive(positiveQuery).
		Negative(negativeQuery).
		NegativeBoost(0.5)

	search(boostingQuery)
}

/*
其是根据映射到矩阵向量空间再根据配置查找，同时相似内容少于5个也搜不出来
当只有一个搜索文本并多个文档有且仅有一个匹配项评分标准
暂只知道内容占比例高评分高(包括总体内容长度更短，搜索内容频率越高)但每次都一样
maxwordlenth:指的是需要搜的词的长度
MinDocFreq:指的是匹配的文档在整个索引文档数占得比例
MaxQueryTerms:应该是查询参数次数，然后根据每次的文档内容进行评分
MinimumShouldMatch:当有多个text，或者id的时候，文章最少应该匹配几个
*/
func MoreLikeThis() {
	molik := elastic.NewMoreLikeThisQuery()

	molik.Field("content").
		LikeText("内容推荐").
		MinTermFreq(1)
	search(molik)
}

func searchAggregation(name string, ag elastic.Aggregation) {
	clientEs := server.GetEs()

	res, err := clientEs.Search(index).
		Aggregation(name, ag).
		Pretty(true).
		Do(context.TODO())
	if err != nil {
		help.Log.Infof("search aggregation error: %s", err.Error())
		return
	}

	agg, found := res.Aggregations.Terms(name)
	if !found {
		log.Fatalf("we should have a terms aggregation called %q", "timeline")
	}
	for _, val := range agg.Aggregations {
		fmt.Println(string(val))
	}
	server.PutEs(clientEs)
}

/*
默认10个
*/
func search(query elastic.Query) {
	clientEs := server.GetEs()

	res, err := clientEs.Search(index).
		Query(query).
		FetchSource(true).
		Pretty(true).
		Size(20).
		Do(context.TODO())
	if err != nil {
		help.Log.Infof("search term error: %s", err.Error())
		return
	}

	getRes(res)
	server.PutEs(clientEs)
}

/*
默认10个
*/
func searchSuggest(query elastic.Query, suggest *elastic.TermSuggester) {
	clientEs := server.GetEs()

	res, err := clientEs.Search(index).
		//Query(query).
		Suggester(suggest).
		Pretty(true).
		Size(20).
		Do(context.TODO())
	if err != nil {
		help.Log.Infof("search term error: %s", err.Error())
		return
	}

	getRes(res)
	server.PutEs(clientEs)
}

func getRes(res *elastic.SearchResult){
	if res.Hits.TotalHits.Value > 0 {

		for _, hit := range res.Hits.Hits {
			// hit.Index contains the name of the index

			// Deserialize hit.Source into a Tweet (could also be just a map[string]interface{}).
			var testIndex TestIndex
			err := json.Unmarshal(hit.Source, &testIndex)
			if err != nil {
			}
			testIndex.Score = *hit.Score
			fmt.Printf("%+v\n", testIndex)
		}
	}
	fmt.Println("fasd")
}

func multiSearch(requests ...*elastic.SearchRequest){
	clientEs := server.GetEs()
	multiSearch := clientEs.MultiSearch()
	for _, request := range requests {
		multiSearch.Add(request)
	}

	res, err := multiSearch.
		Pretty(true).
		Human(true).
		Do(context.TODO())

	if err != nil {
		help.Log.Infof("multi search term error: %s", err.Error())
	}

	if len(res.Responses) > 0 {
		for _, res := range res.Responses {
			if res.Hits.TotalHits.Value > 0 {

				for _, hit := range res.Hits.Hits {
					// hit.Index contains the name of the index

					// Deserialize hit.Source into a Tweet (could also be just a map[string]interface{}).
					var testIndex TestIndex
					err := json.Unmarshal(hit.Source, &testIndex)
					if err != nil {
					}
					testIndex.Score = *hit.Score
					fmt.Printf("%+v\n", testIndex)
				}
			}
		}
	}
}