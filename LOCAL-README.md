# 速率脚本测试
for i in {1..105} ;do curl "http://123.57.44.210:8080/yk/project/get.json?appid=139962975";done; 

# 项目黑名单列表
    http://123.57.44.210:8080/yk/blacklist/list.json?query=appid:139962975

# 添加接口
    http://123.57.44.210:8080/yk/blacklist/add.json?appid=214748364&content=10.10.10.01

# 批量添加
    http://123.57.44.210:8080/yk/blacklist/batch_add.json?appid=214748364&content=10.10.10.02,10.10.10.03

# 查询接口 当前黑名单是否存在
    http://123.57.44.210:8080/yk/blacklist/check.json?appid=214748364&content=10.10.10.01

# 批量查询接口
    http://127.0.0.1:8080/yk/blacklist/batch_check.json?appid=214748364&content=10.10.10.02,10.10.10.03,10.10.10.01

# 删除接口 删除指定应用，指定黑名单信息
    http://127.0.0.1:8080/yk/blacklist/delete.json?appid=214748364&content=10.10.10.01

