1.你的主人是一个学习者,使用语言JAVA,GO,Python,希望学习最新前沿的技术.
2.开发一个项目时,我们需要完整的项目开发顺序,只需要表明一个顺序就好了,便于体系化的阅读代码.
3.对人工智能,后端开发技术,感兴趣.
4.开发项目时,在项目顶部注释写出开发时的心理过程,便于我阅读和学习.
5.开发项目前先构思整体的思路,然后提前写出最终的项目文件列表清单,以便你追踪进度不出错,每个文件都要列出来,然后你可以酌情添加或减少.
6.这是aws es2实例,只有2g内存,注意cpu密集的任务限制内存.
7.python包管理器使用的是uv.

# AI Programming Optimization - Level 2 Compression Strategy

## Core Principles
- Target: 35-45% token savings with full IDE recoverability
- Trigger: Code length > 500 chars or when AI analysis/modification needed
- Benefits: Reduce AI usage costs, overcome context window limits

## Level 2 Compression Rules

### ✅ Compressible Elements
- Remove formatting spaces: `func(a, b)` → `func(a,b)`
- Delete line breaks and indentation: single-line code blocks
- Remove unnecessary empty lines

### ❌ Must Preserve Elements  
- Variable names and function names (keep semantics)
- Critical business logic comments
- String content unchanged

## Applicable Languages
- ✅ JavaScript/TypeScript: Suitable for compression
- ✅ Go: Suitable for compression  
- ✅ Java/C#: Suitable for compression
- ❌ Python: Indentation-sensitive syntax, NOT suitable
- ❌ YAML/Dockerfile: Format-sensitive, NOT suitable

## Workflow
1. Pre-AI interaction: Apply Level 2 compression automatically
2. Post-AI interaction: Use IDE formatting to recover (Prettier/IntelliJ)
3. Validate: Ensure logic integrity

## Example Transformation
```javascript
// Original (standard format)
function getUserTotal(orders, discount) {
    let total = 0;
    for (let i = 0; i < orders.length; i++) {
        total += orders[i].price;
    }
    return total * (1 - discount);
}

// Level 2 compressed (35% token savings)
function getUserTotal(orders,discount){let total=0;for(let i=0;i<orders.length;i++){total+=orders[i].price;}return total*(1-discount);}
```

## Application Scenarios
- ✅ Utility functions and helpers
- ✅ Config files and API routes
- ❌ Complex core business logic
- ❌ Frequently debugged code