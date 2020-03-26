Make sure you install the following libs:
```
sudo apt-get install libdevmapper-dev
sudo apt-get install btrfs-tools
sudo apt-get install libgpgme-dev
```

It looks that registries cannot be unified.
```
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

```

By running skopeoinspect.Test_inspect

```
GOROOT=/home/stavros/Downloads/go #gosetup
GOPATH=/home/stavros/go #gosetup
/home/stavros/Downloads/go/bin/go test -c -o /tmp/___Test_inspect_in_inspect_test_go github.com/lightbend/cloudflow/test-docker-apis/skopeoinspect #gosetup
/home/stavros/Downloads/go/bin/go tool test2json -t /tmp/___Test_inspect_in_inspect_test_go -test.v -test.run ^Test_inspect$ #gosetup
=== RUN   Test_inspect
{
    "Name": "docker.io/lightbend/spark-aggregation",
    "Digest": "sha256:c8168ed0cfdb0329a601c1f43b2b457ccaae91f132a0a7950b9763fdc0016af9",
    "RepoTags": [
        "123-7784f48-dirty",
        "134-d0ec286-dirty"
    ],
    "Created": "2020-02-24T18:12:49.4312664Z",
    "DockerVersion": "19.03.5",
    "Labels": {
        "com.lightbend.cloudflow.streamlet-descriptors": "eJztVl1v2jAU/SuWnwG10K0V0h5Kt1X91NZ2WrtRoUtyCS6OndpOgCH++64TCPSLsm6T+tC3OD73w8fn3GTCIRHVDI0VWvEm3+QVbp1BiCW6aog2MCJx2lje/Dnh4JwR3dShX15XeCDB2o6CGCk0ACPHNYgigxFQSG0PpNwtl3taWS3xE62tpSqBVj0RdRIwFO7QzFLOSxbtEE4o6qQoPysklG8y6GMMvDnhPaEiNIkRytGmPKx3zWi/9aURbLfkhaiPT3YP1M5V2jrZG4V4e7wVX/Wj00uRHXygND1tYvBxkBlN6ztnCcFBbX4CDP15zh04u1SeT9rcjRNs82abGwy0Cdu80s7z5O8eCS8BNoGgQC3q5Zs9gTIkHB16KZV1YNyF8KvKUlWpVdTm08oydChUqIcfUwOeyOfxkEW+vccDQp12Jd4PcdqBfDpoVuV6yqdTulYJXZSzK9apm10pLUyqnMg5JzbMgKjNtExj7MSadnIQ1X2B8s7yy9hHhca/O1CrhDch3fUgla6TgUx9zncb/L4Yi4yWJWiYpWcVMqdZYnRAiWsEH+CYYIUKbJVg1QLmDwVShDlNHU9RrmLXqPPpSsk/SVzpBXq1wgyDeLSxiYetX93h0fm3ow1lx/Xox+Dzzii5aVyefr25qav3vWOzfbu1nhkWxP6JBxZRL5V+atHc1RcNKfFQx9r11wOGghp9qNonwGv7yEuZbBonK83w3zSfj5fFyF1b7JssFooKPZD893yOsDkBjBTCiGIW64x4YkAfDoiQBTpOUpdDShtERqdJtTuuFrPoUQ/M83oW12/qmKYpwwyJKjakZxKtGZQ9NlmfOvaMe3cOQbi8a7kIKlssg59p7h98lt6c+Cqd+FfT9e1X41X+aqwzXa+nvwHWwMoO"
    },
    "Architecture": "amd64",
    "Os": "linux",
    "Layers": [
        "sha256:5667fdb72017d1fb364744ca1abf7b6f3bbe9c98c3786f294a461c2866db69ab",
        "sha256:d83811f270d56d34a208f721f3dbf1b9242d1900ad8981fc7071339681998a31",
        "sha256:ee671aafb583e2321880e275c94d49a49185006730e871435cd851f42d2a775d",
        "sha256:7fc152dfb3a6b5c9a436b49ff6cd72ed7eb5f1fd349128b50ee04c3c5c2355fb",
        "sha256:c98d85f69841cae2d433669ce99e3abbe09a4f44e51f15c18c6b6b250c83ff05",
        "sha256:7afb4c4d2d728a6e247d7164cc892e36d95dc6a9e7cf141a0918f6107d117cf6",
        "sha256:ab606d2c106489116bc180cc1afb0bfd2776956236a447059993edb64576d989",
        "sha256:dcb4f1eee182304799054596747cfa411ba3a6553c9163ec9de7de3907ee156d",
        "sha256:d2ab84ec7cadf6e55b5be04348487070048d59056781743799b636b2324f9fc0",
        "sha256:fe26bf4575d705e85ea12a2cd36bfed07cfe734c33979016714f2996ccffa6cc",
        "sha256:42311ce8aed78497decc15ec45b865d5f9b097d0a7ac4c5f26c6a0bd8df6513c",
        "sha256:89b356b9ad19ce7e454063ff96d8683b4f2c93f46d49a88de4fb12f455ddf1a5",
        "sha256:6ea2881383401005efab41f2bf88c365527de7f5ce4f0c6d7624b7fce656b84a",
        "sha256:acf36f63361f6d55117dccada4ab988182a965113cc52714f444484af72ffd9b",
        "sha256:d70f6e1e5480b2f950cabf8abad58c5532b767ac8cc8dddf46c38d7588a487f2",
        "sha256:3e4b832fae0fc9d52fd240462b266c830d94d6111db902b01d5253c3b9b6b6ce",
        "sha256:7ca41d142d8254d4604880d94cd92f1c1fd48935ea30b1e378fedc7d7875ef57",
        "sha256:df89d95d6ed669d21e8dfe772af9d749b2b02273aed6fe64711219c2d0b7ab31",
        "sha256:df14ceb1d52a1cfa6671eb8bc106fe5bc6ee68dd5c52f8caffadce9b170a1f42",
        "sha256:7cd9e7bfa7ca8894dd9835b08e352cdbdab5720e0d3ed90b607f364d3fc1cec2",
        "sha256:64747126c29b5a99c2574b6c873afce6d80d22f5e44a2b94353085ac5732d431"
    ],
    "Env": [
        "PATH=/opt/java/openjdk/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
        "LANG=en_US.UTF-8",
        "LANGUAGE=en_US:en",
        "LC_ALL=en_US.UTF-8",
        "JAVA_VERSION=jdk8u222-b10",
        "JAVA_HOME=/opt/java/openjdk",
        "TINI_VERSION=v0.18.0",
        "SPARK_HOME=/opt/spark",
        "SPARK_VERSION=2.4.4"
    ]
}
dockerhub inspect took 151ns
{
    "Name": "eu.gcr.io/bubbly-observer-178213/spark-aggregation",
    "Digest": "sha256:c8168ed0cfdb0329a601c1f43b2b457ccaae91f132a0a7950b9763fdc0016af9",
    "RepoTags": [
        "125-92f3631-dirty",
        "130-1d43c7a-dirty",
        "134-d0ec286-dirty",
        "173-3860d41-dirty",
        "174-8e905dd",
        "184-a48ddfe-dirty",
        "187-f3b25b5-dirty",
        "200-78fa9aa-dirty",
        "203-bb41483-dirty",
        "204-bfd11df-dirty"
    ],
    "Created": "2020-02-24T18:12:49.4312664Z",
    "DockerVersion": "19.03.5",
    "Labels": {
        "com.lightbend.cloudflow.streamlet-descriptors": "eJztVl1v2jAU/SuWnwG10K0V0h5Kt1X91NZ2WrtRoUtyCS6OndpOgCH++64TCPSLsm6T+tC3OD73w8fn3GTCIRHVDI0VWvEm3+QVbp1BiCW6aog2MCJx2lje/Dnh4JwR3dShX15XeCDB2o6CGCk0ACPHNYgigxFQSG0PpNwtl3taWS3xE62tpSqBVj0RdRIwFO7QzFLOSxbtEE4o6qQoPysklG8y6GMMvDnhPaEiNIkRytGmPKx3zWi/9aURbLfkhaiPT3YP1M5V2jrZG4V4e7wVX/Wj00uRHXygND1tYvBxkBlN6ztnCcFBbX4CDP15zh04u1SeT9rcjRNs82abGwy0Cdu80s7z5O8eCS8BNoGgQC3q5Zs9gTIkHB16KZV1YNyF8KvKUlWpVdTm08oydChUqIcfUwOeyOfxkEW+vccDQp12Jd4PcdqBfDpoVuV6yqdTulYJXZSzK9apm10pLUyqnMg5JzbMgKjNtExj7MSadnIQ1X2B8s7yy9hHhca/O1CrhDch3fUgla6TgUx9zncb/L4Yi4yWJWiYpWcVMqdZYnRAiWsEH+CYYIUKbJVg1QLmDwVShDlNHU9RrmLXqPPpSsk/SVzpBXq1wgyDeLSxiYetX93h0fm3ow1lx/Xox+Dzzii5aVyefr25qav3vWOzfbu1nhkWxP6JBxZRL5V+atHc1RcNKfFQx9r11wOGghp9qNonwGv7yEuZbBonK83w3zSfj5fFyF1b7JssFooKPZD893yOsDkBjBTCiGIW64x4YkAfDoiQBTpOUpdDShtERqdJtTuuFrPoUQ/M83oW12/qmKYpwwyJKjakZxKtGZQ9NlmfOvaMe3cOQbi8a7kIKlssg59p7h98lt6c+Cqd+FfT9e1X41X+aqwzXa+nvwHWwMoO"
    },
    "Architecture": "amd64",
    "Os": "linux",
    "Layers": [
        "sha256:5667fdb72017d1fb364744ca1abf7b6f3bbe9c98c3786f294a461c2866db69ab",
        "sha256:d83811f270d56d34a208f721f3dbf1b9242d1900ad8981fc7071339681998a31",
        "sha256:ee671aafb583e2321880e275c94d49a49185006730e871435cd851f42d2a775d",
        "sha256:7fc152dfb3a6b5c9a436b49ff6cd72ed7eb5f1fd349128b50ee04c3c5c2355fb",
        "sha256:c98d85f69841cae2d433669ce99e3abbe09a4f44e51f15c18c6b6b250c83ff05",
        "sha256:7afb4c4d2d728a6e247d7164cc892e36d95dc6a9e7cf141a0918f6107d117cf6",
        "sha256:ab606d2c106489116bc180cc1afb0bfd2776956236a447059993edb64576d989",
        "sha256:dcb4f1eee182304799054596747cfa411ba3a6553c9163ec9de7de3907ee156d",
        "sha256:d2ab84ec7cadf6e55b5be04348487070048d59056781743799b636b2324f9fc0",
        "sha256:fe26bf4575d705e85ea12a2cd36bfed07cfe734c33979016714f2996ccffa6cc",
        "sha256:42311ce8aed78497decc15ec45b865d5f9b097d0a7ac4c5f26c6a0bd8df6513c",
        "sha256:89b356b9ad19ce7e454063ff96d8683b4f2c93f46d49a88de4fb12f455ddf1a5",
        "sha256:6ea2881383401005efab41f2bf88c365527de7f5ce4f0c6d7624b7fce656b84a",
        "sha256:acf36f63361f6d55117dccada4ab988182a965113cc52714f444484af72ffd9b",
        "sha256:d70f6e1e5480b2f950cabf8abad58c5532b767ac8cc8dddf46c38d7588a487f2",
        "sha256:3e4b832fae0fc9d52fd240462b266c830d94d6111db902b01d5253c3b9b6b6ce",
        "sha256:7ca41d142d8254d4604880d94cd92f1c1fd48935ea30b1e378fedc7d7875ef57",
        "sha256:df89d95d6ed669d21e8dfe772af9d749b2b02273aed6fe64711219c2d0b7ab31",
        "sha256:df14ceb1d52a1cfa6671eb8bc106fe5bc6ee68dd5c52f8caffadce9b170a1f42",
        "sha256:7cd9e7bfa7ca8894dd9835b08e352cdbdab5720e0d3ed90b607f364d3fc1cec2",
        "sha256:64747126c29b5a99c2574b6c873afce6d80d22f5e44a2b94353085ac5732d431"
    ],
    "Env": [
        "PATH=/opt/java/openjdk/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
        "LANG=en_US.UTF-8",
        "LANGUAGE=en_US:en",
        "LC_ALL=en_US.UTF-8",
        "JAVA_VERSION=jdk8u222-b10",
        "JAVA_HOME=/opt/java/openjdk",
        "TINI_VERSION=v0.18.0",
        "SPARK_HOME=/opt/spark",
        "SPARK_VERSION=2.4.4"
    ]
}
GCloud inspect took 349ns
{
    "Name": "405074236871.dkr.ecr.eu-west-1.amazonaws.com/stavros-test/sensor-data-scala",
    "Digest": "sha256:3e42ffae0ba8f334e40b9754758668bdd0ed08a1ed63b7e831998b65b701599a",
    "RepoTags": [
        "96-7fe7841-dirty",
        "90-d662d87-dirty",
        "122-e847372-dirty"
    ],
    "Created": "2020-01-16T12:05:00.781233736Z",
    "DockerVersion": "19.03.4",
    "Labels": {
        "com.lightbend.cloudflow.application.zlib": "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"
    },
    "Architecture": "amd64",
    "Os": "linux",
    "Layers": [
        "sha256:5667fdb72017d1fb364744ca1abf7b6f3bbe9c98c3786f294a461c2866db69ab",
        "sha256:d83811f270d56d34a208f721f3dbf1b9242d1900ad8981fc7071339681998a31",
        "sha256:ee671aafb583e2321880e275c94d49a49185006730e871435cd851f42d2a775d",
        "sha256:7fc152dfb3a6b5c9a436b49ff6cd72ed7eb5f1fd349128b50ee04c3c5c2355fb",
        "sha256:c98d85f69841cae2d433669ce99e3abbe09a4f44e51f15c18c6b6b250c83ff05",
        "sha256:7afb4c4d2d728a6e247d7164cc892e36d95dc6a9e7cf141a0918f6107d117cf6",
        "sha256:180b4f90fc041fd34d48e164a6d4e42c2aa3252c134d15481cc83d8676881d23",
        "sha256:cc860c6cba5ab6b95e21d755806df7d0adad64a163ccc88c81be4460d1c4e885",
        "sha256:a3c2a80eb11d2f6e74099291fce33726d866ee737c207c0b3947130c8402df90",
        "sha256:73ef3cb7d16327cd29c97e9476ed91c68dc374656edb254e0f32e80178ca5047",
        "sha256:fa7eaa2ca9a5c8e44e6571d29e3e91b8adb42d25c4ae2bfb8d6759aca216136a",
        "sha256:127f288745c60e7264ecbab16eaa5a0881e335de2ad267c48057c5d305e505b5",
        "sha256:60ef79d344c8d0f0090132bc0ef70db0e4f3dc5725e319914867bd231586ac57",
        "sha256:9e425ce1ed79d0ea977fbfcf50b4f013806eec327e80e850355764baa39748da",
        "sha256:2fb1a3a7891e73fa2ae0a7ee570c35b53fc306e3f0690f72543f247c97e40516",
        "sha256:c764dc8b88994ee48c518c525bf1ef9e5dc9c86093361a97e23cfdba585fc717",
        "sha256:fdfd48b843c73333d4396ca5978e13516cd940fdc6ae440478bd30389e1738c3",
        "sha256:048ce485d607e8ace25f99e7c0589516324f988fa0c2ef5817f399576de8098c",
        "sha256:fbb68404e86b373ff0378a6ad4d8051320e2ed11b760565dd6a52e6a05a3c865",
        "sha256:10aba84a265f49e2288dc4d6306156e3f7ab1375d871f89f8815a674b4079e5e",
        "sha256:1b6bdf1d7355b2090f438670c0d2b6cada645e45eea41f43ff140f7b4c8935d0",
        "sha256:5466de3b58a81eb670a7ee5668d4b2936715e18e89e47177423c70109ad1882d",
        "sha256:7997744eb99b37544bc3cc1deaf83f7cab9f1095db299dcfd51723752743fee9",
        "sha256:813f2aa3f2f28be7e87a9a56b8d98189abdd99c9c4030804e7248857d55d83a9",
        "sha256:6cf29ded5163f3e7866ce2643135ca5feebee518a6bce411b1bca5e40fae2da1",
        "sha256:1873dc69e5b679f3628b46d4d94e82079be1558a2beadad72d5d8b392a5879aa",
        "sha256:63f7279048e00004e0bb5c0578d2925383308693a3fc1dde016d82d149c26c69",
        "sha256:1c7393d26c5d2008bd610c4dbf5126ac034545dc6375890cb1c0edfeb03dbccd",
        "sha256:b5062a374ace8cd817baee030bf5bb649327323b279b4297e364130346056657",
        "sha256:f0b2d24f09729fc7b562d465dafba4c7ab5c65bfd7b19c147756f7c1e7b8bd89",
        "sha256:fea8e8c67e371cd1cb698966e85f1b1298dd0f684d50621fbe6ddf03b6e2cc32"
    ],
    "Env": [
        "PATH=/opt/java/openjdk/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
        "LANG=en_US.UTF-8",
        "LANGUAGE=en_US:en",
        "LC_ALL=en_US.UTF-8",
        "JAVA_VERSION=jdk8u222-b10",
        "JAVA_HOME=/opt/java/openjdk",
        "TINI_VERSION=v0.18.0",
        "SPARK_HOME=/opt/spark",
        "FLINK_VERSION=1.9.1",
        "SPARK_VERSION=2.4.4",
        "SCALA_VERSION=2.12",
        "FLINK_HOME=/opt/flink",
        "FLINK_TGZ=flink-1.9.1-bin-scala_2.12.tgz",
        "FLINK_URL_FILE_PATH=flink/flink-1.9.1/flink-1.9.1-bin-scala_2.12.tgz",
        "FLINK_TGZ_URL=https://mirrors.ocf.berkeley.edu/apache/flink/flink-1.9.1/flink-1.9.1-bin-scala_2.12.tgz"
    ]
}
AWS inspect took 295ns
{
    "Name": "docker.io/lightbend/spark-aggregation",
    "Digest": "sha256:d97920e8e8678854cbbbf27c9330b48f0b290156e69003bd734de6f93899c567",
    "RepoTags": [],
    "Created": "2020-02-24T18:12:49.4312664Z",
    "DockerVersion": "19.03.5",
    "Labels": {
        "com.lightbend.cloudflow.streamlet-descriptors": "eJztVl1v2jAU/SuWnwG10K0V0h5Kt1X91NZ2WrtRoUtyCS6OndpOgCH++64TCPSLsm6T+tC3OD73w8fn3GTCIRHVDI0VWvEm3+QVbp1BiCW6aog2MCJx2lje/Dnh4JwR3dShX15XeCDB2o6CGCk0ACPHNYgigxFQSG0PpNwtl3taWS3xE62tpSqBVj0RdRIwFO7QzFLOSxbtEE4o6qQoPysklG8y6GMMvDnhPaEiNIkRytGmPKx3zWi/9aURbLfkhaiPT3YP1M5V2jrZG4V4e7wVX/Wj00uRHXygND1tYvBxkBlN6ztnCcFBbX4CDP15zh04u1SeT9rcjRNs82abGwy0Cdu80s7z5O8eCS8BNoGgQC3q5Zs9gTIkHB16KZV1YNyF8KvKUlWpVdTm08oydChUqIcfUwOeyOfxkEW+vccDQp12Jd4PcdqBfDpoVuV6yqdTulYJXZSzK9apm10pLUyqnMg5JzbMgKjNtExj7MSadnIQ1X2B8s7yy9hHhca/O1CrhDch3fUgla6TgUx9zncb/L4Yi4yWJWiYpWcVMqdZYnRAiWsEH+CYYIUKbJVg1QLmDwVShDlNHU9RrmLXqPPpSsk/SVzpBXq1wgyDeLSxiYetX93h0fm3ow1lx/Xox+Dzzii5aVyefr25qav3vWOzfbu1nhkWxP6JBxZRL5V+atHc1RcNKfFQx9r11wOGghp9qNonwGv7yEuZbBonK83w3zSfj5fFyF1b7JssFooKPZD893yOsDkBjBTCiGIW64x4YkAfDoiQBTpOUpdDShtERqdJtTuuFrPoUQ/M83oW12/qmKYpwwyJKjakZxKtGZQ9NlmfOvaMe3cOQbi8a7kIKlssg59p7h98lt6c+Cqd+FfT9e1X41X+aqwzXa+nvwHWwMoO"
    },
    "Architecture": "amd64",
    "Os": "linux",
    "Layers": [
        "sha256:a1aa3da2a80a775df55e880b094a1a8de19b919435ad0c71c29a0983d64e65db",
        "sha256:ef1a1ec5bba9f5efcecf38693111c335cafa27f53669a91bee5d3dc17819180c",
        "sha256:6c3332381368f5c277995c2e1d19dc895b8a870ba7d1ccd8a4dbe4a5c26810bc",
        "sha256:e80c789bc6aca44ee043fb65d06ddff70f644086dd99e6c65b04149cd5787d84",
        "sha256:3764a40e8cad782996e55b052776c57eefd5edf3bc09e1086cf484596003f462",
        "sha256:3d36dd6d1b284f9436e2a3993f7163f9501694ed738425b4634bd2f330855f41",
        "sha256:a132357083b5f801417f82fb6fe55f26d2bed0581c73a28808acfee526de8026",
        "sha256:9fa50fe641f6f7efb179635732ed0b423260d41b96264fd5a9b63d84e60f42d3",
        "sha256:a67280d18980f1de58e65ea460aad37c5f2d4f12043c2ece14bf4828cfccb4fc",
        "sha256:c6de247f7d4cf8379b7c00934c34df9bcbcf4571d8dce6ce2ec27224597fbfc1",
        "sha256:c19c1728db9ede413d34b294591cb0c2fb42c4e106be6d167e02aa6cefad6d6a",
        "sha256:704bc4aed94fe5e59ac6723b577d61b33676b325736698a44b0ace3e353c13f4",
        "sha256:98a8c6b66842eb37925cf77a6689b20a6a63ac26707227787093eb918ea85f97",
        "sha256:2531b3d491612e49ca92e2146eec4c09b2b6bb8552555cce9e56034750016cee",
        "sha256:89627b3864af8e6d4b8ff947cea734cdf13f1b4c3aa64b7d6f5b2c27fad79d12",
        "sha256:eb6edb7d311a3ee815bb034439de9fd03527c8a82e7fbefeeadc70a2ff94373d",
        "sha256:0997835d3a3048c023401cb8d475ab8c61aa91168c02eae2d6173e6f4edd99ea",
        "sha256:67ca8b84b95392ceffdf318aa4df5bbe60ac7e02a2f4ea9a57b5507b2a56e1e4",
        "sha256:565abb1180221b78d1d3d72596ea4a663e275c6c277fa8254469af4d940a63ad",
        "sha256:0bdd66a9b761f5724f2e67d4ad1d3ca6c0fa484517c3624f6a30400c10aacd16",
        "sha256:d9dcb69e870d2012a5ffc85dfd23bf6cde1d27a9fcc03dcc962d88af8a243420"
    ],
    "Env": [
        "PATH=/opt/java/openjdk/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
        "LANG=en_US.UTF-8",
        "LANGUAGE=en_US:en",
        "LC_ALL=en_US.UTF-8",
        "JAVA_VERSION=jdk8u222-b10",
        "JAVA_HOME=/opt/java/openjdk",
        "TINI_VERSION=v0.18.0",
        "SPARK_HOME=/opt/spark",
        "SPARK_VERSION=2.4.4"
    ]
}
local daemon took 309ns
--- PASS: Test_inspect (30.43s)
PASS

```
