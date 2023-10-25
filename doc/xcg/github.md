# git本地配置正确的name和email才能提交到github仓库
- git config --global user.name xcgsnowdrop
- git config --global user.email 576478492@qq.com

# git配置

## 查看全局配置项
- 查看所有配置项：git config --global --list
- 查看指定配置项：git config --global <配置项名>

## 新增全局配置项
- git config --global <配置项名> <配置项值>
- git config --global user.name xcgsnowdrop
- git config --global user.email 576478492@qq.com

## 删除全局配置项
- git config --global --unset <配置项名>
- git config --global --unset http.sslverify

# github的SSH访问设置
- 本地生成密钥对：ssh-keygen -t rsa -C "576478492@qq.com"
- 拷贝公钥到github的 Settings==>SSH and GPG keys==> New SSH key

# git rebase master 与 git rebase origin/master的区别
- git rebase master是基于本地的master分支执行变基操作，通常在执行该操作之前，需要将本地master分支更新到最新
- git rebase origin/master是基于本地的远程跟踪分支origin/master执行变基操作，通常在执行该操作之前，需要将origin/master更新到最新
假设我们目前在功能分支feature/xcg上进行了提交，以下两种rebase操作等效
--------------------------
git checkout master // 切换到本地master分支
git pull origin master // 更新本地master分支到最新
git checkout feature/xcg // 切换到本地feature/xcg分支
git rebase master // 在feature/xcg分支上基于本地master分支执行变基操作
--------------------------
git fetch // 更新所有本地的远程跟踪分支，包括origin/master
git rebase origin/master // 在feature/xcg分支上基于origin/master分支执行变基操作
--------------------------