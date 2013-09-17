---
layout: post
title:  "对 Go channel 的研究"
date:   2013-09-10 17:00:00
categories: go
---
总结一下, 向 channel 发东西的规律:

1. send 有可能阻塞, 但不会失败
2. recv 有可能阻塞, 也可能会失败(这里指panic)
3. 带 default: 的 select 的作用是当会出现阻塞或失败时, 立刻返回; 如果没有default: 则等第一个可以继续的
4. channel 的缓冲和现有容量决定了是否阻塞
