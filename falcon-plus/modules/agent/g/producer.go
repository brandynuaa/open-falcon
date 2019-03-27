package g

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"os"
	"time"
)

//var Address = []string{"140.143.81.243:9092"}
//var topic = "mytest"

//同步消息模式
func SyncProducer(srcValue string, kafkaConfig *KafkaConfig) {
	fmt.Println("===============syncProducer=========================")

	config := sarama.NewConfig()
	//config.Producer.Timeout = 5 * time.Second
	setConfig(*config, kafkaConfig)
	config.Producer.Return.Successes = kafkaConfig.Returnsuccesses
	config.Producer.Return.Errors = kafkaConfig.Returnerrors
	config.Producer.Timeout = time.Duration(kafkaConfig.Timeout) * time.Second

	p, err := sarama.NewSyncProducer(kafkaConfig.Address, config)
	if err != nil {
		log.Printf("sarama.NewSyncProducer err, message=%s \n", err)
		return
	}
	defer p.Close()
	value := fmt.Sprintf(srcValue)
	msg := &sarama.ProducerMessage{
		Topic: kafkaConfig.Topic,
		Value: sarama.ByteEncoder(value),
	}
	part, offset, err := p.SendMessage(msg)

	if err != nil {
		log.Printf("send message(%s) err=%s \n", value, err)
	} else {
		fmt.Fprintf(os.Stdout, value+"发送成功，partition=%d, offset=%d \n", part, offset)
	}
}

func AsyncProducer(srcValue string, kafkaConfig *KafkaConfig) {
	fmt.Println("===============asyncProducer=========================")

	config := sarama.NewConfig()

	setConfig(*config, kafkaConfig)

	//使用配置,新建一个异步生产者
	producer, e := sarama.NewAsyncProducer(kafkaConfig.Address, config)
	if e != nil {
		panic(e)
	}
	defer producer.AsyncClose()

	value := fmt.Sprintf(srcValue)
	//发送的消息,主题,key
	msg := &sarama.ProducerMessage{
		Topic: kafkaConfig.Topic,
		//Key:   sarama.StringEncoder("test"),
		Value: sarama.ByteEncoder(value),
	}
	//使用通道发送
	producer.Input() <- msg

	go func(producer sarama.AsyncProducer) {
		for {
			//循环判断哪个通道发送过来数据.
			select {
			case suc := <-producer.Successes():
				fmt.Println("offset: ", suc.Offset, "timestamp: ", suc.Timestamp.String(), "partitions: ", suc.Partition)
				return
			case fail := <-producer.Errors():
				fmt.Println("err: ", fail.Err)
				return
			}
		}
	}(producer)
}

func setConfig(config sarama.Config, kafkaConfig *KafkaConfig) {

	if kafkaConfig.Partitioner == "NewManualPartitioner" {
		//返回一个手动选择分区的分割器,也就是获取msg中指定的`partition`
		config.Producer.Partitioner = sarama.NewManualPartitioner
	} else if kafkaConfig.Partitioner == "NewRandomPartitioner" {
		//通过随机函数随机获取一个分区号
		config.Producer.Partitioner = sarama.NewRandomPartitioner
	} else if kafkaConfig.Partitioner == "NewRoundRobinPartitioner" {
		//环形选择,也就是在所有分区中循环选择一个
		config.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	} else if kafkaConfig.Partitioner == "NewHashPartitioner" {
		//通过msg中的key生成hash值,选择分区
		config.Producer.Partitioner = sarama.NewHashPartitioner
	} else {
		fmt.Println("no need Partitioner")
	}
	if kafkaConfig.RequiredAcks == 0 {
		config.Producer.RequiredAcks = sarama.NoResponse
	} else if kafkaConfig.RequiredAcks == 1 {
		config.Producer.RequiredAcks = sarama.WaitForLocal
	} else if kafkaConfig.RequiredAcks == -1 {
		//等待服务器所有副本都保存成功后的响应
		config.Producer.RequiredAcks = sarama.WaitForAll
	} else {
		fmt.Println("your param is null,use default param")
		config.Producer.RequiredAcks = sarama.WaitForAll
	}
	//config.Producer.Return.Successes = kafkaConfig.Returnsuccesses
	//config.Producer.Return.Errors = kafkaConfig.Returnerrors
	//config.Producer.Timeout = time.Duration(kafkaConfig.Timeout) * time.Second
}
