package dao

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/bilibili/kratos/pkg/log"

	"github.com/jinzhu/gorm"
)

// 此文件为了解决， 分层情况下使多个service公用一个dao
// 例如： 添加用户时需要进行 insert(user), 添加其他时需要写很多 insert(other)
// find , delete , update 同理，在service层调用这些公用的方法 不用重复去写dao逻辑， 减少重复代码
// 使用这些方法需要在 service层 来构建所需要的model结构， 来满足增删改查需要
// 当然如果准备直接在 service层写对数据库的增删改查代码，请忽略此文件
// 因为相对来讲orm封装本身就够简单，此文件解决的是（分层下）service去调用dao的时候，减少dao的代码量

// insert
// args : 表名，数据（表对应的结构体）
// return : interface{}, err
func (d *Dao) Insert(ctx context.Context, tableName string, model interface{}, gOrmDb ...*gorm.DB) (interface{}, error) {
	args, _ := json.Marshal(model)
	t := DbTrace(ctx, create+" "+tableName, 0, "添加 "+tableName+" 表数据"+" 参数为："+string(args))
	db := d.db[0]
	if len(gOrmDb) != 0 {
		db = gOrmDb[0]
	}
	if err := db.Table(tableName).Create(model).Error; err != nil {
		if t != nil {
			t.Finish(&err)
		}
		log.Error("添加数据失败 , %v", err)
		return nil, err
	}
	if t != nil {
		t.Finish(nil)
	}
	return model, nil
}

// find
// args : 表名， 查询参数 （表对应的结构体）
// return : *gorm.DB（为了继续构造查询）， trace（进行dao层面链路跟踪 trace执行完毕应执行 t.Finish）
// 统一的查询， 只构造查询参数， 返回对应的db回去， 可以进行下一步的db构建，分页、排序、Or、数据填充进结构体等
// 分页示例：
// if page > 0 && pageSize > 0 {
//    Db = Db.Limit(pageSize).Offset((page - 1) * pageSize)
// }
func (d *Dao) Find(ctx context.Context, tableName string, args interface{}, gOrmDb ...*gorm.DB) *gorm.DB {
	argsBytes, _ := json.Marshal(args)
	//dbId := GetDbId(len(d.db))
	db := d.db[0]
	if len(gOrmDb) != 0 {
		db = gOrmDb[0]
	}
	t := DbTrace(ctx, find+" "+tableName, dbId, "查询 "+tableName+" 表数据"+" 原始构建参数为："+string(argsBytes))
	if t != nil {
		t.Finish(nil)
	}
	return db.Table(tableName).Where(args)
}

// update
// args : tableName  (表名) ,whereArgs （修改的检索条件）（表对应的结构体）， updateArgs  (要修改的参数的结构)（表对应的结构体）
// return : err, int64(影响行数)
// 修改 先查询出来要修改的数据 ， 再进行修改
func (d *Dao) Update(ctx context.Context, tableName string, whereArgs interface{}, updateArgs interface{}, gOrmDb ...*gorm.DB) (error, int64) {
	whereArgsBytes, _ := json.Marshal(whereArgs)
	updateArgsBytes, _ := json.Marshal(updateArgs)
	db := d.db[0]
	if len(gOrmDb) != 0 {
		db = gOrmDb[0]
	}
	t := DbTrace(ctx, update+" "+tableName, 0, "修改 "+tableName+" 表数据"+";   查询参数为："+string(whereArgsBytes)+";  修改参数为："+string(updateArgsBytes))
	u := db.Table(tableName).Where(whereArgs).Updates(updateArgs)
	if t != nil {
		t.Finish(nil)
	}
	return u.Error, u.RowsAffected
}

// delete
// args : tableName  (表名) ,args （删除条件）（表对应的结构体）
// return : err, int64(影响行数)
// 删除时如果删除 id为空， 删除条件也未找到 则会删除所有数据
// 删除前可以通过查询 来判断删除数据存在不存在， 但是 查询过程会有性能损耗 。故在此强制删除时id不能为空
func (d *Dao) Delete(ctx context.Context, tableName string, args interface{}, gOrmDb ...*gorm.DB) (error, int64) {
	db := d.db[0]
	if len(gOrmDb) != 0 {
		db = gOrmDb[0]
	}
	argsBytes, err := json.Marshal(args)
	t := DbTrace(ctx, delete+" "+tableName, 0, "删除 "+tableName+" 表数据"+" 参数为："+string(argsBytes))
	if err != nil {
		if t != nil {
			t.Finish(&err)
		}
		return err, 0
	}
	var dat map[string]interface{}
	if err = json.Unmarshal(argsBytes, &dat); err != nil {
		if t != nil {
			t.Finish(&err)
		}
		return err, 0
	}
	if dat["ID"] == nil || dat["ID"].(float64) == 0 {
		err = fmt.Errorf("ID Is Nil")
		if t != nil {
			t.Finish(&err)
		}
		return err, 0
	}
	if t != nil {
		t.Finish(nil)
	}
	db = db.Table(tableName).Delete(args)
	return db.Error, db.RowsAffected
}

func (d *Dao) GetDb() *gorm.DB {
	return d.db[0]
}
