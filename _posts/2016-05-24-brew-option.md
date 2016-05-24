---
layout: post
title:  "使用 Homebrew 安装软件时指定编译选项"
date:   2016-05-24 14:00:00
categories: homebrew
---

最近有一个 tcp ssl offloading 的事情要做。
因为公司在反向代理方面一直使用 nginx.
nginx 1.10 stable release 添加了 stream 模块支持了 tcp load balancing.
因此想拿 nginx 来做这个事情.

首先在本地做个测试。homebrew 默认安装的 nginx 没有编译 stream module

{% highlight bash %}
menghan@air:~/Downloads$ nginx -V
nginx version: nginx/1.10.0
built by clang 7.3.0 (clang-703.0.29)
built with OpenSSL 1.0.2g  1 Mar 2016 (running with OpenSSL 1.0.2h  3 May 2016)
TLS SNI support enabled
configure arguments: --prefix=/usr/local/Cellar/nginx/1.10.0 --with-http_ssl_module --with-pcre --with-ipv6 --sbin-path=/usr/local/Cellar/nginx/1.10.0/bin/nginx --with-cc-opt='-I/usr/local/Cellar/pcre/8.38/include -I/usr/local/Cellar/openssl/1.0.2g/include' --with-ld-opt='-L/usr/local/Cellar/pcre/8.38/lib -L/usr/local/Cellar/openssl/1.0.2g/lib' --conf-path=/usr/local/etc/nginx/nginx.conf --pid-path=/usr/local/var/run/nginx.pid --lock-path=/usr/local/var/run/nginx.lock --http-client-body-temp-path=/usr/local/var/run/nginx/client_body_temp --http-proxy-temp-path=/usr/local/var/run/nginx/proxy_temp --http-fastcgi-temp-path=/usr/local/var/run/nginx/fastcgi_temp --http-uwsgi-temp-path=/usr/local/var/run/nginx/uwsgi_temp --http-scgi-temp-path=/usr/local/var/run/nginx/scgi_temp --http-log-path=/usr/local/var/log/nginx/access.log --error-log-path=/usr/local/var/log/nginx/error.log --with-http_gzip_static_module
{% endhighlight %}

homebrew 也提供 nginx-full 可以支持更多的编译选项，安装了发现也没有把 stream 编译进去

{% highlight bash %}
menghan@air:~/Downloads$ nginx -V
nginx version: nginx/1.10.0
built by clang 7.3.0 (clang-703.0.31)
built with OpenSSL 1.0.2h  3 May 2016
TLS SNI support enabled
configure arguments: --prefix=/usr/local/Cellar/nginx-full/1.10.0 --with-http_ssl_module --with-pcre --with-ipv6 --sbin-path=/usr/local/Cellar/nginx-full/1.10.0/bin/nginx --with-cc-opt='-I/usr/local/include -I/usr/local/Cellar/pcre/8.38/include -I/usr/local/Cellar/openssl/1.0.2h_1/include' --with-ld-opt='-L/usr/local/lib -L/usr/local/Cellar/pcre/8.38/lib -L/usr/local/Cellar/openssl/1.0.2h_1/lib' --conf-path=/usr/local/etc/nginx/nginx.conf --pid-path=/usr/local/var/run/nginx.pid --lock-path=/usr/local/var/run/nginx.lock --http-client-body-temp-path=/usr/local/var/run/nginx/client_body_temp --http-proxy-temp-path=/usr/local/var/run/nginx/proxy_temp --http-fastcgi-temp-path=/usr/local/var/run/nginx/fastcgi_temp --http-uwsgi-temp-path=/usr/local/var/run/nginx/uwsgi_temp --http-scgi-temp-path=/usr/local/var/run/nginx/scgi_temp --http-log-path=/usr/local/var/log/nginx/access.log --error-log-path=/usr/local/var/log/nginx/error.log
{% endhighlight %}

man 了一下，发现 homebrew 支持 options 来查看可选的编译选项，将选项放到 install 之后即可

{% highlight bash %}
menghan@air:~/Downloads$ brew options homebrew/nginx/nginx-full | grep stream
--with-push-stream-module
	Compile with support for http push stream module
--with-stream
--with-upstream-order-module
	Compile with support for Order Upstream module
	Compile with support for Upstream Statistics (HAProxy style) module
{% endhighlight %}

于是

{% highlight bash %}
menghan@air:~/Downloads$ brew install homebrew/nginx/nginx-full --with-stream
==> Installing nginx-full from homebrew/nginx
==> Downloading http://nginx.org/download/nginx-1.10.0.tar.gz
Already downloaded: /Library/Caches/Homebrew/nginx-full-1.10.0.tar.gz
==> ./configure --prefix=/usr/local/Cellar/nginx-full/1.10.0 --with-http_ssl_module --with-pcre --with-ipv6 --sbin-path=/usr/local/Cellar/nginx
==> make install
==> Caveats
Docroot is: /usr/local/var/www

The default port has been set in /usr/local/etc/nginx/nginx.conf to 8080 so that
nginx can run without sudo.

nginx will load all files in /usr/local/etc/nginx/servers/.

- Tips -
Run port 80:
 $ sudo chown root:wheel /usr/local/Cellar/nginx-full/1.10.0/bin/nginx
 $ sudo chmod u+s /usr/local/Cellar/nginx-full/1.10.0/bin/nginx
Reload config:
 $ nginx -s reload
Reopen Logfile:
 $ nginx -s reopen
Stop process:
 $ nginx -s stop
Waiting on exit process
 $ nginx -s quit

To have launchd start homebrew/nginx/nginx-full now and restart at login:
  brew services start homebrew/nginx/nginx-full
Or, if you don't want/need a background service you can just run:
  nginx
==> Summary
🍺  /usr/local/Cellar/nginx-full/1.10.0: 7 files, 1006.5K, built in 35 seconds
menghan@air:~/Downloads$ nginx -V
nginx version: nginx/1.10.0
built by clang 7.3.0 (clang-703.0.31)
built with OpenSSL 1.0.2h  3 May 2016
TLS SNI support enabled
configure arguments: --prefix=/usr/local/Cellar/nginx-full/1.10.0 --with-http_ssl_module --with-pcre --with-ipv6 --sbin-path=/usr/local/Cellar/nginx-full/1.10.0/bin/nginx --with-cc-opt='-I/usr/local/include -I/usr/local/Cellar/pcre/8.38/include -I/usr/local/Cellar/openssl/1.0.2h_1/include' --with-ld-opt='-L/usr/local/lib -L/usr/local/Cellar/pcre/8.38/lib -L/usr/local/Cellar/openssl/1.0.2h_1/lib' --conf-path=/usr/local/etc/nginx/nginx.conf --pid-path=/usr/local/var/run/nginx.pid --lock-path=/usr/local/var/run/nginx.lock --http-client-body-temp-path=/usr/local/var/run/nginx/client_body_temp --http-proxy-temp-path=/usr/local/var/run/nginx/proxy_temp --http-fastcgi-temp-path=/usr/local/var/run/nginx/fastcgi_temp --http-uwsgi-temp-path=/usr/local/var/run/nginx/uwsgi_temp --http-scgi-temp-path=/usr/local/var/run/nginx/scgi_temp --http-log-path=/usr/local/var/log/nginx/access.log --error-log-path=/usr/local/var/log/nginx/error.log --with-stream
{% endhighlight %}

可以看到，nginx 已经编译进来了想要的模块
