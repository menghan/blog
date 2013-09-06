---
layout: post
title:  "组内 zipkin 分享"
date:   2013-09-06 17:35:00
categories: zipkin
---

在组内做了 [zipkin][zipkin] 的分享，下面是分享摘要

# Zipkin

Menghan

@menghan412


# Outline

1. Google's Dapper

2. Twitter's Zipkin

3. Douban's ???


# Dapper

Why Google create launch Dapper?

What is Dapper like?

black-box V.S. annotation-based


# Dapper

1. Low overhead

2. Application-level transparency

3. Scalability

4. Analyse quickly


# Dapper design

1. data structure: trace, span, annotation, key-value annotation

2. instrumentation points

3. sample

4. collection: out of band


# Dapper deploy

1. Dapper daemon

2. overhead calculation

3. adaptive sampling

4. adaptive collection

5. D-API


# Dapper learned

1. A sample of just one out of thousands of requests provides sufficient information for many common uses of tracing data

2. integration with exception monitoring

3. coalescing effects

4. offline workload

5. find root cause

6. record kernel level infomation


# Zipkin

http://www.infoq.com/presentations/Zipkin


# Zipkin V.S. Dapper

1. scribe vs daemon

2. initiative vs passive

3. BigTable vs ...


# Zipkin

usage


# Douban

???

# END

[zipkin]: http://twitter.github.io/zipkin/
