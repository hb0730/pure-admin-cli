# pure admin cli

用于快速构建基于vue-pure-admin项目的命令行工具

# 方式

* 命令式
* 交互式

## 命令式

```base
./pure init
```

### example

```base
./pure init -t thin -r gitee -v last -p ./ -f vue-test
```

### Flag Options:

- `-t`,`--template`: 需要克隆的模板: `thin`,`i18n`,`tauri`,`electron`,`admin`,默认:`thin`
- `-r`,`--repo`: 需要克隆的地址: `gitee`,`github`,默认: `github`
- `-v`,`--version`: 克隆的版本,默认`last`,如果选择的的模板是分支类型,则该命令无效
- `-p`,`--path`: 本地目录:默认:`./`
- `-f`,`--force`: 是否强制覆盖

# args
- 第[0]地位为项目名称,必须


## 交互式

```base
./pure new
```


## 查看帮助
```base
./pure -h
```