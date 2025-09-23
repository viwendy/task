# task
Web3 learn task

# task1
    136. 【appearOnlyOnceNum.go】只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。
         找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，
         例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
         回文数【palindromicNum.go】

    20. 有效的括号【isValid.go】-------这个题目没搞明白 （堆栈关系）
        给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
        有效字符串需满足：
        左括号必须用相同类型的右括号闭合。
        左括号必须以正确的顺序闭合。
        每个右括号都有一个对应的相同类型的左括号。

    14. 最长公共前缀【longestCommonPrefix.go】
        编写一个函数来查找字符串数组中的最长公共前缀。
        如果不存在公共前缀，返回空字符串 ""。

    66. 加一【plusOne.go】
        给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。

    26. 删除有序数组中的重复项[removeDuplicates.go]
        给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。
        考虑 nums 的唯一元素的数量为 k ，你需要做以下事情确保你的题解可以被通过：
        更改数组 nums ，使 nums 的前 k 个元素包含唯一元素，并按照它们最初在 nums 中出现的顺序排列。nums 的其余元素与 nums 的大小不重要。
        返回 k 。

    56. 合并区间【mergeArr.go】-------这个题目需要再次打磨
        以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
        请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

# task2
    ✅指针【pointer.go】
    题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
    考察点 ：指针的使用、值传递与引用传递的区别。
    题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
    考察点 ：指针运算、切片操作。

    ✅Goroutine【goroutiner.go】【scheduler.go】
    题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
    考察点 ： go 关键字的使用、协程的并发执行。
    题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
    考察点 ：协程原理、并发任务调度。-------这个题目没搞明白

    ✅面向对象【shape.go】【person.go】
    题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，
    创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
    考察点 ：接口的定义与实现、面向对象编程风格。
    题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。
    为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
    考察点 ：组合的使用、方法接收者。

    ✅Channel【channel.go】【bufferChannel.go】
    题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
    考察点 ：通道的基本使用、协程间通信。
    题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
    考察点 ：通道的缓冲机制。

    ✅锁机制【mutex.go】【atomic.go】
    题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
    考察点 ： sync.Mutex 的使用、并发数据安全。
    题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
    考察点 ：原子操作、并发数据安全。

# task3
    题目1：基本CRUD操作【curd.sql】
    假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
    要求 ：
    编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
    编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
    编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
    编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。

    题目2：事务语句【transaction.sql】
    假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
    要求 ：
    编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。

    题目1：使用SQL扩展库进行查询【employees.go】
    假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
    要求 ：
    编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
    编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。

    题目2：实现类型安全映射【books.go】
    假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
    要求 ：
    定义一个 Book 结构体，包含与 books 表对应的字段。
    编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。

    题目1：模型定义【model.go】
    假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
    要求 ：
    使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
    编写Go代码，使用Gorm创建这些模型对应的数据库表。

    题目2：关联查询【relevantQuery.go】
    基于上述博客系统的模型定义。
    要求 ：
    编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
    编写Go代码，使用Gorm查询评论数量最多的文章信息。

    题目3：钩子函数【hook.go】
    继续使用博客系统的模型。
    要求 ：
    为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
    为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。

# task4 Solidity智能合约基础
    本次作业要求你使用 Go 语言结合 Gin 框架和 GORM 库开发一个个人博客系统的后端，实现博客文章的基本管理功能，包括文章的创建、读取、更新和删除（CRUD）操作，
    同时支持用户认证和简单的评论功能

# task5 Solidity智能合约基础
    【Voting.sol】
    ✅ 创建一个名为Voting的合约，包含以下功能：
    一个mapping来存储候选人的得票数
    一个vote函数，允许用户投票给某个候选人
    一个getVotes函数，返回某个候选人的得票数
    一个resetVotes函数，重置所有候选人的得票数

    【ReverseString.sol】
    ✅ 反转字符串 (Reverse String)
    题目描述：反转一个字符串。输入 "abcde"，输出 "edcba"

    【RomanToInt.sol】
    ✅ 用solidity实现整数转罗马数字

    【MergeArray.sol】
    ✅ 合并两个有序数组 (Merge Sorted Array)
    题目描述：将两个有序数组合并为一个有序数组。

    【BinarySearch.sol】
    二分查找 (Binary Search)
    题目描述：在一个有序数组中查找目标值。

# task6 Solidity 编程进阶
    ✅ 作业 1：ERC20 代币【Erc20Contract.sol】
    任务：参考 openzeppelin-contracts/contracts/token/ERC20/IERC20.sol实现一个简单的 ERC20 代币合约。要求：
    合约包含以下标准 ERC20 功能：
    balanceOf：查询账户余额。
    transfer：转账。
    approve 和 transferFrom：授权和代扣转账。
    使用 event 记录转账和授权操作。
    提供 mint 函数，允许合约所有者增发代币。
    提示：
    使用 mapping 存储账户余额和授权信息。
    使用 event 定义 Transfer 和 Approval 事件。
    部署到sepolia 测试网，导入到自己的钱包

    ✅ 作业2：在测试网上发行一个图文并茂的 NFT
    任务目标
    使用 Solidity 编写一个符合 ERC721 标准的 NFT 合约。
    将图文数据上传到 IPFS，生成元数据链接。
    将合约部署到以太坊测试网（如 Goerli 或 Sepolia）。
    铸造 NFT 并在测试网环境中查看。
    任务步骤
    编写 NFT 合约
    使用 OpenZeppelin 的 ERC721 库编写一个 NFT 合约。
    合约应包含以下功能：
    构造函数：设置 NFT 的名称和符号。
    mintNFT 函数：允许用户铸造 NFT，并关联元数据链接（tokenURI）。
    在 Remix IDE 中编译合约。
    准备图文数据
    准备一张图片，并将其上传到 IPFS（可以使用 Pinata 或其他工具）。
    创建一个 JSON 文件，描述 NFT 的属性（如名称、描述、图片链接等）。
    将 JSON 文件上传到 IPFS，获取元数据链接。
    JSON文件参考 https://docs.opensea.io/docs/metadata-standards
    部署合约到测试网
    在 Remix IDE 中连接 MetaMask，并确保 MetaMask 连接到 Goerli 或 Sepolia 测试网。
    部署 NFT 合约到测试网，并记录合约地址。
    铸造 NFT
    使用 mintNFT 函数铸造 NFT：
    在 recipient 字段中输入你的钱包地址。
    在 tokenURI 字段中输入元数据的 IPFS 链接。
    在 MetaMask 中确认交易。
    查看 NFT
    打开 OpenSea 测试网 或 Etherscan 测试网。
    连接你的钱包，查看你铸造的 NFT。