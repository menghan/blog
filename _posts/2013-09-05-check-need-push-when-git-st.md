---
layout: post
title:  "在 git st 时检查是否需要 push"
date:   2013-09-05 11:38:00
categories: git
---

以前常常有修改但忘了提交, 使用 git 之后, 又经常提交了但忘了 `git push`.
为了避免再犯这个错误, 我把 `git st` 改成还检查是否需要 push.

把这段代码保存成 `git-synced` 放在 `$PATH` 里, 然后在 `~/.gitconfig` 里 `[alias]` 部分加上一句 `st = synced` 即可

{% highlight bash %}
#!/usr/bin/env bash

# assume that all local branches should be pushed to 'origin'
REMOTE=origin

if ! git diff-index --quiet HEAD --; then
        git status
else
        if git remote | grep -q $REMOTE; then
                op=$(git remote show $REMOTE)
                if [[ "$op" =~ "fast-forwardable" ]]; then
                        echo you NEED a push
                elif [[ "$op" =~ "local out of date" ]]; then
                        echo you MAY NEED a push
                fi
        fi
fi
{% endhighlight %}

可以在 [这里][menghanrc] 访问我的 rc 文件.

[menghanrc]: https://github.com/menghan/menghanrc

