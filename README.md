# kafkapipe
kafkapipe pipes msgs between two kafka server(s) via librdkafka.

# dependency
librdkafka [https://github.com/edenhill/librdkafka]

# how to use
`kafkapipe --from from.toml --to to.toml`

# how to setup
see configure samples in examples dir, for more options --> [https://github.com/edenhill/librdkafka/blob/master/CONFIGURATION.md]

### from.toml
<pre>
[consumerConfig]
"bootstrap.servers" = "10.2.2.2:9092"
"group.id" = "xx"
"session.timeout.ms" = "30000"
"fetch.message.max.bytes" = "655350"
"go.events.channel.enable" = true
"go.application.rebalance.enable" = true
"auto.offset.reset" = "earliest"
"heartbeat.interval.ms" = "5000"
"queued.max.messages.kbytes" = "655350"
"queued.min.messages" = "1000"

[topicFrom]
name = "ccc"
</pre>

### to.toml
<pre>
[producerConfig]
"bootstrap.servers" = "10.2.2.2:9092"

[topicTo]
name = "aaa"
</pre>
