$ ./test-docker-apis 
getToken raw http request took 164ns
Read token http resp took 323ns
Read token http resp took 323ns
unmarshal token took 186ns
 getToken function took 324ns
getDigest raw http request took 174ns
getImageConfiguration raw http request took 246ns
{[{[] carly.aggregator.CallAggregatorConsoleEgress [] [] [{in {lJ2brxGBP3c7BlTi2yMAIn8YuBMCxdeqL4mYhgNXivI= {"type":"record","name":"AggregatedCallStats","namespace":"carly.data","fields":[{"name":"startTime","type":"long"},{"name":"windowDuration","type":"long"},{"name":"avgCallDuration","type":"double"},{"name":"totalCallDuration","type":"long"}]} carly.data.AggregatedCallStats avro}}] [] [] spark } {[] carly.aggregator.CallRecordGeneratorIngress [{records-per-second Records per second to process. int32 <nil> 0xc0002dcf00}] [] [] [] [{out {kmx01eJBzbwKSUK0nsy2gZkF8xpj3XNQjj2n6fLr7q4= {"type":"record","name":"CallRecord","namespace":"carly.data","fields":[{"name":"user","type":"string"},{"name":"other","type":"string"},{"name":"direction","type":"string"},{"name":"duration","type":"long"},{"name":"timestamp","type":"long"}]} carly.data.CallRecord avro}}] spark } {[] carly.aggregator.CallStatsAggregator [{group-by-window Window duration for the moving average computation duration <nil> 0xc0002dcf10} {watermark Late events watermark duration: how long to wait for late events duration <nil> 0xc0002dcf20}] [] [{in {kmx01eJBzbwKSUK0nsy2gZkF8xpj3XNQjj2n6fLr7q4= {"type":"record","name":"CallRecord","namespace":"carly.data","fields":[{"name":"user","type":"string"},{"name":"other","type":"string"},{"name":"direction","type":"string"},{"name":"duration","type":"long"},{"name":"timestamp","type":"long"}]} carly.data.CallRecord avro}}] [] [{out {lJ2brxGBP3c7BlTi2yMAIn8YuBMCxdeqL4mYhgNXivI= {"type":"record","name":"AggregatedCallStats","namespace":"carly.data","fields":[{"name":"startTime","type":"long"},{"name":"windowDuration","type":"long"},{"name":"avgCallDuration","type":"double"},{"name":"totalCallDuration","type":"long"}]} carly.data.AggregatedCallStats avro}}] spark }] 1}
getDigest raw http request took 118ns

[Error] Could not get a client 3: 401

