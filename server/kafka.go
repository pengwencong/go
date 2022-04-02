package server

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
	"go/help"
	"math/rand"
	"strconv"
	"sync"
	"time"
)
var kafkaAddr = []string{"106.55.178.129:9092"}

var kafkaConfig *sarama.Config

var kafkaPool = sync.Pool{
	New :func() interface{} {
		ins, _ := newKafka()
		return ins
	},
}

func GetKafka() *sarama.Client {
	return kafkaPool.Get().(*sarama.Client)
}

func PutKafka(instance *sarama.Client) {
	kafkaPool.Put(instance)
}

func InitKafkaConfig() {
	kafkaConfig = sarama.NewConfig()
	kafkaConfig.Version = sarama.MaxVersion

	/*producter:
	MaxMessageBytes,Timeout,CompressionLevel,patitioner(自定义分区器，需要有key),return.success/error,
	Idempotent(幂等，false,broker故障会在重试时写入多次消息),flush（刷新到log的时间，大小，数量间隔配置），
	Retry，Interceptors
	*/
	/*consumer:
	Retry,Group.session(time.Duration)/heartbeat(time.Duration)/(重新分配组内分区给组内消费者)Rebalance.Timeout/retry/
	Strategy（自定义）/Group.Member(加入组时设置的metadata数据)
	maxwaittime/MaxProcessingTime/Return.error/Interceptors/IsolationLevel/
	Fetch.min/default/max(读取数据通过fetch，读取数据限制)
	Offsets.Retention（检查失效offset频率）/retry/AutoCommit.enable/interval(自动提交消费者的offset)/Initial(从头或从尾消费)

	*/
	/*net:

	*/

	/*metadata:

	*/

	//0(发送不等确认),1（发送写入磁盘再确认）,-1（发送写入磁盘且副本同步再确认）
	kafkaConfig.Producer.RequiredAcks = sarama.WaitForAll
	//0(不压缩)，1（zip),2(Snappy),3(LZ4),4(ZSTD)
	//kafkaConfig.Producer.Compression = 1
	//kafkaConfig.Producer.Retry.Max = 3
	//kafkaConfig.Producer.Retry.Backoff = time.Millisecond * 500
	//kafkaConfig.Producer.Return.Successes = true
	//kafkaConfig.Producer.Return.Errors = true

	//kafkaConfig.ChannelBufferSize = 3
	kafkaConfig.ClientID = "kafka"

	//kafkaConfig.Consumer.Retry.Backoff = time.Millisecond * 500
	kafkaConfig.Consumer.Offsets.AutoCommit.Enable = false
	kafkaConfig.Consumer.IsolationLevel = sarama.ReadCommitted
}

func InitKafkaPool(num int) error {
	for i:= 0; i < num; i++ {
		kafkaInstance, err := newKafka()
		if err != nil {
			return err
		}
		kafkaPool.Put(kafkaInstance)
	}

	return nil
}

func newKafka() (*sarama.Client, error) {
	kafkainstance, err := sarama.NewClient(kafkaAddr, kafkaConfig)
	if err != nil {
		help.Log.Errorf("new kafka client error: %s", err.Error())
	}

	return &kafkainstance, err
}

func Acl(){
	acl := sarama.Acl{
		Principal:"peng",
		Host:"106.55.178.129",
		Operation:sarama.AclOperationAll,
		PermissionType:sarama.AclPermissionAllow,
	}

	resource := sarama.Resource{
		ResourceType: sarama.AclResourceTopic,
		ResourceName: "my-topic",
	}

	aclCreation := sarama.AclCreation{
		resource,
		acl,
	}

	slice := []*sarama.AclCreation{&aclCreation}
	createRequest := sarama.CreateAclsRequest{
		AclCreations: slice,
	}

	broker := sarama.NewBroker("106.55.178.129:9092")
	err := broker.Open(nil)
	if err != nil {
		help.Log.Errorf("broker open error: %s\n", err.Error())
	}

	response, err := broker.CreateAcls(&createRequest)
	if err != nil {
		help.Log.Errorf("CreateAcls error: %s\n", err.Error())
	}

	fmt.Printf("CreateAcls %+v", response)
}

func Txn(){
	//cli := GetKafka()


	//r := kafka.NewReader(kafka.ReaderConfig{
	//	Brokers:   []string{"106.55.178.129:9092"},
	//	Topic:     "kafka-go-4abb18c948b9e9",
	//	Partition: 0,
	//	MinBytes:  10e3, // 10KB
	//	MaxBytes:  10e6, // 10MB
	//	IsolationLevel: kafka.ReadCommitted,
	//})
	//
	//for {
	//	m, err := r.ReadMessage(context.Background())
	//	if err != nil {
	//		break
	//	}
	//	fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	//}
	//
	//if err := r.Close(); err != nil {
	//	fmt.Println("failed to close reader:", err)
	//}

	//addr := "106.55.178.129:9092"
	//
	//broker := sarama.NewBroker(addr)
	//err := broker.Open(kafkaConfig)
	//if err != nil {
	//	help.Log.Errorf("broker open error: %s\n", err.Error())
	//}
	//topic := "kafka-go-4abb18c948b9e9"
	//
	//fr := &sarama.FetchRequest{
	//	Isolation: sarama.ReadUncommitted,
	//}
	//fr.AddBlock(topic,0,3,100)
	//frre, err := broker.Fetch(fr)
	//if err != nil {
	//
	//}
	//
	//fmt.Printf("fr: %+v\n", frre)
	//for _, Block := range frre.Blocks{
	//	for _, FetchResponseBlocks := range Block {
	//		for _, Records := range FetchResponseBlocks.RecordsSet {
	//			fmt.Println(Records)
	//		}
	//	}
	//}
	//a := make(map[string]*sarama.TopicDetail, 1)
	//a[topic] = &sarama.TopicDetail{
	//	NumPartitions: 3,
	//	ReplicationFactor: 1,
	//}
	//
	//creTR := &sarama.CreateTopicsRequest{
	//	TopicDetails : a,
	//}
	//
	//resp, err := broker.CreateTopics(creTR)
	//if err != nil {
	//	help.Log.Errorf("broker open error: %s\n", err.Error())
	//}
	//fmt.Printf("s %+v\n", resp)
	//resp, err := broker.ListGroups(&sarama.ListGroupsRequest{})
	//fmt.Println(resp.Groups)
	//err = broker.Open(nil)
	//if err != nil {
	//	help.Log.Errorf("broker open error: %s\n", err.Error())
	//}

	//transactionalId := fmt.Sprintf("kafka-go-transactional-id-%016x", rand.Int63())
	//
	//coordRequest := &sarama.FindCoordinatorRequest{
	//	CoordinatorKey  : transactionalId,
	//	CoordinatorType : sarama.CoordinatorTransaction,
	//}
	//
	//findCoordinaRes, err := broker.FindCoordinator(coordRequest)
	//if err != nil {
	//	help.Log.Errorf("find Coordinator error: %s\n", err.Error())
	//}
	//
	//fmt.Println(findCoordinaRes.Err)
	//fmt.Println(findCoordinaRes.ErrMsg)
	//
	//req := &sarama.InitProducerIDRequest{
	//	TransactionalID: &transactionalId,
	//	TransactionTimeout: time.Second * 10,
	//}
	//
	//pro, err := broker.InitProducerID(req)
	//if err != nil {
	//	help.Log.Errorf("broker init producter id error: %s\n", err.Error())
	//}
	//
	//fmt.Println(pro.Err)
	//
	//topPar := make(map[string][]int32, 10)
	//topPar[topic] = []int32{0,1,2}
	//ParTxn := &sarama.AddPartitionsToTxnRequest {
	//	TransactionalID:transactionalId,
	//	ProducerID: pro.ProducerID,
	//	ProducerEpoch: pro.ProducerEpoch,
	//	TopicPartitions: topPar,
	//}
	//
	//res, err := broker.AddPartitionsToTxn(ParTxn)
	//if err != nil {
	//	help.Log.Errorf("broker add offset txn error: %s\n", err.Error())
	//}
	//
	//for _, val := range res.Errors {
	//	for _, v := range val {
	//		fmt.Println(v.Err)
	//	}
	//}


	//pr := &sarama.ProduceRequest{
	//	TransactionalID: &transactionalId,
	//	RequiredAcks: sarama.WaitForAll,
	//}
	//
	//msg := &sarama.Message{
	//	Value: []byte("test txn"),
	//}
	//
	//pr.AddMessage("my_topic", 0, msg)
	//
	//broker.Produce(pr)
	//PutKafka(cli)
}


func Producer(c *gin.Context){
	client := GetKafka()

	producer, err := sarama.NewAsyncProducerFromClient(*client)
	if err != nil {
		help.Log.Errorf("new producter error: %s\n", err.Error())
	}
	defer producer.Close()

	tick := time.Tick(time.Second * 1)
	time := 2
	for {
		select {
		case <- tick:
			producer.Input() <- &sarama.ProducerMessage{
				Topic: "kafka-go-4abb18c948b9e9",
				Key: nil,
				Value: sarama.StringEncoder(strconv.Itoa(time)),
			}
			time++
		}
		select {
		case err := <-producer.Errors():
			help.Log.Errorf("produce send message error: %s\n", err.Error())
		case res := <-producer.Successes():
			help.Log.Infof("produce success message: %+v\n", res)
		}
	}

	PutKafka(client)
}

func Tex(){
	topic1 := "kafka-go-4abb18c948b9e"
	//topic2 := fmt.Sprintf("kafka-go-%016x", rand.Int63())

	client := &kafka.Client{
		Addr:      kafka.TCP("106.55.178.129:9092"),
		Timeout:   10 * time.Second,
	}

	transactionalID := fmt.Sprintf("kafka-go-transactional-id-%016x", rand.Int63())
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	ipResp, err := client.InitProducerID(ctx, &kafka.InitProducerIDRequest{
		TransactionalID:      transactionalID,
		TransactionTimeoutMs: 10000,
	})
	if err != nil {

	}

	if ipResp.Error != nil {

	}
	fmt.Printf("a%+v\n", ipResp)

	fmt.Println(ipResp.Error)

	//resp, err := client.AddPartitionsToTxn(ctx, &kafka.AddPartitionsToTxnRequest{
	//	TransactionalID: transactionalID,
	//	ProducerID:      ipResp.Producer.ProducerID,
	//	ProducerEpoch:   ipResp.Producer.ProducerEpoch,
	//	Topics: map[string][]kafka.AddPartitionToTxn{
	//		topic1: {
	//			{
	//				Partition: 0,
	//			},
	//		},
	//	},
	//})
	//if err != nil {
	//}
	//fmt.Printf("%+v\n", resp)

	records := make([]kafka.Record, 0, 10)
	for i := 0; i < 10; i++ {
		records = append(records, kafka.Record{
			Time:  time.Now(),
			Value: kafka.NewBytes([]byte("test-message-" + strconv.Itoa(i))),
		})
	}
	ctx, cancel = context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	res, err := client.Produce(ctx, &kafka.ProduceRequest{
		Topic:        topic1,
		RequiredAcks: kafka.RequireAll,
		Records:      kafka.NewRecordReader(records...),
		TransactionalID: transactionalID,
	})
	if err != nil {
	}
	fmt.Printf("res: %+v\n", res)


	resp1, err := client.EndTxn(ctx, &kafka.EndTxnRequest{
		TransactionalID: transactionalID,
		ProducerID:      ipResp.Producer.ProducerID,
		ProducerEpoch:   ipResp.Producer.ProducerEpoch,
		Committed:       false,
	})
	if err != nil {

	}
	fmt.Printf("resp1: %+v\n", resp1)
	//res, err := client.CreateTopics(context.Background(), &kafka.CreateTopicsRequest{
	//	Topics: []kafka.TopicConfig{{
	//		Topic:             topic1,
	//		NumPartitions:     3,
	//		ReplicationFactor: 1,
	//	}},
	//})
	//if err != nil {
	//}
	//fmt.Printf("a%+v\n", res)

	//
	//respc, err := client.FindCoordinator(ctx, &kafka.FindCoordinatorRequest{
	//	Addr:    client.Addr,
	//	Key:     transactionalID,
	//	KeyType: kafka.CoordinatorKeyTypeTransaction,
	//})
	//if err != nil {
	//}
	//
	//fmt.Printf("a%+v\n", respc)
	//
	//transactionCoordinator := kafka.TCP(net.JoinHostPort(respc.Coordinator.Host, strconv.Itoa(int(respc.Coordinator.Port))))
	//client1 := &kafka.Client{
	//	Addr:      transactionCoordinator,
	//	Timeout:   10 * time.Second,
	//}
	//

	//
	//defer func() {
	//	resp, err := client1.EndTxn(ctx, &kafka.EndTxnRequest{
	//		TransactionalID: transactionalID,
	//		ProducerID:      ipResp.Producer.ProducerID,
	//		ProducerEpoch:   ipResp.Producer.ProducerEpoch,
	//		Committed:       false,
	//	})
	//	if err != nil {
	//	}
	//	fmt.Printf("%+v\n", resp)
	//}()

	//
	//for topic, partitions := range resp.Topics {
	//	if topic == topic1 {
	//		if len(partitions) != 3 {
	//			fmt.Errorf("expected 3 partitions in response for topic %s; got: %d", topic, len(partitions))
	//		}
	//	}
	//	//if topic == topic2 {
	//	//	if len(partitions) != 2 {
	//	//		fmt.Errorf("expected 2 partitions in response for topic %s; got: %d", topic, len(partitions))
	//	//	}
	//	//}
	//	for _, partition := range partitions {
	//		if partition.Error != nil {
	//			fmt.Println(partition.Error)
	//		}
	//	}
	//}
}

func Test(c *gin.Context){
	Tex()
}

type consumerGroupHandler struct{}

func (consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h consumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("Message topic: %q partition: %d offset: %d\n value: %s\n", msg.Topic,
			msg.Partition, msg.Offset, msg.Value)
		sess.MarkMessage(msg, "")
	}
	return nil
}

func ConsumerGroup(){
	client := GetKafka()
	consumer, err := sarama.NewConsumerGroupFromClient("my-group1", *client)
	if err != nil {
		help.Log.Errorf("new consumer group error: %s\n", err.Error())
	}
	defer func() {
		if err := consumer.Close(); err != nil {
			help.Log.Errorf("consumer close error: %s\n",err.Error())
		}
	}()
	ctx := context.Background()
	for {
		topics := []string{"kafka-go-4abb18c948b9e"}
		handler := consumerGroupHandler{}

		err := consumer.Consume(ctx, topics, handler)
		if err != nil {
			help.Log.Errorf("group consumer consume error: %s\n", err.Error())
		}
	}

	PutKafka(client)
}

func Consumer(){
	//client := GetKafka()
	consumer, err := sarama.NewConsumer(kafkaAddr, kafkaConfig)
	if err != nil {
		help.Log.Errorf("new consumer error: %s\n", err.Error())
	}
	defer func() {
		if err := consumer.Close(); err != nil {
			help.Log.Errorf("consumer close error: %s\n",err.Error())
		}
	}()

	partitionConsumer, err := consumer.ConsumePartition("kafka-go-4abb18c948b9e", 0, sarama.OffsetOldest)
	if err != nil {
		help.Log.Errorf("consumer partition error: %s\n",err.Error())
	}
	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			help.Log.Errorf("partition consumer close error: %s\n", err.Error())
		}
	}()

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			if msg != nil {
				fmt.Printf("Consumed message offset %d\n, %s\n", msg.Offset, msg.Value)
			}
		}
	}

	fmt.Println("Consumed message close")
	//PutKafka(client)
}