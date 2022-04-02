package elastic

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"go/help"
	"go/server"
)

type su struct {
	Input []string
}

type TestIndex struct {
	Id int `json:"id"`
	Author string `json:"author"`
	Title string `json:"title"`
	Content string `json:"content"`
	Suggest []string `json:"suggest"`
	Location []float32 `json:"location"`
	ViewNum int `json:"viewNum"`
	CommentNum int `json:"commentNum"`
	Weight int `json:"weight"`
	Score float64 `json:"score"`
	CreateTime int `json:"create_time"`
}
type TestIndex1 struct {
	Id int `json:"id"`
	Author string `json:"author"`
	Title string `json:"title"`
	Content string `json:"content"`
	ViewNum int `json:"viewNum"`
	CommentNum int `json:"commentNum"`
	Weight int `json:"weight"`
	Score float64 `json:"score"`
	CreateTime int `json:"create_time"`
}
//fielddata 用于排序、聚合场景，text默认禁用7.16去除
//keyword不能用分词
/*
translog操作日志用于断点回复
flush_threshold_size，超过大小则设置新的提交点
sync_interval刷新落盘间隔
durability 异步、同步：立马刷新，异步根据间隔刷新
retention.size 保留到内存的大小目的加快回复数据速度
retention.age 保留时长（同上）
*/
/*
段合并API：_optimize?max_num_segments=1
主要用于更新不活跃的索引，活跃的会自动根据策略合并，有助于提高搜索效率
setting:
*/

/*
index shard:
未分片节点延迟处理时间
index.unassigned.node_left.delayed_timeout:
分片规则
index.routing.allocation.include.{size,name,host,ip,host_ip,id,tier,publish_ip}
未分片恢复优先级,优先级配置、时间、名称（高到低）
index.priority:
每个节点的分片数：
index.routing.allocation.total_shards_per_node
*/
/*
index.soft_deletes.enabled：true 软删除默认开启
index.soft_deletes.retention_lease.period：12h默认12个小时
*/
/*
设置维度的排序
index.sort.field
index.sort.order
index.sort.mode:min/max
index.sort.missing:_last/_first
*/

/*
索引请求时需要的内存限制，如超过将拒绝请求。因为除了处理索引还有其他需要用到内存，一般设置10%
indexing_pressure.memory.limit:
*/
var index = "test"
var index1 = "test1"
var testMaping = `{
	"settings":{
		"refresh_interval": "20s",
		"number_of_shards": 2,
		"translog":{
			"flush_threshold_size": "500mb",
			"sync_interval": "10s",
			"durability": "async"
		},
		"number_of_replicas": 1
	},
	"mappings":{
		"properties":{
			"id":{"type":"integer"},
			"author":{"type":"keyword"},
			"title":{"type":"text","analyzer":"ik_max_word","search_analyzer": "ik_smart"},
			"content":{"type":"text","analyzer":"ik_max_word"},
			"location":{"type":"geo_point"},
			"suggest":{"type":"completion","analyzer":"ik_max_word","search_analyzer":"ik_max_word"},
			"commentNum":{"type":"integer"},
			"Weight":{"type":"integer"},
			"viewNum":{"type":"integer"},
			"create_time":{"type":"integer"}
		}
	}
}`


func CreateIndex(c *gin.Context) {

	clientES := server.GetEs()
	clientES.DeleteIndex(index).Do(context.TODO())
	createIndex, err := clientES.CreateIndex(index).//返回IndicesCreateService
		Body(testMaping).
		Do(context.TODO())
	if err != nil {
		help.Log.Infof("es create index error: %s", err.Error())
	}
	if !createIndex.Acknowledged {
		// Not acknowledged
	}

	slowLogSettings := `{
		"index.search.slowlog.threshold.query.warn": "2s",
	  	"index.search.slowlog.threshold.query.info": "1s",
	  	"index.search.slowlog.threshold.query.debug": "5s",
	  	"index.search.slowlog.threshold.query.trace": "500ms",
	  	"index.search.slowlog.threshold.fetch.warn": "1s",
	  	"index.search.slowlog.threshold.fetch.info": "800ms",
	  	"index.search.slowlog.threshold.fetch.debug": "500ms",
	  	"index.search.slowlog.threshold.fetch.trace": "200ms",
	  	"index.search.slowlog.level": "info"
	}`

	pusettingResponse, err := clientES.IndexPutSettings("test1").
		BodyString(slowLogSettings).
		Do(context.TODO())
	if err != nil {
		help.Log.Infof("es create index error: %s", err.Error())
	}
	if !pusettingResponse.Acknowledged {
		// Not acknowledged
	}

	server.PutEs(clientES)
}

/*
需要先创建索引，mapping对不上的会忽略，对的上的字段会赋予数据
*/
func Reindex(){
	clientES := server.GetEs()

	reServer := elastic.NewReindexService(clientES)
	res, err := reServer.
		SourceIndex(index).
		DestinationIndex("test1").
		Do(context.TODO())
	if err != nil {

	}

	fmt.Println(res)

	server.PutEs(clientES)
}

/*
数据迁移（reindex）：重新设定mapping、分片
提速：批量，scroll（并行）,slices（分片）
*/

/*
管道在于数据通道存在处理过程
有：
IngestPutPipelineService 存储pipeline
IngestGetPipelineService 获取pipeline
IngestDeletePipelineService 删除pipeline
IngestSimulatePipelineService 模拟pipeline？
*/

/*
处理过程：（没自定义）
默认：uppdercase,lowercase,set(设置一个字段值),
bulkProcessor:像是一个不断接收数据-根据设置提交请求的协程
setBulkActions(1000):每添加1000个request，执行一次bulk操作
setBulkSize(new ByteSizeValue(5, ByteSizeUnit.MB)):每达到5M的请求size时，执行一次bulk操作
setFlushInterval(TimeValue.timeValueSeconds(5)):每5s执行一次bulk操作
*/

/*
建仓库，再建快照。建完快照就已经保存好，可以用restore重新保存
获取可以忽略快照不存在，创建则不行
集群快照需要挂载本系统卷否则会没权限
*/
func SnapShot() {
	clientES := server.GetEs()

	//创建仓库
	service := clientES.SnapshotCreateRepository("test")
	_, err := service.
		Setting("location", "/usr/share/elasticsearch/backups/my_backup").
		Type("fs").
		Do(context.TODO())
	if err != nil {
		fmt.Printf("create  repository err : %s\n", err.Error())
		return
	}

	//创建快照
	snapShotCreate := clientES.SnapshotCreate("test","1second")
	snapShot, err := snapShotCreate.
		Human(true).
		Pretty(true).
		Do(context.TODO())
	if err != nil {
		fmt.Printf("snapShot create err: %s\n", err.Error())
		return
	}

	fmt.Printf("snapShot info ss: %v\n", snapShot)

	//获取快照
	getService := clientES.SnapshotGet("test")
	getResponse, err := getService.
		Snapshot("second").
		IgnoreUnavailable(true).
		Human(true).
		Pretty(true).
		Do(context.TODO())
	if err != nil {
		fmt.Printf("snapshot get err: %s\n", err.Error())
	}

	for _, val := range getResponse.Snapshots {
		fmt.Printf("snapshot data: %+v\n", *val)
	}

	verifyService := elastic.NewSnapshotVerifyRepositoryService(clientES)
	verifyResponse, err := verifyService.Repository("test").
		Human(true).
		Pretty(true).
		Do(context.TODO())
	if err != nil {
		return
	}

	for _, node := range verifyResponse.Nodes {
		fmt.Println(node.Name)
	}

	server.PutEs(clientES)
}

/*
批量操作的目的都是节省网络开销
*/
func Bulk() {
	clientES := server.GetEs()
	buldServer := clientES.Bulk()

	data := TestIndex{
		Id: 4,
		Author:"彭文聪",
		Title:"内容推荐",
		Content:"内容推荐",
		ViewNum: 1,
		CommentNum: 3,
		Weight: 4,
		CreateTime:1647048371,
	}

	buldServer.Add(
		elastic.NewBulkIndexRequest().
			Index(index).
			Doc(data).
			Id("2") )

	_, err := buldServer.
		Human(true).
		Pretty(true).
		Do(context.TODO())
	if err != nil {
		help.Log.Infof("es add document error: %s", err.Error())
	}

	server.PutEs(clientES)
}

func InitData() {
	clientES := server.GetEs()
	//id := 4
	//stpool := []rune("额我就怕目前问题我去过了撒刚放假叫哦我该金融个日哦交割日估计清热加工我就是大V发这个花多少【啊发生了框架搭建了刚和日韩剧发锯齿形，没咋说打开了刚放完法王噶十多个撒口感")
	ss := []string{"兴东 中粮创芯","中粮创芯"}
	lo := []float32{113.91,22.57}
	//ssu := su{
	//	Input:ss,
	//}
	data := TestIndex{
		Id: 1,
		Author:"彭文聪",
		Title:"中粮创芯",
		Content:"中粮创芯",
		Suggest: ss,
		Location:lo,
		ViewNum: 4,
		CommentNum: 4,
		Weight: 1,
		CreateTime:1647048371,
	}
	//data1 := TestIndex{
	//	Id: 1,
	//	Author:"彭文聪",
	//	Title:"内容推荐",
	//	Content:"内容推荐",
	//
	//	ViewNum: 1,
	//	CommentNum: 4,
	//	Weight: 1,
	//	CreateTime:1647048371,
	//}
	_, err := clientES.Index().
		Index(index).
		Id("24").
		BodyJson(data).
		Do(context.TODO())
	if err != nil {
		fmt.Printf("es add 1 %s\n", err.Error())
	}

	//bulk
	//buldServer := clientES.Bulk()
	//ll := len(stpool)
	//
	//yuan := []rune("内容推荐")
	//for i:= 0; i < 3; i++ {
	//	textNum := rand.Intn(40)
	//	var ssst = make([]rune, textNum+40)
	//	e := rand.Intn(textNum+36)
	//	for j := 0; j < textNum+36; j++ {
	//		if j == e {
	//			for _, val := range yuan {
	//				ssst[j] = val
	//				j++
	//			}
	//		}
	//		sst := stpool[rand.Intn(ll)]
	//		ssst[j] = sst
	//	}
	//	data.Id = id
	//	data.Weight = id
	//	data.CommentNum = id - 2
	//	data.Content = string(ssst)
	//	idStr := strconv.Itoa(id)
	//	buldServer.Add(
	//		elastic.NewBulkIndexRequest().
	//		Index(index).
	//		Doc(data).
	//		Id(idStr) )
	//	id++
	//}
	//
	//bulkResponse, err := buldServer.
	//	Human(true).
	//	Pretty(true).
	//	Do(context.TODO())
	//if err != nil {
	//	help.Log.Infof("es add document error: %s", err.Error())
	//}
	//
	//for _, val := range bulkResponse.Succeeded() {
	//	fmt.Println(val.Id)
	//}

	server.PutEs(clientES)
}
