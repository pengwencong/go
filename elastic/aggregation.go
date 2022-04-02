package elastic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"go/server"
)

func Aggregation(c *gin.Context)  {
	//avg()
	//cardinality()
	redis := server.GetRedis()
	res, err := redis.Get("peng").Result()
	if err != nil {

	}
	fmt.Println(res)
	server.PutRedis(redis)
}

/*
sum,min,max,
*/
func avg(){
	//InitData()
	avg := elastic.NewAvgAggregation()
	avg.Field("viewNum").
		Missing("1")

	searchAggregation("avg", avg)

}

/*
聚合查询存在误差，在5%范围之内，通过调整“precision_threshold”参数进行调整默认3000
数字
*/
func cardinality(){
	card := elastic.NewCardinalityAggregation()
	card.Field("viewNum")

	searchAggregation("card",card)
}

/*
统计geo边界
geoCentroid,中心点

*/
func GeoBound(){
	geoBound := elastic.NewGeoBoundsAggregation()
	geoBound.Field("location")
	elastic.NewPercentileRanksAggregation()

	searchAggregation("geoBound", geoBound)
}

/*
百分位概念：例：小明考了50分，整个班级中有70%的人少于50分，则70的百分位是50
Percentiles百分位统计，显示哪个百分位的值向上取
Percentiles rank 显示值以内占了百分几

stat:一起显示min,max,avg,sum

value_count:遍历每一个文档的这个field，如果这个field是普通值，则count值加1，如果是array类型的，则count值加n
*/