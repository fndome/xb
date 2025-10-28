# xb Vector Database - Quick Start

**5 分钟快速上手 sqlxb 向量数据库支持**

---

## 🚀 快速开始

### 1. 定义模型

```go
package main

import (
    "fmt"
    "time"
    "github.com/x-ream/xb"
)

type CodeVector struct {
    Id        int64        `db:"id"`
    Content   string       `db:"content"`
    Embedding sqlxb.Vector `db:"embedding"`  // ⭐ 向量字段
    Language  string       `db:"language"`
    Layer     string       `db:"layer"`
    CreatedAt time.Time    `db:"created_at"`
}

func (CodeVector) TableName() string {
    return "code_vectors"
}
```

---

### 2. 基础向量检索

```go

queryVector := sqlxb.Vector{0.1, 0.2, 0.3, 0.4, 0.5}

sql, args := sqlxb.Of(&CodeVector{}).
    VectorSearch("embedding", queryVector, 10).
    Build().
    SqlOfVectorSearch()

fmt.Println(sql)
// 输出:
// SELECT *, embedding <-> ? AS distance 
// FROM code_vectors 
// ORDER BY distance 
// LIMIT 10
```

---

### 3. 向量 + 标量过滤

```go
sql, args := sqlxb.Of(&CodeVector{}).
    Eq("language", "golang").        // 标量过滤
    Eq("layer", "repository").       // 标量过滤
    VectorSearch("embedding", queryVector, 10).
    Build().
    SqlOfVectorSearch()

// 输出:
// SELECT *, embedding <-> ? AS distance 
// FROM code_vectors 
// WHERE language = ? AND layer = ?
// ORDER BY distance 
// LIMIT 10
```

---

### 4. 使用不同的距离度量

```go
// 余弦距离（默认）
sql, args := sqlxb.Of(&CodeVector{}).
    VectorSearch("embedding", queryVector, 10).
    Build().
    SqlOfVectorSearch()

// L2 距离（欧氏距离）
sql, args := sqlxb.Of(&CodeVector{}).
    VectorSearch("embedding", queryVector, 10).
    VectorDistance(sqlxb.L2Distance).
    Build().
    SqlOfVectorSearch()

// 内积
sql, args := sqlxb.Of(&CodeVector{}).
    VectorSearch("embedding", queryVector, 10).
    VectorDistance(sqlxb.InnerProduct).
    Build().
    SqlOfVectorSearch()
```

---

### 5. 距离阈值过滤

```go
// 只返回距离 < 0.3 的结果
sql, args := sqlxb.Of(&CodeVector{}).
    Eq("language", "golang").
    VectorDistanceFilter("embedding", queryVector, "<", 0.3).
    Build().
    SqlOfVectorSearch()

// 输出:
// SELECT *, embedding <-> ? AS distance
// FROM code_vectors
// WHERE language = ? 
//   AND (embedding <-> ?) < ?
// ORDER BY distance
```

---

### 6. 动态查询（自动忽略 nil）

```go
// 完美利用 sqlxb 的自动忽略特性
func SearchCode(filter SearchFilter) {
    sql, args := sqlxb.Of(&CodeVector{}).
        Eq("language", filter.Language).  // nil? 忽略
        Eq("layer", filter.Layer).        // nil? 忽略
        In("tags", filter.Tags).          // empty? 忽略
        VectorSearch("embedding", filter.Vector, filter.TopK).
        Build().
        SqlOfVectorSearch()
    
    // 无需任何 if 判断！
}
```

---

### 7. 向量插入

```go
code := &CodeVector{
    Content:   "func main() { ... }",
    Embedding: sqlxb.Vector{0.1, 0.2, 0.3},
    Language:  "golang",
    Layer:     "main",
}

sql, args := sqlxb.Of(code).
    Insert(func(ib *sqlxb.InsertBuilder) {
        ib.Set("content", code.Content).
            Set("embedding", code.Embedding).
            Set("language", code.Language).
            Set("layer", code.Layer)
    }).
    Build().
    SqlOfInsert()
```

---

### 8. 向量距离计算

```go
vec1 := sqlxb.Vector{1.0, 0.0, 0.0}
vec2 := sqlxb.Vector{0.0, 1.0, 0.0}

// 余弦距离
dist := vec1.Distance(vec2, sqlxb.CosineDistance)
fmt.Printf("余弦距离: %.4f\n", dist)  // 1.0000

// L2 距离
dist = vec1.Distance(vec2, sqlxb.L2Distance)
fmt.Printf("L2 距离: %.4f\n", dist)  // 1.4142

// 向量归一化
vec := sqlxb.Vector{3.0, 4.0}
normalized := vec.Normalize()
fmt.Println(normalized)  // [0.6, 0.8]
```

---

## 📖 完整示例

### 代码搜索系统

```go

// Repository 层
type CodeVectorRepo struct {
    db *sqlx.DB
}

func (r *CodeVectorRepo) SearchSimilar(
    queryVector sqlxb.Vector,
    language string,
    layer string,
    topK int,
) ([]*CodeVector, error) {
    
    // 使用 sqlxb 构建查询
    sql, args := sqlxb.Of(&CodeVector{}).
        Eq("language", language).
        Eq("layer", layer).
        VectorSearch("embedding", queryVector, topK).
        Build().
        SqlOfVectorSearch()
    
    // 执行查询
    var results []*CodeVector
    err := r.db.Select(&results, sql, args...)
    
    return results, err
}

// Service 层
type CodeSearchService struct {
    repo *CodeVectorRepo
}

func (s *CodeSearchService) SearchCode(query string, filter SearchFilter) ([]*CodeVector, error) {
    // 1. 生成查询向量（实际应用中使用 embedding 模型）
    queryVector := generateEmbedding(query)
    
    // 2. 向量检索
    results, err := s.repo.SearchSimilar(
        queryVector,
        filter.Language,
        filter.Layer,
        filter.TopK,
    )
    
    return results, err
}
```

---

## 🎯 API 对比

### MySQL vs VectorDB - API 完全一致！

```go
// MySQL 查询（现有）
sqlxb.Of(&Order{}).
    Eq("status", 1).
    Gt("amount", 1000).
    Build().
    SqlOfSelect()

// 向量检索（新增）- 完全相同的 API！
sqlxb.Of(&CodeVector{}).
    Eq("language", "golang").
    Gt("created_at", yesterday).
    VectorSearch("embedding", queryVector, 10).
    Build().
    SqlOfVectorSearch()
```

**学习成本**: **0 天**（会用 MySQL 就会用向量数据库）✅

---

## 📚 更多示例

查看完整示例代码：[vector_test.go](../vector_test.go) 和 [qdrant_x_test.go](../qdrant_x_test.go)

---

## 📖 深入学习

- **[向量多样性与 Qdrant](./VECTOR_DIVERSITY_QDRANT.md)** - Qdrant 使用指南
- **[为什么选择 Qdrant](./WHY_QDRANT.md)** - Qdrant vs LanceDB
- **[QdrantX 使用指南](./QDRANT_X_USAGE.md)** - 高级 Qdrant API

---

**5 分钟上手，终身受用！** 🚀

