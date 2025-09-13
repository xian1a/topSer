#!/bin/bash

# API测试脚本

BASE_URL="http://localhost:8080"

echo "=== 测试topService API ==="

# 1. 健康检查
echo "1. 测试健康检查..."
curl -s "${BASE_URL}/health" | jq .
echo ""

# 2. 创建用户
echo "2. 创建用户..."
USER_RESPONSE=$(curl -s -X POST "${BASE_URL}/api/v1/users" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "test_user",
    "email": "test@example.com", 
    "phone": "13800138000"
  }')
echo $USER_RESPONSE | jq .
USER_ID=$(echo $USER_RESPONSE | jq -r '.data.id')
echo "用户ID: $USER_ID"
echo ""

# 3. 获取用户列表
echo "3. 获取用户列表..."
curl -s "${BASE_URL}/api/v1/users" | jq .
echo ""

# 4. 获取单个用户
echo "4. 获取单个用户..."
curl -s "${BASE_URL}/api/v1/users/${USER_ID}" | jq .
echo ""

# 5. 创建产品
echo "5. 创建产品..."
PRODUCT_RESPONSE=$(curl -s -X POST "${BASE_URL}/api/v1/products" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "测试产品",
    "description": "这是一个测试产品",
    "price": 99.99,
    "stock": 50,
    "category": "测试分类"
  }')
echo $PRODUCT_RESPONSE | jq .
PRODUCT_ID=$(echo $PRODUCT_RESPONSE | jq -r '.data.id')
echo "产品ID: $PRODUCT_ID"
echo ""

# 6. 获取产品列表
echo "6. 获取产品列表..."
curl -s "${BASE_URL}/api/v1/products" | jq .
echo ""

# 7. 更新用户
echo "7. 更新用户..."
curl -s -X PUT "${BASE_URL}/api/v1/users/${USER_ID}" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "updated_user",
    "phone": "13900139000"
  }' | jq .
echo ""

# 8. 更新产品
echo "8. 更新产品..."
curl -s -X PUT "${BASE_URL}/api/v1/products/${PRODUCT_ID}" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "更新的产品",
    "price": 199.99
  }' | jq .
echo ""

echo "=== API测试完成 ==="