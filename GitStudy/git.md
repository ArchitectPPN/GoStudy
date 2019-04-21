### Git 操作使用说明

~~~2019-4-15 10:58:42
1. 删除本地分支	git branch -D 分支名
2. 删除远程分支	git push origin --delete 分支名		
~~~

### 回退版本

##### 1. 先显示提交的Log 	git log

~~~
$ git log -3
commit 4dc08bb8996a6ee02f
Author: Mark <xxx@xx.com>
Date:   Wed Sep 7 08:08:53 2016 +0800

    xxxxx

commit 9cac9ba76574da2167
Author: xxx<xx@qq.com>
Date:   Tue Sep 6 22:18:59 2016 +0800

    improved the requst

commit e377f60e28c8b84158
Author: xxx<xxx@qq.com>
Date:   Tue Sep 6 14:42:44 2016 +0800

    changed the password from empty to max123
~~~

##### 2. 回滚到指定版本

~~~2019-4-15 11:00:52
git reset --hard 版本号
~~~

##### 3. 强制提交

~~~2019-4-15 11:01:43
git push -f origin 分支名
~~~

### 添加忽略文件不起作用

##### 1.温柔解决

~~~2019-4-15 11:06:31
1. 示例
	frontend/controllers/XtoolsController.php
	common/components
2. 使用命令(取消文件追踪):
	git update-index --assume-unchanged <取消跟踪的文件>
	这时忽略文件就会生效了
~~~

##### 2.暴力解决(刷新git缓存)

~~~
// 清空git缓存
git rm -r --cached .
// 添加文件
git add .
git commit -a -m "modify .gitignore file"
~~~

### 添加暂存区
~~~
 1. git stash save 暂存名
 2. git stash list 显示出所有的暂存区
 2. git stash apply 暂存区名称
~~~

