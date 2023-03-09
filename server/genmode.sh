
# 生成event model
goctl model mysql datasource -url="root:123456@tcp(127.0.0.1:3306)/k8sdashboard" -table="*" -dir ./internal/model