# 关于sqlc.CachedConn.QueryRowIndexCtx()的用法说明

## 实际用途说明
以`query := fmt.Sprintf("select %s from %s where `mobile` = ? and del_state = ? limit 1", userRows, m.table)`为例
1. 先执行该查询语句的查询流程，即：查缓存，若没有缓存则查mysql数据库，并安装唯一索引key缓存数据，这里指："cache:looklookUsercenter:user:mobile:15618918500"
2. 查询到数据后，将该数据安装主键索引key再次缓存一遍该数据，这里指："cache:looklookUsercenter:user:id:1"

## 调用栈信息
cache.cacheNode.doTake.func1 (c:\Users\BFChainer\go\pkg\mod\github.com\zeromicro\go-zero@v1.5.3\core\stores\cache\cachenode.go:212)
syncx.(*flightGroup).makeCall (c:\Users\BFChainer\go\pkg\mod\github.com\zeromicro\go-zero@v1.5.3\core\syncx\singleflight.go:80)
syncx.(*flightGroup).DoEx (c:\Users\BFChainer\go\pkg\mod\github.com\zeromicro\go-zero@v1.5.3\core\syncx\singleflight.go:52)
cache.cacheNode.doTake (c:\Users\BFChainer\go\pkg\mod\github.com\zeromicro\go-zero@v1.5.3\core\stores\cache\cachenode.go:211)
cache.cacheNode.TakeWithExpireCtx (c:\Users\BFChainer\go\pkg\mod\github.com\zeromicro\go-zero@v1.5.3\core\stores\cache\cachenode.go:169)
cache.(*cacheNode).TakeWithExpireCtx (Unknown Source:1)
sqlc.CachedConn.QueryRowIndexCtx (c:\Users\BFChainer\go\pkg\mod\github.com\zeromicro\go-zero@v1.5.3\core\stores\sqlc\cachedsql.go:160)
model.(*defaultUserModel).FindOneByMobile (d:\workspace\go-zero-looklook\app\usercenter\model\userModel_gen.go:116)
model.(*customUserModel).FindOneByMobile (Unknown Source:1)
main.main (d:\workspace\go-zero-looklook\testModel.go:39)
runtime.main (c:\Program Files\Go\src\runtime\proc.go:267)
runtime.goexit (c:\Program Files\Go\src\runtime\asm_amd64.s:1650)

## 调用栈信息简述
model.(*defaultUserModel).FindOneByMobile ==> sqlc.CachedConn.QueryRowIndexCtx ==> cache.cacheNode.TakeWithExpireCtx ==> cache.cacheNode.doTake ==> 


## 调用栈信息相关代码注解
``````go
func (c cacheNode) doTake(ctx context.Context, v any, key string,
	query func(v any) error, cacheVal func(v any) error) error {
	logger := logx.WithContext(ctx)
	val, fresh, err := c.barrier.DoEx(key, func() (any, error) {
		if err := c.doGetCache(ctx, key, v); err != nil {
			if err == errPlaceholder {
				return nil, c.errNotFound
			} else if err != c.errNotFound {
				// why we just return the error instead of query from db,
				// because we don't allow the disaster pass to the dbs.
				// fail fast, in case we bring down the dbs.
				return nil, err
			}

            // 这里的query函数指github.com/zeromicro/go-zero/core/stores/cache.cacheNode.TakeWithExpireCtx.func1
            // 即调用者cache.cacheNode.TakeWithExpireCtx()中调用cache.cacheNode.doTake()传的实参中的第一个匿名函数
            /**
            * 按照唯一索引去查找数据，并获取到该数据对应的主键索引字段的值，然后再次用主键索引字段值去缓存该数据
            func(val any, expire time.Duration) (err error) {
                primaryKey, err = indexQuery(ctx, cc.db, v)
                if err != nil {
                    return
                }

                found = true
                return cc.cache.SetWithExpireCtx(ctx, keyer(primaryKey), v,
                    expire+cacheSafeGapBetweenIndexAndPrimary)
            }
            **/
            // 而该参数TakeWithExpireCtx.func1调用的indexQuery()实际上指向的是model.(*defaultUserModel).FindOneByMobile()中调用m.QueryRowIndexCtx()传的实参中定义的匿名函数：
            /**
            * 按照唯一索引查询数据并返回该数据的主键索引字段的值
            func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
                query := fmt.Sprintf("select %s from %s where `mobile` = ? and del_state = ? limit 1", userRows, m.table)
                if err := conn.QueryRowCtx(ctx, &resp, query, mobile, globalkey.DelStateNo); err != nil {
                    return nil, err
                }
                return resp.Id, nil
            }
            **/            
			if err = query(v); err == c.errNotFound {
				if err = c.setCacheWithNotFound(ctx, key); err != nil {
					logger.Error(err)
				}
				return nil, c.errNotFound
			} else if err != nil {
				c.stat.IncrementDbFails()
				return nil, err
			}

            // 这里的query函数指github.com/zeromicro/go-zero/core/stores/cache.cacheNode.TakeWithExpireCtx.func1
            // 即调用者cache.cacheNode.TakeWithExpireCtx()中调用doTake传的形参params中的第二个函数
			if err = cacheVal(v); err != nil {
				logger.Error(err)
			}
		}

		return jsonx.Marshal(v)
	})
	if err != nil {
		return err
	}
	if fresh {
		return nil
	}

	// got the result from previous ongoing query.
	// why not call IncrementTotal at the beginning of this function?
	// because a shared error is returned, and we don't want to count.
	// for example, if the db is down, the query will be failed, we count
	// the shared errors with one db failure.
	c.stat.IncrementTotal()
	c.stat.IncrementHit()

	return jsonx.Unmarshal(val.([]byte), v)
}
``````go